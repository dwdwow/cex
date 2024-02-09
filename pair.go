package cex

type PairType string

const (
	SpotPair   PairType = "SPOT"
	FuturePair PairType = "FUTURE"
)

type Pair struct {
	Cex           Cex      `json:"cex" bson:"cex"`
	Type          PairType `json:"type" bson:"type"`
	Asset         string   `json:"asset" bson:"asset"`
	Quote         string   `json:"quote" bson:"quote"`
	PairSymbol    string   `json:"pairSymbol" bson:"pairSymbol"`
	MidSymbol     string   `json:"midSymbol" bson:"midSymbol"`
	QPrecision    int      `json:"qPrecision" bson:"qPrecision"`
	PPrecision    int      `json:"pPrecision" bson:"pPrecision"`
	TakerFeeTier  float64  `json:"takerFeeTier" bson:"takerFeeTier"`
	MakerFeeTier  float64  `json:"makerFeeTier" bson:"makerFeeTier"`
	MinTradeQty   float64  `json:"minTradeQty" bson:"minTradeQty"`
	MinTradeQuote float64  `json:"minTradeQuote" bson:"minTradeQuote"`
	Tradable      bool     `json:"tradable" bson:"tradable"`
	CanMarket     bool     `json:"canMarket" bson:"canMarket"`
	CanMargin     bool     `json:"canMargin" bson:"canMargin"`
	IsCross       bool     `json:"isCross" bson:"isCross"`
}