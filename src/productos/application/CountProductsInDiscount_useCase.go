package application

import (
	"233338-R-C2/src/productos/domain"
)

type CountProductsInDiscount struct {
	db domain.IProduct
}

func NewCountProductsInDiscount(db domain.IProduct) *CountProductsInDiscount {
	return &CountProductsInDiscount{db: db}
}

func (cpd *CountProductsInDiscount) Execute() (int, error) {
	return cpd.db.ContarProductosConDescuento()
}
