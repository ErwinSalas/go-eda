package messaging

// MessagePublisher define la interfaz genérica para publicar mensajes.
type MessagePublisher interface {
	Publish(message string) error
}

// MessageSubscriber define la interfaz genérica para suscribirse a eventos/mensajes.
type MessageSubscriber interface {
	Subscribe(endpoint string, protocol string) error
}
