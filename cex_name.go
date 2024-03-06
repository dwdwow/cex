package cex

type Name string

const (
	BINANCE Name = "BINANCE"
)

var cexNames = []Name{BINANCE}

func NotCexName(name Name) bool {
	for _, n := range cexNames {
		if n == name {
			return false
		}
	}
	return true
}
