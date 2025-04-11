package services

// MessagingService define el puerto para los servicios de mensajería
type MessagingService interface {
	// PublishMessage publica un mensaje en una cola o exchange
	PublishMessage(exchange, routingKey string, message interface{}) error

	// SubscribeToQueue se suscribe a una cola para recibir mensajes
	SubscribeToQueue(queueName string, handler func(message []byte) error) error

	// Close cierra la conexión con el servicio de mensajería
	Close() error
}
