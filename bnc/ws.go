package bnc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/dwdwow/props"
	"github.com/gorilla/websocket"
)

type WsDataUnmarshaler func(e WsEvent, isArray bool, data []byte) (any, error)

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

	DataUnmarshaler WsDataUnmarshaler
}

type RawWsClientMsg struct {
	Data []byte `json:"data"`
	Err  error  `json:"err"`
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

	fan *props.Fanout[RawWsClientMsg]

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
		fan:          props.NewFanout[RawWsClientMsg](time.Second),
		latestTokens: make([]int64, cfg.MaxReqPerDur),
		chRestart:    make(chan struct{}),
		reqId:        1000,
		logger:       logger,
	}
}

func (w *RawWsClient) Start() error {
	w.muxStatus.Lock()
	w.status = 1
	w.muxStatus.Unlock()
	err := w.start()
	if err != nil {
		return err
	}
	go w.mainThreadStarter()
	//w.chRestart <- struct{}{}
	return nil
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
		<-w.chRestart
		w.muxStatus.Lock()
		status := w.status
		w.muxStatus.Unlock()
		if status == -1 {
			return
		}
		err := w.start()
		if err != nil {
			w.logger.Error("Cannot Start", "err", err)
			w.fan.Broadcast(RawWsClientMsg{nil, err})
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
	var path string
	if lk != "" {
		path = "/" + lk
	}
	conn, resp, err := dialer.DialContext(w.ctx, w.cfg.Url+path, nil)
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
		t, d, err := w.read()
		if err != nil {
			w.logger.Error("Read message", "err", err)
			w.fan.Broadcast(RawWsClientMsg{nil, err})
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
			w.fan.Broadcast(RawWsClientMsg{Data: d})
		case websocket.BinaryMessage:
			w.logger.Info("Server binary received", "msg", string(d), "binary", d)
		case websocket.CloseMessage:
			w.logger.Info("Server closed message", "msg", string(d))
		}
	}
}

func (w *RawWsClient) write(data []byte) error {
	// cannot read and write concurrently
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
	// cannot read and write concurrently
	w.muxConn.Lock()
	defer w.muxConn.Unlock()
	if w.conn == nil {
		return -1, nil, errors.New("bnc_ws: nil conn")
	}
	return w.conn.ReadMessage()
}

func (w *RawWsClient) Sub() <-chan RawWsClientMsg {
	return w.fan.Sub()
}

func (w *RawWsClient) Unsub(c <-chan RawWsClientMsg) {
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

	var err error

	if strings.Contains(w.cfg.Url, "dstream.binance.com") {
		err = w.conn.WriteJSON(WsSubMsgInt64Id{
			Method: WsMethodSub,
			Params: params,
			Id:     1,
		})
	} else {
		err = w.conn.WriteJSON(WsSubMsg{
			Method: WsMethodSub,
			Params: params,
			Id:     "1",
		})
	}

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

	var err error

	if strings.Contains(w.cfg.Url, "dstream.binance.com") {
		err = w.conn.WriteJSON(WsSubMsgInt64Id{
			Method: WsMethodUnsub,
			Params: params,
			Id:     1,
		})
	} else {
		err = w.conn.WriteJSON(WsSubMsg{
			Method: WsMethodUnsub,
			Params: params,
			Id:     "1",
		})
	}

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
	if w.user == nil || w.cfg.ListenKeyUrl == "" {
		return "", nil
	}
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

type WsClientMsg struct {
	Event   WsEvent `json:"event"`
	Data    any     `json:"data"`
	IsArray bool    `json:"isArray"`
	Err     error   `json:"err"`
}

type WsClientSubscriptionMsg[D any] struct {
	D   D     `json:"d"`
	Err error `json:"err"`
}

type WsClientSubscription[D any] struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	rawCh     <-chan WsClientMsg
	ch        chan WsClientSubscriptionMsg[D]
}

