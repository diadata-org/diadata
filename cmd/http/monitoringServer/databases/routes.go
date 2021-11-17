package databases

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup) {
	databaseUri := rg.Group("/databases")

	databaseUri.GET("/", GetAllDatabaseStates)
	databaseUri.GET("/influx", GetInfluxState)
	databaseUri.GET("/postgres", GetPostgresState)
	databaseUri.GET("/kafka", GetKafkaState)
	databaseUri.GET("/redis", GetRedisState)

}
