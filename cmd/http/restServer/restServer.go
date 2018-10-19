package main

import (
	"github.com/appleboy/gin-jwt"
	_ "github.com/diadata-org/diadata/api/docs"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/diadata-org/diadata/pkg/http/restServer/diaApi"
	"github.com/diadata-org/diadata/pkg/http/restServer/kafkaApi"
	"github.com/diadata-org/diadata/pkg/model"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

// @title diadata.org API
// @version 1.0
// @description | 
// #@description  The world's crowd-driven financial data community has a professional API made for you.
// #@description # Decentral and transparent by design
// W@description ith our decentral approach to data verification, you can gain a deep insight into current and past pricing, volume and exchange info so you can make the right decisions to stay ahead of the game.
//
// ## Find the right data for your needs
// Show your users the most transparent data on the market with our API. Whether you're building a financial service, a portfolio management tool, a new media offering, or more, we have the most advanced and updated data on the market for your product.
// For Oracle usage see [github](https://github.com/diadata-org/diadata/blob/master/documentation/methodology/Oracles.md).
//
// ## Backtest your strategies
// Use the most efficient and transparent crypto data to run simulations and backtest your trading or investing strategies. With crowd-aggregated hundreds of exchanges you can be sure that you're getting the right picture every single time.
//
// ## Run Experiments
// Build your own models with our data, to further your interest or just for fun. With our flexible and powerful API, we provide you with a set of data that will help you draw insights and make conclusions.
//
// ## Request your data
// Set a bounty on gitcoin.io or drop us [line](mailto:API@diadata.org
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.diadata.org
// @BasePath /

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// GetTradesBlock godoc
// @Summary GetTradesBlock get messages in topic
// @Description GetTradesBlock
// @Tags kafka
// @Accept  json
// @Produce  json
// @Param offset query int false "offset, default is last"
// @Param elements query int false "number of elements"
// @Success 200 {object} dia.TradesBlock	"success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /kafka/tradesBlock [get]
func GetTradesBlock(c *gin.Context) {
	kafkaApi.Process(c, kafkaHelper.TopicTradesBlock)
}

// GetFiltersBlock godoc
// @Summary GetFiltersBlock get messages in topic
// @Description GetFiltersBlock
// @Tags kafka
// @Accept  json
// @Produce  json
// @Param offset query int false "offset, default is last"
// @Param elements query int false "number of elements"
// @Success 200 {object} dia.FiltersBlock	"success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /kafka/filtersBlock [get]
func GetFiltersBlock(c *gin.Context) {
	kafkaApi.Process(c, kafkaHelper.TopicFiltersBlock)
}

// GetTrades godoc
// @Summary GetTrades get messages in topic
// @Description GetTrades
// @Tags kafka
// @Accept  json
// @Produce  json
// @Param offset query int false "offset, default is last"
// @Param elements query int false "number of elements"
// @Success 200 {object} dia.TradesBlock	"success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /kafka/trades [get]
func GetTrades(c *gin.Context) {
	kafkaApi.Process(c, kafkaHelper.TopicTrades)
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config := dia.GetConfigApi()

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "party zone",
		Key:         []byte(config.SecretKey), // TOFIX: this secret key should be different from the accepted apikey and secret key downstairs
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if userID == config.ApiKey && password == config.SecretKey { // Temporary: only 1 valid key so far.
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			log.Warning("Authenticator ErrFailedAuthentication")
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == config.ApiKey {
				return true
			}
			log.Warning("Authorizator rejected")
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	kafka := r.Group("/kafka")
	{
		kafka.GET("/tradesBlock", GetTradesBlock)
		kafka.GET("/filtersBlock", GetFiltersBlock)
		kafka.GET("/trades", GetTrades)
	}

	store, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}
	diaApiEnv := &diaApi.Env{store}

	diaAuth := r.Group("/v1")
	diaAuth.Use(authMiddleware.MiddlewareFunc())
	{
		diaAuth.POST("/supply", diaApiEnv.PostSupply)
	}

	dia := r.Group("/v1")
	{
		dia.GET("/quotation/:symbol", diaApiEnv.GetQuotation)
		dia.GET("/supply/:symbol", diaApiEnv.GetSupply)
		dia.GET("/symbol/:symbol", diaApiEnv.GetSymbolDetails)
		dia.GET("/coins", diaApiEnv.GetCoins)
		dia.GET("/pairs", diaApiEnv.GetPairs)
		dia.GET("/chartPoints/:filter/:exchange/:symbol", diaApiEnv.GetChartPoints)
		dia.GET("/chartPointsAllExchanges/:filter/:symbol", diaApiEnv.GetChartPointsAllExchanges)

	}

	r.Use(static.Serve("/v1/chart", static.LocalFile("/charts", true)))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
