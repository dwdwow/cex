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
	LastUpdateId int64      `json:"lastUpdateId" bson:"lastUpdateId"`
	Asks         [][]string `json:"asks" bson:"asks"`
	Bids         [][]string `json:"bids" bson:"bids"`

	// futures order book fields
	E int64 `json:"e" bson:"e"` // Message output time
	T int64 `json:"t" bson:"t"` // Transaction time

	Code int    `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
}

type OrderBook struct {
	LastUpdateId int64
	Asks         [][]float64
	Bids         [][]float64

	// futures order book fields
	E int64 `json:"e" bson:"e"` // Message output time
	T int64 `json:"t" bson:"t"` // Transaction time
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
	RateLimitType string `json:"rateLimitType" bson:"rateLimitType"` // ORDERS REQUEST_WEIGHT
	Interval      string `json:"interval" bson:"interval"`           // SECOND MINUTE
	IntervalNum   int    `json:"intervalNum" bson:"intervalNum"`
	Limit         int    `json:"limit" bson:"limit"`
	// just for spot
	Count int `json:"count" bson:"count"`
}

type Exchange struct {
	Symbol                          string           `json:"symbol" bson:"symbol"`
	Status                          ExchangeStatus   `json:"status" bson:"status"`
	BaseAsset                       string           `json:"baseAsset" bson:"baseAsset"`
	BaseAssetPrecision              int64            `json:"baseAssetPrecision" bson:"baseAssetPrecision"`
	QuoteAsset                      string           `json:"quoteAsset" bson:"quoteAsset"`
	QuotePrecision                  int64            `json:"quotePrecision" bson:"quotePrecision"`
	QuoteAssetPrecision             int64            `json:"quoteAssetPrecision" bson:"quoteAssetPrecision"`
	OrderTypes                      []OrderType      `json:"orderTypes" bson:"orderTypes"`
	IcebergAllowed                  bool             `json:"icebergAllowed" bson:"icebergAllowed"`
	OcoAllowed                      bool             `json:"ocoAllowed" bson:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed      bool             `json:"quoteOrderQtyMarketAllowed" bson:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool             `json:"allowTrailingStop" bson:"allowTrailingStop"`
	CancelReplaceAllowed            bool             `json:"cancelReplaceAllowed" bson:"cancelReplaceAllowed"`
	IsSpotTradingAllowed            bool             `json:"isSpotTradingAllowed" bson:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool             `json:"isMarginTradingAllowed" bson:"isMarginTradingAllowed"`
	Filters                         []map[string]any `json:"filters" bson:"filters"`
	Permissions                     []PairType       `json:"permissions" bson:"permissions"`
	DefaultSelfTradePreventionMode  string           `json:"defaultSelfTradePreventionMode" bson:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string         `json:"allowedSelfTradePreventionModes" bson:"allowedSelfTradePreventionModes"`

	// just for future pair
	Pair              string   `json:"pair" bson:"pair"`
	ContractType      string   `json:"contractType" bson:"contractType"`
	DeliveryData      int64    `json:"deliveryData" bson:"deliveryData"`
	OnboardDate       int64    `json:"onboardDate" bson:"onboardDate"`
	MarginAsset       string   `json:"marginAsset" bson:"marginAsset"`
	UnderlyingType    string   `json:"underlyingType" bson:"underlyingType"`
	UnderlyingSubType []string `json:"underlyingSubType" bson:"underlyingSubType"`
	SettlePlan        int      `json:"settlePlan" bson:"settlePlan"`
}

type FuturesExchangeInfoAsset struct {
	Asset           string `json:"asset" bson:"asset"`
	MarginAvailable bool   `json:"marginAvailable" bson:"marginAvailable"` // whether the asset can be used as margin in Multi-Assets mode
	// binance doc show that AutoAssetExchange can be int or null...
	AutoAssetExchange any `json:"autoAssetExchange" bson:"autoAssetExchange"` // auto-exchange threshold in Multi-Assets margin mode
}

type ExchangeInfo struct {
	Timezone        string              `json:"timezone" bson:"timezone"`
	ServerTime      int64               `json:"serverTime" bson:"serverTime"`
	RateLimits      []ExchangeRateLimit `json:"rateLimits" bson:"rateLimits"`
	ExchangeFilters []map[string]string `json:"exchangeFilters" bson:"exchangeFilters"`
	Symbols         []Exchange          `json:"symbols" bson:"symbols"`

	// just for futures
	Assets []FuturesExchangeInfoAsset `json:"assets" bson:"assets"`
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

// FuturesFundingRateHistoriesParams
// Limit, default 100, max 1000
type FuturesFundingRateHistoriesParams struct {
	Symbol    string `s2m:"symbol,omitempty"`
	StartTime int64  `s2m:"startTime,omitempty"` // Timestamp in ms to get funding rate from INCLUSIVE.
	EndTime   int64  `s2m:"endTime,omitempty"`   // Timestamp in ms to get funding rate until INCLUSIVE.
	Limit     int    `s2m:"limit,omitempty"`     // Default 100; max 1000
}

type FuturesFundingRateHistory struct {
	Symbol      string  `json:"symbol" bson:"symbol"`
	FundingTime int64   `json:"fundingTime" bson:"fundingTime"`
	FundingRate float64 `json:"fundingRate,string" bson:"fundingRate,string"`
	MarkPrice   string  `json:"markPrice" bson:"markPrice"` // mark price maybe empty string
}

var FuturesFundingRateHistoriesConfig = cex.ReqConfig[FuturesFundingRateHistoriesParams, []FuturesFundingRateHistory]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/fundingRate",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesFundingRateHistory]),
}

type FuturesFundingRateInfo struct {
	Symbol                   string  `json:"symbol" bson:"symbol"`
	AdjustedFundingRateCap   float64 `json:"adjustedFundingRateCap,string" bson:"adjustedFundingRateCap,string"`
	AdjustedFundingRateFloor float64 `json:"adjustedFundingRateFloor,string" bson:"adjustedFundingRateFloor,string"`
	FundingIntervalHours     float64 `json:"fundingIntervalHours" bson:"fundingIntervalHours"`
	Disclaimer               bool    `json:"disclaimer" bson:"disclaimer"` // ignore
}

// FuturesFundingRateInfosConfig
// Query funding rate info for symbols that had FundingRateCap/ FundingRateFloor / fundingIntervalHours adjustment
// Be careful! Some symbols have no funding rate info!!!
var FuturesFundingRateInfosConfig = cex.ReqConfig[cex.NilReqData, []FuturesFundingRateInfo]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/fundingInfo",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesFundingRateInfo]),
}

type FuturesFundingRatesParams struct {
	Symbol string `s2m:"symbol"`
}

type FuturesFundingRate struct {
	Symbol               string  `json:"symbol" bson:"symbol"`
	MarkPrice            float64 `json:"markPrice,string" bson:"markPrice,string"`
	IndexPrice           float64 `json:"indexPrice,string" bson:"indexPrice,string"`
	EstimatedSettlePrice float64 `json:"estimatedSettlePrice,string" bson:"estimatedSettlePrice,string"`
	LastFundingRate      float64 `json:"lastFundingRate,string" bson:"lastFundingRate,string"` // This is the Latest funding rate
	NextFundingTime      int64   `json:"nextFundingTime" bson:"nextFundingTime"`
	InterestRate         float64 `json:"interestRate,string" bson:"interestRate,string"`
	Time                 int64   `json:"time" bson:"time"`
}

var FuturesFundingRatesConfig = cex.ReqConfig[FuturesFundingRatesParams, []FuturesFundingRate]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/premiumIndex",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesFundingRate]),
}
