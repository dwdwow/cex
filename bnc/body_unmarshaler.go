package bnc

import (
	"encoding/json"
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

	if code == 0 {
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

	errCtm := cexCustomErrCodes[code]
	if errCtm == nil {
		errCtm = fmt.Errorf("%v, %v", code, msg)
	}

	return &cex.RespBodyUnmarshalerError{
		CexErrCode: code,
		CexErrMsg:  msg,
		Err:        fmt.Errorf("bnc: %v", errCtm),
	}
}

func PageBodyUnmarshaler[Slice any](body []byte) (Slice, *cex.RespBodyUnmarshalerError) {
	page := new(Page[Slice])
	err := json.Unmarshal(body, page)
	var serr *cex.RespBodyUnmarshalerError
	if err != nil {
		serr = &cex.RespBodyUnmarshalerError{
			CexErrCode: 0,
			CexErrMsg:  "",
			Err:        fmt.Errorf("%w: %w", cex.ErrJsonUnmarshal, err),
		}
	}
	return page.Rows, serr
}