package domain

// IMessageBus define la interfaz para la comunicación con un Message Broker
type IMessageBus interface {
	Publish(queue string, message []byte) error
}
