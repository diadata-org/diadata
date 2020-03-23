package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"net/url"
	/*"os"
	"os/signal"
	"strings"*/
	"sync"
	//"syscall"
	"time"
	"errors"

	"github.com/gorilla/websocket"
	utils "github.com/diadata-org/diadata/internal/pkg/scraper-utils"
	//"github.com/diadata-org/diadata/pkg/dia"
	zap "go.uber.org/zap"
	log "github.com/sirupsen/logrus"
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
	//log.Debugf("sent message [%s]: %s", market, message)
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
		/*err = c.Close()
		if err != nil {
			return err
		}*/
		log.Infof("gracefully shutdown deribit scraper on market: %s", market)
		time.Sleep(time.Duration(retryIn) * time.Second)
		return nil
	default:
		return fmt.Errorf("unknown connection type, expected gorilla/websocket, got: %T", c)
	}
}

/*// Authenticate - authenticates with your access key and access secret to retrieve the trade details
func (s *DeribitScraper) Authenticate(market string, websocketConnection interface{}) error {
	switch c := websocketConnection.(type) {
	case *websocket.Conn:
		err := s.send(&map[string]interface{}{
			"method": "public/auth",
			"params": &map[string]string{
				"grant_type": "client_credentials",
				// "scope":         "session:apiconsole-bq7yiuifb88",
				"client_id":     s.AccessKey,
				"client_secret": s.AccessSecret,
			},
			"jsonrpc": "2.0",
		}, market, c)
		return err
	default:
		return fmt.Errorf("unknown connection type, expected gorilla/websocket, got: %T", c)
	}
}
*/

// Scrape starts a websocket scraper for market
func (s *DeribitScraper) Scrape(market string) {
	err := s.validateMarket(market, s.MarketKind)
	if err != nil {
		return
	}
	s.validateRefreshEveryToken()

	/*
	// this block is for listening to sigterms and interupts
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	userCancelled := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		userCancelled <- true
	}()
	*/

	//for {
		// immediately invoked function expression for easy clenup with defer
	//	func() {
			/*refreshToken := ""
			failedToRefreshToken := make(chan interface{})*/
			/*u := url.URL{Scheme: "wss", Host: "www.deribit.com", Path: "/ws/api/v2/"}
			log.Debugf("connecting to [%s], market: [%s]", u.String(), market)
			ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Errorf("could not dial the websocket: %s", err)
				time.Sleep(time.Duration(retryIn) * time.Second)
				return
			}
			defer s.ScraperClose(market, ws)*/
			// 1. authenticate
/*			err = s.Authenticate(market, s.WsConnection)
			if err != nil {
				log.Errorf("could not authenticate. retrying, err: %s", err)
				return
			}*/
			//time.Sleep(time.Second)
			// 2. subscribe to channel depending on the market kind. for futures, collect trades, for options, collect ob
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
			// 3. refresh the token more often than 900 seconds
			/*tick := time.NewTicker(time.Duration(s.RefreshTokenEvery) * time.Second) // every RefreshTokenEvery seconds we have to refresh token
			defer tick.Stop()
			// we require a separate goroutine for ticker, so that we can refresh our access token everyRefreshToken seconds
			go func() {
				for {
					select {
					case <-tick.C:
						isRefreshed, err := s.refreshToken(refreshToken, market, s.WsConnection)
						if err != nil {
							close(failedToRefreshToken)
							time.Sleep(time.Duration(60) * time.Minute) // something very long
						}
						maxRetryAttempts := 5
						if !isRefreshed {
							for i := 1; i < maxRetryAttempts; i++ {
								isRefreshed, err := s.refreshToken(refreshToken, market, s.WsConnection)
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
			}()*/
			/*for {
				select {
				case <-userCancelled:
					log.Infof("received interrupt, gracefully shutting down")
					s.ScraperClose(market, s.WsConnection)
					os.Exit(0)
				case <-failedToRefreshToken:
					log.Errorf("failed to refresh token numerous times. restarting the scraper")
					time.Sleep(time.Duration(retryIn) * time.Second)
					return
				}
			}*/
			/*default:
					_, message, err := s.WsConnection.ReadMessage() // this code is blocking. that is why we need big sleep time in the refreshToken goroutine
					if err != nil {
						log.Errorf("problem reading deribit on [%s], err: %s", market, err)
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
						log.Debugf("obtained a new refresh token on [%s], updating '%s'", market, decodedMsg.Result.RefreshToken)
						refreshToken = decodedMsg.Result.RefreshToken
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
							InstrumentName:  parsedResult.Params.Data.InstrumentName,
							// Caution: The response contains the Unix timestamp in Milliseconds
							ObservationTime: time.Unix(parsedResult.Params.Data.Timestamp / 1000, (parsedResult.Params.Data.Timestamp % 1000) * 1e6),
							AskPrice:        parsedResult.Params.Data.Asks[0][0],
							BidPrice:        parsedResult.Params.Data.Bids[0][0],
							AskSize:         parsedResult.Params.Data.Asks[0][1],
							BidSize:         parsedResult.Params.Data.Bids[0][1],
						}
						err := s.DataStore.SaveOptionOrderbookDatumInflux(orderbookEntry)
						if err != nil {
							log.Errorf("Error writing into influxdb: %s", err)
							return
						}
						log.Debug("Write msg to db: ", orderbookEntry)
					} else {
						// only save the messages if it is a control message
						//log.Debugf("saving new message on [%s]", market)
						log.Debugf(strMessage)
					}*/
			//	}
		//	}
		//}()
	//}
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
