package service

import (
	"ecommerce/models"
	"ecommerce/repository"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CheckoutService struct {
	cartRepo    *repository.CartRepository
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
}

func NewCheckoutService(cartRepo *repository.CartRepository, orderRepo *repository.OrderRepository, productRepo *repository.ProductRepository) *CheckoutService {
	return &CheckoutService{
		cartRepo:    cartRepo,
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

var (
	ErrCartIsEmpty          = errors.New("cart is empty")
	ErrFailedToFetchCart    = errors.New("failed to fetch cart items")
	ErrFailedToFetchProduct = errors.New("failed to fetch product details")
	ErrFailedToCreateOrder  = errors.New("failed to create order")
	ErrFailedToClearCart    = errors.New("failed to clear cart")
)

func (s *CheckoutService) Checkout(req models.CheckoutRequest) (models.CheckoutResponse, error) {
	cartItems, err := s.cartRepo.GetCartItems(req.UserId)
	if err != nil {
		return models.CheckoutResponse{}, ErrFailedToFetchCart
	}

	if len(cartItems) == 0 {
		return models.CheckoutResponse{}, ErrCartIsEmpty
	}

	productIds := make([]string, len(cartItems))
	for i, cartItem := range cartItems {
		productIds[i] = cartItem.ProductId
	}

	products, err := s.productRepo.GetProductsByIds(productIds)
	if err != nil {
		return models.CheckoutResponse{}, ErrFailedToFetchProduct
	}

	productMap := make(map[string]models.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	var totalPrice float64
	orderItems := make([]models.OrderItem, 0, len(cartItems))

	for _, cartItem := range cartItems {
		product, exists := productMap[cartItem.ProductId]
		if !exists {
			return models.CheckoutResponse{}, fmt.Errorf("product %s not found", cartItem.ProductId)
		}

		totalPrice += float64(cartItem.Quantity) * product.Price
		orderItems = append(orderItems, models.OrderItem{
			Id:        uuid.NewString(),
			ProductId: cartItem.ProductId,
			Quantity:  cartItem.Quantity,
			Price:     product.Price,
		})
	}

	orderId := uuid.NewString()
	orderParams := models.CreateOrderParams{
		UserId:          req.UserId,
		OrderId:         orderId,
		TotalPrice:      totalPrice,
		Status:          req.Status,
		ShippingAddress: req.ShippingAddress,
		PaymentMethod:   req.PaymentMethod,
		OrderItems:      orderItems,
	}

	if err := s.orderRepo.CreateOrder(orderParams); err != nil {
		return models.CheckoutResponse{}, ErrFailedToCreateOrder
	}

	if err := s.cartRepo.ClearCart(req.UserId); err != nil {
		return models.CheckoutResponse{}, ErrFailedToClearCart
	}

	return models.CheckoutResponse{
		OrderId:    orderId,
		TotalPrice: totalPrice,
	}, nil
}
