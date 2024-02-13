package bnc

import "testing"

func TestFuChangePositionMode(t *testing.T) {
	testConfig(ChangePositionModConfig, ChangePositionModParams{DualSidePosition: SmallTrue})
}
