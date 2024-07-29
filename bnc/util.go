package bnc

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dwdwow/cex"
)

// ExchangeInfoToPair switch ExchangeInfo.Symbols to cex.Pair.
// If exchange is contract, should be careful that it is delivery contract.
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
	var isPerpetual bool
	if info.ContractType == "" {
		pairType = cex.PairTypeSpot
		sybMid = SpotSymbolMid
		minTradeQuote = 10
		// TODO not correct
		//makerFeeTier, takerFeeTier = SpotMakerFeeTier, SpotTakerFeeTier
	} else {
		pairType = cex.PairTypeFutures
		sybMid = FuturesSymbolMid
		minTradeQuote = 20
		if info.ContractType == "PERPETUAL" {
			isPerpetual = true
		}
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
		Tradable:      info.Status == ExchangeTrading || info.ContractStatus == ExchangeTrading,
		CanMarket:     true,
		CanMargin:     info.IsMarginTradingAllowed,
		IsPerpetual:   isPerpetual,

		ContractSize: info.ContractSize,
		ContractType: info.ContractType,

		DeliveryDate: info.DeliveryDate,
		OnboardDate:  info.OnboardDate,
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

func getWsEvent(data []byte) (event WsEvent, isArray, ok bool) {
	sd := string(data)
	isArray = strings.Index(sd, "[") == 0
	ss := strings.Split(sd, ",")
	for _, s := range ss {
		r := strings.Split(s, "\"e\":")
		if len(r) == 2 {
			return WsEvent(strings.ReplaceAll(r[1], "\"", "")), isArray, true
		}
	}
	return "", isArray, false
}

func unmarshal[T any](data []byte) (t T, err error) {
	err = json.Unmarshal(data, &t)
	return
}
