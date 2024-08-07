package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type ListenKeyResponse struct {
	ListenKey string `json:"listenKey"`
}

type ListenKeyParams struct {
	ListenKey string `s2m:"listenKey"`
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
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[ListenKeyResponse]),
}

var KeepListenKeyConfig = cex.ReqConfig[ListenKeyParams, string]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          "",
		Path:             "",
		Method:           http.MethodPut,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[string]),
}

var DeleteListenKeyConfig = cex.ReqConfig[ListenKeyParams, string]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          "",
		Path:             "",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[string]),
}
