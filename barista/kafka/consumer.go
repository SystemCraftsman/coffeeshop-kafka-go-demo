package kafka

import (
	"coffeeshop-kafka-go-demo/barista/model"
	"coffeeshop-kafka-go-demo/config"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Consumer struct {
	Consumer *kafka.Consumer
}

func NewConsumer(cfg config.Config) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBootstrapServers,
		"group.id":          "barista-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	if err := c.Subscribe(cfg.CoffeeshopTopic, nil); err != nil {
		return nil, err
	}
	return &Consumer{Consumer: c}, nil
}

func (c *Consumer) ConsumeOrders() {
	for {
		msg, err := c.Consumer.ReadMessage(-1)
		if err == nil {
			var order model.Beverage
			if err := json.Unmarshal(msg.Value, &order); err == nil {
				log.Printf("Preparing drink: %+v", order)
			}
		} else {
			log.Printf("Consumer error: %v", err)
		}
	}
}
