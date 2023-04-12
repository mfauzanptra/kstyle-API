package handler

import (
	"kstyleAPI/features/member"
	"kstyleAPI/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type memberControl struct {
	srv member.MemberService
}

func New(srv member.MemberService) member.MemberHandler {
	return &memberControl{
		srv: srv,
	}
}

func (mc *memberControl) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := InsertMemberReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		res, err := mc.srv.Insert(*ToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "duplicated") {
				return c.JSON(http.StatusConflict, helper.ErrorResponse("username already exists"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success insert new member",
		})
	}
}
func (mc *memberControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := InsertMemberReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		IdMember := c.Param("id")
		cnv, err := strconv.Atoi(IdMember)
		if err != nil {
			log.Println("update post param error")
			return c.JSON(http.StatusBadRequest, "wrong url parameter")
		}

		res, err := mc.srv.Update(uint(cnv), *ToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusConflict, helper.ErrorResponse("member does not exist"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success update member data",
		})
	}
}
func (mc *memberControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		IdMember := c.Param("id")
		cnv, err := strconv.Atoi(IdMember)
		if err != nil {
			log.Println("update post param error")
			return c.JSON(http.StatusBadRequest, "wrong url parameter")
		}

		err = mc.srv.Delete(uint(cnv))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusConflict, helper.ErrorResponse("member does not exist"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete member data",
		})
	}
}
func (mc *memberControl) GetMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := mc.srv.GetMembers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success get members",
		})
	}
}
