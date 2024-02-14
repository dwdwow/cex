package cex

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type SimpleTraderFunc func(OrderType, OrderSide, string, string, float64, float64) (*resty.Response, *Order, *RequestError)
type LimitTraderFunc func(string, string, float64, float64) (*resty.Response, *Order, *RequestError)
type MarketTraderFunc func(string, string, float64) (*resty.Response, *Order, *RequestError)

type SpotTrader interface {
	NewSpotOrder(asset, quote string, tradeType OrderType, side OrderSide, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewSpotLimitBuyOrder(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewSpotLimitSellOrder(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewSpotMarketBuyOrder(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
	NewSpotMarketSellOrder(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
}

type FuTrader interface {
	NewFutureOrder(asset, quote string, tradeType OrderType, side OrderSide, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewFutureLimitBuyOrder(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewFutureLimitSellOrder(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewFutureMarketBuyOrder(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
	NewFutureMarketSellOrder(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
}

type Trader interface {
	QueryOrder(*Order) (*resty.Response, *RequestError)
	CancelOrder(*Order) (*resty.Response, *RequestError)
	WaitOrder(context.Context, *Order) (*resty.Response, *RequestError)
	SpotTrader
	FuTrader
}
