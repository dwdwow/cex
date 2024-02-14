package bnc

import (
	"testing"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/props"
)

func testPubConfig[ReqDataType, RespDataType any](
	config cex.ReqConfig[ReqDataType, RespDataType],
	reqData ReqDataType,
	opts ...cex.ReqOpt,
) {
	resp, ob, err := cex.Request(EmptyUser(), config, reqData, opts...)
	_ = resp
	props.PanicIfNotNil(err)
	props.PrintlnIndent(ob)
}

func TestSpotOrderBook(t *testing.T) {
	testPubConfig(SpotOrderBookConfig, OrderBookParams{
		Symbol: "ETHUSDT",
		Limit:  0,
	})
}

func TestFuturesOrderBook(t *testing.T) {
	testPubConfig(FuturesOrderBookConfig, OrderBookParams{
		Symbol: "ETHUSDT",
		Limit:  0,
	})
}

func TestSpotExchangeInfo(t *testing.T) {
	testPubConfig(SpotExchangeInfosConfig, nil)
}

func TestFuturesExchangeInfo(t *testing.T) {
	testPubConfig(FuturesExchangeInfosConfig, nil)
}

func TestFuturesFundingRateHistories(t *testing.T) {
	testPubConfig(FuturesFundingRateHistoriesConfig, FuturesFundingRateHistoriesParams{
		Symbol:    "",
		StartTime: 0,
		EndTime:   0,
		Limit:     0,
	})
}

func TestFuturesFundingRateInfos(t *testing.T) {
	testPubConfig(FuturesFundingRateInfosConfig, nil)
}

func TestFuturesFundingRates(t *testing.T) {
	testPubConfig(FuturesFundingRatesConfig, FuturesFundingRatesParams{Symbol: ""})
}
