package bnc

import (
	"math"
	"net/http"

	"github.com/dwdwow/cex"
)

type FuturesChangePositionModParams struct {
	DualSidePosition SmallBool `s2m:"dualSidePosition"`
}

var FuturesChangePositionModeConfig = cex.ReqConfig[FuturesChangePositionModParams, CodeMsg]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/positionSide/dual",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[CodeMsg]),
}

type FuturesCurrentPositionModeResponse struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

var FuturesPositionModeConfig = cex.ReqConfig[cex.NilReqData, FuturesCurrentPositionModeResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/positionSide/dual",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesCurrentPositionModeResponse]),
}

type FuturesChangeMultiAssetsModeParams struct {
	MultiAssetsMargin SmallBool `s2m:"multiAssetsMargin"`
}

var FuturesChangeMultiAssetsModeConfig = cex.ReqConfig[FuturesChangeMultiAssetsModeParams, CodeMsg]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/multiAssetsMargin",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[CodeMsg]),
}

type FuCurrentMultiAssetsModeResponse struct {
	MultiAssetsMargin bool `json:"multiAssetsMargin"`
}

var FuturesCurrentMultiAssetsModeConfig = cex.ReqConfig[cex.NilReqData, FuCurrentMultiAssetsModeResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/multiAssetsMargin",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuCurrentMultiAssetsModeResponse]),
}

type FuturesNewOrderParams struct {
	Symbol                  string                  `s2m:"symbol,omitempty" json:"symbol,omitempty"`
	PositionSide            FuturesPositionSide     `s2m:"positionSide,omitempty" json:"positionSide,omitempty"`
	Type                    OrderType               `s2m:"type,omitempty" json:"type,omitempty"`
	Side                    OrderSide               `s2m:"side,omitempty" json:"side,omitempty"`
	Quantity                float64                 `s2m:"quantity,omitempty" json:"quantity,omitempty"`
	Price                   float64                 `s2m:"price,omitempty" json:"price,omitempty"`
	TimeInForce             TimeInForce             `s2m:"timeInForce,omitempty" json:"timeInForce,omitempty"`
	NewClientOrderId        string                  `s2m:"newClientOrderId,omitempty" json:"newClientOrderId,omitempty"`
	ReduceOnly              SmallBool               `s2m:"reduceOnly,omitempty" json:"reduceOnly,omitempty"`                           // "true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
	ClosePosition           bool                    `s2m:"closePosition,omitempty" json:"closePosition,omitempty"`                     //	true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
	StopPrice               float64                 `s2m:"stopPrice,omitempty" json:"stopPrice,omitempty"`                             // Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ActivationPrice         float64                 `s2m:"activationPrice,omitempty" json:"activationPrice,omitempty"`                 // Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)
	CallbackRate            float64                 `s2m:"callbackRate,omitempty" json:"callbackRate,omitempty"`                       // Used with TRAILING_STOP_MARKET orders, min 0.1, max 5 where 1 for 1%
	WorkingType             FuturesWorkingType      `s2m:"workingType,omitempty" json:"workingType,omitempty"`                         // stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE".Default "CONTRACT_PRICE"
	PriceProtect            BigBool                 `s2m:"priceProtect,omitempty" json:"priceProtect,omitempty"`                       // "TRUE" or "FALSE", default "FALSE".Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	NewOrderRespType        OrderResponseType       `s2m:"newOrderRespType,omitempty" json:"newOrderRespType,omitempty"`               // "ACK", "RESULT", default "ACK"
	PriceMatch              string                  `s2m:"priceMatch,omitempty" json:"priceMatch,omitempty"`                           //  only available for LIMIT/STOP/TAKE_PROFIT order, can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20. Can't be passed together with price
	SelfTradePreventionMode SelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty" json:"selfTradePreventionMode,omitempty"` // NONE:No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire maker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers , default NONE
	GoodTillDate            int64                   `s2m:"goodTillDate,omitempty" json:"goodTillDate,omitempty"`
}

