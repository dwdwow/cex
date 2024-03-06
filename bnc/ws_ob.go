package bnc

import (
	"context"
	"strings"
	"time"

	"github.com/dwdwow/props"
	"github.com/dwdwow/ws/wsclt"
)

const (
	maxTopicNumPerWs = 30
)

var safeMapObDataBuffer = props.NewSafeRWMap[string, []WsDepthMsg]()
var lastObQueryFailTsMilli = props.SafeRWData[int64]{}

func topicSuber(client *wsclt.BaseClient, topics []string) error {
	subMsg := WsSubMsg{
		Method: WsMethodSub,
		Params: topics,
		Id:     time.Now().UnixMilli(),
	}
	err := client.WriteJSON(subMsg)
	if err != nil {
		return err
	}
	return nil
}

func topicUnsuber(client *wsclt.BaseClient, topics []string) error {
	subMsg := WsSubMsg{
		Method: WsMethodUnsub,
		Params: topics,
		Id:     time.Now().UnixMilli(),
	}
	err := client.WriteJSON(subMsg)
	if err != nil {
		return err
	}
	return nil
}

func ping(ctx context.Context, client *wsclt.BaseClient) {
	for {
		select {
		case <-time.After(time.Second * 60):
		case <-ctx.Done():
			return
		}
		err := client.WriteJSON(map[string]string{"method": "PING"})
		if err != nil {
			return
		}
	}
}

func pong(client *wsclt.BaseClient) {
	err := client.Pong([]byte("pong"))
	if err != nil {
		return
	}
}

func CreateObTopic(symbol string) string {
	return strings.ToLower(symbol) + "@depth@100ms"
}
