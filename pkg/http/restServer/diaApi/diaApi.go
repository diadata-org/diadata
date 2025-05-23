package diaApi

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/gqlclient"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	DECIMALS_CACHE                = make(map[dia.Asset]uint8)
	ASSET_CACHE                   = make(map[string]dia.Asset)
	QUOTATION_CACHE               = make(map[string]*models.AssetQuotation)
	BLOCKCHAINS                   = make(map[string]dia.BlockChain)
	ASSETQUOTATION_LOOKBACK_HOURS = 24 * 7
)

type Env struct {
	DataStore models.Datastore
	RelDB     models.RelDB
	signer    *utils.AssetQuotationSigner
}

func init() {
	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("Error connecting to asset DB: ", err)
		return
	}
	chains, err := relDB.GetAllBlockchains(false)
	if err != nil {
		log.Fatal("get all chains: ", err)
	}
	for _, chain := range chains {
		BLOCKCHAINS[chain.Name] = chain
	}
}

func NewEnv(ds models.Datastore, rdb models.RelDB, signer *utils.AssetQuotationSigner) *Env {
	return &Env{DataStore: ds, RelDB: rdb, signer: signer}
}

// GetAssetQuotation returns quotation of asset with highest market cap among
// all assets with symbol ticker @symbol.
func (env *Env) GetAssetQuotation(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	var (
		err               error
		asset             dia.Asset
		quotationExtended models.AssetQuotationFull
	)

	// Time for quotation is now by default.
	timestampInt, err := strconv.ParseInt(c.DefaultQuery("timestamp", strconv.Itoa(int(time.Now().Unix()))), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("could not parse Unix timestamp"))
		return
	}
	timestamp := time.Unix(timestampInt, 0)

	// An asset is uniquely defined by blockchain and address.
	asset, err = env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	// Get quotation for asset.
	quotation, err := env.DataStore.GetAssetQuotation(asset, dia.CRYPTO_ZERO_UNIX_TIME, timestamp)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	quotationYesterday, err := env.DataStore.GetAssetQuotation(asset, dia.CRYPTO_ZERO_UNIX_TIME, timestamp.AddDate(0, 0, -1))
	if err != nil {
		log.Warn("get quotation yesterday: ", err)
	} else {
		quotationExtended.PriceYesterday = quotationYesterday.Price
	}

	volumeYesterday, err := env.DataStore.Get24HoursAssetVolume(asset)
	if err != nil {
		log.Warn("get volume yesterday: ", err)
	} else {
		quotationExtended.VolumeYesterdayUSD = *volumeYesterday
	}

	// Appropriate formatting.
	quotationExtended.Symbol = quotation.Asset.Symbol
	quotationExtended.Name = quotation.Asset.Name
	quotationExtended.Address = quotation.Asset.Address
	quotationExtended.Blockchain = quotation.Asset.Blockchain
	quotationExtended.Price = quotation.Price
	quotationExtended.Time = quotation.Time
	quotationExtended.Source = quotation.Source

	signedData, err := env.signer.Sign(quotation.Asset.Symbol, quotation.Asset.Address, quotation.Asset.Blockchain, quotation.Price, quotationExtended.Time)
	if err != nil {
		log.Warn("error signing data: ", err)
	}
	quotationExtended.Signature = signedData

	c.JSON(http.StatusOK, quotationExtended)

}

// GetQuotation returns quotation of asset with highest market cap among
// all assets with symbol ticker @symbol.
func (env *Env) GetQuotation(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	symbol := c.Param("symbol")

	timestamp := time.Now()
	var quotationExtended models.AssetQuotationFull
	// Fetch underlying assets for symbol
	assets, err := env.RelDB.GetTopAssetByVolume(symbol)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}
	if len(assets) == 0 {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	topAsset := assets[0]
	quotation, err := env.DataStore.GetAssetQuotation(topAsset, dia.CRYPTO_ZERO_UNIX_TIME, timestamp)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	quotationYesterday, err := env.DataStore.GetAssetQuotation(topAsset, dia.CRYPTO_ZERO_UNIX_TIME, timestamp.AddDate(0, 0, -1))
	if err != nil {
		log.Warn("get quotation yesterday: ", err)
	} else {
		quotationExtended.PriceYesterday = quotationYesterday.Price
	}
	volumeYesterday, err := env.DataStore.Get24HoursAssetVolume(topAsset)
	if err != nil {
		log.Warn("get volume yesterday: ", err)
	} else {
		quotationExtended.VolumeYesterdayUSD = *volumeYesterday
	}
	quotationExtended.Symbol = quotation.Asset.Symbol
	quotationExtended.Name = quotation.Asset.Name
	quotationExtended.Address = quotation.Asset.Address
	quotationExtended.Blockchain = quotation.Asset.Blockchain
	quotationExtended.Price = quotation.Price
	quotationExtended.Time = quotation.Time
	quotationExtended.Source = quotation.Source

	c.JSON(http.StatusOK, quotationExtended)
}

func (env *Env) GetAssetMap(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	timestamp := time.Now()
	var quotations []models.AssetQuotationFull
	// Fetch underlying assets for symbol
	asset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	// get assetid
	assetid, err := env.RelDB.GetAssetID(asset)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	// get groupId
	group_id, err := env.RelDB.GetAssetMap(assetid)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	assets, err := env.RelDB.GetAssetByGroupID(group_id)
	if err != nil || len(assets) == 0 {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}

	for _, topAsset := range assets {
		var quotationExtended models.AssetQuotationFull

		quotation, err := env.DataStore.GetAssetQuotation(topAsset, dia.CRYPTO_ZERO_UNIX_TIME, timestamp)
		if err != nil {
			log.Warn("get quotation: ", err)
		}
		quotationYesterday, err := env.DataStore.GetAssetQuotation(topAsset, dia.CRYPTO_ZERO_UNIX_TIME, timestamp.AddDate(0, 0, -1))
		if err != nil {
			log.Warn("get quotation yesterday: ", err)
		} else {
			quotationExtended.PriceYesterday = quotationYesterday.Price
		}
		volumeYesterday, err := env.RelDB.GetLastAssetVolume24H(topAsset)
		if err != nil {
			log.Warn("get volume yesterday: ", err)
		} else {
			quotationExtended.VolumeYesterdayUSD = volumeYesterday
		}
		quotationExtended.Symbol = topAsset.Symbol
		quotationExtended.Name = topAsset.Name
		quotationExtended.Address = topAsset.Address
		quotationExtended.Blockchain = topAsset.Blockchain
		quotationExtended.Price = quotation.Price
		quotationExtended.Time = quotation.Time
		quotationExtended.Source = quotation.Source
		quotations = append(quotations, quotationExtended)
	}

	c.JSON(http.StatusOK, quotations)
}

// GetSupply returns latest supply of token with @symbol
func (env *Env) GetSupply(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	symbol := c.Param("symbol")

	s, err := env.DataStore.GetLatestSupply(symbol, &env.RelDB)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, s)
	}
}

// GetSupply returns latest supply of token with @symbol
func (env *Env) GetAssetSupply(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*time.Hour))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("parse time range"))
		return
	}

	if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 30*24*time.Hour))
		return
	}

	values, err := env.DataStore.GetSupplyInflux(dia.Asset{Address: address, Blockchain: blockchain}, starttime, endtime)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
			return
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
	}

	// Fetch decimals from local cache implementation.
	for i := range values {
		values[i].Asset.Decimals = env.getDecimalsFromCache(DECIMALS_CACHE, values[i].Asset)
	}

	if len(values) == 1 {
		c.JSON(http.StatusOK, values[0])
	} else {
		c.JSON(http.StatusOK, values)
	}

}

// GetSupplies returns a time range of supplies of token with @symbol
func (env *Env) GetSupplies(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	symbol := c.Param("symbol")
	starttimeStr := c.DefaultQuery("starttime", "noRange")
	endtimeStr := c.Query("endtime")

	var starttime, endtime time.Time

	if starttimeStr == "noRange" || endtimeStr == "" {
		endtime = time.Now()
		starttime = endtime.AddDate(0, 0, -30)
	} else {
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	}
	if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 30*24*time.Hour))
		return
	}

	s, err := env.DataStore.GetSupply(symbol, starttime, endtime, &env.RelDB)
	if len(s) == 0 {
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, s)
}