type FuturesOrder struct {
	// common
	Symbol                  string                  `json:"symbol"`
	OrderId                 int64                   `json:"orderId"`
	ClientOrderId           string                  `json:"clientOrderId"`
	Type                    OrderType               `json:"type"`
	PositionSide            FuturesPositionSide     `json:"positionSide"`
	Side                    OrderSide               `json:"side"`
	OrigQty                 float64                 `json:"origQty,string"`
	Price                   float64                 `json:"price,string"` // orig price
	ExecutedQty             float64                 `json:"executedQty,string"`
	AvgPrice                float64                 `json:"avgPrice,string"`
	ReduceOnly              bool                    `json:"reduceOnly"`
	Status                  OrderStatus             `json:"status"`
	StopPrice               float64                 `json:"stopPrice,string"`
	ClosePosition           bool                    `json:"closePosition"`
	TimeInForce             TimeInForce             `json:"timeInForce"`
	OrigType                OrderType               `json:"origType"`
	UpdateTime              int64                   `json:"updateTime"`
	WorkingType             FuturesWorkingType      `json:"workingType"`
	PriceProtect            bool                    `json:"priceProtect"`
	PriceMatch              string                  `json:"priceMatch"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode"`
	GoodTillDate            int64                   `json:"goodTillDate"`

	// new order
	CumQuote      float64 `json:"cumQuote,string"`
	ActivatePrice float64 `json:"activatePrice,string"`
	PriceRate     float64 `json:"priceRate,string"`

	// new, modify order
	CumQty float64 `json:"cumQty,string"`

	// modify order
	Pair    string `json:"pair"`    // same as symbol
	CumBase string `json:"cumBase"` // same as CumQuote? should verify

	// query order
	Time int64 `json:"time"`

	// place, modify multi orders
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var FuturesNewOrderConfig = cex.ReqConfig[FuturesNewOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

type FuturesModifyOrderParams struct {
	OrderId           int64     `s2m:"orderId,omitempty"`
	OrigClientOrderId string    `s2m:"origClientOrderId,omitempty"`
	Symbol            string    `s2m:"symbol,omitempty"`
	Side              OrderSide `s2m:"side,omitempty"` // needs to be same as origin order
	Quantity          float64   `s2m:"quantity,omitempty"`
	Price             float64   `s2m:"price,omitempty"`
	PriceMatch        string    `s2m:"priceMatch,omitempty"`
}

var FuturesModifyOrderConfig = cex.ReqConfig[FuturesModifyOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodPut,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

// FuturesNewMultiOrdersOrderParams is different with FuturesNewOrderParams.
// All fields are string.
// Binance doc example just show that quantity and price are string.
// Do not know if other float/int fields are string or not.
type FuturesNewMultiOrdersOrderParams struct {
	Symbol                  string                  `s2m:"symbol,omitempty" json:"symbol,omitempty"`
	PositionSide            FuturesPositionSide     `s2m:"positionSide,omitempty" json:"positionSide,omitempty"`
	Type                    OrderType               `s2m:"type,omitempty" json:"type,omitempty"`
	Side                    OrderSide               `s2m:"side,omitempty" json:"side,omitempty"`
	Quantity                string                  `s2m:"quantity,omitempty" json:"quantity,omitempty"`
	Price                   string                  `s2m:"price,omitempty" json:"price,omitempty"`
	TimeInForce             TimeInForce             `s2m:"timeInForce,omitempty" json:"timeInForce,omitempty"`
	NewClientOrderId        string                  `s2m:"newClientOrderId,omitempty" json:"newClientOrderId,omitempty"`
	ReduceOnly              SmallBool               `s2m:"reduceOnly,omitempty" json:"reduceOnly,omitempty"`                           // "true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
	ClosePosition           bool                    `s2m:"closePosition,omitempty" json:"closePosition,omitempty"`                     //	true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
	StopPrice               string                  `s2m:"stopPrice,omitempty" json:"stopPrice,omitempty"`                             // Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ActivationPrice         string                  `s2m:"activationPrice,omitempty" json:"activationPrice,omitempty"`                 // Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)
	CallbackRate            string                  `s2m:"callbackRate,omitempty" json:"callbackRate,omitempty"`                       // Used with TRAILING_STOP_MARKET orders, min 0.1, max 5 where 1 for 1%
	WorkingType             FuturesWorkingType      `s2m:"workingType,omitempty" json:"workingType,omitempty"`                         // stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE".Default "CONTRACT_PRICE"
	PriceProtect            BigBool                 `s2m:"priceProtect,omitempty" json:"priceProtect,omitempty"`                       // "TRUE" or "FALSE", default "FALSE".Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	NewOrderRespType        OrderResponseType       `s2m:"newOrderRespType,omitempty" json:"newOrderRespType,omitempty"`               // "ACK", "RESULT", default "ACK"
	PriceMatch              string                  `s2m:"priceMatch,omitempty" json:"priceMatch,omitempty"`                           //  only available for LIMIT/STOP/TAKE_PROFIT order, can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20. Can't be passed together with price
	SelfTradePreventionMode SelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty" json:"selfTradePreventionMode,omitempty"` // NONE:No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire maker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers , default NONE
	GoodTillDate            string                  `s2m:"goodTillDate,omitempty" json:"goodTillDate,omitempty"`
}

type FuturesPlaceMultiOrdersParams struct {
	BatchOrders []FuturesNewMultiOrdersOrderParams `s2m:"batchOrders"` // max 5 orders
}

// FuturesPlaceMultiOrdersConfig
// Response []FuturesOrder may contain failed orders with error code and msg.
// TODO should add ErrFuMultiOrdersAllFailed or ErrFuMultiOrdersSomeFailed?
var FuturesPlaceMultiOrdersConfig = cex.ReqConfig[FuturesPlaceMultiOrdersParams, []FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/batchOrders",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesOrder]),
}

type FuturesModifyMultiOrdersOrderParams struct {
	OrderId           string    `s2m:"orderId,omitempty"`
	OrigClientOrderId string    `s2m:"origClientOrderId,omitempty"`
	Symbol            string    `s2m:"symbol,omitempty"`
	Side              OrderSide `s2m:"side,omitempty"` // needs to be same as origin order
	Quantity          string    `s2m:"quantity,omitempty"`
	Price             string    `s2m:"price,omitempty"`
	PriceMatch        string    `s2m:"priceMatch,omitempty"`
}

type FuturesModifyMultiOrdersParams struct {
	BatchOrders []FuturesModifyMultiOrdersOrderParams `s2m:"batchOrders"`
}

var FuturesModifyMultiOrdersConfig = cex.ReqConfig[FuturesModifyMultiOrdersParams, []FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/batchOrders",
		Method:           http.MethodPut,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesOrder]),
}

type FuturesOrderModifyHistoriesParams struct {
	Symbol            string `s2m:"symbol,omitempty"`
	OrderId           int64  `s2m:"orderId,omitempty"`
	OrigClientOrderId string `s2m:"origClientOrderId,omitempty"`
	StartTime         int64  `s2m:"startTime,omitempty"`
	EndTime           int64  `s2m:"endTime,omitempty"`
	Limit             int    `s2m:"limit,omitempty"` // Default 1000; max 1000
}

type FuturesOrderModifyHistory struct {
	AmendmentId   int    `json:"amendmentId"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	Time          int64  `json:"time"` // Order modification time
	Amendment     struct {
		Price struct {
			Before float64 `json:"before,string"`
			After  float64 `json:"after,string"`
		} `json:"price"`
		OrigQty struct {
			Before float64 `json:"before,string"`
			After  float64 `json:"after,string"`
		} `json:"origQty"`
		Count int `json:"count"` // Order modification count, representing the number of times the order has been modified
	} `json:"amendment"`
	PriceMatch string `json:"priceMatch"`
}

var FuturesOrderModifyHistoriesConfig = cex.ReqConfig[FuturesOrderModifyHistoriesParams, []FuturesOrderModifyHistory]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/orderAmendment",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesOrderModifyHistory]),
}

type FuturesQueryOrCancelOrderParams struct {
	Symbol string `s2m:"symbol,omitempty"`

	// If canceling all orders, ignore.
	OrderId           int64  `s2m:"orderId,omitempty"`
	OrigClientOrderId string `s2m:"origClientOrderId,omitempty"`
}

var FuturesQueryOrderConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var FuturesCancelOrderConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var FuturesCancelAllOpenOrdersConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, CodeMsg]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/allOpenOrders",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[CodeMsg]),
}

type FuturesCancelMultiOrdersParams struct {
	Symbol string `s2m:"symbol,omitempty"`
	// Do not set orderIdList and origClientOrderIdList together
	OrderIdList           []int64  `s2m:"orderIdList,omitempty"`           // max length: 10
	OrigClientOrderIdList []string `s2m:"origClientOrderIdList,omitempty"` // max length: 10
}

var FuturesCancelMultiOrdersConfig = cex.ReqConfig[FuturesCancelMultiOrdersParams, []FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/batchOrders",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesOrder]),
}

type FuturesAutoCancelAllOpenOrdersParams struct {
	Symbol string `s2m:"symbol,omitempty"`
	// millisecond
	// system will check all countdowns approximately every 10 milliseconds
	// 0 to cancel timer, do not omit empty
	CountdownTime int64 `s2m:"countdownTime"`
}

type FuturesAutoCancelAllOpenOrdersResponse struct {
	Symbol        string `json:"symbol"`
	CountdownTime int64  `json:"countdownTime,string"`
}

var FuturesAutoCancelAllOpenOrdersConfig = cex.ReqConfig[FuturesAutoCancelAllOpenOrdersParams, FuturesAutoCancelAllOpenOrdersResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/countdownCancelAll",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesAutoCancelAllOpenOrdersResponse]),
}

var FuturesCurrentOpenOrderConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/openOrder",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var FuturesCurrentAllOpenOrdersConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, []FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/openOrders",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesOrder]),
}

// FuturesAllOrdersParams
// If orderId is set, it will get orders >= that orderId. Otherwise, most recent orders are returned.
// The query time period must be less than 7 days( default as the recent 7 days).
type FuturesAllOrdersParams struct {
	Symbol    string `s2m:"symbol,omitempty"`
	OrderId   int64  `s2m:"orderId,omitempty"`
	StartTime int64  `s2m:"startTime,omitempty"`
	EndTime   int64  `s2m:"endTime,omitempty"`
	Limit     int    `s2m:"limit,omitempty"` // default: 500, max: 1000
}

// FuturesAllOrdersConfig
// These orders will not be found:
// order status is CANCELED or EXPIRED AND order has NO filled trade AND created time + 3 days < current time
// order create time + 90 days < current time
var FuturesAllOrdersConfig = cex.ReqConfig[FuturesAllOrdersParams, []FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/allOrders",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesOrder]),
}

type FuturesAccountBalance struct {
	AccountAlias       string  `json:"accountAlias"`
	Asset              string  `json:"asset"`
	Balance            float64 `json:"balance,string"`
	CrossWalletBalance float64 `json:"crossWalletBalance,string"`
	CrossUnPnl         float64 `json:"crossUnPnl,string"`
	AvailableBalance   float64 `json:"availableBalance,string"`
	MaxWithdrawAmount  float64 `json:"maxWithdrawAmount,string"`
	MarginAvailable    bool    `json:"marginAvailable"`
	UpdateTime         int64   `json:"updateTime"`
}

var FuturesAccountBalancesConfig = cex.ReqConfig[cex.NilReqData, []FuturesAccountBalance]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV2 + "/balance",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesAccountBalance]),
}

type FuturesAccountAsset struct {
	Asset                  string  `json:"asset"`
	WalletBalance          float64 `json:"walletBalance,string"`
	UnrealizedProfit       float64 `json:"unrealizedProfit,string"`
	MarginBalance          float64 `json:"marginBalance,string"`
	MaintMargin            float64 `json:"maintMargin,string"`
	InitialMargin          float64 `json:"initialMargin,string"`
	PositionInitialMargin  float64 `json:"positionInitialMargin,string"`
	OpenOrderInitialMargin float64 `json:"openOrderInitialMargin,string"`
	CrossWalletBalance     float64 `json:"crossWalletBalance,string"`
	CrossUnPnl             float64 `json:"crossUnPnl,string"`
	AvailableBalance       float64 `json:"availableBalance,string"`
	MaxWithdrawAmount      float64 `json:"maxWithdrawAmount,string"`
	MarginAvailable        bool    `json:"marginAvailable"`
	UpdateTime             int64   `json:"updateTime"`
}

type FuturesAccountPosition struct {
	Symbol                 string              `json:"symbol"`
	InitialMargin          float64             `json:"initialMargin,string"`
	MaintMargin            float64             `json:"maintMargin,string"`
	UnrealizedProfit       float64             `json:"unrealizedProfit,string"`
	PositionInitialMargin  float64             `json:"positionInitialMargin,string"`
	OpenOrderInitialMargin float64             `json:"openOrderInitialMargin,string"`
	Leverage               float64             `json:"leverage,string"`
	Isolated               bool                `json:"isolated"`
	EntryPrice             float64             `json:"entryPrice,string"`
	MaxNotional            float64             `json:"maxNotional,string"`
	BidNotional            float64             `json:"bidNotional,string"`
	AskNotional            float64             `json:"askNotional,string"`
	PositionSide           FuturesPositionSide `json:"positionSide"`
	SignPositionAmt        float64             `json:"positionAmt,string"` // long: > 0, short: < 0
	UpdateTime             int64               `json:"updateTime"`

	// multi asset mode
	BreakEvenPrice string `json:"breakEvenPrice"`
}

func (p FuturesAccountPosition) AbsPositionAmt() float64 {
	return math.Abs(p.SignPositionAmt)
}

type FuturesAccount struct {
	FeeTier                     float64                  `json:"feeTier"`
	CanTrade                    bool                     `json:"canTrade"`
	CanDeposit                  bool                     `json:"canDeposit"`
	CanWithdraw                 bool                     `json:"canWithdraw"`
	UpdateTime                  int64                    `json:"updateTime"`
	MultiAssetsMargin           bool                     `json:"multiAssetsMargin"`
	TradeGroupId                int64                    `json:"tradeGroupId"`
	TotalInitialMargin          float64                  `json:"totalInitialMargin,string"`
	TotalMaintMargin            float64                  `json:"totalMaintMargin,string"`
	TotalWalletBalance          float64                  `json:"totalWalletBalance,string"`
	TotalUnrealizedProfit       float64                  `json:"totalUnrealizedProfit,string"`
	TotalMarginBalance          float64                  `json:"totalMarginBalance,string"`
	TotalPositionInitialMargin  float64                  `json:"totalPositionInitialMargin,string"`
	TotalOpenOrderInitialMargin float64                  `json:"totalOpenOrderInitialMargin,string"`
	TotalCrossWalletBalance     float64                  `json:"totalCrossWalletBalance,string"`
	TotalCrossUnPnl             float64                  `json:"totalCrossUnPnl,string"`
	AvailableBalance            float64                  `json:"availableBalance,string"`
	MaxWithdrawAmount           float64                  `json:"maxWithdrawAmount,string"`
	Assets                      []FuturesAccountAsset    `json:"assets"`
	Positions                   []FuturesAccountPosition `json:"positions"`
}

var FuturesAccountConfig = cex.ReqConfig[cex.NilReqData, FuturesAccount]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV2 + "/account",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesAccount]),
}

type FuturesChangeInitialLeverageParams struct {
	Symbol   string `s2m:"symbol"`
	Leverage int    `s2m:"leverage"`
}

type FuturesChangeInitialLeverageResponse struct {
	Symbol           string  `json:"symbol"`
	Leverage         int     `json:"leverage"`
	MaxNotionalValue float64 `json:"maxNotionalValue,string"`
}

var FuturesChangeInitialLeverageConfig = cex.ReqConfig[FuturesChangeInitialLeverageParams, FuturesChangeInitialLeverageResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/leverage",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesChangeInitialLeverageResponse]),
}

type FuturesChangeMarginTypeParams struct {
	Symbol     string            `s2m:"symbol"`
	MarginType FuturesMarginType `s2m:"marginType"`
}

var FuturesChangeMarginTypeConfig = cex.ReqConfig[FuturesChangeMarginTypeParams, CodeMsg]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/marginType",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[CodeMsg]),
}

type FuturesModifyIsolatedPositionMarginParams struct {
	Symbol       string                  `s2m:"symbol,omitempty"`
	PositionSide FuturesPositionSide     `s2m:"positionSide,omitempty"`
	Amount       float64                 `s2m:"amount,omitempty"`
	Type         FuturesModifyMarginType `s2m:"type,omitempty"` // 1: add position margin; 2: reduce position margin
}

type FuturesModifyIsolatedPositionMarginResponse struct {
	Amount float64                 `json:"amount"`
	Code   int                     `json:"code"`
	Msg    string                  `json:"msg"`
	Type   FuturesModifyMarginType `json:"type"`
}

var FuturesModifyIsolatedPositionMarginConfig = cex.ReqConfig[FuturesModifyIsolatedPositionMarginParams, FuturesModifyIsolatedPositionMarginResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/positionMargin",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesModifyIsolatedPositionMarginResponse]),
}

type FuturesPositionMarginChangeHistoriesParams struct {
	Symbol    string                  `s2m:"symbol,omitempty"`
	Type      FuturesModifyMarginType `s2m:"type,omitempty"`
	StartTime int64                   `s2m:"startTime,omitempty"`
	EndTime   int64                   `s2m:"endTime,omitempty"`
	Limit     int                     `s2m:"limit,omitempty"` // default: 500
}

type FuturesPositionMarginChangeHistory struct {
	Symbol       string                  `json:"symbol"`
	Type         FuturesModifyMarginType `json:"type"`
	DeltaType    FuturesMarginDeltaType  `json:"deltaType"`
	Amount       float64                 `json:"amount,string"`
	Asset        string                  `json:"asset"`
	Time         int64                   `json:"time"`
	PositionSide FuturesPositionSide     `json:"positionSide"`
}

var FuturesPositionMarginChangeHistoriesConfig = cex.ReqConfig[FuturesPositionMarginChangeHistoriesParams, []FuturesPositionMarginChangeHistory]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/positionMargin/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesPositionMarginChangeHistory]),
}

type FuturesPositionsParams struct {
	Symbol string `s2m:"symbol"`
}

type FuturesPosition struct {
	Symbol           string                     `json:"symbol"`
	PositionSide     string                     `json:"positionSide"`
	EntryPrice       float64                    `json:"entryPrice,string"`
	BreakEvenPrice   float64                    `json:"breakEvenPrice,string"`
	MarginType       FuturesMarginLowerCaseType `json:"marginType"`
	IsAutoAddMargin  SmallBool                  `json:"isAutoAddMargin"`
	IsolatedMargin   float64                    `json:"isolatedMargin,string"`
	Leverage         float64                    `json:"leverage,string"`
	LiquidationPrice float64                    `json:"liquidationPrice,string"`
	MarkPrice        float64                    `json:"markPrice,string"`
	MaxNotionalValue float64                    `json:"maxNotionalValue,string"`
	SignPositionAmt  float64                    `json:"positionAmt,string"` // long: > 0, short: < 0
	Notional         float64                    `json:"notional,string"`
	IsolatedWallet   float64                    `json:"isolatedWallet,string"`
	UnRealizedProfit float64                    `json:"unRealizedProfit,string"`
	UpdateTime       int                        `json:"updateTime"`
}

func (p FuturesPosition) AbsPositionAmt() float64 {
	return math.Abs(p.SignPositionAmt)
}

var FuturesPositionsConfig = cex.ReqConfig[FuturesPositionsParams, []FuturesPosition]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV2 + "/positionRisk",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesPosition]),
}

type FuturesAccountTradeListParams struct {
	Symbol    string `s2m:"symbol,omitempty"`
	OrderId   int64  `s2m:"orderId,omitempty"` // This can only be used in combination with symbol
	StartTime int64  `s2m:"startTime,omitempty"`
	EndTime   int64  `s2m:"endTime,omitempty"`
	FromId    int64  `s2m:"fromId,omitempty"` // Trade id to fetch from.Default gets most recent trades.
	Limit     int    `s2m:"limit,omitempty"`  // Default 500 max 1000.
}

type FuturesTradeHistory struct {
	Id              int64               `json:"id"`
	OrderId         int64               `json:"orderId"`
	Symbol          string              `json:"symbol"`
	Buyer           bool                `json:"buyer"`
	Maker           bool                `json:"maker"`
	PositionSide    FuturesPositionSide `json:"positionSide"`
	Side            OrderSide           `json:"side"`
	Qty             float64             `json:"qty,string"`
	Price           float64             `json:"price,string"`
	QuoteQty        float64             `json:"quoteQty,string"`
	RealizedPnl     float64             `json:"realizedPnl,string"`
	Commission      float64             `json:"commission,string"`
	CommissionAsset string              `json:"commissionAsset"`
	Time            int64               `json:"time"`
}

var FuturesAccountTradeListConfig = cex.ReqConfig[FuturesAccountTradeListParams, []FuturesTradeHistory]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/userTrades",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesTradeHistory]),
}

type FuturesIncomeHistoriesParams struct {
	Symbol     string            `s2m:"symbol,omitempty"`
	IncomeType FuturesIncomeType `s2m:"incomeType,omitempty"`
	StartTime  int64             `s2m:"startTime,omitempty"` // Timestamp in ms to get funding from INCLUSIVE.
	EndTime    int64             `s2m:"endTime,omitempty"`   // Timestamp in ms to get funding until INCLUSIVE.
	Page       int               `s2m:"page,omitempty"`
	Limit      int               `s2m:"limit,omitempty"` // Default 100 max 1000
}

type FuturesIncome struct {
	Symbol     string            `json:"symbol"`
	IncomeType FuturesIncomeType `json:"incomeType"`
	Income     float64           `json:"income,string"`
	Asset      string            `json:"asset"`
	Info       string            `json:"info"`
	Time       int64             `json:"time"`
	TranId     int64             `json:"tranId"`
	TradeId    string            `json:"tradeId"`
}

var FuturesIncomeHistoriesConfig = cex.ReqConfig[FuturesIncomeHistoriesParams, []FuturesIncome]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/income",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesIncome]),
}

type FuturesCommissionRateParams struct {
	Symbol string `s2m:"symbol"`
}

type FuturesCommissionRate struct {
	Symbol              string  `json:"symbol"`
	MakerCommissionRate float64 `json:"makerCommissionRate,string"`
	TakerCommissionRate float64 `json:"takerCommissionRate,string"`
}

var FuturesCommissionRateConfig = cex.ReqConfig[FuturesCommissionRateParams, FuturesCommissionRate]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/commissionRate",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesCommissionRate]),
}
