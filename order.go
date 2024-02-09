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

type OrderErr struct {
	Err  string `json:"err" bson:"err"`
	Time int64  `json:"time" bson:"time"`
}

func (oe OrderErr) Error() error {
	return errors.New(oe.Err)
}

type Order struct {
	// popular by user self
	Asset    string  `json:"asset" bson:"asset"`
	Quote    string  `json:"quote" bson:"quote"`
	OriQty   float64 `json:"oriQty" bson:"oriQty"`
	OriPrice float64 `json:"oriPrice" bson:"oriPrice"`

	// popular by user or code
	Cex       Cex       `json:"cex" bson:"cex"`
	PairType  PairType  `json:"pairType" bson:"pairType"`
	TradeType TradeType `json:"tradeType" bson:"tradeType"`
	TradeSide TradeSide `json:"tradeSide" bson:"tradeSide"`

	// popular by code
	Symbol        string `json:"symbol" bson:"symbol"`
	TimeInForce   string `json:"timeInForce" bson:"timeInForce"`
	ClientOrderId string `json:"clientOrderId" bson:"clientOrderId"`
	SendTime      int64  `json:"sendTime" bson:"sendTime"`
	RespTime      int64  `json:"respTime" bson:"respTime"`
	LocalId       string `json:"localId" bson:"localId"`
	ApiKey        string `json:"apiKey" bson:"apiKey"`

	// popular as response
	OrderId string      `json:"orderId" bson:"orderId"`
	Status  OrderStatus `json:"status" bson:"status"`

	// popular a order result
	FilledQty      float64 `json:"filledQty" bson:"filledQty"`
	FilledAvgPrice float64 `json:"filledAvgPrice" bson:"filledAvgPrice"`
	FilledQuote    float64 `json:"filledQuote" bson:"filledQuote"`

	// calculate as raw order or popular by code
	FeeTier float64 `json:"feeTier" bson:"feeTier"`

	RawOrder any `json:"rawOrder" bson:"rawOrder"`

	Err OrderErr `json:"err" bson:"err"`
}
