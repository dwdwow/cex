package bnc

const (
	ApiBaseUrl  = "https://api.binance.com"
	ApiV3       = "/api/v3"
	SapiV1      = "/sapi/v1"
	FapiBaseUrl = "https://fapi.binance.com"
	FapiV1      = "/fapi/v1"
	FapiV2      = "/fapi/v2"
)

const (
	SpotSymbolMid    = ""
	FuturesSymbolMid = ""
)

//const (
//	SpotMakerFeeTier    = 0.00075 * 0.8 // bnb and return fee 0.8
//	SpotTakerFeeTier    = 0.00075 * 0.8
//	FuturesMakerFeeTier = 0.00016 * 0.9
//	FuturesTakerFeeTier = 0.00016 * 0.9
//)

type BigBool string

const (
	BigTrue  BigBool = "TRUE"
	BigFalse BigBool = "FALSE"
)

type SmallBool string

const (
	SmallTrue  SmallBool = "true"
	SmallFalse SmallBool = "false"
)

type AccountType string

const (
	AccountTypeSpot AccountType = "SPOT"
)

type WalletType int

const (
	SpotWallet    = 0
	FundingWallet = 1
)

type TranType string

const (
	TranType_MAIN_UMFUTURE                 TranType = "MAIN_UMFUTURE"                 // Spot account transfer to USDⓈ-M Futures account
	TranType_MAIN_CMFUTURE                 TranType = "MAIN_CMFUTURE"                 // Spot account transfer to COIN-M Futures account
	TranType_MAIN_MARGIN                   TranType = "MAIN_MARGIN"                   // Spot account transfer to Margin（cross）account
	TranType_UMFUTURE_MAIN                 TranType = "UMFUTURE_MAIN"                 // USDⓈ-M Futures account transfer to Spot account
	TranType_UMFUTURE_MARGIN               TranType = "UMFUTURE_MARGIN"               // USDⓈ-M Futures account transfer to Margin（cross）account
	TranType_CMFUTURE_MAIN                 TranType = "CMFUTURE_MAIN"                 // COIN-M Futures account transfer to Spot account
	TranType_CMFUTURE_MARGIN               TranType = "CMFUTURE_MARGIN"               // COIN-M Futures account transfer to Margin(cross) account
	TranType_MARGIN_MAIN                   TranType = "MARGIN_MAIN"                   // Margin（cross）account transfer to Spot account
	TranType_MARGIN_UMFUTURE               TranType = "MARGIN_UMFUTURE"               // Margin（cross）account transfer to USDⓈ-M Futures
	TranType_MARGIN_CMFUTURE               TranType = "MARGIN_CMFUTURE"               // Margin（cross）account transfer to COIN-M Futures
	TranType_ISOLATEDMARGIN_MARGIN         TranType = "ISOLATEDMARGIN_MARGIN"         // Isolated margin account transfer to Margin(cross) account
	TranType_MARGIN_ISOLATEDMARGIN         TranType = "MARGIN_ISOLATEDMARGIN"         // Margin(cross) account transfer to Isolated margin account
	TranType_ISOLATEDMARGIN_ISOLATEDMARGIN TranType = "ISOLATEDMARGIN_ISOLATEDMARGIN" // Isolated margin account transfer to Isolated margin account
	TranType_MAIN_FUNDING                  TranType = "MAIN_FUNDING"                  // Spot account transfer to Funding account
	TranType_FUNDING_MAIN                  TranType = "FUNDING_MAIN"                  // Funding account transfer to Spot account
	TranType_FUNDING_UMFUTURE              TranType = "FUNDING_UMFUTURE"              // Funding account transfer to UMFUTURE account
	TranType_UMFUTURE_FUNDING              TranType = "UMFUTURE_FUNDING"              // UMFUTURE account transfer to Funding account
	TranType_MARGIN_FUNDING                TranType = "MARGIN_FUNDING"                // MARGIN account transfer to Funding account
	TranType_FUNDING_MARGIN                TranType = "FUNDING_MARGIN"                // Funding account transfer to Margin account
	TranType_FUNDING_CMFUTURE              TranType = "FUNDING_CMFUTURE"              // Funding account transfer to CMFUTURE account
	TranType_CMFUTURE_FUNDING              TranType = "CMFUTURE_FUNDING"              // CMFUTURE account transfer to Funding account
	TranType_MAIN_OPTION                   TranType = "MAIN_OPTION"                   // Spot account transfer to Options account
	TranType_OPTION_MAIN                   TranType = "OPTION_MAIN"                   // Options account transfer to Spot account
	TranType_UMFUTURE_OPTION               TranType = "UMFUTURE_OPTION"               // USDⓈ-M Futures account transfer to Options account
	TranType_OPTION_UMFUTURE               TranType = "OPTION_UMFUTURE"               // Options account transfer to USDⓈ-M Futures account
	TranType_MARGIN_OPTION                 TranType = "MARGIN_OPTION"                 // Margin（cross）account transfer to Options account
	TranType_OPTION_MARGIN                 TranType = "OPTION_MARGIN"                 // Options account transfer to Margin（cross）account
	TranType_FUNDING_OPTION                TranType = "FUNDING_OPTION"                // Funding account transfer to Options account
	TranType_OPTION_FUNDING                TranType = "OPTION_FUNDING"                // Options account transfer to Funding account
	TranType_MAIN_PORTFOLIO_MARGIN         TranType = "MAIN_PORTFOLIO_MARGIN"         // Spot account transfer to Portfolio Margin account
	TranType_PORTFOLIO_MARGIN_MAIN         TranType = "PORTFOLIO_MARGIN_MAIN"         // Portfolio Margin account transfer to Spot account
	TranType_MAIN_ISOLATED_MARGIN          TranType = "MAIN_ISOLATED_MARGIN"          // Spot account transfer to Isolated margin account
	TranType_ISOLATED_MARGIN_MAIN          TranType = "ISOLATED_MARGIN_MAIN"          // Isolated margin account transfer to Spot account
)

