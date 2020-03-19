package scrapers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	zap "go.uber.org/zap"
)

// DeribitOptionsScraper - used to maintain the order book and save it every x seconds
type DeribitOptionsScraper struct {
	deribitScraper    DeribitScraper
	//collectMetaEvery int16 // minutes, polls the rest api to see if there are any new options and saves any new options to the same file
	optionsWaitGroup  *sync.WaitGroup
	scraperIsRunning   bool
	ScraperIsRunningMu sync.Mutex
}

type AllDeribitOptionsScrapers struct {
	Scrapers         []DeribitOptionsScraper
	collectMetaEvery int16
	ds               *models.DB
}

type deribitInstrument struct {
	InstrumentName      string  `json:"instrument_name"`
	Kind                string  `json:"kind"`
	TickSize            float64 `json:"tick_size"`
	TakerCommission     float64 `json:"taker_commission"`
	Strike              float64 `json:"strike"`
	SettlementPeriod    string  `json:"settlement_period"`
	QuoteCurrency       string  `json:"quote_currency"`
	OptionType          string  `json:"option_type"`
	MinTradeAmount      float64 `json:"min_trade_amount"`
	MakerCommission     float64 `json:"maker_commission"`
	IsActive            bool    `json:"is_active"`
	ExpirationTimestamp int64   `json:"expiration_timestamp"`
	CreationTimestamp   int64   `json:"creation_timestamp"`
	ContractSize        float64 `json:"contract_size"`
	BaseCurrency        string  `json:"base_currency"`
}

type deribitOrderbookDatum struct {
	Params struct {
		Data struct {
			Timestamp      int64         `json:"timestamp"`
			InstrumentName string        `json:"instrument_name"`
			ChangeID       int64         `json:"change_id"`
			Bids           [1][2]float64 `json:"bids"`
			Asks           [1][2]float64 `json:"asks"`
		} `json:"data"`
	} `json:"params"`
}

type deribitInstruments struct {
	Result []deribitInstrument `json:"result"`
}

const deribitOptionsMetaFilename string = "deribit-options-meta.txt"

func NewAllDeribitOptionsScrapers(owg *sync.WaitGroup, markets []string, accessKey string, accessSecret string) AllDeribitOptionsScrapers {
	result := AllDeribitOptionsScrapers{}
	ds, err := models.NewDataStore()
	if err != nil {
		return result
	}
	result.collectMetaEvery = 6 // hours
	for _, market := range markets {
		newScraper := NewDeribitOptionsScraper(owg, market, accessKey, accessSecret)
		result.Scrapers = append(result.Scrapers, *newScraper)
	}
	result.ds = ds
	return result
}

// NewDeribitOptionsScraper - returns an instance of an options scraper.
func NewDeribitOptionsScraper(owg *sync.WaitGroup, market string, accessKey string, accessSecret string) *DeribitOptionsScraper {
	wg := sync.WaitGroup{}
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	ds, err := models.NewDataStore()
	if err != nil {
		logger.Errorf("Could not create datastore: %s", err)
		return nil
	}

	var scraper DeribitScraper = DeribitScraper{
		WaitGroup: &wg,
		Markets:   []string{market}, // e.g. []string{"BTC-20DEC19-5750-C", "BTC-20DEC19-7500-P"}
		Logger:    logger,
		DataStore: ds,

		AccessKey:    accessKey,
		AccessSecret: accessSecret,

		// expiry is 900 seconds
		RefreshTokenEvery: 800,
		MarketKind:        DeribitOption, // DO NOT change this.
	}

	//owg := sync.WaitGroup{}

	var optionsScraper DeribitOptionsScraper = DeribitOptionsScraper{
		deribitScraper:   scraper,
		//collectMetaEvery: 6, // hours
		optionsWaitGroup: owg,
	}

	return &optionsScraper
}

// Authenticate - authenticates
func (s *DeribitOptionsScraper) Authenticate(market string, websocketConnection interface{}) error {
	return s.deribitScraper.Authenticate(market, websocketConnection)
}

// ScraperClose - responsible for closing out the scraper for a market
func (s *DeribitOptionsScraper) ScraperClose(market string, websocketConnection interface{}) error {
	return s.deribitScraper.ScraperClose(market, websocketConnection)
}

