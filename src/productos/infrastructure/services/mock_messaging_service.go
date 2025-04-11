package services

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// MockMessagingService simula un servicio de mensajería como RabbitMQ
type MockMessagingService struct {
	queues         map[string][]func(message []byte) error
	messages       map[string][][]byte
	isConnected    bool
	processingTime time.Duration
	mu             sync.RWMutex
}

// NewMockMessagingService crea una nueva instancia del servicio de mensajería simulado
func NewMockMessagingService() *MockMessagingService {
	return &MockMessagingService{
		queues:         make(map[string][]func(message []byte) error),
		messages:       make(map[string][][]byte),
		isConnected:    true,
		processingTime: 500 * time.Millisecond, // Simular latencia de red
	}
}

// PublishMessage publica un mensaje en una cola o exchange simulado
func (m *MockMessagingService) PublishMessage(exchange, routingKey string, message interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isConnected {
		return fmt.Errorf("servicio de mensajería no conectado")
	}

	// Simular tiempo de procesamiento
	time.Sleep(m.processingTime)

	// Convertir el mensaje a JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error al serializar el mensaje: %v", err)
	}

	// Registrar la publicación
	queueName := routingKey
	if exchange != "" {
		queueName = fmt.Sprintf("%s.%s", exchange, routingKey)
	}

	log.Printf("[MensajeríaSimulada] Mensaje publicado en %s: %s", queueName, string(messageBytes))

	// Almacenar el mensaje
	if _, exists := m.messages[queueName]; !exists {
		m.messages[queueName] = [][]byte{}
	}
	m.messages[queueName] = append(m.messages[queueName], messageBytes)

	// Notificar a los suscriptores
	if handlers, exists := m.queues[queueName]; exists {
		for _, handler := range handlers {
			// Ejecutar handlers en goroutines para simular procesamiento asíncrono
			go func(h func([]byte) error, msg []byte) {
				if err := h(msg); err != nil {
					log.Printf("[MensajeríaSimulada] Error al procesar mensaje: %v", err)
				}
			}(handler, messageBytes)
		}
	}

	return nil
}

// SubscribeToQueue se suscribe a una cola simulada
func (m *MockMessagingService) SubscribeToQueue(queueName string, handler func(message []byte) error) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isConnected {
		return fmt.Errorf("servicio de mensajería no conectado")
	}

	// Registrar el suscriptor
	if _, exists := m.queues[queueName]; !exists {
		m.queues[queueName] = []func(message []byte) error{}
	}
	m.queues[queueName] = append(m.queues[queueName], handler)

	log.Printf("[MensajeríaSimulada] Nuevo suscriptor registrado para la cola %s", queueName)

	// Entregar mensajes pendientes al nuevo suscriptor
	if messages, exists := m.messages[queueName]; exists {
		for _, msg := range messages {
			// Copia del mensaje para evitar problemas de concurrencia
			messageCopy := make([]byte, len(msg))
			copy(messageCopy, msg)

			go func(h func([]byte) error, msg []byte) {
				time.Sleep(m.processingTime) // Simular latencia de red
				if err := h(msg); err != nil {
					log.Printf("[MensajeríaSimulada] Error al procesar mensaje pendiente: %v", err)
				}
			}(handler, messageCopy)
		}
	}

	return nil
}

// Close cierra la conexión simulada
func (m *MockMessagingService) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isConnected {
		return fmt.Errorf("servicio ya desconectado")
	}

	m.isConnected = false
	log.Printf("[MensajeríaSimulada] Conexión cerrada")
	return nil
}

// SetProcessingTime permite configurar el tiempo de procesamiento simulado
func (m *MockMessagingService) SetProcessingTime(duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.processingTime = duration
}
