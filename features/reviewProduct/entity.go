package reviewproduct

import "github.com/labstack/echo/v4"

type Core struct {
	ID         uint
	IdProduct  uint
	IdMember   uint
	DescReview string
}

type ReviewHandler interface {
	Insert() echo.HandlerFunc
	Like() echo.HandlerFunc
	Unlike() echo.HandlerFunc
}

type ReviewService interface {
	Insert(newReview Core) (Core, error)
	Like(IdReview, IdMember uint) error
	Unlike(IdReview, IdMember uint) error
}

type ReviewData interface {
	Insert(newReview Core) (Core, error)
	Like(IdReview, IdMember uint) error
	Unlike(IdReview, IdMember uint) error
}
