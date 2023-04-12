package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	member "kstyleAPI/features/member/data"
	product "kstyleAPI/features/product/data"
	rp "kstyleAPI/features/reviewProduct/data"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.DBUser, ac.DBPass, ac.DBHost, ac.DBPort, ac.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(member.Member{})
	db.AutoMigrate(product.Product{})
	db.AutoMigrate(rp.ReviewProduct{})
	db.AutoMigrate(rp.LikeReview{})
}
