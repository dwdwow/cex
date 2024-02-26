package bnc

import (
	"sync"
	"testing"
	"time"

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

func TestQuerySpotPairs(t *testing.T) {
	pairs, _, err := QuerySpotPairs()
	props.PanicIfNotNil(err)
	props.PrintlnIndent(pairs)
}

func TestQueryFuturesPairs(t *testing.T) {
	pairs, _, err := QueryFuturesPairs()
	props.PanicIfNotNil(err)
	props.PrintlnIndent(pairs)
	for _, pair := range pairs {
		if !pair.IsPerpetual {
			t.Log(pair.PairSymbol)
		}
	}
}

func TestQueryFuturesExchangeInfo(t *testing.T) {
	publicTestChecker(QueryFuturesExchangeInfo())
}

func TestQueryFundingRateHistories(t *testing.T) {
	publicTestChecker(QueryFundingRateHistories("RNDRUSDT", 0, 0, 0))
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

func TestQueryKline(t *testing.T) {
	now := time.Now().Unix() * 1000
	start := now - 10*1000
	end := now - 1000
	res, err := QuerySpotKline("ETHUSDT", "1s", 0, end)
	t.Log(start, end, res)
	props.PanicIfNotNil(err)
}

func TestQueryKlineAsync(t *testing.T) {
	exchange, err := QuerySpotExchangeInfo()
	props.PanicIfNotNil(err)
	sybs := exchange.Symbols
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		for _, syb := range sybs {
			wg.Add(1)
			syb := syb
			go func() {
				now := time.Now().UnixMilli()
				var err error
				for i := 0; i < 3; i++ {
					_, err = QuerySpotKline(syb.Symbol, "1m", now-time.Hour.Milliseconds(), now)
					if err == nil {
						break
					}
				}
				props.PanicIfNotNil(err)
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
