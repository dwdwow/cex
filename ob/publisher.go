package ob

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/dwdwow/spub"
	"github.com/redis/go-redis/v9"
)

type Publisher struct {
	spub.Publisher[Data]
	producer *Producer
	logger   *slog.Logger
}

func NewPublisher(publisher spub.Publisher[Data], msgHandler CexWsMsgHandler, logger *slog.Logger) *Publisher {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	logger = logger.With("publisher", fmt.Sprintf("%v/%v", msgHandler.Name(), msgHandler.Type()))
	return &Publisher{Publisher: publisher, producer: NewProducer(msgHandler, publisher, logger), logger: logger}
}

func NewSimplePublisher(msgHandler CexWsMsgHandler, logger *slog.Logger) *Publisher {
	publisher := spub.NewSimplePublisher[Data](NewSimplePublisherChannelUtil(), spub.SimpleRcvCapOption[Data](100))
	return NewPublisher(publisher, msgHandler, logger)
}

func NewRedisPublisher(rOpts *redis.Options, msgHandler CexWsMsgHandler, logger *slog.Logger) *Publisher {
	publisher := spub.NewRedisPublisher[Data](rOpts, NewRedisPublisherChannelUtil(), RedisMsgUnmarshal, logger, spub.RedisRcvCapOption[Data](100))
	return NewPublisher(publisher, msgHandler, logger)
}

func NewRedisConsumer(rOpts *redis.Options, logger *slog.Logger) spub.ConsumerService[Data] {
	return spub.NewRedisConsumer[Data](rOpts, NewRedisPublisherChannelUtil(), RedisMsgUnmarshal, logger, spub.RedisConsumerRcvCapOption[Data](100))
}

func (p *Publisher) Start(ctx context.Context) error {
	err := p.Publisher.Start(ctx)
	if err != nil {
		return err
	}
	err = p.producer.Start(ctx)
	if err != nil {
		return err
	}
	return nil
}
