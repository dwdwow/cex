package bnc

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/dwdwow/props"
	"github.com/gorilla/websocket"
)

type WsCfg struct {
	// ws url without streams and auth tokens
	Url          string
	AuthTokenUrl string

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
	authKey    string

	conn   *websocket.Conn
	fanout *props.Fanout[[]byte]

	muxReqToken   sync.Mutex
	crrTokenIndex int
	latestTokens  []int64

	logger *slog.Logger
}

func NewWs(cfg WsCfg, user *User, logger *slog.Logger) *Ws {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Ws{
		ctx:          ctx,
		ctxCancel:    cancel,
		cfg:          cfg,
		user:         user,
		fanout:       props.NewFanout[[]byte](time.Second),
		latestTokens: make([]int64, cfg.MaxReqPerDur),
	}
}

func (w *Ws) authTokenKeeper(ctx context.Context, token string) {
	for {
		select {
		case <-ctx.Done():
		case <-time.After(time.Minute * 30):
		}
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
