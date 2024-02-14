package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type OrderBookParams struct {
	Symbol string `s2m:"symbol,omitempty"`
	Limit  int    `s2m:"limit,omitempty"` // default 100, max 5000
}

type RawOrderBook struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	Asks         [][]string `json:"asks"`
	Bids         [][]string `json:"bids"`

	// futures order book fields
	E int64 `json:"e"` // Message output time
	T int64 `json:"t"` // Transaction time

	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OrderBook struct {
	LastUpdateId int64
	Asks         [][]float64
	Bids         [][]float64

	// futures order book fields
	E int64 `json:"e"` // Message output time
	T int64 `json:"t"` // Transaction time
}

var SpotOrderBookConfig = cex.ReqConfig[OrderBookParams, OrderBook]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/depth",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   obBodyUnmsher,
}

var FuturesOrderBookConfig = cex.ReqConfig[OrderBookParams, OrderBook]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/depth",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   obBodyUnmsher,
}
