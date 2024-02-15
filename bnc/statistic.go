package bnc

import (
	"fmt"
	"time"

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
// within 8000 or 4000 hours
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
		yr := 24 / hours * 365 * his.FundingRate
		frs = append(frs, yr)
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

// CountSomeFundingRateTimeSeries takes long time, need log process.
func CountSomeFundingRateTimeSeries(infos []FuturesFundingRateInfo, start, end int64) (result []FrTsStatisticalResult, uncounted []FuturesFundingRateInfo) {
	for _, info := range infos {
		fmt.Println("counting", info.Symbol)
		var res FrTsStatisticalResult
		var errCount error
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * time.Duration(i+1))
			res, errCount = CountFundingRateTimeSeries(info, start, end)
			if errCount == nil {
				result = append(result, res)
				fmt.Println("counted", info.Symbol)
				break
			}
			fmt.Println("err", errCount)
		}
		if errCount != nil {
			fmt.Println("can not count", info.Symbol)
			uncounted = append(uncounted, info)
		}
	}
	return
}

// CountAllFundingRateTimeSeries takes long time, need log process.
func CountAllFundingRateTimeSeries(start, end int64) (result []FrTsStatisticalResult, uncounted []FuturesFundingRateInfo, err error) {
	fmt.Println("start to count all funding rate time series")
	infos, err := QueryAllFundingRateInfos()
	if err != nil {
		return
	}
	fmt.Println("exchange amount: ", len(infos))
	result, uncounted = CountSomeFundingRateTimeSeries(infos, start, end)
	if len(uncounted) != 0 {
		result, uncounted = CountSomeFundingRateTimeSeries(uncounted, start, end)
	}
	return
}
