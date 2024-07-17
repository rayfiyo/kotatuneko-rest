package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Mode        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	DB_DSN string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load the .env: %v", err)
	}

	Mode = os.Getenv("MODE")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")

	DB_DSN = "host=" + DB_HOST +
		" user=" + DB_USER +
		" password=" + DB_PASSWORD +
		" dbname=" + DB_NAME
}
