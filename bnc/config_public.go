package bnc

type OrderBookParams struct {
	Symbol string `s2m:"symbol,omitempty"`
	Limit  int    `s2m:"limit,omitempty"` // default 100, max 5000
}

type RawOrderBook struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type OrderBook struct {
	LastUpdateId int64
	Bids         [][]float64
	Asks         [][]float64
}
