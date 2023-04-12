package handler

import (
	"kstyleAPI/features/product"
	"kstyleAPI/helper"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type productControl struct {
	srv product.ProductService
}

func New(srv product.ProductService) product.ProductHandler {
	return &productControl{
		srv: srv,
	}
}

func (pc *productControl) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := ProdReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		res, err := pc.srv.Insert(*ToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "duplicated") {
				return c.JSON(http.StatusConflict, helper.ErrorResponse("product already exists"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success insert new product",
		})
	}
}
func (pc *productControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := ProdReq{}

		if err := c.Bind(&input); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong input"))
		}

		IdProd := c.Param("id")
		cnv, err := strconv.Atoi(IdProd)
		if err != nil {
			log.Println("update post param error")
			return c.JSON(http.StatusBadRequest, "wrong url parameter")
		}

		res, err := pc.srv.Update(uint(cnv), *ToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusConflict, helper.ErrorResponse("product does not exist"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success update product data",
		})
	}
}
func (pc *productControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		IdProd := c.Param("id")
		cnv, err := strconv.Atoi(IdProd)
		if err != nil {
			log.Println("update post param error")
			return c.JSON(http.StatusBadRequest, "wrong url parameter")
		}

		err = pc.srv.Delete(uint(cnv))
		if err != nil {
			if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusConflict, helper.ErrorResponse("product does not exist"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete product data",
		})
	}
}
func (pc *productControl) GetProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := pc.srv.GetProducts()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal server error"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success get products",
		})
	}
}
func (pc *productControl) GetProductById() echo.HandlerFunc {
	return func(c echo.Context) error {
		IdProd := c.Param("id")
		cnv, err := strconv.Atoi(IdProd)
		if err != nil {
			log.Println("update post param error")
			return c.JSON(http.StatusBadRequest, "wrong url parameter")
		}

		res, err := pc.srv.GetProductById(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("internal error"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success get product by id",
		})
	}
}
