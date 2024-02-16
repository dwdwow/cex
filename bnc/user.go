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
// Account API
// ------------------------------------------------------------

func (u *User) Coins() (*resty.Response, []Coin, *cex.RequestError) {
	return cex.Request(u, CoinInfoConfig, nil)
}

func (u *User) SpotAccount() (*resty.Response, SpotAccount, *cex.RequestError) {
	return cex.Request(u, SpotAccountConfig, nil)
}

func (u *User) Transfer(tranType TransferType, asset string, amount float64) (*resty.Response, UniversalTransferResp, *cex.RequestError) {
	return cex.Request(u, UniversalTransferConfig, UniversalTransferParams{Type: tranType, Asset: asset, Amount: amount})
}

func (u *User) FuturesAccount() (*resty.Response, FuAccount, *cex.RequestError) {
	return cex.Request(u, FuAccountConfig, nil)
}

func (u *User) FuturesPositions(symbol string) (*resty.Response, []FuPosition, *cex.RequestError) {
	return cex.Request(u, FuPositionsConfig, FuPositionsParams{Symbol: symbol})
}

// ------------------------------------------------------------
// Account API
// ============================================================

// ============================================================
// Flexible Simple Earn API
// ------------------------------------------------------------

func (u *User) SimpleEarnFlexibleProducts(asset string) (*resty.Response, Page[[]FlexibleProduct], *cex.RequestError) {
	return cex.Request(u, FlexibleProductConfig, FlexibleProductListParams{Asset: asset})
}

func (u *User) SimpleEarnFlexiblePositions(asset, productId string) (*resty.Response, Page[[]SimpleEarnFlexiblePosition], *cex.RequestError) {
	return cex.Request(u, SimpleEarnFlexiblePositionsConfig, SimpleEarnFlexiblePositionsParams{Asset: asset, ProductId: productId})
}

func (u *User) SimpleEarnFlexibleRedeem(productId string, redeemAll bool, amount float64, destAccount SimpleEarnFlexibleRedeemDestination) (*resty.Response, FlexibleRedeemResponse, *cex.RequestError) {
	return cex.Request(u, FlexibleRedeemConfig, FlexibleRedeemParams{ProductId: productId, RedeemAll: redeemAll, Amount: amount, DestAccount: destAccount})
}

// ------------------------------------------------------------
// Flexible Simple Earn API
// ============================================================

// ============================================================
// Flexible Loan API
// ------------------------------------------------------------

func (u *User) CryptoLoanFlexibleOngoingOrders(loanCoin, collateralCoin string) (*resty.Response, Page[[]FlexibleOngoingOrder], *cex.RequestError) {
	return cex.Request(u, FlexibleOngoingOrdersConfig, FlexibleOngoingOrdersParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin})
}

func (u *User) CryptoLoanIncomeHistories(asset string, incomeType CryptoLoanIncomeType) (*resty.Response, []CryptoLoanIncomeHistory, *cex.RequestError) {
	return cex.Request(u, CryptoLoansIncomeHistoriesConfig, CryptoLoansIncomeHistoriesParams{Asset: asset, Type: incomeType})
}

func (u *User) CryptoLoanFlexibleBorrow(loanCoin string, collateralCoin string, loanAmount, collateralAmount float64) (*resty.Response, FlexibleBorrowResult, *cex.RequestError) {
	return cex.Request(u, FlexibleBorrowConfig, FlexibleBorrowParams{LoanCoin: loanCoin, LoanAmount: loanAmount, CollateralCoin: collateralCoin, CollateralAmount: collateralAmount})
}

func (u *User) CryptoLoanFlexibleBorrowHistories(loanCoin, collateralCoin string) (*resty.Response, Page[[]FlexibleBorrowHistory], *cex.RequestError) {
	return cex.Request(u, FlexibleBorrowHistoriesConfig, FlexibleBorrowHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, Limit: 100})
}

func (u *User) CryptoLoanFlexibleRepay(loanCoin, collateralCoin string, repayAmount float64, collateralReturn, fullRepayment BigBool) (*resty.Response, FlexibleRepayResult, *cex.RequestError) {
	return cex.Request(u, FlexibleRepayConfig, FlexibleRepayParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, RepayAmount: repayAmount, CollateralReturn: collateralReturn, FullRepayment: fullRepayment})
}

