package foreignscrapers

// Please implement the scraping of coingecko quotations here.

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"net/http"
	"io/ioutil"
	"time"

	"github.com/diadata-org/diadata/pkg/models"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	utils "github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)



type CoinIds []struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
}

type CoinIs struct{
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LastUpdated string `json: "last_updated"`
	Tickers  []map[string]interface{} `json:"tickers"`
	Market  map[string]interface{} `json:"market_data"`
}


var _coingeckourl string = "https://api.coingecko.com/api/v3"

type CoingeckoScraper struct {
	exchangeName string
	// signaling channels for session initialization and finishing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
    datastore    models.Datastore
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*CoingeckoPairScraper
	chanTrades   chan *dia.Trade
}

func NewCoingeckoScraper(datastore models.Datastore) *CoingeckoScraper {
	s := &CoingeckoScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*CoingeckoPairScraper),
		ticker:       time.NewTicker(refreshDelay),
		datastore:    datastore,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	go s.mainLoop()
	return s
}


// mainLoop runs in a goroutine until channel s is closed.
func (scraper  *CoingeckoScraper) mainLoop() {
	for {
		select {
		case <-s.ticker.C:
			s.Update()
		case <-s.shutdown: // user requested shutdown
			log.Println("CoingeckoScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

func (scraper *CoingeckoScraper) Update() error {
	log.Printf("Executing CoingeckoScraper update")
	scraper.run = true
	layout := "2006-01-02T15:04:05"

	quotation := &models.Quotation{
		[]models.ForeignQuotation{},
	}

	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("Coingecko: No data to scrape")
			log.Error(scraper.error.Error())
			break
		}
	
		for key, pairScraper := range scraper.pairScrapers {
	
			pairs := pair.split()
			sPairs := strings.Split(key, "--")
			url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s?localization=false&developer_data=false", sPair[1])
			coinsTemp := CoinIs{}
			bodyData, err := readCoingeckoCoins(url)
			err = json.Unmarshal(bodyData, &scoinsTemp)
	
			if err != nil {
				log.Println(err)
				return err
			}
	
			currentPrices := confirmTemp1.Market["current_price"].(map[string]interface{})
			usdPrice, err := strconv.ParseFloat(currentPrices["usd"], 64)
			if err != nil {
				return fmt.Errorf("error parsing rate %$ %v", "Price", err)
			}
			time, _ := time.Parse(layout, confirmTemp1.LastUpdated)
			
			t := &dia.Trade{
				Pair:   pair,
				Symbol: pair,
				Price:  usdPrice,
				Volume: 0,
				Time:   time,
				Source: "Coingecko",
			}

			scraper.chanTrades <- t
	
			for _, ticker := range coinsTemp.Tickers{
	
				timeStamp, _ := time.Parse(layout, tricker["TimeStamp"])
				f := models.ForeignQuotation{
					Symbol: ticker["base"]
					Name               string
					Price: ticker["last"].(float64)
					PriceYesterday: ticker["last"].(float64)
					VolumeYesterdayUSD: ticker["volume"].(float64)
					Source: ticker["trade_url"]
					Time: timeStamp
					ITIN: ""
				}
				
				quotation.Foreign = append(quotation.Foreign, f)	
			}
			s.datastore.SetCurrencyChange(quotation)
	}
	
}
// Channel returns a channel that can be used to receive trades/pricing information
func (scraper  *CoingeckoScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (scraper  *CoingeckoScraper) Close() error {
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


func (scraper *CoingeckoScraper) createPairs() (pairs map[string]string) {
	pairs = make(map[string]string)
	url := "https://api.coingecko.com/api/v3/coins/list"
	coins, _ := scraper.readCoingeckoCoins(url)
	coinsIdTemp := CoinIds{}
	err = json.Unmarshal(coins, &scoinsIdTemp)

	if err != nil {
		log.Println(err)
		return pairs
	}
	for _, token1 := range tokens {
		
		pair := token1.ID
		pairs[pair] = token1.Symbol + "--" token1.ID
	
	}
	return pairs
}

func (scraper *CoingeckoScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	pairmap := scraper.createPairs()
	for k, v := range pairmap {
		idSymbol := String.Split(v , "--")
		pairs = append(pairs, dia.Pair{
			Symbol:     idSymbol[0] ,
			ForeignName: v,
			Exchange:    "",
		})

	}
	return
}


func (scraper *CoingeckoScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("coingeckoScraper is closed")
	}

	pairScraper := &CoingeckoPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}

func (scraper *CoingeckoScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *CoingeckoScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type CoingeckoPairScraper struct {
	parent *CoingeckoScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *CoingeckoPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *CoingeckoPairScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *CoingeckoPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *CoingeckoPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}

