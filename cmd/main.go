package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"cpf-cnpj-validator/internal/app"
	"cpf-cnpj-validator/internal/infrastructure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Arquivo .env não encontrado, usando variáveis do sistema.")
	}

	infrastructure.NewDatabase()

	router := app.SetupRoutes()

	port := "8080"
	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