func (u *User) CryptoLoanFlexibleRepaymentHistories(loanCoin, collateralCoin string) (*resty.Response, Page[[]FlexibleRepaymentHistory], *cex.RequestError) {
	return cex.Request(u, FlexibleRepaymentHistoriesConfig, FlexibleRepaymentHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, Limit: 100})
}

func (u *User) CryptoLoanFlexibleAdjustLtv(loanCoin, collateralCoin string, adjustmentAmount float64, direction LTVAdjustDirection) (*resty.Response, FlexibleLoanAdjustLtvResult, *cex.RequestError) {
	return cex.Request(u, FlexibleLoanAdjustLtvConfig, FlexibleAdjustLtvParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, AdjustmentAmount: adjustmentAmount, Direction: direction})
}

func (u *User) CryptoLoanFlexibleAdjustLtvHistories(loanCoin, collateralCoin string) (*resty.Response, Page[[]FlexibleAdjustLtvHistory], *cex.RequestError) {
	return cex.Request(u, FlexibleAdjustLtvHistoriesConfig, FlexibleAdjustLtvHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, Limit: 100})
}

func (u *User) CryptoLoanFlexibleLoanAssets(loanCoin string) (*resty.Response, Page[[]FlexibleLoanAsset], *cex.RequestError) {
	return cex.Request(u, FlexibleLoanAssetsConfig, FlexibleLoanAssetsParams{loanCoin})
}

func (u *User) CryptoLoanFlexibleCollateralAssets(collateralCoin string) (*resty.Response, Page[[]FlexibleCollateralCoin], *cex.RequestError) {
	return cex.Request(u, FlexibleCollateralCoinsConfig, FlexibleCollateralCoinsParams{collateralCoin})
}

// ------------------------------------------------------------
// Flexible Loan API
// ============================================================

// ============================================================
// cex.Trader Interface Implementations
// ------------------------------------------------------------

func (u *User) QueryOrder(order *cex.Order) (*resty.Response, *cex.RequestError) {
	return u.queryOrd(order)
}

func (u *User) CancelOrder(order *cex.Order) (*resty.Response, *cex.RequestError) {
	return u.cancelOrd(order)
}

func (u *User) WaitOrder(ctx context.Context, order *cex.Order) (*resty.Response, *cex.RequestError) {
	return u.waitOrd(ctx, order)
}

func (u *User) NewSpotOrder(asset, quote string, tradeType cex.OrderType, orderSide cex.OrderSide, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.newSpotOrd(asset, quote, tradeType, orderSide, qty, price)
}

func (u *User) NewSpotLimitBuyOrder(asset, quote string, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideBuy, qty, price)
}

func (u *User) NewSpotLimitSellOrder(asset, quote string, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideSell, qty, price)
}

func (u *User) NewSpotMarketBuyOrder(asset, quote string, qty float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideBuy, qty, 0)
}

func (u *User) NewSpotMarketSellOrder(asset, quote string, qty float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewSpotOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideSell, qty, 0)
}

func (u *User) NewFuturesOrder(asset, quote string, tradeType cex.OrderType, orderSide cex.OrderSide, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.newFuOrd(asset, quote, tradeType, orderSide, qty, price)
}

func (u *User) NewFuturesLimitBuyOrder(asset, quote string, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideBuy, qty, price)
}

func (u *User) NewFuturesLimitSellOrder(asset, quote string, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeLimit, cex.OrderSideSell, qty, price)
}

func (u *User) NewFuturesMarketBuyOrder(asset, quote string, qty float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideBuy, qty, 0)
}

func (u *User) NewFuturesMarketSellOrder(asset, quote string, qty float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	return u.NewFuturesOrder(asset, quote, cex.OrderTypeMarket, cex.OrderSideSell, qty, 0)
}

// ------------------------------------------------------------
// cex.Trader Interface Implementations
// ============================================================

// ============================================================
// Spot API
// ------------------------------------------------------------

func (u *User) CancelSpotOrder(asset, quote string, orderId int64, cltOrdId string) (*resty.Response, SpotOrder, *cex.RequestError) {
	return cex.Request(u, SpotCancelOrderConfig, SpotCancelOrderParams{Symbol: asset + quote, OrderId: orderId, OrigClientOrderId: cltOrdId})
}