func (env *Env) GetDiaTotalSupply(c *gin.Context) {
	q, err := env.DataStore.GetDiaTotalSupply()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
	}
}

func (env *Env) GetDiaCirculatingSupply(c *gin.Context) {
	q, err := env.DataStore.GetDiaCirculatingSupply()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
	}
}

// Get24hVolume if no times are set use the last 24h
func (env *Env) Get24hVolume(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	v, err := env.DataStore.Get24HoursExchangeVolume(c.Param("exchange"))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, *v)
}

// GetExchanges is the delegate method for fetching all exchanges available in Postgres.
func (env *Env) GetExchanges(c *gin.Context) {

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*time.Hour))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("parse time range"))
		return
	}
	if starttime.Before(endtime.AddDate(0, 0, -30)) {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("time range is limited to 30 days"))
		return
	}

	type exchangeReturn struct {
		Name          string
		Volume24h     float64
		Trades        int64
		Pairs         int
		Type          string
		Blockchain    string
		ScraperActive bool
	}
	var exchangereturns []exchangeReturn
	exchanges, err := env.RelDB.GetAllExchanges()
	if len(exchanges) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	dexPoolCountMap, err := env.RelDB.GetAllDEXPoolsCount()
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	for _, exchange := range exchanges {
		var numPairs int

		vol, err := env.DataStore.GetVolumeInflux(dia.Asset{}, exchange.Name, starttime, endtime)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}

		numTrades, err := env.DataStore.GetNumTrades(exchange.Name, "", "", starttime, endtime)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}

		if models.GetExchangeType(exchange) == "DEX" {
			numPairs = dexPoolCountMap[exchange.Name]
		} else {
			numPairs, err = env.RelDB.GetNumPairs(exchange)
			if err != nil {
				restApi.SendError(c, http.StatusInternalServerError, nil)
			}
		}

		exchangereturn := exchangeReturn{
			Name:          exchange.Name,
			Volume24h:     *vol,
			Trades:        numTrades,
			Pairs:         numPairs,
			Blockchain:    exchange.BlockChain.Name,
			ScraperActive: exchange.ScraperActive,
		}
		exchangereturn.Type = models.GetExchangeType(exchange)
		exchangereturns = append(exchangereturns, exchangereturn)

	}

	sort.Slice(exchangereturns, func(i, j int) bool {
		return exchangereturns[i].Volume24h > exchangereturns[j].Volume24h
	})

	c.JSON(http.StatusOK, exchangereturns)
}

// GetAssetChartPoints queries for filter points of asset given by address and blockchain.
func (env *Env) GetAssetChartPoints(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	filter := c.Param("filter")
	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	exchange := c.Query("exchange")

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(7*24*time.Hour))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("parse time range"))
		return
	}

	if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(14*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 14*24*time.Hour))
		return
	}

	p, err := env.DataStore.GetFilterPointsAsset(filter, exchange, address, blockchain, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// GetChartPoints returns Filter points for given symbol -> Deprecated?
func (env *Env) GetChartPoints(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	filter := c.Param("filter")
	exchange := c.Param("exchange")
	symbol := c.Param("symbol")
	scale := c.Query("scale")

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(7*24*time.Hour))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("parse time range"))
		return
	}

	if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 30*24*time.Hour))
		return
	}

	p, err := env.DataStore.GetFilterPoints(filter, exchange, symbol, scale, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// GetChartPointsAllExchanges returns filter points across all exchanges.
func (env *Env) GetChartPointsAllExchanges(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	filter := c.Param("filter")
	symbol := c.Param("symbol")
	scale := c.Query("scale")

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(7*24*time.Hour))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("parse time range"))
		return
	}

	if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 30*24*time.Hour))
		return
	}

	p, err := env.DataStore.GetFilterPoints(filter, "", symbol, scale, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}

}

func (env *Env) GetFilterPerSource(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	type priceOnExchange struct {
		Price     float64   `json:"Price"`
		Exchange  string    `json:"Exchange"`
		Timestamp time.Time `json:"Time"`
	}

	type localReturn struct {
		Asset  dia.Asset         `json:"Asset"`
		Prices []priceOnExchange `json:"PricePerExchange"`
	}

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)
	filter := c.Param("filter")

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(30)*time.Minute)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
		return
	}

	assetQuotations, err := env.DataStore.GetFilterAllExchanges(filter, address, blockchain, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
		return
	}

	var lr localReturn
	lr.Asset = env.getAssetFromCache(ASSET_CACHE, blockchain, address)

	for _, aq := range assetQuotations {
		var pe priceOnExchange
		pe.Exchange = aq.Source
		pe.Price = aq.Price
		pe.Timestamp = aq.Time
		lr.Prices = append(lr.Prices, pe)
	}
	c.JSON(http.StatusOK, lr)

}

// GetAllSymbols returns all Symbols on @exchange.
// If @exchange is not set, it returns all symbols across all exchanges.
// If @top is set to an integer, only the top @top symbols w.r.t. trading volume are returned. This is
// only active if @exchange is not set.
func (env *Env) GetAllSymbols(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	var (
		s            []string
		numSymbols   int64
		sortedAssets []dia.AssetVolume
		err          error
	)

	substring := c.Param("substring")
	exchange := c.DefaultQuery("exchange", "noRange")
	numSymbolsString := c.Query("top")

	if numSymbolsString != "" {
		numSymbols, err = strconv.ParseInt(numSymbolsString, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, errors.New("number of symbols must be an integer"))
		}
	}

	// Filter results by substring. @exchange is disabled.
	if substring != "" {
		s, err = env.RelDB.GetExchangeSymbols("", substring)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find symbols"))
		}
		s = utils.UniqueStrings(s)

		sort.Strings(s)
		// Sort all symbols by volume, append if they have no volume.
		sortedAssets, err = env.RelDB.GetSortedAssetSymbols(int64(0), int64(0), substring)
		if err != nil {
			log.Error("get assets with volume: ", err)
		}
		var sortedSymbols []string
		for _, assetvol := range sortedAssets {
			sortedSymbols = append(sortedSymbols, assetvol.Asset.Symbol)
		}
		sortedSymbols = utils.UniqueStrings(sortedSymbols)
		allSymbols := utils.UniqueStrings(append(sortedSymbols, s...))

		c.JSON(http.StatusOK, allSymbols)
		return
	}

	if exchange == "noRange" {
		if numSymbolsString != "" {
			// -- Get top @numSymbols symbols across all exchanges. --
			sortedAssets, err = env.RelDB.GetAssetsWithVOL(time.Now().AddDate(0, -1, 0), numSymbols, int64(0), false, "")
			if err != nil {
				log.Error("get assets with volume: ", err)
			}
			for _, assetvol := range sortedAssets {
				s = append(s, assetvol.Asset.Symbol)
			}
			c.JSON(http.StatusOK, s)
		} else {
			// -- Get all symbols across all exchanges. --
			s, err = env.RelDB.GetExchangeSymbols("", "")
			if err != nil {
				restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find symbols"))
			}
			s = utils.UniqueStrings(s)

			sort.Strings(s)
			// Sort all symbols by volume, append if they have no volume.
			sortedAssets, err = env.RelDB.GetAssetsWithVOL(time.Now().AddDate(0, -1, 0), numSymbols, int64(0), false, "")
			if err != nil {
				log.Error("get assets with volume: ", err)
			}
			var sortedSymbols []string
			for _, assetvol := range sortedAssets {
				sortedSymbols = append(sortedSymbols, assetvol.Asset.Symbol)
			}
			sortedSymbols = utils.UniqueStrings(sortedSymbols)
			allSymbols := utils.UniqueStrings(append(sortedSymbols, s...))

			c.JSON(http.StatusOK, allSymbols)
		}
	} else {
		// -- Get all symbols on @exchange. --
		symbols, err := env.RelDB.GetExchangeSymbols(exchange, "")
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find symbols"))
		}
		c.JSON(http.StatusOK, symbols)
	}

}

