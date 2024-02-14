package bnc

import (
	"errors"

	"github.com/dwdwow/cex"
)

func queryOrderBook(config cex.ReqConfig[OrderBookParams, OrderBook], symbol string, limit int) (OrderBook, error) {
	_, ob, err := cex.Request(emptyUser, config, OrderBookParams{symbol, limit})
	if err != nil {
		return OrderBook{}, errors.New(err.Error())
	}
	return ob, nil
}

func QuerySpotOrderBook(symbol string, limit int) (OrderBook, error) {
	return queryOrderBook(SpotOrderBookConfig, symbol, limit)
}

func QueryFuturesOrderBook(symbol string, limit int) (OrderBook, error) {
	return queryOrderBook(FuturesOrderBookConfig, symbol, limit)
}

func queryExchangeInfo(config cex.ReqConfig[cex.NilReqData, ExchangeInfo]) (ExchangeInfo, error) {
	_, ob, err := cex.Request(emptyUser, config, nil)
	if err != nil {
		return ExchangeInfo{}, errors.New(err.Error())
	}
	return ob, nil
}

func QuerySpotExchangeInfo() (ExchangeInfo, error) {
	return queryExchangeInfo(SpotExchangeInfosConfig)
}

func QueryFuturesExchangeInfo() (ExchangeInfo, error) {
	return queryExchangeInfo(FuturesExchangeInfosConfig)
}

func queryInfoAboutFundingRate[Req any, Resp any](config cex.ReqConfig[Req, Resp], params Req) (Resp, error) {
	_, resp, err := cex.Request(emptyUser, config, params)
	if err != nil {
		return resp, errors.New(err.Error())
	}
	return resp, nil
}

func QueryFundingRateHistories(symbol string, startTime, endTime int64, limit int) ([]FuturesFundingRateHistory, error) {
	return queryInfoAboutFundingRate(FuturesFundingRateHistoriesConfig, FuturesFundingRateHistoriesParams{Symbol: symbol, StartTime: startTime, EndTime: endTime, Limit: limit})
}

// QueryFundingRateInfos
// Query funding rate info for symbols that had FundingRateCap/ FundingRateFloor / fundingIntervalHours adjustment
// Be careful! Some symbols have no funding rate info!!!
func QueryFundingRateInfos() ([]FuturesFundingRateInfo, error) {
	return queryInfoAboutFundingRate(FuturesFundingRateInfosConfig, nil)
}

func QueryFundingRates() ([]FuturesFundingRate, error) {
	return queryInfoAboutFundingRate(FuturesFundingRatesConfig, FuturesFundingRatesParams{Symbol: ""})
}
