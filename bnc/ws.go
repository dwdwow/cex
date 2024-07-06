package bnc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"sync"
	"time"

	"github.com/dwdwow/props"
	"github.com/gorilla/websocket"
)

type WsCfg struct {
	// ws url without streams and auth tokens
	Url          string
	ListenKeyUrl string

	// ex. spot 1024, portfolio margin / futures 200
	// normally just for public websocket
	MaxStream int

	// binance has incoming massage limitation
	// ex. spot 5/s, futures 10/s
	ReqDur       time.Duration
	MaxReqPerDur int
}

type RawWsClient struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	cfg WsCfg

	user *User

	muxAuthKey sync.Mutex
	listenKey  string

	muxConn sync.Mutex
	conn    *websocket.Conn

	fan *props.Fanout[[]byte]

	muxReqToken   sync.Mutex
	crrTokenIndex int
	latestTokens  []int64

	muxStream sync.Mutex
	stream    []string

	chRestart chan struct{}

	muxStatus sync.Mutex
	// -1: closed, 1: started
	status int

	muxReqId sync.Mutex
	reqId    int64

	logger *slog.Logger
}

func NewRawWsClient(cfg WsCfg, user *User, logger *slog.Logger) *RawWsClient {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	logger = logger.With("ws", "bnc_raw_ws", "url", cfg.Url)
	return &RawWsClient{
		cfg:          cfg,
		user:         user,
		fan:          props.NewFanout[[]byte](time.Second),
		latestTokens: make([]int64, cfg.MaxReqPerDur),
		chRestart:    make(chan struct{}),
		reqId:        1000,
		logger:       logger,
	}
}

func (w *RawWsClient) Start() {
	w.muxStatus.Lock()
	w.status = 1
	w.muxStatus.Unlock()
	go w.mainThreadStarter()
	w.chRestart <- struct{}{}
}

func (w *RawWsClient) Close() {
	w.muxStatus.Lock()
	w.status = -1
	w.muxStatus.Unlock()
	if w.ctxCancel != nil {
		w.ctxCancel()
	}
	w.muxConn.Lock()
	if w.conn != nil {
		_ = w.conn.Close()
	}
	w.muxConn.Unlock()
}

func (w *RawWsClient) mainThreadStarter() {
	for {
		_ = <-w.chRestart
		w.muxStatus.Lock()
		status := w.status
		w.muxStatus.Unlock()
		if status == -1 {
			return
		}
		err := w.start()
		if err != nil {
			w.logger.Error("Cannot Start", "err", err)
		}
	}
}

