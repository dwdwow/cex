package bnc

import (
	"net/http"

	"github.com/dwdwow/cex"
)

type CoinNetworkInfo struct {
	AddressRegex            string `json:"addressRegex" bson:"addressRegex"`
	Coin                    string `json:"coin" bson:"coin"`
	DepositDesc             string `json:"depositDesc" bson:"depositDesc"`
	DepositEnable           bool   `json:"depositEnable" bson:"depositEnable"`
	IsDefault               bool   `json:"isDefault" bson:"isDefault"`
	MemoRegex               string `json:"memoRegex" bson:"memoRegex"`
	MinConfirm              int    `json:"minConfirm" bson:"minConfirm"`
	Name                    string `json:"name" bson:"name"`
	Network                 string `json:"network" bson:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus" bson:"resetAddressStatus"`
	SpecialTips             string `json:"specialTips" bson:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm" bson:"unLockConfirm"`
	WithdrawDesc            string `json:"withdrawDesc" bson:"withdrawDesc"`
	WithdrawEnable          bool   `json:"withdrawEnable" bson:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee" bson:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple" bson:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax" bson:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin" bson:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress" bson:"sameAddress"`
	EstimatedArrivalTime    int    `json:"estimatedArrivalTime" bson:"estimatedArrivalTime"`
	Busy                    bool   `json:"busy" bson:"busy"`
}

type Coin struct {
	Coin             string            `json:"coin" bson:"coin"`
	DepositAllEnable bool              `json:"depositAllEnable" bson:"depositAllEnable"`
	Free             float64           `json:"free,string" bson:"free,string"`
	Freeze           float64           `json:"freeze,string" bson:"freeze,string"`
	Ipoable          float64           `json:"ipoable,string" bson:"ipoable,string"`
	Ipoing           float64           `json:"ipoing,string" bson:"ipoing,string"`
	IsLegalMoney     bool              `json:"isLegalMoney" bson:"isLegalMoney"`
	Locked           float64           `json:"locked,string" bson:"locked,string"`
	Name             string            `json:"name" bson:"name"`
	NetworkList      []CoinNetworkInfo `json:"networkList" bson:"networkList"`
}

var CoinInfoConfig = cex.ReqConfig[cex.NilReqData, []Coin]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/capital/config/getall",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]Coin]),
}

type SpotBalance struct {
	Asset  string  `json:"asset" bson:"asset"`
	Free   float64 `json:"free,string" bson:"free,string"`
	Locked float64 `json:"locked,string" bson:"locked,string"`
}

type SpotAccount struct {
	MakerCommission  float64 `json:"makerCommission" bson:"makerCommission"`
	TakerCommission  float64 `json:"takerCommission" bson:"takerCommission"`
	BuyerCommission  float64 `json:"buyerCommission" bson:"buyerCommission"`
	SellerCommission float64 `json:"sellerCommission" bson:"sellerCommission"`
	CommissionRates  struct {
		Maker  float64 `json:"maker,string" bson:"maker,string"`
		Taker  float64 `json:"taker,string" bson:"taker,string"`
		Buyer  float64 `json:"buyer,string" bson:"buyer,string"`
		Seller float64 `json:"seller,string" bson:"seller,string"`
	} `json:"commissionRates" bson:"commissionRates" bson:"commissionRates"`
	CanTrade                   bool          `json:"canTrade" bson:"canTrade"`
	CanWithdraw                bool          `json:"canWithdraw" bson:"canWithdraw"`
	CanDeposit                 bool          `json:"canDeposit" bson:"canDeposit"`
	Brokered                   bool          `json:"brokered" bson:"brokered"`
	RequireSelfTradePrevention bool          `json:"requireSelfTradePrevention" bson:"requireSelfTradePrevention"`
	UpdateTime                 int64         `json:"updateTime" bson:"updateTime"`
	AccountType                AccountType   `json:"accountType" bson:"accountType"`
	Balances                   []SpotBalance `json:"balances" bson:"balances"`
	Permissions                []PairType    `json:"permissions" bson:"permissions"`
}

var SpotAccountConfig = cex.ReqConfig[cex.NilReqData, SpotAccount]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/account",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[SpotAccount]),
}

type UniversalTransferParams struct {
	Type       TransferType `s2m:"type,omitempty"`
	Asset      string       `s2m:"asset,omitempty"`
	Amount     float64      `s2m:"amount,omitempty"`
	FromSymbol string       `s2m:"fromSymbol,omitempty"`
	ToSymbol   string       `s2m:"toSymbol,omitempty"`
}

type UniversalTransferResp struct {
	TranId int64 `json:"tranId,omitempty" bson:"tranId,omitempty"`
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
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[UniversalTransferResp]),
}

// =============================================
// Wallet
// ---------------------------------------------

type WithdrawParams struct {
	Coin               string     `s2m:"coin,omitempty"`
	WithdrawOrderId    string     `s2m:"withdrawOrderId,omitempty"`
	Network            Network    `s2m:"network,omitempty"`
	Address            string     `s2m:"address,omitempty"`
	AddressTag         string     `s2m:"addressTag,omitempty"`
	Amount             float64    `s2m:"amount,omitempty"`
	TransactionFeeFlag bool       `s2m:"transactionFeeFlag,omitempty"` //	When making internal transfer, true for returning the fee to the destination account; false for returning the fee back to the departure account. Default false.
	Name               string     `s2m:"name,omitempty"`               //	Description of the address. Space in name should be encoded into %20.
	WalletType         WalletType `s2m:"walletType,omitempty"`
}

type WithdrawResult struct {
	Id string `json:"id"`
}

var WithdrawConfig = cex.ReqConfig[WithdrawParams, WithdrawResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/capital/withdraw/apply",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[WithdrawResult]),
}

