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

type Ws struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	cfg WsCfg

	user *User

	muxAuthKey sync.Mutex
	listenKey  string

	muxConn sync.Mutex
	conn    *websocket.Conn

	fanout *props.Fanout[[]byte]

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

func NewWs(cfg WsCfg, user *User, logger *slog.Logger) *Ws {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	logger = logger.With("ws", "binance", "url", cfg.Url)
	return &Ws{
		cfg:          cfg,
		user:         user,
		fanout:       props.NewFanout[[]byte](time.Second),
		latestTokens: make([]int64, cfg.MaxReqPerDur),
		chRestart:    make(chan struct{}),
		reqId:        1000,
	}
}

func (w *Ws) Start() {
	w.muxStatus.Lock()
	w.status = 1
	w.muxStatus.Unlock()
	go w.mainThreadStarter()
	w.chRestart <- struct{}{}
}

func (w *Ws) Close() {
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

func (w *Ws) start() error {
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
	conn, resp, err := dialer.DialContext(w.ctx, w.cfg.Url+"?listenKey="+lk, nil)
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

func (w *Ws) mainThreadStarter() {
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

func (w *Ws) connListener(conn *websocket.Conn) {
	defer func() {
		w.chRestart <- struct{}{}
	}()
	if conn == nil {
		w.logger.Error("Nil conn")
		return
	}
	for {
		// cannot read and write concurrently
		w.muxConn.Lock()
		t, d, err := conn.ReadMessage()
		w.muxConn.Unlock()
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
			w.fanout.Broadcast(d)
		case websocket.BinaryMessage:
			w.logger.Info("Server binary received", "msg", string(d), "binary", d)
		case websocket.CloseMessage:
			w.logger.Info("Server closed message", "msg", string(d))
		}
	}
}

func (w *Ws) write(data []byte) error {
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

func (w *Ws) read() (msgType int, data []byte, err error) {
	w.muxConn.Lock()
	defer w.muxConn.Unlock()
	if w.conn == nil {
		return -1, nil, errors.New("bnc_ws: nil conn")
	}
	return w.conn.ReadMessage()
}

func (w *Ws) Sub() <-chan []byte {
	return w.fanout.NewOuter()
}

func (w *Ws) Unsub(c <-chan []byte) {
	w.fanout.RemoveOuter(c)
}

func (w *Ws) SubStream(params []string) error {
	if !w.muxConn.TryLock() {
		return errors.New("bnc_ws: conn is busy")
	}

	err := w.conn.WriteJSON(WsSubMsg{
		Method: WsMethodSub,
		Params: params,
		Id:     1,
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

func (w *Ws) UnsubStream(params []string) error {
	if !w.muxConn.TryLock() {
		return errors.New("bnc_ws: conn is busy")
	}

	err := w.conn.WriteJSON(WsSubMsg{
		Method: WsMethodUnsub,
		Params: params,
		Id:     1,
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

func (w *Ws) newAndKeepListenKey() (string, error) {
	lk, err := w.newListenKey()
	if err != nil {
		return "", err
	}
	w.listenKey = lk
	go w.listenKeyKeeper(w.ctx)
	return lk, nil
}

func (w *Ws) newListenKey() (string, error) {
	_, lk, err := w.user.NewListenKey(w.cfg.ListenKeyUrl)
	return lk.ListenKey, err.Err
}

func (w *Ws) listenKeyKeeper(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		case <-time.After(time.Minute * 20):
			_, _, err := w.user.KeepListenKey(w.cfg.ListenKeyUrl, w.listenKey)
			if err.IsNotNil() {
				w.logger.Error("Cannot Keep Listen Key", "err", err.Error())
			}
		}
	}
}

func (w *Ws) canWriteMsg() bool {
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

func (w *Ws) Request(method WsMethod, params []any) (id int64, err error) {
	w.muxReqId.Lock()
	w.reqId++
	id = w.reqId
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
