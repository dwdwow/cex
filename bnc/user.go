package bnc

import "github.com/dwdwow/cex"

type User struct {
	api cex.Api
}

func NewUser(apiKey, secretKey string) User {
	return User{cex.Api{
		Cex:       cex.BINANCE,
		ApiKey:    apiKey,
		SecretKey: secretKey,
	}}
}
