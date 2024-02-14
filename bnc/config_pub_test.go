package bnc

import (
	"testing"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/props"
)

func testPubConfig[ReqDataType, RespDataType any](
	config cex.ReqConfig[ReqDataType, RespDataType],
	reqData ReqDataType,
	opts ...cex.ReqOpt,
) {
	resp, ob, err := cex.Request(EmptyUser(), config, reqData, opts...)
	_ = resp
	props.PanicIfNotNil(err)
	props.PrintlnIndent(ob)
}

func TestOrderBook(t *testing.T) {
	testPubConfig(OrderBookConfig, OrderBookParams{
		Symbol: "ETHUSDT",
		Limit:  0,
	})
}
