package bnc

import (
	"fmt"
	"testing"
)

func TestGetWsEvent(t *testing.T) {
	fmt.Println(getWsEvent([]byte("{\"e\":\"balanceUpdate\",\"E\":1720250097625,\"a\":\"PEPE\",\"d\":\"-100000.00\",\"T\":1720250097624}")))
}
