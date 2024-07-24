package bnc

// WsCMIndexPriceStream
// url: wss://dstream.binance.com/ws
// event: WsEventIndexPriceUpdate
type WsCMIndexPriceStream struct {
	EventType WsEvent `json:"e"`
	EventTime int64   `json:"E"`
	Pair      string  `json:"i"`
	Price     float64 `json:"p,string"`
}

// WsMarkPriceStream
// url: wss://dstream.binance.com/ws
// event: WsEventMarkPriceUpdate
type WsMarkPriceStream struct {
	EventType            WsEvent `json:"e"`
	EventTime            int64   `json:"E"`
	Symbol               string  `json:"s"`
	MarkPrice            float64 `json:"p,string"`
	EstimatedSettlePrice float64 `json:"P,string"`
	IndexPrice           float64 `json:"i,string"`
	FundingRate          float64 `json:"r,string"`
	NextFundingTime      int64   `json:"T,string"`
}

type WsStrategyOrder struct {
	Symbol               string              `json:"s"`
	ClientOrderId        string              `json:"c"`
	Id                   int64               `json:"si"`
	Side                 OrderSide           `json:"S"`
	Type                 string              `json:"st"`
	TimeInForce          TimeInForce         `json:"f"`
	Qty                  float64             `json:"q,string"`
	Price                float64             `json:"p,string"`
	StopPrice            float64             `json:"sp,string"`
	Status               OrderStatus         `json:"os"`
	OrderBookTime        int64               `json:"T"`
	UpdateTime           int64               `json:"ut"`
	IsThisReduceOnly     bool                `json:"R"`
	StopPriceWorkingType string              `json:"wt"`
	PositionSide         FuturesPositionSide `json:"ps"`
	Cp                   bool                `json:"cp"` // If Close-All, pushed with conditional order
	ActivationPrice      float64             `json:"AP,string"`
	CallbackRate         float64             `json:"cr,string"`
	OrderId              int64               `json:"i"`
	STPMode              string              `json:"V"`
	Gtd                  int64               `json:"gtd"`
}

type WsBusinessUnit string

const (
	WsBusinessUnitUM WsBusinessUnit = "UM"
	WsBusinessUnitCM WsBusinessUnit = "CM"
)

// WsConditionalOrderTradeUpdate
// url: https://papi.binance.com/ws
// event: WsEventConditionalOrderTradeUpdate
type WsConditionalOrderTradeUpdate struct {
	EventType       WsEvent         `json:"e"`
	EventTime       int64           `json:"E"`
	TransactionTime int64           `json:"T"`
	BusinessUnit    WsBusinessUnit  `json:"fs"`
	StrategyOrder   WsStrategyOrder `json:"so"`
}

type WsOpenOrderLossUpdate struct {
	Asset      string  `json:"a"`
	SignAmount float64 `json:"o"`
}

// WsOpenOrderLoss
// url: https://papi.binance.com/ws
// event: WsEventOpenOrderLoss
type WsOpenOrderLoss struct {
	EventType WsEvent                 `json:"e"`
	EventTime int64                   `json:"E"`
	OrderLoss []WsOpenOrderLossUpdate `json:"O"`
}

// WsMarginAccountUpdate
// url: https://papi.binance.com/ws
// event: WsEventOutboundAccountPosition
type WsMarginAccountUpdate WsSpotAccountUpdate

type LiabilityType string

const (
	LiabilityTypeBorrow LiabilityType = "BORROW"
)

// WsLiabilityChange
// url: https://papi.binance.com/ws
// event: WsEventLiabilityChange
type WsLiabilityChange struct {
	EventType      WsEvent       `json:"e"`
	EventTime      int64         `json:"E"`
	Asset          string        `json:"a"`
	Type           LiabilityType `json:"t"`
	TxId           int64         `json:"tx"`
	Principal      float64       `json:"p,string"`
	Interest       float64       `json:"i,string"`
	TotalLiability float64       `json:"l,string"`
}

// WsPmFuturesOrderTradeUpdate
// url: https://papi.binance.com/ws
// event: WsEventOrderTradeUpdate
type WsPmFuturesOrderTradeUpdate struct {
	EventType       WsEvent                `json:"e"`
	EventTime       int64                  `json:"E"`
	TransactionTime int64                  `json:"T"`
	BusinessUnit    WsBusinessUnit         `json:"fs"`
	AccountAlias    string                 `json:"i"` // Account Alias,ignore for UM
	Order           WsOrderExecutionReport `json:"o"`
}

