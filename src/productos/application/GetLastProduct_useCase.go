package application

import (
	"233338-R-C2/src/productos/domain"
	"233338-R-C2/src/productos/domain/entities"
)

type GetLastProduct struct {
	db domain.IProduct
}

func NewGetLastProduct(db domain.IProduct) *GetLastProduct {
	return &GetLastProduct{db: db}
}

func (glp *GetLastProduct) Execute() (*entities.Product, error) {
	return glp.db.ObtenerUltimoProducto()
}
