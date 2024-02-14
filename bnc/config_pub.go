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

type ExchangeRateLimit struct {
	RateLimitType string `json:"rateLimitType"` // ORDERS REQUEST_WEIGHT
	Interval      string `json:"interval"`      // SECOND MINUTE
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	// just for spot
	Count int `json:"count"`
}

type Exchange struct {
	Symbol                          string           `json:"symbol"`
	Status                          ExchangeStatus   `json:"status"`
	BaseAsset                       string           `json:"baseAsset"`
	BaseAssetPrecision              int64            `json:"baseAssetPrecision"`
	QuoteAsset                      string           `json:"quoteAsset"`
	QuotePrecision                  int64            `json:"quotePrecision"`
	QuoteAssetPrecision             int64            `json:"quoteAssetPrecision"`
	OrderTypes                      []OrderType      `json:"orderTypes"`
	IcebergAllowed                  bool             `json:"icebergAllowed"`
	OcoAllowed                      bool             `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed      bool             `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool             `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool             `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed            bool             `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool             `json:"isMarginTradingAllowed"`
	Filters                         []map[string]any `json:"filters"`
	Permissions                     []TradeType      `json:"permissions"`
	DefaultSelfTradePreventionMode  string           `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string         `json:"allowedSelfTradePreventionModes"`

	// just for future pair
	Pair              string   `json:"pair"`
	ContractType      string   `json:"contractType"`
	DeliveryData      int64    `json:"deliveryData"`
	OnboardDate       int64    `json:"onboardDate"`
	MarginAsset       string   `json:"marginAsset"`
	UnderlyingType    string   `json:"underlyingType"`
	UnderlyingSubType []string `json:"underlyingSubType"`
	SettlePlan        int      `json:"settlePlan"`
}

type FuturesExchangeInfoAsset struct {
	Asset           string `json:"asset"`
	MarginAvailable bool   `json:"marginAvailable"` // whether the asset can be used as margin in Multi-Assets mode
	// binance doc show that AutoAssetExchange can be int or null...
	AutoAssetExchange any `json:"autoAssetExchange"` // auto-exchange threshold in Multi-Assets margin mode
}

type ExchangeInfo struct {
	Timezone        string              `json:"timezone"`
	ServerTime      int64               `json:"serverTime"`
	RateLimits      []ExchangeRateLimit `json:"rateLimits"`
	ExchangeFilters []map[string]string `json:"exchangeFilters"`
	Symbols         []Exchange          `json:"symbols"`

	// just for futures
	Assets []FuturesExchangeInfoAsset `json:"assets"`
}

var SpotExchangeInfosConfig = cex.ReqConfig[cex.NilReqData, ExchangeInfo]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/exchangeInfo",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[ExchangeInfo]),
}

var FuturesExchangeInfosConfig = cex.ReqConfig[cex.NilReqData, ExchangeInfo]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/exchangeInfo",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[ExchangeInfo]),
}
