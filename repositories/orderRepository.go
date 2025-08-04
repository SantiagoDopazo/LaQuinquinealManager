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
