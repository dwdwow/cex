package bnc

import "testing"

func TestFuChangePositionMode(t *testing.T) {
	testConfig(FuChangePositionModeConfig, ChangePositionModParams{DualSidePosition: SmallFalse})
}

func TestFuCurrentPositionMode(t *testing.T) {
	// TODO retest
	// return {"code":-1022,"msg":"Signature for this request is not valid."}
	testConfig(FuPositionModeConfig, nil)
}

func TestFuChangeMultiAssetsMode(t *testing.T) {
	testConfig(FuChangeMultiAssetsModeConfig, FuChangeMultiAssetsModeParams{MultiAssetsMargin: SmallFalse})
}

func TestFuCurrentMultiAssetsMode(t *testing.T) {
	testConfig(FuCurrentMultiAssetsModeConfig, nil)
}

func TestFuNewOrder(t *testing.T) {
	testConfig(FuNewOrderConfig, FuNewOrderParams{
		Symbol:                  "ETHUSDT",
		Side:                    OrderSideBuy,
		PositionSide:            FuPosBoth,
		Type:                    OrderTypeLimit,
		TimeInForce:             TimeInForceGtc,
		Quantity:                0.02,
		Price:                   1500,
		ReduceOnly:              "",
		NewClientOrderId:        "asdfljksdhkf",
		StopPrice:               0,
		ClosePosition:           false,
		ActivationPrice:         0,
		CallbackRate:            0,
		WorkingType:             "",
		PriceProtect:            "",
		NewOrderRespType:        "",
		PriceMatch:              "",
		SelfTradePreventionMode: "",
		GoodTillDate:            0,
	})
}

func TestFuModifyOrder(t *testing.T) {
	testConfig(FuModifyOrderConfig, FuModifyOrderParams{
		OrderId:           0,
		OrigClientOrderId: "asdfljksdhkf",
		Symbol:            "ETHUSDT",
		Side:              OrderSideBuy,
		Quantity:          0.02,
		Price:             1200,
		PriceMatch:        "",
	})
}

func TestFuPlaceMultiOrders(t *testing.T) {
	testConfig(FuPlaceMultiOrdersConfig, FuPlaceMultiOrdersParams{
		BatchOrders: []FuNewMultiOrdersOrderParams{
			{
				Symbol:                  "ETHUSDT",
				PositionSide:            FuPosBoth,
				Type:                    OrderTypeLimit,
				Side:                    OrderSideBuy,
				Quantity:                "0.02",
				Price:                   "1500",
				TimeInForce:             TimeInForceGtc,
				NewClientOrderId:        "ashjkdg111",
				ReduceOnly:              "",
				ClosePosition:           false,
				StopPrice:               "",
				ActivationPrice:         "",
				CallbackRate:            "",
				WorkingType:             "",
				PriceProtect:            "",
				NewOrderRespType:        "",
				PriceMatch:              "",
				SelfTradePreventionMode: "",
				GoodTillDate:            "",
			},
			{
				Symbol:                  "ETHUSDT",
				PositionSide:            FuPosBoth,
				Type:                    OrderTypeLimit,
				Side:                    OrderSideBuy,
				Quantity:                "0.02",
				Price:                   "1700",
				TimeInForce:             TimeInForceGtc,
				NewClientOrderId:        "ashjkdg111",
				ReduceOnly:              "",
				ClosePosition:           false,
				StopPrice:               "",
				ActivationPrice:         "",
				CallbackRate:            "",
				WorkingType:             "",
				PriceProtect:            "",
				NewOrderRespType:        "",
				PriceMatch:              "",
				SelfTradePreventionMode: "",
				GoodTillDate:            "",
			},
			{
				Symbol:                  "ETHUSDT",
				PositionSide:            FuPosBoth,
				Type:                    OrderTypeLimit,
				Side:                    OrderSideBuy,
				Quantity:                "0.02",
				Price:                   "1900",
				TimeInForce:             TimeInForceGtc,
				NewClientOrderId:        "ashjkdg11",
				ReduceOnly:              "",
				ClosePosition:           false,
				StopPrice:               "",
				ActivationPrice:         "",
				CallbackRate:            "",
				WorkingType:             "",
				PriceProtect:            "",
				NewOrderRespType:        "",
				PriceMatch:              "",
				SelfTradePreventionMode: "",
				GoodTillDate:            "",
			},
		},
	})
}

func TestFuOrderModifyHistories(t *testing.T) {
	testConfig(FuOrderModifyHistoriesConfig, FuOrderModifyHistoriesParams{
		Symbol:            "ETHUSDT",
		OrderId:           0,
		OrigClientOrderId: "",
		StartTime:         0,
		EndTime:           0,
		Limit:             0,
	})
}
