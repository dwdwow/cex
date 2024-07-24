package bnc

import (
	"fmt"
	"time"
)

func SpotWsPublicMsgUnmarshaler(e WsEvent, data []byte) (any, error) {
	switch e {
	case WsEventAggTrade:
		return unmarshal[WsAggTradeStream](data)
	case WsEventTrade:
		return unmarshal[WsTradeStream](data)
	case WsEventKline:
		return unmarshal[WsKlineStream](data)
	case WsEventDepthUpdate:
		return unmarshal[WsDepthStream](data)
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

func UmFuturesWsPublicMsgUnmarshaler(e WsEvent, data []byte) (any, error) {
	switch e {
	case WsEventAggTrade:
		return unmarshal[WsAggTradeStream](data)
	case WsEventMarkPriceUpdate:
		return unmarshal[WsMarkPriceUpdate](data)
	case WsEventKline:
		return unmarshal[WsKlineStream](data)
	case WsEventDepthUpdate:
		return unmarshal[WsDepthStream](data)
	default:
		return nil, fmt.Errorf("bnc: unknown event %v", e)
	}
}

var UmFuturesWsCfg = WsCfg{
	Url:             FutureWsBaseUrl,
	MaxStream:       1024,
	ReqDur:          time.Second,
	MaxReqPerDur:    5,
	DataUnmarshaler: UmFuturesWsPublicMsgUnmarshaler,
}
