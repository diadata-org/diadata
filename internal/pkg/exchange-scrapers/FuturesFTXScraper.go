package scrapers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	utils "github.com/diadata-org/diadata/internal/pkg/scraper-utils"
	writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
)

var allFuturesMarketsFTX = []string{"BTC-PERP", "ETH-PERP", "LINK-PERP",
	"EOS-PERP", "BNB-PERP", "BCH-PERP", "XRP-PERP", "BSV-PERP", "ALGO-PERP",
	"DRGN-PERP", "HT-PERP", "LTC-PERP", "LEO-PERP", "SHIT-PERP", "TRX-PERP",
	"USDT-PERP", "EXCH-PERP", "BTMX-PERP", "ALT-PERP", "ADA-PERP", "MID-PERP",
	"OKB-PERP", "MATIC-PERP", "ATOM-PERP", "ETC-PERP", "TOMO-PERP", "DOGE-PERP"}

const scrapeDataSaveLocationFTX = ""

// FTXFuturesScraper - scrapes the futures from the FTX exchange
type FTXFuturesScraper struct {
	Markets   []string
	WaitGroup *sync.WaitGroup
	Writer    writers.Writer
	Logger    *log.Logger
}

type tradeMessageFTX struct {
	Type string `json:"type"`
}

func (s *FTXFuturesScraper) send(message *map[string]string, market string, websocketConn *websocket.Conn) error {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		return err
	}
	s.Logger.Printf("[DEBUG] sent message [%s]: %s", market, message)
	return nil
}

// Authenticate - placeholder here, since we do not need to authneticate the connection.
func (s *FTXFuturesScraper) Authenticate(market string, connection interface{}) error { return nil }

// ScraperClose - safely closes the scraper; We pass the interface connection as the second argument
// primarily for the reason that Huobi scraper does not use the gorilla websocket; It uses golang's x websocket;
// If we did not define this method in our FuturesScraper interface, we could have easily used the pointer
// to gorilla websocket here; However, to make FuturesScraper more ubiquituous, we need an interface here.
func (s *FTXFuturesScraper) ScraperClose(market string, connection interface{}) error {
	switch c := connection.(type) {
	case *websocket.Conn:
		err := c.WriteJSON(map[string]string{"op": "unsubscribe", "channel": "trades", "market": market})
		if err != nil {
			return err
		}
		err = c.Close()
		if err != nil {
			return err
		}
		time.Sleep(time.Duration(retryIn) * time.Second)
		return nil
	default:
		return fmt.Errorf("unknown connection type, expected gorilla/websocket, got: %T", connection)
	}
}

// Scrape starts a websocket scraper for market
func (s *FTXFuturesScraper) Scrape(market string) {
	s.validateMarket(market)
	for {
		// immediately invoked function expression for easy clenup with defer
		func() {
			u := url.URL{Scheme: "wss", Host: "ftx.com", Path: "/ws"}
			s.Logger.Printf("[DEBUG] connecting to [%s], market: [%s]", u.String(), market)
			ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				s.Logger.Printf("[ERROR] could not fial ftx websocket: %s", err)
				return
			}
			defer s.ScraperClose(market, ws)
			err = s.send(&map[string]string{"market": market, "channel": "trades", "op": "subscribe"}, market, ws)
			if err != nil {
				s.Logger.Printf("[ERROR] could not send a channel subscription message. retrying, err: %s", err)
				return
			}
			err = s.send(&map[string]string{"op": "ping"}, market, ws)
			if err != nil {
				s.Logger.Printf("[ERROR] could not send an initial ping message. retrying, err: %s", err)
				return
			}
			tick := time.NewTicker(15 * time.Second) // every 15 seconds we have to ping FTX
			defer tick.Stop()
			// we require a separate goroutine for ticker, because ReadMessage is blocking
			// and we may fail sending ping before we get any message on a market, thus
			// forcing FTX to close our websocket out.
			go func() {
				for {
					select {
					case <-tick.C:
						s.send(&map[string]string{"op": "ping"}, market, ws)
					}
				}
			}()
			for {
				_, message, err := ws.ReadMessage()
				decodedMsg := tradeMessageFTX{}
				if err != nil {
					s.Logger.Printf("[ERROR] problem reading ftx on [%s], err: %s", market, err)
					return
				}
				err = json.Unmarshal(message, &decodedMsg)
				if err != nil {
					s.Logger.Printf("[ERROR] could not unmarshal ftx message on [%s], err: %s", market, err)
					return
				}
				s.Logger.Printf("[DEBUG] received new message: %s", message)
				if decodedMsg.Type != "subscribed" && decodedMsg.Type != "pong" && decodedMsg.Type != "unsubscribed" {
					s.Logger.Printf("[DEBUG] saving new message on [%s]", market)
					_, err = s.Writer.Write(string(message)+"|", scrapeDataSaveLocationFTX+s.Writer.GetWriteFileName("ftx", market))
					if err != nil {
						s.Logger.Printf("[ERROR] could not write to file, err: %s", err)
						return
					}
				}
			}
		}()
	}
}

// ScrapeMarkets - will scrape the markets specified during instantiation
func (s *FTXFuturesScraper) ScrapeMarkets() {
	for _, market := range s.Markets {
		s.WaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.WaitGroup.Wait()
}

func (s *FTXFuturesScraper) validateMarket(market string) {
	containsMarket := utils.Contains(&allFuturesMarketsFTX, market)
	if !containsMarket {
		panic(fmt.Sprintf("Market %s is not available", market))
	}
}

// example usage
// func main() {
// 	wg := sync.WaitGroup{}
// 	writer := writers.FileWriter{}
// 	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

// 	var ftxScraper scrapers.FuturesScraper = &scrapers.FTXFuturesScraper{
// 		WaitGroup: &wg,
// 		Markets:   []string{"BNB-PERP", "ETH-PERP", "BTC-PERP", "EOS-PERP"},
// 		Writer:    &writer,
// 		Logger:    logger,
// 	}
// 	ftxScraper.ScrapeMarkets()

// 	wg.Wait()
// }
