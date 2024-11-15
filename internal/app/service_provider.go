package app

import (
	"context"
	"log"

	"github.com/Dnlbb/notifier/internal/client/kafka"
	"github.com/Dnlbb/notifier/internal/client/kafka/consumer"
	"github.com/Dnlbb/notifier/internal/config"
	"github.com/Dnlbb/notifier/internal/service"
	"github.com/Dnlbb/notifier/internal/service/consumer/sender"
	"github.com/Dnlbb/platform_common/pkg/closer"
	"github.com/IBM/sarama"
)

type serviceProvider struct {
	kafkaConsumerConf config.KafkaConsumerConfig
	senderConf        config.Sender

	senderConsumer service.ConsumerService

	consumer             kafka.Consumer
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler *consumer.GroupHandler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) KafkaConsumerConfig() config.KafkaConsumerConfig {
	if s.kafkaConsumerConf == nil {
		cfg, err := config.NewKafkaConsumerConfig()
		if err != nil {
			log.Fatalf("failed to get kafka consumer config: %s", err.Error())
		}

		s.kafkaConsumerConf = cfg
	}

	return s.kafkaConsumerConf
}

func (s *serviceProvider) SenderConfig() config.Sender {
	if s.senderConf == nil {
		cfg, err := config.NewSenderConf()
		if err != nil {
			log.Fatalf("failed to get kafka sender config: %s", err.Error())
		}

		s.senderConf = cfg
	}

	return s.senderConf
}

func (s *serviceProvider) SenderConsumer(ctx context.Context) service.ConsumerService {
	if s.senderConsumer == nil {
		s.senderConsumer = sender.NewService(s.Consumer(), s.SenderConfig())
	}

	return s.senderConsumer
}

func (s *serviceProvider) Consumer() kafka.Consumer {
	if s.consumer == nil {
		s.consumer = consumer.NewConsumer(
			s.ConsumerGroup(),
			s.ConsumerGroupHandler(),
		)
		closer.Add(s.consumer.Close)
	}

	return s.consumer
}

func (s *serviceProvider) ConsumerGroup() sarama.ConsumerGroup {
	if s.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			s.KafkaConsumerConfig().Brokers(),
			s.KafkaConsumerConfig().GroupID(),
			s.KafkaConsumerConfig().Config(),
		)
		if err != nil {
			log.Fatalf("failed to create consumer group: %v", err)
		}

		s.consumerGroup = consumerGroup
	}

	return s.consumerGroup
}

func (s *serviceProvider) ConsumerGroupHandler() *consumer.GroupHandler {
	if s.consumerGroupHandler == nil {
		s.consumerGroupHandler = consumer.NewGroupHandler()
	}

	return s.consumerGroupHandler
}
