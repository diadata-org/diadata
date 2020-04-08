package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	utils "github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	zap "go.uber.org/zap"
)

const scrapeDataSaveLocationDeribit = ""

type deribitRefreshMessage struct {
	Result struct {
		RefreshToken string `json:"refresh_token"`
	} `json:"result"`
}

type deribitErrorMessage struct {
	Error struct {
		Message string `json:"message"`
		Code    int64  `json:"code"`
	} `json:"error"`
}

type ParsedDeribitResponse struct {
	Jsonrpc string                      `json:"jsonrpc"`
	Method  string                      `json:"method"`
	Params  ParsedDeribitResponseParams `json:"params"`
}

type ParsedDeribitResponseParams struct {
	Channel string                            `json:"channel"`
	Data    ParsedDeribitOptionOrderbookEntry `json:"data"`
}

type ParsedDeribitOptionOrderbookEntry struct {
	Timestamp      int64       `json:"timestamp"`
	InstrumentName string      `json:"instrument_name"`
	ChangeId       int64       `json:"change_id"`
	Bids           [][]float64 `json:"bids"`
	Asks           [][]float64 `json:"asks"`
}

// NewDeribitFuturesScraper - creates a deribit futures scraper for you for the markets that you supply. Some of the markets available are: "BTC-PERPETUAL" and "ETH-PERPETUAL".
func NewDeribitFuturesScraper(markets []string, accessKey string, accessSecret string) FuturesScraper {
	wg := sync.WaitGroup{}
	logger := zap.NewExample().Sugar() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	var scraper DeribitScraper = DeribitScraper{
		WaitGroup: &wg,
		Markets:   markets, // e.g. []string{"BTC-PERPETUAL", "ETH-PERPETUAL"}
		Logger:    logger,

		AccessKey:    accessKey,
		AccessSecret: accessSecret,

		// expiry is 900 seconds
		RefreshTokenEvery: 800,
		MarketKind:        DeribitFuture, // DO NOT change this.
	}

	return &scraper
}

func (s *DeribitScraper) send(message *map[string]interface{}, websocketConn *websocket.Conn) error {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		return err
	}
	return nil
}

// ScraperClose - responsible for closing out the scraper for a market
func (s *DeribitScraper) ScraperClose(market string, websocketConnection interface{}) error {
	switch c := websocketConnection.(type) {
	case *websocket.Conn:
		err := c.WriteJSON(map[string]string{"op": "unsubscribe", "channel": "trades", "market": market})
		if err != nil {
			return err
		}
		log.Infof("gracefully shutdown deribit scraper on market: %s", market)
		time.Sleep(time.Duration(retryIn) * time.Second)
		return nil
	default:
		return fmt.Errorf("unknown connection type, expected gorilla/websocket, got: %T", c)
	}
}

// Scrape starts a websocket scraper for market
func (s *DeribitScraper) Scrape(market string) {
	err := s.validateMarket(market, s.MarketKind)
	if err != nil {
		return
	}
	s.validateRefreshEveryToken()

	optionRequest := &map[string]interface{}{
		"method": "public/subscribe",
		"params": &map[string]interface{}{
			"channels": []string{"book." + market + ".none.1.100ms"}, // will give us orderbook snapshots every 100 ms
		},
		"jsonrpc": "2.0",
		"id":      0,
	}
	futureRequest := &map[string]interface{}{
		"method": "private/subscribe",
		"params": &map[string]interface{}{
			"channels": []string{"trades." + market + ".raw"},
		},
		"jsonrpc": "2.0",
		"id":      0,
	}

	switch s.MarketKind {
	case DeribitFuture:
		err = s.send(futureRequest, s.WsConnection)
	case DeribitOption:
		err = s.send(optionRequest, s.WsConnection)
	default:
		panic("unknown market kind")
	}

	if err != nil {
		log.Errorf("could not send ws message. restarting the websocket, err: %s", err)
		return
	}
}

// ScrapeMarkets - will scrape the markets specified during instantiation
func (s *DeribitScraper) ScrapeMarkets() {
	for _, market := range s.Markets {
		s.WaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.WaitGroup.Wait()
}

// marketKind can be "future" or "option"
func (s *DeribitScraper) validateMarket(market string, marketKind DeribitScraperKind) error {
	allFuturesMarketsDeribit, err := allDeribitMarketsOfKind(marketKind)
	if err != nil {
		return err
	}
	containsMarket := utils.Contains(&allFuturesMarketsDeribit, market)
	if !containsMarket {
		return errors.New(market + " market is unavailable")
	}
	return nil
}

func (s *DeribitScraper) validateRefreshEveryToken() {
	if s.RefreshTokenEvery >= 900 {
		panic("When you use https://testapp.deribit.com/api_console, you will see that upon a successful authentication, the response will include expiresIn field. Which is set at 900. This means that the token we generated is only valid for 900 seconds, and we have to refresh it before then.")
	}
}

func deribitMarkets(market string, marketKind DeribitScraperKind) ([]string, error) {
	if market != "BTC" && market != "ETH" {
		panic("unsupported market. only btc & eth are supported")
	}
	resp, err := http.Get("https://www.deribit.com/api/v2/public/get_instruments?currency=" + market)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	decodedMsg := deribitInstruments{}
	err = json.Unmarshal(body, &decodedMsg)
	if err != nil {
		return nil, err
	}
	allMarkets := []string{}
	for _, market := range decodedMsg.Result {
		switch marketKind {
		case DeribitFuture:
			if market.Kind == "future" {
				allMarkets = append(allMarkets, market.InstrumentName)
			}
		case DeribitOption:
			if market.Kind == "option" {
				allMarkets = append(allMarkets, market.InstrumentName)
			}
		default:
			panic("unknown market kind")
		}
	}
	return allMarkets, nil
}

func allDeribitMarketsOfKind(marketKind DeribitScraperKind) ([]string, error) {
	BTCMarkets, err := deribitMarkets("BTC", marketKind)
	if err != nil {
		return nil, fmt.Errorf("could not fetch btc futures markets, err: %s", err)
	}
	ETHMarkets, err := deribitMarkets("ETH", marketKind)
	if err != nil {
		return nil, fmt.Errorf("could not fetch eth futures markets, err: %s", err)
	}
	return append(BTCMarkets, ETHMarkets...), nil
}
