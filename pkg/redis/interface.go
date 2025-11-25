package pkgredis

import (
	"context"
	"time"
)

type (
	IRedis interface {
		Set(ctx context.Context, request SetRequest) (*SetResponse, error)
		Get(ctx context.Context, request GetRequest) (*GetResponse, error)
		Del(ctx context.Context, request DelRequest) (*DelResponse, error)
		Incr(ctx context.Context, request IncrRequest) (*IncrResponse, error)
		Expire(ctx context.Context, request ExpireRequest) (*ExpireResponse, error)
		SetJSON(ctx context.Context, key string, value any, ttl time.Duration) error
		GetJSON(ctx context.Context, key string, dest any) (bool, error)
		GetUsageCacheVersion(ctx context.Context) (int64, error)
		BumpUsageCacheVersion(ctx context.Context) error
		Publish(ctx context.Context, channel string, message string) error
		Subscribe(ctx context.Context, channel string) (<-chan string, error)
	}
)
