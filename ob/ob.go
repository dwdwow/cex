package ob

import (
	"errors"
	"fmt"
	"time"

	"github.com/dwdwow/cex"
	"github.com/mohae/deepcopy"
)

type PQ []float64

func (pq PQ) P() (float64, error) {
	if len(pq) != 2 {
		return 0, errors.New("cex: PQ len != 2")
	}
	return pq[0], nil
}

func (pq PQ) Q() (float64, error) {
	if len(pq) != 2 {
		return 0, errors.New("cex: PQ len != 2")
	}
	return pq[1], nil
}

// Book
// one side book
type Book []PQ

func (b Book) Copy() Book {
	nb := deepcopy.Copy(b)
	return nb.(Book)
}

type Data struct {
	Cex         cex.Name     `json:"cex" bson:"cex"`
	Type        cex.PairType `json:"type" bson:"type"`
	Symbol      string       `json:"symbol" bson:"symbol"`
	Version     string       `json:"version" bson:"version"`
	Time        int64        `json:"time" bson:"time"`
	Asks        Book         `json:"asks" bson:"asks"`
	Bids        Book         `json:"bids" bson:"bids"`
	Empty       bool         `json:"empty" bson:"empty"`
	EmptyReason string       `json:"emptyReason" bson:"emptyReason"`
}

func (o *Data) Copy() *Data {
	no := deepcopy.Copy(o)
	return no.(*Data)
}

func (o *Data) SetEmpty(reason string) {
	o.Asks = Book{}
	o.Bids = Book{}
	o.Empty = true
	o.EmptyReason = reason
	o.Time = time.Now().UnixMilli()
}

func (o *Data) SetBook(ask bool, book Book, version string) {
	if ask {
		o.SetAskBook(book, version)
	} else {
		o.SetBidBook(book, version)
	}
}

func (o *Data) SetAskBook(askBook Book, version string) {
	o.Asks = askBook
	o.Version = version
	o.Time = time.Now().UnixMilli()
}

func (o *Data) SetBidBook(bidBook Book, version string) {
	o.Bids = bidBook
	o.Version = version
	o.Time = time.Now().UnixMilli()
}

func (o *Data) UpdateDeltas(ask bool, delta Book, version string) error {
	if ask {
		return o.UpdateAskDeltas(delta, version)
	} else {
		return o.UpdateBidDeltas(delta, version)
	}
}

func (o *Data) UpdateAskDeltas(deltaData Book, version string) error {
	return o.updateDeltas(deltaData, true, version)
}

func (o *Data) UpdateBidDeltas(deltaData Book, version string) error {
	return o.updateDeltas(deltaData, false, version)
}

func (o *Data) updateDeltas(newData Book, isAsk bool, version string) error {
	for _, v := range newData {
		price, qty := v[0], v[1]
		if price == 0 {
			continue
		}
		if price < 0 {
			return fmt.Errorf("cex: ob price %v < 0", price)
		}
		if qty < 0 {
			return fmt.Errorf("cex: ob qty %v < 0", qty)
		}
		var book Book
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

func (o *Data) String() string {
	return fmt.Sprintf("ob-%v-%v-%v", o.Cex, o.Type, o.Symbol)
}

func UpdateOneBookDelta(price, qty float64, oldBook Book, isAsk bool) Book {
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
