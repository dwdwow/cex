package bnc

import "testing"

func TestFuChangePositionMode(t *testing.T) {
	testConfig(FuturesChangePositionModeConfig, FuturesChangePositionModParams{DualSidePosition: SmallFalse})
}

func TestFuCurrentPositionMode(t *testing.T) {
	// TODO retest
	// return {"code":-1022,"msg":"Signature for this request is not valid."}
	testConfig(FuturesPositionModeConfig, nil)
}

func TestFuChangeMultiAssetsMode(t *testing.T) {
	testConfig(FuturesChangeMultiAssetsModeConfig, FuturesChangeMultiAssetsModeParams{MultiAssetsMargin: SmallFalse})
}

func TestFuCurrentMultiAssetsMode(t *testing.T) {
	testConfig(FuturesCurrentMultiAssetsModeConfig, nil)
}

func TestFuNewOrder(t *testing.T) {
	testConfig(FuturesNewOrderConfig, FuturesNewOrderParams{
		Symbol:                  "ETHUSDT",
		Side:                    OrderSideBuy,
		PositionSide:            FuturesPositionSideBoth,
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
	testConfig(FuturesModifyOrderConfig, FuturesModifyOrderParams{
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
	testConfig(FuturesPlaceMultiOrdersConfig, FuturesPlaceMultiOrdersParams{
		BatchOrders: []FuturesNewMultiOrdersOrderParams{
			{
				Symbol:                  "ETHUSDT",
				PositionSide:            FuturesPositionSideBoth,
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
				PositionSide:            FuturesPositionSideBoth,
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
				PositionSide:            FuturesPositionSideBoth,
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
	testConfig(FuturesOrderModifyHistoriesConfig, FuturesOrderModifyHistoriesParams{
		Symbol:            "ETHUSDT",
		OrderId:           0,
		OrigClientOrderId: "",
		StartTime:         0,
		EndTime:           0,
		Limit:             0,
	})
}

func TestFuModifyMultiOrders(t *testing.T) {
	testConfig(FuturesModifyMultiOrdersConfig, FuturesModifyMultiOrdersParams{
		BatchOrders: []FuturesModifyMultiOrdersOrderParams{
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
	testConfig(FuturesQueryOrderConfig, FuturesQueryOrCancelOrderParams{
		Symbol:            "ETHUSDT",
		OrderId:           8389765651924056174,
		OrigClientOrderId: "",
	})
}

func TestFuCancelOrder(t *testing.T) {
	testConfig(FuturesCancelOrderConfig, FuturesQueryOrCancelOrderParams{
		Symbol:            "ETHUSDT",
		OrderId:           8389765651928248535,
		OrigClientOrderId: "asdfljksdhkf",
	})
}

func TestFuCancelAllOpenOrders(t *testing.T) {
	testConfig(FuturesCancelAllOpenOrdersConfig, FuturesQueryOrCancelOrderParams{
		Symbol: "ETHUSDT",
	})
}

func TestFuCancelMultiOrders(t *testing.T) {
	testConfig(FuturesCancelMultiOrdersConfig, FuturesCancelMultiOrdersParams{
		Symbol:                "ETHUSDT",
		OrderIdList:           []int64{8389765651930173670, 8389765651930173671, 8389765651930173669},
		OrigClientOrderIdList: []string{"ashjkdg111", "ashjkdg1112", "ashjkdg11"},
	})
}

func TestFuAutoCancelAllOpenOrders(t *testing.T) {
	testConfig(FuturesAutoCancelAllOpenOrdersConfig, FuturesAutoCancelAllOpenOrdersParams{
		Symbol:        "ETHUSDT",
		CountdownTime: 10 * 1000,
	})
}

func TestFuCurrentOpenOrder(t *testing.T) {
	testConfig(FuturesCurrentOpenOrderConfig, FuturesQueryOrCancelOrderParams{
		Symbol:            "ETHUSDT",
		OrderId:           0,
		OrigClientOrderId: "",
	})
}

func TestFuCurrentAllOpenOrders(t *testing.T) {
	testConfig(FuturesCurrentAllOpenOrdersConfig, FuturesQueryOrCancelOrderParams{
		Symbol: "ETHUSDT",
	})
}

func TestFuAllOrders(t *testing.T) {
	testConfig(FuturesAllOrdersConfig, FuturesAllOrdersParams{
		Symbol: "ETHUSDT",
	})
}

func TestFuAccountBalances(t *testing.T) {
	testConfig(FuturesAccountBalancesConfig, nil)
}

func TestFuAccount(t *testing.T) {
	testConfig(FuturesAccountConfig, nil)
}

func TestChangeInitialLeverage(t *testing.T) {
	testConfig(FuturesChangeInitialLeverageConfig, FuturesChangeInitialLeverageParams{
		Symbol:   "ETHUSDT",
		Leverage: 10,
	})
}

func TestChangeMarginType(t *testing.T) {
	testConfig(FuturesChangeMarginTypeConfig, FuturesChangeMarginTypeParams{
		Symbol:     "ETHUSDT",
		MarginType: FuturesMarginTypeIsolated,
	})
}

func TestModifyIsolatedPositionMargin(t *testing.T) {
	testConfig(FuturesModifyIsolatedPositionMarginConfig, FuturesModifyIsolatedPositionMarginParams{
		Symbol:       "ETHUSDT",
		PositionSide: FuturesPositionSideBoth,
		Amount:       10,
		Type:         FuturesAddMargin,
	})
}

func TestFuPositionMarginChangeHistories(t *testing.T) {
	testConfig(FuturesPositionMarginChangeHistoriesConfig, FuturesPositionMarginChangeHistoriesParams{
		Symbol:    "ETHUSDT",
		Type:      0,
		StartTime: 0,
		EndTime:   0,
		Limit:     0,
	})
}

func TestFuPositions(t *testing.T) {
	testConfig(FuturesPositionsConfig, FuturesPositionsParams{Symbol: ""})
}

func TestFuTradeList(t *testing.T) {
	testConfig(FuturesAccountTradeListConfig, FuturesAccountTradeListParams{
		Symbol:    "ETHUSDT",
		OrderId:   0,
		StartTime: 0,
		EndTime:   0,
		FromId:    0,
		Limit:     0,
	})
}

func TestFuIncomeHistories(t *testing.T) {
	testConfig(FuturesIncomeHistoriesConfig, FuturesIncomeHistoriesParams{
		Symbol:     "",
		IncomeType: "",
		StartTime:  0,
		EndTime:    0,
		Page:       0,
		Limit:      0,
	})
}

func TestFuCommission(t *testing.T) {
	testConfig(FuturesCommissionRateConfig, FuturesCommissionRateParams{Symbol: "ETHUSDT"})
}

func TestFlexibleRedeem(t *testing.T) {
	testConfig(SimpleEarnFlexibleRedeemConfig, SimpleEarnFlexibleRedeemParams{
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
