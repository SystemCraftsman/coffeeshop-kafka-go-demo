
package main

import (
    "confluent-kafka-go-demo/barista/kafka"
    "confluent-kafka-go-demo/config"
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
