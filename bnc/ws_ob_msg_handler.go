package bnc

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dwdwow/cex"
	"github.com/dwdwow/cex/ob"
	"github.com/dwdwow/props"
	"github.com/dwdwow/ws/wsclt"
	"github.com/gorilla/websocket"
)

type WsSpObMsgHandler struct {
	mgClt            *wsclt.MergedClient
	svDataCacheBySyb props.SafeRWMap[string, []WsDepthMsg]
	gettingObSybs    props.SafeRWMap[string, bool]
	obCacheBySyb     props.SafeRWMap[string, ob.Data]
}

func NewWsSpObMsgHandler(logger *slog.Logger) *WsSpObMsgHandler {
	mgClt := wsclt.
		NewMergedClient(WsBaseUrl, true, maxTopicNumPerWs, logger).
		SetTopicSuber(topicSuber).
		SetTopicUnsuber(topicUnsuber).
		SetPong(pong)
	return &WsSpObMsgHandler{
		mgClt: mgClt,
	}
}

func (w *WsSpObMsgHandler) Name() cex.Name {
	return cex.BINANCE
}

func (w *WsSpObMsgHandler) Type() cex.PairType {
	return cex.PairTypeSpot
}

func (w *WsSpObMsgHandler) Client() *wsclt.MergedClient {
	return w.mgClt
}

func (w *WsSpObMsgHandler) Topics(symbols ...string) []string {
	var topics []string
	for _, s := range symbols {
		topics = append(topics, CreateObTopic(s))
	}
	return topics
}

func (w *WsSpObMsgHandler) Handle(msg wsclt.MergedClientMsg) ([]ob.Data, error) {
	return w.handle(msg)
}

func (w *WsSpObMsgHandler) handle(msg wsclt.MergedClientMsg) ([]ob.Data, error) {
	if msg.Err != nil {
		// set ob data to empty
		var obs []ob.Data
		topics := msg.Client.Topics()
		for _, topic := range topics {
			topicSplit := strings.Split(topic, "@depth")
			if len(topicSplit) != 2 {
				// should not get here
				fmt.Println("unexpect: binance ob ws msg handle: can not parse topic", topic)
				continue
			}
			symbol := topicSplit[0]
			empty := ob.Empty(cex.BINANCE, cex.PairTypeSpot, symbol)
			empty.EmptyReason = "ws error, " + msg.Err.Error()
			w.obCacheBySyb.SetKV(symbol, empty)
			obs = append(obs, empty)
		}
		return obs, nil
	}
	if msg.MsgType != websocket.TextMessage {
		return nil, fmt.Errorf("binance: ws receive unknown msg type %v", msg.MsgType)
	}
	msgData := msg.Data
	data := new(WsDepthMsg)
	err := json.Unmarshal(msgData, data)
	if err != nil {
		return nil, fmt.Errorf("binance: ws msg unmarshal, msg: %v, %w", string(msgData), err)
	}
	if data.EventType == WsEDepthUpdate {
		obData := w.update(*data)
		return []ob.Data{obData}, nil
	}
	return nil, nil
}

func (w *WsSpObMsgHandler) update(depthData WsDepthMsg) ob.Data {
	symbol := depthData.Symbol
	err := w.cacheRawData(depthData)
	if err != nil {
		return w.setEmpty(symbol, err.Error())
	}
	if w.needQueryOb(symbol) {
		err = w.queryOb(symbol)
		if err != nil {
			return w.setEmpty(symbol, err.Error())
		}
	}
	o := w.updateOb(depthData)
	w.obCacheBySyb.SetKV(symbol, o)
	return o
}

func (w *WsSpObMsgHandler) setEmpty(symbol, reason string) ob.Data {
	empty := ob.Empty(cex.BINANCE, cex.PairTypeSpot, symbol)
	empty.EmptyReason = reason
	w.obCacheBySyb.SetKV(symbol, empty)
	return empty
}

func (w *WsSpObMsgHandler) cacheRawData(depthData WsDepthMsg) error {
	symbol := depthData.Symbol
	oldCache := w.svDataCacheBySyb.GetV(symbol)
	if len(oldCache) > 100 {
		// clear cache
		w.svDataCacheBySyb.SetKV(symbol, nil)
		return errors.New("server data cache > 1000")
	}
	newCache := append(oldCache, depthData)
	sort.Slice(newCache, func(i, j int) bool {
		iLastId := newCache[i].LastId
		jFirstId := newCache[j].FirstId
		return iLastId+1 == jFirstId
	})
	w.svDataCacheBySyb.SetKV(symbol, newCache)
	return nil
}

func (w *WsSpObMsgHandler) needQueryOb(symbol string) bool {
	obData, ok := w.obCacheBySyb.GetVWithOk(symbol)
	return !ok || obData.Empty
}

