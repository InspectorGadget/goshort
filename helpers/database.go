package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDatabaseDsn() string {
	godotenv.Load()

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	if DB_USERNAME == "" || DB_PASSWORD == "" || DB_HOST == "" || DB_PORT == "" || DB_NAME == "" {
		return ""
	}

	return DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
}
