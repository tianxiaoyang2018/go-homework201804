package cache

import (
	"context"
	"time"
)

type ICacheClient interface {
	Get(ctx context.Context, key string) (string, error)
	Exist(ctx context.Context, key string) (bool, error)
	Gets(ctx context.Context, keys []string) ([]interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration ...time.Duration) error
	Sets(ctx context.Context, kvs map[string]interface{}, expiration ...time.Duration) error
	Del(ctx context.Context, key []string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
}