type DepositAddressParams struct {
	Coin    string  `s2m:"coin,omitempty"`
	Network Network `s2m:"network,omitempty"`
	Amount  float64 `s2m:"amount,omitempty"`
}

type DepositAddress struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	Url     string `json:"url"`
}

var DepositAddressConfig = cex.ReqConfig[DepositAddressParams, DepositAddress]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/capital/deposit/address",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[DepositAddress]),
}

// ---------------------------------------------
// Wallet
// =============================================

// =============================================
// Flexible Simple Earn
// ---------------------------------------------

type SimpleEarnFlexibleProductListParams struct {
	Asset   string `s2m:"asset,omitempty"`
	Current int64  `s2m:"current,omitempty"`
	Size    int64  `s2m:"size,omitempty"` // default:10, max: 100
}

type SimpleEarnFlexibleProduct struct {
	Asset                      string            `json:"asset" bson:"asset"`
	LatestAnnualPercentageRate float64           `json:"latestAnnualPercentageRate,string" bson:"latestAnnualPercentageRate,string"`
	TierAnnualPercentageRate   map[string]string `json:"tierAnnualPercentageRate" bson:"tierAnnualPercentageRate"`
	AirDropPercentageRate      float64           `json:"airDropPercentageRate,string" bson:"airDropPercentageRate,string"`
	CanPurchase                bool              `json:"canPurchase" bson:"canPurchase"`
	CanRedeem                  bool              `json:"canRedeem" bson:"canRedeem"`
	IsSoldOut                  bool              `json:"isSoldOut" bson:"isSoldOut"`
	Hot                        bool              `json:"hot" bson:"hot"`
	MinPurchaseAmount          float64           `json:"minPurchaseAmount,string" bson:"minPurchaseAmount,string"`
	ProductId                  string            `json:"productId" bson:"productId"`
	SubscriptionStartTime      int64             `json:"subscriptionStartTime" bson:"subscriptionStartTime"`
	Status                     string            `json:"status" bson:"status"`
}

var SimpleEarnFlexibleProductConfig = cex.ReqConfig[SimpleEarnFlexibleProductListParams, Page[[]SimpleEarnFlexibleProduct]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/simple-earn/flexible/list",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]SimpleEarnFlexibleProduct]]),
}

type SimpleEarnFlexibleRedeemParams struct {
	ProductId   string                              `s2m:"productId,omitempty"`
	RedeemAll   bool                                `s2m:"redeemAll,omitempty"` //	true or false, default to false
	Amount      float64                             `s2m:"amount,omitempty"`    //	if redeemAll is false, amount is mandatory
	DestAccount SimpleEarnFlexibleRedeemDestination `s2m:"destAccount,omitempty"`
}

type SimpleEarnFlexibleRedeemResponse struct {
	RedeemId int64 `json:"redeemId,omitempty" bson:"redeemId,omitempty"`
	Success  bool  `json:"success,omitempty" bson:"success,omitempty"`
}

var SimpleEarnFlexibleRedeemConfig = cex.ReqConfig[SimpleEarnFlexibleRedeemParams, SimpleEarnFlexibleRedeemResponse]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/simple-earn/flexible/redeem",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[SimpleEarnFlexibleRedeemResponse]),
}

type SimpleEarnFlexiblePositionsParams struct {
	Asset     string `s2m:"asset,omitempty"`
	ProductId string `s2m:"productId,omitempty"`
	Current   int    `s2m:"current,omitempty"`
	Size      int    `s2m:"size,omitempty"` // default 10, max 100
}

type SimpleEarnFlexiblePosition struct {
	TotalAmount                    float64           `json:"totalAmount,string" bson:"totalAmount,string"`
	TierAnnualPercentageRate       map[string]string `json:"tierAnnualPercentageRate" bson:"tierAnnualPercentageRate"`
	LatestAnnualPercentageRate     float64           `json:"latestAnnualPercentageRate,string" bson:"latestAnnualPercentageRate,string"`
	YesterdayAirdropPercentageRate float64           `json:"yesterdayAirdropPercentageRate,string" bson:"yesterdayAirdropPercentageRate,string"`
	Asset                          string            `json:"asset" bson:"asset"`               // raw symbol, is not with prefix, LD
	AirDropAsset                   string            `json:"airDropAsset" bson:"airDropAsset"` // do not know meanings of this
	CanRedeem                      bool              `json:"canRedeem" bson:"canRedeem"`
	CollateralAmount               float64           `json:"collateralAmount,string" bson:"collateralAmount,string"` // is the amount of crypto loans
	ProductId                      string            `json:"productId" bson:"productId"`
	YesterdayRealTimeRewards       float64           `json:"yesterdayRealTimeRewards,string" bson:"yesterdayRealTimeRewards,string"`
	CumulativeBonusRewards         float64           `json:"cumulativeBonusRewards,string" bson:"cumulativeBonusRewards,string"`
	CumulativeRealTimeRewards      float64           `json:"cumulativeRealTimeRewards,string" bson:"cumulativeRealTimeRewards,string"`
	CumulativeTotalRewards         float64           `json:"cumulativeTotalRewards,string" bson:"cumulativeTotalRewards,string"`
	AutoSubscribe                  bool              `json:"autoSubscribe" bson:"autoSubscribe"`
}

var SimpleEarnFlexiblePositionsConfig = cex.ReqConfig[SimpleEarnFlexiblePositionsParams, Page[[]SimpleEarnFlexiblePosition]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/simple-earn/flexible/position",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]SimpleEarnFlexiblePosition]]),
}

type SimpleEarnFlexibleRateHistoryParams struct {
	ProductId string `s2m:"productId,omitempty"`
	StartTime int64  `s2m:"startTime,omitempty"`
	EndTime   int64  `s2m:"endTime,omitempty"`
	Current   int    `s2m:"current,omitempty"` // Currently querying page. Start from 1. Default:1
	Size      int    `s2m:"size,omitempty"`    // Default:10, Max:100
}

