package pgkratelimit

type (
	RateLimitRequest struct {
		ID            string
		Limit         int64
		WindowSeconds int64
	}
)
