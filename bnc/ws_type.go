package bnc

const (
	WsBaseUrl       = "wss://stream.binance.com:9443/ws"
	FutureWsBaseUrl = "wss://fstream.binance.com/ws"
)

type WsMethod string

const (
	WsMethodSub     WsMethod = "SUBSCRIBE"
	WsMethodUnsub   WsMethod = "UNSUBSCRIBE"
	WsMethodRequest WsMethod = "REQUEST"
)

type WsEvent string

const (
	WsEDepthUpdate                  WsEvent = "depthUpdate"
	WsTrade                         WsEvent = "trade"
	WsAggTrade                      WsEvent = "aggTrade"
	WsMarginCall                    WsEvent = "MARGIN_CALL"
	WsAccountUpdate                 WsEvent = "ACCOUNT_UPDATE"
	WsOrderTradeUpdate              WsEvent = "ORDER_TRADE_UPDATE"
	WsAccountConfigUpdate           WsEvent = "ACCOUNT_CONFIG_UPDATE"
	WsStrategyUpdate                WsEvent = "STRATEGY_UPDATE"
	WsGridUpdate                    WsEvent = "GRID_UPDATE"
	WsConditionalOrderTriggerReject WsEvent = "CONDITIONAL_ORDER_TRIGGER_REJECT"

	WsEventOutboundAccountPosition WsEvent = "outboundAccountPosition"
	WsEventBalanceUpdate           WsEvent = "balanceUpdate"
	WsEventExecutionReport         WsEvent = "executionReport"
	WsEventListStatus              WsEvent = "listStatus"
	WsEventListenKeyExpired        WsEvent = "listenKeyExpired"
)

type WsSubMsg struct {
	Method WsMethod `json:"method"`
	Params []string `json:"params"`
	Id     int64    `json:"id"`
}

type WsReqMsg struct {
	Method WsMethod `json:"method"`
	Params []any    `json:"params"`
	Id     int64    `json:"id"`
}

type WsRespMsg[R any] struct {
	Result R     `json:"result"`
	Id     int64 `json:"id"`
}

type WsDepthMsg struct {
	EventType WsEvent    `json:"e"`
	EventTime int64      `json:"E"`
	Symbol    string     `json:"s"`
	FirstId   int64      `json:"U"`
	LastId    int64      `json:"u"`
	Bids      [][]string `json:"b"`
	Asks      [][]string `json:"a"`

	// just for future ob
	TxTime  int64 `json:"T"`
	PLastId int64 `json:"pu"`
}

type WsTradeStream struct {
	EventType     WsEvent `json:"e"`
	EventTime     int64   `json:"E"`
	Symbol        string  `json:"s"`
	TradeID       int64   `json:"t"`
	Price         float64 `json:"p,string"`
	Quantity      float64 `json:"q,string"`
	BuyerOrderID  int64   `json:"b"`
	SellerOrderID int64   `json:"a"`
	TradeTime     int64   `json:"T"`
	IsBuyerMaker  bool    `json:"m"`
}

type WsFuAggTradeStream struct {
	EventType    WsEvent `json:"e"`
	EventTime    int64   `json:"E"`
	Symbol       string  `json:"s"`
	AggID        int64   `json:"a"`
	Price        float64 `json:"p,string"`
	Quantity     float64 `json:"q,string"`
	FirstTradeId int64   `json:"f"`
	LastTradeId  int64   `json:"l"`
	TradeTime    int64   `json:"T"`
	IsBuyerMaker bool    `json:"m"`
}
