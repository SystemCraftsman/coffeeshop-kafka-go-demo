package model

type Order struct {
	Product string `json:"product"`
	Name    string `json:"name"`
	OrderID string `json:"orderId"`
}
