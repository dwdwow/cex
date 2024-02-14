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

func TestQueryFundingRates(t *testing.T) {
	publicTestChecker(QueryFundingRates())
}
