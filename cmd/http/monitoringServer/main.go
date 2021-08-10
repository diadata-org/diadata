package main

import (
	"github.com/diadata-org/diadata/http/monitoringServer/nodes"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.GET("/monitoring/nodes", GetNodeInformations)

	// This environment variable is either set in docker-compose or empty
	err := engine.Run(utils.Getenv("LISTEN_PORT", ":8080"))
	if err != nil {
		log.Error(err)
	}
}

func GetNodeInformations(context *gin.Context) {
	nodeInfos := nodes.GetNodeInformation()
	if len(nodeInfos) == 0 {
		restApi.SendError(context, http.StatusInternalServerError, nil)
	}
	context.JSON(http.StatusOK, nodeInfos)
}
