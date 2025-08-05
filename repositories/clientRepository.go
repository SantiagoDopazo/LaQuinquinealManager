package repositories

import (
	"laquinquenal/db"
	"laquinquenal/models"
)

type ClientRepository struct{}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{}
}

func (repo *ClientRepository) CreateClient(client *models.Client) error {
	return db.DB.Create(client).Error
}

func (repo *ClientRepository) GetClientByID(id uint) (*models.Client, error) {
	var client models.Client
	err := db.DB.First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (repo *ClientRepository) GetAllClients() ([]models.Client, error) {
	var clients []models.Client
	err := db.DB.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (repo *ClientRepository) UpdateClient(client *models.Client) error {
	return db.DB.Save(client).Error
}

func (repo *ClientRepository) DeleteClient(id uint) error {
	return db.DB.Delete(&models.Client{}, id).Error
}

func (repo *ClientRepository) GetClientByCUIT(cuit string) (*models.Client, error) {
	var client models.Client
	err := db.DB.Where("cuit = ?", cuit).First(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (repo *ClientRepository) GetClientsWithOrderCount() ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	err := db.DB.Table("clients").
		Select("clients.*, COUNT(orders.id) as order_count").
		Joins("LEFT JOIN orders ON clients.id = orders.client_id").
		Group("clients.id").
		Scan(&results).Error

	return results, err
}
