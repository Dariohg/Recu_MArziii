package infrastructure

import (
	"233338-R-C2/src/productos/application"
	"233338-R-C2/src/productos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
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

// isValidEmail valida el formato del correo electrónico
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
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

	// Verificar si se proporcionó un correo electrónico válido
	if product.Email != "" && !isValidEmail(product.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El formato del correo electrónico no es válido"})
		return
	}

	// Verificar si se debe enviar notificación
	shouldNotify := c.Query("notify") == "true" || product.Email != ""

	// Guardar el producto en la base de datos
	if err := apc.useCase.Execute(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Enviar notificación si se debe hacer
	notificationSent := false
	if shouldNotify && apc.notifyUseCase != nil && product.Email != "" {
		go func() {
			err := apc.notifyUseCase.Execute(&product)
			if err != nil {
				c.Error(err)
			}
		}()
		notificationSent = true
	}

	c.JSON(http.StatusCreated, gin.H{
		"producto":             product,
		"notificacion_enviada": notificationSent,
	})
}
