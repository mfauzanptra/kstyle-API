package services

import (
	"errors"
	"kstyleAPI/features/product"
	"strings"
)

type productUseCase struct {
	qry product.ProductData
}

func New(pd product.ProductData) product.ProductService {
	return &productUseCase{
		qry: pd,
	}
}

func (puc *productUseCase) Insert(newProduct product.Core) (product.Core, error) {
	res, err := puc.qry.Insert(newProduct)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "duplicated product"
		} else {
			msg = "server error"
		}
		return product.Core{}, errors.New(msg)
	}

	return res, nil
}
func (puc *productUseCase) Update(IdProd uint, updProduct product.Core) (product.Core, error) {
	updProduct.ID = IdProd
	res, err := puc.qry.Update(updProduct)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product not found"
		} else {
			msg = "server error"
		}
		return product.Core{}, errors.New(msg)
	}

	return res, nil
}
func (puc *productUseCase) Delete(id uint) error {
	if err := puc.qry.Delete(id); err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product not found"
		} else {
			msg = "server error"
		}
		return errors.New(msg)
	}

	return nil
}
func (puc *productUseCase) GetProducts() ([]product.Core, error) {
	res, err := puc.qry.GetProducts()
	if err != nil {
		return []product.Core{}, err
	}

	return res, nil
}
func (puc *productUseCase) GetProductById(id uint) (product.ProductRes, error) {
	res, err := puc.qry.GetProductById(id)
	if err != nil {
		return product.ProductRes{}, err
	}

	return res, nil
}
