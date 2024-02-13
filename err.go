package cex

import "errors"

// These std errors should be wrapped by other errors,
// which can help callers to analyse error details.
var (
	ErrUnexpected = errors.New("unexpected")

	ErrJsonMarshal   = errors.New("json marshal err")
	ErrJsonUnmarshal = errors.New("json unmarshal err")
	ErrS2M           = errors.New("s2m switch err")

	// ErrHTTPCexInnerUnknownStatus
	// Cex may occur inner error, which means user do not know the request status.
	// Under this situation, it is important not to retry immediately.
	// User should check the request status again after seconds.
	ErrHTTPCexInnerUnknownStatus = errors.New("http cex inner unknown status")

	ErrHTTPCodeNotInEnum = errors.New("http code is not in enum")
	ErrHTTPBadRequest    = errors.New("http bad request")
	ErrHTTPForbidden     = errors.New("http forbidden")
	ErrHTTPNotFound      = errors.New("http not found")
	ErrHTTPTooFrequency  = errors.New("http too frequency")
	ErrHTTPIpBanned      = errors.New("http ip is banned")

	ErrInvalidTimestamp    = errors.New("invalid timestamp")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrOrderRejected       = errors.New("order is rejected")
	ErrUnknownOrder        = errors.New("unknown order")
)