type SimpleEarnFlexibleRateHistory struct {
	ProductId            string  `json:"productId"`
	Asset                string  `json:"asset"`
	AnnualPercentageRate float64 `json:"annualPercentageRate,string"`
	Time                 int64   `json:"time"`
}

var SimpleEarnFlexibleRateHistoryConfig = cex.ReqConfig[SimpleEarnFlexibleRateHistoryParams, Page[[]SimpleEarnFlexibleRateHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/simple-earn/flexible/history/rateHistory",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]SimpleEarnFlexibleRateHistory]]),
}

// ---------------------------------------------
// Flexible Simple Earn
// =============================================

// =============================================
// Crypto Flexible Loans
// ---------------------------------------------

type CryptoLoansIncomeHistoriesParams struct {
	Asset     string               `s2m:"asset,omitempty"`
	Type      CryptoLoanIncomeType `s2m:"type,omitempty"`
	StartTime int64                `s2m:"startTime,omitempty"`
	EndTime   int64                `s2m:"endTime,omitempty"`
	Limit     int                  `s2m:"limit,omitempty"` // default:20, max:100
}

type CryptoLoanIncomeHistory struct {
	Asset     string               `json:"asset" bson:"asset"`
	Type      CryptoLoanIncomeType `json:"type" bson:"type"`
	Amount    float64              `json:"amount,string" bson:"amount,string"`
	Timestamp int64                `json:"timestamp" bson:"timestamp"`
	TranId    string               `json:"tranId" bson:"tranId"`
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
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]CryptoLoanIncomeHistory]),
}

type CryptoLoanFlexibleBorrowParams struct {
	LoanCoin         string  `s2m:"loanCoin,omitempty"`
	LoanAmount       float64 `s2m:"loanAmount,omitempty"`
	CollateralCoin   string  `s2m:"collateralCoin,omitempty"`
	CollateralAmount float64 `s2m:"collateralAmount,omitempty"`
}

type CryptoLoanFlexibleBorrowResult struct {
	LoanCoin         string                         `json:"loanCoin" bson:"loanCoin"`
	LoanAmount       float64                        `json:"loanAmount,string" bson:"loanAmount,string"`
	CollateralCoin   string                         `json:"collateralCoin" bson:"collateralCoin"`
	CollateralAmount float64                        `json:"collateralAmount,string" bson:"collateralAmount,string"`
	Status           CryptoLoanFlexibleBorrowStatus `json:"status" bson:"status"`
}

var CryptoLoanFlexibleBorrowConfig = cex.ReqConfig[CryptoLoanFlexibleBorrowParams, CryptoLoanFlexibleBorrowResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/borrow",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[CryptoLoanFlexibleBorrowResult]),
}

type CryptoLoanFlexibleOngoingOrdersParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	Current        int    `s2m:"current,omitempty"` // default: 1, max: 1000
	Limit          int    `s2m:"limit,omitempty"`   // default: 10, max: 100
}

type CryptoLoanFlexibleOngoingOrder struct {
	LoanCoin         string  `json:"loanCoin" bson:"loanCoin"`
	TotalDebt        float64 `json:"totalDebt,string" bson:"totalDebt,string"`
	CollateralCoin   string  `json:"collateralCoin" bson:"collateralCoin"`
	CollateralAmount float64 `json:"collateralAmount,string" bson:"collateralAmount,string"`
	CurrentLTV       float64 `json:"currentLTV,string" bson:"currentLTV,string"`
}

var CryptoLoanFlexibleOngoingOrdersConfig = cex.ReqConfig[CryptoLoanFlexibleOngoingOrdersParams, Page[[]CryptoLoanFlexibleOngoingOrder]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/ongoing/orders",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]CryptoLoanFlexibleOngoingOrder]]),
}

type CryptoLoanFlexibleBorrowHistoriesParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	StartTime      int64  `s2m:"startTime,omitempty"`
	EndTime        int64  `s2m:"endTime,omitempty"`
	Current        int64  `s2m:"current,omitempty"` // default: 1, max: 1000
	Limit          int64  `s2m:"limit,omitempty"`   // default: 10, max: 100
}

type CryptoLoanFlexibleBorrowHistory struct {
	LoanCoin                string                         `json:"loanCoin" bson:"loanCoin"`
	InitialLoanAmount       string                         `json:"initialLoanAmount" bson:"initialLoanAmount"`
	CollateralCoin          string                         `json:"collateralCoin" bson:"collateralCoin"`
	InitialCollateralAmount string                         `json:"initialCollateralAmount" bson:"initialCollateralAmount"`
	BorrowTime              int64                          `json:"borrowTime,string" bson:"borrowTime,string"`
	Status                  CryptoLoanFlexibleBorrowStatus `json:"status" bson:"status"`
}

var CryptoLoanFlexibleBorrowHistoriesConfig = cex.ReqConfig[CryptoLoanFlexibleBorrowHistoriesParams, Page[[]CryptoLoanFlexibleBorrowHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/borrow/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]CryptoLoanFlexibleBorrowHistory]]),
}

type CryptoLoanFlexibleRepayParams struct {
	LoanCoin         string  `s2m:"loanCoin,omitempty"`
	CollateralCoin   string  `s2m:"collateralCoin,omitempty"`
	RepayAmount      float64 `s2m:"repayAmount,omitempty"`
	CollateralReturn BigBool `s2m:"collateralReturn,omitempty"`
	FullRepayment    BigBool `s2m:"fullRepayment,omitempty"`
}

