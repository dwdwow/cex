package bnc

import (
	"github.com/dwdwow/cex"
	"github.com/dwdwow/s2m"
	"net/url"
	"strconv"
	"time"
)

type User struct {
	api cex.Api
}

func NewUser(apiKey, secretKey string) User {
	return User{
		api: cex.Api{
			Cex:       cex.BINANCE,
			ApiKey:    apiKey,
			SecretKey: secretKey,
		},
	}
}

func (u User) sign(data any) (query string, err error) {
	return signReqData(data, u.api.SecretKey)
}

func signReqData(data any, key string) (query string, err error) {
	m, err := s2m.ToStrMapWithErr(data)
	if err != nil {
		return
	}
	val := url.Values{
		"timestamp": []string{strconv.FormatInt(time.Now().UnixMilli(), 10)},
	}
	for k, v := range m {
		val.Set(k, v)
	}
	query = val.Encode()
	sig := cex.SignByHmacSHA512ToHex(query, key)
	query += "&signature=" + sig
	return
}
