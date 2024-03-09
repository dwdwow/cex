package bnc

import (
	"context"
	"sync"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/ob"
	"github.com/dwdwow/spub"
)

type suber struct {
	fuObMux         sync.Mutex
	fuObCtx         context.Context
	fuObCtxCanceler context.CancelFunc
	fuObProducer    *ob.Producer
	fuObPuber       spub.Publisher[ob.Data]
}

func (s *suber) subOb(ctx context.Context, pairType cex.PairType, symbol string) (sub spub.Subscription[ob.Data], err error) {
	s.fuObMux.Lock()
	defer s.fuObMux.Unlock()
	if s.fuObProducer == nil {
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
		s.fuObPuber = publisher
		s.fuObProducer = producer
		s.fuObCtx = obCtx
		s.fuObCtxCanceler = cancel
	}
	id, err := ob.ID(cex.BINANCE, pairType, symbol)
	if err != nil {
		return
	}
	return s.fuObPuber.Subscribe(ctx, id)
}

func (s *suber) closeOb() {
	s.fuObMux.Lock()
	defer s.fuObMux.Unlock()
	if s.fuObCtxCanceler != nil {
		s.fuObCtxCanceler()
		s.fuObCtx = nil
		s.fuObCtxCanceler = nil
		s.fuObProducer = nil
		s.fuObPuber = nil
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
