package routes

import (
	"laquinquenal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, healthController *controllers.HealthCheckController, orderController *controllers.OrderController, clientController *controllers.ClientController) {
	router.GET("/health_check", healthController.Check)

	// Client routes
	router.GET("/clients", clientController.GetAllClients) // Con filtros opcionales: ?include_order_count=true
	router.POST("/clients", clientController.CreateClient)
	router.GET("/clients/:identifier", clientController.GetClient) // Por ID numérico o CUIT
	router.PUT("/clients/:id", clientController.UpdateClient)
	router.DELETE("/clients/:id", clientController.DeleteClient)

	// Order routes
	router.GET("/orders", orderController.GetAllOrders)
	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders/:id", orderController.GetOrderByID)
	router.PUT("/orders/:id", orderController.UpdateOrder)
	router.DELETE("/orders/:id", orderController.DeleteOrder)
}
