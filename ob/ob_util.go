package ob

import (
	"errors"
)

var (
	ErrOverOb = errors.New("ob: over orderbook")
)

func BQ(book Book, pos int) (PQ, error) {
	if pos > len(book)-1 {
		return nil, ErrOverOb
	}
	return book[pos], nil
}

func Price(o Data) (float64, error) {
	al := len(o.Asks)
	bl := len(o.Bids)
	if al*bl == 0 {
		if al != 0 {
			return o.Asks[0][0], nil
		}
		if bl != 0 {
			return o.Bids[0][0], nil
		}
		return 0, errors.New("ob: no price")
	}
	if len(o.Asks[0]) < 2 {
		return 0, errors.New("ob: length of book unit < 2")
	}
	if len(o.Bids[0]) < 2 {
		return 0, errors.New("ob: length of book unit < 2")
	}
	p := (o.Asks[0][0] + o.Bids[0][0]) / 2
	if p <= 0 {
		return 0, errors.New("ob: price <= 0")
	}
	return p, nil
}

func ToPrice(book Book, targetPrice float64) (asset, quote float64, err error) {
	if len(book) == 0 {
		return
	}
	first := book[0]
	fp, err := first.P()
	if err != nil {
		return
	}
	last := book[len(book)-1]
	lp, err := last.P()
	if err != nil {
		return
	}
	ask := lp >= fp
	if (targetPrice >= fp) != ask {
		return
	}
	for _, v := range book {
		p, q := v[0], v[1]
		if (targetPrice >= p) != ask {
			break
		}
		asset += q
		value := p * q
		quote += value
	}
	return
}

func ToPos(book Book, pos int) (asset, quote float64, err error) {
	for i, pq := range book {
		if i > pos {
			break
		}
		var p, q float64
		p, err = pq.P()
		if err != nil {
			return
		}
		q, err = pq.Q()
		if err != nil {
			return
		}
		asset += q
		quote += p * q
	}
	return
}

func AssetQty(book Book, quoteQty float64) float64 {
	rquo := quoteQty
	assst := 0.0
	for _, v := range book {
		p, q := v[0], v[1]
		if p == 0 {
			continue
		}
		value := p * q
		if rquo <= value {
			q = rquo / p
			rquo = 0
		} else {
			rquo -= value
		}
		assst += q
		if rquo == 0 {
			break
		}
	}
	return assst
}

func QuoteQty(book Book, assetQty float64) float64 {
	rqty := assetQty
	quote := 0.0
	for _, v := range book {
		p, q := v[0], v[1]
		//bp, bq := big.NewFloat(p), big.NewFloat(q)
		if rqty <= q {
			q = rqty
			//rqty = big.NewFloat(0)
			rqty = 0
		} else {
			//rqty.Sub(rqty, bq)
			rqty -= q
		}
		//quote := big.NewFloat(0)
		//quote.Mul(bp, bq)
		//quo.Add(quo, quote)
		quote += p * q
		if rqty == 0 {
			break
		}
	}
	//quote, _ := quo.Float64()
	return quote
}

// AdjustByTakerFee copy raw ob data, and adjust by taker fee
// if fee >= 1, panic
func AdjustByTakerFee(rawOb Data, fee float64) Data {
	if fee >= 1 {
		panic("ob: adjust ob data by taker fee, but fee >= 1")
	}
	o := rawOb.Copy()
	for _, v := range o.Asks {
		addFee := 1 + fee
		v[0] = v[0] * addFee
	}
	for _, v := range o.Bids {
		subFee := 1 - fee
		v[0] = v[0] * subFee
	}
	return *o
}

func AdjustBookByAsset(book Book, assetQty float64) (Book, error) {
	rq := assetQty
	var nb Book
	for _, pq := range book {
		p, err := pq.P()
		if err != nil {
			return nil, err
		}
		q, err := pq.Q()
		if err != nil {
			return nil, err
		}
		nq := q
		//bq := big.NewFloat(q)
		var bre bool
		if rq != q {
			nq = rq
			bre = true
		} else {
			rq -= q
		}
		npq := PQ{p, nq}
		nb = append(nb, npq)
		if bre {
			break
		}
	}
	return nb, nil
}

func AdjustBookByQuote(book Book, quoteQty float64) (Book, error) {
	rq := quoteQty
	var nb Book
	for _, pq := range book {
		p, err := pq.P()
		if err != nil {
			return nil, err
		}
		if p == 0 {
			return nil, errors.New("p is 0")
		}
		q, err := pq.Q()
		if err != nil {
			return nil, err
		}
		v := p * q
		nq := q
		var bre bool
		if rq != v {
			nq = rq / p
			bre = true
		} else {
			rq -= v
		}
		npq := PQ{p, nq}
		nb = append(nb, npq)
		if bre {
			break
		}
	}
	return nb, nil
}

func CutByQty(book Book, qty float64) Book {
	book = book.Copy()
	remainingQty := qty
	iStartCut := -1
	firstPqQty := 0.0
	for i, pq := range book {
		iStartCut = i
		q := pq[1]
		if remainingQty > q {
			remainingQty -= q
			if remainingQty < 0 {
				remainingQty = 0
			}
			continue
		}
		firstPqQty = q - remainingQty
		remainingQty = 0
		if firstPqQty > 0 {
			break
		}
	}
	if remainingQty > 0 || firstPqQty <= 0 || iStartCut < 0 {
		return Book{}
	}
	nb := book[iStartCut:]
	nb[0][1] = firstPqQty
	return nb
}
