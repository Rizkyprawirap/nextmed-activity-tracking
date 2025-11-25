package pkgredis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

func (r *pkgRedis) Get(ctx context.Context, request GetRequest) (*GetResponse, error) {
	data, err := r.Client.Get(ctx, request.Key).Result()
	if err == redis.Nil {
		return nil, errors.New("key does not exist")
	} else if err != nil {
		return nil, err
	}

	return &GetResponse{
		Result: data,
	}, nil
}
