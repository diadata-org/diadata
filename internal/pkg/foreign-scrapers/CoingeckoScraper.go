package foreignscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type CoinIds []struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type CoinIs struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	Symbol      string                   `json:"symbol"`
	LastUpdated string                   `json: "last_updated"`
	Tickers     []map[string]interface{} `json:"tickers"`
	Market      map[string]interface{}   `json:"market_data"`
}

//var _coingeckourl string = "https://api.coingecko.com/api/v3"

type CoingeckoScraper struct {
	error     error
	datastore *models.DB
}

func NewCoingeckoScraper(datastore *models.DB) *CoingeckoScraper {
	s := &CoingeckoScraper{
		datastore: datastore,
		error:     nil,
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CoingeckoScraper) mainLoop() {

	err := scraper.Update()
	if err != nil {
		log.Errorln("Failed to scrape Coingecko:", err)
		scraper.error = err
	}

}

func (scraper *CoingeckoScraper) Update() error {
	log.Printf("Executing CoingeckoScraper update")
	layout := "2006-01-02T15:04:05.000Z"
	visited := make(map[string]string) // Id and Symbol
	tommorrow := time.Now().Add(24 * time.Hour)

	for true {
		tokenList := scraper.FetchAvailableSymbols()
		//err := nil

		presentTime := time.Now()
		timeAheadbyAday := tommorrow.Before(presentTime) //To update once a day the historical day.
		if timeAheadbyAday {
			tommorrow = presentTime.Add(24 * time.Hour)
		}

		for _, tokenSet := range tokenList {
			id := tokenSet.ID
			url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s?localization=false&developer_data=false", id)
			coinsTemp := CoinIs{}
			bodyData, err := scraper.readCoingeckoCoins(url)
			if err != nil {
				log.Println(err)
				continue
			}

			err = json.Unmarshal(bodyData, &coinsTemp)
			if err != nil {
				log.Errorln("Failed to unmarshal bodyData:", err)
				continue
			}

			if coinsTemp.Market["current_price"] == nil {
				continue
			}

			currentPrices := coinsTemp.Market["current_price"].(map[string]interface{})
			if currentPrices["usd"] == nil {
				fmt.Errorf("error parsing rate Current Price Map empty  for Token (%s): %v", id, err)
				continue
			}

			usdPrice := currentPrices["usd"].(float64)

			timeStamp, _ := time.Parse(layout, coinsTemp.LastUpdated)

			// Yesterday data
			t := time.Now().AddDate(0, 0, -1)
			seenSymbol, ok := visited[tokenSet.ID]

			if !ok || timeAheadbyAday {
				url2 := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%02d-%02d-%d",
					tokenSet.ID, t.Day(), t.Month(), t.Year())
				yesterdayData, err := scraper.readCoingeckoCoins(url2)
				if err != nil {
					log.Errorln("Failed to  scrape yesterdayData:", err)
					continue
				}
				yesterdayHistory := CoinIs{}
				err = json.Unmarshal(yesterdayData, &yesterdayHistory)
				if err != nil {
					log.Errorln("Failed to unmarshal yesterdayHistory :", tokenSet.ID, err)
					continue
				}
				if yesterdayHistory.Market["current_price"] == nil {
					log.Errorln("Current_price Map empty for yesterdayHistory  Token: ", tokenSet.ID)
					continue
				}

				tradePrices := yesterdayHistory.Market["current_price"].(map[string]interface{})
				if tradePrices["usd"] == nil {
					log.Errorln("Yesterday current_price Map empty Token:", tokenSet.ID)
					continue
				}
				yesterdayPriceUSD := tradePrices["usd"].(float64)
				tradeVolumes := yesterdayHistory.Market["total_volume"].(map[string]interface{})
				if tradeVolumes["usd"] == nil {
					log.Errorln("Yesterday Volume Map empty Token:", tokenSet.ID)
					continue
				}
				yesterdayvolumeUSD := tradeVolumes["usd"].(float64)

				foreignQuotation := models.ForeignQuotation{
					Symbol:             coinsTemp.Symbol,
					Name:               coinsTemp.Name,
					Price:              usdPrice,
					PriceYesterday:     &yesterdayPriceUSD,
					VolumeYesterdayUSD: &yesterdayvolumeUSD,
					Source:             "Coingecko",
					Time:               timeStamp,
					ITIN:               "",
				}
				scraper.datastore.SaveForeignQuotationInflux(foreignQuotation)

			} else {

				backVal, err := scraper.datastore.GetForeignQuotationInflux(seenSymbol)

				if err != nil {
					log.Errorln("Failed to get information from Influxdb:", err)
					return err
				}

				yesterdayPriceUSD := backVal.PriceYesterday
				yesterdayvolumeUSD := backVal.VolumeYesterdayUSD

				foreignQuotation := models.ForeignQuotation{
					Symbol:             coinsTemp.Symbol,
					Name:               coinsTemp.Name,
					Price:              usdPrice,
					PriceYesterday:     yesterdayPriceUSD,
					VolumeYesterdayUSD: yesterdayvolumeUSD,
					Source:             "Coingecko",
					Time:               timeStamp,
					ITIN:               "",
				}

				scraper.datastore.SaveForeignQuotationInflux(foreignQuotation)

			}

		}
		time.Sleep(1 * time.Hour)

	}
	return nil

}

func (scraper *CoingeckoScraper) readCoingeckoCoins(url string) ([]byte, error) {

	response, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return []byte{}, fmt.Errorf("HTTP Response Error %d\n", response.StatusCode)
	}

	// Read the response body
	Bodydata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	//we dont have list of pairs, to get poairs we will get all aavailable assets and create pairs
	// Get available assets
	return Bodydata, nil

}

func (scraper *CoingeckoScraper) FetchAvailableSymbols() CoinIds {
	url := "https://api.coingecko.com/api/v3/coins/list"
	coins, _ := scraper.readCoingeckoCoins(url)
	tokens := CoinIds{}
	err := json.Unmarshal(coins, &tokens)

	if err != nil {
		log.Println(err)
		return tokens
	}
	return tokens
}
