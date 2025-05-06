package bnc

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/ob"
	"github.com/dwdwow/mathy"
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
	DeliveryDate      int64    `json:"deliveryDate"`
	OnboardDate       int64    `json:"onboardDate" bson:"onboardDate"`
	MarginAsset       string   `json:"marginAsset" bson:"marginAsset"`
	UnderlyingType    string   `json:"underlyingType" bson:"underlyingType"`
	UnderlyingSubType []string `json:"underlyingSubType" bson:"underlyingSubType"`
	SettlePlan        int64    `json:"settlePlan" bson:"settlePlan"`

	// just for cm futures
	OrderType             []string       `json:"OrderType"`
	TimeInForce           []string       `json:"timeInForce"`
	LiquidationFee        string         `json:"liquidationFee"`
	MarketTakeBound       string         `json:"marketTakeBound"`
	ContractStatus        ExchangeStatus `json:"contractStatus"`
	ContractSize          float64        `json:"contractSize"`
	PricePrecision        int64          `json:"pricePrecision"`
	QuantityPrecision     int64          `json:"quantityPrecision"`
	EqualQtyPrecision     int64          `json:"equalQtyPrecision"`
	TriggerProtect        string         `json:"triggerProtect"`
	MaintMarginPercent    string         `json:"maintMarginPercent"`
	RequiredMarginPercent string         `json:"requiredMarginPercent"`
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

var CMFuturesExchangeInfosConfig = cex.ReqConfig[cex.NilReqData, ExchangeInfo]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          DapiBaseUrl,
		Path:             DapiV1 + "/exchangeInfo",
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
	KlineInterval1s  KlineInterval = "1s"
	KlineInterval1m  KlineInterval = "1m"
	KlineInterval3m  KlineInterval = "3m"
	KlineInterval5m  KlineInterval = "5m"
	KlineInterval15m KlineInterval = "15m"
	KlineInterval30m KlineInterval = "30m"
	KlineInterval1h  KlineInterval = "1h"
	KlineInterval2h  KlineInterval = "2h"
	KlineInterval4h  KlineInterval = "4h"
	KlineInterval6h  KlineInterval = "6h"
	KlineInterval8h  KlineInterval = "8h"
	KlineInterval12h KlineInterval = "12h"
	KlineInterval1d  KlineInterval = "1d"
	KlineInterval3d  KlineInterval = "3d"
	KlineInterval1w  KlineInterval = "1w"
	KlineInterval1M  KlineInterval = "1M"
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

func (k *Kline) CSVRow() string {
	cells := []string{
		strconv.FormatInt(k.OpenTime, 10),
		mathy.BN(k.OpenPrice).Round(20).String(),
		mathy.BN(k.HighPrice).Round(20).String(),
		mathy.BN(k.LowPrice).Round(20).String(),
		mathy.BN(k.ClosePrice).Round(20).String(),
		mathy.BN(k.Volume).Round(8).String(),
		strconv.FormatInt(k.CloseTime, 10),
		mathy.BN(k.QuoteAssetVolume).Round(8).String(),
		strconv.FormatInt(k.TradesNumber, 10),
		mathy.BN(k.TakerBuyBaseAssetVolume).Round(8).String(),
		mathy.BN(k.TakerBuyQuoteAssetVolume).Round(8).String(),
		"unused",
	}
	return strings.Join(cells, ",")
}

func (k *Kline) CSVRowWithPriceRoundPlaces(roundPlaces int32) string {
	cells := []string{
		strconv.FormatInt(k.OpenTime, 10),
		mathy.BN(k.OpenPrice).Round(roundPlaces).String(),
		mathy.BN(k.HighPrice).Round(roundPlaces).String(),
		mathy.BN(k.LowPrice).Round(roundPlaces).String(),
		mathy.BN(k.ClosePrice).Round(roundPlaces).String(),
		mathy.BN(k.Volume).Round(8).String(),
		strconv.FormatInt(k.CloseTime, 10),
		mathy.BN(k.QuoteAssetVolume).Round(8).String(),
		strconv.FormatInt(k.TradesNumber, 10),
		mathy.BN(k.TakerBuyBaseAssetVolume).Round(8).String(),
		mathy.BN(k.TakerBuyQuoteAssetVolume).Round(8).String(),
		"unused",
	}
	return strings.Join(cells, ",")
}

