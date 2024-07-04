package bnc

import (
	"testing"
	"time"
)

func TestWs_newFreqToken(t *testing.T) {
	dur := time.Second
	maxreqPerDur := 5
	ws := NewWs(WsCfg{ReqDur: dur, MaxReqPerDur: maxreqPerDur}, nil, nil)
	for i := 0; i < maxreqPerDur*2; i++ {
		time.Sleep(dur / time.Duration(maxreqPerDur*2))
		ok := ws.newReqToken()
		if i < maxreqPerDur == ok {
			t.Log(i, ok)
		} else {
			t.Error(i, ok)
		}
	}
	time.Sleep(dur)
	for i := 0; i < maxreqPerDur*2; i++ {
		time.Sleep(dur / time.Duration(maxreqPerDur*2))
		ok := ws.newReqToken()
		if i < maxreqPerDur == ok {
			t.Log(i, ok)
		} else {
			t.Error(i, ok)
		}
	}
}