func NewWsClientSubscription[D any](ch <-chan WsClientMsg) *WsClientSubscription[D] {
	ctx, ctxCancel := context.WithCancel(context.Background())
	return &WsClientSubscription[D]{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		rawCh:     ch,
		ch:        make(chan WsClientSubscriptionMsg[D], 1),
	}
}

func (w *WsClientSubscription[D]) start() {
	go func() {
		for {
			var rawMsg WsClientMsg
			select {
			case <-w.ctx.Done():
				return
			case rawMsg = <-w.rawCh:
			}
			var msg WsClientSubscriptionMsg[D]
			if rawMsg.Err != nil {
				msg.Err = rawMsg.Err
			} else {
				var ok bool
				msg.D, ok = rawMsg.Data.(D)
				if !ok {
					msg.Err = errors.New("invalid data type")
				}
			}
			w.ch <- msg
		}
	}()
}

func (w *WsClientSubscription[D]) close() {
	w.ctxCancel()
}

func (w *WsClientSubscription[D]) Chan() <-chan WsClientSubscriptionMsg[D] {
	return w.ch
}

type WsClient struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	wsCfg WsCfg
	rawWs *RawWsClient

	user *User

	muxFan sync.Mutex
	mfan   map[string]*props.Fanout[WsClientMsg]

	logger *slog.Logger
}

func NewWsClient(cfg WsCfg, user *User, logger *slog.Logger) *WsClient {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &WsClient{
		ctx:       ctx,
		ctxCancel: cancel,
		wsCfg:     cfg,
		user:      user,
		mfan:      map[string]*props.Fanout[WsClientMsg]{},
		logger:    logger,
	}
}

func (w *WsClient) Start() error {
	return w.start()
}

func (w *WsClient) start() error {
	rawWs := NewRawWsClient(w.wsCfg, w.user, w.logger)
	w.rawWs = rawWs
	ch := rawWs.Sub()
	if err := rawWs.Start(); err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			case msg := <-ch:
				w.dataHandler(msg)
			}
		}
	}()
	return nil
}

const mfanKeyAll = "__all__"

func (w *WsClient) dataHandler(msg RawWsClientMsg) {
	if msg.Err != nil {
		w.sendToAll(WsClientMsg{Data: msg.Data, Err: msg.Err})
		return
	}
	data := msg.Data
	e, isArray, ok := getWsEvent(data)
	if !ok {
		if string(data) == "{\"result\":null,\"id\":\"1\"}" ||
			string(data) == "{\"result\":null,\"id\":1}" {
			return
		}
		w.logger.Error("Can not get event", "data", string(data))
		return
	}
	w.muxFan.Lock()
	defer w.muxFan.Unlock()

	var specFans []*props.Fanout[WsClientMsg]
	singleFan := w.mfan[string(e)]
	allFan := w.mfan[mfanKeyAll]
	var d any = data
	var err error
	if w.rawWs.cfg.DataUnmarshaler != nil {
		d, err = w.rawWs.cfg.DataUnmarshaler(e, isArray, data)
		if err != nil {
			w.logger.Error("Can not unmarshal msg", "err", err, "data", string(data))
			return
		}
		specFans = w.specFans(e, isArray, d)
	}
	newMsg := WsClientMsg{Event: e, Data: d, IsArray: isArray}
	if singleFan != nil {
		singleFan.Broadcast(newMsg)
	}
	if allFan != nil {
		allFan.Broadcast(newMsg)
	}
	for _, fan := range specFans {
		if fan != nil {
			fan.Broadcast(newMsg)
		}
	}
}

