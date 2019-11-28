package scrapers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	utils "github.com/diadata-org/diadata/internal/pkg/scraper-utils"
	writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
)

const scrapeDataSaveLocationDeribit = ""

// DeribitFuturesScraper - scrapes the futures from the Deribit exchange
type DeribitFuturesScraper struct {
	Markets   []string
	WaitGroup *sync.WaitGroup
	Writer    writers.Writer
	Logger    *log.Logger

	// required for deribit to:
	// 1. authenticate (trades is a private channel)
	// 2. referesh the token from step 1., so that the channel isn't closed
	AccessKey    string
	AccessSecret string

	RefereshToken     string // this is the token that we will use to refresh our existing AccessToken that we receive by authenticating, note that we do not need to use this AccessToken anywhere in our case
	RefreshTokenEvery int16  // how often we refresh the token (in seconds)
}

type messageDeribit struct {
	Result struct {
		RefreshToken string `json:"refresh_token"`
	} `json:"result"`
}

type messageGETDeribit struct {
	Result []struct {
		InstrumentName string `json:"instrument_name"`
	} `json:"result"`
}

func (s *DeribitFuturesScraper) send(message *map[string]interface{}, market string, websocketConn *websocket.Conn) error {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		s.Logger.Printf("[ERROR] problem sending Deribit message on market: [%s], err: %s", market, err)
		return err
	}
	s.Logger.Printf("[DEBUG] sent message [%s]: %s", market, message)
	return nil
}

// ScraperClose - responsible for closing out the scraper for a market
func (s *DeribitFuturesScraper) ScraperClose(market string, websocketConnection interface{}) error {
	switch c := websocketConnection.(type) {
	case *websocket.Conn:
		err := c.WriteJSON(map[string]string{"op": "unsubscribe", "channel": "trades", "market": market})
		if err != nil {
			s.Logger.Printf("[ERROR] failed to send close message for [%s] websocket, err: %s", market, err)
			return err
		}
		err = c.Close()
		if err != nil {
			s.Logger.Printf("[ERROR] failed to close the websocket for [%s]", market)
			return err
		}
		time.Sleep(time.Duration(retryIn) * time.Second)
		return nil
	default:
		s.Logger.Printf("[ERROR] unknown connection type: %T. Expected gorilla/websocket pointer.", c)
		return fmt.Errorf("unknown connection type: %T", c)
	}
}

// Authenticate - authenticates with your access key and access secret to retrieve the trade details
func (s *DeribitFuturesScraper) Authenticate(market string, websocketConnection interface{}) error {
	switch c := websocketConnection.(type) {
	case *websocket.Conn:
		err := s.send(&map[string]interface{}{
			"method": "public/auth",
			"params": &map[string]string{
				"grant_type": "client_credentials",
				"client_id":     s.AccessKey,
				"client_secret": s.AccessSecret,
			},
			"jsonrpc": "2.0",
		}, market, c)
		return err
	default:
		s.Logger.Printf("[ERROR] unknown connection type: %T. Expected gorilla/websocket pointer.", c)
		return fmt.Errorf("unknown connection type: %T", c)
	}
}

// when we authenticate, we get back a refresh token that we use to keep alive our websocket connection
func (s *DeribitFuturesScraper) refreshToken(market string, websocketConn *websocket.Conn) (bool, error) {
	if s.RefereshToken == "" {
		return false, nil
	}
	// refresh
	err := s.send(&map[string]interface{}{
		"method": "public/auth",
		"params": &map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": s.RefereshToken,
		},
		"jsonrpc": "2.0",
	}, market, websocketConn)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Scrape starts a websocket scraper for market
