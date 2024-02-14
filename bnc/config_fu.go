package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type ChangePositionModParams struct {
	DualSidePosition SmallBool `s2m:"dualSidePosition"`
}

var FuChangePositionModeConfig = cex.ReqConfig[ChangePositionModParams, CodeMsg]{
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

type CurrentPositionModeResult struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

var FuPositionModeConfig = cex.ReqConfig[cex.NilReqData, CurrentPositionModeResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/positionSide/dual",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[CurrentPositionModeResult]),
}

type FuChangeMultiAssetsModeParams struct {
	MultiAssetsMargin SmallBool `s2m:"multiAssetsMargin"`
}

var FuChangeMultiAssetsModeConfig = cex.ReqConfig[FuChangeMultiAssetsModeParams, CodeMsg]{
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

var FuCurrentMultiAssetsModeConfig = cex.ReqConfig[cex.NilReqData, FuCurrentMultiAssetsModeResponse]{
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

type FuNewOrderParams struct {
	Symbol                  string                  `s2m:"symbol,omitempty" json:"symbol,omitempty"`
	PositionSide            FuPositionSide          `s2m:"positionSide,omitempty" json:"positionSide,omitempty"`
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
	WorkingType             FuWorkingType           `s2m:"workingType,omitempty" json:"workingType,omitempty"`                         // stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE".Default "CONTRACT_PRICE"
	PriceProtect            BigBool                 `s2m:"priceProtect,omitempty" json:"priceProtect,omitempty"`                       // "TRUE" or "FALSE", default "FALSE".Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	NewOrderRespType        SpotOrderResponseType   `s2m:"newOrderRespType,omitempty" json:"newOrderRespType,omitempty"`               // "ACK", "RESULT", default "ACK"
	PriceMatch              string                  `s2m:"priceMatch,omitempty" json:"priceMatch,omitempty"`                           //  only available for LIMIT/STOP/TAKE_PROFIT order, can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20. Can't be passed together with price
	SelfTradePreventionMode SelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty" json:"selfTradePreventionMode,omitempty"` // NONE:No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire maker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers , default NONE
	GoodTillDate            int64                   `s2m:"goodTillDate,omitempty" json:"goodTillDate,omitempty"`
}

type FuOrder struct {
	// common
	ClientOrderId           string         `json:"clientOrderId"`
	ExecutedQty             float64        `json:"executedQty,string"`
	OrderId                 int            `json:"orderId"`
	AvgPrice                float64        `json:"avgPrice,string"`
	OrigQty                 float64        `json:"origQty,string"`
	Price                   float64        `json:"price,string"`
	ReduceOnly              bool           `json:"reduceOnly"`
	Side                    OrderSide      `json:"side"`
	PositionSide            FuPositionSide `json:"positionSide"`
	Status                  OrderStatus    `json:"status"`
	StopPrice               float64        `json:"stopPrice,string"`
	ClosePosition           bool           `json:"closePosition"`
	Symbol                  string         `json:"symbol"`
	TimeInForce             TimeInForce    `json:"timeInForce"`
	Type                    OrderType      `json:"type"`
	OrigType                OrderType      `json:"origType"`
	UpdateTime              int64          `json:"updateTime"`
	WorkingType             FuWorkingType  `json:"workingType"`
	PriceProtect            bool           `json:"priceProtect"`
	PriceMatch              string         `json:"priceMatch"`
	SelfTradePreventionMode string         `json:"selfTradePreventionMode"`
	GoodTillDate            int64          `json:"goodTillDate"`

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

var FuNewOrderConfig = cex.ReqConfig[FuNewOrderParams, FuOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuOrder]),
}

type FuModifyOrderParams struct {
	OrderId           int64     `s2m:"orderId,omitempty"`
	OrigClientOrderId string    `s2m:"origClientOrderId,omitempty"`
	Symbol            string    `s2m:"symbol,omitempty"`
	Side              OrderSide `s2m:"side,omitempty"` // needs to be same as origin order
	Quantity          float64   `s2m:"quantity,omitempty"`
	Price             float64   `s2m:"price,omitempty"`
	PriceMatch        string    `s2m:"priceMatch,omitempty"`
}

var FuModifyOrderConfig = cex.ReqConfig[FuModifyOrderParams, FuOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodPut,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuOrder]),
}

