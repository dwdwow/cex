package cex

import "errors"

var (
	ErrHttpUnknown  = errors.New("cex std err: http unknown error")
	ErrBadRequest   = errors.New("cex std err: http bad request")
	ErrForbidden    = errors.New("cex std err: http forbidden")
	ErrNotFound     = errors.New("cex std err: http not found")
	ErrTooFrequency = errors.New("cex std err: http too frequency")
	ErrIpBanned     = errors.New("cex std err: http ip is banned")

	ErrInvalidTimestamp    = errors.New("cex stg err: invalid timestamp")
	ErrInsufficientBalance = errors.New("cex stg err: insufficient balance")
	ErrOrderRejected       = errors.New("cex stg err: order is rejected")
)
