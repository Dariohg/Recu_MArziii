package entities

import "time"

type Product struct {
	ID            int       `json:"id"`
	Nombre        string    `json:"nombre"`
	Precio        int       `json:"precio"`
	Codigo        string    `json:"codigo"`
	Descuento     bool      `json:"descuento"`
	Email         string    `json:"email"`
	FechaCreacion time.Time `json:"fecha_creacion,omitempty"`
}
