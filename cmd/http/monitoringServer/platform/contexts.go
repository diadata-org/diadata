package platform

import (
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPlatformStatus(context *gin.Context) {
	states := GetAllStates()
	if len(states) == 0 {
		restApi.SendError(context, http.StatusNotFound, nil)
	}
	context.JSON(http.StatusOK, states)
}
