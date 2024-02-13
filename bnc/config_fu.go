package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type ChangePositionModParams struct {
	DualSidePosition SmallBool `json:"dualSidePosition"`
}

var FuChangePositionModeConfig = cex.ReqConfig[ChangePositionModParams, CodeMsg]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/positionSide/dual",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[CodeMsg]),
}

type CurrentPositionModeResult struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

var FuPositionModeConfig = cex.ReqConfig[cex.NilReqData, CurrentPositionModeResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          FapiBaseUrl,
		Path:             FapiV1 + "/positionSide/dual",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[CurrentPositionModeResult]),
}
