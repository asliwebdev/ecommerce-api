package models

import "time"

type Cart struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ProductId string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}

type AddProductRequest struct {
	UserId    string `json:"user_id" binding:"required,uuid"`
	ProductId string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type UpdateProductQuantity struct {
	UserId    string `json:"user_id" binding:"required,uuid"`
	ProductId string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	Increase  bool   `json:"increase" binding:"required"`
}
