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
	NewFuturesOrder(asset, quote string, tradeType OrderType, side OrderSide, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewFuturesLimitBuyOrder(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewFuturesLimitSellOrder(asset, quote string, qty, price float64) (*resty.Response, *Order, *RequestError)
	NewFuturesMarketBuyOrder(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
	NewFuturesMarketSellOrder(asset, quote string, qty float64) (*resty.Response, *Order, *RequestError)
}

type Trader interface {
	QueryOrder(*Order) (*resty.Response, *RequestError)
	CancelOrder(*Order) (*resty.Response, *RequestError)
	WaitOrder(context.Context, *Order) *RequestError
	SpotTrader
	FuTrader
}
