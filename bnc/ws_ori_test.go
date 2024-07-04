package bnc

import (
	"testing"
	"time"
)

func TestWs_newFreqToken(t *testing.T) {
	dur := time.Second
	maxTokenNum := 5
	ws := NewWs(time.Second, maxTokenNum)
	for i := 0; i < maxTokenNum*2; i++ {
		time.Sleep(dur / time.Duration(maxTokenNum*2))
		ok := ws.newReqToken()
		if i < maxTokenNum == ok {
			t.Log(i, ok)
		} else {
			t.Error(i, ok)
		}
	}
	time.Sleep(dur)
	for i := 0; i < maxTokenNum*2; i++ {
		time.Sleep(dur / time.Duration(maxTokenNum*2))
		ok := ws.newReqToken()
		if i < maxTokenNum == ok {
			t.Log(i, ok)
		} else {
			t.Error(i, ok)
		}
	}
}
