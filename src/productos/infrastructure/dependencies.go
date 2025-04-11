package infrastructure

import (
	"233338-R-C2/src/productos/application"
	infraServices "233338-R-C2/src/productos/infrastructure/services"
	"github.com/gin-gonic/gin"
)

func ConfigureProductRoutes(r *gin.Engine) {

	mysql := NewMySQL()

	// Servicios
	bcryptService := infraServices.NewBcryptService()
	messagingService := infraServices.NewMockMessagingService()

	// Configurar el servicio de correo basado en el servicio de mensajería
	emailService := infraServices.NewMockEmailService(messagingService)

	// Casos de uso
	addProduct := application.NewAddProduct(mysql, bcryptService)
	getLastProduct := application.NewGetLastProduct(mysql)
	countProductsInDiscount := application.NewCountProductsInDiscount(mysql)
	listProduct := application.NewListProduct(mysql)
	notifyProductAdded := application.NewNotifyProductAdded(emailService)

	// Controladores
	addProductController := NewAddProductController(addProduct, notifyProductAdded)
	isNewProductAddedController := NewIsNewProductAddedController(getLastProduct)
	countProductsInDiscountController := NewCountProductsInDiscountController(countProductsInDiscount)
	listProductController := NewListProductController(listProduct)

	// Rutas
	api := r.Group("/api")
	{
		api.POST("/addProducto", addProductController.Execute)
		api.GET("/isNewProductAdded", isNewProductAddedController.Execute)
		api.GET("/countProductsInDiscount", countProductsInDiscountController.Execute)
		api.GET("/productos", listProductController.Execute)
	}
}
