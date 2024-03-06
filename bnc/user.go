package bnc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/s2m"
	"github.com/go-resty/resty/v2"
)

type UserConfig struct {
	fuPosSide FuturesPositionSide
}

type User struct {
	api cex.Api
	cfg UserConfig
}

type UserOpt func(*User)

func UserOptPositionSide(side FuturesPositionSide) func(*User) {
	return func(user *User) {
		user.cfg.fuPosSide = side
	}
}

func NewUser(apiKey, secretKey string, opts ...UserOpt) *User {
	user := &User{
		api: cex.Api{Cex: cex.BINANCE, ApiKey: apiKey, SecretKey: secretKey},
		cfg: UserConfig{},
	}
	for _, opt := range opts {
		opt(user)
	}
	return user
}

var emptyUser = &User{}

func EmptyUser() *User {
	return emptyUser
}

// ============================================================
// User Getter
// ------------------------------------------------------------

func (u *User) Api() cex.Api {
	return u.api
}

func (u *User) Config() UserConfig {
	return u.cfg
}

// ------------------------------------------------------------
// User Getter
// ============================================================

// ============================================================
// Account API
// ------------------------------------------------------------

func (u *User) Coins(opts ...cex.CltOpt) (*resty.Response, []Coin, cex.RequestError) {
	return cex.Request(u, CoinInfoConfig, nil, opts...)
}

func (u *User) SpotAccount(opts ...cex.CltOpt) (*resty.Response, SpotAccount, cex.RequestError) {
	return cex.Request(u, SpotAccountConfig, nil, opts...)
}

func (u *User) Transfer(tranType TransferType, asset string, amount float64, opts ...cex.CltOpt) (*resty.Response, UniversalTransferResp, cex.RequestError) {
	return cex.Request(u, UniversalTransferConfig, UniversalTransferParams{Type: tranType, Asset: asset, Amount: amount}, opts...)
}

func (u *User) FuturesAccount(opts ...cex.CltOpt) (*resty.Response, FuturesAccount, cex.RequestError) {
	return cex.Request(u, FuturesAccountConfig, nil, opts...)
}

func (u *User) FuturesPositions(symbol string, opts ...cex.CltOpt) (*resty.Response, []FuturesPosition, cex.RequestError) {
	return cex.Request(u, FuturesPositionsConfig, FuturesPositionsParams{Symbol: symbol}, opts...)
}

// ------------------------------------------------------------
// Account API
// ============================================================

// ============================================================
// Flexible Simple Earn API
// ------------------------------------------------------------

func (u *User) SimpleEarnFlexibleProducts(asset string, opts ...cex.CltOpt) (*resty.Response, Page[[]SimpleEarnFlexibleProduct], cex.RequestError) {
	return cex.Request(u, SimpleEarnFlexibleProductConfig, SimpleEarnFlexibleProductListParams{Asset: asset, Size: 100}, opts...)
}

func (u *User) SimpleEarnFlexiblePositions(asset, productId string, opts ...cex.CltOpt) (*resty.Response, Page[[]SimpleEarnFlexiblePosition], cex.RequestError) {
	return cex.Request(u, SimpleEarnFlexiblePositionsConfig, SimpleEarnFlexiblePositionsParams{Asset: asset, ProductId: productId, Size: 100}, opts...)
}

func (u *User) SimpleEarnFlexibleRedeem(productId string, redeemAll bool, amount float64, destAccount SimpleEarnFlexibleRedeemDestination, opts ...cex.CltOpt) (*resty.Response, SimpleEarnFlexibleRedeemResponse, cex.RequestError) {
	return cex.Request(u, SimpleEarnFlexibleRedeemConfig, SimpleEarnFlexibleRedeemParams{ProductId: productId, RedeemAll: redeemAll, Amount: amount, DestAccount: destAccount}, opts...)
}

// ------------------------------------------------------------
// Flexible Simple Earn API
// ============================================================

// ============================================================
// Flexible Loan API
// ------------------------------------------------------------

func (u *User) CryptoLoanFlexibleOngoingOrders(loanCoin, collateralCoin string, opts ...cex.CltOpt) (*resty.Response, Page[[]CryptoLoanFlexibleOngoingOrder], cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleOngoingOrdersConfig, CryptoLoanFlexibleOngoingOrdersParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, Limit: 100}, opts...)
}

