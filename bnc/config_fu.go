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
	DualSidePosition bool `json:"dualSidePosition" bson:"dualSidePosition"`
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
	MultiAssetsMargin bool `json:"multiAssetsMargin" bson:"multiAssetsMargin"`
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
	Symbol                  string                  `json:"symbol" bson:"symbol"`
	OrderId                 int64                   `json:"orderId" bson:"orderId"`
	ClientOrderId           string                  `json:"clientOrderId" bson:"clientOrderId"`
	Type                    OrderType               `json:"type" bson:"type"`
	PositionSide            FuturesPositionSide     `json:"positionSide" bson:"positionSide"`
	Side                    OrderSide               `json:"side" bson:"side"`
	OrigQty                 float64                 `json:"origQty,string" bson:"origQty,string"`
	Price                   float64                 `json:"price,string" bson:"price,string"` // orig price
	ExecutedQty             float64                 `json:"executedQty,string" bson:"executedQty,string"`
	AvgPrice                float64                 `json:"avgPrice,string" bson:"avgPrice,string"`
	ReduceOnly              bool                    `json:"reduceOnly" bson:"reduceOnly"`
	Status                  OrderStatus             `json:"status" bson:"status"`
	StopPrice               float64                 `json:"stopPrice,string" bson:"stopPrice,string"`
	ClosePosition           bool                    `json:"closePosition" bson:"closePosition"`
	TimeInForce             TimeInForce             `json:"timeInForce" bson:"timeInForce"`
	OrigType                OrderType               `json:"origType" bson:"origType"`
	UpdateTime              int64                   `json:"updateTime" bson:"updateTime"`
	WorkingType             FuturesWorkingType      `json:"workingType" bson:"workingType"`
	PriceProtect            bool                    `json:"priceProtect" bson:"priceProtect"`
	PriceMatch              string                  `json:"priceMatch" bson:"priceMatch"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode" bson:"selfTradePreventionMode"`
	GoodTillDate            int64                   `json:"goodTillDate" bson:"goodTillDate"`

	// new order
	CumQuote      float64 `json:"cumQuote,string" bson:"cumQuote,string"`
	ActivatePrice float64 `json:"activatePrice,string" bson:"activatePrice,string"`
	PriceRate     float64 `json:"priceRate,string" bson:"priceRate,string"`

	// new, modify order
	CumQty float64 `json:"cumQty,string" bson:"cumQty,string"`

	// modify order
	Pair    string `json:"pair" bson:"pair"`       // same as symbol
	CumBase string `json:"cumBase" bson:"cumBase"` // same as CumQuote? should verify

	// query order
	Time int64 `json:"time" bson:"time"`

	// place, modify multi orders
	Code int    `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
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
	AmendmentId   int    `json:"amendmentId" bson:"amendmentId"`
	Symbol        string `json:"symbol" bson:"symbol"`
	Pair          string `json:"pair" bson:"pair"`
	OrderId       int64  `json:"orderId" bson:"orderId"`
	ClientOrderId string `json:"clientOrderId" bson:"clientOrderId"`
	Time          int64  `json:"time" bson:"time"` // Order modification time
	Amendment     struct {
		Price struct {
			Before float64 `json:"before,string" bson:"before,string"`
			After  float64 `json:"after,string" bson:"after,string"`
		} `json:"price" bson:"price"`
		OrigQty struct {
			Before float64 `json:"before,string" bson:"before,string"`
			After  float64 `json:"after,string" bson:"after,string"`
		} `json:"origQty" bson:"origQty"`
		Count int `json:"count" bson:"count"` // Order modification count, representing the number of times the order has been modified
	} `json:"amendment" bson:"amendment"`
	PriceMatch string `json:"priceMatch" bson:"priceMatch"`
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
	Symbol        string `json:"symbol" bson:"symbol"`
	CountdownTime int64  `json:"countdownTime,string" bson:"countdownTime,string"`
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
	AccountAlias       string  `json:"accountAlias" bson:"accountAlias"`
	Asset              string  `json:"asset" bson:"asset"`
	Balance            float64 `json:"balance,string" bson:"balance,string"`
	CrossWalletBalance float64 `json:"crossWalletBalance,string" bson:"crossWalletBalance,string"`
	CrossUnPnl         float64 `json:"crossUnPnl,string" bson:"crossUnPnl,string"`
	AvailableBalance   float64 `json:"availableBalance,string" bson:"availableBalance,string"`
	MaxWithdrawAmount  float64 `json:"maxWithdrawAmount,string" bson:"maxWithdrawAmount,string"`
	MarginAvailable    bool    `json:"marginAvailable" bson:"marginAvailable"`
	UpdateTime         int64   `json:"updateTime" bson:"updateTime"`
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
	Asset                  string  `json:"asset" bson:"asset"`
	WalletBalance          float64 `json:"walletBalance,string" bson:"walletBalance,string"`
	UnrealizedProfit       float64 `json:"unrealizedProfit,string" bson:"unrealizedProfit,string"`
	MarginBalance          float64 `json:"marginBalance,string" bson:"marginBalance,string"`
	MaintMargin            float64 `json:"maintMargin,string" bson:"maintMargin,string"`
	InitialMargin          float64 `json:"initialMargin,string" bson:"initialMargin,string"`
	PositionInitialMargin  float64 `json:"positionInitialMargin,string" bson:"positionInitialMargin,string"`
	OpenOrderInitialMargin float64 `json:"openOrderInitialMargin,string" bson:"openOrderInitialMargin,string"`
	CrossWalletBalance     float64 `json:"crossWalletBalance,string" bson:"crossWalletBalance,string"`
	CrossUnPnl             float64 `json:"crossUnPnl,string" bson:"crossUnPnl,string"`
	AvailableBalance       float64 `json:"availableBalance,string" bson:"availableBalance,string"`
	MaxWithdrawAmount      float64 `json:"maxWithdrawAmount,string" bson:"maxWithdrawAmount,string"`
	MarginAvailable        bool    `json:"marginAvailable" bson:"marginAvailable"`
	UpdateTime             int64   `json:"updateTime" bson:"updateTime"`
}

type FuturesAccountPosition struct {
	Symbol                 string              `json:"symbol" bson:"symbol"`
	InitialMargin          float64             `json:"initialMargin,string" bson:"initialMargin,string"`
	MaintMargin            float64             `json:"maintMargin,string" bson:"maintMargin,string"`
	UnrealizedProfit       float64             `json:"unrealizedProfit,string" bson:"unrealizedProfit,string"`
	PositionInitialMargin  float64             `json:"positionInitialMargin,string" bson:"positionInitialMargin,string"`
	OpenOrderInitialMargin float64             `json:"openOrderInitialMargin,string" bson:"openOrderInitialMargin,string"`
	Leverage               float64             `json:"leverage,string" bson:"leverage,string"`
	Isolated               bool                `json:"isolated" bson:"isolated"`
	EntryPrice             float64             `json:"entryPrice,string" bson:"entryPrice,string"`
	MaxNotional            float64             `json:"maxNotional,string" bson:"maxNotional,string"`
	BidNotional            float64             `json:"bidNotional,string" bson:"bidNotional,string"`
	AskNotional            float64             `json:"askNotional,string" bson:"askNotional,string"`
	PositionSide           FuturesPositionSide `json:"positionSide" bson:"positionSide"`
	SignPositionAmt        float64             `json:"positionAmt,string" bson:"positionAmt,string"` // long: > 0, short: < 0
	UpdateTime             int64               `json:"updateTime" bson:"updateTime"`

	// multi asset mode
	BreakEvenPrice string `json:"breakEvenPrice" bson:"breakEvenPrice"`
}

func (p FuturesAccountPosition) AbsPositionAmt() float64 {
	return math.Abs(p.SignPositionAmt)
}

type FuturesAccount struct {
	FeeTier                     float64                  `json:"feeTier" bson:"feeTier"`
	CanTrade                    bool                     `json:"canTrade" bson:"canTrade"`
	CanDeposit                  bool                     `json:"canDeposit" bson:"canDeposit"`
	CanWithdraw                 bool                     `json:"canWithdraw" bson:"canWithdraw"`
	UpdateTime                  int64                    `json:"updateTime" bson:"updateTime"`
	MultiAssetsMargin           bool                     `json:"multiAssetsMargin" bson:"multiAssetsMargin"`
	TradeGroupId                int64                    `json:"tradeGroupId" bson:"tradeGroupId"`
	TotalInitialMargin          float64                  `json:"totalInitialMargin,string" bson:"totalInitialMargin,string"`
	TotalMaintMargin            float64                  `json:"totalMaintMargin,string" bson:"totalMaintMargin,string"`
	TotalWalletBalance          float64                  `json:"totalWalletBalance,string" bson:"totalWalletBalance,string"`
	TotalUnrealizedProfit       float64                  `json:"totalUnrealizedProfit,string" bson:"totalUnrealizedProfit,string"`
	TotalMarginBalance          float64                  `json:"totalMarginBalance,string" bson:"totalMarginBalance,string"`
	TotalPositionInitialMargin  float64                  `json:"totalPositionInitialMargin,string" bson:"totalPositionInitialMargin,string"`
	TotalOpenOrderInitialMargin float64                  `json:"totalOpenOrderInitialMargin,string" bson:"totalOpenOrderInitialMargin,string"`
	TotalCrossWalletBalance     float64                  `json:"totalCrossWalletBalance,string" bson:"totalCrossWalletBalance,string"`
	TotalCrossUnPnl             float64                  `json:"totalCrossUnPnl,string" bson:"totalCrossUnPnl,string"`
	AvailableBalance            float64                  `json:"availableBalance,string" bson:"availableBalance,string"`
	MaxWithdrawAmount           float64                  `json:"maxWithdrawAmount,string" bson:"maxWithdrawAmount,string"`
	Assets                      []FuturesAccountAsset    `json:"assets" bson:"assets"`
	Positions                   []FuturesAccountPosition `json:"positions" bson:"positions"`
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
	Symbol           string  `json:"symbol" bson:"symbol"`
	Leverage         int     `json:"leverage" bson:"leverage"`
	MaxNotionalValue float64 `json:"maxNotionalValue,string" bson:"maxNotionalValue,string"`
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
	Amount float64                 `json:"amount" bson:"amount"`
	Code   int                     `json:"code" bson:"code"`
	Msg    string                  `json:"msg" bson:"msg"`
	Type   FuturesModifyMarginType `json:"type" bson:"type"`
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
	Symbol       string                  `json:"symbol" bson:"symbol"`
	Type         FuturesModifyMarginType `json:"type" bson:"type"`
	DeltaType    FuturesMarginDeltaType  `json:"deltaType" bson:"deltaType"`
	Amount       float64                 `json:"amount,string" bson:"amount,string"`
	Asset        string                  `json:"asset" bson:"asset"`
	Time         int64                   `json:"time" bson:"time"`
	PositionSide FuturesPositionSide     `json:"positionSide" bson:"positionSide"`
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
	Symbol string `s2m:"symbol,omitempty"`
}

type FuturesPosition struct {
	Symbol           string  `json:"symbol" bson:"symbol"`
	PositionSide     string  `json:"positionSide" bson:"positionSide"`
	EntryPrice       float64 `json:"entryPrice,string" bson:"entryPrice,string"`
	Leverage         float64 `json:"leverage,string" bson:"leverage,string"`
	LiquidationPrice float64 `json:"liquidationPrice,string" bson:"liquidationPrice,string"`
	MarkPrice        float64 `json:"markPrice,string" bson:"markPrice,string"`
	MaxNotionalValue float64 `json:"maxNotionalValue,string" bson:"maxNotionalValue,string"`
	SignPositionAmt  float64 `json:"positionAmt,string" bson:"positionAmt,string"` // long: > 0, short: < 0
	Notional         float64 `json:"notional,string" bson:"notional,string"`
	UnRealizedProfit float64 `json:"unRealizedProfit,string" bson:"unRealizedProfit,string"`
	UpdateTime       int     `json:"updateTime" bson:"updateTime"`

	// only for futures account, not for portfolio margin account
	BreakEvenPrice  float64                    `json:"breakEvenPrice,string" bson:"breakEvenPrice,string"`
	MarginType      FuturesMarginLowerCaseType `json:"marginType" bson:"marginType"`
	IsAutoAddMargin SmallBool                  `json:"isAutoAddMargin" bson:"isAutoAddMargin"`
	IsolatedMargin  float64                    `json:"isolatedMargin,string" bson:"isolatedMargin,string"`
	IsolatedWallet  float64                    `json:"isolatedWallet,string" bson:"isolatedWallet,string"`
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
	Id              int64               `json:"id" bson:"id"`
	OrderId         int64               `json:"orderId" bson:"orderId"`
	Symbol          string              `json:"symbol" bson:"symbol"`
	Buyer           bool                `json:"buyer" bson:"buyer"`
	Maker           bool                `json:"maker" bson:"maker"`
	PositionSide    FuturesPositionSide `json:"positionSide" bson:"positionSide"`
	Side            OrderSide           `json:"side" bson:"side"`
	Qty             float64             `json:"qty,string" bson:"qty,string"`
	Price           float64             `json:"price,string" bson:"price,string"`
	QuoteQty        float64             `json:"quoteQty,string" bson:"quoteQty,string"`
	RealizedPnl     float64             `json:"realizedPnl,string" bson:"realizedPnl,string"`
	Commission      float64             `json:"commission,string" bson:"commission,string"`
	CommissionAsset string              `json:"commissionAsset" bson:"commissionAsset"`
	Time            int64               `json:"time" bson:"time"`
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
	Symbol     string            `json:"symbol" bson:"symbol"`
	IncomeType FuturesIncomeType `json:"incomeType" bson:"incomeType"`
	Income     float64           `json:"income,string" bson:"income,string"`
	Asset      string            `json:"asset" bson:"asset"`
	Info       string            `json:"info" bson:"info"`
	Time       int64             `json:"time" bson:"time"`
	TranId     int64             `json:"tranId" bson:"tranId"`
	TradeId    string            `json:"tradeId" bson:"tradeId"`
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
	Symbol              string  `json:"symbol" bson:"symbol"`
	MakerCommissionRate float64 `json:"makerCommissionRate,string" bson:"makerCommissionRate,string"`
	TakerCommissionRate float64 `json:"takerCommissionRate,string" bson:"takerCommissionRate,string"`
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
