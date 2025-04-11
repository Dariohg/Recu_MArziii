package application

import (
	"233338-R-C2/src/productos/domain/entities"
	"233338-R-C2/src/productos/domain/services"
	"errors"
)

// NotifyProductAdded es el caso de uso para notificar cuando se agrega un producto
type NotifyProductAdded struct {
	emailService services.EmailService
}

// NewNotifyProductAdded crea una nueva instancia del caso de uso
func NewNotifyProductAdded(emailService services.EmailService) *NotifyProductAdded {
	return &NotifyProductAdded{
		emailService: emailService,
	}
}

// Execute ejecuta el caso de uso para notificar sobre un producto agregado
// usando el correo almacenado en el producto
func (n *NotifyProductAdded) Execute(product *entities.Product) error {
	if product.Email == "" {
		return errors.New("el producto no tiene un correo electrónico asociado")
	}

	// Convertir el producto a un mapa para enviarlo al servicio de correo
	productData := map[string]interface{}{
		"id":        product.ID,
		"nombre":    product.Nombre,
		"precio":    product.Precio,
		"codigo":    product.Codigo,
		"descuento": product.Descuento,
	}

	return n.emailService.SendProductNotification(product.Email, productData)
}
