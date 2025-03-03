package infrastructure

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var (
	rabbitInstance *RabbitMQ
	rabbitOnce     sync.Once
)

func NewRabbitMQ() (*RabbitMQ, error) {
	rabbitOnce.Do(func() {
		rabbitUser := os.Getenv("RABBITMQ_USER")
		rabbitPass := os.Getenv("RABBITMQ_PASS")
		rabbitHost := os.Getenv("RABBITMQ_HOST")
		rabbitPort := os.Getenv("RABBITMQ_PORT")

		amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUser, rabbitPass, rabbitHost, rabbitPort)

		conn, err := amqp.Dial(amqpURL)
		if err != nil {
			log.Printf("❌ Failed to connect to RabbitMQ: %v", err)
			return
		}

		channel, err := conn.Channel()
		if err != nil {
			log.Printf("❌ Failed to open a channel: %v", err)
			return
		}

		rabbitInstance = &RabbitMQ{
			Connection: conn,
			Channel:    channel,
		}
	})

	if rabbitInstance == nil {
		return nil, fmt.Errorf("❌ failed to initialize RabbitMQ")
	}

	return rabbitInstance, nil
}

func (r *RabbitMQ) DeclareQueue(queueName string) error {
	_, err := r.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (r *RabbitMQ) PublishMessage(queueName, message string) error {
	if r.Channel == nil {
		return fmt.Errorf("❌ RabbitMQ channel is not initialized")
	}
	return r.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Connection != nil {
		r.Connection.Close()
	}
}
