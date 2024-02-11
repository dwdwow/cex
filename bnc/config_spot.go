package bnc

import (
	"github.com/dwdwow/cex"
	"net/http"
)

type CoinInfo struct {
	Coin             string            `json:"coin"`
	DepositAllEnable bool              `json:"depositAllEnable"`
	Free             float64           `json:"free,string"`
	Freeze           float64           `json:"freeze,string"`
	Ipoable          float64           `json:"ipoable,string"`
	Ipoing           float64           `json:"ipoing,string"`
	IsLegalMoney     bool              `json:"isLegalMoney"`
	Locked           float64           `json:"locked,string"`
	Name             string            `json:"name"`
	NetworkList      []CoinNetworkInfo `json:"networkList"`
}

type CoinNetworkInfo struct {
	AddressRegex            string `json:"addressRegex"`
	Coin                    string `json:"coin"`
	DepositDesc             string `json:"depositDesc"`
	DepositEnable           bool   `json:"depositEnable"`
	IsDefault               bool   `json:"isDefault"`
	MemoRegex               string `json:"memoRegex"`
	MinConfirm              int    `json:"minConfirm"`
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	SpecialTips             string `json:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm"`
	WithdrawDesc            string `json:"withdrawDesc"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress"`
	EstimatedArrivalTime    int    `json:"estimatedArrivalTime"`
	Busy                    bool   `json:"busy"`
}

var CoinInfoConfig = cex.ReqConfig[cex.EmptyReqData, []CoinInfo]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:               ApiBaseUrl,
		Path:                  SapiV1 + "/capital/config/getall",
		Method:                http.MethodGet,
		IsUserData:            true,
		UserTimeInterval:      0,
		IpTimeInterval:        0,
		HttpStatusCodeChecker: HttpStatusCodeChecker,
		CexCustomCodeChecker:  CustomRespCodeChecker,
	},
}

type SpotBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free,string"`
	Locked float64 `json:"locked,string"`
}

type SpotAccountInfo struct {
	MakerCommission  float64 `json:"makerCommission" bson:"makerCommission"`
	TakerCommission  float64 `json:"takerCommission" bson:"takerCommission"`
	BuyerCommission  float64 `json:"buyerCommission" bson:"buyerCommission"`
	SellerCommission float64 `json:"sellerCommission" bson:"sellerCommission"`
	CommissionRates  struct {
		Maker  float64 `json:"maker,string" bson:"maker"`
		Taker  float64 `json:"taker,string" bson:"taker"`
		Buyer  float64 `json:"buyer,string" bson:"buyer"`
		Seller float64 `json:"seller,string" bson:"seller"`
	} `json:"commissionRates" bson:"commissionRates"`
	CanTrade                   bool          `json:"canTrade" bson:"canTrade"`
	CanWithdraw                bool          `json:"canWithdraw" bson:"canWithdraw"`
	CanDeposit                 bool          `json:"canDeposit" bson:"canDeposit"`
	Brokered                   bool          `json:"brokered" bson:"brokered"`
	RequireSelfTradePrevention bool          `json:"requireSelfTradePrevention" bson:"requireSelfTradePrevention"`
	UpdateTime                 int64         `json:"updateTime" bson:"updateTime"`
	AccountType                AcctType      `json:"accountType" bson:"accountType"`
	Balances                   []SpotBalance `json:"balances" bson:"balances"`
	Permissions                []TradeType   `json:"permissions" bson:"permissions"`
}

var SpotAccountConfig = cex.ReqConfig[cex.EmptyReqData, SpotAccountInfo]{
	ReqBaseConfig: cex.ReqBaseConfig{
		BaseUrl:               ApiBaseUrl,
		Path:                  SapiV1 + "/account",
		Method:                http.MethodGet,
		IsUserData:            true,
		UserTimeInterval:      0,
		IpTimeInterval:        0,
		HttpStatusCodeChecker: HttpStatusCodeChecker,
		CexCustomCodeChecker:  CustomRespCodeChecker,
	},
}
