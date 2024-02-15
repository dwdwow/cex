package bnc

import (
	"testing"

	"github.com/dwdwow/props"
)

func publicTestChecker(v any, err error) {
	props.PanicIfNotNil(err)
	props.PrintlnIndent(v)
}

func TestQuerySpotOrderBook(t *testing.T) {
	publicTestChecker(QuerySpotOrderBook("ETHUSDT", 10))
}

func TestQueryFuturesOrderBook(t *testing.T) {
	publicTestChecker(QueryFuturesOrderBook("ETHUSDT", 10))
}

func TestQuerySpotExchangeInfo(t *testing.T) {
	publicTestChecker(QuerySpotExchangeInfo())
}

func TestQueryFuturesExchangeInfo(t *testing.T) {
	publicTestChecker(QueryFuturesExchangeInfo())
}

func TestQueryFundingRateHistories(t *testing.T) {
	publicTestChecker(QueryFundingRateHistories("ETHUSDT", 0, 0, 0))
}

func TestQueryFundingRateInfos(t *testing.T) {
	publicTestChecker(QueryFundingRateInfos())
}

func TestQueryAllFundingRateInfos(t *testing.T) {
	infos, err := QueryAllFundingRateInfos()
	props.PanicIfNotNil(err)
	for _, info := range infos {
		if info.FundingIntervalHours == 4 {
			props.PrintlnIndent(info)
		}
		//if info.Symbol == "ETHUSDT" {
		//	props.PrintlnIndent(info)
		//	return
		//}
	}
}

func TestQueryFundingRates(t *testing.T) {
	publicTestChecker(QueryFundingRates())
}
