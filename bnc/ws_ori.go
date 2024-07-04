package bnc

import (
	"sync"
	"time"

	"github.com/dwdwow/props"
	"github.com/gorilla/websocket"
)

type Ws struct {
	conn   *websocket.Conn
	fanout props.Fanout[[]byte]

	muxReqToken   sync.Mutex
	reqTokenDur   time.Duration
	crrTokenIndex int
	latestTokens  []int64
}

func NewWs(reqFreqDur time.Duration, maxReqPerDur int) *Ws {
	return &Ws{
		reqTokenDur:  reqFreqDur,
		latestTokens: make([]int64, maxReqPerDur),
	}
}

func (w *Ws) newReqToken() bool {
	w.muxReqToken.Lock()
	defer w.muxReqToken.Unlock()
	t := time.Now().UnixMilli()
	withinDurNum := 0
	for _, v := range w.latestTokens {
		if t-v < w.reqTokenDur.Milliseconds() {
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
