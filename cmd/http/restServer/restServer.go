package main

import (
	"time"

	"github.com/diadata-org/diadata/pkg/utils"

	jwt "github.com/appleboy/gin-jwt/v2"
	//jwt "github.com/blockstatecom/gin-jwt"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/kafkaHelper"
	"github.com/diadata-org/diadata/pkg/http/restServer/diaApi"
	"github.com/diadata-org/diadata/pkg/http/restServer/kafkaApi"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Title diadata.org API
// @Version 1.0
// @description The world's crowd-driven financial data community has a professional API made for you.
// @description <h2>Decentral and transparent by design</h2>
// @description With our decentral approach to data verification, you can gain a deep insight into current and past pricing, volume and exchange info so you can make the right decisions to stay ahead of the game.
// @description
// @description <h3>Find the right data for your needs</h3>
// @description Show your users the most transparent data on the market with our API. Whether you're building a financial service, a portfolio management tool, a new media offering, or more, we have the most advanced and updated data on the market for your product.
// @description For Oracle usage see [github](https://github.com/diadata-org/diadata/tree/master/documentation/methodology/oracles.md).
// @description
// @description <h3>Backtest your strategies</h3>
// @description Use the most efficient and transparent crypto data to run simulations and backtest your trading or investing strategies. With crowd-aggregated hundreds of exchanges you can be sure that you're getting the right picture every single time.
// @description
// @description <h3>Run Experiments</h3>
// @description Build your own models with our data, to further your interest or just for fun. With our flexible and powerful API, we provide you with a set of data that will help you draw insights and make conclusions.
// @description
// @description <h3>Request your data</h3>
// @description Set a bounty on gitcoin.io or drop us [line](mailto:API@diadata.org).
// @license.name GNU GPLv3
// @Host api.diadata.org
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

func GetTradesBlock(c *gin.Context) {
	kafkaApi.Process(c, kafkaHelper.TopicTradesBlock)
}

func GetFiltersBlock(c *gin.Context) {
	kafkaApi.Process(c, kafkaHelper.TopicFiltersBlock)
}

func GetTrades(c *gin.Context) {
	kafkaApi.Process(c, kafkaHelper.TopicTrades)
}

const (
	cachingTime1Sec   = 1 * time.Second
	cachingTime20Secs = 20 * time.Second
	cachingTimeShort  = time.Minute * 2
	// cachingTimeMedium = time.Minute * 10
	cachingTimeLong = time.Minute * 100
)

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}

