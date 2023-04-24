package validator

import "log"

func ValidateConfig() {
	cond := ValidateLoadEnv() && ValidatePostgres() && ValidateRedis()
	if !cond {
		log.Fatal("validator-error: Unable to validate config.")
	}
}