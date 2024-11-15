package kafka

import (
	"context"

	"github.com/Dnlbb/notifier/internal/client/kafka/consumer"
)

type Consumer interface {
	Consume(ctx context.Context, topicName string, handler consumer.Handler) error
	Close() error
}
