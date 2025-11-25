package pgkratelimit

import "context"

type IPkgRateLimit interface {
	RateLimit(ctx context.Context, request RateLimitRequest) error
}
