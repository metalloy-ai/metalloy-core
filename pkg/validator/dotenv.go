package validator

import (
	"log"
	"logiflowCore/internal/config"
)

func ValidateLoadEnv(cfg config.Setting) bool {
	if cfg.ApiVersion != "v1" {
		log.Fatal("config-error: Unable to load environment variables.")
		return false
	}
	return true
}