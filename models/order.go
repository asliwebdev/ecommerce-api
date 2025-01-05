package models

import "time"

type PaymentMethod string

const (
	PaymentMethodCard PaymentMethod = "CARD"
	PaymentMethodCash PaymentMethod = "CASH"
)

type Order struct {
	Id              string        `json:"id"`
	UserId          string        `json:"user_id"`
	TotalPrice      float64       `json:"total_price"`
	Status          string        `json:"status"`
	ShippingAddress string        `json:"shipping_address"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type OrderItem struct {
	Id        string    `json:"id"`
	OrderId   string    `json:"order_id"`
	ProductId string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateOrderParams struct {
	UserId          string        `json:"user_id"`
	OrderId         string        `json:"order_id"`
	TotalPrice      float64       `json:"total_price"`
	Status          string        `json:"status"`
	ShippingAddress string        `json:"shipping_address"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	OrderItems      []OrderItem   `json:"order_items"`
}

type UpdateOrderRequest struct {
	OrderId         string        `json:"order_id"`
	TotalPrice      float64       `json:"total_price"`
	Status          string        `json:"status"`
	ShippingAddress string        `json:"shipping_address"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
}

type OrderResponse struct {
	Order
	Items []OrderItem `json:"items"`
}
