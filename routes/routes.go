package routes

import (
	"laquinquenal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, healthController *controllers.HealthCheckController, orderController *controllers.OrderController) {
	router.GET("/health_check", healthController.Check)

	router.POST("/orders", orderController.CreateOrder)
}
