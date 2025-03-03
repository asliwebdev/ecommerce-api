package models

import "time"

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price" binding:"required,gt=0"`
	Stock       int       `json:"stock" binding:"required,gte=0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
