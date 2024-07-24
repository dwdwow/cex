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
	DataUnmarshaler: SpotWsPublicMsgUnmarshaler,
}

func UmFuturesWsPublicMsgUnmarshaler(e WsEvent, data []byte) (any, error) {
	switch e {
	case WsEventAggTrade:
		return unmarshal[WsAggTradeStream](data)
	case WsEventMarkPriceUpdate:
		return unmarshal[WsMarkPriceStream](data)
	case WsEventForceOrder:
		return unmarshal[WsLiquidationOrderStream](data)
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

func CmFuturesWsPublicMsgUnmarshaler(e WsEvent, data []byte) (any, error) {
	switch e {
	case WsEventAggTrade:
		return unmarshal[WsAggTradeStream](data)
	case WsEventIndexPriceUpdate:
		return unmarshal[WsCMIndexPriceStream](data)
	case WsEventMarkPriceUpdate:
		return unmarshal[WsMarkPriceStream](data)
	case WsEventForceOrder:
		return unmarshal[WsLiquidationOrderStream](data)
	case WsEventKline:
		return unmarshal[WsKlineStream](data)
	case WsEventDepthUpdate:
		return unmarshal[WsDepthStream](data)
	default:
		return nil, fmt.Errorf("bnc: unknown event %v", e)
	}
}

var CmFuturesWsCfg = WsCfg{
	Url:             CMFutureWsBaseUrl,
	MaxStream:       1024,
	ReqDur:          time.Second,
	MaxReqPerDur:    5,
	DataUnmarshaler: CmFuturesWsPublicMsgUnmarshaler,
}
