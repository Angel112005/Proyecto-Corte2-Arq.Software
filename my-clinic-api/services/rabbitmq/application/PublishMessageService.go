package application

import (
	"log"

	"my-clinic-api/services/rabbitmq/domain"
)

// PublishMessageService maneja el envío de mensajes a RabbitMQ
type PublishMessageService struct {
	messageBus domain.IMessageBus
}

// NewPublishMessageService crea una nueva instancia de servicio
func NewPublishMessageService(bus domain.IMessageBus) *PublishMessageService {
	return &PublishMessageService{messageBus: bus}
}

// Execute envía un mensaje a la cola especificada
func (p *PublishMessageService) Execute(queue string, message []byte) {
	err := p.messageBus.Publish(queue, message)
	if err != nil {
		log.Printf("❌ Error al publicar mensaje: %s", err)
	}
}
