package repository

import (
	"database/sql"
	"ecommerce/models"
	"fmt"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(params models.CreateOrderParams) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec(
		`INSERT INTO orders (id, user_id, total_price, status, shipping_address, payment_method) 
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		params.OrderId, params.UserId, params.TotalPrice, params.Status, params.ShippingAddress, params.PaymentMethod,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create order: %w", err)
	}

	for _, item := range params.OrderItems {
		_, err = tx.Exec(
			"INSERT INTO order_items (id, order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4, $5)",
			item.Id, params.OrderId, item.ProductId, item.Quantity, item.Price,
		)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create order item: %w", err)
		}
	}

	return tx.Commit()
}

func (r *OrderRepository) GetOrderById(orderId string) (models.OrderResponse, error) {
	var order models.OrderResponse

	row := r.DB.QueryRow(
		`SELECT id, user_id, total_price, status, shipping_address, payment_method, created_at, updated_at 
		 FROM orders WHERE id = $1`,
		orderId,
	)
	err := row.Scan(&order.Id, &order.UserId, &order.TotalPrice, &order.Status, &order.ShippingAddress, &order.PaymentMethod, &order.CreatedAt, &order.UpdatedAt)
	if err == sql.ErrNoRows {
		return order, fmt.Errorf("order not found")
	} else if err != nil {
		return order, fmt.Errorf("failed to fetch order: %w", err)
	}

	rows, err := r.DB.Query("SELECT id, product_id, quantity, price FROM order_items WHERE order_id = $1", orderId)
	if err != nil {
		return order, fmt.Errorf("failed to fetch order items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		err := rows.Scan(&item.Id, &item.ProductId, &item.Quantity, &item.Price)
		if err != nil {
			return order, fmt.Errorf("failed to scan order item: %w", err)
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

func (r *OrderRepository) GetAllOrders(userId string) ([]models.Order, error) {
	orders := []models.Order{}

	rows, err := r.DB.Query(
		`SELECT id, total_price, status, shipping_address, payment_method, created_at, updated_at 
		 FROM orders WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return orders, fmt.Errorf("failed to fetch orders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.Id, &order.TotalPrice, &order.Status, &order.ShippingAddress, &order.PaymentMethod, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return orders, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepository) UpdateOrder(params models.UpdateOrderRequest) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec(
		`UPDATE orders 
		 SET total_price = $1, status = $2, shipping_address = $3, payment_method = $4, updated_at = CURRENT_TIMESTAMP 
		 WHERE id = $5`,
		params.TotalPrice, params.Status, params.ShippingAddress, params.PaymentMethod, params.OrderId,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update order: %w", err)
	}
	return tx.Commit()
}

func (r *OrderRepository) DeleteOrder(orderId string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	_, err = tx.Exec("DELETE FROM order_items WHERE order_id = $1", orderId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete order items: %w", err)
	}

	_, err = tx.Exec("DELETE FROM orders WHERE id = $1", orderId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return tx.Commit()
}
