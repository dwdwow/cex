package bnc

import (
	"errors"
	"net/http"

	"github.com/dwdwow/cex"
)

/**
Binance spot custom codes are conflict to future's.
Future endpoint hase many meanings about HTTP status code 503.

HTTP Ref:
	Spot    : https://binance-docs.github.io/apidocs/spot/en/#general-api-information
	Futures : https://binance-docs.github.io/apidocs/futures/en/#general-api-information

Error Codes Ref:
Spot    : https://binance-docs.github.io/apidocs/spot/en/#error-codes
Futures : https://binance-docs.github.io/apidocs/futures/en/#error-codes

All Endpoints:
	HTTP 4XX return codes are used for malformed requests; the issue is on the sender's side.
	HTTP 403 return code is used when the WAF Limit (Web Application Firewall) has been violated.
	HTTP 429 return code is used when breaking a request rate limit.
	HTTP 418 return code is used when an IP has been auto-banned for continuing to send requests after receiving 429 codes.

Spot Endpoint:
	HTTP 5XX return codes use for internal errors; the issue is on Binance's side.
	It is important to NOT treat this as a failure operation;
	the execution status is UNKNOWN and could have been a success.

Future Endpoint:
	HTTP 5XX return codes are used for internal errors; the issue is on Binance's side.
	If there is an error message "Request occur unknown error.", please retry later.
	HTTP 503 return code is used when:
		1. If there is an error message "Unknown error, please check your request or try again later."
		   returned in the response, the API successfully sent the request but not get a response
		   within the timeout period. It is important to NOT treat this as a failure operation;
		   the execution status is UNKNOWN and could have been a success;
		2. If there is an error message "Service Unavailable." returned in the response,
		   it means this is a failure API operation and the service might be unavailable at the moment,
		   you need to retry later.
		3. If there is an error message "Internal error; unable to process your request.
		   Please try again." returned in the response, it means this is a failure API operation,
		   and you can resend your request if you need.
		4. If there is an error message "Server is currently overloaded with other requests. Please try again in a few minutes."
		   returned in the response, it means this is a failure API operation, and you can resend your request if you need.
*/

// =============================================
// HTTP Errors
// ---------------------------------------------

var (
	ErrPartiallySucceeds = errors.New("cancelReplace order partially succeeds")
)

// httpErrCodes
var httpErrCodes = map[int]error{
	http.StatusForbidden:       cex.ErrHTTPForbidden,
	http.StatusBadRequest:      cex.ErrHTTPBadRequest,
	http.StatusNotFound:        cex.ErrHTTPNotFound,
	http.StatusTooManyRequests: cex.ErrHTTPTooFrequency,
	http.StatusTeapot:          cex.ErrHTTPIpBanned,

	409: ErrPartiallySucceeds,
}

func HTTPStatusCodeChecker(code int) error {
	if code == 200 {
		return nil
	}
	// Spot Endpoint:
	// If status code >= 500, status is unknown.
	// Binance document indicate that user can ignore.
	// ref: https://binance-docs.github.io/apidocs/spot/en/#general-api-information
	//
	// Future Endpoint:
	// While status code is 5XX,
	// if error msg is not "Unknown error, please check your request or try again later.",
	// should retry latter.
	// ref: https://binance-docs.github.io/apidocs/futures/en/#general-api-information
	//
	// Although there are differences between spot and future endpoint, just ignore.
	// Binance spot document may be not thorough? I do not know.
	if code > 499 {
		return cex.ErrHTTPCexInnerUnknownStatus
	}
	err := httpErrCodes[code]
	if err != nil {
		return err
	}
	return cex.ErrHTTPCodeNotInEnum
}

// ---------------------------------------------
// HTTP Errors
// =============================================

// =============================================
// Common Custom Errors
// ---------------------------------------------

var (
	ErrCexInnerProblems = errors.New("an unknown error occured while processing the request")
)

// ---------------------------------------------
// Common Custom Errors
// =============================================

// =============================================
// Binance Spot Custom Errors
// ---------------------------------------------

var (
	ErrSpotOrderCancelReplacePartiallyFailed = errors.New("order cancel-replace partially failed")
	ErrSpotOrderCancelReplaceFailed          = errors.New("order cancel-replace failed")
	ErrSpotOrderWouldImmediatelyMatchAndTake = errors.New("order would immediately match and take")
	ErrSpotOrderNotAttempted                 = errors.New("order is not attempted")
)

var spotCexCustomErrCodes = map[int]error{
	-1000: ErrCexInnerProblems,
	-1021: cex.ErrInvalidTimestamp,
	-2010: ErrSpotOrderWouldImmediatelyMatchAndTake,
	-2011: cex.ErrUnknownOrder,
	-2021: ErrSpotOrderCancelReplacePartiallyFailed,
	-2022: ErrSpotOrderCancelReplaceFailed,
}

func SpotCodeMsgChecker(code int) error {
	return spotCexCustomErrCodes[code]
}

// ---------------------------------------------
// Binance Spot Custom Errors
// =============================================

// =============================================
// Binance Future Custom Errors
// ---------------------------------------------

var (
	ErrFutureNoNeedToChangePositionSide = errors.New("no need to change position side")
)

var fuCexCustomErrCodes = map[int]error{
	-1000: ErrCexInnerProblems,
	-1021: cex.ErrInvalidTimestamp,
	-4059: ErrFutureNoNeedToChangePositionSide,
}

func FutureCodeMsgChecker(code int) error {
	return FutureCodeMsgChecker(code)
}

// ---------------------------------------------
// Binance Future Custom Errors
// =============================================