// Scrape - scrapes the options markets (meta + order book data)
func (s *DeribitOptionsScraper) Scrape(market string) {
	// calling the futures scraper to collect the options order book data
	s.ScraperIsRunningMu.Lock()
	log.Info("Is ", market, " running already? ", s.scraperIsRunning)
	start := !s.scraperIsRunning
	s.scraperIsRunning = true
	s.ScraperIsRunningMu.Unlock()
	if start {
		go func() {
			log.Info("Start new scraper for ", market)
			s.deribitScraper.Scrape(market) // this will work forever and it will close the scraper inside of it
			log.Info("Scraper died for ", market)
			s.ScraperIsRunningMu.Lock()
			s.scraperIsRunning = false
			s.ScraperIsRunningMu.Unlock()
		}()
	}
	s.optionsWaitGroup.Wait()
}

func (s *AllDeribitOptionsScrapers) GetMetas() {
	tick := time.NewTicker(time.Duration(s.collectMetaEvery) * time.Hour)
	defer tick.Stop()
	log.Info("GetMetas")
	s.GetAndStoreOptionsMeta("BTC")
	s.GetAndStoreOptionsMeta("ETH")
	go func() {
		for {
			select {
			case <-tick.C:
				s.GetAndStoreOptionsMeta("BTC")
				s.GetAndStoreOptionsMeta("ETH")
			}
		}
	}()
}

// ScrapeMarkets - scrapes all the optiosn markets
//func (s *DeribitOptionsScraper) ScrapeMarkets() {
func (s *AllDeribitOptionsScrapers) ScrapeMarkets() {
	/*for _, market := range s.deribitScraper.Markets {
		s.optionsWaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.optionsWaitGroup.Wait()*/
	for {
		for _, scraper := range s.Scrapers {
			scraper.optionsWaitGroup.Add(1)
			//log.Info("Scrapers running: ", s.Scrapers)
			//log.Info("scraper: ", scraper)
			go scraper.Scrape(scraper.deribitScraper.Markets[0])
		}
		time.Sleep(10 * time.Second)
	}
}

// note, this function requires meta to be stored in a file
func (s *AllDeribitOptionsScrapers) MetaOnOptionIsAvailable(option deribitInstrument) (bool, error) {
	optionMetas, err := s.ds.GetOptionMeta(option.BaseCurrency)
	if err != nil {
		return false, err
	}
	for _, optionMeta := range optionMetas {
		if (optionMeta.InstrumentName == option.InstrumentName) {
			return true, nil
		}
	}
	return false, nil
}

// GetAndStoreOptionsMeta - will make a REST API call to obtain options meta, it then checks if we already have this data. If we do not, we save the new data.
func (s *AllDeribitOptionsScrapers) GetAndStoreOptionsMeta(market string) error {
	if market != "BTC" && market != "ETH" {
		panic("unsupported deribit market. only btc and eth are supported")
	}
	resp, err := http.Get("https://www.deribit.com/api/v2/public/get_instruments?currency=" + market)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	decodedMsg := deribitInstruments{}
	err = json.Unmarshal(body, &decodedMsg)
	if err != nil {
		return err
	}
	for _, instrument := range decodedMsg.Result {
		if instrument.Kind != "option" {
			continue
		}
		available, err := s.MetaOnOptionIsAvailable(instrument)
		if err != nil {
			return err
		}
		if !available {
			/*line, err := json.Marshal(instrument)
			if err != nil {
				return err
			}*/
			//s.deribitScraper.Logger.Debugf("new option, writing to redis. option: %s", line)
			optionType := dia.CallOption
			if instrument.OptionType == "put" {
				optionType = dia.PutOption
			}
			optionMeta := dia.OptionMeta {
				InstrumentName: instrument.InstrumentName,
				BaseCurrency:   instrument.BaseCurrency,
				ExpirationTime: time.Unix(instrument.ExpirationTimestamp/1000, (instrument.ExpirationTimestamp%1000)*1e6),
				StrikePrice:    instrument.Strike,
				OptionType:     optionType,
			}
			s.ds.SetOptionMeta(&optionMeta)
		}
	}
	return nil
}
