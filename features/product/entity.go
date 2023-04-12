package product

import "github.com/labstack/echo/v4"

type Core struct {
	ID          uint    `json:"id"`
	NameProduct string  `json:"name_product"`
	Price       float64 `json:"price"`
}

type ProductHandler interface {
	Insert() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetProducts() echo.HandlerFunc
	GetProductById() echo.HandlerFunc
}

type ProductService interface {
	Insert(newMember Core) (Core, error)
	Update(IdProd uint, updMember Core) (Core, error)
	Delete(id uint) error
	GetProducts() ([]Core, error)
	GetProductById(id uint) (ProductRes, error)
}

type ProductData interface {
	Insert(newMember Core) (Core, error)
	Update(updMember Core) (Core, error)
	Delete(id uint) error
	GetProducts() ([]Core, error)
	GetProductById(id uint) (ProductRes, error)
}
