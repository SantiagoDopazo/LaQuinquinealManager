package routes

import (
    "laquinquenal/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, healthController *controllers.HealthCheckController) {
    router.GET("/health_check", healthController.check)
}