package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type HealthCheckController struct {}

func NewHealthCheckController() *HealthCheckController {
    return &HealthCheckController{}
}

func (controller *HealthCheckController) check(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "La Quinquenal API is up",
	})
}