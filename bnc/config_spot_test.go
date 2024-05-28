package bnc

import (
	"fmt"
	"testing"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/props"
)

func readApiKey() cex.Api {
	apiKeys := cex.MustReadApiKey()
	apiKey, ok := apiKeys["2022TEST"]
	if !ok {
		panic("no binance api key")
	}
	return apiKey
}

func readVIPPortmarApiKey() cex.Api {
	apiKeys := cex.MustReadApiKey()
	apiKey, ok := apiKeys["HUANGYAN"]
	if !ok {
		panic("no binance api key")
	}
	return apiKey
}

func testConfig[ReqDataType, RespDataType any](
	config cex.ReqConfig[ReqDataType, RespDataType],
	reqData ReqDataType,
	opts ...cex.CltOpt,
) {
	apiKey := readApiKey()
	user := NewUser(apiKey.ApiKey, apiKey.SecretKey, UserOptPositionSide(FuturesPositionSideBoth))
	_, respData, err := cex.Request(user, config, reqData, opts...)
	props.PanicIfNotNil(err.Err)
	props.PrintlnIndent(respData)
}

func TestCoinInfo(t *testing.T) {
	testConfig(CoinInfoConfig, nil)
}

func TestSpotAccount(t *testing.T) {
	testConfig(SpotAccountConfig, nil)
}

func TestUniversalTransfer(t *testing.T) {
	testConfig(UniversalTransferConfig, UniversalTransferParams{
		Type:       TransferTypeMainUmfuture,
		Asset:      "USDT",
		Amount:     10,
		FromSymbol: "",
		ToSymbol:   "",
	})
}

func TestFlexibleProduct(t *testing.T) {
	testConfig(SimpleEarnFlexibleProductConfig, SimpleEarnFlexibleProductListParams{
		Asset: "BTC",
	})
}

func TestCryptoLoansIncomeHistories(t *testing.T) {
	testConfig(CryptoLoansIncomeHistoriesConfig, CryptoLoansIncomeHistoriesParams{})
}

func TestFlexibleBorrow(t *testing.T) {
	testConfig(CryptoLoanFlexibleBorrowConfig, CryptoLoanFlexibleBorrowParams{
		LoanCoin:         "USDT",
		LoanAmount:       100,
		CollateralCoin:   "ETH",
		CollateralAmount: 0,
	})
}

