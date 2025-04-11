package adapters

import (
	"log"
	"my-clinic-api/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, ch, err := config.ConnectRabbitMQ()
	if err != nil {
		log.Println("❌ No se pudo inicializar RabbitMQAdapter:", err)
		return nil, err
	}

	return &RabbitMQAdapter{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQAdapter) Publish(queue string, msg []byte) error {
	
	_, err := r.ch.QueueDeclare(
		queue,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	)
	if err != nil {
		return err
	}

	return r.ch.Publish(
		"",    // exchange
		queue, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
}
