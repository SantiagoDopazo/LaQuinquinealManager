package server

import (
	"laquinquenal/controllers"
	"laquinquenal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine

	healthController *controllers.HealthCheckController
	orderController  *controllers.OrderController

	port string
}

func NewServer() *Server {
	server := &Server{
		port: ":8080",
	}

	server.initializeControllers()
	server.initializeRoutes()

	return server
}

func (server *Server) initializeControllers() {
	server.healthController = controllers.NewHealthCheckController()
	server.orderController = controllers.NewOrderController()
}

func (server *Server) initializeRoutes() {
	server.router = gin.Default()

	routes.SetupRoutes(server.router, server.healthController, server.orderController)
}

func (server *Server) Start() error {
	log.Println("Starting La Quinquenal API server on", server.port)
	return server.router.Run(server.port)
}

func (server *Server) SetPort(port string) {
	server.port = port
}
