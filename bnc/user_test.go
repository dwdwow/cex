package bnc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/props"
	"github.com/go-resty/resty/v2"
)

func newTestUser() *User {
	apiKey := readApiKey()
	return NewUser(apiKey.ApiKey, apiKey.SecretKey, UserOptPositionSide(FuturesPositionSideBoth))
}

func newTestVIPPortmarUser() *User {
	apiKey := readVIPPortmarApiKey()
	return NewUser(apiKey.ApiKey, apiKey.SecretKey, UserOptPositionSide(FuturesPositionSideBoth))
}

func userTestChecker[RespData any](resp *resty.Response, respData RespData, err cex.RequestError) {
	props.PanicIfNotNil(err.Err)
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
	userTestChecker(newTestUser().Transfer(TransferTypeMainUmfuture, "USDT", 10))
}

func TestUser_SimpleEarnFlexibleProducts(t *testing.T) {
	userTestChecker(newTestUser().SimpleEarnFlexibleProducts("ETH"))
}

func TestUser_SimpleEarnFlexiblePositions(t *testing.T) {
	userTestChecker(newTestUser().SimpleEarnFlexiblePositions("ETH", "ETH001"))
}

func TestUser_SimpleEarnFlexibleRedeem(t *testing.T) {
	userTestChecker(newTestUser().SimpleEarnFlexibleRedeem("ETH001", false, 0.02, SimpleEarnFlexibleRedeemDestinationSpot))
}

func TestUser_SimpleEarnFlexibleRateHistories(t *testing.T) {
	userTestChecker(newTestUser().SimpleEarnFlexibleRateHistories("USDT001", time.Now().UnixMilli()-time.Hour.Milliseconds()*100, 0))
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
	userTestChecker(newTestUser().CryptoLoanFlexibleRepay("USDT", "ETH", 100, BigTrue, BigFalse))
}

func TestUser_CryptoLoanFlexibleRepaymentHistories(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleRepaymentHistories("USDT", "ETH"))
}

func TestUser_CryptoLoanFlexibleAdjustLtv(t *testing.T) {
	userTestChecker(newTestUser().CryptoLoanFlexibleAdjustLtv("USDT", "ETH", 0.03, LTVReduced))
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

func TestUser_NewSpotOrder(t *testing.T) {
	userTestChecker(newTestUser().NewSpotOrder("ETH", "USDT", cex.OrderTypeLimit, cex.OrderSideBuy, 0.01, 1500))
}

func TestUser_QuerySpotOrder(t *testing.T) {
	userTestChecker(newTestUser().QuerySpotOrder("ETHUSDT", 0, ""))
}

func TestUser_CancelSpotOrder(t *testing.T) {
	userTestChecker(newTestUser().CancelSpotOrder("ETHUSDT", 0, ""))
}

func TestUser_NewSpotLimitBuyOrder(t *testing.T) {
	userTestChecker(newTestUser().NewSpotLimitBuyOrder("ETH", "USDT", 0.01, 1600))
}

func TestUser_NewSpotLimitSellOrder(t *testing.T) {
	userTestChecker(newTestUser().NewSpotLimitSellOrder("ETH", "USDT", 0.01, 3000))
}

func TestUser_NewSpotMarketBuyOrder(t *testing.T) {
	userTestChecker(newTestUser().NewSpotMarketBuyOrder("ETH", "USDT", 0.01))
}

func TestUser_NewSpotMarketSellOrder(t *testing.T) {
	userTestChecker(newTestUser().NewSpotMarketSellOrder("ETH", "USDT", 0.01))
}

func TestUser_NewFuturesOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesOrder("ETH", "USDT", cex.OrderTypeLimit, cex.OrderSideBuy, 0.01, 1500))
}

func TestUser_QueryFuturesOrder(t *testing.T) {
	userTestChecker(newTestUser().QueryFuturesOrder("ETHUSDT", 0, ""))
}

func TestUser_CancelFuturesOrder(t *testing.T) {
	userTestChecker(newTestUser().CancelFuturesOrder("ETHUSDT", 0, ""))
}

func TestUser_NewFuturesLimitBuyOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesLimitBuyOrder("ETH", "USDT", 0.01, 1600))
}

func TestUser_NewFuturesLimitSellOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesLimitSellOrder("ETH", "USDT", 0.01, 3000))
}

func TestUser_NewFuturesMarketBuyOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesMarketBuyOrder("ETH", "USDT", 0.01))
}

func TestUser_NewFuturesMarketSellOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesMarketSellOrder("ETH", "USDT", 0.01))
}

func TestUser_NewFuturesCMOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesCMOrder("ETH", "BTC", cex.OrderTypeLimit, cex.OrderSideBuy, 0.01, 1500))
}

func TestUser_QueryFuturesCMOrder(t *testing.T) {
	userTestChecker(newTestUser().QueryFuturesOrder("ETHBTC", 0, ""))
}

func TestUser_CancelFuturesCMOrder(t *testing.T) {
	userTestChecker(newTestUser().CancelFuturesOrder("ETHBTC", 0, ""))
}

func TestUser_NewFuturesLimitBuyCMOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesLimitBuyCMOrder("ETH", "BTC", 0.01, 1600))
}

func TestUser_NewFuturesLimitSellCMOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesLimitSellCMOrder("ETH", "BTC", 0.01, 3000))
}

func TestUser_NewFuturesMarketBuyCMOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesMarketBuyCMOrder("ETH", "BTC", 0.01))
}

func TestUser_NewFuturesMarketSellCMOrder(t *testing.T) {
	userTestChecker(newTestUser().NewFuturesMarketSellCMOrder("ETH", "BTC", 0.01))
}

func TestUser_QueryOrder(t *testing.T) {
	_, ord, err := newTestUser().NewSpotLimitBuyOrder("ETH", "USDT", 0.01, 1900)
	props.PanicIfNotNil(err.Err)
	props.PrintlnIndent(ord)
	_, err = newTestUser().QueryOrder(ord)
	props.PanicIfNotNil(err.Err)
	props.PrintlnIndent(ord)
	_, err = newTestUser().CancelOrder(ord)
	props.PanicIfNotNil(err.Err)
	props.PrintlnIndent(ord)
}

func TestUser_WaitOrder(t *testing.T) {
	fmt.Println("new order")
	_, ord, err := newTestUser().NewSpotLimitBuyOrder("ETH", "USDT", 0.01, 1900)
	props.PanicIfNotNil(err.Err)
	props.PrintlnIndent(ord)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 10)
		cancel()
	}()
	fmt.Println("wait order")
	chErr := newTestUser().WaitOrder(ctx, ord)
	props.PrintlnIndent(ord)
	err = <-chErr
	if err.IsNotNil() && err.Is(context.Canceled) {
		fmt.Println("ctx canceled")
		fmt.Println("cancel order")
		_, err = newTestUser().CancelOrder(ord)
		if err.Is(cex.ErrUnknownOrder) {
			panic(err)
		}
		props.PrintlnIndent(ord)
	}
	fmt.Println("query order")
	_, err = newTestUser().queryOrd(ord)
	props.PanicIfNotNil(err.Err)
	props.PrintlnIndent(ord)
}

func TestUser_Withdraw(t *testing.T) {
	userTestChecker(newTestUser().Withdraw("BOME", NetworkSol, "", 600))
}

func TestUser_DepositAddress(t *testing.T) {
	userTestChecker(newTestUser().DepositAddress("BOME", NetworkSol))
}

func TestUser_VIPLoanOngoingOrders(t *testing.T) {
	userTestChecker(newTestVIPPortmarUser().VIPLoanOngoingOrders(0, 0, "", ""))
}
