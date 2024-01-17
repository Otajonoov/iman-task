package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnDB() (*pgxpool.Pool, error) {
	cfg := Load()

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	pool, err := pgxpool.New(context.Background(), psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return pool, nil
}
