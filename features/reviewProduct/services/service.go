package services

import (
	reviewproduct "kstyleAPI/features/reviewProduct"
)

type reviewService struct {
	qry reviewproduct.ReviewData
}

func New(data reviewproduct.ReviewData) reviewproduct.ReviewService {
	return &reviewService{
		qry: data,
	}
}

func (rs *reviewService) Insert(newReview reviewproduct.Core) (reviewproduct.Core, error) {
	res, err := rs.qry.Insert(newReview)
	if err != nil {
		return reviewproduct.Core{}, err
	}

	return res, nil
}

func (rs *reviewService) Like(IdReview, IdMember uint) error {
	err := rs.qry.Like(IdReview, IdMember)
	if err != nil {
		return err
	}

	return nil
}

func (rs *reviewService) Unlike(IdReview, IdMember uint) error {
	err := rs.qry.Unlike(IdReview, IdMember)
	if err != nil {
		return err
	}

	return nil
}
