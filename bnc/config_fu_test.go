package bnc

import "testing"

func TestFuChangePositionMode(t *testing.T) {
	testConfig(FuChangePositionModeConfig, ChangePositionModParams{DualSidePosition: SmallFalse})
}

func TestFuCurrentPositionMode(t *testing.T) {
	// TODO retest
	// return {"code":-1022,"msg":"Signature for this request is not valid."}
	testConfig(FuPositionModeConfig, nil)
}

func TestFuChangeMultiAssetsMode(t *testing.T) {
	testConfig(FuChangeMultiAssetsModeConfig, FuChangeMultiAssetsModeParams{MultiAssetsMargin: SmallFalse})
}
