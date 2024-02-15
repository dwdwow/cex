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

type TransferType string

const (
	TransferTypeMainUmfuture                 TransferType = "MAIN_UMFUTURE"                 // Spot account transfer to USDⓈ-M Futures account
	TransferTypeMainCmfuture                 TransferType = "MAIN_CMFUTURE"                 // Spot account transfer to COIN-M Futures account
	TransferTypeMainMargin                   TransferType = "MAIN_MARGIN"                   // Spot account transfer to Margin（cross）account
	TransferTypeUmfutureMain                 TransferType = "UMFUTURE_MAIN"                 // USDⓈ-M Futures account transfer to Spot account
	TransferTypeUmfutureMargin               TransferType = "UMFUTURE_MARGIN"               // USDⓈ-M Futures account transfer to Margin（cross）account
	TransferTypeCmfutureMain                 TransferType = "CMFUTURE_MAIN"                 // COIN-M Futures account transfer to Spot account
	TransferTypeCmfutureMargin               TransferType = "CMFUTURE_MARGIN"               // COIN-M Futures account transfer to Margin(cross) account
	TransferTypeMarginMain                   TransferType = "MARGIN_MAIN"                   // Margin（cross）account transfer to Spot account
	TransferTypeMarginUmfuture               TransferType = "MARGIN_UMFUTURE"               // Margin（cross）account transfer to USDⓈ-M Futures
	TransferTypeMarginCmfuture               TransferType = "MARGIN_CMFUTURE"               // Margin（cross）account transfer to COIN-M Futures
	TransferTypeIsolatedmarginMargin         TransferType = "ISOLATEDMARGIN_MARGIN"         // Isolated margin account transfer to Margin(cross) account
	TransferTypeMarginIsolatedmargin         TransferType = "MARGIN_ISOLATEDMARGIN"         // Margin(cross) account transfer to Isolated margin account
	TransferTypeIsolatedmarginIsolatedmargin TransferType = "ISOLATEDMARGIN_ISOLATEDMARGIN" // Isolated margin account transfer to Isolated margin account
	TransferTypeMainFunding                  TransferType = "MAIN_FUNDING"                  // Spot account transfer to Funding account
	TransferTypeFundingMain                  TransferType = "FUNDING_MAIN"                  // Funding account transfer to Spot account
	TransferTypeFundingUmfuture              TransferType = "FUNDING_UMFUTURE"              // Funding account transfer to UMFUTURE account
	TransferTypeUmfutureFunding              TransferType = "UMFUTURE_FUNDING"              // UMFUTURE account transfer to Funding account
	TransferTypeMarginFunding                TransferType = "MARGIN_FUNDING"                // MARGIN account transfer to Funding account
	TransferTypeFundingMargin                TransferType = "FUNDING_MARGIN"                // Funding account transfer to Margin account
	TransferTypeFundingCmfuture              TransferType = "FUNDING_CMFUTURE"              // Funding account transfer to CMFUTURE account
	TransferTypeCmfutureFunding              TransferType = "CMFUTURE_FUNDING"              // CMFUTURE account transfer to Funding account
	TransferTypeMainOption                   TransferType = "MAIN_OPTION"                   // Spot account transfer to Options account
	TransferTypeOptionMain                   TransferType = "OPTION_MAIN"                   // Options account transfer to Spot account
	TransferTypeUmfutureOption               TransferType = "UMFUTURE_OPTION"               // USDⓈ-M Futures account transfer to Options account
	TransferTypeOptionUmfuture               TransferType = "OPTION_UMFUTURE"               // Options account transfer to USDⓈ-M Futures account
	TransferTypeMarginOption                 TransferType = "MARGIN_OPTION"                 // Margin（cross）account transfer to Options account
	TransferTypeOptionMargin                 TransferType = "OPTION_MARGIN"                 // Options account transfer to Margin（cross）account
	TransferTypeFundingOption                TransferType = "FUNDING_OPTION"                // Funding account transfer to Options account
	TransferTypeOptionFunding                TransferType = "OPTION_FUNDING"                // Options account transfer to Funding account
	TransferTypeMainPortfolioMargin          TransferType = "MAIN_PORTFOLIO_MARGIN"         // Spot account transfer to Portfolio Margin account
	TransferTypePortfolioMarginMain          TransferType = "PORTFOLIO_MARGIN_MAIN"         // Portfolio Margin account transfer to Spot account
	TransferTypeMainIsolatedMargin           TransferType = "MAIN_ISOLATED_MARGIN"          // Spot account transfer to Isolated margin account
	TransferTypeIsolatedMarginMain           TransferType = "ISOLATED_MARGIN_MAIN"          // Isolated margin account transfer to Spot account
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
	FuturesIncomeTypeTransfer                 FuturesIncomeType = "TRANSFER"
	FuturesIncomeTypeWelcomeBonus             FuturesIncomeType = "WELCOME_BONUS"
	FuturesIncomeTypeRealizedPnl              FuturesIncomeType = "REALIZED_PNL"
	FuturesIncomeTypeFundingFee               FuturesIncomeType = "FUNDING_FEE"
	FuturesIncomeTypeCommission               FuturesIncomeType = "COMMISSION"
	FuturesIncomeTypeInsuranceClear           FuturesIncomeType = "INSURANCE_CLEAR"
	FuturesIncomeTypeReferralKickback         FuturesIncomeType = "REFERRAL_KICKBACK"
	FuturesIncomeTypeCommissionRebate         FuturesIncomeType = "COMMISSION_REBATE"
	FuturesIncomeTypeApiRebate                FuturesIncomeType = "API_REBATE"
	FuturesIncomeTypeContestReward            FuturesIncomeType = "CONTEST_REWARD"
	FuturesIncomeTypeCrossCollateralTransfer  FuturesIncomeType = "CROSS_COLLATERAL_TRANSFER"
	FuturesIncomeTypeOptionsPremiumFee        FuturesIncomeType = "OPTIONS_PREMIUM_FEE"
	FuturesIncomeTypeOptionsSettleProfit      FuturesIncomeType = "OPTIONS_SETTLE_PROFIT"
	FuturesIncomeTypeInternalTransfer         FuturesIncomeType = "INTERNAL_TRANSFER"
	FuturesIncomeTypeAutoExchange             FuturesIncomeType = "AUTO_EXCHANGE"
	FuturesIncomeTypeDeliveredSettelment      FuturesIncomeType = "DELIVERED_SETTELMENT"
	FuturesIncomeTypeCoinSwapDeposit          FuturesIncomeType = "COIN_SWAP_DEPOSIT"
	FuturesIncomeTypeCoinSwapWithdraw         FuturesIncomeType = "COIN_SWAP_WITHDRAW"
	FuturesIncomeTypePositionLimitIncreaseFee FuturesIncomeType = "POSITION_LIMIT_INCREASE_FEE"
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
	OrderResponseTypeAck    OrderResponseType = "ACK"
	OrderResponseTypeResult OrderResponseType = "RESULT"
	// OrderResponseTypeFull is only for spot new order response
	OrderResponseTypeFull OrderResponseType = "FULL"
)

