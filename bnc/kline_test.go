package bnc

import (
	"strconv"
	"testing"

	"github.com/dwdwow/props"
)

func TestKlineStringToAny(t *testing.T) {
	strKline := []string{
		"2697584323480",
		"125165.3425",
		"3214.1345",
		"154163.1551",
		"5345432.145",
		"15151.1564253",
		"14535431515",
		"1453.13425",
		"15351",
		"125512",
		"3214554.1345",
		"",
	}

	var rawKline []any

	for i, v := range strKline {
		if i == 0 || i == 6 || i == 8 {
			j, err := strconv.ParseInt(v, 10, 64)
			props.PanicIfNotNil(err)
			rawKline = append(rawKline, j)
		} else {
			rawKline = append(rawKline, v)
		}
	}

	structKline, err := UnmarshalRawKline(rawKline)
	props.PanicIfNotNil(err)

	kline, err := KlineStringToAny(strKline)

	props.PanicIfNotNil(err)

	if KlineOpenTime(kline) != structKline.OpenTime {
		panic("open time")
	}

	if KlineCloseTime(kline) != structKline.CloseTime {
		panic("close time")
	}

	if KlineTradesNumber(kline) != structKline.TradesNumber {
		panic("trades number")
	}

	if KlineOpenPrice(kline) != structKline.OpenPrice {
		panic("open price")
	}

	if KlineHighPrice(kline) != structKline.HighPrice {
		panic("high price")
	}

	if KlineLowPrice(kline) != structKline.LowPrice {
		panic("low price")
	}

	if KlineClosePrice(kline) != structKline.ClosePrice {
		panic("close price")
	}

	if KlineVolume(kline) != structKline.Volume {
		panic("volume")
	}

	if KlineQuoteAssetVolume(kline) != structKline.QuoteAssetVolume {
		panic("quote asset volume")
	}

	if KlineTakerBuyBaseAssetVolume(kline) != structKline.TakerBuyBaseAssetVolume {
		panic("taker buy base asset volume")
	}

	if KlineTakerBuyQuoteAssetVolume(kline) != structKline.TakerBuyQuoteAssetVolume {
		panic("taker buy quote asset volume")
	}

}
