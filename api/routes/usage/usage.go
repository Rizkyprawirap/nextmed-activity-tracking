package routesusage

import (
	"context"
	"net/http"

	apicontrollersusage "github.com/Rizkyprawirap/nextmed-activity-tracking/api/controllers/usage"
	apimiddlewareratelimit "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware/rate_limit"

	auth "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware/auth"
	pkgerrors "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/errors"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	authValidation auth.IMiddlewareAuth,
	rateLimiter apimiddlewareratelimit.IMiddlewareRateLimiter,
	controller apicontrollersusage.IControllerUsage,
) {
	ctx := context.Background()
	g := r.Group("/usage")
	g.Use(authValidation.Validate(ctx))

	g.GET("/daily", rateLimiter.Use(apimiddlewareratelimit.RateLimitConfig{}), func(c *gin.Context) {
		var request GetDailyUsageRequest

		if err := c.ShouldBindQuery(&request); err != nil {
			appErr := pkgerrors.BadRequest("invalid request", err)
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		response, err := controller.GetDailyUsage(
			c.Request.Context(),
			apicontrollersusage.GetDailyUsageRequest{
				APIKey: request.APIKey,
			},
		)

		if err != nil {
			appErr := pkgerrors.InternalServerError("error at running controller", err)
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "success get daily requests",
				"data":    response,
			},
		)
	})

	g.GET("/top", rateLimiter.Use(apimiddlewareratelimit.RateLimitConfig{}), func(c *gin.Context) {
		response, err := controller.GetTopClientUsage(
			c.Request.Context(),
			apicontrollersusage.GetTopClientUsageRequest{},
		)

		if err != nil {
			appErr := pkgerrors.InternalServerError("error at running controller", err)
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "success get top client",
				"data":    response,
			},
		)
	})
}
