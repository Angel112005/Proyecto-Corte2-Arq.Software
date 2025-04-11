
package config

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	url := "amqp://agmc:112005@54.157.231.159:5672/"
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, fmt.Errorf("❌ Error al conectar con RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("❌ Error al abrir canal: %w", err)
	}

	log.Println("✅ Conectado a RabbitMQ correctamente.")
	return conn, ch, nil
}
