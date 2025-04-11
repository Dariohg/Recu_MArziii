package infrastructure

import (
	"233338-R-C2/src/productos/application"
	"233338-R-C2/src/productos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddProductController struct {
	useCase       *application.AddProduct
	notifyUseCase *application.NotifyProductAdded
}

func NewAddProductController(useCase *application.AddProduct, notifyUseCase *application.NotifyProductAdded) *AddProductController {
	return &AddProductController{
		useCase:       useCase,
		notifyUseCase: notifyUseCase,
	}
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

	// Verificar si se debe enviar notificación
	shouldNotify := c.Query("notify") == "true"

	notifyEmail := c.Query("email")
	if shouldNotify && notifyEmail == "" {
		notifyEmail = "admin@ejemplo.com"
	}

	if err := apc.useCase.Execute(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enviar notificación si se solicitó
	if shouldNotify && apc.notifyUseCase != nil {
		go func() {
			if err := apc.notifyUseCase.Execute(&product, notifyEmail); err != nil {
				// Solo registrar el error, no afecta la respuesta principal
				c.Error(err)
			}
		}()
	}

	c.JSON(http.StatusCreated, gin.H{
		"producto":   product,
		"notificado": shouldNotify,
	})
}
