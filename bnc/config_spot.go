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

// =============================================
// Crypto Flexible Loans
// ---------------------------------------------

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

type FlexibleBorrowHistoriesParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	StartTime      int64  `s2m:"startTime,omitempty"`
	EndTime        int64  `s2m:"endTime,omitempty"`
	Current        int64  `s2m:"current,omitempty"` // default: 1, max: 1000
	Limit          int64  `s2m:"limit,omitempty"`   // default: 10, max: 100
}

type FlexibleBorrowHistory struct {
	LoanCoin                string               `json:"loanCoin"`
	InitialLoanAmount       string               `json:"initialLoanAmount"`
	CollateralCoin          string               `json:"collateralCoin"`
	InitialCollateralAmount string               `json:"initialCollateralAmount"`
	BorrowTime              int64                `json:"borrowTime,string"`
	Status                  FlexibleBorrowStatus `json:"status"`
}

var FlexibleBorrowHistoriesConfig = cex.ReqConfig[FlexibleBorrowHistoriesParams, Page[[]FlexibleBorrowHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/borrow/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]FlexibleBorrowHistory]]),
}

type FlexibleRepayParams struct {
	LoanCoin         string  `s2m:"loanCoin,omitempty"`
	CollateralCoin   string  `s2m:"collateralCoin,omitempty"`
	RepayAmount      float64 `s2m:"repayAmount,omitempty"`
	CollateralReturn BigBool `s2m:"collateralReturn,omitempty"`
	FullRepayment    BigBool `s2m:"fullRepayment,omitempty"`
}

type FlexibleRepayResult struct {
	LoanCoin            string              `json:"loanCoin"`
	CollateralCoin      string              `json:"collateralCoin"`
	RemainingDebt       string              `json:"remainingDebt"`
	RemainingCollateral string              `json:"remainingCollateral"`
	FullRepayment       bool                `json:"fullRepayment"`
	CurrentLTV          string              `json:"currentLTV"`
	RepayStatus         FlexibleRepayStatus `json:"repayStatus"`
}

var FlexibleRepayConfig = cex.ReqConfig[FlexibleRepayParams, FlexibleRepayResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/repay",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[FlexibleRepayResult]),
}

type FlexibleRepaymentHistoriesParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	StartTime      int64  `s2m:"startTime,omitempty"`
	EndTime        int64  `s2m:"endTime,omitempty"`
	Current        int64  `s2m:"current,omitempty"` //	start from 1; default: 1; max: 1000
	Limit          int64  `s2m:"limit,omitempty"`   // default: 10; max: 100
}

type FlexibleRepaymentHistory struct {
	LoanCoin         string              `json:"loanCoin"`
	RepayAmount      float64             `json:"repayAmount,string"`
	CollateralCoin   string              `json:"collateralCoin"`
	CollateralReturn float64             `json:"collateralReturn,string"`
	RepayStatus      FlexibleRepayStatus `json:"repayStatus"`
	RepayTime        int64               `json:"repayTime,string"`
}

var FlexibleRepaymentHistoriesConfig = cex.ReqConfig[FlexibleRepaymentHistoriesParams, Page[[]FlexibleRepaymentHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/repay/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]FlexibleRepaymentHistory]]),
}

type FlexibleAdjustLtvParams struct {
	LoanCoin         string             `s2m:"loanCoin,omitempty"`
	CollateralCoin   string             `s2m:"collateralCoin,omitempty"`
	AdjustmentAmount float64            `s2m:"adjustmentAmount,omitempty"`
	Direction        LTVAdjustDirection `s2m:"direction,omitempty"`
}

type FlexibleLoanAdjustLtvResult struct {
	LoanCoin         string             `json:"loanCoin"`
	CollateralCoin   string             `json:"collateralCoin"`
	Direction        LTVAdjustDirection `json:"direction"`
	AdjustmentAmount float64            `json:"adjustmentAmount,string"`
	CurrentLTV       float64            `json:"currentLTV,string"`
}

var FlexibleLoanAdjustLtvConfig = cex.ReqConfig[FlexibleAdjustLtvParams, FlexibleLoanAdjustLtvResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/adjust/ltv",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[FlexibleLoanAdjustLtvResult]),
}

type FlexibleAdjustLtvHistoriesParams struct {
	LoanCoin       string `s2m:"loanCoin,omitempty"`
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
	StartTime      int64  `s2m:"startTime,omitempty"`
	EndTime        int64  `s2m:"endTime,omitempty"`
	Current        int64  `s2m:"current,omitempty"` // start from 1; default: 1; max: 1000
	Limit          int64  `s2m:"limit,omitempty"`   // default: 10; max: 100
}

type FlexibleAdjustLtvHistory struct {
	LoanCoin         string `json:"loanCoin"`
	CollateralCoin   string `json:"collateralCoin"`
	Direction        string `json:"direction"`
	CollateralAmount string `json:"collateralAmount"`
	PreLTV           string `json:"preLTV"`
	AfterLTV         string `json:"afterLTV"`
	AdjustTime       int64  `json:"adjustTime,string"`
}

var FlexibleAdjustLtvHistoriesConfig = cex.ReqConfig[FlexibleAdjustLtvHistoriesParams, Page[[]FlexibleAdjustLtvHistory]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/ltv/adjustment/history",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]FlexibleAdjustLtvHistory]]),
}

type FlexibleLoanAssetsParams struct {
	LoanCoin string `s2m:"loanCoin,omitempty"`
}

type FlexibleLoanAsset struct {
	LoanCoin             string  `json:"loanCoin"`
	FlexibleInterestRate float64 `json:"flexibleInterestRate,string"`
	FlexibleMinLimit     float64 `json:"flexibleMinLimit,string"`
	FlexibleMaxLimit     float64 `json:"flexibleMaxLimit,string"`
}

