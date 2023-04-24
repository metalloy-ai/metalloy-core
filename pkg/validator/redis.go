package validator

import (
	"os"

	"github.com/go-redis/redis"
)

func ValidateRedis() bool {
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PSW"),
		DB:       0,
	})
	_, err := conn.Ping().Result()
    return err == nil
}