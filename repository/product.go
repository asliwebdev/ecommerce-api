package repository

import (
	"database/sql"
	"fmt"

	"ecommerce/models"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (p *ProductRepository) CreateProduct(product *models.Product) error {
	id := uuid.NewString()

	_, err := p.DB.Exec("INSERT INTO products (id, name, description, price, stock) VALUES ($1, $2, $3, $4, $5)",
		id, product.Name, product.Description, product.Price, product.Stock)
	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}

	return nil
}

func (p *ProductRepository) GetAllProducts() ([]models.Product, error) {
	rows, err := p.DB.Query("SELECT id, name, description, price, stock, created_at, updated_at FROM products")
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through products: %w", err)
	}

	return products, nil
}

func (p *ProductRepository) GetProductById(id string) (*models.Product, error) {
	var product models.Product
	err := p.DB.QueryRow("SELECT id, name, description, price, stock, created_at, updated_at FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found: %w", err)
		}
		return nil, fmt.Errorf("error fetching product by ID: %w", err)
	}

	return &product, nil
}

func (r *ProductRepository) GetProductsByIds(ids []string) ([]models.Product, error) {
	query := `
		SELECT id, name, description, price, stock, created_at, updated_at 
		FROM products 
		WHERE id = ANY($1)`
	rows, err := r.DB.Query(query, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductRepository) UpdateProduct(id string, product *models.Product) error {
	_, err := p.DB.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock = $4, updated_at = $5 WHERE id = $6",
		product.Name, product.Description, product.Price, product.Stock, product.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}

func (p *ProductRepository) DeleteProduct(id string) error {
	_, err := p.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}

func (p *ProductRepository) ProductExists(productId string) (bool, error) {
	var exists bool
	err := p.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM products WHERE id = $1)`, productId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking product existence: %w", err)
	}
	return exists, nil
}
