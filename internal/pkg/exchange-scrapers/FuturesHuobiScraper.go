package scrapers

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	utils "github.com/nazariyv/diadata-scrapers/pkg/utils"
	writers "github.com/nazariyv/diadata-scrapers/pkg/writers"
	"golang.org/x/net/websocket"
)

// --------------------------------- Config --------------------------------------------------
// Huobi API configuration
const (
	scrapeDataSaveLocationHuobi string = "" // location in which to save. ensure that the folder exists.
	// files will be created automatically

	// API Endpoints
	marketURLHuobi     string = "https://api.hbdm.com"
	wsURLHuobi         string = "wss://www.hbdm.com/ws"
	pingMsgLengthHuobi int    = 22
)

var (
	allowedMarketsHuobi     = []string{"BTC", "ETC", "ETH", "EOS", "LTC", "BCH", "XRP", "TRX", "BSV"}
	allowedFrequenciesHuobi = []string{"CW", "NW", "CQ"}
	bufferHuobi             bytes.Buffer
)

// ---------------

// HuobiFuturesScraper - scrapes huobi's futures markets
type HuobiFuturesScraper struct {
	Markets   []string // markets to scrape. To scrape all, call AllFuturesMarketsHuobi()
	WaitGroup *sync.WaitGroup
	Writer    writers.Writer // an interface to write the messages
	Logger    *log.Logger
}

// --------------------------------------------------------------------------------------------

// sends a byte message to huobi websocket
func (s *HuobiFuturesScraper) send(message []byte, market string, websocketConn *websocket.Conn) (int, error) {
	n, err := websocketConn.Write(message)
	s.Logger.Printf("[DEBUG] [%s] Send: %s\n", market, message)
	return n, err
}

// pongs back the huobi websocket client to keep the connection alive
func (s *HuobiFuturesScraper) pong(time string, market string, websocketConn *websocket.Conn) (int, error) {
	n, err := s.send([]byte(fmt.Sprintf("{\"pong\":%s}", time)), market, websocketConn)
	return n, err
}

// Authenticate - not required for Huobi to scrape the futures data.
func (s *HuobiFuturesScraper) Authenticate(market string, connection interface{}) error { return nil }

// ScraperClose - clean up after the scraper.
func (s *HuobiFuturesScraper) ScraperClose(market string, connection interface{}) error {
	switch c := connection.(type) {
	case *websocket.Conn:
		message := []byte("{\"Unsub\":\"market." + market + ".trade.detail\"}")
		_, err := s.send(message, market, c)
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
		s.Logger.Printf("[ERROR] unknown connection type: %T. Expected golang.org/x/net/websocket pointer.", connection)
		return fmt.Errorf("unknown connection type: %T", connection)
	}
}

// Scrape starts a websocket scraper for market
func (s *HuobiFuturesScraper) Scrape(market string) {
	s.validateMarket(market)
	for {
		// IIFE for easy cleanup with defer
		func() {
			ws, err := websocket.Dial(wsURLHuobi, "", "http://www.google.com")
			// defer inside of the function will cleanup before the next run
			defer s.ScraperClose(market, ws)
			if err != nil {
				// an error opening is fatal. let this kill the programme
				s.Logger.Printf("[ERROR] [%s] %v\n", market, err)
				return
			}
			// subscribe to the trade detail channel
			message := []byte("{\"Sub\":\"market." + market + ".trade.detail\"}")
			_, err = s.send(message, market, ws)
			if err != nil {
				s.Logger.Printf("[ERROR] Problem subscriping to the [%s] trade channel. %v\n", market, err)
				return
			}
			// create the conduit for the received messages
			var msg = make([]byte, 512)
			for {
				m, err := ws.Read(msg)
				if err != nil {
					s.Logger.Printf("[ERROR] [%s] %v\n", market, err)
					// an error reading means we may have lost the connection
					// return out and just try again
					return
				}
				newmsg := msg[:m]
				unzipmsg, err := parseGzip(newmsg)
				if err != nil {
					s.Logger.Printf("[ERROR] [%s] %v\n", market, err)
					return
				}
				s.Logger.Printf("[DEBUG] [%s] byteLen:%d, unzipLen:%d %s\n", market, m, len(unzipmsg), unzipmsg)
				if len(unzipmsg) == pingMsgLengthHuobi {
					if "ping" == string(unzipmsg[2:6]) {
						_, err := s.pong(string(unzipmsg[8:21]), market, ws)
						if err != nil {
							s.Logger.Printf("[ERROR] [%s] Problem ponging the websocket server. %v\n", market, err)
							return
						}
					}
				} else {
					// ensure that scrapeDataSaveLocation exists
					_, err := s.Writer.Write(string(unzipmsg)+"|", scrapeDataSaveLocationHuobi+s.Writer.GetWriteFileName("huobi", market))
					if err != nil {
						s.Logger.Printf("[ERROR] [%s] problem saving to %s. %v\n", market, s.Writer.GetWriteFileName("huobi", market), err)
						return
					}
				}
			}
		}()
	}
	// s.waitGroup.Done()
}

// ScrapeMarkets - will scrape the markets specified during instantiation
func (s *HuobiFuturesScraper) ScrapeMarkets() {
	for _, market := range s.Markets {
		s.WaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.WaitGroup.Wait()
}

// ------------- Huobi util functions -------------------

// AllFuturesMarketsHuobi - returns all the futures markets tradable on Huobi.
// Lists all of the Huobi Futures markets. TODO: add a REST HTTP call to obtain the list
// of trdabale markets.
func AllFuturesMarketsHuobi() []string {
	allContracts := []string{}
	for _, market := range allowedMarketsHuobi {
		for _, maturity := range allowedFrequenciesHuobi {
			allContracts = append(allContracts, market+"_"+maturity)
		}
	}
	return allContracts
}

// Ensures that the provided market to scrape is supported by Huobi
// This function is only required if we manually write out the markets.
// This function will be removed when we make AllHuobiFuturesMarkets make
// a rest request
func (s *HuobiFuturesScraper) validateMarket(market string) {
	parts := strings.Split(market, "_")
	if len(parts) != 2 {
		panic("Incorrect market provided. Should be of the form symbol_frequency")
	}
	tradeSymbol := parts[0]
	tradeFrequency := parts[1]
	containsSymbol := utils.Contains(&allowedMarketsHuobi, tradeSymbol)
	if !containsSymbol {
		panic("Provided trade symbol is not supported")
	}
	containsFreq := utils.Contains(&allowedFrequenciesHuobi, tradeFrequency)
	if !containsFreq {
		panic("Provided frequency is not supported")
	}
}

// Huobi websocket API sends back gzips, this will parse it.
func parseGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	defer r.Close()
	if err != nil {
		fmt.Println("[parseGzip] NewReader error:", err)
		return nil, err
	}
	unzipped, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("[parseGzip] ioutil.ReadAll error:", err)
		return nil, err
	}
	return unzipped, nil
}
