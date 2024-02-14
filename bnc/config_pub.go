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
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`

	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type OrderBook struct {
	LastUpdateId int64
	Bids         [][]float64
	Asks         [][]float64
}

var OrderBookConfig = cex.ReqConfig[OrderBookParams, OrderBook]{
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
