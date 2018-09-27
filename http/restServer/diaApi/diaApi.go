package diaApi

import (
	"encoding/json"
	"errors"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers"
	"github.com/diadata-org/api-golang/http/restApi"
	"github.com/diadata-org/api-golang/services/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
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
				log.Error("received supply:", t)
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

		for _, symbol := range symbols {

			var c1 Coin

			log.Println("Adding symbol", symbol)
			supply, _ := env.DataStore.GetSupply(symbol)
			price, _ := env.DataStore.GetQuotation(symbol)
			volume, _ := env.DataStore.GetVolume(symbol)
			if price != nil {
				c1.Price = price.Price
				c1.Name = price.Name
				c1.Symbol = price.Symbol
				if price.PriceYesterday != nil {
					c1.PriceYesterday = price.PriceYesterday
				}
				c1.Time = price.Time
				c1.VolumeYesterdayUSD = volume
			}
			if supply != nil {
				c1.CirculatingSupply = supply.CirculatingSupply
			}
			c1.VolumeYesterdayUSD, _ = env.DataStore.GetVolume(symbol)
			coins.Coins = append(coins.Coins, c1)
		}
		c.JSON(http.StatusOK, coins)
	}
}
