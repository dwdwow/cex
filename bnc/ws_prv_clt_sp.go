package bnc

import (
	"time"
)

type WsSpotBalance struct {
	Asset  string `json:"a"`
	Free   string `json:"f"`
	Locked string `json:"l"`
}

type WsSpotAccountUpdate struct {
	EventType  string          `json:"e"`
	EventTime  int64           `json:"E"`
	UpdateTime int64           `json:"u"`
	Balances   []WsSpotBalance `json:"B"`
}

type WsSpotBalanceUpdate struct {
	EventType    string `json:"e"`
	EventTime    int64  `json:"E"`
	Asset        string `json:"a"`
	BalanceDelta string `json:"d"`
	ClearTime    int64  `json:"T"`
}

type WsSpotOrderExecutionReport struct {
	EventType               string                  `json:"e"`
	EventTime               int64                   `json:"E"`
	Symbol                  string                  `json:"s"`
	ClientOrderId           string                  `json:"c"`
	Side                    OrderSide               `json:"S"`
	Type                    OrderType               `json:"o"`
	TimeInForce             TimeInForce             `json:"f"`
	Qty                     float64                 `json:"q,string"`
	Price                   float64                 `json:"p,string"`
	StopPrice               float64                 `json:"P,string"`
	IcebergQty              float64                 `json:"F,string"`
	OrderListId             int64                   `json:"g"`
	OriginalClientId        string                  `json:"C"`
	ExecutionType           OrderExecutionType      `json:"x"`
	Status                  OrderStatus             `json:"X"`
	RejectReason            string                  `json:"r"`
	OrderId                 int64                   `json:"i"`
	LastExecutedQty         float64                 `json:"l,string"`
	FilledQty               float64                 `json:"z,string"`
	LastExecutedPrice       float64                 `json:"L,string"`
	CommissionAmt           float64                 `json:"n,string"`
	CommissionAsset         string                  `json:"N"`
	Time                    int64                   `json:"T"`
	TradeId                 int64                   `json:"t"`
	PreventedMatchId        int64                   `json:"v"`
	Ignore                  int64                   `json:"I"`
	IsOrderOnTheBook        bool                    `json:"w"`
	IsMaker                 bool                    `json:"m"`
	Ignore1                 bool                    `json:"M"`
	CreationTime            int64                   `json:"O"`
	FilledQuote             float64                 `json:"Z,string"`
	LastExecutedQuote       float64                 `json:"Y,string"`
	QuoteOrderQty           float64                 `json:"Q,string"`
	WorkingTime             int64                   `json:"W"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"V"`
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

var SpotPrivateWsCfg = WsCfg{
	Url:          WsBaseUrl,
	ListenKeyUrl: SpotListenKeyUrl,
	MaxStream:    1024,
	ReqDur:       time.Second,
	MaxReqPerDur: 10,
}