func (s *DeribitFuturesScraper) Scrape(market string) {
	s.validateMarket(market)
	s.validateRefreshEveryToken()
	for {
		// immediately invoked function expression for easy clenup with defer
		func() {
			u := url.URL{Scheme: "wss", Host: "www.deribit.com", Path: "/ws/api/v2/"}
			s.Logger.Printf("[DEBUG] connecting to [%s], market: [%s]", u.String(), market)
			ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				s.Logger.Printf("[ERROR] dial: %s", err)
				return
			}
			defer s.ScraperClose(market, ws)
			// 1. authenticate
			err = s.Authenticate(market, ws)
			if err != nil {
				s.Logger.Printf("[ERROR] could not authenticate. Retrying.")
				return
			}
			time.Sleep(time.Second)
			// 2. subscribe to channel
			err = s.send(&map[string]interface{}{
				"method": "private/subscribe",
				"params": &map[string]interface{}{
					"channels": []string{"trades." + market + ".raw"},
				},
				"jsonrpc": "2.0",
				"id":      0,
			}, market, ws)
			if err != nil {
				s.Logger.Printf("[ERROR] could not send ws message. Restarting the websocket.")
				return
			}
			// 3. refresh the token more often than 900 seconds
			tick := time.NewTicker(time.Duration(s.RefreshTokenEvery) * time.Second) // every RefreshTokenEvery seconds we have to refresh token
			defer tick.Stop()
			// we require a separate goroutine for ticker, so that we can refresh our access token everyRefreshToken seconds
			go func() {
				for {
					select {
					case <-tick.C:
						s.refreshToken(market, ws)
					}
				}
			}()
			for {
				_, message, err := ws.ReadMessage()
				decodedMsg := messageDeribit{}
				if err != nil {
					s.Logger.Printf("[ERROR] problem reading Deribit on [%s], err: %s", market, err)
					return
				}
				strMessage := string(message)
				s.Logger.Printf("[DEBUG] received new message: %v", strMessage)
				// check if the received message containss the refresh_token json key
				if strings.Contains(strMessage, "refresh_token") {
					err = json.Unmarshal(message, &decodedMsg)
					if err != nil {
						s.Logger.Printf("[ERROR] problem unmarshalling the message that contains refresh_token")
						return
					}
					s.Logger.Printf("[INFO] obtained a new refresh token on [%s], updating '%s'", market, s.RefereshToken)
					s.RefereshToken = decodedMsg.Result.RefreshToken
				} else {
					// only save the messages if the message does not contain thre refresh_token string
					s.Logger.Printf("[DEBUG] saving new message on [%s]", market)
					_, err = s.Writer.Write(strMessage+"|", scrapeDataSaveLocationDeribit+s.Writer.GetWriteFileName("deribit", market))
					if err != nil {
						s.Logger.Printf("[ERROR] issue saving to file on [%s], err: %s", market, err)
						return
					}
				}
			}
		}()
	}
}

// ScrapeMarkets - will scrape the markets specified during instantiation
func (s *DeribitFuturesScraper) ScrapeMarkets() {
	for _, market := range s.Markets {
		s.WaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.WaitGroup.Wait()
}

func (s *DeribitFuturesScraper) validateMarket(market string) {
	allFuturesMarketsDeribit, err := allDeribitFuturesMarkets()
	if err != nil {
		panic(fmt.Sprintf("Could not fetch any futures markets, err: %s", err))
	}
	containsMarket := utils.Contains(&allFuturesMarketsDeribit, market)
	if !containsMarket {
		panic(fmt.Sprintf("Market %s is not available", market))
	}
}

func (s *DeribitFuturesScraper) validateRefreshEveryToken() {
	if s.RefreshTokenEvery >= 900 {
		panic("When you use https://testapp.deribit.com/api_console, you will see that upon a successful authentication, the response will include expiresIn field. Which is set at 900. This means that the token we generated is only valid for 900 seconds, and we have to refresh it before then.")
	}
}

// this will fetch all the futures markets, either ETH or BTC. We use this method to validate that the markets
// that the user instantiated the scraper with are valid. This is done in the validateMarket method.
func makeFuturesMarketsRequest(market string) ([]string, error) {
	if market != "BTC" && market != "ETH" {
		panic("Unsupported market. Only BTC & ETH are supported.")
	}
	resp, err := http.Get("https://www.deribit.com/api/v2/public/get_instruments?currency=" + market)
	if err != nil {
		fmt.Printf("[ERROR] issue making a GET request to obtain all tradable futures markets. Err: %s", err)
		return []string{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[ERROR] issue reading the GET request body. Err: %s", err)
		return []string{}, err
	}
	decodedMsg := messageGETDeribit{}
	err = json.Unmarshal(body, &decodedMsg)
	if err != nil {
		fmt.Printf("[ERROR] issue decoding the JSON response. Err: %s", err)
	}
	allMarkets := []string{}
	for _, market := range decodedMsg.Result {
		allMarkets = append(allMarkets, market.InstrumentName)
	}
	// fmt.Println("[INFO] the decoded message is:", decodedMsg)
	return allMarkets, nil
}

func allDeribitFuturesMarkets() ([]string, error) {
	BTCMarkets, err := makeFuturesMarketsRequest("BTC")
	if err != nil {
		return []string{}, fmt.Errorf("Could not fetch BTC futures markets, err:%s", err)
	}
	ETHMarkets, err := makeFuturesMarketsRequest("ETH")
	if err != nil {
		return []string{}, fmt.Errorf("Could not fetch ETH futures markets, err:%s", err)
	}
	return append(BTCMarkets, ETHMarkets...), nil
}

// example usage
// package main

// import (
// 	"log"
// 	"os"
// 	"sync"

// 	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
// 	writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
// )

// func main() {
// 	wg := sync.WaitGroup{}
// 	writer := writers.FileWriter{}
// 	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

// 	var deribitScraper scrapers.FuturesScraper = &scrapers.DeribitFuturesScraper{
// 		WaitGroup: &wg,
// 		Markets:   []string{"BTC-PERPETUAL", "ETH-PERPETUAL", "ETH-6DEC19-220-C", "ETH-13DEC19-180-C"},
// 		Writer:    &writer,
// 		Logger:    logger,

// 		AccessKey:    "yourDeribitAccessKey",
// 		AccessSecret: "yourDeribiAccessSecret",

// 		RefreshTokenEvery: 60,
// 	}
// 	deribitScraper.ScrapeMarkets()

// 	wg.Wait()
// }

