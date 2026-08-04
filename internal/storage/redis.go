package storage

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

/*
Create connection to redis from argument connStr.

If connStr is invalid - returns error and do os.Exit(1)

If have ping timeout - returns error and do os.Exit(1)
*/
func ConnectRedis(ctx context.Context, connStr string) *redis.Client {
	redis_url, err := redis.ParseURL(connStr)
	if err != nil {
		log.Fatalf("Redis ParseURL error: %v", err)
	}
	rdb := redis.NewClient(redis_url)

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis fatal connection error: %v", err)
	}

	return rdb
}
