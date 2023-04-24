package validator

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func ValidatePostgres() bool {
	client, err := pgx.Connect(context.Background(), os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	query := "SELECT NOW()"
	var example time.Time
	
	err = client.QueryRow(ctx, query).Scan(&example)
	if err != nil {
		log.Fatal("Unable to connect to Postgres.", err)
		return false
	}
	return true
}