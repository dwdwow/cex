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

type FastKline []any

func KlineStringToAny(k []string) (FastKline, error) {
	if len(k) != 12 {
		return nil, fmt.Errorf("bnc: string kline to any kline, length %v != 12", len(k))
	}
	var err error
	kline := make(FastKline, 11)
	if kline[0], err = strconv.ParseInt(k[0], 10, 64); err != nil {
		return nil, err
	}
	if kline[1], err = strconv.ParseFloat(k[1], 64); err != nil {
		return nil, err
	}
	if kline[2], err = strconv.ParseFloat(k[2], 64); err != nil {
		return nil, err
	}
	if kline[3], err = strconv.ParseFloat(k[3], 64); err != nil {
		return nil, err
	}
	if kline[4], err = strconv.ParseFloat(k[4], 64); err != nil {
		return nil, err
	}
	if kline[5], err = strconv.ParseFloat(k[5], 64); err != nil {
		return nil, err
	}
	if kline[6], err = strconv.ParseInt(k[6], 10, 64); err != nil {
		return nil, err
	}
	if kline[7], err = strconv.ParseFloat(k[7], 64); err != nil {
		return nil, err
	}
	if kline[8], err = strconv.ParseInt(k[8], 10, 64); err != nil {
		return nil, err
	}
	if kline[9], err = strconv.ParseFloat(k[9], 64); err != nil {
		return nil, err
	}
	if kline[10], err = strconv.ParseFloat(k[10], 64); err != nil {
		return nil, err
	}
	return kline, nil
}

func (k FastKline) OpenTime() int64 {
	return k[0].(int64)
}

func (k FastKline) CloseTime() int64 {
	return k[6].(int64)
}

func (k FastKline) TradesNumber() int64 {
	return k[8].(int64)
}

func (k FastKline) OpenPrice() float64 {
	return k[1].(float64)
}

func (k FastKline) HighPrice() float64 {
	return k[2].(float64)
}

func (k FastKline) LowPrice() float64 {
	return k[3].(float64)
}

func (k FastKline) ClosePrice() float64 {
	return k[4].(float64)
}

func (k FastKline) Volume() float64 {
	return k[5].(float64)
}

func (k FastKline) QuoteAssetVolume() float64 {
	return k[7].(float64)
}

func (k FastKline) TakerBuyBaseAssetVolume() float64 {
	return k[9].(float64)
}

func (k FastKline) TakerBuyQuoteAssetVolume() float64 {
	return k[10].(float64)
}
