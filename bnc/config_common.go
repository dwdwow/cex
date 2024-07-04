package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type ListenKeyResponse struct {
	ListenKey string `json:"listenKey"`
}

var NewListenKeyConfig = cex.ReqConfig[cex.NilReqData, ListenKeyResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          "",
		Path:             "",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   cex.StdBodyUnmarshaler[ListenKeyResponse],
}

var KeepListenKeyConfig = cex.ReqConfig[cex.NilReqData, ListenKeyResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          "",
		Path:             "",
		Method:           http.MethodPut,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   nil,
}

var DeleteListenKeyConfig = cex.ReqConfig[cex.NilReqData, ListenKeyResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          "",
		Path:             "",
		Method:           http.MethodPut,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   nil,
}
