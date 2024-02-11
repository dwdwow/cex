package cex

import "errors"

var (
	ErrTooFrequency    = errors.New("cex: too frequency")
	ErrIpBanned        = errors.New("cex: ip is banned")
	ErrOutOfRecvWindow = errors.New("cex: out of recv window")

	ErrInsufficientBalance = errors.New("cex: insufficient balance")
	ErrOrderRejected       = errors.New("cex: order is rejected")
)
