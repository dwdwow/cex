package bnc

import (
	"fmt"
	"strings"

	"github.com/dwdwow/cex"
)

func ExchangeInfoToPair(info Exchange) (cex.Pair, error) {
	filters := info.Filters
	var pPrec, qPrec int
	for _, filter := range filters {
		t := filter["filterType"]
		switch t := t.(type) {
		case string:
			if t == "PRICE_FILTER" {
				tick := filter["tickSize"]
				switch tick := tick.(type) {
				case string:
					s, err := GetPrecJustForBinanceFilter(tick)
					if err != nil {
						return cex.Pair{}, err
					}
					pPrec = s
				default:
					// should not get here
					return cex.Pair{}, fmt.Errorf("exchange info tickSize type is not string, tick size %v", tick)
				}
			} else if t == "LOT_SIZE" {
				step := filter["stepSize"]
				switch step := step.(type) {
				case string:
					s, err := GetPrecJustForBinanceFilter(step)
					if err != nil {
						return cex.Pair{}, err
					}
					qPrec = s
				default:
					// should not get here
					return cex.Pair{}, fmt.Errorf("exchange info stepSize type is not string, step size %v", step)
				}
			}
		default:
			// should not get here
			return cex.Pair{}, fmt.Errorf("exchange info filter type is not string, type %v", t)
		}
	}
	var pairType cex.PairType
	var sybMid string
	var makerFeeTier, takerFeeTier, minTradeQuote float64
	if info.ContractType == "" {
		pairType = cex.SpotPair
		sybMid = SpotSymbolMid
		minTradeQuote = 10
		// TODO not correct
		//makerFeeTier, takerFeeTier = SpotMakerFeeTier, SpotTakerFeeTier
	} else {
		pairType = cex.FuturePair
		sybMid = FuturesSymbolMid
		minTradeQuote = 20
		// TODO not correct
		//makerFeeTier, takerFeeTier = FuturesMakerFeeTier, FuturesTakerFeeTier
	}
	pair := cex.Pair{
		Cex:           cex.BINANCE,
		Type:          pairType,
		Asset:         info.BaseAsset,
		Quote:         info.QuoteAsset,
		PairSymbol:    info.Symbol,
		MidSymbol:     sybMid,
		QPrecision:    qPrec,
		PPrecision:    pPrec,
		TakerFeeTier:  takerFeeTier,
		MakerFeeTier:  makerFeeTier,
		MinTradeQty:   0,
		MinTradeQuote: minTradeQuote,
		Tradable:      info.Status == ExchangeTrading,
		CanMarket:     true,
		CanMargin:     info.IsMarginTradingAllowed,
	}
	return pair, nil
}

func GetPrecJustForBinanceFilter(size string) (int, error) {
	ss := strings.Split(size, ".")
	s0 := ss[0]
	i1 := strings.Index(s0, "1")
	if i1 != -1 {
		return i1 - len(s0) + 1, nil
	}
	if len(ss) == 1 {
		return 0, fmt.Errorf("unknown size: %v", size)
	}
	s1 := ss[1]
	f1 := strings.Index(s1, "1")
	if f1 != -1 {
		return f1 + 1, nil
	}
	return 0, fmt.Errorf("unknown size: %v", size)
}
