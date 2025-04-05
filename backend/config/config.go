package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}
