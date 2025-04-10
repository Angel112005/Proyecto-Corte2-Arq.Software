package config

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ConnectRabbitMQ establece la conexión con RabbitMQ y devuelve el objeto de conexión y canal.
func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	rabbitMQURL := "amqp://agmc:112005@54.157.231.159:5672/"

	// Conectar a RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, nil, fmt.Errorf("❌ Error al conectar con RabbitMQ: %w", err)
	}

	// Abrir canal
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("❌ Error al abrir canal RabbitMQ: %w", err)
	}

	log.Println("✅ Conexión a RabbitMQ establecida correctamente.")
	return conn, ch, nil
}
