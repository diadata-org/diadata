package platform

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/platform")

	router.GET("/status", GetPlatformStatus)

}