// -----------------------------------------------------------------------------
// POOLS AND LIQUIDITY
// -----------------------------------------------------------------------------

func (env *Env) GetPoolsByAsset(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)
	asset := env.getAssetFromCache(ASSET_CACHE, blockchain, address)

	liquidityThreshold, err := strconv.ParseFloat(c.DefaultQuery("liquidityThreshold", "10"), 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot parse liquidityThreshold"))
		return
	}

	liquidityThresholdUSD, err := strconv.ParseFloat(c.DefaultQuery("liquidityThresholdUSD", "10000"), 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot parse liquidityThresholdUSD"))
		return
	}

	// Set liquidity threshold measured in native currency to 1 in order to filter out noise.
	pools, err := env.RelDB.GetPoolsByAsset(asset, liquidityThreshold, 0)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find pool"))
		return
	}

	type poolInfo struct {
		Exchange          string
		Blockchain        string
		Address           string
		Time              time.Time
		TotalLiquidityUSD float64
		Message           string
		Liquidity         []dia.AssetLiquidity
	}
	var result []poolInfo

	// Get total liquidity for each filtered pool.
	for _, pool := range pools {
		var pi poolInfo

		totalLiquidity, lowerBound := pool.GetPoolLiquidityUSD()

		// In case we can determine USD liquidity and it's below the threshold, continue.
		if !lowerBound && totalLiquidity < liquidityThresholdUSD {
			continue
		}

		pi.Exchange = pool.Exchange.Name
		pi.Blockchain = pool.Blockchain.Name
		pi.Address = pool.Address
		pi.TotalLiquidityUSD = totalLiquidity
		pi.Time = pool.Time
		for i := range pool.Assetvolumes {
			var al dia.AssetLiquidity = dia.AssetLiquidity(pool.Assetvolumes[i])
			pi.Liquidity = append(pi.Liquidity, al)
		}
		if lowerBound {
			pi.Message = "No US-Dollar price information on one or more pool assets available."
		}
		result = append(result, pi)
	}

	// Sort by total USD liquidity.
	sort.Slice(result, func(m, n int) bool {
		return result[m].TotalLiquidityUSD > result[n].TotalLiquidityUSD
	})

	c.JSON(http.StatusOK, result)
}

func (env *Env) GetPoolLiquidityByAddress(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	pool, err := env.RelDB.GetPoolByAddress(blockchain, address)
	if err != nil {
		log.Info("err: ", err)
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find pool"))
		return
	}

	// Get total liquidity.
	var (
		totalLiquidity float64
		lowerBound     bool
	)
	totalLiquidity, lowerBound = pool.GetPoolLiquidityUSD()

	type localReturn struct {
		Exchange          string
		Blockchain        string
		Address           string
		Time              time.Time
		TotalLiquidityUSD float64
		Message           string
		Liquidity         []dia.AssetLiquidity
	}

	var l localReturn
	if lowerBound {
		l.Message = "No US-Dollar price information on one or more pool assets available."
	}
	l.TotalLiquidityUSD = totalLiquidity
	l.Exchange = pool.Exchange.Name
	l.Blockchain = pool.Blockchain.Name
	l.Address = pool.Address
	l.Time = pool.Time
	for i := range pool.Assetvolumes {
		var al dia.AssetLiquidity = dia.AssetLiquidity(pool.Assetvolumes[i])
		l.Liquidity = append(l.Liquidity, al)
	}

	c.JSON(http.StatusOK, l)

}

func (env *Env) GetPoolSlippage(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	blockchain := c.Param("blockchain")
	addressPool := normalizeAddress(c.Param("addressPool"), blockchain)
	addressAsset := normalizeAddress(c.Param("addressAsset"), blockchain)
	poolType := c.Param("poolType")
	priceDeviationInt, err := strconv.ParseInt(c.Param("priceDeviation"), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("error parsing priceDeviation"))
		return
	}
	if priceDeviationInt < 0 || priceDeviationInt >= 1000 {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("priceDeviation measured in per mille is out of range"))
		return
	}
	priceDeviation := float64(priceDeviationInt) / 1000

	type localReturn struct {
		VolumeRequired float64
		AssetIn        string
		Exchange       string
		Blockchain     string
		Address        string
		Time           time.Time
		Liquidity      []dia.AssetLiquidity
	}

	pool, err := env.RelDB.GetPoolByAddress(blockchain, addressPool)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find pool"))
		return
	}
	var l localReturn
	l.Exchange = pool.Exchange.Name
	l.Blockchain = pool.Blockchain.Name
	l.Address = pool.Address
	l.Time = pool.Time
	for i := range pool.Assetvolumes {
		var al dia.AssetLiquidity = dia.AssetLiquidity(pool.Assetvolumes[i])
		l.Liquidity = append(l.Liquidity, al)
	}

	var (
		assetInIndex int
		foundAsset   bool
	)
	for i := range pool.Assetvolumes {
		if pool.Assetvolumes[i].Asset.Address == addressAsset {
			assetInIndex = i
			foundAsset = true
		}
	}
	if !foundAsset {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("asset %s not in pool", addressAsset))
		return
	}
	l.AssetIn = pool.Assetvolumes[assetInIndex].Asset.Symbol

	switch poolType {
	case "UniswapV2":
		l.VolumeRequired = pool.Assetvolumes[assetInIndex].Volume * (1/(1-priceDeviation) - 1)
	}

	c.JSON(http.StatusOK, l)
}

func (env *Env) GetPoolPriceImpact(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	blockchain := c.Param("blockchain")
	addressPool := normalizeAddress(c.Param("addressPool"), blockchain)
	addressAsset := normalizeAddress(c.Param("addressAsset"), blockchain)
	poolType := c.Param("poolType")
	priceDeviationInt, err := strconv.ParseInt(c.Param("priceDeviation"), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("error parsing priceDeviation"))
		return
	}
	if priceDeviationInt < 0 || priceDeviationInt >= 1000 {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("priceDeviation measured in per mille is out of range"))
		return
	}
	priceDeviation := float64(priceDeviationInt) / 1000

	type localReturn struct {
		VolumeRequired float64
		AssetIn        string
		Exchange       string
		Blockchain     string
		Address        string
		Time           time.Time
		Liquidity      []dia.AssetLiquidity
	}

	pool, err := env.RelDB.GetPoolByAddress(blockchain, addressPool)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find pool"))
		return
	}
	var l localReturn
	l.Exchange = pool.Exchange.Name
	l.Blockchain = pool.Blockchain.Name
	l.Address = pool.Address
	l.Time = pool.Time
	for i := range pool.Assetvolumes {
		var al dia.AssetLiquidity = dia.AssetLiquidity(pool.Assetvolumes[i])
		l.Liquidity = append(l.Liquidity, al)
	}

	var (
		assetInIndex int
		foundAsset   bool
	)
	for i := range pool.Assetvolumes {
		if pool.Assetvolumes[i].Asset.Address == addressAsset {
			assetInIndex = i
			foundAsset = true
		}
	}
	if !foundAsset {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("asset %s not in pool", addressAsset))
		return
	}
	l.AssetIn = pool.Assetvolumes[assetInIndex].Asset.Symbol

	switch poolType {
	case "UniswapV2":
		l.VolumeRequired = pool.Assetvolumes[assetInIndex].Volume * (1/math.Sqrt(1-priceDeviation) - 1)
	}

	c.JSON(http.StatusOK, l)
}

