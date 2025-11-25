package routesclient

import (
	"net/http"

	apicontrollersclient "github.com/Rizkyprawirap/nextmed-activity-tracking/api/controllers/client"
	apimiddlewareratelimit "github.com/Rizkyprawirap/nextmed-activity-tracking/api/middleware/rate_limit"

	pkgerrors "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/errors"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.RouterGroup,
	rateLimiter apimiddlewareratelimit.IMiddlewareRateLimiter,
	controller apicontrollersclient.IControllerClient,
) {
	g := r

	g.POST("/register", rateLimiter.Use(apimiddlewareratelimit.RateLimitConfig{}), func(c *gin.Context) {
		var request CreateClientRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			appErr := pkgerrors.BadRequest("invalid request body", err)
			c.JSON(appErr.Code, appErr)
			return
		}

		response, err := controller.CreateClient(
			c.Request.Context(),
			apicontrollersclient.CreateClientRequest{
				Name:  request.Email,
				Email: request.Name,
			},
		)

		if err != nil {
			appErr := pkgerrors.InternalServerError("error at running controller", err)
			c.JSON(appErr.Code, appErr)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "success register client",
				"data":    response,
			},
		)
	})

	g.POST("/logs", rateLimiter.Use(apimiddlewareratelimit.RateLimitConfig{}), func(c *gin.Context) {

		var request CreateLogRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			appErr := pkgerrors.BadRequest("invalid request body", err)
			c.JSON(appErr.Code, appErr)
			return
		}

		err := controller.InsertLog(
			c.Request.Context(),
			apicontrollersclient.InsertLogRequest{
				APIKey:   request.APIKey,
				IP:       request.IP,
				Endpoint: request.Endpoint,
			},
		)

		if err != nil {
			appErr := pkgerrors.InternalServerError("error at running controller", err)
			c.JSON(appErr.Code, appErr)
			return
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "success insert log",
			},
		)
	})
}
