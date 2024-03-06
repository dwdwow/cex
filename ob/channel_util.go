package ob

import (
	"errors"
	"strings"
)

type ChannelMarshal func(Data) (string, error)
type ChannelUnmarshal func(string) (Data, error)

type ChannelUtil struct {
	marshal   ChannelMarshal
	unmarshal ChannelUnmarshal
}

func NewChannelUtil(marshal ChannelMarshal, unmarshal ChannelUnmarshal) *ChannelUtil {
	return &ChannelUtil{
		marshal:   marshal,
		unmarshal: unmarshal,
	}
}

func (u *ChannelUtil) Marshal(data Data) (string, error) {
	if u.marshal == nil {
		return "", errors.New("marshal is nil")
	}
	return u.marshal(data)
}

func (u *ChannelUtil) Unmarshal(s string) (Data, error) {
	if u.unmarshal == nil {
		return Data{}, errors.New("unmarshal is nil")
	}
	return u.unmarshal(s)
}

func SimplePublisherChannelMarshal(d Data) (string, error) {
	return ID(d.Cex, d.Type, d.Symbol)
}

func SimplePublisherChannelUnmarshal(c string) (Data, error) {
	cn, ot, syb, err := ParseID(c)
	od := Data{
		Cex:    cn,
		Type:   ot,
		Symbol: syb,
	}
	return od, err
}

func NewSimplePublisherChannelUtil() *ChannelUtil {
	return NewChannelUtil(SimplePublisherChannelMarshal, SimplePublisherChannelUnmarshal)
}

const RedisPublisherChannelPrefix = "redis_ob_channel:"

func RedisPublisherChannelMarshal(d Data) (string, error) {
	id, err := ID(d.Cex, d.Type, d.Symbol)
	if err != nil {
		return "", err
	}
	return RedisPublisherChannelPrefix + id, nil
}

func RedisPublisherChannelUnmarshal(c string) (Data, error) {
	split := strings.Split(c, RedisPublisherChannelPrefix)
	if len(split) != 2 {
		return Data{}, errors.New("ob channel util: invalid channel " + c)
	}
	id := split[1]
	cn, ot, syb, err := ParseID(id)
	od := Data{
		Cex:    cn,
		Type:   ot,
		Symbol: syb,
	}
	return od, err
}

func NewRedisPublisherChannelUtil() *ChannelUtil {
	return NewChannelUtil(RedisPublisherChannelMarshal, RedisPublisherChannelUnmarshal)
}
