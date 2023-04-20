package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}
}