func (k *Kline) FromCSVRow(row string) error {
	raw := strings.Split(row, ",")
	if len(raw) < 12 {
		return errors.New("invalid kline csv raw")
	}
	var err error
	k.OpenTime, err = strconv.ParseInt(raw[0], 10, 64)
	if err != nil {
		return err
	}
	k.OpenPrice, err = strconv.ParseFloat(raw[1], 64)
	if err != nil {
		return err
	}
	k.HighPrice, err = strconv.ParseFloat(raw[2], 64)
	if err != nil {
		return err
	}
	k.LowPrice, err = strconv.ParseFloat(raw[3], 64)
	if err != nil {
		return err
	}
	k.ClosePrice, err = strconv.ParseFloat(raw[4], 64)
	if err != nil {
		return err
	}
	k.Volume, err = strconv.ParseFloat(raw[5], 64)
	if err != nil {
		return err
	}
	k.CloseTime, err = strconv.ParseInt(raw[6], 10, 64)
	if err != nil {
		return err
	}
	k.QuoteAssetVolume, err = strconv.ParseFloat(raw[7], 64)
	if err != nil {
		return err
	}
	k.TradesNumber, err = strconv.ParseInt(raw[8], 10, 64)
	if err != nil {
		return err
	}
	k.TakerBuyBaseAssetVolume, err = strconv.ParseFloat(raw[9], 64)
	if err != nil {
		return err
	}
	k.TakerBuyQuoteAssetVolume, err = strconv.ParseFloat(raw[10], 64)
	if err != nil {
		return err
	}
	return nil
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

type SpotTrade struct {
	Id           int64   `json:"id"`
	Price        float64 `json:"price,string"`
	Qty          float64 `json:"qty,string"`
	QuoteQty     float64 `json:"quoteQty,string"`
	Time         int64   `json:"time"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	IsBestMatch  bool    `json:"isBestMatch"`
}

type SpotTradesParams struct {
	Symbol string `s2m:"symbol,omitempty"`
	Limit  int    `s2m:"limit,omitempty"` // default 500; max 1000
}

var SpotTradesConfig = cex.ReqConfig[SpotTradesParams, []SpotTrade]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/trades",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]SpotTrade]),
}

type AggTrades struct {
	Id           int64   `json:"a"`
	Price        float64 `json:"p,string"`
	Qty          float64 `json:"q,string"`
	FirstTradeId int64   `json:"f"`
	LastTradeId  int64   `json:"l"`
	Time         int64   `json:"T"`
	IsBuyerMaker bool    `json:"m"`
	IsBestMatch  bool    `json:"M"`
}

func (a *AggTrades) CSVRow() string {
	cells := []string{
		strconv.FormatInt(a.Id, 10),
		mathy.BN(a.Price).Round(20).String(),
		mathy.BN(a.Qty).Round(20).String(),
		strconv.FormatInt(a.FirstTradeId, 10),
		strconv.FormatInt(a.LastTradeId, 10),
		strconv.FormatInt(a.Time, 10),
		strconv.FormatBool(a.IsBuyerMaker),
		strconv.FormatBool(a.IsBestMatch),
	}
	return strings.Join(cells, ",")
}

func (a *AggTrades) FromCSVRow(row string) error {
	raw := strings.Split(row, ",")
	if len(raw) < 7 {
		return errors.New("invalid agg trades csv raw")
	}
	var err error
	a.Id, err = strconv.ParseInt(raw[0], 10, 64)
	if err != nil {
		return err
	}
	a.Price, err = strconv.ParseFloat(raw[1], 64)
	if err != nil {
		return err
	}
	a.Qty, err = strconv.ParseFloat(raw[2], 64)
	if err != nil {
		return err
	}
	a.FirstTradeId, err = strconv.ParseInt(raw[3], 10, 64)
	if err != nil {
		return err
	}
	a.LastTradeId, err = strconv.ParseInt(raw[4], 10, 64)
	if err != nil {
		return err
	}
	a.Time, err = strconv.ParseInt(raw[5], 10, 64)
	if err != nil {
		return err
	}
	a.IsBuyerMaker, err = strconv.ParseBool(raw[6])
	if err != nil {
		return err
	}
	if len(raw) > 7 {
		a.IsBestMatch, err = strconv.ParseBool(raw[7])
		if err != nil {
			return err
		}
	}
	return nil
}

type AggTradesParams struct {
	Symbol    string `s2m:"symbol,omitempty"`
	FromId    int64  `s2m:"fromId,omitempty"`
	StartTime int64  `s2m:"startTime,omitempty"`
	EndTime   int64  `s2m:"endTime,omitempty"`
	Limit     int    `s2m:"limit,omitempty"` // default 500; max 1000
}

type SpotAggTrades AggTrades

type SpotAggTradesParams AggTradesParams

var SpotAggTradesConfig = cex.ReqConfig[SpotAggTradesParams, []SpotAggTrades]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/aggTrades",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]SpotAggTrades]),
}

type UmFuturesAggTradesParams AggTradesParams

type UmFuturesAggTrades AggTrades

var UmFuturesAggTradesConfig = cex.ReqConfig[UmFuturesAggTradesParams, []UmFuturesAggTrades]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/aggTrades",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]UmFuturesAggTrades]),
}

type UmOpenInterestParams struct {
	Symbol string `s2m:"symbol"`
}

type UmOpenInterest struct {
	Symbol       string  `json:"symbol"`
	OpenInterest float64 `json:"openInterest,string"`
	Time         int64   `json:"time"`
}

var UmOpenInterestConfig = cex.ReqConfig[UmOpenInterestParams, UmOpenInterest]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl: FapiBaseUrl,
		Path:    FapiV1 + "/openInterest",
		Method:  http.MethodGet,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[UmOpenInterest]),
}

type UmOpenInterestStatisticsParams struct {
	Symbol    string              `s2m:"symbol"`
	Period    FuturesStaticPeriod `s2m:"period"`
	Limit     int64               `s2m:"limit,omitempty"`
	StartTime int64               `s2m:"startTime,omitempty"`
	EndTime   int64               `s2m:"endTime,omitempty"`
}

type UmOpenInterestStatistics struct {
	Symbol               string  `json:"symbol"`
	SumOpenInterest      float64 `json:"sumOpenInterest,string"`
	SumOpenInterestValue float64 `json:"sumOpenInterestValue,string"`
	Timestamp            int64   `json:"timestamp"`
}

var UmOpenInterestStatisticsConfig = cex.ReqConfig[UmOpenInterestStatisticsParams, []UmOpenInterestStatistics]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl: FapiBaseUrl,
		Path:    FuturesData + "/openInterestHist",
		Method:  http.MethodGet,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]UmOpenInterestStatistics]),
}

type UmLongShortRatioParams struct {
	Symbol    string              `s2m:"symbol"`
	Period    FuturesStaticPeriod `s2m:"period"`
	Limit     int64               `s2m:"limit,omitempty"`
	StartTime int64               `s2m:"startTime,omitempty"`
	EndTime   int64               `s2m:"endTime,omitempty"`
}

type UmLongShortRatio struct {
	Symbol         string  `json:"symbol"`
	LongShortRatio float64 `json:"longShortRatio,string"`
	LongAccount    float64 `json:"longAccount,string"`
	ShortAccount   float64 `json:"shortAccount,string"`
	Timestamp      int64   `json:"timestamp"`
}

var UmTopLongShortPositionRatioConfig = cex.ReqConfig[UmLongShortRatioParams, []UmLongShortRatio]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FuturesData + "/topLongShortPositionRatio",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]UmLongShortRatio]),
}

var UmTopLongShortAccountRatioConfig = cex.ReqConfig[UmLongShortRatioParams, []UmLongShortRatio]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl: FapiBaseUrl,
		Path:    FuturesData + "/topLongShortAccountRatio",
		Method:  http.MethodGet,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]UmLongShortRatio]),
}

var UmGlobalLongShortAccountRatioConfig = cex.ReqConfig[UmLongShortRatioParams, []UmLongShortRatio]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FuturesData + "/globalLongShortAccountRatio",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]UmLongShortRatio]),
}

type Um24hrTicker struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstId            int64   `json:"firstId"`
	LastId             int64   `json:"lastId"`
	Count              int64   `json:"count"`
}

type Um24hrTickerParams struct {
	Symbol string `s2m:"symbol"`
}

var Um24hrTickerConfig = cex.ReqConfig[Um24hrTickerParams, Um24hrTicker]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl: FapiBaseUrl,
		Path:    FapiV1 + "/ticker/24hr",
		Method:  http.MethodGet,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[Um24hrTicker]),
}

var Um24hrTickersConfig = cex.ReqConfig[cex.NilReqData, []Um24hrTicker]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl: FapiBaseUrl,
		Path:    FapiV1 + "/ticker/24hr",
		Method:  http.MethodGet,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]Um24hrTicker]),
}
