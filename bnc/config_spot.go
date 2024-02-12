package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type CoinInfo struct {
	Coin             string            `json:"coin"`
	DepositAllEnable bool              `json:"depositAllEnable"`
	Free             float64           `json:"free,string"`
	Freeze           float64           `json:"freeze,string"`
	Ipoable          float64           `json:"ipoable,string"`
	Ipoing           float64           `json:"ipoing,string"`
	IsLegalMoney     bool              `json:"isLegalMoney"`
	Locked           float64           `json:"locked,string"`
	Name             string            `json:"name"`
	NetworkList      []CoinNetworkInfo `json:"networkList"`
}

type CoinNetworkInfo struct {
	AddressRegex            string `json:"addressRegex"`
	Coin                    string `json:"coin"`
	DepositDesc             string `json:"depositDesc"`
	DepositEnable           bool   `json:"depositEnable"`
	IsDefault               bool   `json:"isDefault"`
	MemoRegex               string `json:"memoRegex"`
	MinConfirm              int    `json:"minConfirm"`
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	SpecialTips             string `json:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm"`
	WithdrawDesc            string `json:"withdrawDesc"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress"`
	EstimatedArrivalTime    int    `json:"estimatedArrivalTime"`
	Busy                    bool   `json:"busy"`
}

var CoinInfoConfig = cex.ReqConfig[cex.EmptyReqData, []CoinInfo]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/capital/config/getall",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[[]CoinInfo]),
}

type SpotBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free,string"`
	Locked float64 `json:"locked,string"`
}

type SpotAccount struct {
	MakerCommission  float64 `json:"makerCommission" bson:"makerCommission"`
	TakerCommission  float64 `json:"takerCommission" bson:"takerCommission"`
	BuyerCommission  float64 `json:"buyerCommission" bson:"buyerCommission"`
	SellerCommission float64 `json:"sellerCommission" bson:"sellerCommission"`
	CommissionRates  struct {
		Maker  float64 `json:"maker,string" bson:"maker"`
		Taker  float64 `json:"taker,string" bson:"taker"`
		Buyer  float64 `json:"buyer,string" bson:"buyer"`
		Seller float64 `json:"seller,string" bson:"seller"`
	} `json:"commissionRates" bson:"commissionRates"`
	CanTrade                   bool          `json:"canTrade" bson:"canTrade"`
	CanWithdraw                bool          `json:"canWithdraw" bson:"canWithdraw"`
	CanDeposit                 bool          `json:"canDeposit" bson:"canDeposit"`
	Brokered                   bool          `json:"brokered" bson:"brokered"`
	RequireSelfTradePrevention bool          `json:"requireSelfTradePrevention" bson:"requireSelfTradePrevention"`
	UpdateTime                 int64         `json:"updateTime" bson:"updateTime"`
	AccountType                AcctType      `json:"accountType" bson:"accountType"`
	Balances                   []SpotBalance `json:"balances" bson:"balances"`
	Permissions                []TradeType   `json:"permissions" bson:"permissions"`
}

var SpotAccountConfig = cex.ReqConfig[cex.EmptyReqData, SpotAccount]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/account",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[SpotAccount]),
}

type UniversalTransferParams struct {
	Type       TranType `s2m:"type,omitempty"`
	Asset      string   `s2m:"asset,omitempty"`
	Amount     float64  `s2m:"amount,omitempty"`
	FromSymbol string   `s2m:"fromSymbol,omitempty"`
	ToSymbol   string   `s2m:"toSymbol,omitempty"`
}

type UniversalTransferResp struct {
	TranId int64 `json:"tranId,omitempty"`
}

var UniversalTransferConfig = cex.ReqConfig[UniversalTransferParams, UniversalTransferResp]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/asset/transfer",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[UniversalTransferResp]),
}

type FlexibleProductListParams struct {
	Asset   string `s2m:"asset,omitempty"`
	Current int64  `s2m:"current,omitempty"`
	Size    int64  `s2m:"size,omitempty"`
}

type FlexibleProduct struct {
	Asset                      string            `json:"asset"`
	LatestAnnualPercentageRate float64           `json:"latestAnnualPercentageRate,string"`
	TierAnnualPercentageRate   map[string]string `json:"tierAnnualPercentageRate"`
	AirDropPercentageRate      float64           `json:"airDropPercentageRate,string"`
	CanPurchase                bool              `json:"canPurchase"`
	CanRedeem                  bool              `json:"canRedeem"`
	IsSoldOut                  bool              `json:"isSoldOut"`
	Hot                        bool              `json:"hot"`
	MinPurchaseAmount          float64           `json:"minPurchaseAmount,string"`
	ProductId                  string            `json:"productId"`
	SubscriptionStartTime      int64             `json:"subscriptionStartTime"`
	Status                     string            `json:"status"`
}

var FlexibleProductConfig = cex.ReqConfig[FlexibleProductListParams, Page[[]FlexibleProduct]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/simple-earn/flexible/list",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]FlexibleProduct]]),
}

type CryptoLoansIncomeHistoriesParams struct {
	Asset     string               `s2m:"asset,omitempty"`
	Type      CryptoLoanIncomeType `s2m:"type,omitempty"`
	StartTime int64                `s2m:"startTime,omitempty"`
	EndTime   int64                `s2m:"endTime,omitempty"`
	Limit     int                  `s2m:"limit,omitempty"`
}

type CryptoLoanIncomeHistory struct {
	Asset     string               `json:"asset"`
	Type      CryptoLoanIncomeType `json:"type"`
	Amount    float64              `json:"amount,string"`
	Timestamp int64                `json:"timestamp"`
	TranId    string               `json:"tranId"`
}

var CryptoLoansIncomeHistoriesConfig = cex.ReqConfig[CryptoLoansIncomeHistoriesParams, []CryptoLoanIncomeHistory]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/income",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[[]CryptoLoanIncomeHistory]),
}

type FlexibleBorrowParams struct {
	LoanCoin         string  `s2m:"loanCoin,omitempty"`
	LoanAmount       float64 `s2m:"loanAmount,omitempty"`
	CollateralCoin   string  `s2m:"collateralCoin,omitempty"`
	CollateralAmount float64 `s2m:"collateralAmount,omitempty"`
}

type FlexibleBorrowResult struct {
	LoanCoin         string               `json:"loanCoin"`
	LoanAmount       float64              `json:"loanAmount,string"`
	CollateralCoin   string               `json:"collateralCoin"`
	CollateralAmount float64              `json:"collateralAmount,string"`
	Status           FlexibleBorrowStatus `json:"status"`
}

var FlexibleBorrowConfig = cex.ReqConfig[FlexibleBorrowParams, FlexibleBorrowResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/borrow",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[FlexibleBorrowResult]),
}

type FlexibleOngoingOrdersParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	Current        int    `s2m:"current,omitempty"` // default: 1, max: 1000
	Limit          int    `s2m:"limit,omitempty"`   // default: 10, max: 100
}

type FlexibleOngoingOrder struct {
	LoanCoin         string  `json:"loanCoin" bson:"loanCoin"`
	TotalDebt        float64 `json:"totalDebt,string" bson:"totalDebt"`
	CollateralCoin   string  `json:"collateralCoin" bson:"collateralCoin"`
	CollateralAmount float64 `json:"collateralAmount,string" bson:"collateralAmount"`
	CurrentLTV       float64 `json:"currentLTV,string" bson:"currentLTV"`
}

var FlexibleOngoingOrdersConfig = cex.ReqConfig[FlexibleOngoingOrdersParams, Page[[]FlexibleOngoingOrder]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/ongoing/orders",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]FlexibleOngoingOrder]]),
}
