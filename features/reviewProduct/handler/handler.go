package handler

import (
	reviewproduct "kstyleAPI/features/reviewProduct"
	"kstyleAPI/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type reviewControl struct {
	srv reviewproduct.ReviewService
}

func New(srv reviewproduct.ReviewService) reviewproduct.ReviewHandler {
	return &reviewControl{
		srv: srv,
	}
}

func (rc *reviewControl) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := ReviewReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		res, err := rc.srv.Insert(*ToCore(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success insert new review",
		})
	}
}

func (rc *reviewControl) Like() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LikeReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		err := rc.srv.Like(input.IdReview, input.IdMember)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success like review",
		})
	}
}

func (rc *reviewControl) Unlike() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LikeReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		err := rc.srv.Unlike(input.IdReview, input.IdMember)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success unlike review",
		})
	}
}
