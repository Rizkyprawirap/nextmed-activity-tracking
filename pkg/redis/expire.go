package pkgredis

import (
	"context"
	"time"
)

func (r *pkgRedis) Expire(ctx context.Context, request ExpireRequest) (*ExpireResponse, error) {
	res, err := r.Client.Expire(ctx, request.Key, time.Duration(request.Seconds)*time.Second).Result()
	if err != nil {
		return nil, err
	}

	return &ExpireResponse{
		Success: res,
	}, nil
}
