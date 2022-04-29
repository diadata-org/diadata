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

func main(liveness probe, readiness probe) {

	livenessProbe = liveness
	readinessProbe = readiness

	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.GET("/ready", execReadiness)
	engine.GET("/live", execLiveness)
	// This environment variable is either set in docker-compose or empty
	err := engine.Run(utils.Getenv("LISTEN_PORT_PROBES", ":2345"))
	if err != nil {
		log.Error(err)
	}
}

func execReadiness(context *gin.Context) {
	executeProbe(context, readinessProbe)
}

func execLiveness(context *gin.Context) {
	executeProbe(context, livenessProbe)
}

func executeProbe(context *gin.Context, fn probe) bool {
	if fn() {
		context.JSON(200, gin.H{"message": "success"})
		return true
	}
	restApi.SendError(context, http.StatusInternalServerError, nil)
	return false
}
