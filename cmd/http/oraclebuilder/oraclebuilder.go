package main

import (
	"net/http"
	"strconv"

	"github.com/99designs/keyring"
	builderUtils "github.com/diadata-org/diadata/http/oraclebuilder/utils"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
)

var log = logrus.New()

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	relStore, err := models.NewPostgresDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore", err)
	}
	k8bridgeurl := utils.Getenv("K8SBRIDGE_URL", "127.0.0.1:50051")
	signerurl := utils.Getenv("SIGNER_URL", "signer.dia-oracle-feeder:50052")
	diaRestURL := utils.Getenv("DIA_REST_URL", "https://api.diadata.org")
	diaGraphqlURL := utils.Getenv("DIA_GRAPHQL_URL", "https://api.diadata.org/graphql/query")
	rateLimitOracleCreationString := utils.Getenv("RATE_LIMIT_ORACLE_CREATION", "4")
	rateLimitOracleCreation, _ := strconv.ParseInt(rateLimitOracleCreationString, 10, 64)

	oraclebaseimage := utils.Getenv("ORACLE_BASE_IMAGE", "us.icr.io/dia-registry/oracles/oracle-baseimage:latest")
	oraclenamespace := utils.Getenv("ORACLE_NAMESPACE", "dia-oracle-feeder")
	oracleMonitoringUser := utils.Getenv("ORACLE_MONITORING_USER", "user")
	oracleMonitoringPassword := utils.Getenv("ORACLE_MONITORING_PASSWORD", "password")
	affinity := utils.Getenv("ORACLE_FEEDER_AFFINITY", "default")

	ph := builderUtils.NewPodHelper(oraclebaseimage, oraclenamespace, affinity, signerurl, diaRestURL, diaGraphqlURL)

	ring, _ := keyring.Open(keyring.Config{
		ServiceName:     "oraclebuilder",
		Server:          k8bridgeurl,
		AllowedBackends: []keyring.BackendType{keyring.K8Secret},
	})

	oracle := NewEnv(relStore, ph, ring, int(rateLimitOracleCreation))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	routerGroup := r.Group("/oraclebuilder")

	routerGroup.POST("create", oracle.Create)
	routerGroup.GET("/list", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to List your oracles") }, oracle.Auth, oracle.List)
	routerGroup.GET("/view", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to List your oracles") }, oracle.Auth, oracle.View)
	routerGroup.DELETE("/delete", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to delete oracle") }, oracle.Auth, oracle.Delete)
	routerGroup.PATCH("/restart", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to restart oracle feeder") }, oracle.Auth, oracle.Restart)
	routerGroup.PATCH("/pause", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to pause oracle feeder") }, oracle.Auth, oracle.Pause)
	routerGroup.GET("/whitelist", oracle.Whitelist)
	routerGroup.GET("/stats", oracle.Stats)
	routerGroup.GET("/create", oracle.ViewLimit)

	authMiddleware := basicAuth(oracleMonitoringUser, oracleMonitoringPassword)

	routerGroup.GET("/listAll", authMiddleware, oracle.ListAll)

	port := utils.Getenv("LISTEN_PORT", ":8080")

	executionMode := utils.Getenv("EXEC_MODE", "")
	if executionMode == "production" {
		err = r.Run(port)
		if err != nil {
			log.Error(err)
		}
	} else {
		err = r.Run(":8081")
		if err != nil {
			log.Error(err)
		}
	}

}
func basicAuth(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()

		if !hasAuth || user != username || pass != password {
			c.Header("WWW-Authenticate", "Basic realm=Restricted")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