// FuNewMultiOrdersOrderParams is different with FuNewOrderParams.
// All fields are string.
// Binance doc example just show that quantity and price are string.
// Do not know if other float/int fields are string or not.
type FuNewMultiOrdersOrderParams struct {
	Symbol                  string                  `s2m:"symbol,omitempty" json:"symbol,omitempty"`
	PositionSide            FuPositionSide          `s2m:"positionSide,omitempty" json:"positionSide,omitempty"`
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
	WorkingType             FuWorkingType           `s2m:"workingType,omitempty" json:"workingType,omitempty"`                         // stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE".Default "CONTRACT_PRICE"
	PriceProtect            BigBool                 `s2m:"priceProtect,omitempty" json:"priceProtect,omitempty"`                       // "TRUE" or "FALSE", default "FALSE".Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	NewOrderRespType        SpotOrderResponseType   `s2m:"newOrderRespType,omitempty" json:"newOrderRespType,omitempty"`               // "ACK", "RESULT", default "ACK"
	PriceMatch              string                  `s2m:"priceMatch,omitempty" json:"priceMatch,omitempty"`                           //  only available for LIMIT/STOP/TAKE_PROFIT order, can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20. Can't be passed together with price
	SelfTradePreventionMode SelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty" json:"selfTradePreventionMode,omitempty"` // NONE:No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire maker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers , default NONE
	GoodTillDate            string                  `s2m:"goodTillDate,omitempty" json:"goodTillDate,omitempty"`
}

type FuPlaceMultiOrdersParams struct {
	BatchOrders []FuNewMultiOrdersOrderParams `s2m:"batchOrders"` // max 5 orders
}

// FuPlaceMultiOrdersConfig
// Response []FuOrder may contain failed orders with error code and msg.
// TODO should add ErrFuMultiOrdersAllFailed or ErrFuMultiOrdersSomeFailed?
var FuPlaceMultiOrdersConfig = cex.ReqConfig[FuPlaceMultiOrdersParams, []FuOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/batchOrders",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuOrder]),
}

type FuModifyMultiOrdersOrderParams struct {
	OrderId           string    `s2m:"orderId,omitempty"`
	OrigClientOrderId string    `s2m:"origClientOrderId,omitempty"`
	Symbol            string    `s2m:"symbol,omitempty"`
	Side              OrderSide `s2m:"side,omitempty"` // needs to be same as origin order
	Quantity          string    `s2m:"quantity,omitempty"`
	Price             string    `s2m:"price,omitempty"`
	PriceMatch        string    `s2m:"priceMatch,omitempty"`
}

type FuModifyMultiOrdersParams struct {
	BatchOrders []FuModifyMultiOrdersOrderParams `s2m:"batchOrders"`
}

var FuModifyMultiOrdersConfig = cex.ReqConfig[FuModifyMultiOrdersParams, []FuOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/batchOrders",
		Method:           http.MethodPut,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuOrder]),
}

type FuOrderModifyHistoriesParams struct {
	Symbol            string `s2m:"symbol,omitempty"`
	OrderId           int64  `s2m:"orderId,omitempty"`
	OrigClientOrderId string `s2m:"origClientOrderId,omitempty"`
	StartTime         int64  `s2m:"startTime,omitempty"`
	EndTime           int64  `s2m:"endTime,omitempty"`
	Limit             int    `s2m:"limit,omitempty"` // Default 1000; max 1000
}

type FuOrderModifyHistory struct {
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

var FuOrderModifyHistoriesConfig = cex.ReqConfig[FuOrderModifyHistoriesParams, []FuOrderModifyHistory]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/orderAmendment",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuOrderModifyHistory]),
}

type FuQueryOrCancelOrderParams struct {
	Symbol string `s2m:"symbol,omitempty"`

	// If canceling all orders, ignore.
	OrderId           int64  `s2m:"orderId,omitempty"`
	OrigClientOrderId string `s2m:"origClientOrderId,omitempty"`
}

var FuQueryOrderConfig = cex.ReqConfig[FuQueryOrCancelOrderParams, FuOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuOrder]),
}

var FuCancelOrderConfig = cex.ReqConfig[FuQueryOrCancelOrderParams, FuOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuOrder]),
}

var FuCancelAllOpenOrdersConfig = cex.ReqConfig[FuQueryOrCancelOrderParams, CodeMsg]{
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

type FuCancelMultiOrdersParams struct {
	Symbol                string   `s2m:"symbol,omitempty"`
	OrderIdList           []int64  `s2m:"orderIdList,omitempty"`           // max length: 10
	OrigClientOrderIdList []string `s2m:"origClientOrderIdList,omitempty"` // max length: 10
}

var FuCancelMultiOrdersConfig = cex.ReqConfig[FuCancelMultiOrdersParams, []FuOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/batchOrders",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuOrder]),
}
