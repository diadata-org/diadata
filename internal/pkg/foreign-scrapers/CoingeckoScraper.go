package foreignscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type CoinIds []struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	CurrentPrice float64 `json:"current_price"`
	MarketCap    float64 `json:"market_cap"`
	LastUpdated  string  `json: "last_updated"`
}

type CoinIs struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	Symbol      string                   `json:"symbol"`
	LastUpdated string                   `json: "last_updated"`
	Tickers     []map[string]interface{} `json:"tickers"`
	Market      map[string]interface{}   `json:"market_data"`
}

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
	setFetchAvailaSymbol := 0
	var tokenList CoinIds
	afterThisCountSleep := 98
	var err error
	for true {
		if setFetchAvailaSymbol == 0 {
			tokenList, err = scraper.FetchAvailableSymbols()
			if err != nil {
				log.Println(err)
				time.Sleep(10 * time.Minute)
			}
			setFetchAvailaSymbol = 1
			time.Sleep(10 * time.Second)
		}

		presentTime := time.Now()
		timeAheadbyAday := tommorrow.Before(presentTime) //To update once a day the historical day.
		if timeAheadbyAday {
			tommorrow = presentTime.Add(24 * time.Hour)
			setFetchAvailaSymbol = 0
		}
		time.Sleep(40 * time.Second)

		for k, tokenSet := range tokenList {
			t := time.Now().AddDate(0, 0, -1)
			seenSymbol, ok := visited[tokenSet.ID]
			timeStamp, _ := time.Parse(layout, tokenSet.LastUpdated)

			if !ok || timeAheadbyAday {
				log.Printf("Fetch new data directly from Coingecko")

				if (k+1)%afterThisCountSleep == 0 {
					time.Sleep(1 * time.Minute)
				}
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
					Symbol:             tokenSet.Symbol,
					Name:               tokenSet.Name,
					Price:              tokenSet.CurrentPrice,
					PriceYesterday:     &yesterdayPriceUSD,
					VolumeYesterdayUSD: &yesterdayvolumeUSD,
					Source:             "Coingecko",
					Time:               timeStamp,
					ITIN:               "",
				}
				scraper.datastore.SaveForeignQuotationInflux(foreignQuotation)

			} else {
				log.Printf("Fetch from DB")

				backVal, err := scraper.datastore.GetForeignQuotationInflux(seenSymbol)

				if err != nil {
					log.Errorln("Failed to get information from Influxdb:", err)
					return err
				}

				yesterdayPriceUSD := backVal.PriceYesterday
				yesterdayvolumeUSD := backVal.VolumeYesterdayUSD

				foreignQuotation := models.ForeignQuotation{
					Symbol:             tokenSet.Symbol,
					Name:               tokenSet.Name,
					Price:              tokenSet.CurrentPrice,
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
		fmt.Errorf("HTTP Response Error %d\n", response.StatusCode)
		log.Printf("Because of the above error message system would sleep for 1 minute and try again")
		countWait := 1
		time.Sleep(1 * time.Minute)
		for true {
			fmt.Println("Failed: Downloading data and Read")
			if countWait == 5 {
				break
			}
			countWait++

			response, err = http.Get(url)

			if err != nil {
				return []byte{}, err
			}

			defer response.Body.Close()

			if response.StatusCode == 200 {
				break
			}
			time.Sleep(1 * time.Minute)
		}
		if response.StatusCode != 200 {
			log.Printf("Very serious issue , looks like site is down")
			return []byte{}, fmt.Errorf("HTTP Response Error %d\n", response.StatusCode)

		}

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

func (scraper *CoingeckoScraper) FetchAvailableSymbols() (CoinIds, error) {
	var s CoinIds
	pageContentLength := 100
	page := 1
	layout := "2006-01-02T15:04:05.000Z"
	added := make(map[string]bool) // Id and Symbol
	for true {

		url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=%d&sparkline=false", page)
		page++
		XMLdata, err := scraper.readCoingeckoCoins(url)
		if err != nil {
			fmt.Println("Failed: Downloading data and Read")
		}

		confirmTemp := CoinIds{}
		jsonErr := json.Unmarshal(XMLdata, &confirmTemp)
		if jsonErr != nil {
			fmt.Println(jsonErr)
			return s, jsonErr
		}

		signalTrigger := 0
		for _, i := range confirmTemp {
			if i.CurrentPrice == 0 || i.MarketCap == 0 {
				signalTrigger++
				continue
			}

			timeStamp, _ := time.Parse(layout, i.LastUpdated)
			lastTradeAcceptableRange := time.Now().AddDate(0, -4, 0)

			if lastTradeAcceptableRange.Before(timeStamp) {
				continue
			}
			if strings.Contains(i.ID, "3X-") || strings.Contains(i.ID, "long") {
				continue
			}
			if strings.Contains(i.ID, "futures") || strings.Contains(i.ID, "short") {
				continue
			}
			_, ok := added[i.Symbol]
			if !ok {
				s = append(s, i)
				added[i.Symbol] = true
			}
			signalTrigger = 0

		}
		if signalTrigger == 100 {
			break
		}

		if len(confirmTemp) < pageContentLength {
			break
		}

	}
	return s, nil
}
