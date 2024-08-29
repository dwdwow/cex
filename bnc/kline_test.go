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

// TestSimpleKlineToCSVRow tests the ToCSVRow function of SimpleKline
func TestSimpleKlineToCSVRow(t *testing.T) {
	tests := []struct {
		name     string
		kline    SimpleKline
		expected string
	}{
		{
			name: "Standard kline",
			kline: SimpleKline{
				1609459200000, // OpenTime (2021-01-01 00:00:00 UTC)
				30000.0,       // OpenPrice
				31000.0,       // HighPrice
				29500.0,       // LowPrice
				30500.0,       // ClosePrice
				100.5,         // Volume
				1609462800000, // CloseTime (2021-01-01 01:00:00 UTC)
				3065250.0,     // QuoteAssetVolume
				1000,          // TradesNumber
				60.3,          // TakerBuyBaseAssetVolume
				1839150.0,     // TakerBuyQuoteAssetVolume
			},
			expected: "1609459200000,30000,31000,29500,30500,100.5,1609462800000,3065250,1000,60.3,1839150",
		},
		{
			name: "Kline with zero values",
			kline: SimpleKline{
				1609459200000, // OpenTime
				0,             // OpenPrice
				0,             // HighPrice
				0,             // LowPrice
				0,             // ClosePrice
				0,             // Volume
				1609462800000, // CloseTime
				0,             // QuoteAssetVolume
				0,             // TradesNumber
				0,             // TakerBuyBaseAssetVolume
				0,             // TakerBuyQuoteAssetVolume
			},
			expected: "1609459200000,0,0,0,0,0,1609462800000,0,0,0,0",
		},
		{
			name: "Kline with large numbers",
			kline: SimpleKline{
				1609459200000,    // OpenTime
				1000000.123456,   // OpenPrice
				2000000.234567,   // HighPrice
				500000.345678,    // LowPrice
				1500000.456789,   // ClosePrice
				1000000000.56789, // Volume
				1609462800000,    // CloseTime
				1500000000.67891, // QuoteAssetVolume
				1000000,          // TradesNumber
				500000000.78912,  // TakerBuyBaseAssetVolume
				750000000.89123,  // TakerBuyQuoteAssetVolume
			},
			expected: "1609459200000,1000000.123456,2000000.234567,500000.345678,1500000.456789,1000000000.56789,1609462800000,1500000000.67891,1000000,500000000.78912,750000000.89123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.kline.ToCSVRow()
			if result != tt.expected {
				t.Errorf("ToCSVRow() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCSVDataToSimpleKlines(t *testing.T) {
	tests := []struct {
		name     string
		csvData  []byte
		expected []SimpleKline
		wantErr  bool
	}{
		{
			name: "Valid CSV data with header",
			csvData: []byte(`openTime,openPrice,highPrice,lowPrice,closePrice,volume,closeTime,quoteAssetVolume,tradesNumber,takerBuyBaseAssetVolume,takerBuyQuoteAssetVolume,unused
1609459200000,30000,31000,29500,30500,100.5,1609462800000,3065250,1000,60.3,1839150,
1609462800000,30500,32000,30000,31500,150.75,1609466400000,4748625,1500,90.45,2848725,`),
			expected: []SimpleKline{
				{1609459200000, 30000, 31000, 29500, 30500, 100.5, 1609462800000, 3065250, 1000, 60.3, 1839150},
				{1609462800000, 30500, 32000, 30000, 31500, 150.75, 1609466400000, 4748625, 1500, 90.45, 2848725},
			},
			wantErr: false,
		},
		{
			name: "Valid CSV data without header",
			csvData: []byte(`1609459200000,30000,31000,29500,30500,100.5,1609462800000,3065250,1000,60.3,1839150,
1609462800000,30500,32000,30000,31500,150.75,1609466400000,4748625,1500,90.45,2848725,`),
			expected: []SimpleKline{
				{1609459200000, 30000, 31000, 29500, 30500, 100.5, 1609462800000, 3065250, 1000, 60.3, 1839150},
				{1609462800000, 30500, 32000, 30000, 31500, 150.75, 1609466400000, 4748625, 1500, 90.45, 2848725},
			},
			wantErr: false,
		},
		{
			name:     "Empty CSV data",
			csvData:  []byte(``),
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Invalid CSV data (missing column)",
			csvData: []byte(`1609459200000,30000,31000,29500,30500,100.5,1609462800000,3065250,1000,60.3,
1609462800000,30500,32000,30000,31500,150.75,1609466400000,4748625,1500,90.45,2848725,`),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CSVDataToSimpleKlines(tt.csvData)
			if (err != nil) != tt.wantErr {
				t.Errorf("CSVDataToSimpleKlines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("CSVDataToSimpleKlines() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSimpleKlinesToCSVData(t *testing.T) {
	tests := []struct {
		name     string
		klines   []SimpleKline
		expected string
	}{
		{
			name: "Valid SimpleKlines",
			klines: []SimpleKline{
				{1609459200000, 30000, 31000, 29500, 30500, 100.5, 1609462800000, 3065250, 1000, 60.3, 1839150},
				{1609462800000, 30500, 32000, 30000, 31500, 150.75, 1609466400000, 4748625, 1500, 90.45, 2848725},
			},
			expected: "1609459200000,30000,31000,29500,30500,100.5,1609462800000,3065250,1000,60.3,1839150\n1609462800000,30500,32000,30000,31500,150.75,1609466400000,4748625,1500,90.45,2848725",
		},
		{
			name:     "Empty SimpleKlines",
			klines:   []SimpleKline{},
			expected: "",
		},
		{
			name: "Single SimpleKline",
			klines: []SimpleKline{
				{1609459200000, 30000, 31000, 29500, 30500, 100.5, 1609462800000, 3065250, 1000, 60.3, 1839150},
			},
			expected: "1609459200000,30000,31000,29500,30500,100.5,1609462800000,3065250,1000,60.3,1839150",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleKlinesToCSVData(tt.klines)
			if string(got) != tt.expected {
				t.Errorf("SimpleKlinesToCSVData() = %v, want %v", string(got), tt.expected)
			}
		})
	}
}

func TestSimpleKlinesToCSVDataAndBack(t *testing.T) {
	tests := []struct {
		name   string
		klines []SimpleKline
	}{
		{
			name: "Multiple klines",
			klines: []SimpleKline{
				{1609459200000, 30000, 31000, 29500, 30500, 100.5, 1609462800000, 3065250, 1000, 60.3, 1839150},
				{1609462800000, 30500, 32000, 30000, 31500, 150.75, 1609466400000, 4748625, 1500, 90.45, 2848725},
			},
		},
		{
			name:   "Empty klines",
			klines: []SimpleKline{},
		},
		{
			name: "Single kline",
			klines: []SimpleKline{
				{1609459200000, 30000, 31000, 29500, 30500, 100.5, 1609462800000, 3065250, 1000, 60.3, 1839150},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert SimpleKlines to CSV data
			csvData := SimpleKlinesToCSVData(tt.klines)

			// Convert CSV data back to SimpleKlines
			gotKlines, err := CSVDataToSimpleKlines(csvData)
			if err != nil {
				t.Errorf("CSVDataToSimpleKlines() error = %v", err)
				return
			}

			// Compare original klines with the ones converted back from CSV
			if !reflect.DeepEqual(gotKlines, tt.klines) {
				t.Errorf("Conversion mismatch. Got %v, want %v", gotKlines, tt.klines)
			}
		})
	}
}
