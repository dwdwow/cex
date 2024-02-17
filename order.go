package cex

type OrderType string

const (
	OrderTypeLimit  OrderType = "LIMIT"
	OrderTypeMarket OrderType = "MARKET"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

type OrderStatus string

const (
	OrderStatusNew             OrderStatus = "NEW"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusRejected        OrderStatus = "REJECTED"
	OrderStatusExpired         OrderStatus = "EXPIRED"
)

// Order
// Every field value in Order must be certain.
type Order struct {
	// popular by user or code
	Cex       Name      `json:"cex" bson:"cex"`
	PairType  PairType  `json:"pairType" bson:"pairType"`
	OrderType OrderType `json:"orderType" bson:"orderType"`
	OrderSide OrderSide `json:"orderSide" bson:"orderSide"`

	// popular by code
	Symbol        string `json:"symbol" bson:"symbol"`
	TimeInForce   string `json:"timeInForce" bson:"timeInForce"`
	ClientOrderId string `json:"clientOrderId" bson:"clientOrderId"`
	ApiKey        string `json:"apiKey" bson:"apiKey"`

	// popular by user self
	OriQty   float64 `json:"oriQty" bson:"oriQty"`
	OriPrice float64 `json:"oriPrice" bson:"oriPrice"`

	// popular as response
	OrderId string      `json:"orderId" bson:"orderId"`
	Status  OrderStatus `json:"status" bson:"status"`

	// popular a order result
	FilledQty      float64 `json:"filledQty" bson:"filledQty"`
	FilledQuote    float64 `json:"filledQuote" bson:"filledQuote"`
	FilledAvgPrice float64 `json:"filledAvgPrice" bson:"filledAvgPrice"`

	RawOrder any `json:"rawOrder" bson:"rawOrder"`

	//Asset    string  `json:"asset" bson:"asset"`
	//Quote    string  `json:"quote" bson:"quote"`
}

func (o *Order) IsFinished() bool {
	if o == nil {
		return false
	}
	switch o.Status {
	case OrderStatusRejected, OrderStatusExpired, OrderStatusFilled, OrderStatusCanceled:
		return true
	}
	return false
}
