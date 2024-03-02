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

	if kline.OpenTime() != structKline.OpenTime {
		panic("open time")
	}

	if kline.CloseTime() != structKline.CloseTime {
		panic("close time")
	}

	if kline.TradesNumber() != structKline.TradesNumber {
		panic("trades number")
	}

	if kline.OpenPrice() != structKline.OpenPrice {
		panic("open price")
	}

	if kline.HighPrice() != structKline.HighPrice {
		panic("high price")
	}

	if kline.LowPrice() != structKline.LowPrice {
		panic("low price")
	}

	if kline.ClosePrice() != structKline.ClosePrice {
		panic("close price")
	}

	if kline.Volume() != structKline.Volume {
		panic("volume")
	}

	if kline.QuoteAssetVolume() != structKline.QuoteAssetVolume {
		panic("quote asset volume")
	}

	if kline.TakerBuyBaseAssetVolume() != structKline.TakerBuyBaseAssetVolume {
		panic("taker buy base asset volume")
	}

	if kline.TakerBuyQuoteAssetVolume() != structKline.TakerBuyQuoteAssetVolume {
		panic("taker buy quote asset volume")
	}

}