type CryptoLoanFlexibleRepayResult struct {
	LoanCoin            string                    `json:"loanCoin" bson:"loanCoin"`
	CollateralCoin      string                    `json:"collateralCoin" bson:"collateralCoin"`
	RemainingDebt       string                    `json:"remainingDebt" bson:"remainingDebt"`
	RemainingCollateral string                    `json:"remainingCollateral" bson:"remainingCollateral"`
	FullRepayment       bool                      `json:"fullRepayment" bson:"fullRepayment"`
	CurrentLTV          string                    `json:"currentLTV" bson:"currentLTV"`
	RepayStatus         CryptoFlexibleRepayStatus `json:"repayStatus" bson:"repayStatus"`
}

var CryptoLoanFlexibleRepayConfig = cex.ReqConfig[CryptoLoanFlexibleRepayParams, CryptoLoanFlexibleRepayResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/repay",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[CryptoLoanFlexibleRepayResult]),
}

type CryptoLoanFlexibleRepaymentHistoriesParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	StartTime      int64  `s2m:"startTime,omitempty"`
	EndTime        int64  `s2m:"endTime,omitempty"`
	Current        int64  `s2m:"current,omitempty"` //	start from 1; default: 1; max: 1000
	Limit          int64  `s2m:"limit,omitempty"`   // default: 10; max: 100
}

type CryptoLoanFlexibleRepaymentHistory struct {
	LoanCoin         string                    `json:"loanCoin" bson:"loanCoin"`
	RepayAmount      float64                   `json:"repayAmount,string" bson:"repayAmount,string"`
	CollateralCoin   string                    `json:"collateralCoin" bson:"collateralCoin"`
	CollateralReturn float64                   `json:"collateralReturn,string" bson:"collateralReturn,string"`
	RepayStatus      CryptoFlexibleRepayStatus `json:"repayStatus" bson:"repayStatus"`
	RepayTime        int64                     `json:"repayTime,string" bson:"repayTime,string"`
}

var CryptoLoanFlexibleRepaymentHistoriesConfig = cex.ReqConfig[CryptoLoanFlexibleRepaymentHistoriesParams, Page[[]CryptoLoanFlexibleRepaymentHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/repay/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]CryptoLoanFlexibleRepaymentHistory]]),
}

type CryptoLoanFlexibleAdjustLtvParams struct {
	LoanCoin         string             `s2m:"loanCoin,omitempty"`
	CollateralCoin   string             `s2m:"collateralCoin,omitempty"`
	AdjustmentAmount float64            `s2m:"adjustmentAmount,omitempty"`
	Direction        LTVAdjustDirection `s2m:"direction,omitempty"`
}

type CryptoLoanFlexibleLoanAdjustLtvResult struct {
	LoanCoin         string             `json:"loanCoin" bson:"loanCoin"`
	CollateralCoin   string             `json:"collateralCoin" bson:"collateralCoin"`
	Direction        LTVAdjustDirection `json:"direction" bson:"direction"`
	AdjustmentAmount float64            `json:"adjustmentAmount,string" bson:"adjustmentAmount,string"`
	CurrentLTV       float64            `json:"currentLTV,string" bson:"currentLTV,string"`
}

var CryptoLoanFlexibleLoanAdjustLtvConfig = cex.ReqConfig[CryptoLoanFlexibleAdjustLtvParams, CryptoLoanFlexibleLoanAdjustLtvResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/adjust/ltv",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[CryptoLoanFlexibleLoanAdjustLtvResult]),
}

type CryptoLoanFlexibleAdjustLtvHistoriesParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	StartTime      int64  `s2m:"startTime,omitempty"`
	EndTime        int64  `s2m:"endTime,omitempty"`
	Current        int64  `s2m:"current,omitempty"` // start from 1; default: 1; max: 1000
	Limit          int64  `s2m:"limit,omitempty"`   // default: 10; max: 100
}

type CryptoLoanFlexibleAdjustLtvHistory struct {
	LoanCoin         string `json:"loanCoin" bson:"loanCoin"`
	CollateralCoin   string `json:"collateralCoin" bson:"collateralCoin"`
	Direction        string `json:"direction" bson:"direction"`
	CollateralAmount string `json:"collateralAmount" bson:"collateralAmount"`
	PreLTV           string `json:"preLTV" bson:"preLTV"`
	AfterLTV         string `json:"afterLTV" bson:"afterLTV"`
	AdjustTime       int64  `json:"adjustTime,string" bson:"adjustTime,string"`
}

var CryptoLoanFlexibleAdjustLtvHistoriesConfig = cex.ReqConfig[CryptoLoanFlexibleAdjustLtvHistoriesParams, Page[[]CryptoLoanFlexibleAdjustLtvHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/ltv/adjustment/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]CryptoLoanFlexibleAdjustLtvHistory]]),
}

type CryptoLoanFlexibleLoanAssetsParams struct {
	LoanCoin string `s2m:"loanCoin,omitempty"`
}

type CryptoLoanFlexibleLoanAsset struct {
	LoanCoin             string  `json:"loanCoin" bson:"loanCoin"`
	FlexibleInterestRate float64 `json:"flexibleInterestRate,string" bson:"flexibleInterestRate,string"`
	FlexibleMinLimit     float64 `json:"flexibleMinLimit,string" bson:"flexibleMinLimit,string"`
	FlexibleMaxLimit     float64 `json:"flexibleMaxLimit,string" bson:"flexibleMaxLimit,string"`
}

var CryptoLoanFlexibleLoanAssetsConfig = cex.ReqConfig[CryptoLoanFlexibleLoanAssetsParams, Page[[]CryptoLoanFlexibleLoanAsset]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/loanable/data",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]CryptoLoanFlexibleLoanAsset]]),
}

type CryptoLoanFlexibleCollateralCoinsParams struct {
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
}

