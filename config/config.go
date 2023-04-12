package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
	jwtKey string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}

	err := godotenv.Load("local.env")
	if err != nil {
		log.Println("error reading environment", err.Error())
		return nil
	}

	app.DBUser = os.Getenv("DBUSER")
	app.DBPass = os.Getenv("DBPASS")
	app.DBHost = os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	app.DBPort, err = strconv.Atoi(port)
	if err != nil {
		fmt.Println("Error saat convert", err.Error())
		return nil
	}
	app.DBName = os.Getenv("DBNAME")
	app.jwtKey = os.Getenv("JWTKEY")

	return &app
}
