package main

import (
	"233338-R-C2/src/productos/infrastructure"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	infrastructure.ConfigureProductRoutes(r)

	puerto := os.Getenv("PUERTO")
	if puerto == "" {
		puerto = "8080"
	}

	log.Printf("Servidor iniciando en el puerto %s", puerto)
	if err := r.Run(":" + puerto); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