type CryptoLoanFlexibleCollateralCoin struct {
	CollateralCoin string  `json:"collateralCoin" bson:"collateralCoin"`
	InitialLTV     float64 `json:"initialLTV,string" bson:"initialLTV,string"`
	MarginCallLTV  float64 `json:"marginCallLTV,string" bson:"marginCallLTV,string"`
	LiquidationLTV float64 `json:"liquidationLTV,string" bson:"liquidationLTV,string"`
	MaxLimit       float64 `json:"maxLimit,string" bson:"maxLimit,string"`
}

var CryptoLoanFlexibleCollateralCoinsConfig = cex.ReqConfig[CryptoLoanFlexibleCollateralCoinsParams, Page[[]CryptoLoanFlexibleCollateralCoin]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV2 + "/loan/flexible/collateral/data",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]CryptoLoanFlexibleCollateralCoin]]),
}

// ---------------------------------------------
// Crypto Flexible Loans
// =============================================

// =============================================
// VIP Flexible Loans
// ---------------------------------------------

type VIPLoanOngoingOrder struct {
	OrderId                          int64   `json:"orderId"`
	LoanCoin                         string  `json:"loanCoin"`
	TotalDebt                        float64 `json:"totalDebt,string"`
	LoanRate                         string  `json:"loanRate"` // maybe Flexible Rate
	ResidualInterest                 float64 `json:"residualInterest,string"`
	CollateralAccountId              string  `json:"collateralAccountId"`
	CollateralCoin                   string  `json:"collateralCoin"`
	TotalCollateralValueAfterHaircut float64 `json:"totalCollateralValueAfterHaircut,string"`
	LockedCollateralValue            float64 `json:"lockedCollateralValue,string"`
	CurrentLTV                       float64 `json:"currentLTV,string"`
	ExpirationTime                   int64   `json:"expirationTime"`
	LoanDate                         string  `json:"loanDate"`
	LoanTerm                         string  `json:"loanTerm"`
	InitialLtv                       string  `json:"initialLtv"`     // x%
	MarginCallLtv                    string  `json:"marginCallLtv"`  // x%
	LiquidationLtv                   string  `json:"liquidationLtv"` // x%
}

type VIPLoanOngoingOrderParams struct {
	OrderId             int64  `s2m:"orderId,omitempty"`
	CollateralAccountId int64  `s2m:"collateralAccountId,omitempty"`
	LoanCoin            string `s2m:"loanCoin,omitempty"`
	CollateralCoin      string `s2m:"collateralCoin,omitempty"`
	Current             int64  `s2m:"current,omitempty"` //	Currently querying page. Start from 1, Default:1, Max: 1000.
	Limit               int64  `s2m:"limit,omitempty"`   //	Default: 10, Max: 100
}

var VIPLoanOngoingOrderQueryConfig = cex.ReqConfig[VIPLoanOngoingOrderParams, Page[[]VIPLoanOngoingOrder]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/ongoing/orders",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]VIPLoanOngoingOrder]]),
}

type VIPLoanRepayParams struct {
	OrderId int64   `s2m:"orderId,omitempty"`
	Amount  float64 `s2m:"amount,omitempty"`
}

type VIPLoanRepayStatus string

const (
	VIPLoanRepayStatusRepaid   VIPLoanRepayStatus = "Repaid"
	VIPLoanRepayStatusRepaying VIPLoanRepayStatus = "Repaying"
	VIPLoanRepayStatusFailed   VIPLoanRepayStatus = "Failed"
)

type VIPLoanOrderStatus string

const (
	VIPLoanOrderStatusRepaid           VIPLoanOrderStatus = "Repaid"
	VIPLoanOrderStatusRepaying         VIPLoanOrderStatus = "Repaying"
	VIPLoanOrderStatusFailed           VIPLoanOrderStatus = "Failed"
	VIPLoanOrderStatusAccruingInterest VIPLoanOrderStatus = "Accruing_Interest"
	VIPLoanOrderStatusOverdue          VIPLoanOrderStatus = "Overdue"
	VIPLoanOrderStatusLiquidating      VIPLoanOrderStatus = "Liquidating"
	VIPLoanOrderStatusLiquidated       VIPLoanOrderStatus = "Liquidated"
	VIPLoanOrderStatusPending          VIPLoanOrderStatus = "Pending"
)

type VIPLoanRepayResult struct {
	LoanCoin           string             `json:"loanCoin"`
	RepayAmount        float64            `json:"repayAmount,string"`
	RemainingPrincipal float64            `json:"remainingPrincipal,string"`
	RemainingInterest  float64            `json:"remainingInterest,string"`
	CurrentLTV         float64            `json:"currentLTV,string"`
	CollateralCoin     string             `json:"collateralCoin"`
	RepayStatus        VIPLoanRepayStatus `json:"repayStatus"`
}

var VIPLoanRepayConfig = cex.ReqConfig[VIPLoanRepayParams, VIPLoanRepayResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/repay",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[VIPLoanRepayResult]),
}

type VIPLoanRepayHistoryParams struct {
	OrderId   int64  `s2m:"orderId,omitempty"`
	LoanCoin  string `s2m:"loanCoin,omitempty"`
	StartTime int64  `s2m:"startTime,omitempty"`
	EndTime   int64  `s2m:"endTime,omitempty"`
	Current   int64  `s2m:"current,omitempty"` //	Currently querying page. Start from 1, Default:1, Max: 1000
	Limit     int64  `s2m:"limit,omitempty"`   //	Default: 10, Max: 100
}

type VIPLoanRepayHistory struct {
	LoanCoin       string             `json:"loanCoin"`
	RepayAmount    float64            `json:"repayAmount,string"`
	CollateralCoin string             `json:"collateralCoin"`
	RepayStatus    VIPLoanRepayStatus `json:"repayStatus"`
	LoanDate       int64              `json:"loanDate,string"`
	RepayTime      int64              `json:"repayTime,string"`
	OrderId        string             `json:"orderId"`
}

