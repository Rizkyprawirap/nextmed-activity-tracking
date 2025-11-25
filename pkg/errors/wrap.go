package pkgerrors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context, err error) {
	if err == nil {
		return
	}

	appErr, ok := err.(*AppError)
	if !ok {
		c.JSON(http.StatusInternalServerError, AppError{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Detail:  err.Error(),
		})
		return
	}

	c.JSON(appErr.Code, appErr)
}
