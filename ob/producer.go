package ob

import (
	"context"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/spub"
	"github.com/dwdwow/ws/wsclt"
)

type CexWsMsgHandler interface {
	Name() cex.Name
	Type() cex.PairType
	Client() *wsclt.MergedClient
	Topics(symbols ...string) []string
	Handle(wsclt.MergedClientMsg) ([]Data, error)
}

type Producer struct {
	c             CexWsMsgHandler
	ps            spub.ProducerService[Data]
	mgClt         *wsclt.MergedClient
	msgCh         chan wsclt.MergedClientMsg
	muxChannelMap sync.Mutex
	channelMap    map[string]bool
	ncWatcher     spub.Subscription[[]string]
	logger        *slog.Logger
}

func NewProducer(c CexWsMsgHandler, producerService spub.ProducerService[Data], logger *slog.Logger) *Producer {
	msgCh := make(chan wsclt.MergedClientMsg, 1000)
	mgClt := c.Client().SetMsgCh(msgCh)
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return &Producer{
		c:          c,
		ps:         producerService,
		mgClt:      mgClt,
		msgCh:      msgCh,
		channelMap: map[string]bool{},
		logger:     logger,
	}
}

func (w *Producer) Start(ctx context.Context) error {
	ncWatcher, err := w.ps.WatchChannels()
	if err != nil {
		return err
	}
	w.ncWatcher = ncWatcher
	go w.receive(ctx)
	go w.channelsWatcher(ctx)
	go w.channelsLooper(ctx)
	return nil
}

func (w *Producer) receive(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-w.msgCh:
			obs, err := w.c.Handle(msg)
			if err != nil {
				w.logger.Error("Can not handle ws msg", "err", err)
				continue
			}
			for _, o := range obs {
				logger := w.logger.With("symbol", o.Symbol)
				if o.Empty {
					logger.Error("Empty ob", "reason", o.EmptyReason)
				}
				channel, err := w.ps.ChannelUtil().Marshal(o)
				if err != nil {
					w.logger.Error("Can not get channel", "err", err)
					continue
				}
				ctx, cancel := context.WithTimeout(context.TODO(), 50*time.Millisecond)
				err = w.ps.Publish(ctx, channel, o)
				cancel()
				if err != nil {
					logger.Error("Can not publish ob", "err", err)
					continue
				}
			}
		}
	}
}

func (w *Producer) channelsWatcher(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case cs := <-w.ncWatcher.Chan():
			// avoid block
			go w.csHandler(cs)
		}
	}
}

func (w *Producer) channelsLooper(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(5 * time.Second):
			nctx, cancel := context.WithTimeout(ctx, time.Second)
			cs, err := w.ps.ConsumerChannels(nctx)
			cancel()
			if err != nil {
				w.logger.Error("Can not get all consumer channels", "err", err)
				return
			}
			go w.csHandler(cs)
		}
	}
}

func (w *Producer) csHandler(cs []string) {
	w.muxChannelMap.Lock()
	defer w.muxChannelMap.Unlock()
	var symbols []string
	for _, c := range cs {
		ok := w.channelMap[c]
		if ok {
			continue
		}
		// set true firstly
		w.channelMap[c] = true
		co, err := w.ps.ChannelUtil().Unmarshal(c)
		if err != nil {
			w.logger.Error("Can not parse channel", "err", err, "channel", c)
			continue
		}
		if co.Cex != w.c.Name() {
			continue
		}
		if co.Type != w.c.Type() {
			continue
		}
		symbols = append(symbols, co.Symbol)
	}
	if len(symbols) == 0 {
		return
	}
	topics := w.c.Topics(symbols...)
	err := w.mgClt.Sub(topics)
	if err != nil {
		w.logger.Error("Can not sub topics by merged client", "err", err)
	}
}
