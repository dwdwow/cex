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

func KlineStringToAny(k []string) ([]any, error) {
	if len(k) != 12 {
		return nil, fmt.Errorf("bnc: string kline to any kline, length %v != 12", len(k))
	}
	var err error
	kline := make([]any, 11)
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

func KlineOpenTime(kline []any) int64 {
	return kline[0].(int64)
}

func KlineCloseTime(kline []any) int64 {
	return kline[6].(int64)
}

func KlineTradesNumber(kline []any) int64 {
	return kline[8].(int64)
}

func KlineOpenPrice(kline []any) float64 {
	return kline[1].(float64)
}

func KlineHighPrice(kline []any) float64 {
	return kline[2].(float64)
}

func KlineLowPrice(kline []any) float64 {
	return kline[3].(float64)
}

func KlineClosePrice(kline []any) float64 {
	return kline[4].(float64)
}

func KlineVolume(kline []any) float64 {
	return kline[5].(float64)
}

func KlineQuoteAssetVolume(kline []any) float64 {
	return kline[7].(float64)
}

func KlineTakerBuyBaseAssetVolume(kline []any) float64 {
	return kline[9].(float64)
}

func KlineTakerBuyQuoteAssetVolume(kline []any) float64 {
	return kline[10].(float64)
}
