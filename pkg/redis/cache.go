package pkgredis

import (
	"context"
	"encoding/json"
	"time"
)

func (r *pkgRedis) SetJSON(ctx context.Context, key string, value any, ttl time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.Set(ctx, SetRequest{
		Key:   key,
		Value: string(jsonData),
	})
	if err != nil {
		return err
	}

	_, err = r.Expire(ctx, ExpireRequest{
		Key:     key,
		Seconds: int64(ttl.Seconds()),
	})
	return err
}

func (r *pkgRedis) GetJSON(ctx context.Context, key string, dest any) (bool, error) {
	resp, err := r.Get(ctx, GetRequest{Key: key})
	if err != nil {
		return false, err
	}

	if resp.Result == "" {
		return false, nil
	}

	if err := json.Unmarshal([]byte(resp.Result), dest); err != nil {
		return false, err
	}

	return true, nil
}

func (r *pkgRedis) GetUsageCacheVersion(ctx context.Context) (int64, error) {
	resp, err := r.Incr(ctx, IncrRequest{
		Key: "cache:usage:version",
	})
	if err != nil {
		return 1, err
	}
	return resp.Result, nil
}

func (r *pkgRedis) BumpUsageCacheVersion(ctx context.Context) error {
	_, err := r.Incr(ctx, IncrRequest{
		Key: "cache:usage:version",
	})
	return err
}
