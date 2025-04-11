package infrastructure

import (
	"233338-R-C2/src/productos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CountProductsInDiscountController struct {
	useCase *application.CountProductsInDiscount
}

func NewCountProductsInDiscountController(useCase *application.CountProductsInDiscount) *CountProductsInDiscountController {
	return &CountProductsInDiscountController{useCase: useCase}
}

func (cpdc *CountProductsInDiscountController) Execute(c *gin.Context) {
	// Obtenemos el último conteo conocido (si existe)
	lastKnownCountStr := c.DefaultQuery("lastCount", "-1")
	lastKnownCount, err := strconv.Atoi(lastKnownCountStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor de conteo inválido"})
		return
	}

	// Configurar timeout y ticker para el long polling
	timeout := time.After(30 * time.Second)          // Timeout después de 30 segundos
	ticker := time.NewTicker(500 * time.Millisecond) // Verificar cada 500ms
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			// Si se alcanza el timeout, respondemos con el último conteo conocido
			count, err := cpdc.useCase.Execute()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"mensaje":   "Timeout alcanzado sin cambios",
				"count":     count,
				"timestamp": time.Now(),
				"cambio":    count != lastKnownCount,
			})
			return

		case <-ticker.C:
			// Cada tick, verificamos si hay cambios en el conteo
			currentCount, err := cpdc.useCase.Execute()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Si hay un cambio o es la primera consulta, respondemos inmediatamente
			if currentCount != lastKnownCount || lastKnownCount == -1 {
				c.JSON(http.StatusOK, gin.H{
					"count":     currentCount,
					"timestamp": time.Now(),
					"cambio":    lastKnownCount != -1, // No indicamos cambio si es la primera consulta
				})
				return
			}
			// Si no hay cambios, continuamos esperando
		}
	}
}
