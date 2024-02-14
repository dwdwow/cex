package cex

import "github.com/go-resty/resty/v2"

type SimpleTraderFunc func(TradeType, TradeSide, string, string, float64, float64) (*resty.Response, *Order, *RequestError)
type LimitTraderFunc func(string, string, float64, float64) (*resty.Response, *Order, *RequestError)
type MarketTraderFunc func(string, string, float64) (*resty.Response, *Order, *RequestError)

type SpotTrader interface {
	SpotTrade(tradeType TradeType, side TradeSide, asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	SpotLimitBuy(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	SpotLimitSell(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	SpotMarketBuy(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
	SpotMarketSell(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
}

type FuTrader interface {
	FuTrade(tradeType TradeType, side TradeSide, asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	FuLimitBuy(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	FuLimitSell(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	FuMarketBuy(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
	FuMarketSell(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
}

type Trader interface {
	Trade(*Order) (*resty.Response, *RequestError)
	Query(*Order) (*resty.Response, *RequestError)
	Cancel(*Order) (*resty.Response, *RequestError)
	Wait(*Order) (*resty.Response, *RequestError)
	SpotTrader
	FuTrader
}