type TransferDirection int

const (
	TransferMainToCross TransferDirection = 1
	TransferCrossToMain TransferDirection = 2
)

type PairType string

const (
	PairTypeSpot   = "SPOT"
	PairTypeMargin = "MARGIN"
)

type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

type OrderType string

const (
	OrderTypeLimit           OrderType = "LIMIT"
	OrderTypeMarket          OrderType = "MARKET"
	OrderTypeStopLoss        OrderType = "STOP_LOSS"
	OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"
	OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"

	// just for futures
	OrderTypeStop               OrderType = "STOP"
	OrderTypeStopMarket         OrderType = "STOP_MARKET"
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket           = "TRAILING_STOP_MARKET"
)

type OrderStatus string

const (
	OrderStatusNew             OrderStatus = "NEW"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusPendingCancel   OrderStatus = "PENDING_CANCEL"
	OrderStatusRejected        OrderStatus = "REJECTED"
	OrderStatusExpired         OrderStatus = "EXPIRED"
)

type OrderExecutionType string

const (
	OrderExecutionTypeNew      OrderExecutionType = "NEW"
	OrderExecutionTypeCanceled OrderExecutionType = "CANCELED"
	OrderExecutionTypeReplaced OrderExecutionType = "REPLACED"
	OrderExecutionRejected     OrderExecutionType = "REJECTED"
	OrderExecutionTrade        OrderExecutionType = "TRADE"
	OrderExecutionExpired      OrderExecutionType = "EXPIRED"
)

type MarginOrderSideEffectType string

const (
	MarginOrderSideEffectTypeNoSideEffect MarginOrderSideEffectType = "NO_SIDE_EFFECT"
	MarginOrderSideEffectTypeMarginBuy    MarginOrderSideEffectType = "MARGIN_BUY"
	MarginOrderSideEffectTypeAutoRepay    MarginOrderSideEffectType = "AUTO_REPAY"
)

type FuturesPositionSide string

const (
	FuturesPositionSideBoth  = "BOTH"
	FuturesPositionSideLong  = "LONG"
	FuturesPositionSideShort = "SHORT"
)

type FuturesMarginType string

const (
	FuturesMarginTypeIsolated FuturesMarginType = "ISOLATED"
	FuturesMarginTypeCrossed  FuturesMarginType = "CROSSED"
)

type FuturesMarginLowerCaseType string

const (
	FuturesMarginLowerCaseIsolated FuturesMarginType = "isolated"
	FuturesMarginLowerCaseCross    FuturesMarginType = "cross"
)

type FuturesWorkingType string

const (
	FuturesWorkingTypeMarkPrice     FuturesWorkingType = "MARK_PRICE"
	FuturesWorkingTypeContractPrice FuturesWorkingType = "CONTRACT_PRICE"
)