func (env *Env) GetPriceImpactSimulation(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	poolType := c.Param("poolType")
	priceDeviationInt, err := strconv.ParseInt(c.Param("priceDeviation"), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("error parsing priceDeviation"))
		return
	}
	if priceDeviationInt < 0 || priceDeviationInt >= 1000 {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("priceDeviation measured in per mille is out of range"))
		return
	}
	priceDeviation := float64(priceDeviationInt) / 1000
	liquidityA, err := strconv.ParseFloat(c.Param("liquidityA"), 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("error parsing liquidityA"))
		return
	}
	if liquidityA <= 0 {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("liquidity must be a non-negative number"))
		return
	}
	liquidityB, err := strconv.ParseFloat(c.Param("liquidityB"), 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("error parsing liquidityB"))
		return
	}
	if liquidityB <= 0 {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("liquidity must be a non-negative number"))
		return
	}

	type dummyLiquidity struct {
		Asset     string
		Liquidity float64
	}

	type localReturn struct {
		PriceDeviation  float64
		PriceAssetA     float64
		PriceAssetB     float64
		VolumesRequired []struct {
			AssetIn                string
			VolumeRequired         float64
			InitialPriceAssetIn    float64
			ResultingPriceAssetIn  float64
			ResultingPriceAssetOut float64
		}
		Liquidity []dummyLiquidity
	}

	l := []dummyLiquidity{
		{Asset: "A", Liquidity: liquidityA},
		{Asset: "B", Liquidity: liquidityB},
	}
	var lr localReturn
	lr.PriceDeviation = priceDeviation
	lr.PriceAssetA = liquidityB / liquidityA
	lr.PriceAssetB = liquidityA / liquidityB

	switch poolType {
	case "UniswapV2":
		volRequiredA := liquidityA * (1/math.Sqrt(1-priceDeviation) - 1)
		volRequiredB := liquidityB * (1/math.Sqrt(1-priceDeviation) - 1)
		lr.VolumesRequired = append(lr.VolumesRequired, struct {
			AssetIn                string
			VolumeRequired         float64
			InitialPriceAssetIn    float64
			ResultingPriceAssetIn  float64
			ResultingPriceAssetOut float64
		}{
			"A",
			volRequiredA,
			liquidityB / liquidityA,
			liquidityA * liquidityB / math.Pow(volRequiredA+liquidityA, 2),
			math.Pow(volRequiredA+liquidityA, 2) / liquidityA / liquidityB,
		})
		lr.Liquidity = l
		lr.VolumesRequired = append(lr.VolumesRequired, struct {
			AssetIn                string
			VolumeRequired         float64
			InitialPriceAssetIn    float64
			ResultingPriceAssetIn  float64
			ResultingPriceAssetOut float64
		}{
			"B",
			volRequiredB,
			liquidityA / liquidityB,
			liquidityB * liquidityA / math.Pow(volRequiredB+liquidityB, 2),
			math.Pow(volRequiredB+liquidityB, 2) / liquidityA / liquidityB,
		})

	}

	c.JSON(http.StatusOK, lr)
}

// -----------------------------------------------------------------------------
// EXCHANGE PAIRS
// -----------------------------------------------------------------------------

func (env *Env) GetExchangePairs(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	exchange, err := env.RelDB.GetExchange(c.Param("exchange"))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	var (
		filterVerified bool
		verified       bool
	)
	verifiedString := c.Query("verified")
	if verifiedString != "" {
		verified, err = strconv.ParseBool(verifiedString)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		filterVerified = true
	}

	pairs, err := env.RelDB.GetPairsForExchange(exchange, filterVerified, verified)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	sort.Slice(pairs, func(m, n int) bool {
		return pairs[m].Symbol < pairs[n].Symbol
	})
	c.JSON(http.StatusOK, pairs)

}

func (env *Env) GetAssetPairs(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)
	var (
		filterVerified bool
		verified       bool
		err            error
	)
	verifiedString := c.Query("verified")
	if verifiedString != "" {
		verified, err = strconv.ParseBool(verifiedString)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		filterVerified = true
	}

	pairs, err := env.RelDB.GetPairsForAsset(dia.Asset{Address: address, Blockchain: blockchain}, filterVerified, verified)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	sort.Slice(pairs, func(m, n int) bool { return pairs[m].Exchange < pairs[n].Exchange })
	c.JSON(http.StatusOK, pairs)

}

func (env *Env) SearchAsset(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	querystring := c.Param("query")
	var (
		assets = []dia.Asset{}
		err    error
	)

	switch {
	case len(querystring) > 4 && strings.Contains(querystring[0:2], "0x"):
		assets, err = env.RelDB.GetAssetsByAddress(querystring)
		if err != nil {
			// restApi.SendError(c, http.StatusInternalServerError, errors.New("eror getting asset"))
			log.Errorln("error getting GetAssetsByAddress", err)
		}

	case len(querystring) > 4 && !strings.Contains(querystring[0:2], "0x"):
		assets, err = env.RelDB.GetAssetsBySymbolName(querystring, querystring)
		if err != nil {
			// restApi.SendError(c, http.StatusInternalServerError, errors.New("eror getting asset"))
			log.Errorln("error getting GetAssetsBySymbolName", err)

		}

	case len(querystring) <= 4:
		assets, err = env.RelDB.GetAssetsBySymbolName(querystring, querystring)
		if err != nil {
			// restApi.SendError(c, http.StatusInternalServerError, errors.New("eror getting asset"))
			log.Errorln("error getting GetAssetsBySymbolName", err)

		}
	}
	c.JSON(http.StatusOK, assets)
}

func (env *Env) GetTopAssets(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	numAssetsString := c.Param("numAssets")
	pageString := c.DefaultQuery("Page", "1")
	onlycexString := c.DefaultQuery("Cex", "false")
	blockchain := c.DefaultQuery("Network", "")

	var (
		numAssets    int64
		sortedAssets []dia.AssetVolume
		err          error
		pageNumber   int64
		offset       int64
	)

	pageNumber, err = strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("page of assets must be an integer"))
	}

	onlycex, err := strconv.ParseBool(onlycexString)
	if err != nil {
		log.Fatal(err)
	}

	numAssets, err = strconv.ParseInt(numAssetsString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("number of assets must be an integer"))
	}

	offset = (pageNumber - 1) * numAssets

	sortedAssets, err = env.RelDB.GetAssetsWithVOL(time.Now().AddDate(0, 0, -7), numAssets, offset, onlycex, blockchain)
	if err != nil {
		log.Error("get assets with volume: ", err)

	}
	var assets = []dia.TopAsset{}

	for _, v := range sortedAssets {
		var sources = make(map[string][]string)

		aqf := dia.TopAsset{}
		aqf.Asset = v.Asset
		quotation, err := env.DataStore.GetAssetQuotationLatest(aqf.Asset, dia.CRYPTO_ZERO_UNIX_TIME)
		if err != nil {
			log.Warn("quotation: ", err)
		} else {
			aqf.Price = quotation.Price

		}
		aqf.Volume = v.Volume

		sources["CEX"], err = env.RelDB.GetAssetSource(v.Asset, true)
		if err != nil {
			log.Warn("get GetAssetSource: ", err)
		}
		sources["DEX"], err = env.RelDB.GetAssetSource(v.Asset, false)
		if err != nil {
			log.Warn("get GetAssetSource: ", err)
		}
		aqf.Source = sources

		quotationYesterday, err := env.DataStore.GetAssetQuotation(aqf.Asset, dia.CRYPTO_ZERO_UNIX_TIME, time.Now().AddDate(0, 0, -1))
		if err != nil {
			log.Warn("get quotation yesterday: ", err)
		} else {
			aqf.PriceYesterday = quotationYesterday.Price
		}

		assets = append(assets, aqf)

	}
	c.JSON(http.StatusOK, assets)
}

// GetQuotedAssets is the delegate method to fetch all assets that have an asset quotation
// dating back at most 7 days.
func (env *Env) GetQuotedAssets(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	endtime := time.Now()
	starttime := endtime.AddDate(0, 0, -7)
	assetvolumes, err := env.RelDB.GetAssetsWithVolByBlockchain(starttime, endtime, c.Query("blockchain"))
	if err != nil {
		log.Error("get assets with volume: ", err)
	}

	c.JSON(http.StatusOK, assetvolumes)
}

