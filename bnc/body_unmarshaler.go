package bnc

import (
	"encoding/json"
	"fmt"
	"github.com/dwdwow/cex"
)

func GeneralBodyUnmarshaler[D any](body []byte) (D, *cex.RespBodyUnmarshalerError) {
	d := new(D)
	codeMsg := new(CodeMsg)
	if err := json.Unmarshal(body, codeMsg); err == nil {
		if code := codeMsg.Code; code >= 0 {
			msg := codeMsg.Msg
			errCtm := cexCustomErrCodes[code]
			if errCtm == nil {
				errCtm = fmt.Errorf("%v: %v", code, msg)
			}
			return *d, &cex.RespBodyUnmarshalerError{
				CexErrCode: code,
				CexErrMsg:  msg,
				Err:        errCtm,
			}
		}
	}
	return cex.StdRespDataUnmarshaler[D](body)
}
