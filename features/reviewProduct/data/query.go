package data

import (
	reviewproduct "kstyleAPI/features/reviewProduct"
	"log"

	"gorm.io/gorm"
)

type reviewData struct {
	db *gorm.DB
}

func New(db *gorm.DB) reviewproduct.ReviewData {
	return &reviewData{
		db: db,
	}
}

func (rd *reviewData) Insert(newReview reviewproduct.Core) (reviewproduct.Core, error) {
	cnv := CoreToData(newReview)
	err := rd.db.Create(&cnv).Error
	if err != nil {
		log.Println("error create new review: ", err)
		return reviewproduct.Core{}, err
	}

	newReview.ID = cnv.ID

	return newReview, nil
}

func (rd *reviewData) Like(IdReview, IdMember uint) error {
	like := LikeReview{
		IdReview: IdReview,
		IdMember: IdMember,
	}
	err := rd.db.Create(&like).Error
	if err != nil {
		return err
	}

	return nil
}

func (rd *reviewData) Unlike(IdReview, IdMember uint) error {
	like := LikeReview{}
	err := rd.db.Where("id_review = ? AND id_member = ?", IdReview, IdMember).Delete(&like).Error
	if err != nil {
		return err
	}

	return nil
}
