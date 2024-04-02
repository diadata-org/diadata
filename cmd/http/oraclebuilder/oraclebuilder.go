package main

import (
	"net/http"
	"strconv"

	"github.com/99designs/keyring"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-contrib/cors"

	k8sbridge "github.com/diadata-org/diadata/pkg/dia/helpers/k8sbridge/protoc"
)

var log = logrus.New()

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	relStore, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore", err)
	}

	ds, err := models.NewDataStoreWithoutInflux()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}
	k8bridgeurl := utils.Getenv("K8SBRIDGE_URL", "127.0.0.1:50051")

	rateLimitOracleCreationString := utils.Getenv("RATE_LIMIT_ORACLE_CREATION", "4")
	rateLimitOracleCreation, _ := strconv.ParseInt(rateLimitOracleCreationString, 10, 64)

	oracleMonitoringUser := utils.Getenv("ORACLE_MONITORING_USER", "user")
	oracleMonitoringPassword := utils.Getenv("ORACLE_MONITORING_PASSWORD", "password")

	routerPath := utils.Getenv("ROUTER_PATH", "/oraclebuilder")

	conn, err := grpc.Dial(k8bridgeurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Error while connecting to the signer service: %v", err)
		panic("signer connection error")

	}
	k8sbridgeClient := k8sbridge.NewK8SHelperClient(conn)

	ring, _ := keyring.Open(keyring.Config{
		ServiceName:     "oraclebuilder",
		Server:          k8bridgeurl,
		AllowedBackends: []keyring.BackendType{keyring.K8Secret},
	})

	oracle := NewEnv(relStore, ds, k8sbridgeClient, ring, int(rateLimitOracleCreation))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	routerGroup := r.Group(routerPath)

	routerGroup.POST("create", oracle.Create)
	routerGroup.GET("/create", oracle.ViewLimit)
	routerGroup.GET("/list", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to List your oracles") }, oracle.Auth, oracle.List)
	routerGroup.GET("/view", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to List your oracles") }, oracle.Auth, oracle.View)
	routerGroup.DELETE("/delete", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to delete oracle") }, oracle.Auth, oracle.Delete)
	routerGroup.PATCH("/restart", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to restart oracle feeder") }, oracle.Auth, oracle.Restart)
	routerGroup.PATCH("/pause", func(ctx *gin.Context) { ctx.Set("message", "Verify its your address to pause oracle feeder") }, oracle.Auth, oracle.Pause)
	routerGroup.GET("/whitelist", oracle.Whitelist)
	routerGroup.GET("/stats", oracle.Stats)
	routerGroup.GET("/dashboard", oracle.Dashboard)

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
