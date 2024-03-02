package bnc

import (
	"fmt"
	"strconv"
)

var klineMapKeys = []string{
	"openTime",
	"openPrice",
	"highPrice",
	"lowPrice",
	"closePrice",
	"volume",
	"closeTime",
	"quoteAssetVolume",
	"tradesNumber",
	"takerBuyBaseAssetVolume",
	"takerBuyQuoteAssetVolume",
	"unused",
}

// SimpleKline is convenient to save in file
type SimpleKline []any

func KlineStringToAny(k []string) (SimpleKline, error) {
	if len(k) != 12 {
		return nil, fmt.Errorf("bnc: string kline to any kline, length %v != 12", len(k))
	}
	var err error
	kline := make(SimpleKline, 11)
	for i, v := range k[:11] {
		switch i {
		case 0, 6, 8:
			if kline[i], err = strconv.ParseInt(v, 10, 64); err != nil {
				return nil, err
			}
		default:
			if kline[i], err = strconv.ParseFloat(v, 64); err != nil {
				return nil, err
			}
		}
	}
	return kline, nil
}

func (k SimpleKline) OpenTime() int64 {
	return k[0].(int64)
}

func (k SimpleKline) CloseTime() int64 {
	return k[6].(int64)
}

func (k SimpleKline) TradesNumber() int64 {
	return k[8].(int64)
}

func (k SimpleKline) OpenPrice() float64 {
	return k[1].(float64)
}

func (k SimpleKline) HighPrice() float64 {
	return k[2].(float64)
}

func (k SimpleKline) LowPrice() float64 {
	return k[3].(float64)
}

func (k SimpleKline) ClosePrice() float64 {
	return k[4].(float64)
}

func (k SimpleKline) Volume() float64 {
	return k[5].(float64)
}

func (k SimpleKline) QuoteAssetVolume() float64 {
	return k[7].(float64)
}

func (k SimpleKline) TakerBuyBaseAssetVolume() float64 {
	return k[9].(float64)
}

func (k SimpleKline) TakerBuyQuoteAssetVolume() float64 {
	return k[10].(float64)
}
