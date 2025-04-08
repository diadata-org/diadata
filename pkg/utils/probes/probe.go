package probes

import (
	"errors"
	"net/http"

	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type probe func() bool

var livenessProbe probe
var readinessProbe probe

func Start(liveness probe, readiness probe) {

	log.Info("Ready and Live probes loading")
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
	log.Info("Ready and Live probes starting")
}

func execReadiness(context *gin.Context) {
	executeProbe(context, readinessProbe, http.StatusTooEarly)
}

func execLiveness(context *gin.Context) {
	executeProbe(context, livenessProbe, http.StatusInternalServerError)
}

func executeProbe(context *gin.Context, fn probe, errorCode int) bool {
	if fn() {
		context.JSON(http.StatusOK, gin.H{"message": "success"})
		return true
	}
	restApi.SendError(context, errorCode, errors.New("not live"))
	return false
}
