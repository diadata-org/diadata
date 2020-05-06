package scrapers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// DeribitOptionsScraper - used to maintain the order book and save it every x seconds
type DeribitOptionsScraper struct {
	deribitScraper     *DeribitScraper
	optionsWaitGroup   *sync.WaitGroup
	ScraperIsRunning   bool
	ScraperIsRunningMu sync.Mutex
}

type AllDeribitOptionsScrapers struct {
	Scrapers          []*DeribitOptionsScraper
	collectMetaEvery  int16
	ds                *models.DB
	accessKey         string
	accessSecret      string
	owg               *sync.WaitGroup
	WsConnection      *websocket.Conn
	refreshToken      string
	RefreshTokenEvery int16
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

type deribitInstruments struct {
	Result []deribitInstrument `json:"result"`
}

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
	result.RefreshTokenEvery = 800
	return result
}

// NewDeribitOptionsScraper - returns an instance of an options scraper.
func NewDeribitOptionsScraper(ds *models.DB, owg *sync.WaitGroup, market string, accessKey string, accessSecret string, ws *websocket.Conn) DeribitOptionsScraper {
	wg := sync.WaitGroup{}
	logger := zap.NewExample().Sugar()
	optionsScraper := DeribitOptionsScraper{}
	defer logger.Sync()

	var scraper = DeribitScraper{
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

	optionsScraper.deribitScraper = &scraper
	optionsScraper.optionsWaitGroup = owg
	return optionsScraper
}

func (s *AllDeribitOptionsScrapers) send(message *map[string]interface{}, websocketConn *websocket.Conn) error {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		return err
	}
	return nil
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

// ScrapeMarkets - scrapes all the options markets
func (s *AllDeribitOptionsScrapers) ScrapeMarkets() {
	err := s.Authenticate(s.WsConnection)
	if err != nil {
		log.Errorf("could not authenticate. retrying, err: %s", err)
		return
	}
	log.Info("Authentication success")
	go s.refreshWsToken()
	go func() {
		for {
			s.handleWsMessage()
		}
		time.Sleep(30 * time.Second)
	}()
	for {
		for _, scraper := range s.Scrapers {
			scraper.optionsWaitGroup.Add(1)
			time.Sleep(10 * time.Second)
			go scraper.Scrape(scraper.deribitScraper.Markets[0])
		}
	}
}

func (s *AllDeribitOptionsScrapers) handleWsMessage() {
	_, message, err := s.WsConnection.ReadMessage() // this code is blocking. that is why we need big sleep time in the refreshToken goroutine
	if err != nil {
		log.Errorf("problem reading deribit, err: %s", err)
		return
	}
	strMessage := string(message)
	log.Debugf("received new message: %v", strMessage)
	// check if the received message contains the refresh_token json key
	if strings.Contains(strMessage, "refresh_token") {
		decodedMsg := deribitRefreshMessage{}
		err = json.Unmarshal(message, &decodedMsg)
		if err != nil {
			log.Errorf("problem unmarshalling the message: %s, err: %s", message, err)
			return
		}
		log.Debugf("obtained a new refresh token, updating '%s'", decodedMsg.Result.RefreshToken)
		s.refreshToken = decodedMsg.Result.RefreshToken
	} else if strings.Contains(strMessage, "error") {
		decodedMsg := deribitErrorMessage{}
		err = json.Unmarshal(message, &decodedMsg)
		if err != nil {
			log.Errorf("problem unmarshalling the message: %s, err: %s", message, err)
			return
		}
		if decodedMsg.Error.Message != "" {
			if decodedMsg.Error.Code == 13009 {
				panic("wrong or expired authorization token or bad signature. For example, please check scope of the token, \"connection\" scope can't be reused for other connections.")
			}
		}
	} else if strings.Contains(strMessage, `"method":"subscription"`) {
		// Magic happens here :-)
		// only save the messages if the message does not contain thre refresh_token string and no error
		//log.Debugf("saving new orderbook message on [%s]", market)
		parsedResult := ParsedDeribitResponse{}
		err = json.Unmarshal(message, &parsedResult)
		if err != nil {
			log.Errorf("problem unmarshalling the message: %s, err: %s", message, err)
			return
		}
		if len(parsedResult.Params.Data.Bids) == 0 ||
			len(parsedResult.Params.Data.Asks) == 0 {
			log.Errorf("No bid or ask in message %s", message)
			return
		}
		orderbookEntry := dia.OptionOrderbookDatum{
			InstrumentName: parsedResult.Params.Data.InstrumentName,
			// Caution: The response contains the Unix timestamp in Milliseconds
			ObservationTime: time.Unix(parsedResult.Params.Data.Timestamp/1000, (parsedResult.Params.Data.Timestamp%1000)*1e6),
			AskPrice:        parsedResult.Params.Data.Asks[0][0],
			BidPrice:        parsedResult.Params.Data.Bids[0][0],
			AskSize:         parsedResult.Params.Data.Asks[0][1],
			BidSize:         parsedResult.Params.Data.Bids[0][1],
		}
		err := s.ds.SaveOptionOrderbookDatumInflux(orderbookEntry)
		if err != nil {
			log.Errorf("Error writing into influxdb: %s", err)
			return
		}
		log.Debug("Write msg to db: ", orderbookEntry)
	} else {
		// only save the messages if it is a control message
		//log.Debugf("saving new message on [%s]", market)
		log.Debugf(strMessage)
	}
}

// Authenticate - authenticates with your access key and access secret to retrieve the trade details
func (s *AllDeribitOptionsScrapers) Authenticate(websocketConnection interface{}) error {
	switch c := websocketConnection.(type) {
	case *websocket.Conn:
		err := s.send(&map[string]interface{}{
			"method": "public/auth",
			"params": &map[string]string{
				"grant_type":    "client_credentials",
				"client_id":     s.accessKey,
				"client_secret": s.accessSecret,
			},
			"jsonrpc": "2.0",
		}, c)
		return err
	default:
		return fmt.Errorf("unknown connection type, expected gorilla/websocket, got: %T", c)
	}
}

// when we authenticate, we get back a refresh token that we use to keep alive our websocket connection
func (s *AllDeribitOptionsScrapers) handleRefreshToken(previousToken string, websocketConn *websocket.Conn) (bool, error) {
	if previousToken == "" {
		return false, nil
	}
	// refresh
	err := s.send(&map[string]interface{}{
		"method": "public/auth",
		"params": &map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": previousToken,
		},
		"jsonrpc": "2.0",
	}, websocketConn)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *AllDeribitOptionsScrapers) refreshWsToken() {
	failedToRefreshToken := make(chan interface{})
	// 3. refresh the token more often than 900 seconds
	tick := time.NewTicker(time.Duration(s.RefreshTokenEvery) * time.Second) // every RefreshTokenEvery seconds we have to refresh token
	defer tick.Stop()
	// we require a separate goroutine for ticker, so that we can refresh our access token everyRefreshToken seconds
	for {
		select {
		case <-tick.C:
			isRefreshed, err := s.handleRefreshToken(s.refreshToken, s.WsConnection)
			if err != nil {
				close(failedToRefreshToken)
				time.Sleep(time.Duration(60) * time.Minute) // something very long
			}
			maxRetryAttempts := 5
			if !isRefreshed {
				for i := 1; i < maxRetryAttempts; i++ {
					isRefreshed, err := s.handleRefreshToken(s.refreshToken, s.WsConnection)
					if isRefreshed {
						break
					}
					if err != nil {
						close(failedToRefreshToken)
						time.Sleep(time.Duration(60) * time.Minute) // something very long
					}
				}
			}
		}
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
		if optionMeta.InstrumentName == option.InstrumentName {
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
	body, err := utils.GetRequest("https://www.deribit.com/api/v2/public/get_instruments?currency=" + market)
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
			optionMeta := dia.OptionMeta{
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
