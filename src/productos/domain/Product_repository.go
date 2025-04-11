package domain

import "233338-R-C2/src/productos/domain/entities"

type IProduct interface {
	Guardar(*entities.Product) error
	ObtenerTodos() ([]entities.Product, error)
	ObtenerPorId(id int) (*entities.Product, error)
	ObtenerUltimoProducto() (*entities.Product, error)
	ContarProductosConDescuento() (int, error)
}
