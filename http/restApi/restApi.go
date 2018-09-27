package restApi

import (
	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, errorCode int, err error) {
	c.JSON(errorCode,
		&APIError{
			ErrorCode:    errorCode,
			ErrorMessage: err.Error(),
		})
}
