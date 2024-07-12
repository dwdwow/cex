package bnc

// WsMarkPriceUpdate 3000ms or 1000ms
type WsMarkPriceUpdate struct {
	EventType            WsEvent `json:"e"`
	EventTime            int64   `json:"E"`
	Symbol               string  `json:"s"`
	MarkPrice            string  `json:"p"`
	IndexPrice           string  `json:"i"`
	EstimatedSettlePrice string  `json:"P"` // Estimated Settle Price, only useful in the last hour before the settlement starts
	FundingRate          string  `json:"r"`
	NextFundingTime      int64   `json:"T"`
}

type WsCMIndexPriceStream struct {
	EventType WsEvent `json:"e"`
	EventTime int64   `json:"E"`
	Pair      string  `json:"i"`
	Price     float64 `json:"p,string"`
}

type WsCMMarkPriceStream struct {
	EventType            WsEvent `json:"e"`
	EventTime            int64   `json:"E"`
	Symbol               string  `json:"s"`
	MarkPrice            float64 `json:"p,string"`
	EstimatedSettlePrice float64 `json:"P,string"`
	IndexPrice           float64 `json:"i,string"`
	FundingRate          float64 `json:"r,string"`
	NextFundingRate      float64 `json:"T,string"`
}
