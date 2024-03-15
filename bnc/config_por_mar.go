package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type PortfolioMarginAccountBalanceParams struct {
	Symbol string `s2m:"symbol,omitempty"`
}

type PortfolioMarginBalance struct {
	Asset               string  `json:"asset"`
	TotalWalletBalance  float64 `json:"totalWalletBalance,string"`
	CrossMarginAsset    float64 `json:"crossMarginAsset,string"`
	CrossMarginBorrowed float64 `json:"crossMarginBorrowed,string"`
	CrossMarginFree     float64 `json:"crossMarginFree,string"`
	CrossMarginInterest float64 `json:"crossMarginInterest,string"`
	CrossMarginLocked   float64 `json:"crossMarginLocked,string"`
	UmWalletBalance     float64 `json:"umWalletBalance,string"`
	UmUnrealizedPNL     float64 `json:"umUnrealizedPNL,string"`
	CmWalletBalance     float64 `json:"cmWalletBalance,string"`
	CmUnrealizedPNL     float64 `json:"cmUnrealizedPNL,string"`
	UpdateTime          int64   `json:"updateTime"`
}

var PortfolioMarginNewOrderConfig = cex.ReqConfig[FuturesNewOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/um/order",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var PortfolioMarginQueryOrderConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/um/order",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var PortfolioMarginCancelOrderConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/um/order",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}
