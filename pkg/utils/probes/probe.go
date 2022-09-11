package probes

import (
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type probe func() bool

var livenessProbe probe
var readinessProbe probe

func Start(liveness probe, readiness probe) {

	log.Infoln("Ready and Live probes loading")
	livenessProbe = liveness
	readinessProbe = readiness

	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.GET("/ready", execReadiness)
	engine.GET("/live", execLiveness)
	// This environment variable is either set in docker-compose or empty
	go func() {
		err := engine.Run(utils.Getenv("LISTEN_PORT_PROBES", ":2345"))
		if err != nil {
			log.Error(err)
		}
	}()
	log.Infoln("Ready and Live probes starting")
}

func execReadiness(context *gin.Context) {
	executeProbe(context, readinessProbe)
}

func execLiveness(context *gin.Context) {
	executeProbe(context, livenessProbe)
}

func executeProbe(context *gin.Context, fn probe) bool {
	log.Infoln("probe has been started")
	if fn() {
		context.JSON(http.StatusOK, gin.H{"message": "success"})
		return true
	}
	restApi.SendError(context, http.StatusInternalServerError, nil)
	return false
}
