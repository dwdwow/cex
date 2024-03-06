package bnc

import (
	"context"
	"fmt"
	"testing"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/ob"
	"github.com/dwdwow/props"
	"github.com/dwdwow/spub"
)

func TestFuObWs(t *testing.T) {
	publisher := spub.NewSimplePublisher(ob.NewSimplePublisherChannelUtil(), spub.SimpleRcvCapOption[ob.Data](100))
	err := publisher.Start(context.TODO())
	props.PanicIfNotNil(err)
	wsCex := NewWsFuObMsgHandler(nil)
	obWs := ob.NewProducer(wsCex, publisher, nil)
	err = obWs.Start(context.TODO())
	props.PanicIfNotNil(err)
	id, err := ob.ID(cex.BINANCE, cex.PairTypeFutures, "ETHUSDT")
	props.PanicIfNotNil(err)
	sub, err := publisher.Subscribe(context.TODO(), id)
	props.PanicIfNotNil(err)
	for {
		o, closed, err := sub.Receive(context.TODO())
		props.PanicIfNotNil(err)
		if closed {
			panic("closed")
		}
		if o.Empty {
			fmt.Println(o.EmptyReason)
		} else {
			fmt.Println(o.Asks[0])
		}
	}
}
