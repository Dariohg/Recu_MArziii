package infrastructure

import (
	"233338-R-C2/src/productos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type IsNewProductAddedController struct {
	useCase *application.GetLastProduct
}

func NewIsNewProductAddedController(useCase *application.GetLastProduct) *IsNewProductAddedController {
	return &IsNewProductAddedController{useCase: useCase}
}

func (ipc *IsNewProductAddedController) Execute(c *gin.Context) {
	// Obtenemos el último ID conocido (si existe)
	lastKnownIDStr := c.DefaultQuery("lastId", "0")
	lastKnownID, err := strconv.Atoi(lastKnownIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Header("Content-Type", "application/json")

	// Obtener el último producto
	product, err := ipc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		c.JSON(http.StatusOK, gin.H{
			"nuevoProducto": false,
			"mensaje":       "No hay productos disponibles",
			"timestamp":     time.Now(),
		})
		return
	}

	// Verificar si hay un nuevo producto
	if product.ID > lastKnownID {
		c.JSON(http.StatusOK, gin.H{
			"nuevoProducto": true,
			"producto":      product,
			"timestamp":     time.Now(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"nuevoProducto":    false,
			"ultimoProductoId": product.ID,
			"timestamp":        time.Now(),
		})
	}
}
