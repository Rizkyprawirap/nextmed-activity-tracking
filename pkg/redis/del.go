package pkgredis

import (
	"context"
)

func (r *pkgRedis) Del(ctx context.Context, request DelRequest) (*DelResponse, error) {
	err := r.Client.Del(ctx, request.Key).Err()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
