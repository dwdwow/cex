package bnc

const (
	ApiBaseUrl = "https://api.binance.com"
	ApiV3      = "/api/v3"
	SapiV1     = "/sapi/v1"

	FutureApiBaseUrl = "https://fapi.binance.com"
	FapiV1           = "/fapi/v1"
	FapiV2           = "/fapi/v2"
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
