package service

import (
	"errors"
	"time"

	"ecommerce/models"
	"ecommerce/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.productRepo.CreateProduct(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAllProducts()
}

func (s *ProductService) GetProductById(id string) (*models.Product, error) {
	return s.productRepo.GetProductById(id)
}

func (s *ProductService) UpdateProduct(id string, product *models.Product) error {
	exists, err := s.productRepo.ProductExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("product with ID " + id + " not found")
	}

	product.UpdatedAt = time.Now()

	return s.productRepo.UpdateProduct(id, product)
}

func (s *ProductService) DeleteProduct(id string) error {
	exists, err := s.productRepo.ProductExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("product with ID " + id + " not found")
	}

	return s.productRepo.DeleteProduct(id)
}
