
package main

import (
    "confluent-kafka-go-demo/coffeeshop/api"
    "confluent-kafka-go-demo/coffeeshop/kafka"
    "confluent-kafka-go-demo/config"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    cfg := config.LoadConfig()

    producer, err := kafka.NewProducer(cfg)
    if err != nil {
        log.Fatalf("Failed to create producer: %v", err)
    }
    defer producer.Writer.Close()

    handler := api.NewOrderHandler(producer)

    router := gin.Default()
    router.POST("/order", handler.CreateOrder)
    router.Run(":8080")
}
