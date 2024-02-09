package cex

type SimpleTraderFunc func(TradeType, TradeSide, string, string, float64, float64) (*Order, error)
type LimitTraderFunc func(string, string, float64, float64) (*Order, error)
type MarketTraderFunc func(string, string, float64) (*Order, error)

type SpotTrader interface {
	SpotTrade(tradeType TradeType, side TradeSide, asset, quote string, qty, price float64) (*Order, error)
	SpotLimitBuy(asset, quote string, qty, price float64) (*Order, error)
	SpotLimitSell(asset, quote string, qty, price float64) (*Order, error)
	SpotMarketBuy(asset, quote string, qty float64) (*Order, error)
	SpotMarketSell(asset, quote string, qty float64) (*Order, error)
}

type FuTrader interface {
	FuTrade(tradeType TradeType, side TradeSide, asset, quote string, qty, price float64) (*Order, error)
	FuLimitBuy(asset, quote string, qty, price float64) (*Order, error)
	FuLimitSell(asset, quote string, qty, price float64) (*Order, error)
	FuMarketBuy(asset, quote string, qty float64) (*Order, error)
	FuMarketSell(asset, quote string, qty float64) (*Order, error)
}

type Trader interface {
	Trade(*Order) error
	Update(*Order) error
	Cancel(*Order) error
	Wait(*Order) error
	SpotTrader
	FuTrader
}
