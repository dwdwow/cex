package bnc

import (
	"fmt"
	"strconv"
	"time"

	"github.com/montanaflynn/stats"
	"github.com/xuri/excelize/v2"
)

type FrTsStatisticalResult struct {
	FrInfo    FuturesFundingRateInfo
	StartTime int64
	EndTime   int64
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
	return FrTsStatisticalResult{frInfo, start, end, mean, std}, nil
}

// CountSomeFundingRateTimeSeries takes long time, need log process.
func CountSomeFundingRateTimeSeries(infos []FuturesFundingRateInfo, start, end int64) (result []FrTsStatisticalResult, uncounted []FuturesFundingRateInfo) {
	fmt.Println("start to count all funding rate time series")
	infosLen := len(infos)
	fmt.Println("exchanges amount: ", infosLen)
	for j, info := range infos {
		syb := info.Symbol
		num := fmt.Sprintf("%v/%v/%v", len(result), j+1, infosLen)
		fmt.Println(num, "counting", syb)
		var res FrTsStatisticalResult
		var errCount error
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * time.Duration(i+1))
			res, errCount = CountFundingRateTimeSeries(info, start, end)
			if errCount == nil {
				result = append(result, res)
				fmt.Println(num, "counted", syb)
				break
			}
			fmt.Println("err", errCount)
		}
		if errCount != nil {
			fmt.Println(num, "can not count", syb)
			uncounted = append(uncounted, info)
		}
	}
	fmt.Println(len(result), "are counted")
	return
}

// CountAllFundingRateTimeSeries takes long time, will log process.
func CountAllFundingRateTimeSeries(start, end int64) (result []FrTsStatisticalResult, uncounted []FuturesFundingRateInfo, err error) {
	infos, err := QueryAllFundingRateInfos()
	if err != nil {
		return
	}
	result, uncounted = CountSomeFundingRateTimeSeries(infos, start, end)
	if len(uncounted) != 0 {
		var newRes []FrTsStatisticalResult
		newRes, uncounted = CountSomeFundingRateTimeSeries(uncounted, start, end)
		result = append(result, newRes...)
	}
	return
}

func CountAllBncFundingRateTimeSeriesAndSaveToExcel(days int64, path string) error {
	_crtFrs, err := QueryFundingRates()
	if err != nil {
		return err
	}

	crtFrs := map[string]FuturesFundingRate{}

	for _, fr := range _crtFrs {
		crtFrs[fr.Symbol] = fr
	}

	_frInfos, err := QueryAllFundingRateInfos()
	if err != nil {
		return nil
	}

	frInfos := map[string]FuturesFundingRateInfo{}

	for _, info := range _frInfos {
		frInfos[info.Symbol] = info
	}

	spotExchangeInfo, err := QuerySpotExchangeInfo()
	if err != nil {
		return err
	}

	spotExInfos := map[string]Exchange{}

	for _, info := range spotExchangeInfo.Symbols {
		spotExInfos[info.Symbol] = info
	}

	interval := days * 24 * 60 * 60 * 1000
	now := time.Now().UnixMilli()
	start := now - interval
	end := now
	results, uncounted, err := CountAllFundingRateTimeSeries(start, end)
	if err != nil {
		return err
	}

	for _, unc := range uncounted {
		results = append(results, FrTsStatisticalResult{FrInfo: unc, StartTime: start, EndTime: end})
	}

	return writeFrStatisticalResultsToExcel(path, crtFrs, frInfos, spotExInfos, results)
}

func writeFrStatisticalResultsToExcel(path string, crtFrs map[string]FuturesFundingRate, frInfos map[string]FuturesFundingRateInfo, spotExes map[string]Exchange, results []FrTsStatisticalResult) error {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	const sheet1Name = "Sheet1"

	err := f.SetSheetRow(sheet1Name, "A1", &[]string{"symbol", "now", "mean", "std", "has spot"})
	if err != nil {
		return err
	}

	for i, res := range results {
		row := strconv.FormatInt(int64(i+2), 10)
		syb := res.FrInfo.Symbol
		crtFr := crtFrs[syb]
		frInfo, ok := frInfos[syb]
		hours := frInfo.FundingIntervalHours
		if !ok || hours <= 0 {
			fmt.Printf("%v has no fr info, %+v", syb, res)
			continue
		}
		_, hasSpot := spotExes[syb]
		yr := 24 / hours * 365 * crtFr.LastFundingRate
		err := f.SetSheetRow(sheet1Name, "A"+row, &[]any{syb, yr, res.Mean, res.Std, hasSpot})
		if err != nil {
			return err
		}
	}

	return f.SaveAs(fmt.Sprintf("%v/cex_fr_statistics_%v.xlsx", path, time.Now().UnixMilli()))
}
