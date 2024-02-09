package bnc

import (
	"github.com/dwdwow/cex"
	"github.com/dwdwow/s2m"
	"github.com/go-resty/resty/v2"
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

// ============================== req handler start ==============================

func (u User) MakeReq(config cex.ReqBaseConfig, reqData any, opts ...cex.ReqOpt) (*resty.Request, error) {
	query, err := u.sign(reqData)
	if err != nil {
		return nil, err
	}
	clt := resty.New().
		SetBaseURL(config.BaseUrl).
		SetHeader("X-MBX-APIKEY", u.api.ApiKey)
	req := clt.R().
		SetQueryString(query)
	for _, opt := range opts {
		opt(clt, req)
	}
	return req, nil
}

func (u User) CheckResp(response *resty.Response, request *resty.Request) error {
	return nil
}

// ============================== req handler end ==============================

// ============================== sign start ==============================

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

// ============================== sign end ==============================
