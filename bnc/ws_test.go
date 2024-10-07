package bnc

import (
	"testing"
	"time"

	"github.com/dwdwow/props"
)

func TestWs_newFreqToken(t *testing.T) {
	dur := time.Second
	maxReqPerDur := 5
	ws := NewRawWsClient(WsCfg{ReqDur: dur, MaxReqPerDur: maxReqPerDur}, nil, nil)
	for i := 0; i < maxReqPerDur*2; i++ {
		time.Sleep(dur / time.Duration(maxReqPerDur*2))
		ok := ws.canWriteMsg()
		if i < maxReqPerDur == ok {
			t.Log(i, ok)
		} else {
			t.Error(i, ok)
		}
	}
	time.Sleep(dur)
	for i := 0; i < maxReqPerDur*2; i++ {
		time.Sleep(dur / time.Duration(maxReqPerDur*2))
		ok := ws.canWriteMsg()
		if i < maxReqPerDur == ok {
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
		t.Logf("%+v", msg.Data)
	}
}

func TestSpotPublicWsClientAggTrade(t *testing.T) {
	ws := NewWsClient(SpotPublicWsCfg, nil, nil)
	err := ws.Start()
	props.PanicIfNotNil(err)
	err = ws.SubAggTradeStream("BTCUSDT")
	props.PanicIfNotNil(err)
	sub, err := ws.SubAggTrade("BTCUSDT")
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
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
		t.Logf("%+v", msg.Data)
	}
}

func TestCmFuturesPublicWsClient(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubStream([]string{"btcusd_perp@kline_1m", "ethusd_perp@kline_1m"})
	chAll, err := ws.Sub(string(WsEventKline))
	props.PanicIfNotNil(err)
	for {
		msg := <-chAll
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubTrade(t *testing.T) {
	ws := NewWsClient(SpotPublicWsCfg, nil, nil)
	err := ws.Start()
	props.PanicIfNotNil(err)
	err = ws.SubTradeStream("ETHUSDT", "BTCUSDT")
	props.PanicIfNotNil(err)
	sub, err := ws.SubTrade("ETHUSDT")
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubAggTrade(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubAggTradeStream("BTCUSD_PERP", "ETHUSD_PERP")
	props.PanicIfNotNil(err)
	sub, err := ws.SubAggTrade("ETHUSD_PERP")
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubKline(t *testing.T) {
	ws := NewWsClient(SpotPublicWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubKlineStream(KlineInterval1s, "ETHUSDT", "BTCUSDT")
	props.PanicIfNotNil(err)
	sub, err := ws.SubKline("ETHUSDT", KlineInterval1s)
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubDepthUpdate(t *testing.T) {
	ws := NewWsClient(SpotPublicWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubDepthUpdateStream100ms("ETHUSDT", "BTCUSDT")
	props.PanicIfNotNil(err)
	sub, err := ws.SubDepthUpdate100ms("BTCUSDT")
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubMarkPrice1s(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubMarkPriceStream3s("ETHUSD_PERP", "BTCUSD_PERP")
	props.PanicIfNotNil(err)
	sub, err := ws.SubMarkPrice3s("ETHUSD_PERP")
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubAllMarkPrice1s(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubAllMarkPriceStream1s()
	props.PanicIfNotNil(err)
	sub, err := ws.SubAllMarkPrice1s()
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubCMIndexPrice1s(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubCMIndexPriceStream3s("ETHUSD", "BTCUSD")
	props.PanicIfNotNil(err)
	sub, err := ws.SubCMIndexPrice3s("ETHUSD")
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubLiquidationOrder(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubLiquidationOrderStream("ETHUSD_PERP", "BTCUSD_PERP")
	props.PanicIfNotNil(err)
	sub, err := ws.SubLiquidationOrder("")
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}

func TestWsClient_SubAllMarketLiquidationOrder(t *testing.T) {
	ws := NewWsClient(CmFuturesWsCfg, nil, nil)
	err := ws.start()
	props.PanicIfNotNil(err)
	err = ws.SubAllMarketLiquidationOrderStream()
	props.PanicIfNotNil(err)
	sub, err := ws.SubAllMarketLiquidationOrder()
	props.PanicIfNotNil(err)
	for {
		msg := <-sub.Chan()
		if msg.Err != nil {
			t.Error(msg.Err)
			break
		}
		t.Logf("%+v", msg.Data)
	}
}
