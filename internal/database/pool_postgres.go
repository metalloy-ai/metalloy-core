package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"metalloyCore/internal/config"
)

func GetPool(cfg config.Setting) *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), cfg.PG_URL) 
	if err != nil {
		log.Fatal(err)
	}
	return pool
}
