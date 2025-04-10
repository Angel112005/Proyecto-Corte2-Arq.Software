package infrastructure

import (
	"log"
	"my-clinic-api/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQ estructura la conexión y canal
type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// NewRabbitMQ usa la conexión centralizada de `config/`
func NewRabbitMQ() (*RabbitMQ, error) {
	conn, ch, err := config.ConnectRabbitMQ()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

// Publish envía mensajes a la cola de RabbitMQ
func (r *RabbitMQ) Publish(queue string, message []byte) error {
	q, err := r.ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("❌ Error al declarar la cola: %s", err)
		return err
	}

	err = r.ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		log.Printf("❌ Error al publicar mensaje en RabbitMQ: %s", err)
		return err
	}

	log.Printf("📢 Mensaje enviado a RabbitMQ: %s", message)
	return nil
}

// Close cierra la conexión de RabbitMQ
func (r *RabbitMQ) Close() {
	r.ch.Close()
	r.conn.Close()
}