type FuturesModifyMarginType int

const (
	FuturesAddMargin    FuturesModifyMarginType = 1
	FuturesReduceMargin FuturesModifyMarginType = 2
)

type FuturesMarginDeltaType string

const (
	FuturesMarginModifyTypeTrade      FuturesMarginDeltaType = "TRADE"
	FuturesMarginModifyTypeUserAdjust FuturesMarginDeltaType = "USER_ADJUST"
)

type FuturesIncomeType string

const (
	FuturesIncomeType_TRANSFER                    FuturesIncomeType = "TRANSFER"
	FuturesIncomeType_WELCOME_BONUS               FuturesIncomeType = "WELCOME_BONUS"
	FuturesIncomeType_REALIZED_PNL                FuturesIncomeType = "REALIZED_PNL"
	FuturesIncomeType_FUNDING_FEE                 FuturesIncomeType = "FUNDING_FEE"
	FuturesIncomeType_COMMISSION                  FuturesIncomeType = "COMMISSION"
	FuturesIncomeType_INSURANCE_CLEAR             FuturesIncomeType = "INSURANCE_CLEAR"
	FuturesIncomeType_REFERRAL_KICKBACK           FuturesIncomeType = "REFERRAL_KICKBACK"
	FuturesIncomeType_COMMISSION_REBATE           FuturesIncomeType = "COMMISSION_REBATE"
	FuturesIncomeType_API_REBATE                  FuturesIncomeType = "API_REBATE"
	FuturesIncomeType_CONTEST_REWARD              FuturesIncomeType = "CONTEST_REWARD"
	FuturesIncomeType_CROSS_COLLATERAL_TRANSFER   FuturesIncomeType = "CROSS_COLLATERAL_TRANSFER"
	FuturesIncomeType_OPTIONS_PREMIUM_FEE         FuturesIncomeType = "OPTIONS_PREMIUM_FEE"
	FuturesIncomeType_OPTIONS_SETTLE_PROFIT       FuturesIncomeType = "OPTIONS_SETTLE_PROFIT"
	FuturesIncomeType_INTERNAL_TRANSFER           FuturesIncomeType = "INTERNAL_TRANSFER"
	FuturesIncomeType_AUTO_EXCHANGE               FuturesIncomeType = "AUTO_EXCHANGE"
	FuturesIncomeType_DELIVERED_SETTELMENT        FuturesIncomeType = "DELIVERED_SETTELMENT"
	FuturesIncomeType_COIN_SWAP_DEPOSIT           FuturesIncomeType = "COIN_SWAP_DEPOSIT"
	FuturesIncomeType_COIN_SWAP_WITHDRAW          FuturesIncomeType = "COIN_SWAP_WITHDRAW"
	FuturesIncomeType_POSITION_LIMIT_INCREASE_FEE FuturesIncomeType = "POSITION_LIMIT_INCREASE_FEE"
)

type BSwapOrderStatus string

const (
	BSwapStatusProcess       BSwapOrderStatus = "PROCESS"
	BSwapStatusAcceptSuccess BSwapOrderStatus = "ACCEPT_SUCCESS"
	BSwapStatusSuccess       BSwapOrderStatus = "SUCCESS"
	BSwapStatusFail          BSwapOrderStatus = "FAIL"
)

type TimeInForce string

const (
	TimeInForceNone TimeInForce = ""
	TimeInForceGtc  TimeInForce = "GTC"
	TimeInForceIoc  TimeInForce = "IOC"
)

// OrderResponseType is the response of JSON type. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
type OrderResponseType string

const (
	SpotOrderResponseTypeAck    OrderResponseType = "ACK"
	SpotOrderResponseTypeResult OrderResponseType = "RESULT"
	SpotOrderResponseTypeFull   OrderResponseType = "FULL"
)

type SelfTradePreventionMode string

const (
	SelfTradePreventionModeExpireTaker SelfTradePreventionMode = "EXPIRE_TAKER"
	SelfTradePreventionModeExpireMaker SelfTradePreventionMode = "EXPIRE_MAKER"
	SelfTradePreventionModeExpireBoth  SelfTradePreventionMode = "EXPIRE_BOTH"
	SelfTradePreventionModeNone        SelfTradePreventionMode = "NONE"
)