func (u *User) QuerySpotOrder(asset, quote string, orderId int64, cltOrdId string) (*resty.Response, SpotOrder, *cex.RequestError) {
	return cex.Request(u, SpotQueryOrderConfig, SpotQueryOrderParams{Symbol: asset + quote, OrderId: orderId, OrigClientOrderId: cltOrdId})
}

// ------------------------------------------------------------
// Spot API
// ============================================================

// ============================================================
// Futures API
// ------------------------------------------------------------

func (u *User) CancelFuturesOrder(asset, quote string, orderId int64, cltOrdId string) (*resty.Response, FuOrder, *cex.RequestError) {
	return cex.Request(u, FuCancelOrderConfig, FuQueryOrCancelOrderParams{Symbol: asset + quote, OrderId: orderId, OrigClientOrderId: cltOrdId})
}

func (u *User) QueryFuturesOrder(asset, quote string, orderId int64, cltOrdId string) (*resty.Response, FuOrder, *cex.RequestError) {
	return cex.Request(u, FuQueryOrderConfig, FuQueryOrCancelOrderParams{Symbol: asset + quote, OrderId: orderId, OrigClientOrderId: cltOrdId})
}

// ------------------------------------------------------------
// Futures API
// ============================================================

// ============================================================
// Private Trade Functions
// ------------------------------------------------------------

func (u *User) newSpotOrd(asset, quote string, orderType cex.OrderType, orderSide cex.OrderSide, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
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
	})
	ord := SwitchSpotOrderToCexOrder(rawOrd)
	ord.ApiKey = u.api.ApiKey
	return resp, &ord, err
}

