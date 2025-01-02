package repository

import (
	"database/sql"
	"ecommerce/models"
	"fmt"

	"github.com/google/uuid"
)

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepo(db *sql.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) AddProductToCart(cart *models.AddProductRequest) error {
	var existingQuantity int
	err := r.DB.QueryRow("SELECT quantity FROM carts WHERE user_id = $1 AND product_id = $2", cart.UserId, cart.ProductId).Scan(&existingQuantity)

	if err == sql.ErrNoRows {
		cartId := uuid.NewString()
		_, err := r.DB.Exec("INSERT INTO carts (id, user_id, product_id, quantity) VALUES ($1, $2, $3, $4)",
			cartId, cart.UserId, cart.ProductId, cart.Quantity)
		if err != nil {
			return fmt.Errorf("error adding product %s to user %s's cart: %w", cart.ProductId, cart.UserId, err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking if product exists in cart for user %s: %w", cart.UserId, err)
	} else {
		newQuantity := existingQuantity + cart.Quantity
		_, err := r.DB.Exec("UPDATE carts SET quantity = $1 WHERE user_id = $2 AND product_id = $3",
			newQuantity, cart.UserId, cart.ProductId)
		if err != nil {
			return fmt.Errorf("error updating quantity for product %s in user %s's cart: %w", cart.ProductId, cart.UserId, err)
		}
	}

	return nil
}

func (r *CartRepository) GetCartItems(userId string) ([]models.Cart, error) {
	rows, err := r.DB.Query("SELECT id, user_id, product_id, quantity, created_at FROM carts WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching cart items for user %s: %w", userId, err)
	}
	defer rows.Close()

	var cartItems []models.Cart
	for rows.Next() {
		var cartItem models.Cart
		if err := rows.Scan(&cartItem.Id, &cartItem.UserId, &cartItem.ProductId, &cartItem.Quantity, &cartItem.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning cart item: %w", err)
		}
		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through cart items for user %s: %w", userId, err)
	}

	return cartItems, nil
}

func (r *CartRepository) RemoveProductFromCart(userId, productId string) error {
	_, err := r.DB.Exec("DELETE FROM carts WHERE user_id = $1 AND product_id = $2", userId, productId)
	if err != nil {
		return fmt.Errorf("error removing product %s from cart: %w", productId, err)
	}
	return nil
}

func (r *CartRepository) ClearCart(userId string) error {
	_, err := r.DB.Exec("DELETE FROM carts WHERE user_id = $1", userId)
	if err != nil {
		return fmt.Errorf("error clearing cart for user with ID: %s, %w", userId, err)
	}
	return nil
}

func (r *CartRepository) UpdateProductQuantity(cart *models.UpdateProductQuantity) error {
	var existingQuantity int
	err := r.DB.QueryRow("SELECT quantity FROM carts WHERE user_id = $1 AND product_id = $2", cart.UserId, cart.ProductId).Scan(&existingQuantity)

	if err == sql.ErrNoRows {
		return fmt.Errorf("product %s not found in user %s's cart", cart.ProductId, cart.UserId)
	} else if err != nil {
		return fmt.Errorf("error checking current quantity for product %s in user %s's cart: %w", cart.ProductId, cart.UserId, err)
	}

	var newQuantity int
	if cart.Increase {
		newQuantity = existingQuantity + cart.Quantity
	} else {
		if existingQuantity-cart.Quantity <= 0 {
			return r.RemoveProductFromCart(cart.UserId, cart.ProductId)
		}
		newQuantity = existingQuantity - cart.Quantity
	}

	_, err = r.DB.Exec("UPDATE carts SET quantity = $1 WHERE user_id = $2 AND product_id = $3", newQuantity, cart.UserId, cart.ProductId)
	if err != nil {
		return fmt.Errorf("error updating quantity for product %s in user %s's cart: %w", cart.ProductId, cart.UserId, err)
	}

	return nil
}
