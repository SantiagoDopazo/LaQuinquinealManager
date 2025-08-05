package repositories

import (
	"laquinquenal/db"
	"laquinquenal/models"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	return db.DB.Create(order).Error
}

func (repo *OrderRepository) GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
	err := db.DB.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
