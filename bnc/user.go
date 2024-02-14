package bnc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/s2m"
	"github.com/go-resty/resty/v2"
)

type User struct {
	api cex.Api
}

func NewUser(apiKey, secretKey string) User {
	return User{
		api: cex.Api{
			Cex:       cex.BINANCE,
			ApiKey:    apiKey,
			SecretKey: secretKey,
		},
	}
}

// ============================================================
// Spot API
// ------------------------------------------------------------

func (u User) Coins() (*resty.Response, []Coin, *cex.RequestError) {
	return cex.Request(u, CoinInfoConfig, nil)
}

func (u User) SpotAccount() (*resty.Response, SpotAccount, *cex.RequestError) {
	return cex.Request(u, SpotAccountConfig, nil)
}

func (u User) Transfer(tranType TranType, asset string, amount float64) (*resty.Response, UniversalTransferResp, *cex.RequestError) {
	return cex.Request(u, UniversalTransferConfig, UniversalTransferParams{Type: tranType, Asset: asset, Amount: amount})
}

func (u User) FlexibleProducts(asset string) (*resty.Response, Page[[]FlexibleProduct], *cex.RequestError) {
	return cex.Request(u, FlexibleProductConfig, FlexibleProductListParams{Asset: asset})
}

func (u User) CryptoLoanIncomeHistories(asset string, incomeType CryptoLoanIncomeType) (*resty.Response, []CryptoLoanIncomeHistory, *cex.RequestError) {
	return cex.Request(u, CryptoLoansIncomeHistoriesConfig, CryptoLoansIncomeHistoriesParams{Asset: asset, Type: incomeType})
}

func (u User) FlexibleBorrow(loanCoin string, collateralCoin string, loanAmount, collateralAmount float64) (*resty.Response, FlexibleBorrowResult, *cex.RequestError) {
	return cex.Request(u, FlexibleBorrowConfig, FlexibleBorrowParams{LoanCoin: loanCoin, LoanAmount: loanAmount, CollateralCoin: collateralCoin, CollateralAmount: collateralAmount})
}

func (u User) FlexibleBorrowHistories(loanCoin, collateralCoin string) (*resty.Response, Page[[]FlexibleBorrowHistory], *cex.RequestError) {
	return cex.Request(u, FlexibleBorrowHistoriesConfig, FlexibleBorrowHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin})
}

func (u User) FlexibleRepay(loanCoin, collateralCoin string, repayAmount float64, collateralReturn, fullRepayment BigBool) (*resty.Response, FlexibleRepayResult, *cex.RequestError) {
	return cex.Request(u, FlexibleRepayConfig, FlexibleRepayParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, RepayAmount: repayAmount, CollateralReturn: collateralReturn, FullRepayment: fullRepayment})
}

func (u User) FlexibleRepaymentHistories(loanCoin, collateralCoin string) (*resty.Response, Page[[]FlexibleRepaymentHistory], *cex.RequestError) {
	return cex.Request(u, FlexibleRepaymentHistoriesConfig, FlexibleRepaymentHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin})
}

func (u User) FlexibleAdjustLtv(loanCoin, collateralCoin string, adjustmentAmount float64, direction LTVAdjustDirection) (*resty.Response, FlexibleLoanAdjustLtvResult, *cex.RequestError) {
	return cex.Request(u, FlexibleLoanAdjustLtvConfig, FlexibleAdjustLtvParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin, AdjustmentAmount: adjustmentAmount, Direction: direction})
}

func (u User) FlexibleAdjustLtvHistories(loanCoin, collateralCoin string) (*resty.Response, Page[[]FlexibleAdjustLtvHistory], *cex.RequestError) {
	return cex.Request(u, FlexibleAdjustLtvHistoriesConfig, FlexibleAdjustLtvHistoriesParams{LoanCoin: loanCoin, CollateralCoin: collateralCoin})
}

func (u User) FlexibleLoanAssets(loanCoin string) (*resty.Response, Page[[]FlexibleLoanAsset], *cex.RequestError) {
	return cex.Request(u, FlexibleLoanAssetsConfig, FlexibleLoanAssetsParams{loanCoin})
}

func (u User) FlexibleCollateralAssets(collateralCoin string) (*resty.Response, Page[[]FlexibleCollateralCoin], *cex.RequestError) {
	return cex.Request(u, FlexibleCollateralCoinsConfig, FlexibleCollateralCoinsParams{collateralCoin})
}

// ------------------------------------------------------------
// Spot API
// ============================================================

// ============================================================
// Trade private functions
// ------------------------------------------------------------

// ------------------------------------------------------------
// Trade private functions
// ============================================================

// ------------------------------------------------------------
// ReqMaker
// ============================================================

func (u User) Make(config cex.ReqBaseConfig, reqData any, opts ...cex.ReqOpt) (*resty.Request, error) {
	if config.IsUserData {
		return u.makePrivateReq(config, reqData, opts...)
	} else {
		return u.makePublicReq(config, reqData, opts...)
	}
}

func (u User) makePublicReq(config cex.ReqBaseConfig, reqData any, opts ...cex.ReqOpt) (*resty.Request, error) {
	m, err := s2m.ToStrMap(reqData)
	if err != nil {
		return nil, fmt.Errorf("bnc: make public request, %w", err)
	}
	val := url.Values{}
	for k, v := range m {
		val.Set(k, v)
	}
	clt := resty.New().
		SetBaseURL(config.BaseUrl)
	req := clt.R().
		SetQueryString(val.Encode())
	for _, opt := range opts {
		opt(clt, req)
	}
	return nil, nil
}

func (u User) makePrivateReq(config cex.ReqBaseConfig, reqData any, opts ...cex.ReqOpt) (*resty.Request, error) {
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

func (u User) HandleResp(resp *resty.Response, req *resty.Request) error {
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

func (u User) sign(data any) (query string, err error) {
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
