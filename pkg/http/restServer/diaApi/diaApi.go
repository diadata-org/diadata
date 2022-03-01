package diaApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/indexCalculationService"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var DECIMALS_CACHE = make(map[dia.Asset]uint8)

type Env struct {
	DataStore models.Datastore
	RelDB     models.RelDB
}

// PostSupply deprecated? TO DO
func (env *Env) PostSupply(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("ReadAll"))
	} else {
		var t dia.Supply
		err = json.Unmarshal(body, &t)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		} else {
			if t.Asset.Symbol == "" || t.CirculatingSupply == 0.0 {
				log.Errorln("received supply:", t)
				restApi.SendError(c, http.StatusInternalServerError, errors.New("missing symbol or circulating supply value"))
			} else {
				log.Println("received supply:", t)
				source := dia.Diadata
				if t.Source != "" {
					source = t.Source
				}
				s := &dia.Supply{
					Time:              t.Time,
					Asset:             t.Asset,
					Source:            source,
					CirculatingSupply: t.CirculatingSupply}

				err := env.DataStore.SetSupply(s)

				if err == nil {
					c.JSON(http.StatusOK, s)
				} else {
					restApi.SendError(c, http.StatusInternalServerError, err)
				}
			}
		}
	}
}

// SetQuotation sets a quotation to redis cache. Input must be of the format:
// '["blockchain","address","value"]'
func (env *Env) SetQuotation(c *gin.Context) {

	var quotation models.AssetQuotation
	var input []string
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("ReadAll"))
		return
	}
	err = json.Unmarshal(body, &input)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("unmarshal body"))
		return
	}
	if len(input) != 3 {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("wrong number of inputs"))
		return
	}

	quotation.Asset.Blockchain = input[0]
	quotation.Asset.Address = input[1]
	price, err := strconv.ParseFloat(input[2], 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	quotation.Price = price
	quotation.Source = "diadata.org"
	quotation.Time = time.Now()

	_, err = env.DataStore.SetAssetQuotationCache(&quotation, true)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
}

// GetAssetQuotation returns quotation of asset with highest market cap among
// all assets with symbol ticker @symbol.
func (env *Env) GetAssetQuotation(c *gin.Context) {
	blockchain := c.Param("blockchain")
	address := c.Param("address")
	var err error
	var asset dia.Asset
	var quotationExtended models.AssetQuotationFull
	timestamp := time.Now()

	// An asset is uniquely defined by blockchain and address.
	asset, err = env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	// Get quotation for asset.
	quotation, err := env.DataStore.GetAssetQuotation(asset, timestamp)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	quotationYesterday, err := env.DataStore.GetAssetQuotation(asset, timestamp.AddDate(0, 0, -1))
	if err != nil {
		log.Warn("get quotation yesterday: ", err)
	} else {
		quotationExtended.PriceYesterday = quotationYesterday.Price
	}
	volumeYesterday, err := env.RelDB.GetAssetVolume24H(asset)
	if err != nil {
		log.Warn("get volume yesterday: ", err)
	} else {
		quotationExtended.VolumeYesterdayUSD = volumeYesterday
	}

	// Appropriate formatting.
	quotationExtended.Symbol = quotation.Asset.Symbol
	quotationExtended.Name = quotation.Asset.Name
	quotationExtended.Address = quotation.Asset.Address
	quotationExtended.Blockchain = quotation.Asset.Blockchain
	quotationExtended.Price = quotation.Price
	quotationExtended.Time = quotation.Time
	quotationExtended.Source = quotation.Source

	c.JSON(http.StatusOK, quotationExtended)

}

