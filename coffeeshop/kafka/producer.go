package kafka

import (
	"coffeeshop-kafka-go-demo/coffeeshop/model"
	"coffeeshop-kafka-go-demo/config"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"
)

type Producer struct {
	Writer *kafka.Writer
}

func NewProducer(cfg config.Config) (*Producer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{cfg.KafkaBootstrapServers},
		Topic:        cfg.CoffeeshopTopic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
	})
	return &Producer{Writer: writer}, nil
}

func (p *Producer) ProduceOrder(order model.Order) error {
	value, _ := json.Marshal(order)
	msg := kafka.Message{
		Value: value,
	}
	return p.Writer.WriteMessages(context.Background(), msg)
}
