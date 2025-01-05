package models

type CheckoutRequest struct {
	UserId          string        `json:"user_id" binding:"required,uuid"`
	Status          string        `json:"status"`
	ShippingAddress string        `json:"shipping_address" binding:"required"`
	PaymentMethod   PaymentMethod `json:"payment_method" binding:"required"`
}

type CheckoutResponse struct {
	OrderId    string  `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
}
