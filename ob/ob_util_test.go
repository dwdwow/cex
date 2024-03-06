package ob

import (
	"fmt"
	"testing"
)

func TestAssetQty(t *testing.T) {
	book := Book{
		{1.1, 0.3},
		{1.2, 0.5},
		{1.3, 1.1},
	}
	q := AssetQty(book, 0.9)
	fmt.Println(q)
}

func TestQuoteQty(t *testing.T) {
	book := Book{
		{1.1, 0.3},
		{1.2, 0.5},
		{1.3, 1.1},
	}
	q := QuoteQty(book, 0.9)
	fmt.Println(q)
}

func TestCutByQty(t *testing.T) {
	book := Book{
		{1.1, 0.3},
		{1.2, 0.5},
		{1.3, 1.1},
	}
	nb := CutByQty(book, 0.4)
	fmt.Println(book)
	fmt.Println(nb)
}
