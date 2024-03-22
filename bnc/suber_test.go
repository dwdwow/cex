package bnc

import (
	"context"
	"testing"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/props"
)

func TestSubOb(t *testing.T) {
	pairs, _, err := QueryFuturesPairs()
	props.PanicIfNotNil(err)
	sub, err := SubOb(context.Background(), cex.PairTypeFutures, "ETHUSDT")
	props.PanicIfNotNil(err)
	for _, p := range pairs {
		syb := p.PairSymbol
		go func() {
			err := SubObWithSubscription(context.Background(), sub, cex.PairTypeFutures, syb)
			props.PanicIfNotNil(err)
		}()
	}
	for {
		o, closed, err := sub.Receive(context.TODO())
		props.PanicIfNotNil(err)
		if closed {
			panic("closed")
		}
		if o.Empty {
			t.Error(o.Symbol, o.EmptyReason)
		} else {
			t.Log(o.Symbol, o.Asks[0])
		}
	}
}

func TestSubSpOb(t *testing.T) {
	sub, err := SubOb(context.Background(), cex.PairTypeSpot, "BOMEUSDT")
	props.PanicIfNotNil(err)
	for {
		o, closed, err := sub.Receive(context.TODO())
		props.PanicIfNotNil(err)
		if closed {
			panic("closed")
		}
		if o.Empty {
			t.Error(o.Symbol, o.EmptyReason)
		} else {
			t.Log(o.Symbol, o.Asks[0])
		}
	}
}