func (u *User) CryptoLoanIncomeHistories(asset string, incomeType CryptoLoanIncomeType, opts ...cex.CltOpt) (*resty.Response, []CryptoLoanIncomeHistory, cex.RequestError) {
	return cex.Request(u, CryptoLoansIncomeHistoriesConfig, CryptoLoansIncomeHistoriesParams{Asset: asset, Type: incomeType, Limit: 100}, opts...)
}

func (u *User) CryptoLoanFlexibleBorrow(loanCoin string, collateralCoin string, loanAmount, collateralAmount float64, opts ...cex.CltOpt) (*resty.Response, CryptoLoanFlexibleBorrowResult, cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleBorrowConfig, CryptoLoanFlexibleBorrowParams{LoanCoin: loanCoin, LoanAmount: loanAmount, CollateralCoin: collateralCoin, CollateralAmount: collateralAmount}, opts...)
}

func (u *User) CryptoLoanFlexibleBorrowHistories(loanCoin, collateralCoin string, opts ...cex.CltOpt) (*resty.Response, Page[[]CryptoLoanFlexibleBorrowHistory], cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleBorrowHistoriesConfig, CryptoLoanFlexibleBorrowHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, Limit: 100}, opts...)
}

func (u *User) CryptoLoanFlexibleRepay(loanCoin, collateralCoin string, repayAmount float64, collateralReturn, fullRepayment BigBool, opts ...cex.CltOpt) (*resty.Response, CryptoLoanFlexibleRepayResult, cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleRepayConfig, CryptoLoanFlexibleRepayParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, RepayAmount: repayAmount, CollateralReturn: collateralReturn, FullRepayment: fullRepayment}, opts...)
}

func (u *User) CryptoLoanFlexibleRepaymentHistories(loanCoin, collateralCoin string, opts ...cex.CltOpt) (*resty.Response, Page[[]CryptoLoanFlexibleRepaymentHistory], cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleRepaymentHistoriesConfig, CryptoLoanFlexibleRepaymentHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, Limit: 100}, opts...)
}

func (u *User) CryptoLoanFlexibleAdjustLtv(loanCoin, collateralCoin string, adjustmentAmount float64, direction LTVAdjustDirection, opts ...cex.CltOpt) (*resty.Response, CryptoLoanFlexibleLoanAdjustLtvResult, cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleLoanAdjustLtvConfig, CryptoLoanFlexibleAdjustLtvParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, AdjustmentAmount: adjustmentAmount, Direction: direction}, opts...)
}

func (u *User) CryptoLoanFlexibleAdjustLtvHistories(loanCoin, collateralCoin string, opts ...cex.CltOpt) (*resty.Response, Page[[]CryptoLoanFlexibleAdjustLtvHistory], cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleAdjustLtvHistoriesConfig, CryptoLoanFlexibleAdjustLtvHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, Limit: 100}, opts...)
}

func (u *User) CryptoLoanFlexibleLoanAssets(loanCoin string, opts ...cex.CltOpt) (*resty.Response, Page[[]CryptoLoanFlexibleLoanAsset], cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleLoanAssetsConfig, CryptoLoanFlexibleLoanAssetsParams{loanCoin}, opts...)
}

func (u *User) CryptoLoanFlexibleCollateralAssets(collateralCoin string, opts ...cex.CltOpt) (*resty.Response, Page[[]CryptoLoanFlexibleCollateralCoin], cex.RequestError) {
	return cex.Request(u, CryptoLoanFlexibleCollateralCoinsConfig, CryptoLoanFlexibleCollateralCoinsParams{collateralCoin}, opts...)
}

// ------------------------------------------------------------
// Flexible Loan API
// ============================================================

// ============================================================
// cex.Trader Interface Implementations
// ------------------------------------------------------------

func (u *User) QueryOrder(order *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	return u.queryOrd(order, opts...)
}

func (u *User) CancelOrder(order *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	return u.cancelOrd(order, opts...)
}

func (u *User) WaitOrder(ctx context.Context, order *cex.Order, opts ...cex.CltOpt) chan cex.RequestError {
	return u.waitOrd(ctx, order, opts...)
}

func (u *User) NewSpotOrder(asset, quote string, tradeType cex.OrderType, orderSide cex.OrderSide, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.newSpotOrd(asset, quote, tradeType, orderSide, qty, price, opts...)
}

