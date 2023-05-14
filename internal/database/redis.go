package database

import (
	"sync"

	"github.com/go-redis/redis"

	"metalloyCore/internal/config"
)

var conn *redis.Client
var once sync.Once

func GetRedisClient(cfg config.Setting) *redis.Client {
	once.Do(func() {
		conn = redis.NewClient(&redis.Options{
			Addr:     cfg.REDIS_URL,
			Password: cfg.REDIS_PWS,
			DB:       0,
		})
	})
		
	return conn
}
