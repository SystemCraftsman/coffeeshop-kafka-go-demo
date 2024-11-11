package model

type Beverage struct {
	beverage string
	customer string
	orderId  string
}

func NewBeverage(order *Order) *Beverage {
	return &Beverage{beverage: order.Product, customer: order.Name, orderId: order.OrderID}
}