func (w *RawWsClient) start() error {
	w.logger.Info("Starting")
	if w.ctxCancel != nil {
		w.ctxCancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	w.ctx = ctx
	w.ctxCancel = cancel

	w.muxConn.Lock()
	defer w.muxConn.Unlock()

	lk, err := w.newAndKeepListenKey()

	if err != nil {
		return err
	}

	dialer := websocket.Dialer{}
	conn, resp, err := dialer.DialContext(w.ctx, w.cfg.Url+"/"+lk, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 101 {
		return fmt.Errorf("bnc_ws: response status code %d", resp.StatusCode)
	}
	if w.conn != nil {
		_ = w.conn.Close()
	}
	w.conn = conn

	go w.connListener(conn)

	return nil
}

func (w *RawWsClient) connListener(conn *websocket.Conn) {
	defer func() {
		w.chRestart <- struct{}{}
	}()
	if conn == nil {
		w.logger.Error("Nil conn")
		return
	}
	w.logger.Info("Conn listener started")
	for {
		// cannot read and write concurrently
		t, d, err := w.read()
		if err != nil {
			w.logger.Error("Read message", "err", err)
			return
		}
		switch t {
		case websocket.PingMessage:
			w.logger.Info("Server ping received", "msg", string(d))
			err = w.write(d)
			if err != nil {
				w.logger.Error("Write pong msg", "err", err)
			}
		case websocket.PongMessage:
			w.logger.Info("Server pong received", "msg", string(d))
		case websocket.TextMessage:
			w.logger.Info("Server text message received", "msg", string(d))
			w.fan.Broadcast(d)
		case websocket.BinaryMessage:
			w.logger.Info("Server binary received", "msg", string(d), "binary", d)
		case websocket.CloseMessage:
			w.logger.Info("Server closed message", "msg", string(d))
		}
	}
}

func (w *RawWsClient) write(data []byte) error {
	w.muxConn.Lock()
	defer w.muxConn.Unlock()
	if w.conn == nil {
		return errors.New("bnc_ws: nil conn")
	}
	if !w.canWriteMsg() {
		return errors.New("bnc_ws: too frequent write")
	}
	return w.conn.WriteMessage(websocket.TextMessage, data)
}

func (w *RawWsClient) read() (msgType int, data []byte, err error) {
	w.muxConn.Lock()
	defer w.muxConn.Unlock()
	if w.conn == nil {
		return -1, nil, errors.New("bnc_ws: nil conn")
	}
	return w.conn.ReadMessage()
}

func (w *RawWsClient) Sub() <-chan []byte {
	return w.fan.Sub()
}

func (w *RawWsClient) Unsub(c <-chan []byte) {
	w.fan.Unsub(c)
}

func (w *RawWsClient) SubStream(params []string) error {
	w.muxStream.Lock()
	oldStream := w.stream
	w.muxStream.Unlock()
	if len(oldStream)+len(params) > w.cfg.MaxStream {
		return fmt.Errorf("bnc_ws: too many streams, max is %d", w.cfg.MaxStream)
	}
	if !w.muxConn.TryLock() {
		return errors.New("bnc_ws: conn is busy")
	}

	err := w.conn.WriteJSON(WsSubMsg{
		Method: WsMethodSub,
		Params: params,
		Id:     "1",
	})

	w.muxConn.Unlock()

	if err != nil {
		return err
	}

	w.muxStream.Lock()
	defer w.muxStream.Unlock()
	for _, s := range params {
		if slices.Contains(w.stream, s) {
			continue
		}
		w.stream = append(w.stream, s)
	}

	return nil
}

func (w *RawWsClient) UnsubStream(params []string) error {
	if !w.muxConn.TryLock() {
		return errors.New("bnc_ws: conn is busy")
	}

	err := w.conn.WriteJSON(WsSubMsg{
		Method: WsMethodUnsub,
		Params: params,
		Id:     "1",
	})

	w.muxConn.Unlock()

	if err != nil {
		return err
	}
	w.muxStream.Lock()
	defer w.muxStream.Unlock()
	for _, s := range params {
		i := slices.Index(w.stream, s)
		if i > -1 {
			w.stream = slices.Delete(w.stream, i, i+1)
		}
	}

	return nil
}

func (w *RawWsClient) newAndKeepListenKey() (string, error) {
	w.logger.Info("Getting new listen key")
	lk, err := w.newListenKey()
	if err != nil {
		return "", err
	}
	w.logger.Info("Listen key gotten")
	w.listenKey = lk
	go w.listenKeyKeeper(w.ctx)
	return lk, nil
}

func (w *RawWsClient) newListenKey() (string, error) {
	_, lk, err := w.user.NewListenKey(w.cfg.ListenKeyUrl)
	return lk.ListenKey, err.Err
}

func (w *RawWsClient) listenKeyKeeper(ctx context.Context) {
	w.logger.Info("Listen key keeper started")
	for {
		select {
		case <-ctx.Done():
		case <-time.After(time.Minute * 20):
			w.logger.Info("Keep listening key")
			_, _, err := w.user.KeepListenKey(w.cfg.ListenKeyUrl, w.listenKey)
			if err.IsNotNil() {
				w.logger.Error("Cannot Keep Listen Key", "err", err.Error())
			}
		}
	}
}

func (w *RawWsClient) canWriteMsg() bool {
	w.muxReqToken.Lock()
	defer w.muxReqToken.Unlock()
	t := time.Now().UnixMilli()
	withinDurNum := 0
	for _, v := range w.latestTokens {
		if t-v < w.cfg.ReqDur.Milliseconds() {
			withinDurNum++
		}
	}
	maxTokenNum := len(w.latestTokens)
	if withinDurNum >= maxTokenNum {
		return false
	}
	i := w.crrTokenIndex + 1
	if i >= maxTokenNum {
		i -= maxTokenNum
	}
	w.latestTokens[i] = t
	w.crrTokenIndex = i
	return true
}

func (w *RawWsClient) Request(method WsMethod, params []any) (id string, err error) {
	w.muxReqId.Lock()
	w.reqId++
	id = fmt.Sprintf("%d", w.reqId)
	w.muxReqId.Unlock()
	d, err := json.Marshal(WsReqMsg{
		Method: method,
		Params: params,
		Id:     id,
	})
	if err != nil {
		return
	}
	err = w.write(d)
	return
}

type WsClient struct {
	wsCfg WsCfg
	rawWs *RawWsClient

	user *User

	muxFan sync.Mutex
	mfan   map[string]*props.Fanout[any]

	logger *slog.Logger
}

func NewWsClient(cfg WsCfg, user *User, logger *slog.Logger) *WsClient {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return &WsClient{
		wsCfg:  cfg,
		user:   user,
		mfan:   map[string]*props.Fanout[any]{},
		logger: logger,
	}
}

func (w *WsClient) Start() {
	go w.start()
}

func (w *WsClient) start() {
	rawWs := NewRawWsClient(w.wsCfg, w.user, w.logger)
	w.rawWs = rawWs
	ch := rawWs.Sub()
	rawWs.Start()
	for {
		w.dataHandler(<-ch)
	}
}

const mfanKeyAll = "__all__"

func (w *WsClient) dataHandler(data []byte) {
	e, ok := getWsEvent(data)
	if !ok {
		w.logger.Error("Can not get event", "data", string(data))
		return
	}
	w.muxFan.Lock()
	fan := w.mfan[string(e)]
	allFan := w.mfan[mfanKeyAll]
	w.muxFan.Unlock()
	d, err := UnmarshalSpotPrivateWsMsg(e, data)
	if err != nil {
		w.logger.Error("Can not unmarshal msg", "data", string(data))
		return
	}
	if fan != nil {
		fan.Broadcast(d)
	}
	if allFan != nil {
		allFan.Broadcast(d)
	}
}

func (w *WsClient) event2MfanKey(event WsEvent) string {
	// do not use empty string as mfan key
	if event == "" {
		event = mfanKeyAll
	}
	return string(event)
}

// Sub
// Pass empty string if you want listen all events.
func (w *WsClient) Sub(event WsEvent) <-chan any {
	w.muxFan.Lock()
	defer w.muxFan.Unlock()
	// do not use empty string as mfan key
	key := w.event2MfanKey(event)
	fan := w.mfan[key]
	if fan == nil {
		fan = props.NewFanout[any](time.Second)
		w.mfan[key] = fan
	}
	return fan.Sub()
}

func (w *WsClient) Unsub(event WsEvent, ch <-chan any) {
	w.muxFan.Lock()
	defer w.muxFan.Unlock()
	key := w.event2MfanKey(event)
	fan := w.mfan[key]
	if fan == nil {
		return
	}
	fan.Unsub(ch)
}