// -----------------------------------------------------------------------------
// FIAT CURRENCIES
// -----------------------------------------------------------------------------

// GetFiatQuotations returns several quotations vs USD as published by the ECB
func (env *Env) GetFiatQuotations(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	q, err := env.DataStore.GetCurrencyChange()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
	}
}

func (env *Env) GetTwelvedataFiatQuotations(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	// Parse symbol.
	assets := strings.Split(c.Param("symbol"), "-")
	if len(assets) != 2 {
		restApi.SendError(c, http.StatusNotFound, errors.New("wrong format for forex pair"))
		return
	}
	symbol := assets[0] + "/" + assets[1]

	// Time for quotation is time.Now() by default.
	timestampInt, err := strconv.ParseInt(c.DefaultQuery("timestamp", strconv.Itoa(int(time.Now().Unix()))), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("could not parse Unix timestamp"))
		return
	}
	timestamp := time.Unix(timestampInt, 0)

	var (
		q       models.ForeignQuotation
		errRev  error
		reverse bool
	)

	q, err = env.DataStore.GetForeignQuotationInflux(symbol, "TwelveData", timestamp)
	if err != nil || q.Price == 0 {
		reverse = true
		symbol = assets[1] + "/" + assets[0]
		log.Info("try reverse order: ", symbol)
		q, errRev = env.DataStore.GetForeignQuotationInflux(symbol, "TwelveData", timestamp)
		if errRev != nil || q.Price == 0 {
			if q.Price == 0 {
				errRev = errors.New("not found")
			}
			if errors.Is(errRev, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, errRev)
				return
			} else {
				log.Info(c)
				restApi.SendError(c, http.StatusInternalServerError, errRev)
				return
			}
		}
	}

	response := struct {
		Ticker    string
		Price     float64
		Timestamp time.Time
	}{
		Ticker:    c.Param("symbol"),
		Price:     q.Price,
		Timestamp: q.Time,
	}
	if err == nil && !reverse {
		c.JSON(http.StatusOK, response)
		return
	}
	if errRev == nil && q.Price != 0 {
		response.Price = 1 / q.Price
		c.JSON(http.StatusOK, response)
	}
}

// -----------------------------------------------------------------------------
// STOCKS
// -----------------------------------------------------------------------------

func (env *Env) GetTwelvedataStockQuotations(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	// Time for quotation is time.Now() by default.
	timestampInt, err := strconv.ParseInt(c.DefaultQuery("timestamp", strconv.Itoa(int(time.Now().Unix()))), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("could not parse Unix timestamp"))
		return
	}
	timestamp := time.Unix(timestampInt, 0)

	q, err := env.DataStore.GetForeignQuotationInflux(c.Param("symbol"), "TwelveData", timestamp)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		// Format response.
		response := struct {
			Ticker    string
			Price     float64
			Timestamp time.Time
		}{
			Ticker:    c.Param("symbol"),
			Price:     q.Price,
			Timestamp: q.Time,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (env *Env) GetStockSymbols(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	type sourcedStock struct {
		Stock  models.Stock
		Source string
	}
	var srcStocks []sourcedStock
	stocks, err := env.DataStore.GetStockSymbols()
	log.Info("stocks: ", stocks)

	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		for stock, source := range stocks {
			srcStocks = append(srcStocks, sourcedStock{
				Stock:  stock,
				Source: source,
			})
		}
		c.JSON(http.StatusOK, srcStocks)
	}
}

// GetStockQuotation is the delegate method to fetch the value(s) of
// quotations of asset with @symbol from @source.
// Last value is retrieved. Otional query parameters allow to obtain data in a time range.
func (env *Env) GetStockQuotation(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	source := c.Param("source")
	symbol := c.Param("symbol")
	date := c.Param("time")
	// Add optional query parameters for requesting a range of values
	dateInit := c.DefaultQuery("dateInit", "noRange")
	dateFinal := c.Query("dateFinal")

	if dateInit == "noRange" {
		// Return most recent data point
		var endTime time.Time
		var err error
		if date == "" {
			endTime = time.Now()
		} else {
			// Convert unix time int/string to time
			endTime, err = utils.StrToUnixtime(date)
			if err != nil {
				restApi.SendError(c, http.StatusNotFound, err)
			}
		}
		startTime := endTime.AddDate(0, 0, -1)

		q, err := env.DataStore.GetStockQuotation(source, symbol, startTime, endTime)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q[0])
		}
	} else {
		starttime, err := utils.StrToUnixtime(dateInit)
		if err != nil {
			restApi.SendError(c, http.StatusNotFound, err)
		}
		endtime, err := utils.StrToUnixtime(dateFinal)
		if err != nil {
			restApi.SendError(c, http.StatusNotFound, err)
		}

		q, err := env.DataStore.GetStockQuotation(source, symbol, starttime, endtime)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q)
		}
	}
}

func (env *Env) GetTwelvedataCommodityQuotation(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	timestampInt, err := strconv.ParseInt(c.DefaultQuery("timestamp", strconv.Itoa(int(time.Now().Unix()))), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("could not parse Unix timestamp"))
		return
	}
	timestamp := time.Unix(timestampInt, 0)
	if len(strings.Split(c.Param("symbol"), "-")) != 2 {
		restApi.SendError(c, http.StatusNotFound, errors.New("symbol format not known"))
		return
	}
	symbol := strings.Split(c.Param("symbol"), "-")[0] + "/" + strings.Split(c.Param("symbol"), "-")[1]

	q, err := env.DataStore.GetForeignQuotationInflux(symbol, "TwelveData", timestamp)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		// Format response.
		response := struct {
			Ticker    string
			Name      string
			Price     float64
			Timestamp time.Time
		}{
			Ticker:    c.Param("symbol"),
			Name:      q.Name,
			Price:     q.Price,
			Timestamp: q.Time,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (env *Env) GetTwelvedataETFQuotation(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	timestampInt, err := strconv.ParseInt(c.DefaultQuery("timestamp", strconv.Itoa(int(time.Now().Unix()))), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("could not parse Unix timestamp"))
		return
	}
	timestamp := time.Unix(timestampInt, 0)

	q, err := env.DataStore.GetForeignQuotationInflux(c.Param("symbol"), "TwelveData", timestamp)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		// Format response.
		response := struct {
			Ticker    string
			Name      string
			Price     float64
			Timestamp time.Time
		}{
			Ticker:    c.Param("symbol"),
			Name:      q.Name,
			Price:     q.Price,
			Timestamp: q.Time,
		}
		c.JSON(http.StatusOK, response)
	}
}

// -----------------------------------------------------------------------------
// FOREIGN QUOTATIONS
// -----------------------------------------------------------------------------

// GetForeignQuotation returns several quotations vs USD as published by the ECB
func (env *Env) GetForeignQuotation(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	var (
		q models.ForeignQuotation
	)

	source := c.Param("source")
	symbol := c.Param("symbol")
	date := c.DefaultQuery("time", "noRange")
	var timestamp time.Time

	if date == "noRange" {
		timestamp = time.Now()
	} else {
		t, err := strconv.Atoi(date)
		if err != nil {
			log.Error(err)
		}
		timestamp = time.Unix(int64(t), 0)
	}

	var err error
	q, err = env.DataStore.GetForeignQuotationInflux(symbol, source, timestamp)
	if err != nil || q.Time.Before(time.Unix(1689847252, 0)) {
		// Attempt to fetch quotation for reversed order of symbol string.
		assetsSymbols := strings.Split(symbol, "-")
		if source == "YahooFinance" && len(assetsSymbols) == 2 {
			symbolInflux := assetsSymbols[1] + "-" + assetsSymbols[0]
			q, err = env.DataStore.GetForeignQuotationInflux(symbolInflux, source, timestamp)
			if err != nil {
				restApi.SendError(c, http.StatusInternalServerError, err)
				return
			}
			if q.Price != 0 {
				q.Price = 1 / q.Price
			}
			if q.PriceYesterday != 0 {
				q.PriceYesterday = 1 / q.PriceYesterday
			}
			q.Symbol = symbol
			q.Name = symbol
			c.JSON(http.StatusOK, q)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
	} else {
		c.JSON(http.StatusOK, q)
	}
}

// GetForeignSymbols returns all symbols available for quotation from @source.
func (env *Env) GetForeignSymbols(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	source := c.Param("source")

	q, err := env.DataStore.GetForeignSymbolsInflux(source)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
	}

}

