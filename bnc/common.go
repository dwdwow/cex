package bnc

const (
	ApiBaseUrl = "https://api.binance.com"
	ApiV3      = "/api/v3"
	SapiV1     = "/sapi/v1"

	FapiBaseUrl = "https://fapi.binance.com"
	FapiV1      = "/fapi/v1"
	FapiV2      = "/fapi/v2"
)

const (
	SpotMakerFeeTier   = 0.00075 * 0.8 // bnb and return fee 0.8
	SpotTakerFeeTier   = 0.00075 * 0.8
	FutureMakerFeeTier = 0.00016 * 0.9
	FutureTakerFeeTier = 0.00016 * 0.9
)

const (
	SpotSymbolMid   = ""
	FutureSymbolMid = ""
)

type BigBool string

const (
	TRUE  BigBool = "TRUE"
	FALSE BigBool = "FALSE"
)

type SmallBool string

const (
	SmallTrue  SmallBool = "true"
	SmallFalse SmallBool = "false"
)

type AcctType string

const (
	AcctSpot AcctType = "SPOT"
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

type TransferDire int

const (
	TransferDireMainToCross TransferDire = 1
	TransferDireCrossToMain TransferDire = 2
)

type TradeType string

const (
	TradeSpot   = "SPOT"
	TradeMargin = "MARGIN"
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

	// future
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
	OrderExeNew      OrderExecutionType = "NEW"
	OrderExeCanceled OrderExecutionType = "CANCELED"
	OrderExeReplaced OrderExecutionType = "REPLACED"
	OrderExeRejected OrderExecutionType = "REJECTED"
	OrderExeTrade    OrderExecutionType = "TRADE"
	OrderExeExpired  OrderExecutionType = "EXPIRED"
)

type FuPositionSide string

const (
	FuPosBoth  = "BOTH"
	FuPosLong  = "LONG"
	FuPosShort = "SHORT"
)

type MarginOrderSideEffectType string

const (
	SideEffectTypeNoSideEffect MarginOrderSideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    MarginOrderSideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    MarginOrderSideEffectType = "AUTO_REPAY"
)

type FuMarginType string

const (
	FuMarginIsolated FuMarginType = "isolated"
	FuMarginCross    FuMarginType = "cross"
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

// SpotOrderResponseType is the response of JSON type. ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
type SpotOrderResponseType string

const (
	SpotOrderResponse_ACK    SpotOrderResponseType = "ACK"
	SpotOrderResponse_RESULT SpotOrderResponseType = "RESULT"
	SpotOrderResponse_FULL   SpotOrderResponseType = "FULL"
)

type SpotSelfTradePreventionMode string

const (
	SpotSelfTradePreventionMode_EXPIRE_TAKER SpotSelfTradePreventionMode = "EXPIRE_TAKER"
	SpotSelfTradePreventionMode_EXPIRE_MAKER SpotSelfTradePreventionMode = "EXPIRE_MAKER"
	SpotSelfTradePreventionMode_EXPIRE_BOTH  SpotSelfTradePreventionMode = "EXPIRE_BOTH"
	SpotSelfTradePreventionMode_NONE         SpotSelfTradePreventionMode = "NONE"
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
	DepoStatusNone           DepositStatus = -1
	DepoStatusPending        DepositStatus = 0
	DepoStatusCannotWithdraw DepositStatus = 6
	DepoStatusSuccess        DepositStatus = 1
)

type WithdrawStatus int

const (
	WidrStatusNone WithdrawStatus = iota - 1
	WidrStatusEmail
	WidrStatusCancelled
	WidrStatusAwaiting
	WidrStatusRejected
	WidrStatusProcessing
	WidrStatusFailure
	WidrStatusCompleted
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
	LTVAdDireAdditional LTVAdjustDirection = "ADDITIONAL"
	LTVAdDireReduced    LTVAdjustDirection = "REDUCED"
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
