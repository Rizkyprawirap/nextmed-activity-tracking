package apimiddlewareauth

import (
	"context"
	"fmt"
	"strings"

	pkgerrors "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/errors"
	pkgjwt "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func (m *middlewareAuth) Validate(ctx context.Context) gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		temp := strings.Split(authHeader, "Bearer ")
		if len(temp) < 2 {
			appErr := pkgerrors.Unauthorized("authentication failed", fmt.Errorf("bearer token in Authorization header required"))
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		responseValidate, err := m.PkgJWT.Validate(ctx, pkgjwt.ValidateRequest{
			Token: temp[1],
		})
		if err != nil {
			appErr := pkgerrors.Unauthorized("authentication failed", err)
			c.JSON(appErr.Code, appErr)
			c.Abort()
			return
		}

		c.Set("user", responseValidate.Claims)
		c.Next()

	}

}
