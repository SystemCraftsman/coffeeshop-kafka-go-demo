package main

import (
    "coffeeshop-kafka-go-demo/barista/config"
    "coffeeshop-kafka-go-demo/barista/kafka"
    "log"
)

func main() {
	cfg := config.LoadConfig()

	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Consumer.Close()

	consumer.ConsumeOrders()
}
