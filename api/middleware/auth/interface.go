package apimiddlewareauth

import (
	"context"

	"github.com/gin-gonic/gin"
)

type IMiddlewareAuth interface {
	Validate(ctx context.Context) gin.HandlerFunc
}
