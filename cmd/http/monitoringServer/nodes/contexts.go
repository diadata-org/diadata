package nodes

import (
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNodeInformation(context *gin.Context) {
	nodeInfos := GetAllStates()
	if len(nodeInfos) == 0 {
		restApi.SendError(context, http.StatusNotFound, nil)
	}
	context.JSON(http.StatusOK, nodeInfos)
}
