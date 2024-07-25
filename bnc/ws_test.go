package bnc

import (
	"testing"
	"time"

	"github.com/dwdwow/props"
)

func TestWs_newFreqToken(t *testing.T) {
	dur := time.Second
	maxreqPerDur := 5
	ws := NewRawWsClient(WsCfg{ReqDur: dur, MaxReqPerDur: maxreqPerDur}, nil, nil)
	for i := 0; i < maxreqPerDur*2; i++ {
		time.Sleep(dur / time.Duration(maxreqPerDur*2))
		ok := ws.canWriteMsg()
		if i < maxreqPerDur == ok {
			t.Log(i, ok)
		} else {
			t.Error(i, ok)
		}
	}
	time.Sleep(dur)
	for i := 0; i < maxreqPerDur*2; i++ {
		time.Sleep(dur / time.Duration(maxreqPerDur*2))
		ok := ws.canWriteMsg()
		if i < maxreqPerDur == ok {
			t.Log(i, ok)
		} else {
			t.Error(i, ok)
		}
	}
}

func TestNewWsClient(t *testing.T) {
	ws := NewWsClient(SpotPrivateWsCfg, newTestVIPPortmarUser(), nil)
	err := ws.Start()
	props.PanicIfNotNil(err)
	chAll, err := ws.Sub("")
	props.PanicIfNotNil(err)
	chBalance, err := ws.Sub(string(WsEventBalanceUpdate))
	props.PanicIfNotNil(err)
	for {
		d := <-chAll
		t.Log("all", d)
		d = <-chBalance
		u := d.Data.(WsSpotBalanceUpdate)
		t.Log("balanceUpdate", u)
	}
}

func TestSpotPublicWsClient(t *testing.T) {
	ws := NewWsClient(SpotPublicWsCfg, nil, nil)
	err := ws.Start()
	props.PanicIfNotNil(err)
	err = ws.SubStream([]string{"btcusdt@depth", "ethusdt@depth@100ms"})
	chAll, err := ws.Sub(string(WsEventDepthUpdate))
	props.PanicIfNotNil(err)
	for {
		msg := <-chAll
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Log(msg.Data)
	}
}

func TestUmFuturesPublicWsClient(t *testing.T) {
	ws := NewWsClient(UmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubStream([]string{"btcusdt@depth", "ethusdt@depth@100ms"})
	chAll, err := ws.Sub(string(WsEventDepthUpdate))
	props.PanicIfNotNil(err)
	for {
		msg := <-chAll
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Log(msg.Data)
	}
}

func TestCmFuturesPublicWsClient(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubStream([]string{"btcusd_perp@depth", "ethusd_perp@depth@100ms"})
	chAll, err := ws.Sub(string(WsEventDepthUpdate))
	props.PanicIfNotNil(err)
	for {
		msg := <-chAll
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Log(msg.Data)
	}
}
