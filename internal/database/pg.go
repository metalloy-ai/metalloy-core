package database

import (
	"context"
	"log"
	"logiflowCore/internal/config"

	"github.com/jackc/pgx/v5"
)

func GetClient(cfg config.Setting) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), cfg.PG_URL)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}