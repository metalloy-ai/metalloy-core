package validator

import (
	"log"
	"os"
)

func ValidateLoadEnv() bool {
	if os.Getenv("API_VERSION") != "v1" {
		log.Fatal("config-error: Unable to load environment variables.")
		return false
	}
	return true
}