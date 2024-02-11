package bnc

import (
	"errors"
	"github.com/dwdwow/cex"
	"net/http"
)

// httpErrCodes
// HTTP 5XX return codes are used for internal errors;
// the issue is on Binance's side.
// It is important to NOT treat this as a failure operation;
// the execution status is UNKNOWN and could have been a success.
var httpErrCodes = map[int]error{
	http.StatusForbidden:       cex.ErrForbidden,
	http.StatusBadRequest:      cex.ErrBadRequest,
	http.StatusNotFound:        cex.ErrNotFound,
	http.StatusTooManyRequests: cex.ErrTooFrequency,
	http.StatusTeapot:          cex.ErrIpBanned,

	409: errors.New("cancelReplace order partially succeeds"),
}

func HttpStatusCodeChecker(code int) error {
	if code == 200 || code >= 500 {
		return nil
	}
	err := httpErrCodes[code]
	if err != nil {
		return err
	}
	return cex.ErrHttpUnknown
}

func CustomRespCodeChecker(code int) error {
	return nil
}
