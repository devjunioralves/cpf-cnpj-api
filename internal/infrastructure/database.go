package infrastructure

import (
	"fmt"
	"log"
	"os"

	"cpf-cnpj-api/internal/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", ""),
		getEnv("DB_USER", ""),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_NAME", ""),
		getEnv("DB_PORT", ""),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error on connect database: %v", err)
	}

	fmt.Println("📦 Successfully on connect database!")
	database := &Database{DB: db}

	database.Migrate()

	return database
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}

func (d *Database) Migrate() {
	err := d.DB.AutoMigrate(&models.CpfCnpj{})
	if err != nil {
		log.Fatalf("Error on running migrations: %v", err)
	}
	fmt.Println("✅ Successfully on run migrations!")
}
