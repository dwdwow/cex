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
	Symbol                  string                  `s2m:"symbol,omitempty"`
	Side                    OrderSide               `s2m:"side,omitempty"`
	PositionSide            FuPositionSide          `s2m:"positionSide,omitempty"`
	Type                    OrderType               `s2m:"type,omitempty"`
	TimeInForce             TimeInForce             `s2m:"timeInForce,omitempty"`
	Quantity                float64                 `s2m:"quantity,omitempty"`
	ReduceOnly              SmallBool               `s2m:"reduceOnly,omitempty"` // "true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
	Price                   float64                 `s2m:"price,omitempty"`
	NewClientOrderId        string                  `s2m:"newClientOrderId,omitempty"`
	StopPrice               float64                 `s2m:"stopPrice,omitempty"`               // Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ClosePosition           bool                    `s2m:"closePosition,omitempty"`           //	true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
	ActivationPrice         float64                 `s2m:"activationPrice,omitempty"`         // Used with TRAILING_STOP_MARKET orders, default as the latest price(supporting different workingType)
	CallbackRate            float64                 `s2m:"callbackRate,omitempty"`            // Used with TRAILING_STOP_MARKET orders, min 0.1, max 5 where 1 for 1%
	WorkingType             FuWorkingType           `s2m:"workingType,omitempty"`             // stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE".Default "CONTRACT_PRICE"
	PriceProtect            BigBool                 `s2m:"priceProtect,omitempty"`            // "TRUE" or "FALSE", default "FALSE".Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	NewOrderRespType        SpotOrderResponseType   `s2m:"newOrderRespType,omitempty"`        // "ACK", "RESULT", default "ACK"
	PriceMatch              string                  `s2m:"priceMatch,omitempty"`              //  only available for LIMIT/STOP/TAKE_PROFIT order, can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20. Can't be passed together with price
	SelfTradePreventionMode SelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty"` // NONE:No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire maker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers , default NONE
	GoodTillDate            int64                   `s2m:"goodTillDate,omitempty"`
}

type FuOrderResponse struct {
	ClientOrderId           string         `json:"clientOrderId"`
	CumQty                  float64        `json:"cumQty,string"`
	CumQuote                float64        `json:"cumQuote,string"`
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
	ActivatePrice           float64        `json:"activatePrice,string"`
	PriceRate               float64        `json:"priceRate,string"`
	UpdateTime              int64          `json:"updateTime"`
	WorkingType             FuWorkingType  `json:"workingType"`
	PriceProtect            bool           `json:"priceProtect"`
	PriceMatch              string         `json:"priceMatch"`
	SelfTradePreventionMode string         `json:"selfTradePreventionMode"`
	GoodTillDate            int64          `json:"goodTillDate"`
}

var FuNewOrderConfig = cex.ReqConfig[FuNewOrderParams, FuOrderResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/order",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuOrderResponse]),
}
