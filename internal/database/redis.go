package database

import (
	"metalloyCore/internal/config"

	"github.com/go-redis/redis"
)

func GetRedisClient(cfg config.Setting) *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_URL,
		Password: cfg.REDIS_PWS,
		DB:       0,
	})
	return conn
}
