package service

import (
	"ecommerce/models"
	"ecommerce/repository"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
}

func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) GetAllOrders(userId string) ([]models.Order, error) {
	return s.orderRepo.GetAllOrders(userId)
}

func (s *OrderService) GetOrderById(orderId string) (models.OrderResponse, error) {
	return s.orderRepo.GetOrderById(orderId)
}

func (s *OrderService) UpdateOrder(req models.UpdateOrderRequest) error {
	return s.orderRepo.UpdateOrder(req)
}

func (s *OrderService) DeleteOrder(orderId string) error {
	return s.orderRepo.DeleteOrder(orderId)
}
