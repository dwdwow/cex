package bnctest

import (
	"testing"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/bnc"
	"github.com/dwdwow/cex/test/cextest"
	"github.com/dwdwow/props"
)

func readApiKey() cex.Api {
	apiKeys := cextest.MustReadApiKey()
	apiKey, ok := apiKeys[cex.BINANCE]
	if !ok {
		panic("no binance api key")
	}
	return apiKey
}

func testConfig[ReqDataType, RespDataType any](
	config cex.ReqConfig[ReqDataType, RespDataType],
	reqData ReqDataType,
	opts ...cex.ReqOpt,
) {
	apiKey := readApiKey()
	user := bnc.NewUser(apiKey.ApiKey, apiKey.SecretKey)
	_, respData, err := cex.Request(user, config, reqData, opts...)
	props.PanicIfNotNil(err)
	props.PrintlnIndent(respData)
}

func TestCoinInfo(t *testing.T) {
	testConfig(bnc.CoinInfoConfig, nil)
}

func TestSpotAccount(t *testing.T) {
	testConfig(bnc.SpotAccountConfig, nil)
}

func TestUniversalTransfer(t *testing.T) {
	testConfig(bnc.UniversalTransferConfig, bnc.UniversalTransferParams{
		Type:       bnc.TranType_MAIN_UMFUTURE,
		Asset:      "USDT",
		Amount:     10,
		FromSymbol: "",
		ToSymbol:   "",
	})
}

func TestFlexibleProduct(t *testing.T) {
	testConfig(bnc.FlexibleProductConfig, bnc.FlexibleProductListParams{
		Asset: "BTC",
	})
}

func TestCryptoLoansIncomeHistories(t *testing.T) {
	testConfig(bnc.CryptoLoansIncomeHistoriesConfig, bnc.CryptoLoansIncomeHistoriesParams{})
}

func TestFlexibleBorrow(t *testing.T) {
	testConfig(bnc.FlexibleBorrowConfig, bnc.FlexibleBorrowParams{
		LoanCoin:         "USDT",
		LoanAmount:       100,
		CollateralCoin:   "ETH",
		CollateralAmount: 0,
	})
}

func TestFlexibleOngoingOrders(t *testing.T) {
	testConfig(bnc.FlexibleOngoingOrdersConfig, bnc.FlexibleOngoingOrdersParams{
		LoanCoin:       "USDT",
		CollateralCoin: "ETH",
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleBorrowHistories(t *testing.T) {
	testConfig(bnc.FlexibleBorrowHistoriesConfig, bnc.FlexibleBorrowHistoriesParams{
		LoanCoin:       "USDT",
		CollateralCoin: "ETH",
		StartTime:      0,
		EndTime:        0,
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleRepay(t *testing.T) {
	testConfig(bnc.FlexibleRepayConfig, bnc.FlexibleRepayParams{
		LoanCoin:         "USDT",
		CollateralCoin:   "ETH",
		RepayAmount:      100,
		CollateralReturn: bnc.TRUE,
		FullRepayment:    bnc.FALSE,
	})
}

func TestFlexibleRepayHistories(t *testing.T) {
	testConfig(bnc.FlexibleRepaymentHistoriesConfig, bnc.FlexibleRepaymentHistoriesParams{
		LoanCoin:       "",
		CollateralCoin: "",
		StartTime:      0,
		EndTime:        0,
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleAdjustLtv(t *testing.T) {
	testConfig(bnc.FlexibleLoanAdjustLtvConfig, bnc.FlexibleAdjustLtvParams{
		LoanCoin:         "USDT",
		CollateralCoin:   "ETH",
		AdjustmentAmount: 0.05,
		Direction:        bnc.LTVAdDireReduced,
	})
}

func TestFlexibleAdjustLtvHistories(t *testing.T) {
	testConfig(bnc.FlexibleAdjustLtvHistoriesConfig, bnc.FlexibleAdjustLtvHistoriesParams{
		LoanCoin:       "USDT",
		CollateralCoin: "ETH",
		StartTime:      0,
		EndTime:        0,
		Current:        0,
		Limit:          0,
	})
}

func TestFlexibleLoanAssets(t *testing.T) {
	testConfig(bnc.FlexibleLoanAssetsConfig, bnc.FlexibleLoanAssetsParams{
		LoanCoin: "",
	})
}

func TestFlexibleCollateralCoins(t *testing.T) {
	testConfig(bnc.FlexibleCollateralCoinsConfig, bnc.FlexibleCollateralCoinsParams{
		CollateralCoin: "",
	})
}

func TestNewSpotOrder(t *testing.T) {
	testConfig(bnc.SpotNewOrderConfig, bnc.SpotNewOrderParams{
		Symbol:                  "ETHUSDT",
		Type:                    bnc.OrderTypeLimit,
		Side:                    bnc.OrderSideBuy,
		Quantity:                0.01,
		Price:                   1500,
		TimeInForce:             bnc.TimeInForceGtc,
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
	testConfig(bnc.SpotCancelOrderConfig, bnc.SpotCancelOrderParams{
		Symbol:             "ETHUSDT",
		OrderId:            0,
		OrigClientOrderId:  "",
		NewClientOrderId:   "",
		CancelRestrictions: "",
	})
}

func TestCancelAllSpotOpenOrders(t *testing.T) {
	testConfig(bnc.SpotCancelAllOpenOrdersConfig, bnc.SpotCancelAllOpenOrdersParams{
		Symbol: "ETHUSDT",
	})
}
