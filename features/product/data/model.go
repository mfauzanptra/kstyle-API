package data

import (
	"kstyleAPI/features/product"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	NameProduct string
	Price       float64
}

func ToCore(data Product) product.Core {
	return product.Core{
		ID:          data.ID,
		NameProduct: data.NameProduct,
		Price:       data.Price,
	}
}

func CoreToData(core product.Core) Product {
	return Product{
		ID:          core.ID,
		NameProduct: core.NameProduct,
		Price:       core.Price,
	}
}
