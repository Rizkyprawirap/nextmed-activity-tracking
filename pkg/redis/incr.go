package pkgredis

import (
	"context"
)

func (r *pkgRedis) Incr(ctx context.Context, request IncrRequest) (*IncrResponse, error) {
	res, err := r.Client.Incr(ctx, request.Key).Result()
	if err != nil {
		return nil, err
	}

	return &IncrResponse{
		Result: res,
	}, nil
}