func (u *User) NewSpotLimitBuyOrder(asset, quote string, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideBuy, qty, price, opts...)
}

func (u *User) NewSpotLimitSellOrder(asset, quote string, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideSell, qty, price, opts...)
}

func (u *User) NewSpotMarketBuyOrder(asset, quote string, qty float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideBuy, qty, 0, opts...)
}

func (u *User) NewSpotMarketSellOrder(asset, quote string, qty float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideSell, qty, 0, opts...)
}

func (u *User) NewFuturesOrder(asset, quote string, tradeType cex.OrderType, orderSide cex.OrderSide, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.newFuOrd(asset, quote, tradeType, orderSide, qty, price, opts...)
}

func (u *User) NewFuturesLimitBuyOrder(asset, quote string, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideBuy, qty, price, opts...)
}

func (u *User) NewFuturesLimitSellOrder(asset, quote string, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideSell, qty, price, opts...)
}

func (u *User) NewFuturesMarketBuyOrder(asset, quote string, qty float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideBuy, qty, 0, opts...)
}

func (u *User) NewFuturesMarketSellOrder(asset, quote string, qty float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideSell, qty, 0, opts...)
}

// ------------------------------------------------------------
// cex.Trader Interface Implementations
// ============================================================

// ============================================================
// Spot API
// ------------------------------------------------------------

func (u *User) CancelSpotOrder(symbol string, orderId int64, cltOrdId string, opts ...cex.CltOpt) (*resty.Response, SpotOrder, cex.RequestError) {
	return cex.Request(u, SpotCancelOrderConfig, SpotCancelOrderParams{Symbol: symbol, OrderId: orderId, OrigClientOrderId: cltOrdId}, opts...)
}

func (u *User) QuerySpotOrder(symbol string, orderId int64, cltOrdId string, opts ...cex.CltOpt) (*resty.Response, SpotOrder, cex.RequestError) {
	return cex.Request(u, SpotQueryOrderConfig, SpotQueryOrderParams{Symbol: symbol, OrderId: orderId, OrigClientOrderId: cltOrdId}, opts...)
}

// ------------------------------------------------------------
// Spot API
// ============================================================

// ============================================================
// Futures API
// ------------------------------------------------------------

func (u *User) CancelFuturesOrder(symbol string, orderId int64, cltOrdId string, opts ...cex.CltOpt) (*resty.Response, FuturesOrder, cex.RequestError) {
	return cex.Request(u, FuturesCancelOrderConfig, FuturesQueryOrCancelOrderParams{Symbol: symbol, OrderId: orderId, OrigClientOrderId: cltOrdId}, opts...)
}

func (u *User) QueryFuturesOrder(symbol string, orderId int64, cltOrdId string, opts ...cex.CltOpt) (*resty.Response, FuturesOrder, cex.RequestError) {
	return cex.Request(u, FuturesQueryOrderConfig, FuturesQueryOrCancelOrderParams{Symbol: symbol, OrderId: orderId, OrigClientOrderId: cltOrdId}, opts...)
}

func (u *User) CloseFuturesOrder(symbol string, side OrderSide, opts ...cex.CltOpt) (*resty.Response, FuturesOrder, cex.RequestError) {
	return cex.Request(u, FuturesNewOrderConfig, FuturesNewOrderParams{Symbol: symbol, PositionSide: u.cfg.fuPosSide, Side: side, ReduceOnly: SmallTrue}, opts...)
}

// ------------------------------------------------------------
// Futures API
// ============================================================

// ============================================================
// Private Trade Functions
// ------------------------------------------------------------

func (u *User) newSpotOrd(asset, quote string, orderType cex.OrderType, orderSide cex.OrderSide, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	symbol := asset + quote
	var tif TimeInForce
	if orderType == cex.OrderTypeLimit {
		tif = TimeInForceGtc
	}
	resp, rawOrd, err := cex.Request(u, SpotNewOrderConfig, SpotNewOrderParams{
		Symbol:      symbol,
		Type:        mapStrStr(orderType, ordTypByCexOrdTyp),
		Side:        mapStrStr(orderSide, ordSideByCexOrdSide),
		Quantity:    qty,
		Price:       price,
		TimeInForce: tif,
	}, opts...)
	ord := SwitchSpotOrderToCexOrder(rawOrd)
	ord.ApiKey = u.api.ApiKey
	return resp, &ord, err
}

