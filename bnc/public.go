package bnc

import (
	"errors"

	"github.com/dwdwow/cex"
)

func queryOrderBook(config cex.ReqConfig[OrderBookParams, OrderBook], symbol string, limit int) (OrderBook, error) {
	_, ob, err := cex.Request(emptyUser, config, OrderBookParams{symbol, limit})
	if err.IsNotNil() {
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
	_, info, err := cex.Request(emptyUser, config, nil)
	if err.IsNotNil() {
		return ExchangeInfo{}, errors.New(err.Error())
	}
	return info, nil
}

func QuerySpotExchangeInfo() (ExchangeInfo, error) {
	return queryExchangeInfo(SpotExchangeInfosConfig)
}

// QueryFuturesExchangeInfo
// Deprecated: use QueryUMFuturesExchangeInfo instead
func QueryFuturesExchangeInfo() (ExchangeInfo, error) {
	return queryExchangeInfo(FuturesExchangeInfosConfig)
}

func QueryUMFuturesExchangeInfo() (ExchangeInfo, error) {
	return queryExchangeInfo(FuturesExchangeInfosConfig)
}

func QueryCMFuturesExchangeInfo() (ExchangeInfo, error) {
	return queryExchangeInfo(CMFuturesExchangeInfosConfig)
}

func queryPairs(exInfoQuerier func() (ExchangeInfo, error)) (pairs []cex.Pair, info ExchangeInfo, err error) {
	info, err = exInfoQuerier()
	if err != nil {
		return
	}
	for _, syb := range info.Symbols {
		var pair cex.Pair
		pair, err = ExchangeInfoToPair(syb)
		if err != nil {
			return
		}
		pairs = append(pairs, pair)
	}
	return
}

func QuerySpotPairs() ([]cex.Pair, ExchangeInfo, error) {
	return queryPairs(QuerySpotExchangeInfo)
}

func QueryFuturesPairs() ([]cex.Pair, ExchangeInfo, error) {
	return queryPairs(QueryFuturesExchangeInfo)
}

func QueryCMFuturesPairs() ([]cex.Pair, ExchangeInfo, error) {
	return queryPairs(QueryCMFuturesExchangeInfo)
}

func queryInfoAboutFundingRate[Req any, Resp any](config cex.ReqConfig[Req, Resp], params Req) (Resp, error) {
	_, resp, err := cex.Request(emptyUser, config, params)
	if err.IsNotNil() {
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
	var err error
	result, err = QueryFundingRateInfos()
	if err != nil {
		return nil, err
	}
	frInfoBySyb := map[string]FuturesFundingRateInfo{}
	for _, info := range result {
		frInfoBySyb[info.Symbol] = info
	}
	futuresExchangeInfo, err := QueryFuturesExchangeInfo()
	if err != nil {
		return nil, err
	}
	for _, ex := range futuresExchangeInfo.Symbols {
		info, ok := frInfoBySyb[ex.Symbol]
		if !ok {
			info = FuturesFundingRateInfo{
				Symbol:                   ex.Symbol,
				AdjustedFundingRateCap:   0,
				AdjustedFundingRateFloor: 0,
				FundingIntervalHours:     8,
				Disclaimer:               false,
			}
			result = append(result, info)
		}
	}
	return result, nil
}

func QueryFundingRates() ([]FuturesFundingRate, error) {
	return queryInfoAboutFundingRate(FuturesFundingRatesConfig, FuturesFundingRatesParams{Symbol: ""})
}

func queryKline(config cex.ReqConfig[KlineParams, []Kline], symbol string, interval KlineInterval, start, end, limit int64) ([]Kline, error) {
	_, res, err := cex.Request(emptyUser, config, KlineParams{
		Symbol:    symbol,
		Interval:  interval,
		StartTime: start,
		EndTime:   end,
		TimeZone:  "",
		Limit:     limit,
	})
	return res, err.Err
}

func QuerySpotKline(symbol string, interval KlineInterval, start, end int64) ([]Kline, error) {
	return queryKline(SpotKlineConfig, symbol, interval, start, end, 1000)
}

func QuerySpotKlineWithLimit(symbol string, interval KlineInterval, start, end, limit int64) ([]Kline, error) {
	return queryKline(SpotKlineConfig, symbol, interval, start, end, limit)
}

func QueryFuturesKline(symbol string, interval KlineInterval, start, end int64) ([]Kline, error) {
	return queryKline(FuturesKlineConfig, symbol, interval, start, end, 1000)
}

func QueryFuturesKlineWithLimit(symbol string, interval KlineInterval, start, end, limit int64) ([]Kline, error) {
	return queryKline(FuturesKlineConfig, symbol, interval, start, end, limit)
}

func QueryPortfolioMarginCollateralRates() ([]PortfolioMarginCollateralRate, error) {
	_, data, reqErr := cex.Request(emptyUser, PortfolioMarginCollateralRatesConfig, nil)
	if reqErr.IsNotNil() {
		return nil, reqErr.Err
	}
	return data.Data, nil
}

func QuerySpotPrices() ([]SpotPriceTicker, error) {
	_, data, reqErr := cex.Request(emptyUser, SpotPricesConfig, nil)
	if reqErr.IsNotNil() {
		return nil, reqErr.Err
	}
	return data, nil
}

func QueryFuturesPrices() ([]FuturesPriceTicker, error) {
	_, data, reqErr := cex.Request(emptyUser, FuturesPricesConfig, nil)
	if reqErr.IsNotNil() {
		return nil, reqErr.Err
	}
	return data, nil
}

func QueryCMPremiumIndex(symbol, pair string) ([]CMPremiumIndex, error) {
	_, data, reqErr := cex.Request(emptyUser, CMPremiumIndexConfig, CMPremiumIndexParams{
		Symbol: symbol,
		Pair:   pair,
	})
	if reqErr.IsNotNil() {
		return nil, reqErr.Err
	}
	return data, nil
}
