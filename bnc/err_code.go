package bnc

import (
	"errors"
	"fmt"
	"net/http"
)

// HttpErrCodes
// HTTP 5XX return codes are used for internal errors;
// the issue is on Binance's side.
// It is important to NOT treat this as a failure operation;
// the execution status is UNKNOWN and could have been a success.
var HttpErrCodes = map[int]string{
	http.StatusBadRequest:      "Bad Request",
	http.StatusNotFound:        "Not Found",
	http.StatusForbidden:       "Web Application Firewall Has Bin Violated",
	409:                        "CancelReplace Order Partially Succeeds",
	http.StatusTooManyRequests: "Too Many Requests",
	http.StatusTeapot:          "IP Banned",
}

func HttpStatusCodeChecker(code int) error {
	if code == 200 || code >= 500 {
		return nil
	}
	errMsg, ok := HttpErrCodes[code]
	if ok {
		return errors.New(errMsg)
	}
	return fmt.Errorf("bnc: http status code %v is unknown", code)
}

func CustomRespCodeChecker(code int) error {
	return nil
}
