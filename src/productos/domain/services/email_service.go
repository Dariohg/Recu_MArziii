package services

// EmailService define el puerto para el servicio de correo electrónico
type EmailService interface {
	// SendEmail envía un correo electrónico
	SendEmail(to, subject, body string) error

	// SendProductNotification envía una notificación específica de producto
	SendProductNotification(to string, productData map[string]interface{}) error
}
