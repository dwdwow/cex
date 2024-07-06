package bnc

import (
	"testing"
	"time"
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
	ws.Start()
	chAll := ws.Sub("")
	chBalance := ws.Sub(WsEventBalanceUpdate)
	for {
		d := <-chAll
		t.Log("all", d)
		d = <-chBalance
		t.Log("balanceUpdate", d)
	}
}
