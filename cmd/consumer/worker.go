package main

import (
	"cpf-cnpj-api/internal/infrastructure"
	"cpf-cnpj-api/internal/infrastructure/repositories"
	"log"
	"strconv"
)

func main() {
	rabbitMQ, channel, err := infrastructure.NewRabbitMQ()
	if err != nil {
		log.Fatalf("❌ failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()
	defer channel.Close()

	_, err = channel.QueueDeclare(
		"block_cpf_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("❌ failed to declare queue: %v", err)
	}

	database := infrastructure.NewDatabase()
	db := database.GetDB()
	cpfRepo := repositories.NewCpfCnpjRepositoryGorm(db)

	msgs, err := channel.Consume(
		"block_cpf_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("❌ failed to consume messages: %v", err)
	}

	for msg := range msgs {
		cpfId := string(msg.Body)
		id, err := strconv.ParseUint(cpfId, 10, 32)
		if err != nil {
			log.Printf("❌ Error converting CPF/CNPJ ID %s to uint: %v", cpfId, err)
			continue
		}

		err = cpfRepo.BlockCpfCnpjById(uint(id))
		if err != nil {
			log.Printf("❌ Error blocking CPF/CNPJ with ID %s: %v", cpfId, err)
		} else {
			log.Printf("❌ Blocked CPF/CNPJ with ID %s", cpfId)
		}
	}
}
