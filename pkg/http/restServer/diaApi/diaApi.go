package diaApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	DECIMALS_CACHE = make(map[dia.Asset]uint8)
	ASSET_CACHE    = make(map[string]dia.Asset)
	BLOCKCHAINS    = make(map[string]dia.BlockChain)
)

type Env struct {
	DataStore models.Datastore
	RelDB     models.RelDB
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
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	var (
		err               error
		asset             dia.Asset
		quotationExtended models.AssetQuotationFull
	)
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

func (env *Env) GetAssetMap(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

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

	log.Info("num assets: ", len(assets))
	if len(assets) == 0 {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	log.Info("num assets: ", len(assets))

	for _, topAsset := range assets {
		var quotationExtended models.AssetQuotationFull

		quotation, err := env.DataStore.GetAssetQuotation(topAsset, timestamp)
		if err != nil {
			log.Warn("get quotation: ", err)
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
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

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

	if ok, err := validTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
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
	if ok, err := validTimeRange(starttime, endtime, time.Duration(31*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
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
	if !validateInputParams(c) {
		return
	}

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
	v, err := env.DataStore.GetVolumeInflux(preliminaryAsset, "", starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, *v)
}

// Get24hVolume if no times are set use the last 24h
func (env *Env) Get24hVolume(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

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

	v, err := env.DataStore.Get24HoursExchangeVolume(exchange)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, *v)
}

// GetExchanges is the delegate method for fetching all exchanges available in Postgres.
func (env *Env) GetExchanges(c *gin.Context) {
	type exchangeReturn struct {
		Name       string
		Volume24h  float64
		Trades     int64
		Pairs      int
		Type       string
		Blockchain string
	}
	var exchangereturns []exchangeReturn
	exchanges, err := env.RelDB.GetAllExchanges()
	if len(exchanges) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	for _, exchange := range exchanges {

		vol, err := env.DataStore.Get24HoursExchangeVolume(exchange.Name)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		numTrades, err := env.DataStore.GetNumTradesExchange24H(exchange.Name)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		numPairs, err := env.RelDB.GetNumPairs(exchange)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}

		exchangereturn := exchangeReturn{
			Name:       exchange.Name,
			Volume24h:  *vol,
			Trades:     numTrades,
			Pairs:      numPairs,
			Blockchain: exchange.BlockChain.Name,
		}
		exchangereturn.Type = models.GetExchangeType(exchange)
		exchangereturns = append(exchangereturns, exchangereturn)

	}

	sort.Slice(exchangereturns, func(i, j int) bool {
		return exchangereturns[i].Volume24h > exchangereturns[j].Volume24h
	})

	c.JSON(http.StatusOK, exchangereturns)
}

// GetNFTExchanges is the delegate method for fetching all exchanges available in Postgres.
func (env *Env) GetNFTExchanges(c *gin.Context) {
	type exchangeReturn struct {
		Name       string
		Volume24h  float64
		Trades24h  int64
		Blockchain string
		// NumCollections int64
	}
	var exchangereturns []exchangeReturn
	exchanges, err := env.RelDB.GetAllNFTExchanges()

	log.Infoln("exchanges", exchanges)
	if len(exchanges) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	for _, exchange := range exchanges {

		vol, err := env.RelDB.Get24HoursNFTExchangeVolume(exchange)
		if err != nil {
			log.Errorln("err on Get24HoursNFTExchangeVolume", err)
		}
		numTrades, err := env.RelDB.Get24HoursNFTExchangeTrades(exchange)
		if err != nil {
			log.Errorln("err on Get24HoursNFTExchangeTrades", err)

		}
		// numCollections, err := env.RelDB.GetCollectionCountByExchange(exchange.Name)
		// if err != nil {
		// 	log.Errorln("err on GetCollectionCountByExchange", err)
		// }

		exchangereturn := exchangeReturn{
			Name:       exchange.Name,
			Volume24h:  vol,
			Trades24h:  numTrades,
			Blockchain: exchange.BlockChain.Name,
			// NumCollections: numCollections,
		}
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
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	exchange := c.Query("exchange")
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
	if ok, err := validTimeRange(starttime, endtime, time.Duration(14*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	p, err := env.DataStore.GetFilterPointsAsset(filter, exchange, address, blockchain, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// GetChartPoints godoc
// @Param   scale      query   string     false       "scale 5m 30m 1h 4h 1d 1w"
func (env *Env) GetChartPoints(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

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
	if ok, err := validTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
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
	if !validateInputParams(c) {
		return
	}

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
	if ok, err := validTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
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
	if !validateInputParams(c) {
		return
	}

	var s []string
	var numSymbols int64
	var sortedAssets []dia.AssetVolume
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
			sortedAssets, err = env.RelDB.GetAssetsWithVOL(numSymbols, int64(0), false, "")
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
			sortedAssets, err = env.RelDB.GetAssetsWithVOL(numSymbols, int64(0), false, "")
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

func (env *Env) GetPoolLiquidityByAddress(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}
	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	type localReturn struct {
		Exchange          string
		Blockchain        string
		Address           string
		Time              time.Time
		TotalLiquidityUSD float64
		Liquidity         []dia.AssetVolume
	}

	pool, err := env.RelDB.GetPoolByAddress(blockchain, address)
	if err != nil {
		log.Info("err: ", err)
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cannot find pool"))
		return
	}
	var l localReturn
	l.Exchange = pool.Exchange.Name
	l.Blockchain = pool.Blockchain.Name
	l.Address = pool.Address
	l.Time = pool.Time
	l.Liquidity = pool.Assetvolumes

	// Get total liquidity.
	var totalLiquidity float64
	for _, assetvol := range pool.Assetvolumes {
		price, err := env.DataStore.GetAssetPriceUSDCache(assetvol.Asset)
		if err != nil {
			log.Warnf("no quotation for %v: %v", assetvol.Asset, err)
			totalLiquidity = 0
			break
		}
		totalLiquidity += price * assetvol.Volume
	}
	l.TotalLiquidityUSD = totalLiquidity

	c.JSON(http.StatusOK, l)

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
	address := c.Param("address")
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

		err error
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

func (env *Env) SearchNFTs(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	querystring := c.Param("query")

	var (
		collections = []dia.NFTClass{}
		err         error
	)

	switch {
	case len(querystring) > 4 && strings.Contains(querystring[0:2], "0x"):
		var collection dia.NFTClass
		address := common.HexToAddress(querystring).Hex()
		collection, err = env.RelDB.GetNFTClass(address, dia.ETHEREUM)
		if err != nil {
			log.Errorln("error getting GetNFTByNameSymbol", err)
			restApi.SendError(c, http.StatusInternalServerError, errors.New("Address not valid."))
		}
		collections = append(collections, collection)

	case !strings.Contains(querystring[0:2], "0x"):
		collections, err = env.RelDB.GetNFTClassesByNameSymbol(querystring)
		if err != nil {
			log.Errorln("error getting GetNFTByNameSymbol", err)
			restApi.SendError(c, http.StatusInternalServerError, errors.New("Couldn't find any collections."))
		}

	}

	c.JSON(http.StatusOK, collections)
}

func (env *Env) GetTopAssets(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	numAssetsString := c.Param("numAssets")
	pageString := c.DefaultQuery("Page", "1")
	onlycexString := c.DefaultQuery("Cex", "false")
	blokchain := c.DefaultQuery("Network", "")

	var (
		numAssets    int64
		sortedAssets []dia.AssetVolume
		err          error
		pageNumber   int64
		offset       int64
	)

	pageNumber, err = strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("Page of assets must be an integer"))
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

	sortedAssets, err = env.RelDB.GetAssetsWithVOL(numAssets, offset, onlycex, blokchain)
	if err != nil {
		log.Error("get assets with volume: ", err)

	}
	var assets = []dia.TopAsset{}

	for _, v := range sortedAssets {
		var sources = make(map[string][]string)

		aqf := dia.TopAsset{}
		aqf.Asset = v.Asset
		quotation, err := env.DataStore.GetAssetQuotationLatest(aqf.Asset)
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

		quotationYesterday, err := env.DataStore.GetAssetQuotation(aqf.Asset, time.Now().AddDate(0, 0, -1))
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
	if !validateInputParams(c) {
		return
	}

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
		if ok, err := validTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
			restApi.SendError(c, http.StatusInternalServerError, err)
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

// GetBenchmarkedIndexValue Get benchmarked Index values
func (env *Env) GetBenchmarkedIndexValue(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

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

func (env *Env) GetLastTradeTime(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	exchange := c.Param("exchange")
	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

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
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

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
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")

	q, err := env.RelDB.GetAllNFTClasses(blockchain)
	if len(q) == 0 || err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetNFTClasses returns all NFT classes.
func (env *Env) GetNFTClasses(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

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
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	id := c.Param("id")

	q, err := env.RelDB.GetNFT(address, blockchain, id)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetNFTTrades returns all trades of the unique NFT with given parameters.
func (env *Env) GetNFTTrades(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	id := c.Param("id")
	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*30)*time.Hour)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
		return
	}

	q, err := env.RelDB.GetNFTTrades(address, blockchain, id, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, q)
}

// GetNFTTradesCollection returns all trades of the collection with given parameters.
func (env *Env) GetNFTTradesCollection(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	starttime, endtime, err := utils.MakeTimerange(c.Query("starttime"), c.Query("endtime"), time.Duration(24*30)*time.Hour)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
		return
	}

	q, err := env.RelDB.GetNFTTradesCollection(address, blockchain, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	// Amend output.
	nftClass, err := env.RelDB.GetNFTClass(address, blockchain)
	if err != nil {
		log.Error("get nft class: ", err)
	}

	type nftTradesCollReturn struct {
		Name        string
		Price       float64
		NFTid       string
		FromAddress string
		ToAddress   string
		BundleSale  bool
		BlockNumber uint64
		Timestamp   time.Time
		TxHash      string
		Exchange    string
		Currency    dia.Asset
	}

	var r []nftTradesCollReturn
	for _, trade := range q {
		var t nftTradesCollReturn
		t.Name = nftClass.Name
		price, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(trade.Price), new(big.Float).SetFloat64(math.Pow10(int(trade.Currency.Decimals)))).Float64()
		t.Price = price
		t.Currency = trade.Currency
		t.NFTid = trade.NFT.TokenID
		t.FromAddress = trade.FromAddress
		t.ToAddress = trade.ToAddress
		t.BundleSale = trade.BundleSale
		t.BlockNumber = trade.BlockNumber
		t.Timestamp = trade.Timestamp
		t.TxHash = trade.TxHash
		t.Exchange = trade.Exchange
		r = append(r, t)
	}

	c.JSON(http.StatusOK, r)
}

// GetNFTFloor returns the last floor price of a collection before @timestamp.
func (env *Env) GetNFTFloor(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	timestampUnixString := c.Query("timestamp")
	var timestamp time.Time
	if timestampUnixString != "" {
		timestampUnix, err := strconv.ParseInt(timestampUnixString, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusBadRequest, err)
		}
		timestamp = time.Unix(timestampUnix, 0)
	} else {
		timestamp = time.Now()
	}

	floorWindowSeconds := c.Query("floorWindow")
	var floorWindow int64
	var err error
	if floorWindowSeconds != "" {
		floorWindow, err = strconv.ParseInt(floorWindowSeconds, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusBadRequest, err)
		}
	} else {
		// Set floor window to default 24h.
		floorWindow = 24 * 60 * 60
	}

	// Exclude bundle sales by default.
	bundlesString := c.DefaultQuery("bundles", "false")
	bundles, err := strconv.ParseBool(bundlesString)
	if err != nil {
		log.Error("parse bundles string: ", err)
	}

	nftClass := dia.NFTClass{Address: address, Blockchain: blockchain}

	// Look for floor price. Iterate backwards in time if no sales are found.
	var floor float64
	windowDuration := time.Duration(floorWindow) * time.Second
	stepBackLimit := 40
	floor, err = env.RelDB.GetNFTFloorRecursive(nftClass, timestamp, windowDuration, stepBackLimit, !bundles)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, err)
		return
	}
	type returnStruct struct {
		Floor  float64   `json:"Floor_Price"`
		Time   time.Time `json:"Time"`
		Source string    `json:"Source"`
	}
	var resp returnStruct
	resp.Floor = floor
	resp.Time = timestamp
	resp.Source = dia.Diadata
	c.JSON(http.StatusOK, resp)
}

// GetNFTFloorMA returns the moving average floor price of the nft class over the last 30 days.
func (env *Env) GetNFTFloorMA(c *gin.Context) {

	// NFT collection.
	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	nftClass := dia.NFTClass{Address: address, Blockchain: blockchain}

	// lookback for MA is 30 days per default.
	lookbackString := c.DefaultQuery("lookbackSeconds", "2592000")
	lookbackInt, err := strconv.ParseInt(lookbackString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, nil)
	}

	// floor price window is 24h per default.
	floorWindowString := c.DefaultQuery("floorWindow", "86400")
	floorWindowInt, err := strconv.ParseInt(floorWindowString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, nil)
	}
	floorWindow := time.Duration(floorWindowInt) * time.Second

	// Exclude bundle sales by default.
	bundlesString := c.DefaultQuery("bundles", "false")
	bundles, err := strconv.ParseBool(bundlesString)
	if err != nil {
		log.Error("parse bundles string: ", err)
	}

	endtime := time.Now()
	starttime := endtime.Add(-time.Duration(lookbackInt) * time.Second)
	stepBackLimit := 120

	t := time.Now()
	floorPrices, err := env.RelDB.GetNFTFloorRange(nftClass, starttime, endtime, floorWindow, stepBackLimit, !bundles)
	log.Infof("took %v time to compute floorPrices: %v", time.Since(t), floorPrices)

	cleanFloorPrices, indices := filters.RemoveOutliers(floorPrices, 1.5)
	var floorMA float64
	if len(indices) == 2 {
		floorMA = utils.Average(cleanFloorPrices)
	} else {
		floorMA = utils.Average(floorPrices)
	}
	log.Info("cleaned floorPrices: ", cleanFloorPrices)
	if len(indices) == 2 {
		log.Info("indices: ", indices)
		log.Info("discarded: ", floorPrices[:indices[0]], floorPrices[indices[1]-1:len(floorPrices)-1])
	} else {
		log.Info("nothing discarded.")
	}

	type returnStruct struct {
		Floor  float64   `json:"Moving_Average_Floor_Price"`
		Time   time.Time `json:"Time"`
		Source string    `json:"Source"`
	}
	var resp returnStruct
	resp.Floor = floorMA
	resp.Time = endtime
	resp.Source = dia.Diadata
	c.JSON(http.StatusOK, resp)
}

// GetNFTDownday returns the moving average floor price of the nft class over the last 30 days.
func (env *Env) GetNFTDownday(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	// NFT collection.
	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	nftClass := dia.NFTClass{Address: address, Blockchain: blockchain}

	// lookback for MA is 90 days per default.
	lookbackString := c.DefaultQuery("lookbackSeconds", "7776000")
	lookbackInt, err := strconv.ParseInt(lookbackString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, nil)
	}

	// floor price window is 24h per default.
	floorWindowString := c.DefaultQuery("floorWindow", "86400")
	floorWindowInt, err := strconv.ParseInt(floorWindowString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, nil)
	}
	floorWindow := time.Duration(floorWindowInt) * time.Second

	// Exclude bundle sales by default.
	bundlesString := c.DefaultQuery("bundles", "false")
	bundles, err := strconv.ParseBool(bundlesString)
	if err != nil {
		log.Error("parse bundles string: ", err)
	}

	endtime := time.Now()
	starttime := endtime.Add(-time.Duration(lookbackInt) * time.Second)
	stepBackLimit := 120
	floorPrices, err := env.RelDB.GetNFTFloorRange(nftClass, starttime, endtime, floorWindow, stepBackLimit, !bundles)

	log.Info("floorPrices: ", floorPrices)

	// Compute movement time-series.
	var movement []float64
	var downwardMovement []float64
	for i := range floorPrices {
		if i == len(floorPrices)-1 {
			break
		}
		if floorPrices[i] == 0 {
			continue
		}
		mov := 100 * (floorPrices[i+1] - floorPrices[i]) / floorPrices[i]
		movement = append(movement, mov)
		if mov < 0 {
			downwardMovement = append(downwardMovement, mov)
		}
	}
	log.Info("movement: ", movement)

	type Downward struct {
		WeeklyDrawdown   float64   `json:"Weekly_Drawdown"`
		DowndayAverage   float64   `json:"Downday_Average"`
		DowndayDeviation float64   `json:"Downday_Deviation"`
		Time             time.Time `json:"Time"`
		Source           string    `json:"Source"`
	}
	var response Downward

	response.DowndayAverage = utils.Average(downwardMovement)
	response.DowndayDeviation = utils.StandardDeviation(downwardMovement)

	// Caution: This is only valid for 24h windows.
	// response.WeeklyDrawdown = 0
	// if len(movement) > 7 {
	// 	for i := 7; i < len(movement)-1; i++ {
	// 		diff := movement[i] - movement[i-7]
	// 		if diff < response.WeeklyDrawdown {
	// 			response.WeeklyDrawdown = diff
	// 		}
	// 	}
	// }

	var drawdowns []float64
	if len(movement) > 7 {
		for i := 7; i < len(movement)-1; i++ {
			drawdowns = append(drawdowns, movement[i]-movement[i-1])
		}
	}
	cleanDrawdowns, _ := filters.RemoveOutliers(drawdowns, float64(1.5))
	min := cleanDrawdowns[0]
	for _, x := range cleanDrawdowns {
		if x < min {
			min = x
		}
	}
	response.WeeklyDrawdown = min
	response.Time = endtime
	response.Source = dia.Diadata

	c.JSON(http.StatusOK, response)
}

// GetNFTFloorVola returns the volatility of the moving average floor price of the nft class over the last 90 days.
func (env *Env) GetNFTFloorVola(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	// NFT collection.
	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	nftClass := dia.NFTClass{Address: address, Blockchain: blockchain}

	// Parse query parameter time.
	endtimeString := c.Query("time")
	var endtime time.Time
	if endtimeString != "" {
		endtimeInt, err := strconv.ParseInt(endtimeString, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	} else {
		endtime = time.Now()
	}

	// lookback for volatility is 90 days per default.
	lookbackString := c.DefaultQuery("lookbackSeconds", "7776000")
	lookbackInt, err := strconv.ParseInt(lookbackString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, nil)
	}

	// floor price window is 24h per default.
	floorWindowString := c.DefaultQuery("floorWindow", "86400")
	floorWindowInt, err := strconv.ParseInt(floorWindowString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, nil)
	}
	floorWindow := time.Duration(floorWindowInt) * time.Second

	// Exclude bundle sales by default.
	bundlesString := c.DefaultQuery("bundles", "false")
	bundles, err := strconv.ParseBool(bundlesString)
	if err != nil {
		log.Error("parse bundles string: ", err)
	}

	starttime := endtime.Add(-time.Duration(lookbackInt) * time.Second)
	stepBackLimit := 120
	floorPrices, err := env.RelDB.GetNFTFloorRange(nftClass, starttime, endtime, floorWindow, stepBackLimit, !bundles)
	if err != nil {
		log.Error("get nft floor range: ", err)
	}

	// Get collection name.
	nftClass, err = env.RelDB.GetNFTClass(nftClass.Address, nftClass.Blockchain)
	if err != nil {
		log.Error("get nft class: ", err)
	}

	type FloorStats struct {
		FloorAverage    float64   `json:"Floor_Average"`
		FloorVolatility float64   `json:"Floor_Volatility"`
		Collection      string    `json:"Collection"`
		Time            time.Time `json:"Time"`
		Source          string    `json:"Source"`
	}
	var response FloorStats

	response.FloorAverage = utils.Average(floorPrices)
	response.FloorVolatility = utils.StandardDeviation(floorPrices)
	response.Time = endtime
	response.Collection = nftClass.Name
	response.Source = dia.Diadata

	c.JSON(http.StatusOK, response)
}

func (env *Env) GetTopNFTClasses(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	type localReturn struct {
		Collection   string
		Floor        float64
		FloorMA      float64
		Volume       float64
		Trades       int
		FloorChange  float64
		VolumeChange float64
		TradesChange float64
		Address      string
		Blockchain   string
		Time         time.Time
		Source       string
	}
	var (
		starttime   time.Time
		endtime     time.Time
		returnValue []localReturn
		pageNumber  int64
		offset      int64
	)

	numCollections, err := strconv.Atoi(c.Param("numCollections"))
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	pageString := c.DefaultQuery("page", "1")
	pageNumber, err = strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("Page of assets must be an integer"))
	}
	offset = (pageNumber - 1) * int64(numCollections)

	endtimeString := c.Query("endtime")
	if endtimeString != "" {
		endtimeInt, err := strconv.ParseInt(endtimeString, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	} else {
		endtime = time.Now()
	}
	starttimeString := c.Query("starttime")
	if starttimeString != "" {
		starttimeInt, err := strconv.ParseInt(starttimeString, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
	} else {
		starttime = endtime.AddDate(0, 0, -1)
	}

	exchangesString := c.Query("exchanges")
	var exchanges []string
	if exchangesString != "" {
		exchanges = strings.Split(exchangesString, ",")
	}

	// Exclude bundle sales by default.
	bundlesString := c.DefaultQuery("bundles", "false")
	bundles, err := strconv.ParseBool(bundlesString)
	if err != nil {
		log.Error("parse bundles string: ", err)
	}

	var window24h = time.Duration(24 * 60 * time.Minute)

	nftVolumes, err := env.RelDB.GetTopNFTsEth(numCollections, offset, exchanges, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	for _, nftvolume := range nftVolumes {
		floor, err := env.RelDB.GetNFTFloor(
			dia.NFTClass{Address: nftvolume.Address, Blockchain: nftvolume.Blockchain},
			endtime,
			window24h,
			!bundles,
		)
		if err != nil {
			log.Errorf("get number of nft trades for address %s: %v", nftvolume.Address, err)
		}

		// ------------- Floor MA -------------
		floorPrices, err := env.RelDB.GetNFTFloorRange(
			dia.NFTClass{Address: nftvolume.Address, Blockchain: nftvolume.Blockchain},
			endtime.AddDate(0, 0, -30),
			endtime,
			window24h,
			20,
			!bundles,
		)
		if err != nil {
			log.Errorf("get floor range for address %s: %v", nftvolume.Address, err)
		}

		cleanFloorPrices, indices := filters.RemoveOutliers(floorPrices, 1.5)
		var floorMA float64
		if len(indices) == 2 {
			floorMA = utils.Average(cleanFloorPrices)
		} else {
			floorMA = utils.Average(floorPrices)
		}
		// -------------------------------------

		floorYesterday, err := env.RelDB.GetNFTFloor(
			dia.NFTClass{Address: nftvolume.Address, Blockchain: nftvolume.Blockchain},
			endtime.Add(-window24h),
			window24h,
			!bundles,
		)
		if err != nil {
			log.Errorf("get floor yesterday for address %s: %v", nftvolume.Address, err)
		}

		numTrades, err := env.RelDB.GetNumNFTTrades(nftvolume.Address, nftvolume.Blockchain, "", starttime, endtime)
		if err != nil {
			log.Errorf("get number of nft trades for address %s: %v", nftvolume.Address, err)
		}
		numTradesYesterday, err := env.RelDB.GetNumNFTTrades(nftvolume.Address, nftvolume.Blockchain, "", starttime.Add(-window24h), endtime.Add(-window24h))
		if err != nil {
			log.Errorf("get number of nft trades yesterday for address %s: %v", nftvolume.Address, err)
		}
		volumeYesterday, err := env.RelDB.GetNFTVolume(nftvolume.Address, nftvolume.Blockchain, "", starttime.Add(-window24h), endtime.Add(-window24h))
		if err != nil {
			log.Errorf("get volume yesterday for address %s: %v", nftvolume.Address, err)
		}

		var l localReturn
		l.Collection = nftvolume.Name
		l.Floor = floor
		l.FloorMA = floorMA
		l.Volume = nftvolume.Volume
		if volumeYesterday > 0 {
			l.VolumeChange = (nftvolume.Volume - volumeYesterday) / volumeYesterday * 100
		}
		l.Trades = numTrades
		if numTradesYesterday > 0 {
			l.TradesChange = float64(numTrades-numTradesYesterday) / float64(numTradesYesterday) * 100
		}
		if floorYesterday > 0 {
			l.FloorChange = (floor - floorYesterday) / floorYesterday * 100
		}
		l.Address = nftvolume.Address
		l.Blockchain = nftvolume.Blockchain
		l.Time = endtime
		l.Source = dia.Diadata
		returnValue = append(returnValue, l)
	}

	c.JSON(http.StatusOK, returnValue)

}

func (env *Env) GetNFTVolume(c *gin.Context) {

	type localReturn struct {
		Collection   string
		Floor        float64
		Volume       float64
		Trades       int
		FloorChange  float64
		VolumeChange float64
		TradesChange float64
		Address      string
		Blockchain   string
		Time         time.Time
		Source       string
		Exchanges    []dia.NFTExchangeStats
	}
	var (
		starttime time.Time
		endtime   time.Time
		l         localReturn
	)

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	endtimeString := c.Query("endtime")
	if endtimeString != "" {
		endtimeInt, err := strconv.ParseInt(endtimeString, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	} else {
		endtime = time.Now()
	}
	starttimeString := c.Query("starttime")
	if starttimeString != "" {
		starttimeInt, err := strconv.ParseInt(starttimeString, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
	} else {
		starttime = endtime.AddDate(0, 0, -1)
	}
	timeWindow := endtime.Sub(starttime)

	// Exclude bundle sales by default.
	bundlesString := c.DefaultQuery("bundles", "false")
	bundles, err := strconv.ParseBool(bundlesString)
	if err != nil {
		log.Error("parse bundles string: ", err)
	}

	collection, err := env.RelDB.GetNFTClass(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusBadRequest, err)
		return
	}

	floor, err := env.RelDB.GetNFTFloorRecursive(
		dia.NFTClass{Address: address, Blockchain: blockchain},
		endtime,
		timeWindow,
		10,
		!bundles,
	)
	if err != nil {
		log.Error("get floor: ", err)
	}
	floorYesterday, err := env.RelDB.GetNFTFloorRecursive(
		dia.NFTClass{Address: address, Blockchain: blockchain},
		endtime.Add(-timeWindow),
		timeWindow,
		10,
		!bundles,
	)
	if err != nil {
		log.Error("get floor yesterday: ", err)
	}
	volume, err := env.RelDB.GetNFTVolume(address, blockchain, "", starttime, endtime)
	if err != nil {
		log.Error("get volume: ", err)
	}
	volumeYesterday, err := env.RelDB.GetNFTVolume(address, blockchain, "", starttime.Add(-timeWindow), endtime.Add(-timeWindow))
	if err != nil {
		log.Error("get volume yesterday: ", err)
	}
	numTrades, err := env.RelDB.GetNumNFTTrades(address, blockchain, "", starttime, endtime)
	if err != nil {
		log.Error("get number of nft trades: ", err)
	}
	numTradesYesterday, err := env.RelDB.GetNumNFTTrades(address, blockchain, "", starttime.Add(-timeWindow), endtime.Add(-timeWindow))
	if err != nil {
		log.Error("get number of nft trades yesterday: ", err)
	}

	exchanges, err := env.RelDB.GetNFTExchanges(address, blockchain)
	if err != nil {
		log.Error("get number of nft trades yesterday: ", err)
	}

	for _, exchange := range exchanges {
		numTrades, err := env.RelDB.GetNumNFTTrades(address, blockchain, exchange, starttime, endtime)
		if err != nil {
			log.Error("get number of nft trades: ", err)
		}
		volume, err := env.RelDB.GetNFTVolume(address, blockchain, exchange, starttime, endtime)
		if err != nil {
			log.Error("get number of nft trades: ", err)
		}
		l.Exchanges = append(l.Exchanges, dia.NFTExchangeStats{Exchange: exchange, NumTrades: uint64(numTrades), Volume: volume})
	}

	l.Collection = collection.Name
	l.Floor = floor
	if floorYesterday > 0 {
		l.FloorChange = (floor - floorYesterday) / floorYesterday * 100
	}

	l.Volume = volume
	if volumeYesterday > 0 {
		l.VolumeChange = (volume - volumeYesterday) / volumeYesterday * 100
	}

	l.Trades = numTrades
	if numTradesYesterday > 0 {
		l.TradesChange = float64(numTrades-numTradesYesterday) / float64(numTradesYesterday) * 100
	}
	l.Address = collection.Address
	l.Blockchain = collection.Blockchain
	l.Time = endtime
	l.Source = dia.Diadata

	c.JSON(http.StatusOK, l)

}

func (env *Env) GetFeedStats(c *gin.Context) {
	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")
	var starttime time.Time
	var endtime time.Time

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
	if starttimeStr == "" {
		starttime = endtime.AddDate(0, 0, -1)
	} else {
		starttimeInt, err := strconv.ParseInt(starttimeStr, 10, 64)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
			return
		}
		starttime = time.Unix(starttimeInt, 0)
	}
	if ok, err := validTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	asset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	exchVolumes, err := env.RelDB.GetAggVolumesByExchange(asset, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	pairVolumes, err := env.RelDB.GetAggVolumesByPair(asset, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	tradesDist, err := env.RelDB.GetTradesDistribution(asset, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, nil)
	}

	type localDistType struct {
		NumTradesTotal   int     `json:"NumTradesTotal"`
		NumBins          int     `json:"NumBins"`
		NumLowBins       int     `json:"NumberLowBins"`
		Threshold        int     `json:"Threshold"`
		SizeBinSeconds   int64   `json:"SizeBin"`
		AvgNumPerBin     float64 `json:"AverageNumberPerBin"`
		StdDeviation     float64 `json:"StandardDeviation"`
		TimeRangeSeconds int64   `json:"TimeRangeSeconds"`
	}
	var tradesDistReduced []localDistType
	for _, val := range tradesDist {
		tradesDistReduced = append(tradesDistReduced, localDistType{
			NumTradesTotal:   val.NumTradesTotal,
			NumBins:          int(val.TimeRangeSeconds) / int(val.SizeBinSeconds),
			NumLowBins:       val.NumLowBins,
			Threshold:        val.Threshold,
			SizeBinSeconds:   val.SizeBinSeconds,
			AvgNumPerBin:     val.AvgNumPerBin,
			StdDeviation:     val.StdDeviation,
			TimeRangeSeconds: val.TimeRangeSeconds,
		})
	}

	type localReturn struct {
		Timestamp          time.Time
		TotalVolume        float64
		Price              float64
		TradesDistribution localDistType
		ExchangeVolumes    []dia.ExchangeVolume
		PairVolumes        []dia.PairVolume
	}

	var retVal []localReturn

	// Fill local return type.
	for i := range exchVolumes {
		var l localReturn
		var price float64
		sort.Slice(exchVolumes[i].Volumes, func(m, n int) bool { return exchVolumes[i].Volumes[m].Volume > exchVolumes[i].Volumes[n].Volume })
		l.ExchangeVolumes = exchVolumes[i].Volumes
		// Compute total volume.
		for _, vol := range l.ExchangeVolumes {
			l.TotalVolume += vol.Volume
		}
		sort.Slice(pairVolumes[i].Volumes, func(m, n int) bool { return pairVolumes[i].Volumes[m].Volume > pairVolumes[i].Volumes[n].Volume })
		l.PairVolumes = pairVolumes[i].Volumes
		l.Timestamp = exchVolumes[i].Timestamp
		// Get Price.
		price, err = env.DataStore.GetAssetPriceUSD(asset, l.Timestamp)
		if err != nil {
			log.New().Errorf("get price usd for asset %v: %v", asset, err)
		}
		l.Price = price
		if len(tradesDistReduced) > i {
			l.TradesDistribution = tradesDistReduced[i]
		}
		retVal = append(retVal, l)
	}

	// If no time-range is given, don't return a slice of length 1.
	if endtimeStr == "" && starttimeStr == "" {
		if len(retVal) > 0 {
			c.JSON(http.StatusOK, retVal[0])
		}
	} else {
		c.JSON(http.StatusOK, retVal)
	}

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
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

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
		restApi.SendError(c, http.StatusRequestedRangeNotSatisfiable, errors.New("Requested time-range too large."))
		return
	}
	if ok, err := validTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	asset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	quotations, err := env.DataStore.GetAssetQuotations(asset, starttime, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
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
	address := makeAddressEIP55Compliant(c.Param("address"), blockchain)

	endtime := time.Now()
	starttime := endtime.AddDate(0, 0, -1)

	var quotationExtended localAssetInfoReturn

	asset, err := env.RelDB.GetAsset(address, blockchain)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, err)
		return
	}

	quotation, err := env.DataStore.GetAssetQuotation(asset, endtime)
	if err != nil {
		restApi.SendError(c, http.StatusNotFound, errors.New("no quotation available"))
		return
	}
	quotationYesterday, err := env.DataStore.GetAssetQuotation(asset, starttime)
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
	exchangemap, err := env.DataStore.GetActiveExchangesAndPairs(asset.Address, asset.Blockchain, starttime, endtime)
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

// GetSyntheticAsset
func (env *Env) GetSyntheticAsset(c *gin.Context) {
	var (
		err error
		p   []dia.SynthAssetSupply
	)

	if !validateInputParams(c) {
		return
	}

	blockchain := c.Param("blockchain")
	protocol := c.Param("protocol")

	address := makeAddressEIP55Compliant(c.Query("address"), blockchain)

	starttimeStr := c.Query("starttime")
	endtimeStr := c.Query("endtime")

	limit := 0

	// Set times depending on what is given by the query parameters
	var starttime, endtime time.Time
	if starttimeStr == "" && endtimeStr == "" {
		// Last seven days per default
		starttime = time.Now().AddDate(0, 0, -7)
		endtime = time.Now()
		limit = 1
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
	if ok, err := validTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}

	if address != "" && address != "0x0000000000000000000000000000000000000000" {

		p, err = env.DataStore.GetSynthSupplyInflux(blockchain, protocol, address, limit, starttime, endtime)

	} else {
		synthassets, err := env.DataStore.GetSynthAssets(blockchain, protocol)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, errors.New("no response for quoted timestamp"))
		}
		for _, asset := range synthassets {
			points, _ := env.DataStore.GetSynthSupplyInflux(blockchain, protocol, asset, limit, starttime, endtime)
			if err != nil {
				log.Errorln("GetSynthSupplyInflux", err)
			} else {
				if len(points) > 0 {
					p = append(p, points[0])
				}

			}
		}
	}

	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		var response []map[string]interface{}

		for _, v := range p {
			row := make(map[string]interface{})
			row["Blockchain"] = v.Asset.Blockchain
			row["UnderlyingTokenAddress"] = v.AssetUnderlying.Address
			row["UnderlyingTokenSymbol"] = v.AssetUnderlying.Symbol
			row["SyntheticTokenAddress"] = v.Asset.Address
			row["SyntheticTokenSymbol"] = v.Asset.Symbol
			row["TotalDebt"] = v.TotalDebt
			row["BlockNumber"] = v.BlockNumber
			row["CollateralRatio"] = v.ColleteralRatio
			row["LockedUnderlying"] = v.LockedUnderlying
			row["Supply"] = v.Supply
			row["Protocol"] = v.Protocol
			row["Time"] = v.Time
			response = append(response, row)

		}

		c.JSON(http.StatusOK, response)
	}
}

func validTimeRange(starttime time.Time, endtime time.Time, maxDuration time.Duration) (ok bool, err error) {
	if endtime.Sub(starttime) < maxDuration {
		ok = true
	}
	err = fmt.Errorf("time-range too big. max duration is %v", maxDuration)
	return
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
