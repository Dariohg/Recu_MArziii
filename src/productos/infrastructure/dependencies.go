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

	// Casos de uso
	addProduct := application.NewAddProduct(mysql, bcryptService)
	getLastProduct := application.NewGetLastProduct(mysql)
	countProductsInDiscount := application.NewCountProductsInDiscount(mysql)
	listProduct := application.NewListProduct(mysql)

	// Controladores
	addProductController := NewAddProductController(addProduct)
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
