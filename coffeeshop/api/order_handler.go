
package api

import (
    "confluent-kafka-go-demo/coffeeshop/model"
    "confluent-kafka-go-demo/coffeeshop/kafka"
    "net/http"
    "github.com/gin-gonic/gin"
)

type OrderHandler struct {
    Producer *kafka.Producer
}

func NewOrderHandler(producer *kafka.Producer) *OrderHandler {
    return &OrderHandler{Producer: producer}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var order model.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.Producer.ProduceOrder(order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce order"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "Order placed successfully"})
}
