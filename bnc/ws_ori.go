package bnc

import (
	"sync"
	"time"

	"github.com/dwdwow/props"
	"github.com/gorilla/websocket"
)

type WsCfg struct {
	Url string

	// binance has incoming massage limitation
	// ex. spot 5/s, futures 10/s
	ReqDur       time.Duration
	MaxReqPerDur int
}

type Ws struct {
	cfg WsCfg

	conn   *websocket.Conn
	fanout props.Fanout[[]byte]

	muxReqToken   sync.Mutex
	crrTokenIndex int
	latestTokens  []int64
}

func NewWs(cfg WsCfg) *Ws {
	return &Ws{
		cfg:          cfg,
		latestTokens: make([]int64, cfg.MaxReqPerDur),
	}
}

func (w *Ws) newReqToken() bool {
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
