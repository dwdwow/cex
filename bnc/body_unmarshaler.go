package bnc

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dwdwow/cex"
)

func bodyUnmshWrapper[D any](unmarshaler cex.RespBodyUnmarshaler[D]) cex.RespBodyUnmarshaler[D] {
	return func(body []byte) (D, *cex.RespBodyUnmarshalerError) {
		err := bodyUnmshCodeMsg(body)
		if err != nil {
			return *new(D), err
		}
		return unmarshaler(body)
	}
}

func bodyUnmshCodeMsg(body []byte) *cex.RespBodyUnmarshalerError {
	codeMsg := CodeMsg{}

	_ = json.Unmarshal(body, &codeMsg)

	code := codeMsg.Code
	msg := codeMsg.Msg

	// https://binance-docs.github.io/apidocs/futures/en/#general-api-information
	// Binance general info doc description:
	// If there is an error message "Unknown error, please check your request or try again later."
	// returned in the response, the API successfully sent the request but not get a response
	// within the timeout period. It is important to NOT treat this as a failure operation;

	// https://binance-docs.github.io/apidocs/futures/en/#error-codes
	// Binance error codes description:
	// -1000 UNKNOWN
	// An unknown error occured while processing the request.
	// if code == -1000 || msg == "Unknown error, please check your request or try again later." {
	// 	return nil
	// }

	// spot: 0, code: 0, 200
	if code == 0 || code == 200 {
		return nil
	}

	if code > 0 {
		// should not get here
		return &cex.RespBodyUnmarshalerError{
			CexErrCode: code,
			CexErrMsg:  msg,
			Err: fmt.Errorf(
				"bnc: %w: code: %v, msg: %v",
				cex.ErrUnexpected, code, msg,
			),
		}
	}

	errCtm := spotCexCustomErrCodes[code]
	switch errCtm {
	case ErrFutureNoNeedToChangePositionSide:
		return nil
	default:
	}
	if errCtm == nil {
		errCtm = fmt.Errorf("%v, %v", code, msg)
	}

	return &cex.RespBodyUnmarshalerError{
		CexErrCode: code,
		CexErrMsg:  msg,
		Err:        fmt.Errorf("bnc: %v", errCtm),
	}
}

func spotOrderReplaceUnmarshaler(body []byte) (SpotReplaceOrderResult, *cex.RespBodyUnmarshalerError) {
	errCodeMsgCheck := bodyUnmshCodeMsg(body)
	if errCodeMsgCheck == nil {
		return spotSucceedOrderReplaceUnmarshaler(body)
	}
	return spotFailedOrderReplaceUnmarshaler(body)
}

func spotSucceedOrderReplaceUnmarshaler(body []byte) (SpotReplaceOrderResult, *cex.RespBodyUnmarshalerError) {
	rawResult := new(SpotReplaceOrderRawData)
	result := SpotReplaceOrderResult{}
	err := json.Unmarshal(body, rawResult)
	if err != nil {
		return result, &cex.RespBodyUnmarshalerError{
			CexErrCode: 0,
			CexErrMsg:  "",
			Err:        fmt.Errorf("%w: %w", cex.ErrJsonUnmarshal, err),
		}
	}
	result.OK = true
	result.OrderCancel = rawResult.CancelResponse
	result.OrderNew = rawResult.NewOrderResponse
	return result, nil
}

func spotFailedOrderReplaceUnmarshaler(body []byte) (SpotReplaceOrderResult, *cex.RespBodyUnmarshalerError) {
	result := SpotReplaceOrderResult{}

	rawResult := new(SpotReplaceOrderRawResult)
	unmshErr := json.Unmarshal(body, rawResult)
	if unmshErr != nil {
		return result, &cex.RespBodyUnmarshalerError{
			CexErrCode: 0,
			CexErrMsg:  "",
			Err:        fmt.Errorf("%w: %w", cex.ErrJsonUnmarshal, unmshErr),
		}
	}

	rawData := rawResult.Data

	cancelResult := rawData.CancelResult
	newResult := rawData.NewOrderResult

	rawCancelResp := rawData.CancelResponse
	rawNewResp := rawData.NewOrderResponse

	result.OrderCancel = rawCancelResp
	result.OrderNew = rawNewResp

	// checking cancel result and new result is unnecessary
	// just in case
	rawResCode := rawResult.Code
	if rawResCode == 0 &&
		cancelResult == SpotOrderCancelNewStatus_SUCCESS &&
		newResult == SpotOrderCancelNewStatus_SUCCESS {
		result.OK = true
		return result, nil
	}

	rawResMsg := rawResult.Msg
	rawErr := SpotCodeMsgChecker(rawResCode)
	if rawErr == nil {
		rawErr = errors.New(rawResMsg)
	}

	if cancelResult == SpotOrderCancelNewStatus_NOT_ATTEMPTED {
		result.ErrCancel = ErrSpotOrderNotAttempted
	} else if cancelResult == SpotOrderCancelNewStatus_FAILURE {
		code := rawCancelResp.Code
		err := SpotCodeMsgChecker(code)
		if err == nil {
			err = errors.New(rawCancelResp.Msg)
		}
		result.ErrCancel = err
	}

	if newResult == SpotOrderCancelNewStatus_NOT_ATTEMPTED {
		result.ErrNew = ErrSpotOrderNotAttempted
	} else if newResult == SpotOrderCancelNewStatus_FAILURE {
		code := rawNewResp.Code
		err := SpotCodeMsgChecker(code)
		if err == nil {
			err = errors.New(rawNewResp.Msg)
		}
		result.ErrNew = err
	}

	return result, &cex.RespBodyUnmarshalerError{
		CexErrCode: rawResCode,
		CexErrMsg:  rawResMsg,
		Err:        rawErr,
	}
}
