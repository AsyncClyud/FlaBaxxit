package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

/*
Create connection to database from argument connStr.

If connStr is invalid - returns error and do os.Exit(1)

If cannot create pgxpool - returns error and do os.Exit(1)

If have ping timeout - returns error and do os.Exit(1)
*/
func ConnectDataBase(ctx context.Context, connStr string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Parse config error: %v", err)
	}

	config.MinConns = 10
	config.MaxConns = 45

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Pgxpool creation error: %v", err)
	}

	if ok := db.Ping(ctx); ok != nil {
		log.Fatalf("Pgxpool connection error: %v", err)
	}

	return db
}
