package bnc

import (
	"context"
	"sync"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/ob"
	"github.com/dwdwow/spub"
)

type suber struct {
	obMux         sync.Mutex
	obCtx         context.Context
	obCtxCanceler context.CancelFunc
	obProducer    *ob.Producer
	obPuber       spub.Publisher[ob.Data]
}

func (s *suber) subOb(pairType cex.PairType, symbol string) (sub spub.Subscription[ob.Data], err error) {
	s.obMux.Lock()
	defer s.obMux.Unlock()
	if s.obProducer == nil {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			if err != nil {
				cancel()
			}
		}()
		publisher := spub.NewSimplePublisher(ob.NewSimplePublisherChannelUtil(), spub.SimpleRcvCapOption[ob.Data](100))
		if err = publisher.Start(ctx); err != nil {
			return
		}
		producer := ob.NewProducer(NewWsFuObMsgHandler(nil), publisher, nil)
		if err = producer.Start(ctx); err != nil {
			return
		}
		s.obPuber = publisher
		s.obProducer = producer
		s.obCtx = ctx
		s.obCtxCanceler = cancel
	}
	id, err := ob.ID(cex.BINANCE, pairType, symbol)
	if err != nil {
		return
	}
	return s.obPuber.Subscribe(s.obCtx, id)
}

func (s *suber) closeOb() {
	s.obMux.Lock()
	defer s.obMux.Unlock()
	if s.obCtxCanceler != nil {
		s.obCtxCanceler()
		s.obCtx = nil
		s.obCtxCanceler = nil
		s.obProducer = nil
		s.obPuber = nil
	}
}

var defaultSuber = &suber{}

func SubOb(pairType cex.PairType, symbol string) (spub.Subscription[ob.Data], error) {
	return defaultSuber.subOb(pairType, symbol)
}