// -----------------------------------------------------------------------------
// CUSTOMIZED PRODUCTS
// -----------------------------------------------------------------------------

func (env *Env) GetVwapFirefly(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	foreignname := c.Param("ticker")
	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")

	var starttime, endtime time.Time
	if starttimeStr == "" || endtimeStr == "" {
		starttime = time.Now().Add(time.Duration(-4 * time.Hour))
		endtime = time.Now()
	} else {
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
		if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
			restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 30*24*time.Hour))
			return
		}
	}

	type vwapObj struct {
		Ticker    string
		Value     float64
		Timestamp time.Time
	}
	values, timestamps, err := env.DataStore.GetVWAPFirefly(foreignname, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	if starttimeStr == "" || endtimeStr == "" {
		response := vwapObj{
			Ticker:    foreignname,
			Value:     values[0],
			Timestamp: timestamps[0],
		}
		c.JSON(http.StatusOK, response)
	} else {
		var response []vwapObj
		for i := 0; i < len(values); i++ {
			tmp := vwapObj{
				Ticker:    foreignname,
				Value:     values[i],
				Timestamp: timestamps[i],
			}
			response = append(response, tmp)
		}
		c.JSON(http.StatusOK, response)
	}
}

func (env *Env) GetLastTradeTime(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	exchange := c.Param("exchange")
	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	t, err := env.DataStore.GetLastTradeTimeForExchange(dia.Asset{Address: address, Blockchain: blockchain}, exchange)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, *t)
	}
}

// GetLastTrades returns last N trades of an asset. Defaults to N=1000.
func (env *Env) GetLastTradesAsset(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	numTradesString := c.DefaultQuery("numTrades", "1000")
	exchange := c.Query("exchange")

	var numTrades int64
	var err error
	numTrades, err = strconv.ParseInt(numTradesString, 10, 64)
	if err != nil {
		numTrades = 1000
	}
	if numTrades > 5000 {
		numTrades = 5000
	}

	asset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	q, err := env.DataStore.GetLastTrades(asset, exchange, time.Now(), int(numTrades), true)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
	}
}

// GetMissingExchangeSymbol returns all unverified symbol
func (env *Env) GetMissingExchangeSymbol(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	exchange := c.Param("exchange")

	//symbols, err := api.GetUnverifiedExchangeSymbols(exchange)
	symbols, err := env.RelDB.GetUnverifiedExchangeSymbols(exchange)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, symbols)
	}
}

func (env *Env) GetAsset(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	symbol := c.Param("symbol")

	symbols, err := env.RelDB.GetAssets(symbol)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, symbols)
	}
}

func (env *Env) GetAssetExchanges(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	symbol := c.Param("symbol")

	symbols, err := env.RelDB.GetAssetExchange(symbol)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, symbols)
	}
}

func (env *Env) GetAllBlockchains(c *gin.Context) {
	blockchains, err := env.RelDB.GetAllAssetsBlockchains()
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, blockchains)
	}
}

func (env *Env) GetFeedStats(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	// Return type for the trades distribution statistics.
	type localDistType struct {
		NumTradesTotal   int64   `json:"NumTradesTotal"`
		NumBins          int     `json:"NumBins"`
		NumLowBins       int     `json:"NumberLowBins"`
		Threshold        int     `json:"Threshold"`
		SizeBinSeconds   int64   `json:"SizeBinSeconds"`
		AvgNumPerBin     float64 `json:"AverageNumberPerBin"`
		StdDeviation     float64 `json:"StandardDeviation"`
		TimeRangeSeconds int64   `json:"TimeRangeSeconds"`
	}

	// Return type for pair volumes per exchange.
	type exchangeVolumes struct {
		Exchange    string
		PairVolumes []dia.PairVolume
	}

	// Overall return type.
	type localReturn struct {
		Timestamp          time.Time
		TotalVolume        float64
		Price              float64
		TradesDistribution localDistType
		ExchangeVolumes    []exchangeVolumes
	}

	// ---- Parse / check input ----

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	// Make starttime and endtime from Unix time input strings.
	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*time.Hour))
	if err != nil {
		log.Error("make timerange: ", err)
	}

	// Check whether time range is feasible.
	if starttime.After(endtime) {
		restApi.SendError(c, http.StatusNotAcceptable, fmt.Errorf("endtime must be later than starttime"))
		return
	}
	if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 24*time.Hour))
		return
	}

	tradesBinThreshold, err := strconv.Atoi(c.DefaultQuery("tradesThreshold", "2"))
	if err != nil {
		log.Warn("parse trades bin threshold: ", err)
		tradesBinThreshold = 2
	}

	sizeBinSeconds, err := strconv.Atoi(c.DefaultQuery("sizeBinSeconds", "120"))
	if err != nil {
		log.Warn("parse sizeBinSeconds: ", err)
		sizeBinSeconds = 120
	}

	volumeThreshold, err := strconv.ParseFloat(c.DefaultQuery("volumeThreshold", "0"), 64)
	if err != nil {
		log.Warn("parse volumeThreshold: ", err)
	}

	if sizeBinSeconds < 20 || sizeBinSeconds > 21600 {
		restApi.SendError(
			c,
			http.StatusInternalServerError,
			fmt.Errorf("sizeBinSeconds out of range. Must be between %v and %v", 20*time.Second, 6*time.Hour),
		)
		return
	}

	// ---- Get data for input ----

	asset := env.getAssetFromCache(ASSET_CACHE, blockchain, address)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	volumeMap, err := env.DataStore.GetExchangePairVolumes(asset, starttime, endtime, volumeThreshold)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
		return
	}

	numTradesSeries, err := env.DataStore.GetNumTradesSeries(asset, "", starttime, endtime, strconv.Itoa(sizeBinSeconds)+"s")
	if err != nil {
		log.Error("get number of trades' series: ", err)
	}

	// ---- Fill return types with fetched data -----

	var (
		result localReturn
		// tradesDistribution localDistType
		ev []exchangeVolumes
	)

	for key, value := range volumeMap {

		var e exchangeVolumes
		e.Exchange = key
		e.PairVolumes = value

		// Collect total volume and full asset information.
		for i, v := range value {
			result.TotalVolume += v.Volume
			e.PairVolumes[i].Pair.QuoteToken = env.getAssetFromCache(ASSET_CACHE, blockchain, address)
			e.PairVolumes[i].Pair.BaseToken = env.getAssetFromCache(ASSET_CACHE, v.Pair.BaseToken.Blockchain, v.Pair.BaseToken.Address)
		}

		// Sort pairs per exchange by volume.
		aux := value
		sort.Slice(aux, func(k, l int) bool {
			return aux[k].Volume > aux[l].Volume
		})
		ev = append(ev, e)

		// Sort exchanges by volume.
		sort.Slice(ev, func(k, l int) bool {
			var ExchangeSums []float64
			for _, exchange := range ev {
				var S float64
				for _, vol := range exchange.PairVolumes {
					S += vol.Volume
				}
				ExchangeSums = append(ExchangeSums, S)
			}
			return ExchangeSums[k] > ExchangeSums[l]
		})

	}

	result.ExchangeVolumes = ev
	result.Timestamp = endtime
	result.Price, err = env.DataStore.GetAssetPriceUSD(asset, endtime.Add(-time.Duration(ASSETQUOTATION_LOOKBACK_HOURS)*time.Hour), endtime)
	if err != nil {
		log.Warn("get price for asset: ", err)
	}

	// Trades Distribution values.
	result.TradesDistribution.Threshold = tradesBinThreshold
	result.TradesDistribution.NumBins = len(numTradesSeries)
	result.TradesDistribution.SizeBinSeconds = int64(sizeBinSeconds)
	var numTradesSeriesFloat []float64
	for _, num := range numTradesSeries {
		numTradesSeriesFloat = append(numTradesSeriesFloat, float64(num))
		result.TradesDistribution.NumTradesTotal += num
		if num < int64(tradesBinThreshold) {
			result.TradesDistribution.NumLowBins++
		}
	}
	if len(volumeMap) == 0 {
		result.TradesDistribution.NumBins = int(endtime.Sub(starttime).Seconds()) / sizeBinSeconds
		result.TradesDistribution.NumLowBins = result.TradesDistribution.NumBins
	}
	if len(numTradesSeries) > 0 {
		result.TradesDistribution.AvgNumPerBin = float64(result.TradesDistribution.NumTradesTotal) / float64(len(numTradesSeries))
	}
	result.TradesDistribution.StdDeviation = utils.StandardDeviation(numTradesSeriesFloat)
	result.TradesDistribution.TimeRangeSeconds = int64(endtime.Sub(starttime).Seconds())

	c.JSON(http.StatusOK, result)

}