func TestFlexibleOngoingOrders(t *testing.T) {
	testConfig(CryptoLoanFlexibleOngoingOrdersConfig, CryptoLoanFlexibleOngoingOrdersParams{
		LoanCoin:       "USDT",
		CollateralCoin: "ETH",
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleBorrowHistories(t *testing.T) {
	testConfig(CryptoLoanFlexibleBorrowHistoriesConfig, CryptoLoanFlexibleBorrowHistoriesParams{
		LoanCoin:       "USDT",
		CollateralCoin: "ETH",
		StartTime:      0,
		EndTime:        0,
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleRepay(t *testing.T) {
	testConfig(CryptoLoanFlexibleRepayConfig, CryptoLoanFlexibleRepayParams{
		LoanCoin:         "USDT",
		CollateralCoin:   "ETH",
		RepayAmount:      100,
		CollateralReturn: BigTrue,
		FullRepayment:    BigFalse,
	})
}

func TestFlexibleRepayHistories(t *testing.T) {
	testConfig(CryptoLoanFlexibleRepaymentHistoriesConfig, CryptoLoanFlexibleRepaymentHistoriesParams{
		LoanCoin:       "",
		CollateralCoin: "",
		StartTime:      0,
		EndTime:        0,
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleAdjustLtv(t *testing.T) {
	testConfig(CryptoLoanFlexibleLoanAdjustLtvConfig, CryptoLoanFlexibleAdjustLtvParams{
		LoanCoin:         "USDT",
		CollateralCoin:   "ETH",
		AdjustmentAmount: 0.05,
		Direction:        LTVReduced,
	})
}

func TestFlexibleAdjustLtvHistories(t *testing.T) {
	testConfig(CryptoLoanFlexibleAdjustLtvHistoriesConfig, CryptoLoanFlexibleAdjustLtvHistoriesParams{
		LoanCoin:       "USDT",
		CollateralCoin: "ETH",
		StartTime:      0,
		EndTime:        0,
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleLoanAssets(t *testing.T) {
	testConfig(CryptoLoanFlexibleLoanAssetsConfig, CryptoLoanFlexibleLoanAssetsParams{
		LoanCoin: "",
	})
}

func TestFlexibleCollateralCoins(t *testing.T) {
	testConfig(CryptoLoanFlexibleCollateralCoinsConfig, CryptoLoanFlexibleCollateralCoinsParams{
		CollateralCoin: "",
	})
}

func TestNewSpotOrder(t *testing.T) {
	testConfig(SpotNewOrderConfig, SpotNewOrderParams{
		Symbol:                  "ETHUSDT",
		Type:                    OrderTypeLimit,
		Side:                    OrderSideBuy,
		Quantity:                0.01,
		Price:                   1500,
		TimeInForce:             TimeInForceGtc,
		NewClientOrderId:        "asdfsfhkhuiwe",
		QuoteOrderQty:           0,
		StrategyId:              0,
		StrategyType:            0,
		StopPrice:               0,
		TrailingDelta:           0,
		IcebergQty:              0,
		NewOrderRespType:        "",
		SelfTradePreventionMode: "",
	})
}

func TestCancelSpotOder(t *testing.T) {
	testConfig(SpotCancelOrderConfig, SpotCancelOrderParams{
		Symbol:             "ETHUSDT",
		OrderId:            0,
		OrigClientOrderId:  "",
		NewClientOrderId:   "",
		CancelRestrictions: "",
	})
}

func TestCancelAllSpotOpenOrders(t *testing.T) {
	testConfig(SpotCancelAllOpenOrdersConfig, SpotCancelAllOpenOrdersParams{
		Symbol: "ETHUSDT",
	})
}

func TestSpotQueryOrder(t *testing.T) {
	testConfig(SpotQueryOrderConfig, SpotQueryOrderParams{
		Symbol:            "ETHUSDT",
		OrderId:           0,
		OrigClientOrderId: "",
	})
}

func TestSpotReplaceOrder(t *testing.T) {
	apiKey := readApiKey()
	user := NewUser(apiKey.ApiKey, apiKey.SecretKey)
	_, respData, err := cex.Request(user, SpotReplaceOrderConfig, SpotReplaceOrderParams{
		Symbol:                  "ETHUSDT",
		Type:                    OrderTypeLimit,
		Side:                    OrderSideSell,
		CancelReplaceMode:       SpotCancelReplaceModeStopOnFailure,
		TimeInForce:             TimeInForceGtc,
		Quantity:                10,
		QuoteOrderQty:           0,
		Price:                   3000,
		CancelNewClientOrderId:  "",
		CancelOrigClientOrderId: "",
		CancelOrderId:           15946838304,
		NewClientOrderId:        "",
		StrategyId:              0,
		StrategyType:            0,
		StopPrice:               0,
		TrailingDelta:           0,
		IcebergQty:              0,
		NewOrderRespType:        "",
		SelfTradePreventionMode: "",
		CancelRestrictions:      "",
	})
	if err.IsNotNil() {
		fmt.Println(err)
	}
	props.PrintlnIndent(respData)
}

func TestSpotCurrentOpenOrders(t *testing.T) {
	testConfig(SpotCurrentOpenOrdersConfig, SpotCurrentOpenOrdersParams{Symbol: ""})
}

func TestSpotAllOrders(t *testing.T) {
	testConfig(SpotAllOrdersConfig, SpotAllOrdersParams{
		Symbol:    "ETHUSDT",
		OrderId:   0,
		StartTime: 0,
		EndTime:   0,
		Limit:     0,
	})
}
