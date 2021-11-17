package main

import (
	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/databases"
	"github.com/diadata-org/diadata/http/monitoringServer/helpers"
	"github.com/diadata-org/diadata/http/monitoringServer/nodes"
	"github.com/diadata-org/diadata/http/monitoringServer/platform"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	routerGroup := engine.Group("/monitoring")
	routerGroup.GET("/", GetAllStates)

	databases.AddRoutes(routerGroup)
	nodes.AddRoutes(routerGroup)
	platform.AddRoutes(routerGroup)

	// This environment variable is either set in docker-compose or empty
	err := engine.Run(utils.Getenv("LISTEN_PORT", ":8080"))
	if err != nil {
		log.Error(err)
	}
}

func GetAllStates(context *gin.Context) {
	var states []config.State
	dbStates := databases.GetAllStates()
	states = append(states, helpers.MergeStateSlicesAsSubsection("databases", dbStates))
	nodeStates := nodes.GetAllStates()
	states = append(states, helpers.MergeStateSlicesAsSubsection("nodes", nodeStates))
	platformStates := platform.GetAllStates()
	states = append(states, helpers.MergeStateSlicesAsSubsection("platform", platformStates))

	if len(states) == 0 {
		restApi.SendError(context, http.StatusNotFound, nil)
	}
	context.JSON(http.StatusOK, states)
}
