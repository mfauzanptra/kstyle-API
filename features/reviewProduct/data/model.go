package data

import (
	member "kstyleAPI/features/member/data"
	product "kstyleAPI/features/product/data"
	reviewproduct "kstyleAPI/features/reviewProduct"
)

type ReviewProduct struct {
	ID         uint `gorm:"primaryKey"`
	IdProduct  uint
	Product    product.Product `gorm:"foreignKey:IdProduct"`
	IdMember   uint
	Member     member.Member `gorm:"foreignKey:IdMember"`
	DescReview string
}

type LikeReview struct {
	IdReview      uint
	ReviewProduct ReviewProduct `gorm:"foreignKey:IdReview"`
	IdMember      uint
	Member        member.Member `gorm:"foreignKey:IdMember"`
}

func CoreToData(core reviewproduct.Core) ReviewProduct {
	return ReviewProduct{
		ID:         core.ID,
		IdProduct:  core.IdProduct,
		IdMember:   core.IdMember,
		DescReview: core.DescReview,
	}
}
