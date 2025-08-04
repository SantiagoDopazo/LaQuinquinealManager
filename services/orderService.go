package services

import (
	"errors"
	"laquinquenal/models"
	"laquinquenal/repositories"
	"time"
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

	now := time.Now()
	order.CreatedAt = now
	order.UpdatedAt = now

	if err := service.orderRepo.CreateOrder(order); err != nil {
		return errors.New("failed to create order: " + err.Error())
	}

	return nil
}
