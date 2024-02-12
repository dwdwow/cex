package bnc

import (
	"errors"
	"net/http"

	"github.com/dwdwow/cex"
)

var (
	ErrPartiallySucceeds = errors.New("cancelReplace order partially succeeds")
)

// httpErrCodes
// HTTP 5XX return codes are used for internal errors;
// the issue is on Binance's side.
// It is important to NOT treat this as a failure operation;
// the execution status is UNKNOWN and could have been a success.
var httpErrCodes = map[int]error{
	http.StatusForbidden:       cex.ErrHttpForbidden,
	http.StatusBadRequest:      cex.ErrHttpBadRequest,
	http.StatusNotFound:        cex.ErrHttpNotFound,
	http.StatusTooManyRequests: cex.ErrHttpTooFrequency,
	http.StatusTeapot:          cex.ErrHttpIpBanned,

	409: ErrPartiallySucceeds,
}

func HTTPStatusCodeChecker(code int) error {
	// If status code >= 500, status is unknown.
	// Binance document indicate that user can ignore.
	// https://binance-docs.github.io/apidocs/spot/en/#general-api-information
	if code == 200 || code >= 500 {
		return nil
	}
	err := httpErrCodes[code]
	if err != nil {
		return err
	}
	return cex.ErrHttpUnknown
}

var cexCustomErrCodes = map[int]error{
	-1021: cex.ErrInvalidTimestamp,
}
