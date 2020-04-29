package diaApi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	models "github.com/diadata-org/diadata/pkg/model"
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
			restApi.SendError(c, http.StatusInternalServerError, errors.New("Unmarshal"))
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
					Time:              time.Now(),
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
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, q)
}

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

// GetCompoundedRate is the delegate method to fetch compunded rate values for interest rates
func (env *Env) GetCompoundedRate(c *gin.Context) {
	// Import and cast input from API call
	symbol := c.Param("symbol")
	datestring := c.Param("time")
	date, _ := time.Parse("2006-01-02", datestring)
	// zone := c.Param("zone")
	dpy := c.Param("dpy")
	daysPerYear, err := strconv.Atoi(dpy)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	}
	// Compute compunded rate and return if no error
	q, err := env.DataStore.GetCompoundedRate(symbol, date, daysPerYear)
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

// GetSupply godoc
// @Summary Get supply
// @Description GetSupply
// @Tags dia
// @Accept  json
// @Produce  json
// @Param   symbol     path    string     true        "Some symbol"
// @Success 200 {object} dia.Supply "success"
// @Failure 404 {object} restApi.APIError "Symbol not found"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/supply/:symbol: [get]
func (env *Env) GetSupply(c *gin.Context) {
	symbol := c.Param("symbol")
	s, err := env.DataStore.GetSupply(symbol)
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

// GetPairs godoc
// @Summary Get pairs
// @Description Get pairs
// @Tags dia
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Pairs "success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/pairs/ [get]
func (env *Env) GetPairs(c *gin.Context) {
	p, err := env.DataStore.GetPairs("")
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, &models.Pairs{Pairs: p})
	}
}

// GetSymbol godoc
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

	p, err := env.DataStore.GetFilterPoints(filter, exchange, symbol, scale)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// GetChartPointsAllExchange godoc
// @Summary Get Symbol Details
// @Description Get Symbol Details
// @Tags dia
// @Accept  json
// @Produce  json
// @Param   symbol     path    string     true        "Some symbol"
// @Param   filter     path    string     true        "Some filter"
// @Param   scale      query   string     false       "scale 5m 30m 1h 4h 1d 1w"
// @Success 200 {object} models.Points "success"
// @Failure 404 {object} restApi.APIError "Symbol not found"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/chartPointsAllExchanges/:filter:/:symbol: [get]
func (env *Env) GetChartPointsAllExchanges(c *gin.Context) {
	filter := c.Param("filter")
	symbol := c.Param("symbol")
	scale := c.Query("scale")

	p, err := env.DataStore.GetFilterPoints(filter, "", symbol, scale)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// GetAllSymbols godoc
// @Summary Get all symbols list
// @Description Get all symbols list
// @Tags dia
// @Accept  json
// @Produce  json
// @Param   symbol     path    string     true        "Some symbol"
// @Param   filter     path    string     true        "Some filter"
// @Param   scale      query   string     false       "scale 5m 30m 1h 4h 1d 1w"
// @Success 200 {object} dia.Symbols "success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/symbols [get]
func (env *Env) GetAllSymbols(c *gin.Context) {
	s := env.DataStore.GetAllSymbols()
	if len(s) == 0 {
		restApi.SendError(c, http.StatusInternalServerError, errors.New("cant find symbols"))
	} else {
		c.JSON(http.StatusOK, dia.Symbols{Symbols: s})
	}
}
