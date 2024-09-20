package main

import (
	"net/http"
	"strconv"

	"github.com/99designs/keyring"
	k8sbridge "github.com/diadata-org/diadata/pkg/dia/helpers/k8sbridge/protoc"
	model "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logrus.New()

func main() {
	r := setupRouter()
	port := utils.Getenv("LISTEN_PORT", ":8080")
	executionMode := utils.Getenv("EXEC_MODE", "")

	if executionMode == "production" {
		log.Println("Running server on port", port)

		err := r.Run(port)
		if err != nil {
			log.Error(err)
		}
	} else {
		err := r.Run(":8081")
		if err != nil {
			log.Error(err)
		}
	}
}

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	loopPaymentSecret := utils.Getenv("LOOP_PAYMENT_SECRET", "")

	relStore, err := model.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore", err)
	}

	ds, err := model.NewDataStoreWithoutInflux()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	k8bridgeClient := initializeKubernetesBridgeClient()

	ring, _ := initializeKeyring()

	rateLimitOracleCreation := initializeRateLimitOracleCreation()
	oracle := NewEnv(relStore, ds, k8bridgeClient, ring, rateLimitOracleCreation, loopPaymentSecret)

	// Setup CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))

	routerPath := utils.Getenv("ROUTER_PATH", "/oraclebuilder")
	routerGroup := r.Group(routerPath)

	// endpoints without account
	routerGroup.POST("createAccount", authenticate("Verify its your address to create Account"), oracle.Auth(false), oracle.CreateAccount)
	routerGroup.POST("/account/approveWallet", authenticate("Verify its your address to Approve Adding this Wallet to Team"), oracle.Auth(false), oracle.ApproveWallet)

	// endpoints with and without account
	routerGroup.GET("/account/dashboard", authenticate("Verify its your address to List your oracles and View Account"), oracle.Auth(false), oracle.ListAndViewAccount)

	// Oracle Create By steps
	routerGroup.POST("/oracle/feeder", authenticate("Verify its your address to create and update oracle"), oracle.Auth(true), oracle.CanWrite, oracle.CreateUpdateOracle)
	routerGroup.POST("/oracle/deploy", authenticate("Verify its your address to start feeder"), oracle.Auth(true), oracle.CanWrite, oracle.InitFeeder)

	// Define routes
	routerGroup.POST("create", oracle.CreateOneStep)
	routerGroup.GET("/create", oracle.ViewLimit)

	routerGroup.POST("/account/removeWallet", authenticate("Verify its your address to Remove Wallet"), oracle.Auth(true), oracle.CanWrite, oracle.RemoveWallet)
	routerGroup.POST("/account/addWallet", authenticate("Verify its your address to Add Wallet"), oracle.Auth(true), oracle.CanWrite, oracle.AddTempWallet)

	routerGroup.POST("/account/updateAccess", authenticate("Verify its your address to Update Access"), oracle.Auth(true), oracle.CanWrite, oracle.UpdateAccess)

	routerGroup.POST("/account/view", authenticate("Verify its your address to View Account"), oracle.Auth(true), oracle.CanRead, oracle.ViewAccount)

	// oracle actions
	routerGroup.GET("/list", authenticate("Verify its your address to List your oracles"), oracle.Auth(true), oracle.CanRead, oracle.List)

	routerGroup.GET("/view", authenticate("Verify its your address to List your oracles"), oracle.Auth(true), oracle.CanRead, oracle.View)
	routerGroup.DELETE("/delete", authenticate("Verify its your address to delete oracle"), oracle.Auth(true), oracle.CanWrite, oracle.Delete)
	routerGroup.PATCH("/restart", authenticate("Verify its your address to restart oracle feeder"), oracle.Auth(true), oracle.CanWrite, oracle.Restart)
	routerGroup.PATCH("/pause", authenticate("Verify its your address to pause oracle feeder"), oracle.Auth(true), oracle.CanWrite, oracle.Pause)

	//payment related endpoints
	routerGroup.POST("/paymenthook", oracle.LoopWebHook)
	routerGroup.GET("/paymentStatus", oracle.LoopPaymentStatus)

	routerGroup.GET("/whitelist", oracle.Whitelist)
	routerGroup.GET("/stats", oracle.Stats)
	routerGroup.GET("/chains", oracle.SupportedChains)
	routerGroup.GET("/dashboard", oracle.Dashboard)

	// Middleware for basic authentication
	oracleMonitoringUser := utils.Getenv("ORACLE_MONITORING_USER", "user")
	oracleMonitoringPassword := utils.Getenv("ORACLE_MONITORING_PASSWORD", "password")
	authMiddleware := basicAuth(oracleMonitoringUser, oracleMonitoringPassword)
	routerGroup.GET("/listAll", authMiddleware, oracle.ListAll)

	return r
}

func initializeKubernetesBridgeClient() k8sbridge.K8SHelperClient {
	k8bridgeurl := utils.Getenv("K8SBRIDGE_URL", "127.0.0.1:50051")
	conn, err := grpc.Dial(k8bridgeurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Error while connecting to the signer service: %v", err)
		panic("signer connection error")
	}
	return k8sbridge.NewK8SHelperClient(conn)
}

func initializeKeyring() (keyring.Keyring, error) {
	k8bridgeurl := utils.Getenv("K8SBRIDGE_URL", "127.0.0.1:50051")
	return keyring.Open(keyring.Config{
		ServiceName:     "oraclebuilder",
		Server:          k8bridgeurl,
		AllowedBackends: []keyring.BackendType{keyring.K8Secret},
	})
}

func initializeRateLimitOracleCreation() int {
	rateLimitOracleCreationString := utils.Getenv("RATE_LIMIT_ORACLE_CREATION", "4")
	rateLimitOracleCreation, _ := strconv.ParseInt(rateLimitOracleCreationString, 10, 64)
	return int(rateLimitOracleCreation)
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

func authenticate(message string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("message", message)
	}
}
