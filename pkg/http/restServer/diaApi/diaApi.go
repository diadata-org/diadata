package diaApi

import (
	"encoding/json"
	"errors"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/influxdata/influxdb/client/v2"
	clientInfluxdb "github.com/influxdata/influxdb/client/v2"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

const (
	coinsPerPage = 50
)

type Env struct {
	DataStore models.Datastore
}

type points struct {
	DataPoints []clientInfluxdb.Result
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

				s := &dia.Supply{
					Time:              time.Now(),
					Name:              helpers.NameForSymbol(t.Symbol),
					Symbol:            t.Symbol,
					Source:            dia.Diadata,
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
// @Success 200 {object} diaApi.Pairs "success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/pairs/ [get]
func (env *Env) GetPairs(c *gin.Context) {
	p, err := env.DataStore.GetPairs("")
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, &Pairs{Pairs: p})
	}
}

// GetSymbol godoc
// @Summary Get Symbol Details
// @Description Get Symbol Details
// @Tags dia
// @Accept  json
// @Produce  json
// @Param   symbol     path    string     true        "Some symbol"
// @Success 200 {object} diaApi.SymbolDetails "success"
// @Failure 404 {object} restApi.APIError "Symbol not found"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/symbol/:symbol: [get]
func (env *Env) GetSymbolDetails(c *gin.Context) {
	symbol := c.Param("symbol")

	q, err := env.DataStore.GetQuotation(symbol)
	if err != nil {
		if err == redis.Nil {
			restApi.SendError(c, http.StatusNotFound, err)
		} else {
			restApi.SendError(c, http.StatusInternalServerError, err)
		}
	} else {
		r := &SymbolDetails{
			Coin: Coin{
				Symbol:             q.Symbol,
				Name:               q.Name,
				Price:              q.Price,
				VolumeYesterdayUSD: q.VolumeYesterdayUSD,
				Time:               q.Time,
				PriceYesterday:     q.PriceYesterday,
			},
			Exchanges: []models.SymbolExchangeDetails{},
		}
		r.Change, _ = env.DataStore.GetCurrencyChange()

		s, err := env.DataStore.GetSupply(symbol)
		if err == nil {
			r.Coin.CirculatingSupply = &s.CirculatingSupply
		}

		exs, err := env.DataStore.GetExchangesForSymbol(symbol)
		if err != nil {
			restApi.SendError(c, http.StatusInternalServerError, err)
		} else {
			for _, e := range exs {
				s, err2 := env.DataStore.GetSymbolExchangeDetails(symbol, e)
				if err2 == nil {
					r.Exchanges = append(r.Exchanges, *s)
				}
			}
			c.JSON(http.StatusOK, r)
		}
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
// @Success 200 {object} diaApi.Coins "success"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/coins [get]
func (env *Env) GetCoins(c *gin.Context) {

	symbols, err := env.DataStore.SymbolsWithASupply()
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {

		var coins Coins
		coins.Coins = []Coin{}
		coins.CompleteCoinList = []CoinSymbolAndName{}
		coins.Change, _ = env.DataStore.GetCurrencyChange()

		for _, symbol := range symbols {
			var c1 Coin
			log.Debug("Adding symbol", symbol)
			supply, _ := env.DataStore.GetSupply(symbol)
			price, _ := env.DataStore.GetQuotation(symbol)
			volume, _ := env.DataStore.GetVolume(symbol)

			if price != nil && supply != nil && volume != nil {
				c1.Price = price.Price
				c1.Name = price.Name
				c1.Symbol = price.Symbol
				if price.PriceYesterday != nil {
					c1.PriceYesterday = price.PriceYesterday
				}
				c1.Time = price.Time
				c1.VolumeYesterdayUSD = volume
				if supply != nil {
					c1.CirculatingSupply = &supply.CirculatingSupply
					coins.Coins = append(coins.Coins, c1)
				}
			} else {
				log.Warningln("no price, supply, or volume for ", symbol, price, supply, volume)
			}
		}

		sort.Slice(coins.Coins, func(i, j int) bool {
			return (*coins.Coins[i].CirculatingSupply * coins.Coins[i].Price) > (*coins.Coins[j].CirculatingSupply * coins.Coins[j].Price)
		})
		for _, coin := range coins.Coins {
			coins.CompleteCoinList = append(coins.CompleteCoinList, CoinSymbolAndName{coin.Symbol, coin.Name})
		}
		if len(coins.Coins) > coinsPerPage {
			coins.Coins = coins.Coins[:coinsPerPage]
		}
		c.JSON(http.StatusOK, coins)
	}
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
// @Success 200 {object} diaApi.points "success"
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
		c.JSON(http.StatusOK, points{DataPoints: p})
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
// @Success 200 {object} diaApi.points "success"
// @Failure 404 {object} restApi.APIError "Symbol not found"
// @Failure 500 {object} restApi.APIError "error"
// @Router /v1/chartPointsAllExchanges/:symbol:/:symbol: [get]
func (env *Env) GetChartPointsAllExchanges(c *gin.Context) {
	filter := c.Param("filter")
	symbol := c.Param("symbol")
	scale := c.Query("scale")

	p, err := env.DataStore.GetFilterPoints(filter, "", symbol, scale)
	if err != nil {
		restApi.SendError(c, http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, points{DataPoints: p})
	}
}
