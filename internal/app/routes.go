package app

import (
	"cpf-cnpj-api/internal/domain/services"
	"cpf-cnpj-api/internal/presentation"
	"cpf-cnpj-api/internal/presentation/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(authHandler *presentation.AuthHandler, cpfHandler *presentation.CPFHandler, authService *services.AuthService) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())

	router.GET("/metrics", middlewares.MetricsHandler())

	api := router.Group("/api")
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)

	apiAuthenticated := api.Group("/cpf-cnpj")
	apiAuthenticated.Use(middlewares.AuthMiddleware(authService))
	{
		apiAuthenticated.POST("/", cpfHandler.CreateCPF)
		apiAuthenticated.GET("/", cpfHandler.GetAllCPF)
		apiAuthenticated.POST("/block", cpfHandler.BlockCPF)
	}

	return router
}
