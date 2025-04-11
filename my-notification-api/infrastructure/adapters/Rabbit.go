
package adapters

import (
	"log"
	config "my-notification-api/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// Constructor
func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, ch, err := config.ConnectRabbitMQ()
	if err != nil {
		log.Println("❌ No se pudo conectar a RabbitMQ:", err)
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

func (r *RabbitMQAdapter) Publish(queue string, msg []byte) error {
	_, err := r.ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	return r.ch.Publish(
		"", queue, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
}
