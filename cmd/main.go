package main

import (
	"cpf-cnpj-api/internal/app"
	"cpf-cnpj-api/internal/domain/services"
	"cpf-cnpj-api/internal/infrastructure"
	"cpf-cnpj-api/internal/infrastructure/repositories"
	"cpf-cnpj-api/internal/presentation"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, using system vars.")
	}

	database := infrastructure.NewDatabase()
	db := database.GetDB()

	userRepo := repositories.NewUserRepositoryGorm(db)
	cpfRepo := repositories.NewCpfCnpjRepositoryGorm(db)

	authService := services.NewAuthService(userRepo, os.Getenv("JWT_SECRET"))
	cpfService := services.NewCpfCnpjService(cpfRepo)

	authHandler := presentation.NewAuthHandler(authService)

	rabbitMQ, err := infrastructure.NewRabbitMQ()
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}
	defer func() {
		if err := rabbitMQ.Channel.Close(); err != nil {
			log.Printf("⚠️ Error closing RabbitMQ connection: %v", err)
		}
	}()

	cpfHandler := presentation.NewCPFHandler(cpfService, rabbitMQ)

	router := app.SetupRoutes(authHandler, cpfHandler, authService)

	server := app.NewHttpServer(router)
	server.Start("8080")

}
