package bnc

type WsAggTradeStream struct {
	EventType    WsEvent `json:"e"`
	EventTime    int64   `json:"E"`
	Symbol       string  `json:"s"`
	AggTradeId   int64   `json:"a"`
	Price        float64 `json:"p,string"`
	Qty          float64 `json:"q,string"`
	FirstTradeId int64   `json:"f"`
	LastTradeId  int64   `json:"l"`
	TradeTime    int64   `json:"T"`
	IsMaker      bool    `json:"m"`
	M            bool    `json:"M"` // ignore
}

type WsTradeStream struct {
	EventType    WsEvent `json:"e"`
	EventTime    int64   `json:"E"`
	Symbol       string  `json:"s"`
	TradeID      int64   `json:"t"`
	Price        float64 `json:"p,string"`
	Quantity     float64 `json:"q,string"`
	TradeTime    int64   `json:"T"`
	IsBuyerMaker bool    `json:"m"`
	M            bool    `json:"M"` // ignore

	//
	BuyerOrderID  int64 `json:"b"`
	SellerOrderID int64 `json:"a"`
}

type WsKline struct {
	StartTime             int64         `json:"t"`
	CloseTime             int64         `json:"T"`
	Symbol                string        `json:"s"`
	Interval              KlineInterval `json:"i"`
	FirstTradeId          int64         `json:"f"`
	LastTradeId           int64         `json:"L"`
	OpenPrice             float64       `json:"o,string"`
	ClosePrice            float64       `json:"c,string"`
	HighPrice             float64       `json:"h,string"`
	LowPrice              float64       `json:"l,string"`
	BaseAssetVolume       float64       `json:"v,string"`
	TradesNum             int64         `json:"n"`
	IsKlineClosed         bool          `json:"x"`
	QuoteVolume           float64       `json:"q,string"`
	BaseAssetTakerVolume  float64       `json:"V,string"`
	QuoteAssetTakerVolume float64       `json:"Q,string"`
	B                     string        `json:"B"` // ignore
}

type WsKlineStream struct {
	EventType WsEvent `json:"e"`
	EventTime int64   `json:"E"`
	Symbol    string  `json:"s"`
	Kline     WsKline `json:"k"`
}

type WsDepthStream WsDepthMsg

type WsOrderExecutionReport struct {
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

	// just for margin order
	TrailingDelta      float64 `json:"d"` // Trailing Delta; This is only visible if the order was a trailing stop order.
	TrailingTime       int64   `json:"D"` // Trailing Time; This is only visible if the trailing stop order has been activated.
	MarginStrategyId   int64   `json:"j"`
	MarginStrategyType int64   `json:"J"`
	TradeGroupId       int64   `json:"u"`
	CounterOrderId     int64   `json:"U"`
	PreventedQty       float64 `json:"A,string"`
	LastPreventedQty   float64 `json:"B,string"`

	// just for futures order
	AvgPrice            string              `json:"ap"`
	Sp                  string              `json:"sp"` // ignore
	BidNotional         string              `json:"b"`
	AskNotional         string              `json:"a"`
	IsReduceOnly        bool                `json:"R"`
	PositionSide        FuturesPositionSide `json:"ps"`
	RealizedProfit      string              `json:"rp"`
	FuturesStrategyType string              `json:"st"`
	FuturesStrategyId   int64               `json:"si"`
	Gtd                 int64               `json:"gtd"`
}

type WsListenKeyExpired struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	ListenKey string `json:"listenKey"`
}
