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
				NewClientOrderId:        "ashjkdg1112",
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

func TestFuModifyMultiOrders(t *testing.T) {
	testConfig(FuModifyMultiOrdersConfig, FuModifyMultiOrdersParams{
		BatchOrders: []FuModifyMultiOrdersOrderParams{
			{
				OrderId:           "8389765651923704480",
				OrigClientOrderId: "",
				Symbol:            "ETHUSDT",
				Side:              OrderSideBuy,
				Quantity:          "0.03",
				Price:             "1300",
				PriceMatch:        "",
			},
			{
				OrderId:           "8389765651923827507",
				OrigClientOrderId: "",
				Symbol:            "ETHUSDT",
				Side:              OrderSideBuy,
				Quantity:          "0.01",
				Price:             "1700",
				PriceMatch:        "",
			},
			{
				OrderId:           "8389765651923996553",
				OrigClientOrderId: "",
				Symbol:            "ETHUSDT",
				Side:              OrderSideSell,
				Quantity:          "0.01",
				Price:             "3000",
				PriceMatch:        "",
			},
			{
				OrderId:           "8389765651924056174",
				OrigClientOrderId: "",
				Symbol:            "ETHUSDT",
				Side:              OrderSideSell,
				Quantity:          "0.03",
				Price:             "2800",
				PriceMatch:        "",
			},
		},
	})
}

func TestFuQueryOrder(t *testing.T) {
	testConfig(FuQueryOrderConfig, FuQueryOrCancelOrderParams{
		Symbol:            "ETHUSDT",
		OrderId:           8389765651924056174,
		OrigClientOrderId: "",
	})
}

func TestFuCancelOrder(t *testing.T) {
	testConfig(FuCancelOrderConfig, FuQueryOrCancelOrderParams{
		Symbol:            "ETHUSDT",
		OrderId:           8389765651928248535,
		OrigClientOrderId: "asdfljksdhkf",
	})
}

func TestFuCancelAllOpenOrders(t *testing.T) {
	testConfig(FuCancelAllOpenOrdersConfig, FuQueryOrCancelOrderParams{
		Symbol: "ETHUSDT",
	})
}

func TestFuCancelMultiOrders(t *testing.T) {
	testConfig(FuCancelMultiOrdersConfig, FuCancelMultiOrdersParams{
		Symbol:                "ETHUSDT",
		OrderIdList:           []int64{8389765651930173670, 8389765651930173671, 8389765651930173669},
		OrigClientOrderIdList: []string{"ashjkdg111", "ashjkdg1112", "ashjkdg11"},
	})
}

func TestFuAutoCancelAllOpenOrders(t *testing.T) {
	testConfig(FuAutoCancelAllOpenOrdersConfig, FuAutoCancelAllOpenOrdersParams{
		Symbol:        "ETHUSDT",
		CountdownTime: 10 * 1000,
	})
}

func TestFuCurrentOpenOrder(t *testing.T) {
	testConfig(FuCurrentOpenOrderConfig, FuQueryOrCancelOrderParams{
		Symbol:            "ETHUSDT",
		OrderId:           0,
		OrigClientOrderId: "",
	})
}

func TestFuCurrentAllOpenOrders(t *testing.T) {
	testConfig(FuCurrentAllOpenOrdersConfig, FuQueryOrCancelOrderParams{
		Symbol: "ETHUSDT",
	})
}

func TestFuAllOrders(t *testing.T) {
	testConfig(FuAllOrdersConfig, FuAllOrdersParams{
		Symbol: "ETHUSDT",
	})
}

func TestFuAccountBalances(t *testing.T) {
	testConfig(FuAccountBalancesConfig, nil)
}

func TestFuAccount(t *testing.T) {
	testConfig(FuAccountConfig, nil)
}

func TestChangeInitialLeverage(t *testing.T) {
	testConfig(FuChangeInitialLeverageConfig, FuChangeInitialLeverageParams{
		Symbol:   "ETHUSDT",
		Leverage: 10,
	})
}

func TestChangeMarginType(t *testing.T) {
	testConfig(FuChangeMarginTypeConfig, FuChangeMarginTypeParams{
		Symbol:     "ETHUSDT",
		MarginType: FuMarginIsolated,
	})
}

func TestModifyIsolatedPositionMargin(t *testing.T) {
	testConfig(FuModifyIsolatedPositionMarginConfig, FuModifyIsolatedPositionMarginParams{
		Symbol:       "ETHUSDT",
		PositionSide: FuPosBoth,
		Amount:       10,
		Type:         FuAddMargin,
	})
}

func TestFuPositionMarginChangeHistories(t *testing.T) {
	testConfig(FuPositionMarginChangeHistoriesConfig, FuPositionMarginChangeHistoriesParams{
		Symbol:    "ETHUSDT",
		Type:      0,
		StartTime: 0,
		EndTime:   0,
		Limit:     0,
	})
}

func TestFuPositions(t *testing.T) {
	testConfig(FuPositionsConfig, FuPositionsParams{Symbol: ""})
}

func TestFuTradeList(t *testing.T) {
	testConfig(FuAccountTradeListConfig, FuAccountTradeListParams{
		Symbol:    "ETHUSDT",
		OrderId:   0,
		StartTime: 0,
		EndTime:   0,
		FromId:    0,
		Limit:     0,
	})
}

func TestFuIncomeHistories(t *testing.T) {
	testConfig(FuIncomeHistoriesConfig, FuIncomeHistoriesParams{
		Symbol:     "",
		IncomeType: "",
		StartTime:  0,
		EndTime:    0,
		Page:       0,
		Limit:      0,
	})
}

func TestFuCommission(t *testing.T) {
	testConfig(FuCommissionRateConfig, FuCommissionRateParams{Symbol: "ETHUSDT"})
}

func TestFlexibleRedeem(t *testing.T) {
	testConfig(FlexibleRedeemConfig, FlexibleRedeemParams{
		ProductId:   "ETH001",
		RedeemAll:   false,
		Amount:      0,
		DestAccount: "",
	})
}

func TestSimpleEarnFlexiblePositions(t *testing.T) {
	testConfig(SimpleEarnFlexiblePositionsConfig, SimpleEarnFlexiblePositionsParams{
		Asset:     "ETH",
		ProductId: "",
		Current:   0,
		Size:      0,
	})
}
