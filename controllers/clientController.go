package controllers

import (
	"net/http"
	"strconv"

	"laquinquenal/models"
	"laquinquenal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClientController struct {
	clientService *services.ClientService
}

func NewClientController() *ClientController {
	return &ClientController{
		clientService: services.NewClientService(),
	}
}

func (ctrl *ClientController) CreateClient(c *gin.Context) {
	var client models.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format: " + err.Error(),
		})
		return
	}

	if err := ctrl.clientService.CreateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Client created successfully",
		"data":    client,
	})
}

func (ctrl *ClientController) GetClient(c *gin.Context) {
	identifier := c.Param("identifier")
	if identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Client identifier is required",
		})
		return
	}

	var client *models.Client
	var err error

	if id, parseErr := strconv.ParseUint(identifier, 10, 32); parseErr == nil {
		client, err = ctrl.clientService.GetClientByID(uint(id))
	} else {
		client, err = ctrl.clientService.GetClientByCUIT(identifier)
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Client not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (ctrl *ClientController) GetAllClients(c *gin.Context) {
	includeOrderCount := c.Query("include_order_count") == "true"

	if includeOrderCount {
		clientsWithCount, err := ctrl.clientService.GetClientsWithOrderCount()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve clients with order count: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, clientsWithCount)
	} else {
		clients, err := ctrl.clientService.GetAllClients()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve clients: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, clients)
	}
}

func (ctrl *ClientController) UpdateClient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid client ID format",
		})
		return
	}

	var updatedClient models.Client
	if err := c.ShouldBindJSON(&updatedClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format: " + err.Error(),
		})
		return
	}

	client, err := ctrl.clientService.UpdateClient(uint(id), &updatedClient)
	if err != nil {
		if err.Error() == "client not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Client not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (ctrl *ClientController) DeleteClient(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid client ID format",
		})
		return
	}

	err = ctrl.clientService.DeleteClient(uint(id))
	if err != nil {
		if err.Error() == "client not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Client not found",
			})
			return
		}
		if err.Error() == "cannot delete client with existing orders" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Cannot delete client with existing orders",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Client deleted successfully",
	})
}
