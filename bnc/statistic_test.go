package bnc

import (
	"testing"
	"time"

	"github.com/dwdwow/props"
)

func TestCountFundingRateTimeSeries(t *testing.T) {
	info := FuturesFundingRateInfo{
		Symbol:                   "ETHUSDT",
		AdjustedFundingRateCap:   0,
		AdjustedFundingRateFloor: 0,
		FundingIntervalHours:     8,
		Disclaimer:               false,
	}
	days := 60.0
	interval := int64(days * 24 * 60 * 60 * 1000)
	now := time.Now().UnixMilli()
	res, err := CountFundingRateTimeSeries(info, now-interval, now)
	props.PanicIfNotNil(err)
	props.PrintlnIndent(res)
}

func TestCountAllFundingRateTimeSeries(t *testing.T) {
	days := 60.0
	interval := int64(days * 24 * 60 * 60 * 1000)
	now := time.Now().UnixMilli()
	res, uncounted, err := CountAllFundingRateTimeSeries(now-interval, now)
	props.PanicIfNotNil(err)
	props.PrintlnIndent(res)
	props.PrintlnIndent(uncounted)
}