func (w *WsClient) specFans(e WsEvent, isArray bool, d any) (fans []*props.Fanout[WsClientMsg]) {
	switch e {
	case WsEventAggTrade:
		s, ok := d.(WsAggTradeStream)
		if !ok {
			return
		}
		fan := w.mfan[strings.ToLower(s.Symbol)+"@"+string(WsEventAggTrade)]
		if fan != nil {
			fans = append(fans, fan)
		}
	case WsEventTrade:
		s, ok := d.(WsTradeStream)
		if !ok {
			return
		}
		fan := w.mfan[strings.ToLower(s.Symbol)+"@"+string(WsEventTrade)]
		if fan != nil {
			fans = append(fans, fan)
		}
	case WsEventKline:
		s, ok := d.(WsKlineStream)
		if !ok {
			return
		}
		fan := w.mfan[strings.ToLower(s.Symbol)+"@"+string(WsEventKline)+"_"+string(s.Kline.Interval)]
		if fan != nil {
			fans = append(fans, fan)
		}
	case WsEventDepthUpdate:
		s, ok := d.(WsDepthStream)
		if !ok {
			return
		}
		fan := w.mfan[strings.ToLower(s.Symbol)+"@depth"]
		if fan != nil {
			fans = append(fans, fan)
		}
		fan = w.mfan[strings.ToLower(s.Symbol)+"@depth@100ms"]
		if fan != nil {
			fans = append(fans, fan)
		}
		fan = w.mfan[strings.ToLower(s.Symbol)+"@depth@250ms"]
		if fan != nil {
			fans = append(fans, fan)
		}
		fan = w.mfan[strings.ToLower(s.Symbol)+"@depth@500ms"]
		if fan != nil {
			fans = append(fans, fan)
		}
	case WsEventMarkPriceUpdate:
		if isArray {
			fan := w.mfan["!markPrice@arr@1s"]
			if fan != nil {
				fans = append(fans, fan)
			}
			fan = w.mfan["!markPrice@arr"]
			if fan != nil {
				fans = append(fans, fan)
			}
			return
		}
		s, ok := d.(WsMarkPriceStream)
		if !ok {
			return
		}
		fan := w.mfan[strings.ToLower(s.Symbol)+"@markPrice"]
		if fan != nil {
			fans = append(fans, fan)
		}
		fan = w.mfan[strings.ToLower(s.Symbol)+"@markPrice@1s"]
		if fan != nil {
			fans = append(fans, fan)
		}
	case WsEventIndexPriceUpdate:
		s, ok := d.(WsCMIndexPriceStream)
		if !ok {
			return
		}
		fan := w.mfan[strings.ToLower(s.Pair)+"@indexPrice"]
		if fan != nil {
			fans = append(fans, fan)
		}
		fan = w.mfan[strings.ToLower(s.Pair)+"@indexPrice@1s"]
		if fan != nil {
			fans = append(fans, fan)
		}
	case WsEventForceOrder:
		if isArray {
			fan := w.mfan["!forceOrder@arr"]
			if fan != nil {
				fans = append(fans, fan)
			}
			return
		}
		s, ok := d.(WsLiquidationOrderStream)
		if !ok {
			return
		}
		fan := w.mfan[strings.ToLower(s.Order.Symbol)+"@"+string(WsEventForceOrder)]
		if fan != nil {
			fans = append(fans, fan)
		}
	}
	return
}

func (w *WsClient) sendToAll(msg WsClientMsg) {
	w.muxFan.Lock()
	defer w.muxFan.Unlock()
	for _, fan := range w.mfan {
		if fan != nil {
			fan.Broadcast(msg)
		}
	}
}

func (w *WsClient) event2MfanKey(event string) string {
	// do not use empty string as mfan key
	if event == "" {
		event = mfanKeyAll
	}
	return event
}

// Sub
// Pass empty string if you want listen all events.
// Should SubStream firstly.
func (w *WsClient) Sub(event string) (<-chan WsClientMsg, error) {
	w.muxFan.Lock()
	defer w.muxFan.Unlock()
	// do not use empty string as mfan key
	key := w.event2MfanKey(event)
	fan := w.mfan[key]
	if fan == nil {
		fan = props.NewFanout[WsClientMsg](time.Second)
		w.mfan[key] = fan
	}
	return fan.Sub(), nil
}

