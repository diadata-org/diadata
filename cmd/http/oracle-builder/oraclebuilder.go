package main

import (
	"time"

	//jwt "github.com/blockstatecom/gin-jwt"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/oraclebuilder"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

const (
	cachingTime1Sec   = 1 * time.Second
	cachingTime20Secs = 20 * time.Second
	cachingTimeShort  = time.Minute * 2
	// cachingTimeMedium = time.Minute * 10
	cachingTimeLong = time.Minute * 100
)

var identityKey = "id"

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	relStore, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore", err)
	}

	ob := &oraclebuilder.Env{RelDB: relStore}

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))
	routerGroup := r.Group("/oraclebuilder")

	routerGroup.POST("/", ob.InitiateOracle)

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
