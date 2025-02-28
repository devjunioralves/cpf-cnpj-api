package app

import (
	"cpf-cnpj-api/internal/presentation"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(handler *presentation.CPFHandler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/cpf-cnpj")
	{
		api.POST("/", handler.CreateCPF)
		api.GET("/", handler.GetAllCPF)
	}

	return router
}
