package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type PortfolioMarginAccountAsset struct {
	Asset                  string  `json:"asset"`
	CrossWalletBalance     float64 `json:"crossWalletBalance,string"`
	CrossUnPnl             float64 `json:"crossUnPnl,string"`
	MaintMargin            float64 `json:"maintMargin,string"`
	InitialMargin          float64 `json:"initialMargin,string"`
	PositionInitialMargin  float64 `json:"positionInitialMargin,string"`
	OpenOrderInitialMargin float64 `json:"openOrderInitialMargin,string"`
	UpdateTime             int64   `json:"updateTime"`
}

type PortfolioMarginAccountPosition struct {
	Symbol                 string              `json:"symbol"`
	InitialMargin          float64             `json:"initialMargin,string"`
	MaintMargin            float64             `json:"maintMargin,string"`
	UnrealizedProfit       float64             `json:"unrealizedProfit,string"`
	PositionInitialMargin  float64             `json:"positionInitialMargin,string"`
	OpenOrderInitialMargin float64             `json:"openOrderInitialMargin,string"`
	Leverage               float64             `json:"leverage,string"`
	EntryPrice             float64             `json:"entryPrice,string"`
	MaxNotional            float64             `json:"maxNotional,string"`
	BidNotional            float64             `json:"bidNotional,string"`
	AskNotional            float64             `json:"askNotional,string"`
	PositionSide           FuturesPositionSide `json:"positionSide,string"`
	PositionAmt            float64             `json:"positionAmt,string"`
	BreakEvenPrice         float64             `json:"breakEvenPrice,string"`
	UpdateTime             int                 `json:"updateTime"`
}

type PortfolioMarginAccountDetail struct {
	Assets    []PortfolioMarginAccountAsset    `json:"assets"`
	Positions []PortfolioMarginAccountPosition `json:"positions"`
}

var PortfolioMarginAccountDetailConfig = cex.ReqConfig[cex.NilReqData, PortfolioMarginAccountDetail]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/um/account",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[PortfolioMarginAccountDetail]),
}

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

var PortfolioMarginBalanceConfig = cex.ReqConfig[PortfolioMarginAccountBalanceParams, PortfolioMarginBalance]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/balance",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[PortfolioMarginBalance]),
}

var PortfolioMarginBalancesConfig = cex.ReqConfig[cex.NilReqData, []PortfolioMarginBalance]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/balance",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]PortfolioMarginBalance]),
}

type PortfolioMarginAccountInformation struct {
	UniMMR                   float64                      `json:"uniMMR,string"`
	AccountEquity            float64                      `json:"accountEquity,string"`
	ActualEquity             float64                      `json:"actualEquity,string"`
	AccountInitialMargin     float64                      `json:"accountInitialMargin,string"`
	AccountMaintMargin       float64                      `json:"accountMaintMargin,string"`
	AccountStatus            PortfolioMarginAccountStatus `json:"accountStatus"`
	VirtualMaxWithdrawAmount float64                      `json:"virtualMaxWithdrawAmount,string"`
	TotalAvailableBalance    float64                      `json:"totalAvailableBalance,string"`
	TotalMarginOpenLoss      float64                      `json:"totalMarginOpenLoss,string"`
	UpdateTime               int64                        `json:"updateTime"`
}

var PortfolioMarginAccountInformationConfig = cex.ReqConfig[cex.NilReqData, PortfolioMarginAccountInformation]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/account",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[PortfolioMarginAccountInformation]),
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

var PortfolioMarginPositionsConfig = cex.ReqConfig[FuturesPositionsParams, []FuturesPosition]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/um/positionRisk",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]FuturesPosition]),
}
