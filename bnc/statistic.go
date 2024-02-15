package bnc

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

type FrTsStatisticalResult struct {
	FrInfo    FuturesFundingRateInfo
	StartTime int64
	EndTime   int64
	Now       float64
	Mean      float64
	Std       float64
}

// CountFundingRateTimeSeries
// within 3 years
func CountFundingRateTimeSeries(frInfo FuturesFundingRateInfo, start, end int64) (FrTsStatisticalResult, error) {
	hours := frInfo.FundingIntervalHours
	if hours <= 0 {
		return FrTsStatisticalResult{}, fmt.Errorf("%v funding rate hours %v <= 0", frInfo.Symbol, hours)
	}
	hiss, err := QueryFundingRateHistories(frInfo.Symbol, start, end, 1000)
	if err != nil {
		return FrTsStatisticalResult{}, err
	}
	var frs []float64
	for _, his := range hiss {
		frs = append(frs, his.FundingRate)
	}
	mean, err := stats.Mean(frs)
	if err != nil {
		return FrTsStatisticalResult{}, err
	}
	std, err := stats.StandardDeviation(frs)
	if err != nil {
		return FrTsStatisticalResult{}, err
	}
	return FrTsStatisticalResult{frInfo, start, end, 0, mean, std}, nil
}
