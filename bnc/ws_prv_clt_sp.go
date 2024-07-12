package bnc

import (
	"fmt"
	"math"
	"time"
)

type WsSpotBalance struct {
	Asset  string  `json:"a"`
	Free   float64 `json:"f,string"`
	Locked float64 `json:"l,string"`
}

type WsSpotAccountUpdate struct {
	EventType  string          `json:"e"`
	EventTime  int64           `json:"E"`
	UpdateTime int64           `json:"u"`
	Balances   []WsSpotBalance `json:"B"`

	// just for margin user data stream
	TimeUpdateId int64 `json:"U"`
}

type WsSpotBalanceUpdate struct {
	EventType        string  `json:"e"`
	EventTime        int64   `json:"E"`
	Asset            string  `json:"a"`
	SignBalanceDelta float64 `json:"d,string"`
	ClearTime        int64   `json:"T"`
}

func (w WsSpotBalanceUpdate) AbsBalanceDelta() float64 {
	return math.Abs(w.SignBalanceDelta)
}

type WsSpotListStatusObject struct {
	Symbol        string `json:"s"`
	OrderId       int64  `json:"i"`
	ClientOrderId string `json:"c"`
}

type WsSpotListStatus struct {
	EventType         string                   `json:"e"`
	EventTime         int64                    `json:"E"`
	Symbol            string                   `json:"s"`
	OrderListId       int                      `json:"g"`
	ContingencyType   string                   `json:"c"`
	ListStatusType    string                   `json:"l"`
	ListOrderStatus   string                   `json:"L"`
	ListRejectReason  string                   `json:"r"`
	ListClientOrderId string                   `json:"C"`
	Time              int64                    `json:"T"`
	Objects           []WsSpotListStatusObject `json:"O"`
}

func SpotWsPrivateMsgUnmarshaler(e WsEvent, data []byte) (any, error) {
	switch e {
	case WsEventOutboundAccountPosition:
		return unmarshal[WsSpotAccountUpdate](data)
	case WsEventBalanceUpdate:
		return unmarshal[WsSpotBalanceUpdate](data)
	case WsEventExecutionReport:
		return unmarshal[WsOrderExecutionReport](data)
	case WsEventListStatus:
		return unmarshal[WsSpotListStatus](data)
	case WsEventListenKeyExpired:
		return unmarshal[WsListenKeyExpired](data)
	default:
		return nil, fmt.Errorf("bnc: unknown event %v", e)
	}
}

var SpotPrivateWsCfg = WsCfg{
	Url:             WsBaseUrl,
	ListenKeyUrl:    SpotListenKeyUrl,
	MaxStream:       1024,
	ReqDur:          time.Second,
	MaxReqPerDur:    10,
	DataUnmarshaler: SpotWsPrivateMsgUnmarshaler,
}
