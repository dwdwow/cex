package bnctest

import (
	"encoding/json"
	"fmt"
	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/bnc"
	"github.com/dwdwow/cex/test/cextest"
	"testing"
)

func readApiKey() cex.Api {
	apiKeys := cextest.MustReadApiKey()
	apiKey, ok := apiKeys[cex.BINANCE]
	if !ok {
		panic("no binance api key")
	}
	return apiKey
}

func TestCoinInfo(t *testing.T) {
	apiKey := readApiKey()
	user := bnc.NewUser(apiKey.ApiKey, apiKey.SecretKey)
	_, respData, err := cex.Request(user, bnc.CoinInfoConfig, nil)
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(respData, "", "  ")
	fmt.Println(string(data))
}

func TestSpotAccount(t *testing.T) {
	apiKey := readApiKey()
	user := bnc.NewUser(apiKey.ApiKey, apiKey.SecretKey)
	_, respData, err := cex.Request(user, bnc.SpotAccountConfig, nil)
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(respData, "", "  ")
	fmt.Println(string(data))
}

func TestUniversalTransfer(t *testing.T) {
	apiKey := readApiKey()
	user := bnc.NewUser(apiKey.ApiKey, apiKey.SecretKey)
	_, respData, err := cex.Request(user, bnc.UniversalTransferConfig, bnc.UniversalTransferReq{
		Type:       bnc.TranType_MAIN_UMFUTURE,
		Asset:      "USDT",
		Amount:     10,
		FromSymbol: "",
		ToSymbol:   "",
	})
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(respData, "", "  ")
	fmt.Println(string(data))
}

func TestFlexibleProduct(t *testing.T) {
	apiKey := readApiKey()
	user := bnc.NewUser(apiKey.ApiKey, apiKey.SecretKey)
	_, respData, err := cex.Request(user, bnc.FlexibleProductConfig, bnc.FlexibleProductListReq{
		Asset: "BTC",
	})
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(respData, "", "  ")
	fmt.Println(string(data))
}
