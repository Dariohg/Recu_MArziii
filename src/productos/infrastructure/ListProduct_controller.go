package infrastructure

import (
	"233338-R-C2/src/productos/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListProductController struct {
	useCase *application.ListProduct
}

func NewListProductController(useCase *application.ListProduct) *ListProductController {
	return &ListProductController{useCase: useCase}
}

func (lpc *ListProductController) Execute(c *gin.Context) {
	products, err := lpc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
