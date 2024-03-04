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

// RawKline is binance raw kline data, floats are all string, and time is int64
type RawKline [12]any

// StrKline is data read from csv file
type StrKline [12]string

// SimpleKline is convenient to save in file
type SimpleKline [11]float64

func NewSimpleKlineFromStr(k StrKline) (kline SimpleKline, err error) {
	for i, v := range k[:11] {
		if kline[i], err = strconv.ParseFloat(v, 64); err != nil {
			return
		}
	}
	return kline, nil
}

func NewSimpleKlineFromRaw(k RawKline) (kline SimpleKline, err error) {
	for i, v := range k[:11] {
		switch v := v.(type) {
		case string:
			if kline[i], err = strconv.ParseFloat(v, 64); err != nil {
				return
			}
		case int64:
			kline[i] = float64(v)
		default:
			err = fmt.Errorf("bnc: unknown raw kline data type, %v: %v", i, v)
			return
		}
	}
	return kline, nil
}

func NewSimpleKlineFromStruct(k Kline) (kline SimpleKline) {
	kline[0] = float64(k.OpenTime)
	kline[1] = k.OpenPrice
	kline[2] = k.HighPrice
	kline[3] = k.LowPrice
	kline[4] = k.ClosePrice
	kline[5] = k.Volume
	kline[6] = float64(k.CloseTime)
	kline[7] = k.QuoteAssetVolume
	kline[8] = float64(k.TradesNumber)
	kline[9] = k.TakerBuyBaseAssetVolume
	kline[10] = k.TakerBuyQuoteAssetVolume
	return
}

func (k SimpleKline) NotExist() bool {
	return k[1] == 0
}

func (k SimpleKline) OpenTime() int64 {
	return int64(k[0])
}

func (k SimpleKline) CloseTime() int64 {
	return int64(k[6])
}

func (k SimpleKline) TradesNumber() int64 {
	return int64(k[8])
}

func (k SimpleKline) OpenPrice() float64 {
	return k[1]
}

func (k SimpleKline) HighPrice() float64 {
	return k[2]
}

func (k SimpleKline) LowPrice() float64 {
	return k[3]
}

func (k SimpleKline) ClosePrice() float64 {
	return k[4]
}

func (k SimpleKline) Volume() float64 {
	return k[5]
}

func (k SimpleKline) QuoteAssetVolume() float64 {
	return k[7]
}

func (k SimpleKline) TakerBuyBaseAssetVolume() float64 {
	return k[9]
}

func (k SimpleKline) TakerBuyQuoteAssetVolume() float64 {
	return k[10]
}
