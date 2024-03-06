package ob

import (
	"encoding/json"
	"time"

	"github.com/dwdwow/cex"
)

func Empty(cexName cex.Name, obType cex.PairType, symbol string) Data {
	return Data{
		Cex:     cexName,
		Type:    obType,
		Symbol:  symbol,
		Version: "",
		Time:    time.Now().UnixMilli(),
		Asks:    nil,
		Bids:    nil,
		Empty:   true,
	}
}

func RedisMsgUnmarshal(payload string) (Data, error) {
	o := new(Data)
	err := json.Unmarshal([]byte(payload), o)
	return *o, err
}
