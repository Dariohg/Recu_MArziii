package services

import (
	"233338-R-C2/src/productos/domain/services"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// EmailMessage representa un mensaje de correo electrónico
type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// MockEmailService es un servicio de correo electrónico que utiliza el servicio de mensajería
type MockEmailService struct {
	messagingService services.MessagingService
}

// NewMockEmailService crea una nueva instancia del servicio de correo electrónico
func NewMockEmailService(messagingService services.MessagingService) *MockEmailService {
	service := &MockEmailService{
		messagingService: messagingService,
	}

	// Suscribirse a la cola de correos para procesarlos
	err := messagingService.SubscribeToQueue("emails", service.processEmailMessage)
	if err != nil {
		log.Printf("Error al suscribirse a la cola de correos: %v", err)
	}

	return service
}

// SendEmail envía un correo electrónico a través del servicio de mensajería
func (e *MockEmailService) SendEmail(to, subject, body string) error {
	message := EmailMessage{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	return e.messagingService.PublishMessage("", "emails", message)
}

// SendProductNotification envía una notificación sobre un producto
func (e *MockEmailService) SendProductNotification(to string, productData map[string]interface{}) error {
	// Extraer información del producto
	productName, _ := productData["nombre"].(string)
	productID, _ := productData["id"].(int)

	subject := fmt.Sprintf("Nuevo producto agregado: %s", productName)
	body := fmt.Sprintf(`
Estimado usuario,

Un nuevo producto ha sido agregado a nuestro catálogo:

ID: %d
Nombre: %s

Este correo es una notificación automática.

Saludos,
El equipo de Productos
`, productID, productName)

	return e.SendEmail(to, subject, body)
}

// processEmailMessage procesa un mensaje de correo electrónico de la cola
func (e *MockEmailService) processEmailMessage(messageBytes []byte) error {
	var email EmailMessage
	if err := json.Unmarshal(messageBytes, &email); err != nil {
		return fmt.Errorf("error al deserializar mensaje de correo: %v", err)
	}

	// Simular el envío real de un correo electrónico
	time.Sleep(1 * time.Second)

	log.Printf("[Servicio de Correo] Correo enviado a: %s", email.To)
	log.Printf("[Servicio de Correo] Asunto: %s", email.Subject)
	log.Printf("[Servicio de Correo] Cuerpo: %s", email.Body)

	return nil
}
