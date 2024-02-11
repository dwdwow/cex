package cex

import "errors"

var (
	ErrJsonMarshal   = errors.New("json marshal err")
	ErrJsonUnmarshal = errors.New("json unmarshal err")
	ErrS2M           = errors.New("s2m err")

	ErrHttpUnknown      = errors.New("http unknown error")
	ErrHttpBadRequest   = errors.New("http bad request")
	ErrHttpForbidden    = errors.New("http forbidden")
	ErrHttpNotFound     = errors.New("http not found")
	ErrHttpTooFrequency = errors.New("http too frequency")
	ErrHttpIpBanned     = errors.New("http ip is banned")

	ErrInvalidTimestamp    = errors.New("invalid timestamp")
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrOrderRejected       = errors.New("order is rejected")
)
