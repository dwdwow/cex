package bnc

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
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

func NewEmptySimpleKline() SimpleKline {
	return SimpleKline{}
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
	return k[0] == 0
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

func (k SimpleKline) SetOpenTime(value int64) SimpleKline {
	k[0] = float64(value)
	return k
}

func (k SimpleKline) SetCloseTime(value int64) SimpleKline {
	k[6] = float64(value)
	return k
}

func (k SimpleKline) SetTradesNumber(value int64) SimpleKline {
	k[8] = float64(value)
	return k
}

func (k SimpleKline) SetOpenPrice(value float64) SimpleKline {
	k[1] = value
	return k
}

func (k SimpleKline) SetHighPrice(value float64) SimpleKline {
	k[2] = value
	return k
}

func (k SimpleKline) SetLowPrice(value float64) SimpleKline {
	k[3] = value
	return k
}

func (k SimpleKline) SetClosePrice(value float64) SimpleKline {
	k[4] = value
	return k
}

func (k SimpleKline) SetVolume(value float64) SimpleKline {
	k[5] = value
	return k
}

func (k SimpleKline) SetQuoteAssetVolume(value float64) SimpleKline {
	k[7] = value
	return k
}

func (k SimpleKline) SetTakerBuyBaseAssetVolume(value float64) SimpleKline {
	k[9] = value
	return k
}

func (k SimpleKline) SetTakerBuyQuoteAssetVolume(value float64) SimpleKline {
	k[10] = value
	return k
}

// ToCSVRow converts SimpleKline to a CSV row string
func (k SimpleKline) ToCSVRow() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
		// OpenTime: Unix timestamp in milliseconds
		strconv.FormatInt(int64(k.OpenTime()), 10),
		// OpenPrice: Opening price of the interval
		strconv.FormatFloat(k.OpenPrice(), 'f', -1, 64),
		// HighPrice: Highest price during the interval
		strconv.FormatFloat(k.HighPrice(), 'f', -1, 64),
		// LowPrice: Lowest price during the interval
		strconv.FormatFloat(k.LowPrice(), 'f', -1, 64),
		// ClosePrice: Closing price of the interval
		strconv.FormatFloat(k.ClosePrice(), 'f', -1, 64),
		// Volume: Total trading volume during the interval
		strconv.FormatFloat(k.Volume(), 'f', -1, 64),
		// CloseTime: Unix timestamp in milliseconds for interval close
		strconv.FormatInt(int64(k.CloseTime()), 10),
		// QuoteAssetVolume: Total quote asset volume during the interval
		strconv.FormatFloat(k.QuoteAssetVolume(), 'f', -1, 64),
		// TradesNumber: Number of trades during the interval
		strconv.FormatInt(int64(k.TradesNumber()), 10),
		// TakerBuyBaseAssetVolume: Taker buy base asset volume
		strconv.FormatFloat(k.TakerBuyBaseAssetVolume(), 'f', -1, 64),
		// TakerBuyQuoteAssetVolume: Taker buy quote asset volume
		strconv.FormatFloat(k.TakerBuyQuoteAssetVolume(), 'f', -1, 64),
	)
}

// CSVDataToSimpleKlines converts CSV data to SimpleKline
// CSV data should be in the format:
// "openTime,openPrice,highPrice,lowPrice,closePrice,volume,closeTime,quoteAssetVolume,tradesNumber,takerBuyBaseAssetVolume,takerBuyQuoteAssetVolume,unused"
// if the first line is the header, it will be skipped
func CSVDataToSimpleKlines(data []byte) (klines []SimpleKline, err error) {
	if len(data) == 0 {
		return
	}
	buf := bufio.NewReader(bytes.NewReader(data))
	var line int64
	for {
		l, _, err := buf.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		if len(l) == 0 {
			return nil, errors.New("empty line")
		}
		line++
		parts := strings.Split(string(l), ",")
		_, err = strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			if line == 1 {
				continue
			}
			return nil, err
		}
		var raw RawKline
		for i, p := range parts {
			raw[i] = p
		}
		kline, err := NewSimpleKlineFromRaw(raw)
		if err != nil {
			return nil, err
		}
		klines = append(klines, kline)
	}
	return
}

// SimpleKlinesToCSVData converts a slice of SimpleKline to CSV data without headers
func SimpleKlinesToCSVData(klines []SimpleKline) []byte {
	var buf bytes.Buffer
	for _, kline := range klines {
		buf.WriteString(kline.ToCSVRow() + "\n")
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	return buf.Bytes()
}