// GetAssetUpdates returns the number of updates an oracle with the given parameters
// would have done in the given time-range.
func (env *Env) GetAssetUpdates(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	type localDeviationType struct {
		Time      time.Time `json:"Time"`
		Deviation float64   `json:"Deviation"`
	}
	type localResultType struct {
		UpdateCount   int                  `json:"UpdateCount"`
		UpdatesPer24h float64              `json:"UpdatesPer24h"`
		Asset         dia.Asset            `json:"Asset"`
		Deviations    []localDeviationType `json:"Deviations"`
	}

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	// Deviation in per mille.
	deviation, err := strconv.Atoi(c.Param("deviation"))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	FrequencySeconds, err := strconv.Atoi(c.Param("frequencySeconds"))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*60)*time.Minute)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	if endtime.Sub(starttime) > time.Duration(7*24*60)*time.Minute {
		restApi.SendError(c, http.StatusRequestedRangeNotSatisfiable, errors.New("requested time-range too large"))
		return
	}
	if ok := utils.ValidTimeRange(starttime, endtime, 30*24*time.Hour); !ok {
		restApi.SendError(c, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 30*24*time.Hour))
		return
	}

	asset, errGetAsset := env.RelDB.GetAsset(address, blockchain)
	if errGetAsset != nil {
		restApi.SendError(c, http.StatusInternalServerError, errGetAsset)
		return
	}

	quotations, errGetAssetQuotations := env.DataStore.GetAssetQuotations(asset, starttime, endtime)
	if errGetAssetQuotations != nil {
		restApi.SendError(c, http.StatusInternalServerError, errGetAssetQuotations)
		return
	}

	var lrt localResultType

	lastQuotation := quotations[len(quotations)-1]
	lastValue := lastQuotation.Price
	for i := len(quotations) - 2; i >= 0; i-- {
		var diff float64
		// Oracle did not check for new quotation yet.
		if quotations[i].Time.Sub(lastQuotation.Time) < time.Duration(FrequencySeconds)*time.Second {
			continue
		}
		if lastValue != 0 {
			diff = math.Abs((quotations[i].Price - lastValue) / lastValue)
		} else {
			// Artificially make diff large enough for update (instead of infty).
			diff = float64(deviation)/1000 + 1
		}
		// If deviation is large enough, update values.
		if diff > float64(deviation)/1000 {
			lastQuotation = quotations[i]
			lastValue = lastQuotation.Price

			var ldt localDeviationType
			ldt.Deviation = diff
			ldt.Time = lastQuotation.Time
			lrt.Deviations = append(lrt.Deviations, ldt)
			lrt.UpdateCount++

		}
	}

	lrt.Asset = asset
	lrt.UpdatesPer24h = float64(lrt.UpdateCount) * float64(time.Duration(24*time.Hour).Hours()/endtime.Sub(starttime).Hours())
	c.JSON(http.StatusOK, lrt)
}

// GetAssetInfo returns quotation of asset with highest market cap among
// all assets with symbol ticker @symbol. Additionally information on exchanges and volumes.
func (env *Env) GetAssetInfo(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	type localExchangeInfo struct {
		Name      string
		Volume24h float64
		NumPairs  int
		NumTrades int64
	}

	type localAssetInfoReturn struct {
		Symbol             string
		Name               string
		Address            string
		Blockchain         string
		Price              float64
		PriceYesterday     float64
		VolumeYesterdayUSD float64
		Time               time.Time
		Source             string
		ExchangeInfo       []localExchangeInfo
	}

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*60)*time.Minute)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	var quotationExtended localAssetInfoReturn

	asset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	quotation, err := env.DataStore.GetAssetQuotation(asset, endtime.Add(-time.Duration(ASSETQUOTATION_LOOKBACK_HOURS)*time.Hour), endtime)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	quotationYesterday, err := env.DataStore.GetAssetQuotation(asset, endtime.Add(-time.Duration(ASSETQUOTATION_LOOKBACK_HOURS)*time.Hour), starttime)
	if err != nil {
		log.Warn("get quotation yesterday: ", err)
	} else {
		quotationExtended.PriceYesterday = quotationYesterday.Price
	}
	volumeYesterday, err := env.DataStore.GetVolumeInflux(asset, "", starttime, endtime)
	if err != nil {
		log.Warn("get volume yesterday: ", err)
	} else {
		quotationExtended.VolumeYesterdayUSD = *volumeYesterday
	}
	quotationExtended.Symbol = quotation.Asset.Symbol
	quotationExtended.Name = quotation.Asset.Name
	quotationExtended.Address = quotation.Asset.Address
	quotationExtended.Blockchain = quotation.Asset.Blockchain
	quotationExtended.Price = quotation.Price
	quotationExtended.Time = quotation.Time
	quotationExtended.Source = quotation.Source

	// Get Exchange stats
	exchangemap, _, err := env.DataStore.GetActiveExchangesAndPairs(asset.Address, asset.Blockchain, int64(0), starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}
	var eix []localExchangeInfo
	for exchange, pairs := range exchangemap {
		var ei localExchangeInfo
		ei.Name = exchange
		ei.NumPairs = len(pairs)
		ei.NumTrades, err = env.DataStore.GetNumTrades(exchange, asset.Address, asset.Blockchain, starttime, endtime)
		if err != nil {
			log.Errorf("get number of trades for %s: %v", exchange, err)
		}
		vol, err := env.DataStore.GetVolumeInflux(asset, exchange, starttime, endtime)
		if err != nil {
			log.Errorf("get 24h volume for %s: %v", exchange, err)
		} else {
			ei.Volume24h = *vol
		}
		eix = append(eix, ei)
	}

	sort.Slice(eix, func(i, j int) bool {
		return eix[i].Volume24h > eix[j].Volume24h
	})
	quotationExtended.ExchangeInfo = eix

	c.JSON(http.StatusOK, quotationExtended)
}

