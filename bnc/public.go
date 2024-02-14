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
