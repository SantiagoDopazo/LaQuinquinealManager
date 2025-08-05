package services

import (
	"errors"
	"laquinquenal/models"
	"laquinquenal/repositories"
	"regexp"
	"strings"
)

type ClientService struct {
	clientRepo *repositories.ClientRepository
}

func NewClientService() *ClientService {
	return &ClientService{
		clientRepo: repositories.NewClientRepository(),
	}
}

func (service *ClientService) CreateClient(client *models.Client) error {
	if client.Name == "" {
		return errors.New("client name is required")
	}

	if client.Cuit == "" {
		return errors.New("client CUIT is required")
	}

	client.Cuit = strings.ReplaceAll(client.Cuit, "-", "")
	client.Cuit = strings.ReplaceAll(client.Cuit, " ", "")

	if !service.isValidCUIT(client.Cuit) {
		return errors.New("invalid CUIT format")
	}

	existingClient, _ := service.clientRepo.GetClientByCUIT(client.Cuit)
	if existingClient != nil {
		return errors.New("client with this CUIT already exists")
	}

	client.Name = strings.TrimSpace(client.Name)
	client.Contact = strings.TrimSpace(client.Contact)
	client.Address = strings.TrimSpace(client.Address)
	client.Email = strings.TrimSpace(client.Email)
	client.Phone = strings.TrimSpace(client.Phone)

	if client.Email != "" && !service.isValidEmail(client.Email) {
		return errors.New("invalid email format")
	}

	if err := service.clientRepo.CreateClient(client); err != nil {
		return errors.New("failed to create client: " + err.Error())
	}

	return nil
}

func (service *ClientService) GetClientByID(id uint) (*models.Client, error) {
	return service.clientRepo.GetClientByID(id)
}

func (service *ClientService) GetAllClients() ([]models.Client, error) {
	return service.clientRepo.GetAllClients()
}

func (service *ClientService) UpdateClient(id uint, updatedClient *models.Client) (*models.Client, error) {
	existingClient, err := service.clientRepo.GetClientByID(id)
	if err != nil {
		return nil, errors.New("client not found")
	}

	if updatedClient.Name == "" {
		return nil, errors.New("client name is required")
	}

	if updatedClient.Cuit == "" {
		return nil, errors.New("client CUIT is required")
	}

	updatedClient.Cuit = strings.ReplaceAll(updatedClient.Cuit, "-", "")
	updatedClient.Cuit = strings.ReplaceAll(updatedClient.Cuit, " ", "")

	if !service.isValidCUIT(updatedClient.Cuit) {
		return nil, errors.New("invalid CUIT format")
	}

	if updatedClient.Cuit != existingClient.Cuit {
		existingCUITClient, _ := service.clientRepo.GetClientByCUIT(updatedClient.Cuit)
		if existingCUITClient != nil {
			return nil, errors.New("client with this CUIT already exists")
		}
	}

	updatedClient.ID = existingClient.ID

	updatedClient.Name = strings.TrimSpace(updatedClient.Name)
	updatedClient.Contact = strings.TrimSpace(updatedClient.Contact)
	updatedClient.Address = strings.TrimSpace(updatedClient.Address)
	updatedClient.Email = strings.TrimSpace(updatedClient.Email)
	updatedClient.Phone = strings.TrimSpace(updatedClient.Phone)

	if updatedClient.Email != "" && !service.isValidEmail(updatedClient.Email) {
		return nil, errors.New("invalid email format")
	}

	updatedClient.CreatedAt = existingClient.CreatedAt

	if err := service.clientRepo.UpdateClient(updatedClient); err != nil {
		return nil, errors.New("failed to update client: " + err.Error())
	}

	return updatedClient, nil
}

func (service *ClientService) DeleteClient(id uint) error {
	client, err := service.clientRepo.GetClientByID(id)
	if err != nil {
		return errors.New("client not found")
	}

	if len(client.Orders) > 0 {
		return errors.New("cannot delete client with existing orders")
	}

	if err := service.clientRepo.DeleteClient(id); err != nil {
		return errors.New("failed to delete client: " + err.Error())
	}

	return nil
}

func (service *ClientService) GetClientByCUIT(cuit string) (*models.Client, error) {
	cuit = strings.ReplaceAll(cuit, "-", "")
	cuit = strings.ReplaceAll(cuit, " ", "")

	return service.clientRepo.GetClientByCUIT(cuit)
}

func (service *ClientService) GetClientsWithOrderCount() ([]map[string]interface{}, error) {
	return service.clientRepo.GetClientsWithOrderCount()
}

func (service *ClientService) isValidCUIT(cuit string) bool {
	if len(cuit) != 11 {
		return false
	}

	matched, _ := regexp.MatchString(`^\d{11}$`, cuit)
	return matched
}

func (service *ClientService) isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