func (w *WsClient) Unsub(event string, ch <-chan WsClientMsg) {
	w.muxFan.Lock()
	defer w.muxFan.Unlock()
	key := w.event2MfanKey(event)
	fan := w.mfan[key]
	if fan == nil {
		return
	}
	fan.Unsub(ch)
}

func (w *WsClient) SubStream(events []string) error {
	return w.rawWs.SubStream(events)
}

func wsClientSubEvent[D any](rawChGetter func() (<-chan WsClientMsg, error)) (*WsClientSubscription[D], error) {
	ch, err := rawChGetter()
	if err != nil {
		return nil, err
	}
	return NewWsClientSubscription[D](ch), nil
}

// SubAggTradeStream real time
func (w *WsClient) SubAggTradeStream(symbol string) error {
	return w.SubStream([]string{symbol + "@aggTrade"})
}

// SubAggTrade real time
// if symbol is empty, will listen all aggTrade events
func (w *WsClient) SubAggTrade(symbol string) (*WsClientSubscription[WsAggTradeStream], error) {
	return wsClientSubEvent[WsAggTradeStream](func() (<-chan WsClientMsg, error) {
		if symbol != "" {
			return w.Sub(symbol + "@aggTrade")
		}
		return w.Sub(string(WsEventAggTrade))
	})
}

// SubTradeStream real time
// just for spot ws
func (w *WsClient) SubTradeStream(symbol string) error {
	return w.SubStream([]string{symbol + "@trade"})
}

// SubTrade real time
// if symbol is empty, will listen all trade events
func (w *WsClient) SubTrade(symbol string) (*WsClientSubscription[WsTradeStream], error) {
	return wsClientSubEvent[WsTradeStream](func() (<-chan WsClientMsg, error) {
		if symbol != "" {
			return w.Sub(symbol + "@trade")

		}
		return w.Sub(string(WsEventTrade))
	})
}

// SubKlineStream 1000ms for 1s, 2000ms for others
// 1s just for spot kline
func (w *WsClient) SubKlineStream(symbol string, interval KlineInterval) error {
	return w.SubStream([]string{fmt.Sprintf("%s@kline_%v", symbol, interval)})
}

func (w *WsClient) SubKline(symbol string, interval KlineInterval) (*WsClientSubscription[WsKlineStream], error) {
	return wsClientSubEvent[WsKlineStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(fmt.Sprintf("%s@kline_%v", symbol, interval))
	})
}

// SubDepthUpdateStream 1000ms for spot, 250ms for futures
func (w *WsClient) SubDepthUpdateStream(symbol string) error {
	return w.SubStream([]string{symbol + "@depth"})
}

// SubDepthUpdateStream500ms 500ms
// just for futures ws
func (w *WsClient) SubDepthUpdateStream500ms(symbol string) error {
	return w.SubStream([]string{symbol + "@depth@500ms"})
}

// SubDepthUpdateStream100ms 100ms
func (w *WsClient) SubDepthUpdateStream100ms(symbol string) error {
	return w.SubStream([]string{symbol + "@depth@100ms"})
}

func (w *WsClient) SubDepthUpdate(symbol string) (*WsClientSubscription[WsDepthStream], error) {
	return wsClientSubEvent[WsDepthStream](func() (<-chan WsClientMsg, error) {
		if symbol != "" {
			return w.Sub(symbol + "@depth")
		}
		return w.Sub(string(WsEventDepthUpdate))
	})
}

// SubMarkPriceStream1s 1s
func (w *WsClient) SubMarkPriceStream1s(symbol string) error {
	return w.SubStream([]string{symbol + "@markPrice@1s"})
}

// SubMarkPriceStream3s 3s
func (w *WsClient) SubMarkPriceStream3s(symbol string) error {
	return w.SubStream([]string{symbol + "@markPrice"})
}

func (w *WsClient) SubMarkPrice1s(symbol string) (*WsClientSubscription[WsMarkPriceStream], error) {
	return wsClientSubEvent[WsMarkPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(symbol + "@markPrice@1s")
	})
}

