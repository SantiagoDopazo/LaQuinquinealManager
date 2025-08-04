package services

import (
	"errors"
	"laquinquenal/models"
	"laquinquenal/repositories"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo: repositories.NewOrderRepository(),
	}
}

func (service *OrderService) CreateOrder(order *models.Order) error {
	if order.OrderNumber == "" || order.ClientName == "" {
		return errors.New("order number and client name are required")
	}

	order.TotalPrice = order.UnitPrice * float64(order.Quantity)

	if order.Status == "" {
		order.Status = "pending"
	}

	if err := service.orderRepo.CreateOrder(order); err != nil {
		return errors.New("failed to create order: " + err.Error())
	}

	return nil
}