type WsFuturesWalletBalance struct {
	Asset              string  `json:"a"`
	WalletBalance      float64 `json:"wb,string"`
	CrossWalletBalance float64 `json:"cw,string"`
	BalanceChange      float64 `json:"bc,string"`
}

type WsFuturesPosition struct {
	Symbol                    string              `json:"s"`
	SignPositionQty           float64             `json:"pa"`
	EntryPrice                float64             `json:"ep"`
	PreFeeAccumulatedRealized float64             `json:"cr"`
	UnrealizedPnl             float64             `json:"up"`
	PositionSide              FuturesPositionSide `json:"ps"`
	BreakevenPrice            string              `json:"bep"`
}

type WsPmFuturesAcctUpdateData struct {
	Reason    string                   `json:"m"`
	Balances  []WsFuturesWalletBalance `json:"B"`
	Positions []WsFuturesPosition      `json:"P"`
}

// WsPmFuturesAcctUpdateStream
// url: https://papi.binance.com/ws
// event: WsEventAccountUpdate
type WsPmFuturesAcctUpdateStream struct {
	EventType       WsEvent                   `json:"e"`
	EventTime       int64                     `json:"E"`
	TransactionTime int64                     `json:"T"`
	BusinessUnit    WsBusinessUnit            `json:"fs"`
	AccountAlias    string                    `json:"i"` // Account Alias,ignore for UM
	Data            WsPmFuturesAcctUpdateData `json:"a"`
}

type WsFuturesCfg struct {
	Symbol   string  `json:"s"`
	Leverage float64 `json:"l"`
}

// WsFuturesAccountCfgUpdate
// url: https://papi.binance.com/ws
// event: WsEventAccountConfigUpdate
type WsFuturesAccountCfgUpdate struct {
	EventType       WsEvent        `json:"e"`
	EventTime       int64          `json:"E"`
	TransactionTime int64          `json:"T"`
	BusinessUnit    WsBusinessUnit `json:"fs"`
	Cfg             WsFuturesCfg   `json:"ac"`
}

type WsPmRiskLevelChangeType string

const (
	WsPmRiskLevelChangeTypeMarginCall       WsPmRiskLevelChangeType = "MARGIN_CALL"
	WsPmRiskLevelChangeTypeSupplyMargin     WsPmRiskLevelChangeType = "SUPPLY_MARGIN"
	WsPmRiskLevelChangeTypeReduceOnly       WsPmRiskLevelChangeType = "REDUCE_ONLY"
	WsPmRiskLevelChangeTypeForceLiquidation WsPmRiskLevelChangeType = "FORCE_LIQUIDATION"
)

// WsPmRiskLevelChange
// url: https://papi.binance.com/ws
// event: WsEventRiskLevelChange
type WsPmRiskLevelChange struct {
	EventType                             WsEvent                 `json:"e"`
	EventTime                             int64                   `json:"E"`
	UniMMR                                string                  `json:"u"`
	ChangeType                            WsPmRiskLevelChangeType `json:"s"`
	AccountEquityUSD                      float64                 `json:"eq,string"`
	AccountEquityUSDWithoutCollateralRate float64                 `json:"ae,string"`
	MaintenanceMarginUSD                  float64                 `json:"m,string"`
}

// WsMarginBalanceUpdate
// url: https://papi.binance.com/ws
// event: balanceUpdate
type WsMarginBalanceUpdate struct {
	Asset        string  `json:"a"`
	BalanceDelta float64 `json:"d,string"`
	UpdateId     int64   `json:"U"`
	ClearTime    int64   `json:"T"`
}

type WsLiquidationOrder struct {
	Symbol        string      `json:"s"`
	OrderSide     OrderSide   `json:"S"`
	OrderType     OrderType   `json:"o"`
	TimeInForce   TimeInForce `json:"f"`
	OriQty        float64     `json:"q,string"`
	Price         float64     `json:"p,string"`
	AvgPrice      float64     `json:"ap,string"`
	OrderStatus   OrderStatus `json:"X"`
	LastFilledQty float64     `json:"l,string"`
	FilledQty     float64     `json:"z,string"`
	TradeTime     int64       `json:"T"`
}

type WsLiquidationOrderStream struct {
	EventType WsEvent            `json:"e"`
	EventTime int64              `json:"E"`
	Order     WsLiquidationOrder `json:"o"`
}
