package main

import (
	"github.com/gin-gonic/gin"
)

func handleError(context *gin.Context, status int, errorMsg, logMsg string, logArgs ...interface{}) {
	m := make(map[string]string)
	m["error"] = errorMsg
	context.JSON(status, m)
	log.Errorf(logMsg, logArgs...)
	context.Abort()

}