var FlexibleLoanAssetsConfig = cex.ReqConfig[FlexibleLoanAssetsParams, Page[[]FlexibleLoanAsset]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/loanable/data",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]FlexibleLoanAsset]]),
}

type FlexibleCollateralCoinsParams struct {
	CollateralCoin string `s2m:"collateralCoin,omitempty"`
}

type FlexibleCollateralCoin struct {
	CollateralCoin string  `json:"collateralCoin"`
	InitialLTV     float64 `json:"initialLTV,string"`
	MarginCallLTV  float64 `json:"marginCallLTV,string"`
	LiquidationLTV float64 `json:"liquidationLTV,string"`
	MaxLimit       float64 `json:"maxLimit,string"`
}

var FlexibleCollateralCoinsConfig = cex.ReqConfig[FlexibleCollateralCoinsParams, Page[[]FlexibleCollateralCoin]]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             SapiV1 + "/loan/flexible/collateral/data",
		Method:           http.MethodGet,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[Page[[]FlexibleCollateralCoin]]),
}

// ---------------------------------------------
// Crypto Flexible Loans
// =============================================

// =============================================
// Spot Trading
// ---------------------------------------------

type NewSpotOrderParams struct {
	Symbol                  string                      `s2m:"symbol,omitempty"`
	Type                    OrderType                   `s2m:"type,omitempty"`
	Side                    OrderSide                   `s2m:"side,omitempty"`
	Quantity                float64                     `s2m:"quantity,omitempty"`
	Price                   float64                     `s2m:"price,omitempty"`
	TimeInForce             TimeInForce                 `s2m:"timeInForce,omitempty"`
	NewClientOrderId        string                      `s2m:"newClientOrderId,omitempty"`
	QuoteOrderQty           float64                     `s2m:"quoteOrderQty,omitempty"`
	StrategyId              int64                       `s2m:"strategyId,omitempty"`
	StrategyType            int64                       `s2m:"strategyType,omitempty"`
	StopPrice               float64                     `s2m:"stopPrice,omitempty"`               // Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	TrailingDelta           int64                       `s2m:"trailingDelta,omitempty"`           // Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders. For more details on SPOT implementation on trailing stops, please refer to Trailing Stop FAQ
	IcebergQty              float64                     `s2m:"icebergQty,omitempty"`              // Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
	NewOrderRespType        SpotOrderResponseType       `s2m:"newOrderRespType,omitempty"`        // Set the response JSON. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
	SelfTradePreventionMode SpotSelfTradePreventionMode `s2m:"selfTradePreventionMode,omitempty"` // The allowed enums is dependent on what is configured on the symbol.The possible supported values are EXPIRE_TAKER, EXPIRE_MAKER, EXPIRE_BOTH, NONE.
}

type NewSpotOrderFill struct {
	Price           float64 `json:"price,string"`
	Qty             float64 `json:"qty,string"`
	Commission      float64 `json:"commission,string"`
	CommissionAsset string  `json:"commissionAsset"`
	TradeId         int64   `json:"tradeId"`
}

type SpotOrderResult struct {
	// common
	Symbol                  string                      `json:"symbol"`
	OrderId                 int64                       `json:"orderId"`
	OrderListId             int64                       `json:"orderListId"` // Unless OCO, value will be -1
	ClientOrderId           string                      `json:"clientOrderId"`
	TransactTime            int64                       `json:"transactTime"`
	Price                   string                      `json:"price"`
	OrigQty                 string                      `json:"origQty"`
	ExecutedQty             string                      `json:"executedQty"`
	CummulativeQuoteQty     string                      `json:"cummulativeQuoteQty"`
	Status                  OrderStatus                 `json:"status"`
	TimeInForce             TimeInForce                 `json:"timeInForce"`
	Type                    OrderType                   `json:"type"`
	Side                    OrderSide                   `json:"side"`
	SelfTradePreventionMode SpotSelfTradePreventionMode `json:"selfTradePreventionMode"`
	Fills                   []NewSpotOrderFill          `json:"fills"`

	// new order result
	WorkingTime int64 `json:"workingTime"`

	// cancel order result
	OrigClientOrderId string `json:"origClientOrderId"`
}

var NewSpotOrderConfig = cex.ReqConfig[NewSpotOrderParams, SpotOrderResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/order",
		Method:           http.MethodPost,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[SpotOrderResult]),
}

type CancelSpotOrderParams struct {
	Symbol             string                     `s2m:"symbol,omitempty"`
	OrderId            int64                      `s2m:"orderId,omitempty"`
	OrigClientOrderId  string                     `s2m:"origClientOrderId,omitempty"`
	NewClientOrderId   string                     `s2m:"newClientOrderId,omitempty"`   // Used to uniquely identify this cancel. Automatically generated by default.
	CancelRestrictions SpotOrderCancelRestriction `s2m:"cancelRestrictions,omitempty"` // Supported values: ONLY_NEW - Cancel will succeed if the order status is NEW. ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED
}

var CancelSpotOrderConfig = cex.ReqConfig[CancelSpotOrderParams, SpotOrderResult]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:          ApiBaseUrl,
		Path:             ApiV3 + "/order",
		Method:           http.MethodDelete,
		IsUserData:       true,
		UserTimeInterval: 0,
		IpTimeInterval:   0,
	},
	HTTPStatusCodeChecker: HTTPStatusCodeChecker,
	RespBodyUnmarshaler:   bodyUnmshWrapper(cex.StdBodyUnmarshaler[SpotOrderResult]),
}

// ---------------------------------------------
// Spot Trading
// =============================================