var VIPLoanRepayHistoryConfig = cex.ReqConfig[VIPLoanRepayHistoryParams, Page[[]VIPLoanRepayHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/repay/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]VIPLoanRepayHistory]]),
}

type VIPLoanLockedValue struct {
	CollateralAccountId string `json:"collateralAccountId"`
	CollateralCoin      string `json:"collateralCoin"`
}

type VIPLoanLockedValueQueryParams struct {
	OrderId             int64 `s2m:"orderId,omitempty"`
	CollateralAccountId int64 `s2m:"collateralAccountId,omitempty"`
}

var VIPLoanLockedValueConfig = cex.ReqConfig[VIPLoanLockedValueQueryParams, Page[[][]VIPLoanLockedValue]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/collateral/account",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[][]VIPLoanLockedValue]]),
}

type VIPLoanBorrowParams struct {
	LoanAccountId       int64   `s2m:"loanAccountId,omitempty"`
	LoanCoin            string  `s2m:"loanCoin,omitempty"`
	LoanAmount          float64 `s2m:"loanAmount,omitempty"`
	CollateralAccountId string  `s2m:"collateralAccountId,omitempty"` // Multiple split by ,
	CollateralCoin      string  `s2m:"collateralCoin,omitempty"`      // Multiple split by,
	IsFlexibleRate      BigBool `s2m:"isFlexibleRate,omitempty"`      // Default: TRUE.TRUE: flexible rate FALSE: fixed rate
	LoanTerm            int64   `s2m:"loanTerm,omitempty"`            // Mandatory for fixed rate.Optional for fixed interest rate.Eg: 30/60 days
}

type VIPLoanBorrowResult struct {
	LoanAccountId       string  `json:"loanAccountId"`
	RequestId           string  `json:"requestId"`
	LoanCoin            string  `json:"loanCoin"`
	IsFlexibleRate      YesNo   `json:"isFlexibleRate"`
	LoanAmount          float64 `json:"loanAmount,string"`
	CollateralAccountId string  `json:"collateralAccountId"`
	CollateralCoin      string  `json:"collateralCoin"`
	LoanTerm            string  `json:"loanTerm"`
}

var VIPLoanBorrowConfig = cex.ReqConfig[VIPLoanBorrowParams, VIPLoanBorrowResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/borrow",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[VIPLoanBorrowResult]),
}

type VIPLoanableAsset struct {
	LoanCoin                   string  `json:"loanCoin"`
	FlexibleHourlyInterestRate float64 `json:"_flexibleHourlyInterestRate,string"`
	FlexibleYearlyInterestRate float64 `json:"_flexibleYearlyInterestRate,string"`
	DDailyInterestRate         float64 `json:"_30dDailyInterestRate,string"`
	DYearlyInterestRate        float64 `json:"_30dYearlyInterestRate,string"`
	DDailyInterestRate1        float64 `json:"_60dDailyInterestRate,string"`
	DYearlyInterestRate1       float64 `json:"_60dYearlyInterestRate,string"`
	MinLimit                   float64 `json:"minLimit,string"`
	MaxLimit                   float64 `json:"maxLimit,string"`
	VipLevel                   int     `json:"vipLevel"`
}

type VIPLoanableAssetQueryParams struct {
	LoanCoin string `s2m:"loanCoin,omitempty"`
	VipLevel int    `s2m:"vipLevel,omitempty"`
}

var VIPLoanLoanableAssetsConfig = cex.ReqConfig[VIPLoanableAssetQueryParams, Page[[]VIPLoanableAsset]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/loanable/data",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]VIPLoanableAsset]]),
}

type VIPLoanCollateralAsset struct {
	CollateralCoin    string `json:"collateralCoin"`
	StCollateralRatio string `json:"_1stCollateralRatio"` // x%
	StCollateralRange string `json:"_1stCollateralRange"`
	NdCollateralRatio string `json:"_2ndCollateralRatio"`
	NdCollateralRange string `json:"_2ndCollateralRange"`
	RdCollateralRatio string `json:"_3rdCollateralRatio"`
	RdCollateralRange string `json:"_3rdCollateralRange"`
	ThCollateralRatio string `json:"_4thCollateralRatio"`
	ThCollateralRange string `json:"_4thCollateralRange"`
}

type VIPLoanCollateralAssetQueryParams struct {
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
}

var VIPLoanCollateralAssetsConfig = cex.ReqConfig[VIPLoanCollateralAssetQueryParams, Page[[]VIPLoanCollateralAsset]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/collateral/data",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]VIPLoanCollateralAsset]]),
}

type VIPLoanApplicationStatusQueryParams struct {
	Current int64 `s2m:"current,omitempty"` // Currently querying page. Start from 1, Default:1, Max: 1000
	Limit   int64 `s2m:"limit,omitempty"`   // Default: 10, Max: 100
}

type VIPLoanApplicationStatusInfo struct {
	LoanAccountId       string             `json:"loanAccountId"`
	OrderId             string             `json:"orderId"`
	RequestId           string             `json:"requestId"`
	LoanCoin            string             `json:"loanCoin"`
	LoanAmount          float64            `json:"loanAmount,string"`
	CollateralAccountId string             `json:"collateralAccountId"`
	CollateralCoin      string             `json:"collateralCoin"`
	LoanTerm            string             `json:"loanTerm"`
	Status              VIPLoanOrderStatus `json:"status"`
	LoanDate            int64              `json:"loanDate,string"`
}

var VIPLoanApplicationStatusConfig = cex.ReqConfig[VIPLoanApplicationStatusQueryParams, Page[[]VIPLoanApplicationStatusInfo]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/request/data",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]VIPLoanApplicationStatusInfo]]),
}

type VIPLoanInterestRateQueryParams struct {
	LoanCoin string `s2m:"loanCoin,omitempty"` // Max 10 assets, Multiple split by ","
}

