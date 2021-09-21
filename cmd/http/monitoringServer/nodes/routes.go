package nodes

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/nodes")

	router.GET("/", GetNodeInformation)

}
