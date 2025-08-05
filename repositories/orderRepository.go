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

func (repo *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := db.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (repo *OrderRepository) UpdateOrder(order *models.Order) error {
	return db.DB.Save(order).Error
}

func (repo *OrderRepository) DeleteOrder(id uint) error {
	return db.DB.Delete(&models.Order{}, id).Error
}
