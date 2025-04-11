package application

import (
	"233338-R-C2/src/productos/domain"
	"233338-R-C2/src/productos/domain/entities"
	"233338-R-C2/src/productos/domain/services"
)

type AddProduct struct {
	db                domain.IProduct
	encryptionService services.EncryptionService
}

func NewAddProduct(db domain.IProduct, encryptionService services.EncryptionService) *AddProduct {
	return &AddProduct{
		db:                db,
		encryptionService: encryptionService,
	}
}

func (ap *AddProduct) Execute(product *entities.Product) error {
	encryptedCode, err := ap.encryptionService.Encrypt(product.Codigo)
	if err != nil {
		return err
	}

	product.Codigo = encryptedCode

	err = ap.db.Guardar(product)

	return err
}
