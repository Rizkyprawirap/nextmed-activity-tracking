package apimiddlewareratelimit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	pkgjwt "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/jwt"
	pkgRateLimit "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/rate_limit"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	jwt       pkgjwt.IPkgJWT
	rateLimit pkgRateLimit.IPkgRateLimit
}

func New(jwt pkgjwt.IPkgJWT, rl pkgRateLimit.IPkgRateLimit) IMiddlewareRateLimiter {
	return &rateLimiter{
		jwt:       jwt,
		rateLimit: rl,
	}
}

func (r *rateLimiter) Use(config RateLimitConfig) gin.HandlerFunc {
	return func(c *gin.Context) {

		defaultRateLimit, err := strconv.ParseInt(os.Getenv("RATE_LIMIT"), 10, 64)
		if err != nil {
			panic(fmt.Errorf("invalid RATE_LIMIT: %w", err))
		}
		key := c.ClientIP()

		if u, ok := c.Get("user"); ok && u != nil {
			if id := extractIDFromClaims(u); id != "" {
				key = id
			}
		}

		rateLimit := config.Limit
		if rateLimit <= 0 {
			rateLimit = defaultRateLimit
		}

		err = r.rateLimit.RateLimit(c.Request.Context(), pkgRateLimit.RateLimitRequest{
			ID:            key,
			Limit:         rateLimit,
			WindowSeconds: 60,
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": "too many requests",
			})
			return
		}

		c.Next()
	}
}

func extractIDFromClaims(u any) string {
	if m, ok := u.(map[string]any); ok {
		if v, ok := m["id"].(string); ok && v != "" {
			return v
		}
		if v, ok := m["ID"].(string); ok && v != "" {
			return v
		}
		if v, ok := m["user_id"].(string); ok && v != "" {
			return v
		}
		if v, ok := m["sub"].(string); ok && v != "" {
			return v
		}
	}

	v := reflect.ValueOf(u)
	if v.IsValid() && v.Kind() == reflect.Struct {
		rc := v.FieldByName("RegisteredClaims")
		if rc.IsValid() && rc.Kind() == reflect.Struct {
			if f := rc.FieldByName("ID"); f.IsValid() && f.Kind() == reflect.String {
				if s := f.String(); s != "" {
					return s
				}
			}
			if f := rc.FieldByName("Subject"); f.IsValid() && f.Kind() == reflect.String {
				if s := f.String(); s != "" {
					return s
				}
			}
		}
		data := v.FieldByName("Data")
		if data.IsValid() && data.Kind() == reflect.Slice && data.Type().Elem().Kind() == reflect.Uint8 {
			b := make([]byte, data.Len())
			reflect.Copy(reflect.ValueOf(b), data)
			var mm map[string]any
			if err := json.Unmarshal(b, &mm); err == nil {
				if v, ok := mm["id"].(string); ok && v != "" {
					return v
				}
				if v, ok := mm["ID"].(string); ok && v != "" {
					return v
				}
				if v, ok := mm["user_id"].(string); ok && v != "" {
					return v
				}
				if v, ok := mm["sub"].(string); ok && v != "" {
					return v
				}
			}
		}
	}
	return ""
}
