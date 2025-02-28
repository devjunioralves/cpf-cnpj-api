package main

import (
	"log"

	"github.com/joho/godotenv"

	"cpf-cnpj-api/internal/app"
	"cpf-cnpj-api/internal/domain/services"
	"cpf-cnpj-api/internal/infrastructure"
	"cpf-cnpj-api/internal/infrastructure/repositories"
	"cpf-cnpj-api/internal/presentation"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, using system vars.")
	}

	database := infrastructure.NewDatabase()
	db := database.GetDB()

	repo := repositories.NewCpfCnpjRepositoryGorm(db)
	cpfService := services.NewCpfCnpjService(repo)
	cpfHandler := presentation.NewCPFHandler(cpfService)

	router := app.SetupRoutes(cpfHandler)
	server := app.NewHttpServer(router)
	server.Start("8080")
}
