package cache

import (
	"context"
	"time"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	DeleteByPattern(ctx context.Context, pattern string) error
	Ping(ctx context.Context) error
}
