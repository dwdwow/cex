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

// QueryFundingRateHistories
// limit, default 100, max 1000
func QueryFundingRateHistories(symbol string, startTime, endTime int64, limit int) ([]FuturesFundingRateHistory, error) {
	return queryInfoAboutFundingRate(FuturesFundingRateHistoriesConfig, FuturesFundingRateHistoriesParams{Symbol: symbol, StartTime: startTime, EndTime: endTime, Limit: limit})
}

// QueryFundingRateInfos
// Query funding rate info for symbols that had FundingRateCap/ FundingRateFloor / fundingIntervalHours adjustment
// Be careful! Some symbols have no funding rate info!!!
func QueryFundingRateInfos() ([]FuturesFundingRateInfo, error) {
	return queryInfoAboutFundingRate(FuturesFundingRateInfosConfig, nil)
}

func QueryAllFundingRateInfos() ([]FuturesFundingRateInfo, error) {
	var result []FuturesFundingRateInfo
	frInfos, err := QueryFundingRateInfos()
	if err != nil {
		return nil, err
	}
	frInfoBySyb := map[string]FuturesFundingRateInfo{}
	for _, info := range frInfos {
		frInfoBySyb[info.Symbol] = info
	}
	futuresExchangeInfo, err := QueryFuturesExchangeInfo()
	if err != nil {
		return nil, err
	}
	exchanges := futuresExchangeInfo.Symbols
	for _, ex := range exchanges {
		info, ok := frInfoBySyb[ex.Symbol]
		if ok {
			result = append(result, info)
			continue
		}
		info = FuturesFundingRateInfo{
			Symbol:                   ex.Symbol,
			AdjustedFundingRateCap:   0,
			AdjustedFundingRateFloor: 0,
			FundingIntervalHours:     8,
			Disclaimer:               false,
		}
	}
	return result, nil
}

func QueryFundingRates() ([]FuturesFundingRate, error) {
	return queryInfoAboutFundingRate(FuturesFundingRatesConfig, FuturesFundingRatesParams{Symbol: ""})
}
