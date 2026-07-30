package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDataBase(ctx context.Context, connStr string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Parse config error: %v", err)
	}

	config.MinConns = 10
	config.MaxConns = 45

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Pgxpool connection error: %v", err)
	}

	return db
}
