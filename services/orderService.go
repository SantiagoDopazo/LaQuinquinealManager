package services

import (
	"errors"
	"laquinquenal/models"
	"laquinquenal/repositories"
)

type OrderService struct {
	orderRepo  *repositories.OrderRepository
	clientRepo *repositories.ClientRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo:  repositories.NewOrderRepository(),
		clientRepo: repositories.NewClientRepository(),
	}
}

func (service *OrderService) CreateOrder(order *models.Order) error {
	if order.OrderNumber == "" {
		return errors.New("order number is required")
	}

	if order.ClientID == 0 {
		return errors.New("client ID is required")
	}

	_, err := service.clientRepo.GetClientByID(order.ClientID)
	if err != nil {
		return errors.New("client not found")
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

func (service *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	return service.orderRepo.GetOrderByID(id)
}

func (service *OrderService) GetAllOrders() ([]models.Order, error) {
	return service.orderRepo.GetAllOrders()
}

func (service *OrderService) UpdateOrder(id uint, updatedOrder *models.Order) (*models.Order, error) {
	existingOrder, err := service.orderRepo.GetOrderByID(id)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if updatedOrder.OrderNumber == "" {
		return nil, errors.New("order number is required")
	}

	if updatedOrder.ClientID == 0 {
		return nil, errors.New("client ID is required")
	}

	_, err = service.clientRepo.GetClientByID(updatedOrder.ClientID)
	if err != nil {
		return nil, errors.New("client not found")
	}

	updatedOrder.ID = existingOrder.ID

	updatedOrder.TotalPrice = updatedOrder.UnitPrice * float64(updatedOrder.Quantity)

	if updatedOrder.Status == "" {
		updatedOrder.Status = existingOrder.Status
	}

	updatedOrder.CreatedAt = existingOrder.CreatedAt

	if err := service.orderRepo.UpdateOrder(updatedOrder); err != nil {
		return nil, errors.New("failed to update order: " + err.Error())
	}

	return updatedOrder, nil
}

func (service *OrderService) DeleteOrder(id uint) error {
	_, err := service.orderRepo.GetOrderByID(id)
	if err != nil {
		return errors.New("order not found")
	}

	if err := service.orderRepo.DeleteOrder(id); err != nil {
		return errors.New("failed to delete order: " + err.Error())
	}

	return nil
}