func main() {

	r := gin.New()
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

			if userID == utils.Getenv("HTTP_BASIC_AUTH_USER", config.ApiKey) &&
				password == utils.Getenv("HTTP_BASIC_AUTH_PASSWD", config.SecretKey) { // Temporary: only 1 valid key so far.
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
	if err != nil {
		log.Error("creating middleware: ", err)
	}

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

	memoryStore := persistence.NewInMemoryStore(time.Second)

	store, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}
	relStore, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore", err)
	}
	diaApiEnv := &diaApi.Env{
		DataStore: store,
		RelDB:     *relStore,
	}

	diaAuth := r.Group("/v1")
	diaAuth.Use(authMiddleware.MiddlewareFunc())
	{
		diaAuth.POST("/supply", diaApiEnv.PostSupply)
		diaAuth.POST("/quotation", diaApiEnv.SetQuotation)
	}

	diaGroup := r.Group("/v1")
	{
		// Trades and prices endpoints.
		diaGroup.GET("/quotation/:symbol", cache.CachePageAtomic(memoryStore, cachingTime20Secs, diaApiEnv.GetQuotation))
		diaGroup.GET("/assetQuotation/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTime20Secs, diaApiEnv.GetAssetQuotation))
		diaGroup.GET("/lastTradeTime/:exchange/:blockchain/:address", diaApiEnv.GetLastTradeTime)
		diaGroup.GET("/lastTradesAsset/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetLastTradesAsset))

		// Filters endpoints.
		diaGroup.GET("/chartPoints/:filter/:exchange/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetChartPoints))
		diaGroup.GET("/assetChartPoints/:filter/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetAssetChartPoints))
		diaGroup.GET("/chartPointsAllExchanges/:filter/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetChartPointsAllExchanges))

		// Supply endpoints.
		diaGroup.GET("/supply/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetSupply))
		diaGroup.GET("/assetSupply/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetAssetSupply))
		diaGroup.GET("/supplies/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetSupplies))

		// Asset endpoints.
		diaGroup.GET("/topAssets/:numAssets", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetTopAssets))
		diaGroup.GET("/symbols", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetAllSymbols))
		diaGroup.GET("/symbols/:substring", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetAllSymbols))
		diaGroup.GET("/quotedAssets", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetQuotedAssets))

		// (DEX) pools/liquidity endpoints.
		diaGroup.GET("/poolLiquidity/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetPoolLiquidityByAddress))

		// Pairs endpoints
		diaGroup.GET("/pairsCex/:exchange", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetExchangePairs))
		diaGroup.GET("/pairsAssetCex/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetAssetPairs))

		// Volume endpoints.
		diaGroup.GET("/volume/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetVolume))
		diaGroup.GET("/volume24/:exchange", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.Get24hVolume))
		diaGroup.GET("/feedStats/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetFeedStats))

		// Other endpoints.
		diaGroup.GET("/search/:query", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.SearchAsset))
		diaGroup.GET("/searchnft/:query", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.SearchNFTs))
		diaGroup.GET("/assetInfo/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetAssetInfo))
		diaGroup.GET("/token/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetAsset))

		diaGroup.GET("/missingToken/:exchange", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetMissingExchangeSymbol))
		diaGroup.GET("/tokenexchanges/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetAssetExchanges))

		diaGroup.GET("/exchanges", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetExchanges))
		diaGroup.GET("/NFT/exchanges", cache.CachePageAtomic(memoryStore, cachingTime1Sec, diaApiEnv.GetNFTExchanges))

		diaGroup.GET("/blockchains", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetAllBlockchains))

		// Endpoints for interestrates
		// diaGroup.GET("/interestrates", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetRates))
		// diaGroup.GET("/interestrate/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetInterestRate))
		// diaGroup.GET("/interestrate/:symbol/:time", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetInterestRate))
		// diaGroup.GET("/compoundedRate/:symbol/:dpy", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetCompoundedRate))
		// diaGroup.GET("/compoundedRate/:symbol/:dpy/:time", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetCompoundedRate))
		// diaGroup.GET("/compoundedAvg/:symbol/:days/:dpy", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetCompoundedAvg))
		// diaGroup.GET("/compoundedAvg/:symbol/:days/:dpy/:time", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetCompoundedAvg))
		// diaGroup.GET("/compoundedAvgDIA/:symbol/:days/:dpy", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetCompoundedAvgDIA))
		// diaGroup.GET("/compoundedAvgDIA/:symbol/:days/:dpy/:time", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetCompoundedAvgDIA))

		// Endpoints for fiat currencies
		diaGroup.GET("/fiatQuotations", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetFiatQuotations))

		// // Endpoints for stocks
		// dia.GET("/stockSymbols", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetStockSymbols))
		// dia.GET("/stockQuotation/:source/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetStockQuotation))
		// dia.GET("/stockQuotation/:source/:symbol/:time", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetStockQuotation))

		// Endpoints for foreign sources
		diaGroup.GET("/foreignQuotation/:source/:symbol", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetForeignQuotation))
		diaGroup.GET("/foreignQuotation/:source/:symbol/:time", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetForeignQuotation))
		diaGroup.GET("/foreignSymbols/:source", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetForeignSymbols))

		// Endpoints for customized products
		diaGroup.GET("/custom/vwapFirefly/:ticker", cache.CachePageAtomic(memoryStore, cachingTime20Secs, diaApiEnv.GetVwapFirefly))

		// Index
		diaGroup.GET("/benchmarkedIndexValue/:symbol", diaApiEnv.GetBenchmarkedIndexValue)

		// External supply reports
		diaGroup.GET("/diaTotalSupply", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetDiaTotalSupply))
		diaGroup.GET("/diaCirculatingSupply", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetDiaCirculatingSupply))

		// NFT endpoints.
		diaGroup.GET("/AllNFTClasses/:blockchain", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetAllNFTClasses))
		diaGroup.GET("/NFTClasses/:limit/:offset", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTClasses))
		diaGroup.GET("/NFTCategories", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTCategories))
		diaGroup.GET("/NFT/:blockchain/:address/:id", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFT))
		diaGroup.GET("/NFTTrades/:blockchain/:address/:id", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTTrades))
		diaGroup.GET("/NFTTradesCollection/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTTradesCollection))
		diaGroup.GET("/NFTFloor/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTFloor))
		diaGroup.GET("/NFTFloorMA/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTFloorMA))
		diaGroup.GET("/NFTDownday/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTDownday))
		diaGroup.GET("/NFTVolatility/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTFloorVola))
		diaGroup.GET("/topNFT/:numCollections", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetTopNFTClasses))
		diaGroup.GET("/NFTVolume/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetNFTVolume))
		diaGroup.GET("/assetmap/:blockchain/:address", cache.CachePageAtomic(memoryStore, cachingTimeLong, diaApiEnv.GetAssetMap))
		diaGroup.GET("/assetUpdates/:blockchain/:address/:deviation/:frequencySeconds", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetAssetUpdates))

		// Endpoints for Synthassets

		diaGroup.GET("/synthasset/:blockchain/:protocol", cache.CachePageAtomic(memoryStore, cachingTimeShort, diaApiEnv.GetSyntheticAsset))

	}

	r.Use(static.Serve("/v1/chart", static.LocalFile("/charts", true)))

	AddEndpoints(r)

	// This environment variable is either set in docker-compose or empty
	executionMode := utils.Getenv("EXEC_MODE", "")
	if executionMode == "production" {
		err = r.Run(utils.Getenv("LISTEN_PORT", ":8080"))
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
