package main

import (
	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/databases"
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
	"github.com/diadata-org/diadata/http/monitoringServer/nodes"
	"github.com/diadata-org/diadata/http/monitoringServer/platform"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/diadata-org/diadata/pkg/utils/probes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var startupDone = false

func main() {

	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	routerGroup := engine.Group("/monitoring")
	routerGroup.GET("/", GetAllStates)
	databases.AddRoutes(routerGroup)
	nodes.AddRoutes(routerGroup)
	platform.AddRoutes(routerGroup)
	startupDone = true
	// This environment variable is either set in docker-compose or empty

	go func() {
		err := engine.Run(utils.Getenv("LISTEN_PORT", ":8080"))
		if err != nil {
			log.Error(err)
		}
	}()

	log.Infoln("starting probes")
	probes.Start(live, ready)
}

func ready() bool {
	return startupDone
}

func live() bool {
	if !startupDone {
		return false
	}
	return config.GetKubernetesConnection() != nil
}

func mergeStateSlicesAsSubsection(name string, states []config.State) config.State {
	state := config.GetOperationalHealthState(name)

	for _, oneSlice := range states {
		state.Subsection = append(state.Subsection, oneSlice)
		switch oneSlice.State {
		case enums.HealthStateMajor:
			state.State = enums.HealthStateMajor
		case enums.HealthStateMinor:
			if state.State == enums.HealthStateOperational || state.State == enums.HealthStateMaintenance {
				state.State = enums.HealthStateMinor
			}
		case enums.HealthStateMaintenance:
			if state.State == enums.HealthStateOperational {
				state.State = enums.HealthStateOperational
			}
		}
	}

	return state
}

func GetAllStates(context *gin.Context) {
	var states []config.State
	dbStates := databases.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("databases", dbStates))
	nodeStates := nodes.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("nodes", nodeStates))
	platformStates := platform.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("platform", platformStates))

	if len(states) == 0 {
		restApi.SendError(context, http.StatusNotFound, nil)
	}
	context.JSON(http.StatusOK, states)
}
