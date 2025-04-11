package application

import (
	"233338-R-C2/src/productos/domain"
	"233338-R-C2/src/productos/domain/entities"
)

type ListProduct struct {
	db domain.IProduct
}

func NewListProduct(db domain.IProduct) *ListProduct {
	return &ListProduct{db: db}
}

func (lp *ListProduct) Execute() ([]entities.Product, error) {
	return lp.db.ObtenerTodos()
}