func (w *WsSpObMsgHandler) queryOb(symbol string) error {
	oldCache := w.svDataCacheBySyb.GetV(symbol)
	if len(oldCache) < 10 {
		return errors.New("cache len < 10")
	}
	if time.Now().UnixMilli()-lastObQueryFailTsMilli.Get() < 3000 {
		return errors.New("can not query orderbook within 3000 milliseconds")
	}
	if w.gettingObSybs.SetKV(symbol, true) {
		return errors.New("lock to query orderbook")
	}
	defer w.gettingObSybs.SetKV(symbol, false)
	// because ws orderbook default limit is 1000
	// so limit must be 1000
	rawOrderbook, err := QuerySpotOrderBook(symbol, 1000)
	if err != nil {
		lastObQueryFailTsMilli.Set(time.Now().UnixMilli())
		return err
	}
	obData := ob.Data{
		Cex:     cex.BINANCE,
		Type:    cex.PairTypeSpot,
		Symbol:  symbol,
		Version: strconv.FormatInt(rawOrderbook.LastUpdateId, 10),
		Time:    time.Now().UnixMilli(),
		Asks:    rawOrderbook.Asks,
		Bids:    rawOrderbook.Bids,
		Empty:   false,
	}
	w.obCacheBySyb.SetKV(symbol, obData)
	return nil
}

func (w *WsSpObMsgHandler) updateOb(depthData WsDepthMsg) ob.Data {
	symbol := depthData.Symbol
	buffer := w.svDataCacheBySyb.GetV(symbol)
	empty := ob.Empty(cex.BINANCE, cex.PairTypeSpot, symbol)
	obData, ok := w.obCacheBySyb.GetVWithOk(symbol)
	if !ok || obData.Empty {
		empty.EmptyReason = "unexpect: binance update ob: if !ok || obData.Empty {}"
		return empty
	}
	currentVersion, err := strconv.ParseInt(obData.Version, 10, 64)
	if err != nil {
		empty.EmptyReason = fmt.Sprintln("can not parse ob data version", obData.Version, "err:", err)
		return empty
	}
	if buffer[0].FirstId > currentVersion+1 {
		empty.EmptyReason = fmt.Sprintln("current ob version is small")
		return empty
	}
	lastIndex := 0
	_id := int64(0)
	for i, _depthData := range buffer {
		firstId := _depthData.FirstId
		lastId := _depthData.LastId
		if _id > 0 {
			if firstId == _id+1 {
				_id = lastId
			} else {
				safeMapObDataBuffer.SetKV(symbol, buffer[i:])
				return empty
			}
		} else {
			_id = lastId
		}
		if firstId != currentVersion+1 {
			if firstId <= currentVersion+1 && lastId >= currentVersion+1 {
				// update firstly
			} else {
				if firstId > currentVersion+1 {
					lastIndex = i - 1
					break
				}
				if firstId < currentVersion+1 {
					lastIndex = i
					continue
				}
			}
		}
		lastIndex = i
		asks := _depthData.Asks
		bids := _depthData.Bids
		currentVersion = lastId
		for _, ask := range asks {
			price, err := strconv.ParseFloat(ask[0], 64)
			if err != nil {
				empty.EmptyReason = fmt.Sprintln("can not parse ask price", ask[0], "err:", err)
				return empty
			}
			qty, err := strconv.ParseFloat(ask[1], 64)
			if err != nil {
				empty.EmptyReason = fmt.Sprintln("can not parse ask qty", ask[1], "err:", err)
				return empty
			}
			err = obData.UpdateAskDeltas(ob.Book{{price, qty}}, strconv.FormatInt(currentVersion, 10))
			if err != nil {
				empty.EmptyReason = fmt.Sprintln("can not update ask deltas, err:", err)
				return empty
			}
		}
		for _, bid := range bids {
			price, err := strconv.ParseFloat(bid[0], 64)
			if err != nil {
				empty.EmptyReason = fmt.Sprintln("can not parse bid price", bid[0], "err", err)
				return empty
			}
			qty, err := strconv.ParseFloat(bid[1], 64)
			if err != nil {
				empty.EmptyReason = fmt.Sprintln("can not parse bid qty", bid[1], "err", err)
				return empty
			}
			err = obData.UpdateBidDeltas(ob.Book{{price, qty}}, strconv.FormatInt(currentVersion, 10))
			if err != nil {
				empty.EmptyReason = fmt.Sprintln("can not update bid deltas, err:", err)
				return empty
			}
		}
	}
	if len(buffer) <= lastIndex+1 {
		w.svDataCacheBySyb.SetKV(symbol, []WsDepthMsg{})
	} else {
		w.svDataCacheBySyb.SetKV(symbol, buffer[lastIndex+1:])
	}
	return obData
}
