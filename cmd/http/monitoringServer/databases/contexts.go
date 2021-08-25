package databases

import (
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllDatabaseStates(context *gin.Context) {
	states := GetAllStates()
	if len(states) == 0 {
		restApi.SendError(context, http.StatusNotFound, nil)
	}
	context.JSON(http.StatusOK, states)
}

func GetInfluxState(context *gin.Context) {
	context.JSON(http.StatusOK, InfluxState())
}

func GetPostgresState(context *gin.Context) {
	context.JSON(http.StatusOK, PostgresState())
}

func GetRedisState(context *gin.Context) {
	context.JSON(http.StatusOK, RedisState())
}

func GetKafkaState(context *gin.Context) {
	context.JSON(http.StatusOK, KafkaState())
}