// GetPairsInFeed returns quotation of asset with highest market cap among
// all assets with symbol ticker @symbol. Additionally information on exchanges and volumes.
func (env *Env) GetPairsInFeed(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	type localPairInfo struct {
		ForeignName string
		Exchange    string
		NumTrades   int64
		Quotetoken  dia.Asset
		Basetoken   dia.Asset
	}

	type localAssetInfoReturn struct {
		Symbol             string
		Name               string
		Address            string
		Blockchain         string
		Price              float64
		PriceYesterday     float64
		VolumeYesterdayUSD float64
		Time               time.Time
		Source             string
		PairInfo           []localPairInfo
	}
	var quotationExtended localAssetInfoReturn

	blockchain := c.Param("blockchain")
	address := normalizeAddress(c.Param("address"), blockchain)
	numTradesThreshold, err := strconv.ParseInt(c.Param("numTradesThreshold"), 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*60)*time.Minute)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	asset := env.getAssetFromCache(ASSET_CACHE, blockchain, address)

	quotation, err := env.DataStore.GetAssetQuotation(asset, dia.CRYPTO_ZERO_UNIX_TIME, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	quotationYesterday, err := env.DataStore.GetAssetQuotation(asset, dia.CRYPTO_ZERO_UNIX_TIME, starttime)
	if err != nil {
		log.Warn("get quotation yesterday: ", err)
	} else {
		quotationExtended.PriceYesterday = quotationYesterday.Price
	}
	volumeYesterday, err := env.DataStore.Get24HoursAssetVolume(asset)
	if err != nil {
		log.Warn("get volume yesterday: ", err)
	} else {
		quotationExtended.VolumeYesterdayUSD = *volumeYesterday
	}
	quotationExtended.Symbol = quotation.Asset.Symbol
	quotationExtended.Name = quotation.Asset.Name
	quotationExtended.Address = quotation.Asset.Address
	quotationExtended.Blockchain = quotation.Asset.Blockchain
	quotationExtended.Price = quotation.Price
	quotationExtended.Time = quotation.Time
	quotationExtended.Source = quotation.Source

	// Get Exchange stats
	exchangemap, pairCountMap, err := env.DataStore.GetActiveExchangesAndPairs(asset.Address, asset.Blockchain, numTradesThreshold, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	var eix []localPairInfo
	for exchange, pairs := range exchangemap {
		var ei localPairInfo
		ei.Exchange = exchange

		for _, pair := range pairs {
			ei.NumTrades = pairCountMap[pair.PairExchangeIdentifier(exchange)]
			ei.Quotetoken = asset
			ei.Basetoken = env.getAssetFromCache(ASSET_CACHE, pair.BaseToken.Blockchain, pair.BaseToken.Address)
			ei.ForeignName = ei.Quotetoken.Symbol + "-" + ei.Basetoken.Symbol
			eix = append(eix, ei)
		}

	}

	sort.Slice(eix, func(i, j int) bool {
		return eix[i].NumTrades > eix[j].NumTrades
	})
	quotationExtended.PairInfo = eix

	c.JSON(http.StatusOK, quotationExtended)
}

func (env *Env) GetAvailableAssets(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	assetClass := c.Param("assetClass")

	if assetClass == "CryptoToken" {
		assets, err := env.RelDB.GetAllExchangeAssets(true)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, assets)
	} else {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("unknown asset class"))
		return
	}
}

func validateInputParams(c *gin.Context) bool {

	// Validate input parameters.
	for _, input := range c.Params {
		if containsSpecialChars(input.Value) {
			restApi.SendError(c, http.StatusInternalServerError, errors.New("invalid input params"))
			return false
		}
	}

	// Validate query parameters.
	for _, value := range c.Request.URL.Query() {
		for _, input := range value {
			if containsSpecialChars(input) {
				restApi.SendError(c, http.StatusInternalServerError, errors.New("invalid input params"))
				return false
			}
		}
	}

	return true
}

func containsSpecialChars(s string) bool {
	return strings.ContainsAny(s, "!@#$%^&*()'\"|{}[];><?/`~,")
}

// Returns the EIP55 compliant address in case @blockchain has an Ethereum ChainID.
func makeAddressEIP55Compliant(address string, blockchain string) string {
	if strings.Contains(BLOCKCHAINS[blockchain].ChainID, "Ethereum") {
		return common.HexToAddress(address).Hex()
	}
	return address
}

// Normalize address depending on the blockchain.
func normalizeAddress(address string, blockchain string) string {
	if strings.Contains(BLOCKCHAINS[blockchain].ChainID, "Ethereum") {
		return makeAddressEIP55Compliant(address, blockchain)
	}
	if BLOCKCHAINS[blockchain].Name == dia.OSMOSIS {
		if strings.Contains(address, "ibc-") && len(strings.Split(address, "-")[1]) > 1 {
			return "ibc/" + strings.Split(address, "-")[1]
		}
	}
	return address
}

// getDecimalsFromCache returns the decimals of @asset, either from the map @localCache or from
// Postgres, in which latter case it also adds the decimals to the local cache.
// Remember that maps are always passed by reference.
func (env *Env) getDecimalsFromCache(localCache map[dia.Asset]uint8, asset dia.Asset) uint8 {
	if decimals, ok := localCache[asset]; ok {
		return decimals
	}
	fullAsset, err := env.RelDB.GetAsset(asset.Address, asset.Blockchain)
	if err != nil {
		log.Warnf("could not find asset with address %s on blockchain %s in postgres: ", asset.Address, asset.Blockchain)
	}
	localCache[asset] = fullAsset.Decimals
	return fullAsset.Decimals
}

// getAssetFromCache returns the full asset given by blockchain and address, either from the map @localCache
// or from Postgres, in which latter case it also adds the asset to the local cache.
// Remember that maps are always passed by reference.
func (env *Env) getAssetFromCache(localCache map[string]dia.Asset, blockchain string, address string) dia.Asset {
	if asset, ok := localCache[assetIdentifier(blockchain, address)]; ok {
		return asset
	}
	fullAsset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		log.Warnf("could not find asset with address %s on blockchain %s in postgres: ", address, blockchain)
	}
	localCache[assetIdentifier(blockchain, address)] = fullAsset
	return fullAsset
}

func assetIdentifier(blockchain string, address string) string {
	return blockchain + "-" + address
}

func (env *Env) SearchAssetList(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	querystring := c.Param("query")
	var (
		assets = []dia.AssetList{}
		err    error
	)

	assets, err = env.RelDB.SearchAssetList(querystring)
	if err != nil {
		// restApi.SendError(c, http.StatusInternalServerError, errors.New("eror getting asset"))
		log.Errorln("error getting SearchAssetList", err)
	}

	c.JSON(http.StatusOK, assets)
}

func (env *Env) GetAssetList(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	listname := c.Param("listname")

	assets, err := env.RelDB.GetAssetList(listname)
	if err != nil {
		// restApi.SendError(c, http.StatusInternalServerError, errors.New("eror getting asset"))
		log.Errorln("error getting GetAssetList", err)
	}
	if len(assets) <= 0 {
		restApi.SendError(c, http.StatusNotFound, errors.New("asset missing"))
		return
	}

	c.JSON(http.StatusOK, assets)
}

func (env *Env) GetAssetListBySymbol(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	listname := c.Param("listname")

	querystring := c.Param("symbol")
	querystring = strings.ToUpper(querystring)
	var (
		assets = []dia.AssetList{}
		err    error
	)

	assets, err = env.RelDB.GetAssetListBySymbol(querystring, listname)
	if err != nil {
		// restApi.SendError(c, http.StatusInternalServerError, errors.New("eror getting asset"))
		log.Errorln("error getting SearchAssetList", err)
	}
	if len(assets) <= 0 {
		restApi.SendError(c, http.StatusNotFound, errors.New("asset missing"))
		return
	}

	selectedAsset := assets[0]

	splitted := strings.Split(selectedAsset.AssetName, "-")

	price, _, time, source, err := gqlclient.GetGraphqlAssetQuotationFromDia(splitted[0], splitted[1], 60, selectedAsset)
	if err != nil {
		// restApi.SendError(c, http.StatusInternalServerError, errors.New("eror getting asset"))
		log.Errorln("error getting GetGraphqlAssetQuotationFromDia", err)
	}

	asset := dia.Asset{Symbol: selectedAsset.Symbol, Name: selectedAsset.CustomName, Blockchain: splitted[0], Address: splitted[1]}
	q := models.AssetQuotationFull{Symbol: asset.Symbol, Name: asset.Name, Address: asset.Address, Price: price, Blockchain: asset.Blockchain}

	volumeYesterday, err := env.DataStore.Get24HoursAssetVolume(asset)
	if err != nil {
		log.Errorln("error getting Get24HoursAssetVolume", err)
	}
	q.VolumeYesterdayUSD = *volumeYesterday
	q.Time = time
	q.Source = strings.Join(source, ",")
	c.JSON(http.StatusOK, q)
}
