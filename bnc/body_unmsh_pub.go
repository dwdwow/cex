package bnc

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/dwdwow/cex"
)

func obBodyUnmsher(body []byte) (OrderBook, *cex.RespBodyUnmarshalerError) {
	raw := new(RawOrderBook)
	err := json.Unmarshal(body, raw)
	if err != nil {
		return OrderBook{}, &cex.RespBodyUnmarshalerError{Err: fmt.Errorf("%w: %w", cex.ErrJsonUnmarshal, err)}
	}

	code := raw.Code
	if code != 0 {
		errMsg := SpotCodeMsgChecker(code)
		if errMsg == nil {
			errMsg = errors.New(raw.Msg)
		}
		return OrderBook{}, &cex.RespBodyUnmarshalerError{
			CexErrCode: code,
			CexErrMsg:  raw.Msg,
			Err:        errMsg,
		}
	}

	bids, err := convRawStrBookToFloatBook(raw.Bids)
	if err != nil {
		return OrderBook{}, &cex.RespBodyUnmarshalerError{Err: fmt.Errorf("parse raw orderbook bids, %w", err)}
	}
	asks, err := convRawStrBookToFloatBook(raw.Asks)
	if err != nil {
		return OrderBook{}, &cex.RespBodyUnmarshalerError{Err: fmt.Errorf("parse raw orderbook asks, %w", err)}
	}
	return OrderBook{Bids: bids, Asks: asks, LastUpdateId: raw.LastUpdateId}, nil
}

func convRawStrBookToFloatBook(raw [][]string) ([][]float64, error) {
	var book [][]float64
	for _, pq := range raw {
		if len(pq) != 2 {
			return nil, fmt.Errorf("price and qty in book %v len != 2", pq)
		}
		p, err := strconv.ParseFloat(pq[0], 64)
		if err != nil {
			return nil, fmt.Errorf("parse price %v, %w", pq[0], err)
		}
		q, err := strconv.ParseFloat(pq[1], 64)
		if err != nil {
			return nil, fmt.Errorf("parse qty %v, %w", pq[1], err)
		}
		book = append(book, []float64{p, q})
	}
	return book, nil
}
