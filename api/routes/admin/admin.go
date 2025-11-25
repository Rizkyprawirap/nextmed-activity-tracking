package routesadmin

import (
	apicontrollersadmin "github.com/Rizkyprawirap/nextmed-activity-tracking/api/controllers/admin"
	apimiddlewareratelimit "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware/rate_limit"

	auth "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware/auth"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	authValidation auth.IMiddlewareAuth,
	rateLimiter apimiddlewareratelimit.IMiddlewareRateLimiter,
	controller apicontrollersadmin.IControllerAdmin,
) {
	g := r

	g.POST("/login", rateLimiter.Use(apimiddlewareratelimit.RateLimitConfig{}), func(c *gin.Context) {
		var request AdminLoginRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		response, err := controller.Login(
			c.Request.Context(),
			apicontrollersadmin.GetAdminDetailRequest{
				Email:    request.Email,
				Password: request.Password,
			},
		)

		if err != nil {
			c.JSON(
				500,
				gin.H{
					"code":    500,
					"message": "error at running controller",
					"error":   err.Error(),
				},
			)
			return
		}

		c.JSON(
			200,
			gin.H{
				"code":    200,
				"message": "success login admin",
				"data":    response,
			},
		)
	})
}
