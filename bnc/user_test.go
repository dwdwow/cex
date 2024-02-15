package bnc

import (
	"testing"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/props"
	"github.com/go-resty/resty/v2"
)

func newTestUser() *User {
	apiKey := readApiKey()
	return NewUser(apiKey.ApiKey, apiKey.SecretKey, UserOptPositionSide(FuPosBoth))
}

func userTestChecker[RespData any](resp *resty.Response, respData RespData, err *cex.RequestError) {
	props.PanicIfNotNil(err)
	props.PrintlnIndent(respData)
}

func TestUser_Coins(t *testing.T) {
	userTestChecker(newTestUser().Coins())
}

func TestUser_SpotAccount(t *testing.T) {
	userTestChecker(newTestUser().SpotAccount())
}

func TestUser_FuturesAccount(t *testing.T) {
	userTestChecker(newTestUser().FuturesAccount())
}

func TestUser_Transfer(t *testing.T) {
	userTestChecker(newTestUser().Transfer(TranType_MAIN_UMFUTURE, "USDT", 10))
}

func TestUser_SimpleEarnFlexibleProducts(t *testing.T) {
	userTestChecker(newTestUser().SimpleEarnFlexibleProducts("ETH"))
}

func TestUser_SimpleEarnFlexiblePositions(t *testing.T) {
	userTestChecker(newTestUser().SimpleEarnFlexiblePositions("ETH", "ETH001"))
}

func TestUser_SimpleEarnFlexibleRedeem(t *testing.T) {
	userTestChecker(newTestUser().SimpleEarnFlexibleRedeem("ETH001", false, 0.02, FlexibleRedeemDestSpot))
}

func TestUser_CryptoLoanFlexibleOngoingOrders(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleOngoingOrders("USDT", "ETH"))
}

func TestUser_CryptoLoanIncomeHistories(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanIncomeHistories("", ""))
}

func TestUser_CryptoLoanFlexibleBorrow(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleBorrow("USDT", "ETH", 100, 0))
}

func TestUser_CryptoLoanFlexibleBorrowHistories(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleBorrowHistories("USDT", "ETH"))
}

func TestUser_CryptoLoanFlexibleRepay(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleRepay("USDT", "ETH", 100, TRUE, FALSE))
}

func TestUser_CryptoLoanFlexibleRepaymentHistories(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleRepaymentHistories("USDT", "ETH"))
}

func TestUser_CryptoLoanFlexibleAdjustLtv(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleAdjustLtv("USDT", "ETH", 0.03, LTVAdDireReduced))
}

func TestUser_CryptoLoanFlexibleAdjustLtvHistories(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleAdjustLtvHistories("USDT", "ETH"))
}

func TestUser_CryptoLoanFlexibleLoanAssets(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleLoanAssets(""))
}

func TestUser_CryptoLoanFlexibleCollateralAssets(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleCollateralAssets(""))
}
