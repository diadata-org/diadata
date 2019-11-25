package scrapers

import (
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

func (s *FTXFuturesScraper) send(message *map[string]string, market string, websocketConn *websocket.Conn) {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		s.Logger.Printf("[ERROR] problem sending FTX message on market: [%s], err: %s", market, err)
		return
	}
	s.Logger.Printf("[DEBUG] sent message [%s]: %s", market, message)
}

func (s *FTXFuturesScraper) scraperClose(market string, websocketConn *websocket.Conn) {
	err := websocketConn.WriteJSON(map[string]string{"op": "unsubscribe", "channel": "trades", "market": market})
	if err != nil {
		s.Logger.Printf("[ERROR] failed to send close message for [%s] websocket, err: %s", market, err)
		return
	}
	err = websocketConn.Close()
	if err != nil {
		s.Logger.Printf("[ERROR] failed to close the websocket for [%s]", market)
		return
	}
	time.Sleep(time.Duration(retryIn) * time.Second)
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
				s.Logger.Printf("[ERROR] dial: %s", err)
				return
			}
			defer s.scraperClose(market, ws)
			s.send(&map[string]string{"market": market, "channel": "trades", "op": "subscribe"}, market, ws)
			s.send(&map[string]string{"op": "ping"}, market, ws)
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
				if err != nil {
					s.Logger.Printf("[ERROR] problem reading FTX on [%s], err: %s", market, err)
					return
				}
				_, err = s.Writer.Write(string(message)+"|", scrapeDataSaveLocationFTX+s.Writer.GetWriteFileName("ftx", market))
				s.Logger.Printf("[DEBUG] received new message: %s", message)
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