func (u *User) cancelSpotOrd(ord *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	if ord == nil {
		return nil, cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.CancelSpotOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId, opts...)
	if err.IsNil() {
		UpdateOrderWithRawSpotOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) querySpotOrd(ord *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	if ord == nil {
		return nil, cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.QuerySpotOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId, opts...)
	if err.IsNil() {
		UpdateOrderWithRawSpotOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) newFuOrd(asset, quote string, orderType cex.OrderType, orderSide cex.OrderSide, qty, price float64, opts ...cex.CltOpt) (*resty.Response, *cex.Order, cex.RequestError) {
	symbol := asset + quote
	var tif TimeInForce
	if orderType == cex.OrderTypeLimit {
		tif = TimeInForceGtc
	}
	resp, rawOrd, err := cex.Request(u, FuturesNewOrderConfig, FuturesNewOrderParams{
		Symbol:       symbol,
		PositionSide: u.cfg.fuPosSide,
		Type:         mapStrStr(orderType, ordTypByCexOrdTyp),
		Side:         mapStrStr(orderSide, ordSideByCexOrdSide),
		Quantity:     qty,
		Price:        price,
		TimeInForce:  tif,
	}, opts...)

	ord := SwitchFutureOrderToCexOrder(rawOrd)
	ord.ApiKey = u.api.ApiKey
	return resp, &ord, err
}

func (u *User) cancelFuturesOrd(ord *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	if ord == nil {
		return nil, cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.CancelFuturesOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId, opts...)
	if err.IsNil() {
		UpdateOrderWithRawFuturesOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) queryFuturesOrd(ord *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	if ord == nil {
		return nil, cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.QueryFuturesOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId, opts...)
	if err.IsNil() {
		UpdateOrderWithRawFuturesOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) cancelOrd(ord *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	if ord == nil {
		return nil, cex.RequestError{Err: errors.New("nil order")}
	}
	switch ord.PairType {
	case cex.PairTypeSpot:
		return u.cancelSpotOrd(ord, opts...)
	case cex.PairTypeFutures:
		return u.cancelFuturesOrd(ord, opts...)
	default:
		return nil, cex.RequestError{Err: fmt.Errorf("unknown order pair type %v", ord.PairType)}
	}
}

func (u *User) queryOrd(ord *cex.Order, opts ...cex.CltOpt) (*resty.Response, cex.RequestError) {
	if ord == nil {
		return nil, cex.RequestError{Err: errors.New("nil order")}
	}
	switch ord.PairType {
	case cex.PairTypeSpot:
		return u.querySpotOrd(ord, opts...)
	case cex.PairTypeFutures:
		return u.queryFuturesOrd(ord, opts...)
	default:
		return nil, cex.RequestError{Err: fmt.Errorf("unknown order pair type %v", ord.PairType)}
	}
}

func (u *User) waitOrd(ctx context.Context, ord *cex.Order, opts ...cex.CltOpt) chan cex.RequestError {
	ch := make(chan cex.RequestError, 1)
	if ord == nil {
		ch <- cex.RequestError{Err: errors.New("nil order")}
		return ch
	}
	if ord.IsFinished() {
		ch <- cex.RequestError{}
		return ch
	}
	go func() {
		for {
			_, err := u.queryOrd(ord, opts...)
			if err.IsNil() && ord.IsFinished() {
				ch <- cex.RequestError{}
				return
			}
			select {
			case <-ctx.Done():
				ch <- cex.RequestError{Err: fmt.Errorf("ctxerr: %w, requesterr: %w", ctx.Err(), err.Err)}
				return
			case <-time.After(time.Second):
			}
		}
	}()
	return ch
}

func strOrdIdToInt64(id string) int64 {
	i, _ := strconv.ParseInt(id, 10, 64)
	return i
}

var validQuotes = []string{"USDT", "USDC", "BTC", "ETH", "BNB"}

func SplitPairSymbol(symbol, pairQuote string) (asset, quote string, err error) {
	quotes := validQuotes
	if pairQuote != "" {
		quotes = []string{pairQuote}
	}
	for _, quo := range quotes {
		ass, ok := strings.CutSuffix(symbol, quo)
		if ok && len(ass) > 0 {
			return ass, quo, nil
		}
	}
	err = fmt.Errorf("can not split symbol %v into asset and quote", symbol)
	return
}

var cexOrdTypByOrdTyp = map[OrderType]cex.OrderType{
	OrderTypeMarket: cex.OrderTypeMarket,
	OrderTypeLimit:  cex.OrderTypeLimit,
}

var ordTypByCexOrdTyp = map[cex.OrderType]OrderType{
	cex.OrderTypeMarket: OrderTypeMarket,
	cex.OrderTypeLimit:  OrderTypeLimit,
}

var cexOrdSideByOrdSide = map[OrderSide]cex.OrderSide{
	OrderSideBuy:  cex.OrderSideBuy,
	OrderSideSell: cex.OrderSideSell,
}

var ordSideByCexOrdSide = map[cex.OrderSide]OrderSide{
	cex.OrderSideBuy:  OrderSideBuy,
	cex.OrderSideSell: OrderSideSell,
}

var cexOrdStatusByOrdStatus = map[OrderStatus]cex.OrderStatus{
	OrderStatusNew:             cex.OrderStatusNew,
	OrderStatusPartiallyFilled: cex.OrderStatusPartiallyFilled,
	OrderStatusFilled:          cex.OrderStatusFilled,
	OrderStatusCanceled:        cex.OrderStatusCanceled,
	OrderStatusRejected:        cex.OrderStatusRejected,
	OrderStatusExpired:         cex.OrderStatusExpired,
}

type rawMapCexKV interface {
	~string
}

func mapStrStr[K, V rawMapCexKV](raw K, m map[K]V) V {
	if m == nil {
		return V(raw)
	}
	s, ok := m[raw]
	if !ok {
		return V(raw)
	}
	return s
}

func SwitchSpotOrderToCexOrder(rawOrd SpotOrder) cex.Order {
	ordTyp := mapStrStr(rawOrd.Type, cexOrdTypByOrdTyp)
	ordSide := mapStrStr(rawOrd.Side, cexOrdSideByOrdSide)
	ordStatus := mapStrStr(rawOrd.Status, cexOrdStatusByOrdStatus)

	filledQty := rawOrd.ExecutedQty
	filledQuote := rawOrd.CummulativeQuoteQty
	var avgp float64
	if filledQty != 0 {
		avgp = filledQuote / filledQty
	}

	return cex.Order{
		OriQty:         rawOrd.OrigQty,
		OriPrice:       rawOrd.Price,
		Cex:            cex.BINANCE,
		PairType:       cex.PairTypeSpot,
		OrderType:      ordTyp,
		OrderSide:      ordSide,
		Symbol:         rawOrd.Symbol,
		TimeInForce:    string(rawOrd.TimeInForce),
		ClientOrderId:  rawOrd.ClientOrderId,
		ApiKey:         "",
		OrderId:        strconv.FormatInt(rawOrd.OrderId, 10),
		Status:         ordStatus,
		FilledQty:      filledQty,
		FilledAvgPrice: avgp,
		FilledQuote:    filledQuote,
		RawOrder:       rawOrd,
	}
}

func UpdateOrderWithRawSpotOrder(ord *cex.Order, rawOrd SpotOrder) {
	if ord == nil {
		return
	}
	filledQty := rawOrd.ExecutedQty
	filledQuote := rawOrd.CummulativeQuoteQty
	var avgp float64
	if filledQty != 0 {
		avgp = filledQuote / filledQty
	}
	ord.Status = mapStrStr(rawOrd.Status, cexOrdStatusByOrdStatus)
	ord.FilledQty = filledQty
	ord.FilledQuote = filledQuote
	ord.FilledAvgPrice = avgp
	ord.RawOrder = rawOrd
}

func SwitchFutureOrderToCexOrder(rawOrd FuturesOrder) cex.Order {
	return cex.Order{
		OriQty:         rawOrd.OrigQty,
		OriPrice:       rawOrd.Price,
		Cex:            cex.BINANCE,
		PairType:       cex.PairTypeFutures,
		OrderType:      mapStrStr(rawOrd.Type, cexOrdTypByOrdTyp),
		OrderSide:      mapStrStr(rawOrd.Side, cexOrdSideByOrdSide),
		Symbol:         rawOrd.Symbol,
		TimeInForce:    string(rawOrd.TimeInForce),
		ClientOrderId:  rawOrd.ClientOrderId,
		ApiKey:         "",
		OrderId:        strconv.FormatInt(rawOrd.OrderId, 10),
		Status:         mapStrStr(rawOrd.Status, cexOrdStatusByOrdStatus),
		FilledQty:      rawOrd.ExecutedQty,
		FilledQuote:    rawOrd.CumQuote,
		FilledAvgPrice: rawOrd.AvgPrice,
		RawOrder:       rawOrd,
	}
}

func UpdateOrderWithRawFuturesOrder(ord *cex.Order, rawOrd FuturesOrder) {
	if ord == nil {
		return
	}
	ord.Status = mapStrStr(rawOrd.Status, cexOrdStatusByOrdStatus)
	ord.FilledQty = rawOrd.ExecutedQty
	ord.FilledQuote = rawOrd.CumQuote
	ord.FilledAvgPrice = rawOrd.AvgPrice
	ord.RawOrder = rawOrd
}

// ------------------------------------------------------------
// Private Trade Functions
// ============================================================

// ------------------------------------------------------------
// ReqMaker
// ============================================================

func (u *User) Make(config cex.ReqBaseConfig, reqData any, opts ...cex.CltOpt) (*resty.Request, error) {
	if config.IsUserData {
		return u.makePrivateReq(config, reqData, opts...)
	} else {
		return u.makePublicReq(config, reqData, opts...)
	}
}

func (u *User) makePublicReq(config cex.ReqBaseConfig, reqData any, opts ...cex.CltOpt) (*resty.Request, error) {
	m, err := s2m.ToStrMap(reqData)
	if err != nil {
		return nil, fmt.Errorf("bnc: make public request, %w", err)
	}
	val := url.Values{}
	for k, v := range m {
		val.Set(k, v)
	}
	clt := resty.New().
		SetBaseURL(config.BaseUrl + config.Path + "?" + val.Encode())
	for _, opt := range opts {
		opt(clt)
	}
	req := clt.R()
	return req, nil
}

func (u *User) makePrivateReq(config cex.ReqBaseConfig, reqData any, opts ...cex.CltOpt) (*resty.Request, error) {
	query, err := u.sign(reqData)
	if err != nil {
		return nil, err
	}
	// must compose url by self
	// url.Values composing is alphabetical
	// but binance require signature as the last one
	clt := resty.New().
		SetHeader("X-MBX-APIKEY", u.api.ApiKey).
		SetBaseURL(config.BaseUrl + config.Path + "?" + query)
	for _, opt := range opts {
		opt(clt)
	}
	req := clt.R()
	return req, nil
}

func (u *User) HandleResp(resp *resty.Response, req *resty.Request) error {
	if resp == nil {
		return errors.New("bnc: response checker: response is nil")
	}

	// check http code
	httpCode := resp.StatusCode()
	if httpCode != 200 {
		cexStdErr := HTTPStatusCodeChecker(httpCode)
		if cexStdErr != nil {
			return fmt.Errorf("bnc: http code: %v, status: %v, err: %w", httpCode, resp.Status(), cexStdErr)
		}
	}

	// check binance error code
	body := resp.Body()
	codeMsg := new(CodeMsg)
	if err := json.Unmarshal(body, codeMsg); err != nil {
		// nil err means body is not CodeMsg
		return nil
	}
	if codeMsg.Code >= 0 {
		return nil
	}
	return fmt.Errorf("bnc: msg: %v, code: %v", codeMsg.Msg, codeMsg.Code)
}

// ------------------------------------------------------------
// ReqMaker
// ============================================================

// ============================================================
// Signer
// ------------------------------------------------------------

func (u *User) sign(data any) (query string, err error) {
	return signReqData(data, u.api.SecretKey)
}

func signReqData(data any, key string) (query string, err error) {
	m, err := s2m.ToStrMap(data)
	if err != nil {
		err = fmt.Errorf("%w: %w", cex.ErrS2M, err)
		return
	}
	val := url.Values{
		"timestamp": []string{strconv.FormatInt(time.Now().UnixMilli(), 10)},
	}
	for k, v := range m {
		val.Set(k, v)
	}
	query = val.Encode()
	sig := cex.SignByHmacSHA256ToHex(query, key)
	// binance requires that the signature must be the last one
	query += "&signature=" + sig
	return
}

// ------------------------------------------------------------
// Signer
// ============================================================
