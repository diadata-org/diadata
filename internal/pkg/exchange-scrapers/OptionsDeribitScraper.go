package scrapers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"net/url"

	"github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	zap "go.uber.org/zap"
	"github.com/gorilla/websocket"
)

// DeribitOptionsScraper - used to maintain the order book and save it every x seconds
type DeribitOptionsScraper struct {
	deribitScraper     *DeribitScraper
	//collectMetaEvery int16 // minutes, polls the rest api to see if there are any new options and saves any new options to the same file
	optionsWaitGroup   *sync.WaitGroup
	ScraperIsRunning   bool
	ScraperIsRunningMu sync.Mutex
}

type AllDeribitOptionsScrapers struct {
	Scrapers         []*DeribitOptionsScraper
	collectMetaEvery int16
	ds               *models.DB
	accessKey        string
	accessSecret     string
	owg              *sync.WaitGroup
	WsConnection       *websocket.Conn
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
	u := url.URL{Scheme: "wss", Host: "www.deribit.com", Path: "/ws/api/v2/"}
	log.Debugf("connecting to [%s]", u.String())
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Errorf("could not dial the websocket: %s", err)
		time.Sleep(time.Duration(retryIn) * time.Second)
		return result
	}
	result.WsConnection = ws
	result.collectMetaEvery = 6 // hours
	for _, market := range markets {
		newScraper := NewDeribitOptionsScraper(ds, owg, market, accessKey, accessSecret, ws)
		result.Scrapers = append(result.Scrapers, &newScraper)
	}
	result.ds = ds
	result.accessKey = accessKey
	result.accessSecret = accessSecret
	result.owg = owg
	return result
}

// NewDeribitOptionsScraper - returns an instance of an options scraper.
func NewDeribitOptionsScraper(ds *models.DB, owg *sync.WaitGroup, market string, accessKey string, accessSecret string, ws *websocket.Conn) DeribitOptionsScraper {
	wg := sync.WaitGroup{}
	logger := zap.NewExample().Sugar()
	optionsScraper := DeribitOptionsScraper{}
	defer logger.Sync()

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
		WsConnection:      ws,
	}

	//owg := sync.WaitGroup{}

	optionsScraper.deribitScraper = &scraper
	optionsScraper.optionsWaitGroup = owg
	return optionsScraper
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
	start := !s.ScraperIsRunning
	s.ScraperIsRunning = true
	s.ScraperIsRunningMu.Unlock()
	if start {
		s.deribitScraper.Scrape(market) // this will work forever and it will close the scraper inside of it
		s.ScraperIsRunningMu.Lock()
		s.ScraperIsRunning = false
		s.ScraperIsRunningMu.Unlock()
	}
	s.optionsWaitGroup.Wait()
}

func (s *AllDeribitOptionsScrapers) GetMetas() {
	tick := time.NewTicker(time.Duration(s.collectMetaEvery) * time.Hour)
	defer tick.Stop()
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
	for {
		for _, scraper := range s.Scrapers {
			scraper.optionsWaitGroup.Add(1)
			go scraper.Scrape(scraper.deribitScraper.Markets[0])
		}
		time.Sleep(10 * time.Second)
	}
}

func (s *AllDeribitOptionsScrapers) AddMarket(market string) {
	for _, scraper := range s.Scrapers {
		if scraper.deribitScraper.Markets[0] == market {
			return
		}
	}
	newScraper := NewDeribitOptionsScraper(s.ds, s.owg, market, s.accessKey, s.accessSecret, s.WsConnection)
	s.Scrapers = append(s.Scrapers, &newScraper)
	return
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

func (s *AllDeribitOptionsScrapers) RefreshMetas(currency string) error {
	tick := time.NewTicker(time.Duration(s.collectMetaEvery) * time.Hour)
	defer tick.Stop()
	go func() {
		for {
			metas, err := s.ds.GetOptionMeta(currency)
			if err != nil {
				log.Error(err)
			}
			for _, meta := range metas {
				s.AddMarket(meta.InstrumentName)
			}
			select {
				case <-tick.C:
			}
		}
	}()
	return nil
}

func (s *AllDeribitOptionsScrapers) Close() {
	for _, scraper := range s.Scrapers {
		scraper.ScraperClose(scraper.deribitScraper.Markets[0], s.WsConnection)
	}
	err := s.WsConnection.Close()
	if err != nil {
		log.Error(err)
	}
}
