package database

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	"metalloyCore/internal/config"
)

var (
	poolInstance *pgxpool.Pool
	poolOnce     sync.Once
)

func GetPool(cfg config.Setting) *pgxpool.Pool {
	poolOnce.Do(func() {
		config, err := pgxpool.ParseConfig(cfg.PG_URL)
		if err != nil {
			log.Fatal(err)
		}

		config.MinConns = int32(cfg.NumCPU) / 2
		config.MaxConns = int32(cfg.NumCPU) * 4

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatal(err)
		}
		poolInstance = pool
	})

	return poolInstance
}
