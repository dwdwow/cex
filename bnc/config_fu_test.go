package bnc

import "testing"

func TestFuChangePositionMode(t *testing.T) {
	testConfig(FuChangePositionModeConfig, ChangePositionModParams{DualSidePosition: SmallFalse})
}

func TestCurrentPositionMode(t *testing.T) {
	testConfig(FuPositionModeConfig, nil)
}
