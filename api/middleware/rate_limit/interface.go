package apimiddlewareratelimit

import (
	"github.com/gin-gonic/gin"
)

type IMiddlewareRateLimiter interface {
	Use(config RateLimitConfig) gin.HandlerFunc
}