func (u *User) cancelSpotOrd(ord *cex.Order) (*resty.Response, *cex.RequestError) {
	if ord == nil {
		return nil, &cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.CancelSpotOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId)
	if err == nil {
		UpdateOrderWithRawSpotOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) querySpotOrd(ord *cex.Order) (*resty.Response, *cex.RequestError) {
	if ord == nil {
		return nil, &cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.QuerySpotOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId)
	if err == nil {
		UpdateOrderWithRawSpotOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) newFuOrd(asset, quote string, orderType cex.OrderType, orderSide cex.OrderSide, qty, price float64) (*resty.Response, *cex.Order, *cex.RequestError) {
	symbol := asset + quote
	var tif TimeInForce
	if orderType == cex.OrderTypeLimit {
		tif = TimeInForceGtc
	}
	resp, rawOrd, err := cex.Request(u, FuNewOrderConfig, FuNewOrderParams{
		Symbol:       symbol,
		PositionSide: u.cfg.fuPosSide,
		Type:         mapStrStr(orderType, ordTypByCexOrdTyp),
		Side:         mapStrStr(orderSide, ordSideByCexOrdSide),
		Quantity:     qty,
		Price:        price,
		TimeInForce:  tif,
	})

	ord := SwitchFutureOrderToCexOrder(rawOrd)
	ord.ApiKey = u.api.ApiKey
	return resp, &ord, err
}

func (u *User) cancelFuturesOrd(ord *cex.Order) (*resty.Response, *cex.RequestError) {
	if ord == nil {
		return nil, &cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.CancelFuturesOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId)
	if err == nil {
		UpdateOrderWithRawFuturesOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) queryFuturesOrd(ord *cex.Order) (*resty.Response, *cex.RequestError) {
	if ord == nil {
		return nil, &cex.RequestError{Err: errors.New("nil order")}
	}
	resp, rawOrd, err := u.QueryFuturesOrder(ord.Symbol, strOrdIdToInt64(ord.OrderId), ord.ClientOrderId)
	if err == nil {
		UpdateOrderWithRawFuturesOrder(ord, rawOrd)
	}
	return resp, err
}

func (u *User) cancelOrd(ord *cex.Order) (*resty.Response, *cex.RequestError) {
	if ord == nil {
		return nil, &cex.RequestError{Err: errors.New("nil order")}
	}
	switch ord.PairType {
	case cex.SpotPair:
		return u.cancelSpotOrd(ord)
	case cex.FuturePair:
		return u.cancelFuturesOrd(ord)
	default:
		return nil, &cex.RequestError{Err: fmt.Errorf("unknown order pair type %v", ord.PairType)}
	}
}

func (u *User) queryOrd(ord *cex.Order) (*resty.Response, *cex.RequestError) {
	if ord == nil {
		return nil, &cex.RequestError{Err: errors.New("nil order")}
	}
	switch ord.PairType {
	case cex.SpotPair:
		return u.querySpotOrd(ord)
	case cex.FuturePair:
		return u.queryFuturesOrd(ord)
	default:
		return nil, &cex.RequestError{Err: fmt.Errorf("unknown order pair type %v", ord.PairType)}
	}
}

func (u *User) waitOrd(ctx context.Context, ord *cex.Order) (*resty.Response, *cex.RequestError) {
	if ord == nil {
		return nil, &cex.RequestError{Err: errors.New("nil order")}
	}
	var resp *resty.Response
	var err error
	for {
		select {
		case <-ctx.Done():
			return nil, &cex.RequestError{Err: fmt.Errorf("ctxerr: %w, requesterr: %w", ctx.Err(), err)}
		case <-time.After(time.Second):
		}
		resp, err = u.queryOrd(ord)
		if err == nil && ord.IsFinished() {
			return resp, nil
		}
	}
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
		return any(raw).(V)
	}
	s, ok := m[raw]
	if !ok {
		return any(raw).(V)
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
		PairType:       cex.SpotPair,
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
	ordStatus := mapStrStr(rawOrd.Status, cexOrdStatusByOrdStatus)
	filledQty := rawOrd.ExecutedQty
	filledQuote := rawOrd.CummulativeQuoteQty
	var avgp float64
	if filledQty != 0 {
		avgp = filledQuote / filledQty
	}
	ord.Status = ordStatus
	ord.FilledQty = filledQty
	ord.FilledQuote = filledQuote
	ord.FilledAvgPrice = avgp
}

func SwitchFutureOrderToCexOrder(rawOrd FuOrder) cex.Order {
	ordTyp := mapStrStr(rawOrd.Type, cexOrdTypByOrdTyp)
	ordSide := mapStrStr(rawOrd.Side, cexOrdSideByOrdSide)
	ordStatus := mapStrStr(rawOrd.Status, cexOrdStatusByOrdStatus)

	filledQty := rawOrd.ExecutedQty
	filledQuote := rawOrd.CumQuote
	var avgp float64
	if filledQty != 0 {
		avgp = filledQuote / filledQty
	}

	return cex.Order{
		OriQty:         rawOrd.OrigQty,
		OriPrice:       rawOrd.Price,
		Cex:            cex.BINANCE,
		PairType:       cex.FuturePair,
		OrderType:      ordTyp,
		OrderSide:      ordSide,
		Symbol:         rawOrd.Symbol,
		TimeInForce:    string(rawOrd.TimeInForce),
		ClientOrderId:  rawOrd.ClientOrderId,
		ApiKey:         "",
		OrderId:        strconv.FormatInt(rawOrd.OrderId, 10),
		Status:         ordStatus,
		FilledQty:      filledQty,
		FilledQuote:    filledQuote,
		FilledAvgPrice: avgp,
		RawOrder:       rawOrd,
	}
}

func UpdateOrderWithRawFuturesOrder(ord *cex.Order, rawOrd FuOrder) {
	if ord == nil {
		return
	}
	ordStatus := mapStrStr(rawOrd.Status, cexOrdStatusByOrdStatus)
	filledQty := rawOrd.ExecutedQty
	filledQuote := rawOrd.CumQuote
	var avgp float64
	if filledQty != 0 {
		avgp = filledQuote / filledQty
	}
	ord.Status = ordStatus
	ord.FilledQty = filledQty
	ord.FilledQuote = filledQuote
	ord.FilledAvgPrice = avgp
}

// ------------------------------------------------------------
// Private Trade Functions
// ============================================================

// ------------------------------------------------------------
// ReqMaker
// ============================================================

func (u *User) Make(config cex.ReqBaseConfig, reqData any, opts ...cex.ReqOpt) (*resty.Request, error) {
	if config.IsUserData {
		return u.makePrivateReq(config, reqData, opts...)
	} else {
		return u.makePublicReq(config, reqData, opts...)
	}
}

func (u *User) makePublicReq(config cex.ReqBaseConfig, reqData any, opts ...cex.ReqOpt) (*resty.Request, error) {
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
	req := clt.R()
	for _, opt := range opts {
		opt(clt, req)
	}
	return req, nil
}

func (u *User) makePrivateReq(config cex.ReqBaseConfig, reqData any, opts ...cex.ReqOpt) (*resty.Request, error) {
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
	req := clt.R()
	for _, opt := range opts {
		opt(clt, req)
	}
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
