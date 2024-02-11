package bnc

import (
	"encoding/json"
	"fmt"
	"github.com/dwdwow/cex"
)

func GeneralBodyUnmarshaler[D any](body []byte) (D, *cex.RespBodyUnmarshalerError) {
	d := new(D)
	codeMsg := CodeMsg{}

	if err := json.Unmarshal(body, &codeMsg); err != nil {
		return cex.StdRespDataUnmarshaler[D](body)
	}

	code := codeMsg.Code
	msg := codeMsg.Msg

	if code >= 0 {
		// should not get here
		return *d, &cex.RespBodyUnmarshalerError{
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

	return *d, &cex.RespBodyUnmarshalerError{
		CexErrCode: code,
		CexErrMsg:  msg,
		Err:        fmt.Errorf("bnc: %v", errCtm),
	}
}
