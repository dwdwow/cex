package cex

import "errors"

var (
	ErrOrderRejected = errors.New("cex: order is rejected")
)

type TradeType string

const (
	TradeLimit  TradeType = "LIMIT"
	TradeMarket TradeType = "MARKET"
)

type TradeSide string

const (
	TradeBuy  TradeSide = "BUY"
	TradeSell TradeSide = "SELL"
)

type OrderStatus string

const (
	OrderStatusNew             OrderStatus = "NEW"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusErr             OrderStatus = "ERR"
)

type Order struct {
	// popular by user self
	Asset    string  `json:"asset"`
	Quote    string  `json:"quote"`
	OriQty   float64 `json:"oriQty"`
	OriPrice float64 `json:"oriPrice"`

	// popular by user or code
	Cex       Cex       `json:"cex"`
	PairType  PairType  `json:"pairType"`
	TradeType TradeType `json:"tradeType"`
	TradeSide TradeSide `json:"tradeSide"`

	// popular by code
	Symbol        string `json:"symbol"`
	TimeInForce   string `json:"timeInForce"`
	ClientOrderId string `json:"clientOrderId"`
	SendTsMilli   int64  `json:"sendTsMilli"`
	RspTsMilli    int64  `json:"rspTsMilli"`
	LocalId       string `json:"localId"`
	ApiKey        string `json:"apiKey"`

	// popular by cex server response
	OrderId string      `json:"orderId"`
	Status  OrderStatus `json:"status"`

	// popular as order result
	FilledQty      float64 `json:"filledQty"`
	FilledAvgPrice float64 `json:"filledAvgPrice"`
	FilledQuote    float64 `json:"filledQuote"`

	// calculate as raw order or popular code
	FeeTier float64 `json:"feeTier"`

	Err error `json:"err"`
}
