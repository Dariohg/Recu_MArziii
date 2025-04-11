package application

import (
	"233338-R-C2/src/productos/domain"
	"233338-R-C2/src/productos/domain/entities"
	"233338-R-C2/src/productos/domain/services"
)

type AddProduct struct {
	db                domain.IProduct
	encryptionService services.EncryptionService
	emailService      services.EmailService // Añadimos el servicio de correo
}

func NewAddProduct(
	db domain.IProduct,
	encryptionService services.EncryptionService,
	emailService services.EmailService, // Añadimos el servicio de correo
) *AddProduct {
	return &AddProduct{
		db:                db,
		encryptionService: encryptionService,
		emailService:      emailService, // Inicializamos el servicio de correo
	}
}

func (ap *AddProduct) Execute(product *entities.Product) error {
	// Encriptar el código del producto
	encryptedCode, err := ap.encryptionService.Encrypt(product.Codigo)
	if err != nil {
		return err
	}

	product.Codigo = encryptedCode

	// Guardar el producto en la base de datos
	err = ap.db.Guardar(product)
	if err != nil {
		return err
	}

	// Si el producto tiene un correo electrónico, enviar la notificación inmediatamente
	if product.Email != "" && ap.emailService != nil {
		// Convertir el producto a un mapa para enviarlo al servicio de correo
		productData := map[string]interface{}{
			"id":        product.ID,
			"nombre":    product.Nombre,
			"precio":    product.Precio,
			"codigo":    product.Codigo,
			"descuento": product.Descuento,
		}

		// Enviar la notificación por correo
		// No maneja el error para no interrumpir el flujo principal si el correo falla
		_ = ap.emailService.SendProductNotification(product.Email, productData)
	}

	return nil
}
