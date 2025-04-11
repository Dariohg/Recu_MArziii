package infrastructure

import (
	"233338-R-C2/src/productos/application"
	"233338-R-C2/src/productos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddProductController struct {
	useCase *application.AddProduct
}

func NewAddProductController(useCase *application.AddProduct) *AddProductController {
	return &AddProductController{useCase: useCase}
}

func (apc *AddProductController) Execute(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if product.Nombre == "" || product.Codigo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y código son campos obligatorios"})
		return
	}

	if err := apc.useCase.Execute(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