type VIPLoanInterestRateInfo struct {
	Asset                      string  `json:"asset"`
	FlexibleDailyInterestRate  float64 `json:"flexibleDailyInterestRate,string"`
	FlexibleYearlyInterestRate float64 `json:"flexibleYearlyInterestRate,string"`
	Time                       int64   `json:"time"`
}

var VIPLoanInterestRatesConfig = cex.ReqConfig[VIPLoanInterestRateQueryParams, Page[[]VIPLoanInterestRateInfo]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/vip/request/interestRate",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]VIPLoanInterestRateInfo]]),
}

// ---------------------------------------------
// VIP Flexible Loans
// =============================================

// =============================================
// Spot Trading
// ---------------------------------------------

type SpotNewOrderParams struct {
	Symbol                  string                  `s2m:"symbol,omitempty"`
	Type                    OrderType               `s2m:"type,omitempty"`
	Side                    OrderSide               `s2m:"side,omitempty"`
	Quantity                float64                 `s2m:"quantity,omitempty"`
	Price                   float64                 `s2m:"price,omitempty"`
	TimeInForce             TimeInForce             `s2m:"timeInForce,omitempty"`
	NewClientOrderId        string                  `s2m:"newClientOrderId,omitempty"`
	QuoteOrderQty           float64                 `s2m:"quoteOrderQty,omitempty"`
	StrategyId              int64                   `s2m:"strategyId,omitempty"`
	StrategyType            int64                   `s2m:"strategyType,omitempty"`
	StopPrice               float64                 `s2m:"stopPrice,omitempty"`               // Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	TrailingDelta           int64                   `s2m:"trailingDelta,omitempty"`           // Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. For more details on SPOT implementation on trailing stops, please refer to Trailing Stop FAQ
	IcebergQty              float64                 `s2m:"icebergQty,omitempty"`              // Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
	NewOrderRespType        OrderResponseType       `s2m:"newOrderRespType,omitempty"`        // Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
	SelfTradePreventionMode SelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty"` // The allowed enums is dependent on what is configured on the symbol.The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE.
}

type SpotOrderFill struct {
	Price           float64 `json:"price,string" bson:"price,string"`
	Qty             float64 `json:"qty,string" bson:"qty,string"`
	Commission      float64 `json:"commission,string" bson:"commission,string"`
	CommissionAsset string  `json:"commissionAsset" bson:"commissionAsset"`
	TradeId         int64   `json:"tradeId" bson:"tradeId"`
}

type SpotOrder struct {
	// common
	Symbol                  string                  `json:"symbol" bson:"symbol"`
	OrderId                 int64                   `json:"orderId" bson:"orderId"`
	OrderListId             int64                   `json:"orderListId" bson:"orderListId"` // Unless OCO, value will be -1
	ClientOrderId           string                  `json:"clientOrderId" bson:"clientOrderId"`
	Price                   float64                 `json:"price,string" bson:"price,string"` // origin price, if market order, it is 0
	OrigQty                 float64                 `json:"origQty,string" bson:"origQty,string"`
	ExecutedQty             float64                 `json:"executedQty,string" bson:"executedQty,string"`
	CummulativeQuoteQty     float64                 `json:"cummulativeQuoteQty,string" bson:"cummulativeQuoteQty,string"`
	Status                  OrderStatus             `json:"status" bson:"status"`
	TimeInForce             TimeInForce             `json:"timeInForce" bson:"timeInForce"`
	Type                    OrderType               `json:"type" bson:"type"`
	Side                    OrderSide               `json:"side" bson:"side"`
	SelfTradePreventionMode SelfTradePreventionMode `json:"selfTradePreventionMode" bson:"selfTradePreventionMode"`

	// new order result, replace order result
	TransactTime int64 `json:"transactTime" bson:"transactTime"`

	// new order result, cancel new (replace) order
	Fills []SpotOrderFill `json:"fills" bson:"fills"`

	// new order result, query order, all orders
	WorkingTime int64 `json:"workingTime" bson:"workingTime"`

	// cancel order result, cancel new (replace) order
	OrigClientOrderId string `json:"origClientOrderId" bson:"origClientOrderId"`

	// query order, all orders
	StopPrice         float64 `json:"stopPrice,string" bson:"stopPrice,string"`
	IcebergQty        float64 `json:"icebergQty,string" bson:"icebergQty,string"`
	Time              int64   `json:"time" bson:"time"`
	UpdateTime        int64   `json:"updateTime" bson:"updateTime"`
	IsWorking         bool    `json:"isWorking" bson:"isWorking"`
	OrigQuoteOrderQty float64 `json:"origQuoteOrderQty,string" bson:"origQuoteOrderQty,string"`

	// for cancel replace order response
	Code int    `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
}

var SpotNewOrderConfig = cex.ReqConfig[SpotNewOrderParams, SpotOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/order",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[SpotOrder]),
}

type SpotCancelOrderParams struct {
	Symbol             string                 `s2m:"symbol,omitempty"`
	OrderId            int64                  `s2m:"orderId,omitempty"`
	OrigClientOrderId  string                 `s2m:"origClientOrderId,omitempty"`
	NewClientOrderId   string                 `s2m:"newClientOrderId,omitempty"`   // Used to uniquely identify this cancel. Automatically generated by default.
	CancelRestrictions OrderCancelRestriction `s2m:"cancelRestrictions,omitempty"` // Supported values: ONLY_NEW - Cancel will succeed if the order status is NEW. ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED
}

var SpotCancelOrderConfig = cex.ReqConfig[SpotCancelOrderParams, SpotOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/order",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[SpotOrder]),
}

type SpotCancelAllOpenOrdersParams struct {
	Symbol string `s2m:"symbol,omitempty"`
}

var SpotCancelAllOpenOrdersConfig = cex.ReqConfig[SpotCancelAllOpenOrdersParams, []SpotOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/openOrders",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]SpotOrder]),
}

type SpotQueryOrderParams struct {
	Symbol            string `s2m:"symbol,omitempty"` // must
	OrderId           int64  `s2m:"orderId,omitempty"`
	OrigClientOrderId string `s2m:"origClientOrderId,omitempty"`
}

var SpotQueryOrderConfig = cex.ReqConfig[SpotQueryOrderParams, SpotOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/order",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[SpotOrder]),
}

type SpotReplaceOrderParams struct {
	// must be filled, unless price if type is market
	Symbol                  string                `s2m:"symbol,omitempty"`
	Type                    OrderType             `s2m:"type,omitempty"`
	Side                    OrderSide             `s2m:"side,omitempty"`
	Quantity                float64               `s2m:"quantity,omitempty"`
	Price                   float64               `s2m:"price,omitempty"`
	TimeInForce             TimeInForce           `s2m:"timeInForce,omitempty"`
	CancelOrigClientOrderId string                `s2m:"cancelOrigClientOrderId,omitempty"` // Either the cancelOrigClientOrderId or cancelOrderId must be provided. If both are provided, cancelOrderId takes precedence.
	CancelOrderId           int64                 `s2m:"cancelOrderId,omitempty"`           // Either the cancelOrigClientOrderId or cancelOrderId must be provided. If both are provided, cancelOrderId takes precedence.
	CancelReplaceMode       SpotCancelReplaceMode `s2m:"cancelReplaceMode,omitempty"`       // The allowed values are: STOP_ON_FAILURE - If the cancel request fails, the new order placement will not be attempted. ALLOW_FAILURE - new order placement will be attempted even if cancel request fails.

	QuoteOrderQty           float64                 `s2m:"quoteOrderQty,omitempty"`
	CancelNewClientOrderId  string                  `s2m:"cancelNewClientOrderId,omitempty"` // Used to uniquely identify this cancel. Automatically generated by default.
	NewClientOrderId        string                  `s2m:"newClientOrderId,omitempty"`       // Used to identify the new order.
	StrategyId              int64                   `s2m:"strategyId,omitempty"`
	StrategyType            int64                   `s2m:"strategyType,omitempty"` // The value cannot be less than 1000000.
	StopPrice               float64                 `s2m:"stopPrice,omitempty"`
	TrailingDelta           int64                   `s2m:"trailingDelta,omitempty"`
	IcebergQty              float64                 `s2m:"icebergQty,omitempty"`
	NewOrderRespType        OrderResponseType       `s2m:"newOrderRespType,omitempty"`
	SelfTradePreventionMode SelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty"` // The allowed enums is dependent on what is configured on the symbol. The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE.
	CancelRestrictions      OrderCancelRestriction  `s2m:"cancelRestrictions,omitempty"`      // Supported values: ONLY_NEW - Cancel will succeed if the order status is NEW. ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED.
}

