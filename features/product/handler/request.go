package handler

import "kstyleAPI/features/product"

type ProdReq struct {
	NameProduct string  `json:"name_product"`
	Price       float64 `json:"price"`
}

func ToCore(data ProdReq) *product.Core {
	return &product.Core{
		NameProduct: data.NameProduct,
		Price:       data.Price,
	}
}
