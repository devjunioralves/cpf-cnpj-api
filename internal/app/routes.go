package app

import (
	"cpf-cnpj-api/internal/domain/services"
	"cpf-cnpj-api/internal/presentation"
	"cpf-cnpj-api/internal/presentation/middlewares"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	startTime          = time.Now()
	requestCount int64 = 0
)

func SetupRoutes(authHandler *presentation.AuthHandler, cpfHandler *presentation.CPFHandler, authService *services.AuthService) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.MetricMiddleware())

	router.GET("/metrics", middlewares.MetricsHandler())

	router.Use(func(c *gin.Context) {
		atomic.AddInt64(&requestCount, 1)
		c.Next()
	})

	router.GET("/status", func(c *gin.Context) {
		uptime := time.Since(startTime).String()
		c.JSON(200, gin.H{
			"uptime":        uptime,
			"request_count": atomic.LoadInt64(&requestCount),
		})
	})

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
