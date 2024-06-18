package bnc

import (
	"math"
	"net/http"

	"github.com/dwdwow/cex"
)

// futures account assets in portfolio margin mode
// use portfolio margin account_balance first ?
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

// futures account positions in portfolio margin mode
// use portfolio margin um_position_risk first ?
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
	PositionSide           FuturesPositionSide `json:"positionSide"`
	SignPositionAmt        float64             `json:"positionAmt,string"`
	BreakEvenPrice         float64             `json:"breakEvenPrice,string"`
	UpdateTime             int64               `json:"updateTime"`

	// just for CM position
	MaxQty string `json:"maxQty"`
}

func (p PortfolioMarginAccountPosition) AbsPositionAmt() float64 {
	return math.Abs(p.SignPositionAmt)
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

var PortfolioMarginAccountCMDetailConfig = cex.ReqConfig[cex.NilReqData, PortfolioMarginAccountDetail]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/cm/account",
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

type PortfolioMarginUMPositionRisk struct {
	Symbol           string  `json:"symbol"`
	PositionAmt      float64 `json:"positionAmt,string"`
	EntryPrice       float64 `json:"entryPrice,string"`
	MarkPrice        float64 `json:"markPrice,string"`
	UnRealizedProfit float64 `json:"unRealizedProfit,string"`
	LiquidationPrice float64 `json:"liquidationPrice,string"`
	Leverage         float64 `json:"leverage,string"`
	PositionSide     string  `json:"positionSide"`
	UpdateTime       int64   `json:"updateTime"`
	MaxNotionalValue float64 `json:"maxNotionalValue,string"`
	Notional         float64 `json:"notional,string"`
	BreakEvenPrice   float64 `json:"breakEvenPrice,string"`
}

type PortfolioMarginCMPositionRisk PortfolioMarginUMPositionRisk

//var PortfolioMarginBalanceConfig = cex.ReqConfig[PortfolioMarginAccountBalanceParams, PortfolioMarginBalance]{
//	ReqBaseConfig: cex.ReqBaseConfig{
//		BaseUrl:          PapiBaseUrl,
//		Path:             PapiV1 + "/balance",
//		Method:           http.MethodGet,
//		IsUserData:       true,
//		UserTimeInterval: 0,
//		IpTimeInterval:   0,
//	},
//	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
//	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[PortfolioMarginBalance]),
//}

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

var PortfolioMarginPositionsConfig = cex.ReqConfig[FuturesPositionsParams, []PortfolioMarginUMPositionRisk]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/um/positionRisk",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]PortfolioMarginUMPositionRisk]),
}

var PortfolioMarginNewCMOrderConfig = cex.ReqConfig[FuturesNewOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/cm/order",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var PortfolioMarginQueryCMOrderConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/cm/order",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var PortfolioMarginCancelCMOrderConfig = cex.ReqConfig[FuturesQueryOrCancelOrderParams, FuturesOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/cm/order",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FuturesOrder]),
}

var PortfolioMarginCMPositionsConfig = cex.ReqConfig[FuturesPositionsParams, []PortfolioMarginCMPositionRisk]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/cm/positionRisk",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]PortfolioMarginCMPositionRisk]),
}

type PortfolioMarginBNBTransferParams struct {
	Amount       float64
	TransferSide PortfolioMarginBNBTransferSide
}

type PortfolioMarginBNBTransferResult struct {
	TranId int64
}

var PortfolioMarginBNBTransferConfig = cex.ReqConfig[PortfolioMarginBNBTransferParams, PortfolioMarginBNBTransferResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          PapiBaseUrl,
		Path:             PapiV1 + "/bnb-transfer",
		Method:           http.MethodPost,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[PortfolioMarginBNBTransferResult]),
}

type PortfolioMarginCollateralRate struct {
	Asset          string  `json:"asset"`
	CollateralRate float64 `json:"collateralRate,string"`
}

var PortfolioMarginCollateralRatesConfig = cex.ReqConfig[cex.NilReqData, FrontData[[]PortfolioMarginCollateralRate]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          "https://www.binance.com",
		Path:             "/bapi/margin/v1/public/margin/portfolio/collateral-rate",
		Method:           http.MethodGet,
		IsUserData:       false,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   fuBodyUnmshWrapper(cex.StdBodyUnmarshaler[FrontData[[]PortfolioMarginCollateralRate]]),
}