type SpotReplaceOrderRawData struct {
	CancelResult   SpotOrderCancelNewStatus `json:"cancelResult" bson:"cancelResult"`
	NewOrderResult SpotOrderCancelNewStatus `json:"newOrderResult" bson:"newOrderResult"`
	// response may be SpotOrder, CodeMsg, null
	// if result is SpotOrderCancelNewStatus_SUCCESS, response is SpotOrder
	// if result is SpotOrderCancelNewStatus_FAILURE, response is CodeMsg
	// if result is SpotOrderCancelNewStatus_NOT_ATTEMPTED, response is null
	CancelResponse   SpotOrder `json:"cancelResponse" bson:"cancelResponse"`
	NewOrderResponse SpotOrder `json:"newOrderResponse" bson:"newOrderResponse"`
}

type SpotReplaceOrderRawResult struct {
	Code int    `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
	// result: SUCCESS or FAILURE
	Data SpotReplaceOrderRawData `json:"data" bson:"data"`
}

type SpotReplaceOrderResult struct {
	OK          bool      `json:"OK" bson:"OK"`
	ErrCancel   error     `json:"errCancel" bson:"errCancel"`
	ErrNew      error     `json:"errNew" bson:"errNew"`
	OrderCancel SpotOrder `json:"orderCancel" bson:"orderCancel"`
	OrderNew    SpotOrder `json:"orderNew" bson:"orderNew"`
}

var SpotReplaceOrderConfig = cex.ReqConfig[SpotReplaceOrderParams, SpotReplaceOrderResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/order/cancelReplace",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotOrderReplaceUnmarshaler,
}

type SpotCurrentOpenOrdersParams struct {
	Symbol string `s2m:"symbol,omitempty"`
}

var SpotCurrentOpenOrdersConfig = cex.ReqConfig[SpotCurrentOpenOrdersParams, []SpotOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/openOrders",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]SpotOrder]),
}

// SpotAllOrdersParams
// If orderId is set, it will get orders >= that orderId. Otherwise, most recent orders are returned.
// For some historical orders cummulativeQuoteQty will be < 0, meaning the data is not available at this time.
// If startTime and/or endTime provided, orderId is not required.
// The payload sample does not show all fields that can appear.
type SpotAllOrdersParams struct {
	// Symbol is required
	Symbol    string `s2m:"symbol,omitempty"`
	OrderId   int64  `s2m:"orderId,omitempty"`
	StartTime int64  `s2m:"startTime,omitempty"`
	EndTime   int64  `s2m:"endTime,omitempty"`
	Limit     int64  `s2m:"limit,omitempty"` // Default 500; max 1000.
}

var SpotAllOrdersConfig = cex.ReqConfig[SpotAllOrdersParams, []SpotOrder]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/allOrders",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   spotBodyUnmshWrapper(cex.StdBodyUnmarshaler[[]SpotOrder]),
}

// ---------------------------------------------
// Spot Trading
// =============================================
