package cex

type PairType string

const (
	PairTypeSpot    PairType = "SPOT"
	PairTypeFutures PairType = "FUTURES"
)

func NotPairType(t PairType) bool {
	switch t {
	case PairTypeSpot, PairTypeFutures:
		return false
	}
	return true
}

type Pair struct {
	// must be contained
	Cex        Name     `json:"cex" bson:"cex"`
	Type       PairType `json:"type" bson:"type"`
	Asset      string   `json:"asset" bson:"asset"`
	Quote      string   `json:"quote" bson:"quote"`
	PairSymbol string   `json:"pairSymbol" bson:"pairSymbol"`
	MidSymbol  string   `json:"midSymbol" bson:"midSymbol"`
	QPrecision int      `json:"qPrecision" bson:"qPrecision"`
	PPrecision int      `json:"pPrecision" bson:"pPrecision"`

	// may be omitted
	TakerFeeTier  float64 `json:"takerFeeTier" bson:"takerFeeTier"`
	MakerFeeTier  float64 `json:"makerFeeTier" bson:"makerFeeTier"`
	MinTradeQty   float64 `json:"minTradeQty" bson:"minTradeQty"`
	MinTradeQuote float64 `json:"minTradeQuote" bson:"minTradeQuote"`
	Tradable      bool    `json:"tradable" bson:"tradable"`
	CanMarket     bool    `json:"canMarket" bson:"canMarket"`
	CanMargin     bool    `json:"canMargin" bson:"canMargin"`
	IsCross       bool    `json:"isCross" bson:"isCross"`

	IsPerpetual bool `json:"isPerpetual" bson:"isPerpetual"`
}