// GetQuotation returns quotation of asset with highest market cap among
// all assets with symbol ticker @symbol.
func (env *Env) GetQuotation(c *gin.Context) {
	symbol := c.Param("symbol")
	timestamp := time.Now()
	var quotationExtended models.AssetQuotationFull
	// Fetch underlying assets for symbol
	assets, err := env.RelDB.GetTopAssetByVolume(symbol)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}
	log.Info("num assets: ", len(assets))
	log.Info("error: ", err)
	if len(assets) == 0 {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	topAsset := assets[0]
	quotation, err := env.DataStore.GetAssetQuotation(topAsset, timestamp)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	quotationYesterday, err := env.DataStore.GetAssetQuotation(topAsset, timestamp.AddDate(0, 0, -1))
	if err != nil {
		log.Warn("get quotation yesterday: ", err)
	} else {
		quotationExtended.PriceYesterday = quotationYesterday.Price
	}
	volumeYesterday, err := env.RelDB.GetAssetVolume24H(topAsset)
	if err != nil {
		log.Warn("get volume yesterday: ", err)
	} else {
		quotationExtended.VolumeYesterdayUSD = volumeYesterday
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

func (env *Env) GetPaxgQuotationOunces(c *gin.Context) {
	q, err := env.DataStore.GetPaxgQuotationOunces()
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

func (env *Env) GetPaxgQuotationGrams(c *gin.Context) {
	q, err := env.DataStore.GetPaxgQuotationGrams()
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

// GetSupply returns latest supply of token with @symbol
func (env *Env) GetSupply(c *gin.Context) {
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
	address := c.Param("address")
	blockchain := c.Param("blockchain")
	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")
	var starttime, endtime time.Time
	if starttimeStr != "" && endtimeStr != "" {
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			log.Error("parse starttime: ", err)
		} else {
			starttime = time.Unix(starttimeInt, 0)
		}
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			log.Error("parse starttime: ", err)
		} else {
			endtime = time.Unix(endtimeInt, 0)
		}
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
	symbol := c.Param("symbol")
	starttimeStr := c.DefaultQuery("starttime", "noRange")
	endtimeStr := c.Query("endtime")

	var starttime, endtime time.Time

	if starttimeStr == "noRange" || endtimeStr == "" {
		starttime = time.Unix(1, 0)
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
		if err == redis.Nil {
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
		if err == redis.Nil {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
	}
}

// GetVolume if no times are set use the last 24h
func (env *Env) GetVolume(c *gin.Context) {
	symbol := c.Param("symbol")
	starttimeStr := c.DefaultQuery("starttime", "noRange")
	endtimeStr := c.Query("endtime")

	var starttime, endtime time.Time

	if starttimeStr == "noRange" {
		starttime = time.Unix(1, 0)
	} else {
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
	}
	if endtimeStr == "" {
		endtime = time.Now()
	} else {
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	}

	// TO DO: Adapt to new asset struct
	preliminaryAsset := dia.Asset{
		Symbol: symbol,
	}
	v, err := env.DataStore.GetVolumeInflux(preliminaryAsset, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, v)
}

// Get24hVolume if no times are set use the last 24h
func (env *Env) Get24hVolume(c *gin.Context) {
	exchange := c.Param("exchange")
	// starttimeStr := c.DefaultQuery("starttime", "noRange")
	// endtimeStr := c.Query("endtime")

	// var starttime, endtime time.Time

	// if starttimeStr == "noRange" {
	// 	starttime := time.Now().AddDate(0, 0, -1)
	// } else {
	// 	starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
	// 	if err != nil {
	// 		restApi.SendError(c, http.StatusInternalServerError, err)
	// 		return
	// 	}
	// 	starttime = time.Unix(starttimeInt, 0)
	// }
	// if endtimeStr == "" {
	// 	endtime = time.Now()
	// } else {
	// 	endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
	// 	if err != nil {
	// 		restApi.SendError(c, http.StatusInternalServerError, err)
	// 		return
	// 	}
	// 	endtime = time.Unix(endtimeInt, 0)
	// }

	v, err := env.DataStore.Sum24HoursExchange(exchange)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, v)
}

// GetPairs returns all pairs
func (env *Env) GetPairs(c *gin.Context) {
	p, err := env.RelDB.GetPairs("")
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, &models.Pairs{Pairs: p})
	}
}

// GetExchanges is the delegate method for fetching all
// available trading places.
func (env *Env) GetExchanges(c *gin.Context) {
	q := env.DataStore.GetExchanges()
	if len(q) == 0 {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetChartPoints godoc
// @Param   scale      query   string     false       "scale 5m 30m 1h 4h 1d 1w"
func (env *Env) GetChartPoints(c *gin.Context) {
	filter := c.Param("filter")
	exchange := c.Param("exchange")
	symbol := c.Param("symbol")
	scale := c.Query("scale")
	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")

	// Set times depending on what is given by the query parameters
	var starttime, endtime time.Time
	if starttimeStr == "" && endtimeStr == "" {
		// Last seven days per default
		starttime = time.Now().AddDate(0, 0, -7)
		endtime = time.Now()
	} else if starttimeStr == "" && endtimeStr != "" {
		// zero time if not given
		starttime = time.Time{}
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	} else if starttimeStr != "" && endtimeStr == "" {
		// endtime now if not given
		endtime = time.Now()
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
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

	p, err := env.DataStore.GetFilterPoints(filter, exchange, symbol, scale, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// GetChartPointsAllExchanges godoc
// @Param   scale      query   string     false       "scale 5m 30m 1h 4h 1d 1w"
func (env *Env) GetChartPointsAllExchanges(c *gin.Context) {
	filter := c.Param("filter")
	symbol := c.Param("symbol")
	scale := c.Query("scale")
	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")

	// Set times depending on what is given by the query parameters
	var starttime, endtime time.Time
	if starttimeStr == "" && endtimeStr == "" {
		// Last seven days per default
		starttime = time.Now().AddDate(0, 0, -7)
		endtime = time.Now()
	} else if starttimeStr == "" && endtimeStr != "" {
		// zero time if not given
		starttime = time.Time{}
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	} else if starttimeStr != "" && endtimeStr == "" {
		// endtime now if not given
		endtime = time.Now()
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
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

	p, err := env.DataStore.GetFilterPoints(filter, "", symbol, scale, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// GetAllSymbols returns all Symbols on @exchange.
// If @exchange is not set, it returns all symbols across all exchanges.
// If @top is set to an integer, only the top @top symbols w.r.t. trading volume are returned. This is
// only active if @exchange is not set.
func (env *Env) GetAllSymbols(c *gin.Context) {
	var s []string
	var numSymbols int64
	var sortedAssets []dia.Asset
	var err error

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
		sortedAssets, err = env.RelDB.GetAssetsWithVOL(int64(0), substring)
		if err != nil {
			log.Error("get assets with volume: ", err)
		}
		var sortedSymbols []string
		for _, asset := range sortedAssets {
			sortedSymbols = append(sortedSymbols, asset.Symbol)
		}
		sortedSymbols = utils.UniqueStrings(sortedSymbols)
		allSymbols := utils.UniqueStrings(append(sortedSymbols, s...))

		c.JSON(http.StatusOK, allSymbols)
		return
	}

	if exchange == "noRange" {
		if numSymbolsString != "" {
			// -- Get top @numSymbols symbols across all exchanges. --
			sortedAssets, err = env.RelDB.GetAssetsWithVOL(numSymbols, "")
			if err != nil {
				log.Error("get assets with volume: ", err)
			}
			for _, asset := range sortedAssets {
				s = append(s, asset.Symbol)
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
			sortedAssets, err = env.RelDB.GetAssetsWithVOL(numSymbols, "")
			if err != nil {
				log.Error("get assets with volume: ", err)
			}
			var sortedSymbols []string
			for _, asset := range sortedAssets {
				sortedSymbols = append(sortedSymbols, asset.Symbol)
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
// INDICES
// -----------------------------------------------------------------------------

func (env *Env) GetCviIndex(c *gin.Context) {
	starttimeStr := c.DefaultQuery("starttime", "noRange")
	endtimeStr := c.Query("endtime")
	symbol := c.Query("symbol")

	var starttime, endtime time.Time
	var q []dia.CviDataPoint
	var err error

	if starttimeStr == "noRange" || endtimeStr == "" {
		starttime = time.Unix(0, 0)
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
	}

	q, err = env.DataStore.GetCVIInflux(starttime, endtime, symbol)

	//for i := range q {
	//	q[i].Value /= 2430.5812295231785
	//}
	if len(q) == 0 {
		c.JSON(http.StatusOK, make([]string, 0))
	}
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, q)
}

// GetCryptoDerivative returns all information on a given derivative of class
// @derivativeType and name @name
func (env *Env) GetCryptoDerivative(c *gin.Context) {
	derivativeType := c.Param("type")
	fmt.Println(derivativeType)
	derivativeName := c.Param("name")
	fmt.Println(derivativeName)
	// TO DO
	// 2-step:
	// 1. specify class of derivative
	// 2. get derivative information from derivative given by @derivativeName in class given by item 1.

	// q, err := env.DataStore.GetCryptoIndex(indexType)
	// if err != nil {
	// 	if errors.Is(err, redis.Nil) {
	// 		restApi.SendError(c, http.StatusNotFound, err)
	// 	} else {
	// 		restApi.SendError(c, http.StatusInternalServerError, err)
	// 	}
	// }

	// Depending on return format of GetCryptoIndex, either get additional information
	// on the constituents or directly fill the new type "CryptoIndex"

}

// -----------------------------------------------------------------------------
// DeFi LENDING RATES
// -----------------------------------------------------------------------------

// GetLendingProtocols returns all symbols available in our (redis) database.
// Optional query parameter exchange returns only symbols available on this exchange.
func (env *Env) GetLendingProtocols(c *gin.Context) {
	q, err := env.DataStore.GetDefiProtocols()
	fmt.Println("protocols: ", q)
	if len(q) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetDefiRate is the delegate method to fetch the value(s) of
// the defi lending rate of @asset at the exchange with @protocol.
// Last value is retrieved. Otional query parameters allow to obtain data in a time range.
func (env *Env) GetDefiRate(c *gin.Context) {
	protocol := c.Param("protocol")
	asset := c.Param("asset")
	date := c.Param("time")
	// Add optional query parameters for requesting a range of values
	dateInit := c.DefaultQuery("dateInit", "noRange")
	dateFinal := c.Query("dateFinal")

	if dateInit == "noRange" {
		// Return most recent data point
		endtime := time.Time{}
		var err error
		if date == "" {
			endtime = time.Now()
		} else {
			// Convert unix time int/string to time
			endtime, err = utils.StrToUnixtime(date)
			if err != nil {
				restApi.SendError(c, http.StatusNotFound, err)
			}
		}
		starttime := endtime.AddDate(0, 0, -1)

		q, err := env.DataStore.GetDefiRateInflux(starttime, endtime, asset, protocol)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q[len(q)-1])
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
		q, err := env.DataStore.GetDefiRateInflux(starttime, endtime, asset, protocol)
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

// GetDefiState is the delegate method to fetch the value(s) of
// the defi lending rate of @asset at the exchange with @protocol.
// Last value is retrieved. Otional query parameters allow to obtain data in a time range.
func (env *Env) GetDefiState(c *gin.Context) {
	protocol := c.Param("protocol")
	date := c.Param("time")
	// Add optional query parameters for requesting a range of values
	dateInit := c.DefaultQuery("dateInit", "noRange")
	dateFinal := c.Query("dateFinal")

	if dateInit == "noRange" {
		// Return most recent data point
		endtime := time.Time{}
		var err error
		if date == "" {
			endtime = time.Now()
		} else {
			// Convert unix time int/string to time
			endtime, err = utils.StrToUnixtime(date)
			if err != nil {
				restApi.SendError(c, http.StatusNotFound, err)
			}
		}
		starttime := endtime.AddDate(0, 0, -1)

		q, err := env.DataStore.GetDefiStateInflux(starttime, endtime, protocol)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q[len(q)-1])
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
		q, err := env.DataStore.GetDefiStateInflux(starttime, endtime, protocol)
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

// -----------------------------------------------------------------------------
// FARMING POOLS
// -----------------------------------------------------------------------------

// GetFarmingPools is the delegate method to fetch the value(s) of
// the farming pool information of @protocol.
// Last value is retrieved. Otional query parameters allow to obtain data in a time range.
func (env *Env) GetFarmingPools(c *gin.Context) {
	q, err := env.DataStore.GetFarmingPools()
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

// GetFarmingPoolData is the delegate method to fetch the value(s) of
// the farming pool information of @protocol.
// Last value is retrieved. Otional query parameters allow to obtain data in a time range.
func (env *Env) GetFarmingPoolData(c *gin.Context) {
	protocol := c.Param("protocol")
	poolID := c.Param("poolID")
	date := c.Param("time")
	// Add optional query parameters for requesting a range of values
	dateInit := c.DefaultQuery("dateInit", "noRange")
	dateFinal := c.Query("dateFinal")

	if dateInit == "noRange" {
		// Return most recent data point
		endtime := time.Time{}
		var err error
		if date == "" {
			endtime = time.Now()
		} else {
			// Convert unix time int/string to time
			endtime, err = utils.StrToUnixtime(date)
			if err != nil {
				restApi.SendError(c, http.StatusNotFound, err)
			}
		}
		starttime := endtime.AddDate(0, 0, -1)

		q, err := env.DataStore.GetFarmingPoolData(starttime, endtime, protocol, poolID)
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
		q, err := env.DataStore.GetFarmingPoolData(starttime, endtime, protocol, poolID)
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

// -----------------------------------------------------------------------------
// INTEREST RATES
// -----------------------------------------------------------------------------

// GetInterestRate is the delegate method to fetch the value of
// the interest rate with symbol @symbol at the date @time.
// Optional query parameters allow to obtain data in a time range.
func (env *Env) GetInterestRate(c *gin.Context) {
	symbol := c.Param("symbol")
	date := c.Param("time")
	// Add optional query parameters for requesting a range of values
	dateInit := c.DefaultQuery("dateInit", "noRange")
	dateFinal := c.Query("dateFinal")

	if dateInit == "noRange" {
		q, err := env.DataStore.GetInterestRate(symbol, date)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q)
		}
	} else {
		q, err := env.DataStore.GetInterestRateRange(symbol, dateInit, dateFinal)
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

// GetCompoundedRate is the delegate method to fetch compounded rate values for interest rates
func (env *Env) GetCompoundedRate(c *gin.Context) {

	// Import and cast input from API call
	symbol := c.Param("symbol")
	dpy := c.Param("dpy")
	daysPerYear, err := strconv.Atoi(dpy)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	}
	datestring := c.Param("time")

	// Add optional query parameters for requesting a range of values
	dateInitstring := c.DefaultQuery("dateInit", "noRange")
	dateFinalstring := c.Query("dateFinal")

	// Retrieve rounding convention for @symbol
	rounding := 0

	if dateInitstring == "noRange" {

		var date time.Time
		if datestring == "" {
			// If date is omitted in API call, take most recent date
			date = time.Now()
		} else {
			date, err = time.Parse("2006-01-02", datestring)
		}
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}

		q, err := env.DataStore.GetCompoundedIndex(symbol, date, daysPerYear, rounding)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q)
		}

	} else {

		dateInit, err := time.Parse("2006-01-02", dateInitstring)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
		dateFinal, err := time.Parse("2006-01-02", dateFinalstring)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}

		q, err := env.DataStore.GetCompoundedIndexRange(symbol, dateInit, dateFinal, daysPerYear, rounding)
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

// GetCompoundedAvg is the delegate method to fetch averaged compounded rate values for interest rates
func (env *Env) GetCompoundedAvg(c *gin.Context) {

	tInit := time.Now()

	// Import and cast input from API call
	symbol := c.Param("symbol")
	datestring := c.Param("time")
	date, _ := time.Parse("2006-01-02", datestring)
	days := c.Param("days")
	calDays, err := strconv.Atoi(days)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	}
	dpy := c.Param("dpy")
	daysPerYear, err := strconv.Atoi(dpy)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	}

	// Add optional query parameters for requesting a range of values
	dateInitstring := c.DefaultQuery("dateInit", "noRange")
	dateFinalstring := c.Query("dateFinal")

	rounding := 0

	if dateInitstring == "noRange" {

		// Compute compunded rate and return if no error
		q, err := env.DataStore.GetCompoundedAvg(symbol, date, calDays, daysPerYear, rounding)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q)
		}

	} else {

		dateInit, err := time.Parse("2006-01-02", dateInitstring)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
		dateFinal, err := time.Parse("2006-01-02", dateFinalstring)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}

		q, err := env.DataStore.GetCompoundedAvgRange(symbol, dateInit, dateFinal, calDays, daysPerYear, rounding)
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

	tFinal := time.Now()
	fmt.Println("time elapsed in API call: ", tFinal.Sub(tInit))
}

// GetCompoundedAvgDIA is the delegate method to fetch averaged compounded rate values for interest rates
func (env *Env) GetCompoundedAvgDIA(c *gin.Context) {

	tInit := time.Now()

	// Import and cast input from API call
	symbol := c.Param("symbol")
	datestring := c.Param("time")
	date, _ := time.Parse("2006-01-02", datestring)
	days := c.Param("days")
	calDays, err := strconv.Atoi(days)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	}
	dpy := c.Param("dpy")
	daysPerYear, err := strconv.Atoi(dpy)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	}

	// Add optional query parameters for requesting a range of values
	dateInitstring := c.DefaultQuery("dateInit", "noRange")
	dateFinalstring := c.Query("dateFinal")

	rounding := 0

	if dateInitstring == "noRange" {

		// In this method, there is a rate for every calendar day. Hence, the compounded rate
		// for a particular day can be retrieved by the range method easily.
		dateFinal := date.AddDate(0, 0, 1)
		q, err := env.DataStore.GetCompoundedAvgDIARange(symbol, date, dateFinal, calDays, daysPerYear, rounding)

		if err != nil {
			if errors.Is(err, redis.Nil) {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q)
		}

	} else {

		dateInit, err := time.Parse("2006-01-02", dateInitstring)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
		dateFinal, err := time.Parse("2006-01-02", dateFinalstring)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}

		q, err := env.DataStore.GetCompoundedAvgDIARange(symbol, dateInit, dateFinal, calDays, daysPerYear, rounding)
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

	tFinal := time.Now()
	fmt.Println("time elapsed in API call: ", tFinal.Sub(tInit))
}

// GetRates is the delegate method for fetching all rate types
// present in the (redis) database.
func (env *Env) GetRates(c *gin.Context) {
	q, err := env.DataStore.GetRatesMeta()
	if len(q) == 0 {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, q)
}

// -----------------------------------------------------------------------------
// FIAT CURRENCIES
// -----------------------------------------------------------------------------

// GetFiatQuotations returns several quotations vs USD as published by the ECB
func (env *Env) GetFiatQuotations(c *gin.Context) {
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

// -----------------------------------------------------------------------------
// STOCKS
// -----------------------------------------------------------------------------

func (env *Env) GetStockSymbols(c *gin.Context) {
	type sourcedStock struct {
		Stock  models.Stock
		Source string
	}
	var srcStocks []sourcedStock
	stocks, err := env.DataStore.GetStockSymbols()
	log.Info("stocks: ", stocks)

	if err != nil {
		if err == redis.Nil {
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
	source := c.Param("source")
	symbol := c.Param("symbol")
	date := c.Param("time")
	// Add optional query parameters for requesting a range of values
	dateInit := c.DefaultQuery("dateInit", "noRange")
	dateFinal := c.Query("dateFinal")

	if dateInit == "noRange" {
		// Return most recent data point
		endtime := time.Time{}
		var err error
		if date == "" {
			endtime = time.Now()
		} else {
			// Convert unix time int/string to time
			endtime, err = utils.StrToUnixtime(date)
			if err != nil {
				restApi.SendError(c, http.StatusNotFound, err)
			}
		}
		starttime := endtime.AddDate(0, 0, -1)

		q, err := env.DataStore.GetStockQuotation(source, symbol, starttime, endtime)
		if err != nil {
			if err == redis.Nil {
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
			if err == redis.Nil {
				restApi.SendError(c, http.StatusNotFound, err)
			} else {
				restApi.SendError(c, http.StatusInternalServerError, err)
			}
		} else {
			c.JSON(http.StatusOK, q)
		}
	}
}

// -----------------------------------------------------------------------------
// FOREIGN QUOTATIONS
// -----------------------------------------------------------------------------

// GetForeignQuotation returns several quotations vs USD as published by the ECB
func (env *Env) GetForeignQuotation(c *gin.Context) {
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
	q, err := env.DataStore.GetForeignQuotationInflux(symbol, source, timestamp)
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

// GetForeignSymbols returns all symbols available for quotation from @source, along with their ITIN
func (env *Env) GetForeignSymbols(c *gin.Context) {
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
// CRYPTO INDEX
// -----------------------------------------------------------------------------

func (env *Env) GetCryptoIndex(c *gin.Context) {
	symbol := c.Param("symbol")
	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")
	maxResultsString := c.Query("maxResults")
	var maxResults int
	var err error
	if maxResultsString != "" {
		maxResults, err = strconv.Atoi(maxResultsString)
		if err != nil {
			log.Error("parse maxResults: ", err)
		}
	} else {
		maxResults = 1
	}

	// Set times depending on what is given by the query parameters
	var starttime, endtime time.Time
	if starttimeStr == "" && endtimeStr == "" {
		// Last seven days per default
		starttime = time.Now().AddDate(0, 0, -7)
		endtime = time.Now()
	} else if starttimeStr == "" && endtimeStr != "" {
		// zero time if not given
		starttime = time.Time{}
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	} else if starttimeStr != "" && endtimeStr == "" {
		// endtime now if not given
		endtime = time.Now()
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
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

	q, err := env.DataStore.GetCryptoIndex(starttime, endtime, symbol, maxResults)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// Fetch decimals from local cache implementation.
	for i := range q {
		q[i].Asset.Decimals = env.getDecimalsFromCache(DECIMALS_CACHE, q[i].Asset)
		for j := range q[i].Constituents {
			q[i].Constituents[j].Asset.Decimals = env.getDecimalsFromCache(DECIMALS_CACHE, q[i].Constituents[j].Asset)
		}
	}

	c.JSON(http.StatusOK, q)
}

func (env *Env) GetCryptoIndexValues(c *gin.Context) {
	symbol := c.Param("symbol")
	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")
	maxResults := 0

	// Set times depending on what is given by the query parameters
	var starttime, endtime time.Time
	if starttimeStr == "" && endtimeStr == "" {
		starttime = time.Now().Add(time.Duration(-4 * time.Hour))
		endtime = time.Now()
		maxResults = 1
	} else if starttimeStr == "" && endtimeStr != "" {
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
		starttime = endtime.AddDate(0, 0, -1)
	} else if starttimeStr != "" && endtimeStr == "" {
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
		endtime = starttime.AddDate(0, 0, 1)
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

	q, err := env.DataStore.GetCryptoIndexValues(starttime, endtime, symbol, maxResults)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// local return type
	type indexVals struct {
		Symbol          string
		Address         string
		Blockchain      string
		Value           float64
		CalculationTime time.Time
	}
	var returnIndices []indexVals
	for _, index := range q {
		tmp := indexVals{
			Symbol:          index.Asset.Symbol,
			Address:         index.Asset.Address,
			Blockchain:      index.Asset.Blockchain,
			Value:           index.Value,
			CalculationTime: index.CalculationTime,
		}
		returnIndices = append(returnIndices, tmp)
	}

	c.JSON(http.StatusOK, returnIndices)
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

// GetBenchmarkedIndexValue Get benchmarked Index values
func (env *Env) GetBenchmarkedIndexValue(c *gin.Context) {
	symbol := c.Param("symbol")
	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")

	// Set times depending on what is given by the query parameters
	var starttime, endtime time.Time
	if starttimeStr == "" && endtimeStr == "" {
		// Last seven days per default
		starttime = time.Now().AddDate(0, 0, -7)
		endtime = time.Now()
	} else if starttimeStr == "" && endtimeStr != "" {
		// zero time if not given
		starttime = time.Time{}
		endtimeInt, err := strconv.ParseInt(endtimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	} else if starttimeStr != "" && endtimeStr == "" {
		// endtime now if not given
		endtime = time.Now()
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
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

	q, err := env.DataStore.GetBenchmarkedIndexValuesInflux(symbol, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, q)
}

// GetLastTrades Get last 1000 trades of an asset
func (env *Env) GetLastTrades(c *gin.Context) {
	symbol := c.Param("symbol")

	// First get asset with @symbol with largest market cap.
	topAsset, err := env.DataStore.GetTopAssetByVolume(symbol, &env.RelDB)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
	}

	q, err := env.DataStore.GetLastTrades(topAsset, "", 1000, true)
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

// GetLastTrades returns last N trades of an asset. Defaults to N=1000.
func (env *Env) GetLastTradesAsset(c *gin.Context) {
	blockchain := c.Param("blockchain")
	address := c.Param("address")
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
	// Make the address EIP55 compliant.
	address = common.HexToAddress(address).Hex()

	asset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	q, err := env.DataStore.GetLastTrades(asset, exchange, int(numTrades), true)
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

// PostIndexRebalance Post data must be of the type [][]string.
// Each entry of the 2-d slice corresponds to an asset and has the format [Blockchain, Address].
func (env *Env) PostIndexRebalance(c *gin.Context) {
	indexSymbol := c.Param("symbol")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("ReadAll"))
		return
	}
	var assetlist [][]string
	var assets []dia.Asset
	err = json.Unmarshal(body, &assetlist)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	// Cast assetlist to []dia.Asset
	for _, asset := range assetlist {
		assets = append(assets, dia.Asset{
			Blockchain: asset[0],
			Address:    asset[1],
		})
	}

	// Get constituents information
	constituents, err := indexCalculationService.GetIndexBasket(assets)
	if err != nil {
		log.Error(err)
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// Calculate relative weights
	err = indexCalculationService.CalculateWeights(indexSymbol, &constituents)
	if err != nil {
		log.Error(err)
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	// Get old index
	currIndex, err := env.DataStore.GetCryptoIndex(time.Now().Add(-24*time.Hour), time.Now(), indexSymbol, 1)
	if err != nil {
		log.Error(err)
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	var newIndexValue float64
	var newIndexRawValue float64
	newDivisor := 1.0
	if indexSymbol == "SCIFI" {
		// Determine new divisor
		currIndexRawValue := currIndex[0].Value * currIndex[0].Divisor
		newIndexRawValue = models.GetIndexValue(indexSymbol, constituents)
		newDivisor = (newIndexRawValue * currIndex[0].Divisor) / currIndexRawValue
		newIndexValue = newIndexRawValue / newDivisor
	} else {
		newIndexValue = currIndex[0].Value
		newIndexRawValue = currIndex[0].Value
	}

	// Calculate Base Amount for each constituent
	for i, constituent := range constituents {
		constituents[i].NumBaseTokens = ((constituent.Weight * newIndexValue) / constituent.Price) * 1e16 //((Weight * IndexPrice) / TokenPrice) * 1e18  (divided by 100 because index level is 100 = 1 usd)
	}

	var newIndex models.CryptoIndex
	newIndex.Asset.Name = indexSymbol
	newIndex.Constituents = constituents
	newIndex.Value = newIndexRawValue
	newIndex.Price = currIndex[0].Price
	newIndex.Divisor = newDivisor

	err = env.DataStore.SetCryptoIndex(&newIndex)
	if err != nil {
		log.Error()
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, constituents)
}

// GetMissingExchangeSymbol returns all unverified symbol
func (env *Env) GetMissingExchangeSymbol(c *gin.Context) {
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
	symbol := c.Param("symbol")

	symbols, err := env.RelDB.GetAssets(symbol)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, symbols)
	}
}

func (env *Env) GetAssetExchanges(c *gin.Context) {
	symbol := c.Param("symbol")

	symbols, err := env.RelDB.GetAssetExchange(symbol)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, symbols)
	}
}

func (env *Env) GetAllBlockchains(c *gin.Context) {
	blockchains, err := env.RelDB.GetAllBlockchains()
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, blockchains)
	}
}

// -----------------------------------------------------------------------------
// NFT
// -----------------------------------------------------------------------------

// GetNFTCategories returns all available NFT categories.
func (env *Env) GetNFTCategories(c *gin.Context) {
	q, err := env.RelDB.GetNFTCategories()
	if len(q) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetAllNFTClasses returns all NFT classes.
func (env *Env) GetAllNFTClasses(c *gin.Context) {
	blockchain := c.Param("blockchain")
	q, err := env.RelDB.GetAllNFTClasses(blockchain)
	if len(q) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetNFTClasses returns all NFT classes.
func (env *Env) GetNFTClasses(c *gin.Context) {
	limitString := c.Param("limit")
	offsetString := c.Param("offset")
	limit, err := strconv.ParseUint(limitString, 10, 32)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	offset, err := strconv.ParseUint(offsetString, 10, 32)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	q, err := env.RelDB.GetNFTClasses(limit, offset)
	if len(q) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetNFT returns an NFT.
func (env *Env) GetNFT(c *gin.Context) {
	blockchain := c.Param("blockchain")
	address := c.Param("address")
	id := c.Param("id")
	q, err := env.RelDB.GetNFT(address, blockchain, id)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetNFTTrades returns all trades of the unique NFT with given parameters.
func (env *Env) GetNFTTrades(c *gin.Context) {
	blockchain := c.Param("blockchain")
	// Sanitize address
	address := common.HexToAddress(c.Param("address")).Hex()
	id := c.Param("id")

	nft, err := env.RelDB.GetNFT(address, blockchain, id)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	q, err := env.RelDB.GetNFTTrades(nft)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetNFTPrice30Days returns the average price of the whole nft class over the last 30 days.
func (env *Env) GetNFTPrice30Days(c *gin.Context) {
	blockchain := c.Param("blockchain")
	address := common.HexToAddress(c.Param("address")).Hex()
	nftClass, err := env.RelDB.GetNFTClass(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	avgPrice, err := env.RelDB.GetNFTPrice30Days(nftClass)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, avgPrice)
}
