package bnc

import (
	"encoding/json"
	"fmt"

	"github.com/dwdwow/cex"
)

func fuBodyUnmshWrapper[D any](unmarshaler cex.RespBodyUnmarshaler[D]) cex.RespBodyUnmarshaler[D] {
	return func(body []byte) (D, *cex.RespBodyUnmarshalerError) {
		err := fuBodyUnmshCodeMsg(body)
		if err != nil {
			return *new(D), err
		}
		return unmarshaler(body)
	}
}

func fuBodyUnmshCodeMsg(body []byte) *cex.RespBodyUnmarshalerError {
	codeMsg := CodeMsg{}

	_ = json.Unmarshal(body, &codeMsg)

	code := codeMsg.Code
	msg := codeMsg.Msg

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
	//switch errCtm {
	//case ErrFutureNoNeedToChangePositionSide:
	//	return nil
	//default:
	//}
	if errCtm == nil {
		errCtm = fmt.Errorf("%v, %v", code, msg)
	}

	return &cex.RespBodyUnmarshalerError{
		CexErrCode: code,
		CexErrMsg:  msg,
		Err:        fmt.Errorf("bnc: %v", errCtm),
	}
}
