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

func (s *suber) subOb(ctx context.Context, pairType cex.PairType, symbol string) (sub spub.Subscription[ob.Data], err error) {
	s.obMux.Lock()
	defer s.obMux.Unlock()
	if s.obProducer == nil {
		obCtx, cancel := context.WithCancel(context.Background())
		defer func() {
			if err != nil {
				cancel()
			}
		}()
		publisher := spub.NewSimplePublisher(ob.NewSimplePublisherChannelUtil(), spub.SimpleRcvCapOption[ob.Data](100))
		if err = publisher.Start(obCtx); err != nil {
			return
		}
		producer := ob.NewProducer(NewWsFuObMsgHandler(nil), publisher, nil)
		if err = producer.Start(obCtx); err != nil {
			return
		}
		s.obPuber = publisher
		s.obProducer = producer
		s.obCtx = obCtx
		s.obCtxCanceler = cancel
	}
	id, err := ob.ID(cex.BINANCE, pairType, symbol)
	if err != nil {
		return
	}
	return s.obPuber.Subscribe(ctx, id)
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

func SubOb(ctx context.Context, pairType cex.PairType, symbol string) (spub.Subscription[ob.Data], error) {
	return defaultSuber.subOb(ctx, pairType, symbol)
}

func SubObWithSubsription(ctx context.Context, sub spub.Subscription[ob.Data], pairType cex.PairType, symbol string) error {
	id, err := ob.ID(cex.BINANCE, pairType, symbol)
	if err != nil {
		return err
	}
	return sub.Subscribe(ctx, id)
}

func CloseObSuber() {
	defaultSuber.closeOb()
}