type SelfTradePreventionMode string

const (
	SelfTradePreventionModeExpireTaker SelfTradePreventionMode = "EXPIRE_TAKER"
	SelfTradePreventionModeExpireMaker SelfTradePreventionMode = "EXPIRE_MAKER"
	SelfTradePreventionModeExpireBoth  SelfTradePreventionMode = "EXPIRE_BOTH"
	SelfTradePreventionModeNone        SelfTradePreventionMode = "NONE"
)

type OrderCancelRestriction string

const (
	OrderCancelRestrictionOnlyNew             OrderCancelRestriction = "ONLY_NEW"
	OrderCancelRestrictionOnlyPartiallyFilled OrderCancelRestriction = "ONLY_PARTIALLY_FILLED"
)

type SpotCancelReplaceMode string

const (
	SpotCancelReplaceModeStopOnFailure SpotCancelReplaceMode = "STOP_ON_FAILURE"
	SpotCancelReplaceModeAllowFailure  SpotCancelReplaceMode = "ALLOW_FAILURE"
)

type SpotOrderCancelNewStatus string

const (
	SpotOrderCancelNewStatusSuccess      SpotOrderCancelNewStatus = "SUCCESS"
	SpotOrderCancelNewStatusFailure      SpotOrderCancelNewStatus = "FAILURE"
	SpotOrderCancelNewStatusNotAttempted SpotOrderCancelNewStatus = "NOT_ATTEMPTED"
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

type CryptoLoanFlexibleBorrowStatus string

const (
	CryptoLoanFlexibleBorrowSucceeds   CryptoLoanFlexibleBorrowStatus = "Succeeds"
	CryptoLoanFlexibleBorrowFailed     CryptoLoanFlexibleBorrowStatus = "Failed"
	CryptoLoanFlexibleBorrowProcessing CryptoLoanFlexibleBorrowStatus = "Processing"
)

type CryptoFlexibleRepayStatus string

const (
	CryptoLoanFlexibleRepayRepaid   CryptoFlexibleRepayStatus = "Repaid"
	CryptoLoanFlexibleRepayRepaying CryptoFlexibleRepayStatus = "Repaying"
	CryptoLoanFlexibleRepayFailed   CryptoFlexibleRepayStatus = "Failed"
)

type SimpleEarnFlexibleRedeemDestination string

const (
	SimpleEarnFlexibleRedeemDestinationSpot SimpleEarnFlexibleRedeemDestination = "SPOT"
	SimpleEarnFlexibleRedeemDestinationFund SimpleEarnFlexibleRedeemDestination = "FUND"
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
