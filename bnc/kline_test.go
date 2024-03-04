package bnc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/dwdwow/props"
)

func TestNewSimpleKline(t *testing.T) {
	strKline := StrKline{
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

	var rawKline RawKline

	for i, v := range strKline {
		if i == 0 || i == 6 || i == 8 {
			j, err := strconv.ParseInt(v, 10, 64)
			props.PanicIfNotNil(err)
			rawKline[i] = j
		} else {
			rawKline[i] = v
		}
	}

	structKline, err := UnmarshalRawKline(rawKline)
	props.PanicIfNotNil(err)

	klineFromStr, err := NewSimpleKlineFromStr(strKline)
	props.PanicIfNotNil(err)

	klineFromRaw, err := NewSimpleKlineFromRaw(rawKline)
	props.PanicIfNotNil(err)

	klineFromStruct := NewSimpleKlineFromStruct(structKline)

	for i, v := range klineFromStr {
		if klineFromRaw[i] != v {
			panic("from str != from raw")
		}
		if klineFromStruct[i] != v {
			panic("from struct != from raw")
		}
	}

	if klineFromStr.OpenTime() != structKline.OpenTime {
		panic("open time")
	}

	if klineFromStr.CloseTime() != structKline.CloseTime {
		panic("close time")
	}

	if klineFromStr.TradesNumber() != structKline.TradesNumber {
		panic("trades number")
	}

	if klineFromStr.OpenPrice() != structKline.OpenPrice {
		panic("open price")
	}

	if klineFromStr.HighPrice() != structKline.HighPrice {
		panic("high price")
	}

	if klineFromStr.LowPrice() != structKline.LowPrice {
		panic("low price")
	}

	if klineFromStr.ClosePrice() != structKline.ClosePrice {
		panic("close price")
	}

	if klineFromStr.Volume() != structKline.Volume {
		panic("volume")
	}

	if klineFromStr.QuoteAssetVolume() != structKline.QuoteAssetVolume {
		panic("quote asset volume")
	}

	if klineFromStr.TakerBuyBaseAssetVolume() != structKline.TakerBuyBaseAssetVolume {
		panic("taker buy base asset volume")
	}

	if klineFromStr.TakerBuyQuoteAssetVolume() != structKline.TakerBuyQuoteAssetVolume {
		panic("taker buy quote asset volume")
	}

}

// TestNotExist tests the NotExist method of the SimpleKline type
func TestSimpleKline_NotExist(t *testing.T) {
	tests := []struct {
		name string
		k    SimpleKline
		want bool
	}{
		{
			name: "not_exist",
			k:    SimpleKline{1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: true,
		},
		{
			name: "exist",
			k:    SimpleKline{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: false,
		},
		{
			name: "empty_kline",
			k:    SimpleKline{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.NotExist(); got != tt.want {
				t.Errorf("NotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleKlineJson(t *testing.T) {
	kline := SimpleKline{1111, 222, 333.1235, 12351351.1346, 0, -1, -1235.2345}
	data, err := json.Marshal(kline)
	props.PanicIfNotNil(err)
	if string(data) != `[1111,222,333.1235,12351351.1346,0,-1,-1235.2345,0,0,0,0]` {
		fmt.Println(string(data))
		t.FailNow()
	}
	kline2 := SimpleKline{}
	err = json.Unmarshal(data, &kline2)
	props.PanicIfNotNil(err)
	if reflect.DeepEqual(kline2, kline) {
		fmt.Println(kline2)
		t.FailNow()
	}
}
