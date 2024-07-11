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

//type WsSpotTradeStream struct {
//	EventType WsEvent `json:"e"`
//	EventTime int64   `json:"E"`
//	Symbol    string  `json:"s"`
//	TradeId   int64   `json:"t"`
//	Price     float64 `json:"p,string"`
//	Qty       float64 `json:"q,string"`
//	TradeTime int64   `json:"T"`
//	IsMaker   bool    `json:"m"`
//	M         bool    `json:"M"` // ignore
//}

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