func (w *WsClient) SubMarkPrice3s(symbol string) (*WsClientSubscription[WsMarkPriceStream], error) {
	return wsClientSubEvent[WsMarkPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(symbol + "@markPrice")
	})
}

// SubAllMarkPriceStream1s 1s
// just for um futures
func (w *WsClient) SubAllMarkPriceStream1s() error {
	return w.SubStream([]string{"!markPrice@arr@1s"})
}

// SubAllMarkPriceStream3s 3s
// just for um futures
func (w *WsClient) SubAllMarkPriceStream3s() error {
	return w.SubStream([]string{"!markPrice@arr"})
}

func (w *WsClient) SubAllMarkPrice1s() (*WsClientSubscription[[]WsMarkPriceStream], error) {
	return wsClientSubEvent[[]WsMarkPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub("!markPrice@arr@1s")
	})
}

func (w *WsClient) SubAllMarkPrice3s() (*WsClientSubscription[[]WsMarkPriceStream], error) {
	return wsClientSubEvent[[]WsMarkPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub("!markPrice@arr")
	})
}

func (w *WsClient) SubAllMarkPriceEvents() (*WsClientSubscription[[]WsMarkPriceStream], error) {
	return wsClientSubEvent[[]WsMarkPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(string(WsEventMarkPriceUpdate))
	})
}

// SubCMIndexPriceStream3s 3s
// just for cm futures
func (w *WsClient) SubCMIndexPriceStream3s(pair string) error {
	return w.SubStream([]string{pair + "@indexPrice"})
}

// SubCMIndexPriceStream1s 1s
// just for cm futures
func (w *WsClient) SubCMIndexPriceStream1s(pair string) error {
	return w.SubStream([]string{pair + "@indexPrice@1s"})
}

// SubCMIndexPrice3s
// just for cm futures
func (w *WsClient) SubCMIndexPrice3s(pair string) (*WsClientSubscription[WsCMIndexPriceStream], error) {
	return wsClientSubEvent[WsCMIndexPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(pair + "@indexPrice")
	})
}

// SubCMIndexPrice1s
// just for cm futures
func (w *WsClient) SubCMIndexPrice1s(pair string) (*WsClientSubscription[WsCMIndexPriceStream], error) {
	return wsClientSubEvent[WsCMIndexPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(pair + "@indexPrice@1s")
	})
}

// SubAllCMIndexPriceEvents
// just for cm futures
func (w *WsClient) SubAllCMIndexPriceEvents() (*WsClientSubscription[WsCMIndexPriceStream], error) {
	return wsClientSubEvent[WsCMIndexPriceStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(string(WsEventIndexPriceUpdate))
	})
}

// SubLiquidationOrderStream 1s
func (w *WsClient) SubLiquidationOrderStream(symbol string) error {
	return w.SubStream([]string{symbol + "@forceOrder"})
}

func (w *WsClient) SubLiquidationOrder(symbol string) (*WsClientSubscription[WsLiquidationOrderStream], error) {
	return wsClientSubEvent[WsLiquidationOrderStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(symbol + "@forceOrder")
	})
}

func (w *WsClient) SubAllLiquidationOrderEvents() (*WsClientSubscription[WsLiquidationOrderStream], error) {
	return wsClientSubEvent[WsLiquidationOrderStream](func() (<-chan WsClientMsg, error) {
		return w.Sub(string(WsEventForceOrder))
	})
}

// SubAllMarketLiquidationOrderStream 1s
func (w *WsClient) SubAllMarketLiquidationOrderStream() error {
	return w.SubStream([]string{"!forceOrder@arr"})
}

func (w *WsClient) SubAllMarketLiquidationOrder() (*WsClientSubscription[[]WsLiquidationOrderStream], error) {
	return wsClientSubEvent[[]WsLiquidationOrderStream](func() (<-chan WsClientMsg, error) {
		return w.Sub("!forceOrder@arr")
	})
}
