package storage

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(ctx context.Context, connStr *redis.Options) *redis.Client{
	rdb := redis.NewClient(connStr)
 	if err := rdb.Ping(ctx).Err(); err != nil {
        log.Fatal("Redis fatal connection error:", err)
    }
	return rdb
}
