package bnc

import (
	"fmt"
	"time"
)

type WsSpotAggTradeStream struct {
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

type WsSpotTradeStream struct {
	EventType WsEvent `json:"e"`
	EventTime int64   `json:"E"`
	Symbol    string  `json:"s"`
	TradeId   int64   `json:"t"`
	Price     float64 `json:"p,string"`
	Qty       float64 `json:"q,string"`
	TradeTime int64   `json:"T"`
	IsMaker   bool    `json:"m"`
	M         bool    `json:"M"` // ignore
}

type WsSpotKline struct {
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
	QuoteVolume           string        `json:"q"`
	BaseAssetTakerVolume  string        `json:"V"`
	QuoteAssetTakerVolume string        `json:"Q"`
	B                     string        `json:"B"` // ignore
}

type WsSpotKlineStream struct {
	EventType WsEvent     `json:"e"`
	EventTime int64       `json:"E"`
	Symbol    string      `json:"s"`
	Kline     WsSpotKline `json:"k"`
}

type WsSpotDepthStream WsDepthMsg

func SpotWsPublicMsgUnmarshaler(e WsEvent, data []byte) (any, error) {
	switch e {
	case WsEventAggTrade:
		return unmarshal[WsSpotAggTradeStream](data)
	case WsEventTrade:
		return unmarshal[WsSpotTradeStream](data)
	case WsEventKline:
		return unmarshal[WsSpotKlineStream](data)
	case WsEventDepthUpdate:
		return unmarshal[WsSpotDepthStream](data)
	default:
		return nil, fmt.Errorf("bnc: unknown event %v", e)
	}
}

var SpotPublicWsCfg = WsCfg{
	Url:             WsBaseUrl,
	MaxStream:       1024,
	ReqDur:          time.Second,
	MaxReqPerDur:    5,
	DataUnmarshaler: SpotWsPrivateMsgUnmarshaler,
}
