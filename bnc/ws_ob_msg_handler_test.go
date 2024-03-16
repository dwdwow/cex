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

func TestSpotObWs(t *testing.T) {
	publisher := spub.NewSimplePublisher[ob.Data](ob.NewSimplePublisherChannelUtil(), spub.SimpleRcvCapOption[ob.Data](100))
	err := publisher.Start(context.TODO())
	props.PanicIfNotNil(err)
	wsCex := NewWsSpObMsgHandler(nil)
	obWs := ob.NewProducer(wsCex, publisher, nil)
	err = obWs.Start(context.TODO())
	props.PanicIfNotNil(err)
	id, err := ob.ID(cex.BINANCE, cex.PairTypeSpot, "ETHUSDT")
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

func TestObSimplePublisher(t *testing.T) {
	handler := NewWsSpObMsgHandler(nil)
	publisher := ob.NewSimplePublisher(handler, nil)
	err := publisher.Start(context.TODO())
	props.PanicIfNotNil(err)
	consumer := spub.ConsumerService[ob.Data](publisher)
	c, err := consumer.ChannelUtil().Marshal(ob.Data{Cex: cex.BINANCE, Type: cex.PairTypeSpot, Symbol: "BOMEUSDT"})
	props.PanicIfNotNil(err)
	sub, err := consumer.Subscribe(context.TODO(), c)
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

//func TestObRedisPublisher(t *testing.T) {
//	handler := NewWsObMsgHandler(nil)
//	publisher := ob.NewRedisPublisher(&redis.Options{Addr: ":" + m.I2S(redi.DefaultPort)}, handler, nil)
//	err := publisher.Start(context.TODO())
//	props.PanicIfNotNil(err)
//	consumer := spub.ConsumerService[ob.Data](publisher)
//	c, err := consumer.ChannelUtil().Marshal(ob.Data{Cex: cex.BINANCE, Type: cex.PairTypeSpot, Symbol: "ETHUSDT"})
//	props.PanicIfNotNil(err)
//	//c, err := ob.SimplePublisherChannel(ob.Data{Cex: cex.BINANCE, Type: cex.ObTypeSpot, Symbol: "ETHUSDT"})
//	sub, err := consumer.Subscribe(context.TODO(), c)
//	props.PanicIfNotNil(err)
//	for {
//		o, closed, err := sub.Receive(context.TODO())
//		props.PanicIfNotNil(err)
//		if closed {
//			panic("closed")
//		}
//		if o.Empty {
//			fmt.Println(o.EmptyReason)
//		} else {
//			fmt.Println(o.Asks[0])
//		}
//	}
//}
//
//func TestFutureObSimplePublisher(t *testing.T) {
//	handler := NewWsFutureObMsHandler()
//	publisher := ob.NewSimplePublisher(handler)
//	err := publisher.Start(context.TODO())
//	props.PanicIfNotNil(err)
//	consumer := spub.ConsumerService[ob.Data](publisher)
//	c, err := consumer.ChannelUtil().Marshal(ob.Data{Cex: e.ExBinanceFuture, Type: cex.FuturePair, Symbol: "ETHUSDT"})
//	props.PanicIfNotNil(err)
//	//c, err := ob.SimplePublisherChannel(ob.Data{Cex: cex.BINANCE, Type: cex.ObTypeSpot, Symbol: "ETHUSDT"})
//	sub, err := consumer.Subscribe(context.TODO(), c)
//	props.PanicIfNotNil(err)
//	for {
//		o, closed, err := sub.Receive(context.TODO())
//		props.PanicIfNotNil(err)
//		if closed {
//			panic("closed")
//		}
//		if o.Empty {
//			fmt.Println(o.EmptyReason)
//		} else {
//			fmt.Println(o.Asks[0], o.Asks[1])
//		}
//	}
//}
//
//func TestMergedFutureObSimplePublisher(t *testing.T) {
//	obCh := make(chan ob.Data, 1000)
//	for i := 0; i < 10; i++ {
//		time.Sleep(10 * time.Millisecond)
//		go func() {
//			handler := NewWsFutureObMsHandler()
//			publisher := ob.NewSimplePublisher(handler)
//			err := publisher.Start(context.TODO())
//			props.PanicIfNotNil(err)
//			consumer := spub.ConsumerService[ob.Data](publisher)
//			c, err := consumer.ChannelUtil().Marshal(ob.Data{Cex: e.ExBinanceFuture, Type: cex.FuturePair, Symbol: "ETHUSDT"})
//			props.PanicIfNotNil(err)
//			//c, err := ob.SimplePublisherChannel(ob.Data{Cex: cex.BINANCE, Type: cex.ObTypeSpot, Symbol: "ETHUSDT"})
//			sub, err := consumer.Subscribe(context.TODO(), c)
//			props.PanicIfNotNil(err)
//			for {
//				o, closed, err := sub.Receive(context.TODO())
//				props.PanicIfNotNil(err)
//				if closed {
//					panic("closed")
//				}
//				if o.Empty {
//					fmt.Println(o.EmptyReason)
//				} else {
//					obCh <- o
//				}
//			}
//		}()
//	}
//	for {
//		o := <-obCh
//		fmt.Println(o.Version)
//	}
//}
