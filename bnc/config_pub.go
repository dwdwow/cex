package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/ob"
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
	Asks         ob.Book
	Bids         ob.Book

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

type KlineInterval string

const (
	// KlineInterval1s is only for spot
	KlineInterval1s = "1s"
	KlineInterval1m = "1m"
	KlineInterval1h = "1h"
)

type KlineParams struct {
	Symbol    string        `s2m:"symbol,omitempty"`
	Interval  KlineInterval `s2m:"interval,omitempty"`
	StartTime int64         `s2m:"startTime,omitempty"`
	EndTime   int64         `s2m:"endTime,omitempty"`
	// TimeZone
	// Hours and minutes (e.g. -1:00, 05:45)
	// Only hours (e.g. 0, 8, 4)
	// Accepted range is strictly [-12:00 to +14:00] inclusive
	TimeZone string `s2m:"timeZone,omitempty"`
	Limit    int64  `s2m:"limit,omitempty"`
}

type Kline struct {
	OpenTime                 int64   `json:"openTime,string" bson:"openTime,string"`
	CloseTime                int64   `json:"closeTime,string" bson:"closeTime,string"`
	TradesNumber             int64   `json:"tradesNumber,string" bson:"tradesNumber,string"`
	OpenPrice                float64 `json:"openPrice,string" bson:"openPrice,string"`
	HighPrice                float64 `json:"highPrice,string" bson:"highPrice,string"`
	LowPrice                 float64 `json:"lowPrice,string" bson:"lowPrice,string"`
	ClosePrice               float64 `json:"closePrice,string" bson:"closePrice,string"`
	Volume                   float64 `json:"volume,string" bson:"volume,string"`
	QuoteAssetVolume         float64 `json:"quoteAssetVolume,string" bson:"quoteAssetVolume,string"`
	TakerBuyBaseAssetVolume  float64 `json:"takerBuyBaseAssetVolume,string" bson:"takerBuyBaseAssetVolume,string"`
	TakerBuyQuoteAssetVolume float64 `json:"takerBuyQuoteAssetVolume,string" bson:"takerBuyQuoteAssetVolume,string"`
	Unused                   any     `json:"unused" bson:"unused"`
}

var SpotKlineConfig = cex.ReqConfig[KlineParams, []Kline]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/klines",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(klineBodyUnmsher),
}

var FuturesKlineConfig = cex.ReqConfig[KlineParams, []Kline]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/klines",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(klineBodyUnmsher),
}

type SpotPriceTicker struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

var SpotPricesConfig = cex.ReqConfig[cex.NilReqData, []SpotPriceTicker]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/ticker/price",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]SpotPriceTicker]),
}

type FuturesPriceTicker struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Time   int64   `json:"time,omitempty"`
}

var FuturesPricesConfig = cex.ReqConfig[cex.NilReqData, []FuturesPriceTicker]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV2 + "/ticker/price",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesPriceTicker]),
}

type CMPremiumIndex struct {
	Symbol               string  `json:"symbol"`
	Pair                 string  `json:"pair"`
	MarkPrice            float64 `json:"markPrice,string"`
	IndexPrice           float64 `json:"indexPrice,string"`
	EstimatedSettlePrice float64 `json:"estimatedSettlePrice,string"`
	LastFundingRate      string  `json:"lastFundingRate"`
	InterestRate         string  `json:"interestRate"`
	NextFundingTime      int64   `json:"nextFundingTime"`
	Time                 int64   `json:"time"`
}

type CMPremiumIndexParams struct {
	Symbol string `s2m:"symbol,omitempty"`
	Pair   string `s2m:"pair,omitempty"`
}

var CMPremiumIndexConfig = cex.ReqConfig[CMPremiumIndexParams, []CMPremiumIndex]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          DapiBaseUrl,
		Path:             DapiV1 + "/premiumIndex",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]CMPremiumIndex]),
}
