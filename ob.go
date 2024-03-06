package cex

import (
	"errors"
	"fmt"
	"time"

	"github.com/mohae/deepcopy"
)

func NotPairType(t PairType) bool {
	switch t {
	case PairTypeSpot, PairTypeFutures:
		return false
	}
	return true
}

type PQ []float64

func (pq PQ) Price() (float64, error) {
	if len(pq) != 2 {
		return 0, errors.New("book pq: len != 2")
	}
	return pq[0], nil
}

func (pq PQ) Qty() (float64, error) {
	if len(pq) != 2 {
		return 0, errors.New("book pq: len != 2")
	}
	return pq[1], nil
}

// ObBook
// one side book
type ObBook []PQ

func (b ObBook) Copy() ObBook {
	nb := deepcopy.Copy(b)
	return nb.(ObBook)
}

type ObData struct {
	Cex         Name     `json:"cex" bson:"cex"`
	Type        PairType `json:"type" bson:"type"`
	Symbol      string   `json:"symbol" bson:"symbol"`
	Version     string   `json:"version" bson:"version"`
	Time        int64    `json:"time" bson:"time"`
	Asks        ObBook   `json:"asks" bson:"asks"`
	Bids        ObBook   `json:"bids" bson:"bids"`
	Empty       bool     `json:"empty" bson:"empty"`
	EmptyReason string   `json:"emptyReason" bson:"emptyReason"`
}

func (o *ObData) Copy() *ObData {
	no := deepcopy.Copy(o)
	return no.(*ObData)
}

func (o *ObData) SetEmpty(reason string) {
	o.Asks = ObBook{}
	o.Bids = ObBook{}
	o.Empty = true
	o.EmptyReason = reason
	o.Time = time.Now().UnixMilli()
}

func (o *ObData) SetBook(ask bool, book ObBook, version string) {
	if ask {
		o.SetAskBook(book, version)
	} else {
		o.SetBidBook(book, version)
	}
}

func (o *ObData) SetAskBook(askBook ObBook, version string) {
	o.Asks = askBook
	o.Version = version
	o.Time = time.Now().UnixMilli()
}

func (o *ObData) SetBidBook(bidBook ObBook, version string) {
	o.Bids = bidBook
	o.Version = version
	o.Time = time.Now().UnixMilli()
}

func (o *ObData) UpdateDeltas(ask bool, delta ObBook, version string) error {
	if ask {
		return o.UpdateAskDeltas(delta, version)
	} else {
		return o.UpdateBidDeltas(delta, version)
	}
}

func (o *ObData) UpdateAskDeltas(deltaData ObBook, version string) error {
	return o.updateDeltas(deltaData, true, version)
}

func (o *ObData) UpdateBidDeltas(deltaData ObBook, version string) error {
	return o.updateDeltas(deltaData, false, version)
}

func (o *ObData) updateDeltas(newData ObBook, isAsk bool, version string) error {
	for _, v := range newData {
		price, qty := v[0], v[1]
		if price == 0 {
			continue
		}
		if price < 0 {
			return fmt.Errorf("ob: price %v < 0", price)
		}
		if qty < 0 {
			return fmt.Errorf("ob: qty %v < 0", qty)
		}
		var book ObBook
		if isAsk {
			book = o.Asks
		} else {
			book = o.Bids
		}
		book = UpdateOneBookDelta(price, qty, book, isAsk)
		if isAsk {
			o.Asks = book
		} else {
			o.Bids = book
		}
	}
	o.Version = version
	o.Time = time.Now().UnixMilli()
	return nil
}

func (o *ObData) String() string {
	return fmt.Sprintf("ob-%v-%v-%v", o.Cex, o.Type, o.Symbol)
}

func UpdateOneBookDelta(price, qty float64, oldBook ObBook, isAsk bool) ObBook {
	updated := false
	for i, v := range oldBook {
		oldPrice := v[0]
		if price == oldPrice {
			if qty > 0 {
				oldBook[i] = []float64{price, qty}
			} else {
				oldBook = append(oldBook[:i], oldBook[i+1:]...)
			}
			updated = true
			break
		}
		if (isAsk && price < oldPrice) || (!isAsk && price > oldPrice) {
			if qty > 0 {
				oldBook = append(oldBook, []float64{0, 0})
				copy(oldBook[i+1:], oldBook[i:])
				oldBook[i] = []float64{price, qty}
			}
			updated = true
			break
		}
	}
	if !updated && qty > 0 {
		oldBook = append(oldBook, []float64{price, qty})
	}
	return oldBook
}
