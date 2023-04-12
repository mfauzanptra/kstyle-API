package main

import (
	"kstyleAPI/config"
	"log"

	md "kstyleAPI/features/member/data"
	mh "kstyleAPI/features/member/handler"
	ms "kstyleAPI/features/member/services"

	pd "kstyleAPI/features/product/data"
	ph "kstyleAPI/features/product/handler"
	ps "kstyleAPI/features/product/services"

	rd "kstyleAPI/features/reviewProduct/data"
	rh "kstyleAPI/features/reviewProduct/handler"
	rs "kstyleAPI/features/reviewProduct/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	memberData := md.New(db)
	memberService := ms.New(memberData)
	memberHandler := mh.New(memberService)

	productData := pd.New(db)
	productService := ps.New(productData)
	productHandler := ph.New(productService)

	reviewData := rd.New(db)
	reviewService := rs.New(reviewData)
	reviewHandler := rh.New(reviewService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "- method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	//member
	e.POST("/members", memberHandler.Insert())
	e.PUT("/members/:id", memberHandler.Update())
	e.DELETE("/members/:id", memberHandler.Delete())
	e.GET("/members", memberHandler.GetMembers())

	//product
	e.POST("/products", productHandler.Insert())
	e.PUT("/products/:id", productHandler.Update())
	e.DELETE("/products/:id", productHandler.Delete())
	e.GET("/products", productHandler.GetProducts())
	e.GET("/products/:id", productHandler.GetProductById())

	//review
	e.POST("/reviews", reviewHandler.Insert())
	e.POST("/likes", reviewHandler.Like())
	e.DELETE("/likes", reviewHandler.Unlike())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