type SpotOrderCancelRestriction string

const (
	SpotOrderCancelRestriction_ONLY_NEW              SpotOrderCancelRestriction = "ONLY_NEW"
	SpotOrderCancelRestriction_ONLY_PARTIALLY_FILLED SpotOrderCancelRestriction = "ONLY_PARTIALLY_FILLED"
)

type SpotCancelReplaceMode string

const (
	SpotCancelReplaceMode_STOP_ON_FAILURE SpotCancelReplaceMode = "STOP_ON_FAILURE"
	SpotCancelReplaceMode_ALLOW_FAILURE   SpotCancelReplaceMode = "ALLOW_FAILURE"
)

type SpotOrderCancelNewStatus string

const (
	SpotOrderCancelNewStatus_SUCCESS       SpotOrderCancelNewStatus = "SUCCESS"
	SpotOrderCancelNewStatus_FAILURE       SpotOrderCancelNewStatus = "FAILURE"
	SpotOrderCancelNewStatus_NOT_ATTEMPTED SpotOrderCancelNewStatus = "NOT_ATTEMPTED"
)

type Network string

const (
	NetworkNone Network = ""
	NetworkEth  Network = "ETH"
	NetworkTrx  Network = "TRX"
	NetworkArb1 Network = "ARBITRUM"
	NetworkErr  Network = "NoThisNetwork"
)

type DepositStatus int

const (
	DepositStatusNone           DepositStatus = -1
	DepositStatusPending        DepositStatus = 0
	DepositStatusCannotWithdraw DepositStatus = 6
	DepositStatusSuccess        DepositStatus = 1
)

type WithdrawStatus int

const (
	WithdrawStatusNone WithdrawStatus = iota - 1
	WithdrawStatusEmail
	WithdrawStatusCancelled
	WithdrawStatusAwaiting
	WithdrawStatusRejected
	WithdrawStatusProcessing
	WithdrawStatusFailure
	WithdrawStatusCompleted
)

type SubObErrorCode int

const (
	SubObErrorCon SubObErrorCode = iota + 1
	SubObErrorSend
	SubObErrorRead
	SubObErrorClose
)

type ExchangeStatus string

const (
	ExchangeTrading ExchangeStatus = "TRADING"
)

type LTVAdjustDirection string

const (
	LTVAdditional LTVAdjustDirection = "ADDITIONAL"
	LTVReduced    LTVAdjustDirection = "REDUCED"
)

type CryptoLoanIncomeType string

const (
	CryptoLoanIncomeBorrowIn                         CryptoLoanIncomeType = "borrowIn"
	CryptoLoanIncomeCollateralSpent                  CryptoLoanIncomeType = "collateralSpent"
	CryptoLoanIncomeRepayAmount                      CryptoLoanIncomeType = "repayAmount"
	CryptoLoanIncomeCollateralReturn                 CryptoLoanIncomeType = "collateralReturn"
	CryptoLoanIncomeAddCollateral                    CryptoLoanIncomeType = "addCollateral"
	CryptoLoanIncomeRemoveCollateral                 CryptoLoanIncomeType = "removeCollateral"
	CryptoLoanIncomeCollateralReturnAfterLiquidation CryptoLoanIncomeType = "collateralReturnAfterLiquidation"
)

type FlexibleBorrowStatus string

const (
	FlexibleBorrowSucceeds   FlexibleBorrowStatus = "Succeeds"
	FlexibleBorrowFailed     FlexibleBorrowStatus = "Failed"
	FlexibleBorrowProcessing FlexibleBorrowStatus = "Processing"
)

type FlexibleRepayStatus string

const (
	FlexibleRepayRepaid   FlexibleRepayStatus = "Repaid"
	FlexibleRepayRepaying FlexibleRepayStatus = "Repaying"
	FlexibleRepayFailed   FlexibleRepayStatus = "Failed"
)

type FlexibleRedeemDestType string

const (
	FlexibleRedeemDestSpot FlexibleRedeemDestType = "SPOT"
	FlexibleRedeemDestFund FlexibleRedeemDestType = "FUND"
)

type CodeMsg struct {
	// spot: 0, future: 200
	Code int `json:"code"`
	// spot: "", future: "success"
	Msg string `json:"msg"`
}

type Page[Slice any] struct {
	Rows  Slice `json:"rows"`
	Total int   `json:"total"`
}
