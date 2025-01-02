package service

import (
	"ecommerce/models"
	"ecommerce/repository"
)

type CartService struct {
	cartRepo *repository.CartRepository
}

func NewCartService(cartRepo *repository.CartRepository) *CartService {
	return &CartService{cartRepo: cartRepo}
}

func (s *CartService) AddProduct(cart *models.AddProductRequest) error {
	return s.cartRepo.AddProductToCart(cart)
}

func (s *CartService) GetCart(userId string) ([]models.Cart, error) {
	return s.cartRepo.GetCartItems(userId)
}

func (s *CartService) RemoveProduct(userId, productId string) error {
	return s.cartRepo.RemoveProductFromCart(userId, productId)
}

func (s *CartService) ClearCart(userId string) error {
	return s.cartRepo.ClearCart(userId)
}

func (s *CartService) UpdateProductQuantity(cart *models.UpdateProductQuantity) error {
	return s.cartRepo.UpdateProductQuantity(cart)
}
