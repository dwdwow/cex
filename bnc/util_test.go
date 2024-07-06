package bnc

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dwdwow/props"
)

func TestGetWsEvent(t *testing.T) {
	fmt.Println(getWsEvent([]byte("{\"e\":\"balanceUpdate\",\"E\":1720250097625,\"a\":\"PEPE\",\"d\":\"-100000.00\",\"T\":1720250097624}")))
}

func TestUnmarshal(t *testing.T) {
	type A struct {
		H string
	}
	a := A{"hhhh"}
	data, err := json.Marshal(a)
	props.PanicIfNotNil(err)
	a, err = unmarshal[A](data)
	props.PanicIfNotNil(err)
	props.PrintlnIndent(a)
}
