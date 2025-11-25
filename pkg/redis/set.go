package pkgredis

import (
	"context"
	"time"
)

func (r *pkgRedis) Set(ctx context.Context, request SetRequest) (*SetResponse, error) {
	var exp time.Duration = 24 * time.Hour
	if request.Exp != 0 {
		exp = time.Duration(request.Exp) * time.Second
	}

	err := r.Client.Set(ctx, request.Key, request.Value, exp).Err()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
