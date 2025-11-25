package pgkratelimit

import (
	"context"
	"fmt"

	pkgredis "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/redis"
)

type pkgRateLimit struct {
	Redis pkgredis.IRedis
}

func New(redis pkgredis.IRedis) *pkgRateLimit {
	return &pkgRateLimit{
		Redis: redis,
	}
}

func (p *pkgRateLimit) RateLimit(ctx context.Context, req RateLimitRequest) error {
	key := "rate_limit:" + req.ID

	incrResp, err := p.Redis.Incr(ctx, pkgredis.IncrRequest{Key: key})
	if err != nil {
		return err
	}

	hits := incrResp.Result

	if hits == 1 {
		_, err = p.Redis.Expire(ctx, pkgredis.ExpireRequest{
			Key:     key,
			Seconds: req.WindowSeconds,
		})
		if err != nil {
			return err
		}
	}

	if hits > int64(req.Limit) {

		p.Redis.Del(ctx, pkgredis.DelRequest{
			Key: req.ID,
		})
		return fmt.Errorf("rate limit exceeded")
	}

	return nil
}
