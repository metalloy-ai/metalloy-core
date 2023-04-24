package validator

import (
	"log"
	"logiflowCore/internal/config"
)

func ValidateConfig(cfg config.Setting) {
	cond := ValidateLoadEnv(cfg) && ValidatePostgres(cfg)
	if !cond {
		log.Fatal("validator-error: Unable to validate config.")
	}
}