package diaApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type Env struct {
	DataStore models.Datastore
}

// PostSupply godoc
// @Summary Post the circulating supply
// @Description Post the circulating supply
// @Tags dia
// @Accept  json
// @Produce  json
// @Param Symbol query string true "Coin symbol"
// @Param CirculatingSupply query float64 true "number of coins in circulating supply"
// @Success 200 {object} dia.Supply	"success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/supply [post]
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
			if t.Symbol == "" || t.CirculatingSupply == 0.0 {
				log.Errorln("received supply:", t)
				restApi.SendError(c, http.StatusInternalServerError, errors.New("Missing Symbol or CirculatingSupply value"))
			} else {
				log.Println("received supply:", t)
				source := dia.Diadata
				if t.Source != "" {
					source = t.Source
				}
				s := &dia.Supply{
					Time:              t.Time,
					Name:              helpers.NameForSymbol(t.Symbol),
					Symbol:            t.Symbol,
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

// GetQuotation godoc
// @Summary Get quotation
// @Description GetQuotation
// @Tags dia
// @Accept  json
// @Produce  json
// @Param   symbol     path    string     true        "Some symbol"
// @Success 200 {object} models.Quotation "success"
// @Failure 404 {object} restApi.APIError "Symbol not found"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/quotation/:symbol: [get]
func (env *Env) GetQuotation(c *gin.Context) {
	symbol := c.Param("symbol")
	q, err := env.DataStore.GetQuotation(symbol)
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

// GetSupply returns latest supply of token with @symbol
func (env *Env) GetSupply(c *gin.Context) {
	symbol := c.Param("symbol")
	s, err := env.DataStore.GetLatestSupply(symbol)
	if err != nil {
		if err == redis.Nil {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, s)
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

	s, err := env.DataStore.GetSupplyInflux(symbol, starttime, endtime)
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

	v, err := env.DataStore.GetVolumeInflux(symbol, starttime, endtime)
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
	p, err := env.DataStore.GetPairs("")
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, &models.Pairs{Pairs: p})
	}
}

// GetSymbolDetails godoc
// @Summary Get Symbol Details
// @Description Get Symbol Details
// @Tags dia
// @Accept  json
// @Produce  json
// @Param   symbol     path    string     true        "Some symbol"
// @Success 200 {object} models.SymbolDetails "success"
// @Failure 404 {object} restApi.APIError "Symbol not found"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/symbol/:symbol: [get]
func (env *Env) GetSymbolDetails(c *gin.Context) {
	symbol := c.Param("symbol")

	s, err := env.DataStore.GetSymbolDetails(symbol)
	if err != nil {
		if err == redis.Nil {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, s)
	}
}

func roundUpTime(t time.Time, roundOn time.Duration) time.Time {
	t = t.Round(roundOn)
	if time.Since(t) >= 0 {
		t = t.Add(roundOn)
	}
	return t
}

// GetCoins godoc
// @Summary Get coins
// @Description GetCoins
// @Tags dia
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Coins "success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/coins [get]
func (env *Env) GetCoins(c *gin.Context) {
	coins, err := env.DataStore.GetCoins()
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, coins)
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
// @Summary Get chart points for
// @Description Get Symbol Details
// @Tags dia
// @Accept  json
// @Produce  json
// @Param   symbol     path    string     true        "Some symbol"
// @Param   exchange     path    string     true        "Some exchange"
// @Param   filter     path    string     true        "Some filter"
// @Param   scale      query   string     false       "scale 5m 30m 1h 4h 1d 1w"
// @Success 200 {object} models.Points "success"
// @Failure 404 {object} restApi.APIError "Symbol not found"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/chartPoints/:filter/:exchange:/:symbol: [get]
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

// GetAllSymbols returns all symbols available in our (redis) database.
// Optional query parameter exchange returns only symbols available on this exchange.
func (env *Env) GetAllSymbols(c *gin.Context) {
	exchange := c.DefaultQuery("exchange", "noRange")
	if exchange == "noRange" {
		s := env.DataStore.GetAllSymbols()
		if len(s) == 0 {
			restApi.SendError(c, http.StatusInternalServerError, errors.New("cant find symbols"))
		} else {
			c.JSON(http.StatusOK, dia.Symbols{Symbols: s})
		}
	} else {
		s := env.DataStore.GetSymbolsByExchange(exchange)
		if len(s) == 0 {
			restApi.SendError(c, http.StatusInternalServerError, errors.New("cant find symbols"))
		} else {
			c.JSON(http.StatusOK, dia.Symbols{Symbols: s})
		}
	}

}

// -----------------------------------------------------------------------------
// INDICES
// -----------------------------------------------------------------------------

func (env *Env) GetCviIndex(c *gin.Context) {
	starttimeStr := c.DefaultQuery("starttime", "noRange")
	endtimeStr := c.Query("endtime")

	var starttime, endtime time.Time

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
	q, err := env.DataStore.GetCVIInflux(starttime, endtime)
	for i := range q {
		q[i].Value /= 2430.5812295231785
	}
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
	// 	if err == redis.Nil {
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
			if err == redis.Nil {
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
			if err == redis.Nil {
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
// FARMING POOLS
// -----------------------------------------------------------------------------

// GetFarmingPools is the delegate method to fetch the value(s) of
// the farming pool information of @protocol.
// Last value is retrieved. Otional query parameters allow to obtain data in a time range.
func (env *Env) GetFarmingPools(c *gin.Context) {
	q, err := env.DataStore.GetFarmingPools()
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
			if err == redis.Nil {
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
		q, err := env.DataStore.GetFarmingPoolData(starttime, endtime, protocol, poolID)
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
			if err == redis.Nil {
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

		date := time.Time{}
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
			if err == redis.Nil {
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

// GetCompoundedAvg is the delegate method to fetch averaged compounded rate values for interest rates
func (env *Env) GetCompoundedAvg(c *gin.Context) {

	tInit := time.Now()

	// Import and cast input from API call
	symbol := c.Param("symbol")
	datestring := c.Param("time")
	date, _ := time.Parse("2006-01-02", datestring)
	days := c.Param("days")
	calDays, err := strconv.Atoi(days)
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
			if err == redis.Nil {
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
			if err == redis.Nil {
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
			if err == redis.Nil {
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
			if err == redis.Nil {
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
		if err == redis.Nil {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
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
		if err == redis.Nil {
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
		if err == redis.Nil {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		c.JSON(http.StatusOK, q)
	}

}
