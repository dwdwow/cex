package bnc

import (
	"context"
	"fmt"
	"sync"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/ob"
	"github.com/dwdwow/spub"
)

type suber struct {
	spObMux         sync.Mutex
	spObCtx         context.Context
	spObCtxCanceler context.CancelFunc
	spObProducer    *ob.Producer
	spObPuber       spub.Publisher[ob.Data]

	fuObMux         sync.Mutex
	fuObCtx         context.Context
	fuObCtxCanceler context.CancelFunc
	fuObProducer    *ob.Producer
	fuObPuber       spub.Publisher[ob.Data]
}

func (s *suber) subFuOb(ctx context.Context, symbol string) (sub spub.Subscription[ob.Data], err error) {
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
	id, err := ob.ID(cex.BINANCE, cex.PairTypeFutures, symbol)
	if err != nil {
		return
	}
	return s.fuObPuber.Subscribe(ctx, id)
}

func (s *suber) closeFuOb() {
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

func (s *suber) subSpOb(ctx context.Context, symbol string) (sub spub.Subscription[ob.Data], err error) {
	s.spObMux.Lock()
	defer s.spObMux.Unlock()
	if s.spObProducer == nil {
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
		producer := ob.NewProducer(NewWsSpObMsgHandler(nil), publisher, nil)
		if err = producer.Start(obCtx); err != nil {
			return
		}
		s.spObPuber = publisher
		s.spObProducer = producer
		s.spObCtx = obCtx
		s.spObCtxCanceler = cancel
	}
	id, err := ob.ID(cex.BINANCE, cex.PairTypeSpot, symbol)
	if err != nil {
		return
	}
	return s.spObPuber.Subscribe(ctx, id)
}

func (s *suber) closeSpOb() {
	s.spObMux.Lock()
	defer s.spObMux.Unlock()
	if s.spObCtxCanceler != nil {
		s.spObCtxCanceler()
		s.spObCtx = nil
		s.spObCtxCanceler = nil
		s.spObProducer = nil
		s.spObPuber = nil
	}
}

var defaultSuber = &suber{}

func SubOb(ctx context.Context, pairType cex.PairType, symbol string) (spub.Subscription[ob.Data], error) {
	switch pairType {
	case cex.PairTypeFutures:
		return defaultSuber.subFuOb(ctx, symbol)
	case cex.PairTypeSpot:
		return defaultSuber.subSpOb(ctx, symbol)
	}
	return nil, fmt.Errorf("bnc: unknown pair type %v", pairType)
}

func SubObWithSubscription(ctx context.Context, sub spub.Subscription[ob.Data], pairType cex.PairType, symbol string) error {
	id, err := ob.ID(cex.BINANCE, pairType, symbol)
	if err != nil {
		return err
	}
	return sub.Subscribe(ctx, id)
}

func CloseFuObSuber() {
	defaultSuber.closeFuOb()
}
