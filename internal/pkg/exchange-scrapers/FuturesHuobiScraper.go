package scrapers

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
	utils "github.com/diadata-org/diadata/pkg/utils"
	zap "go.uber.org/zap"
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
	Logger    *zap.SugaredLogger
}

// --------------------------------------------------------------------------------------------

// NewHuobiFuturesScraper - returns an instance of the Huobi scraper
func NewHuobiFuturesScraper(markets []string) FuturesScraper {
	wg := sync.WaitGroup{}
	writer := writers.FileWriter{}
	logger := zap.NewExample().Sugar() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	var scraper FuturesScraper = &HuobiFuturesScraper{
		WaitGroup: &wg,
		Markets:   markets, // []string{"BNB-PERP", "ETH-PERP", "BTC-PERP", "EOS-PERP"}
		Writer:    &writer,
		Logger:    logger,
	}

	return scraper
}

// sends a byte message to huobi websocket
func (s *HuobiFuturesScraper) send(message []byte, market string, websocketConn *websocket.Conn) (int, error) {
	n, err := websocketConn.Write(message)
	s.Logger.Debugf("[%s] send: %s", market, message)
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
			return err
		}
		err = c.Close()
		if err != nil {
			return err
		}
		s.Logger.Infof("gracefully shutdown huobi scraper on market: %s", market)
		time.Sleep(time.Duration(retryIn) * time.Second)
		return nil
	default:
		return fmt.Errorf("unknown connection type, expected golang.org/x/net/websocket, got: %T", connection)
	}
}

// Scrape starts a websocket scraper for market
func (s *HuobiFuturesScraper) Scrape(market string) {
	s.validateMarket(market)

	// this block is for listening to sigterms and interupts
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	userCancelled := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		userCancelled <- true
	}()

	for {
		// IIFE for easy cleanup with defer
		func() {
			ws, err := websocket.Dial(wsURLHuobi, "", "http://www.google.com")
			// defer inside of the function will cleanup before the next run
			defer s.ScraperClose(market, ws)
			if err != nil {
				// an error opening is fatal. let this kill the programme
				s.Logger.Errorf("[%s] %s", market, err)
				time.Sleep(time.Duration(retryIn) * time.Second)
				return
			}
			// subscribe to the trade detail channel
			message := []byte("{\"Sub\":\"market." + market + ".trade.detail\"}")
			_, err = s.send(message, market, ws)
			if err != nil {
				s.Logger.Errorf("problem subscriping to the [%s] trade channel, err: %s", market, err)
				return
			}
			// create the conduit for the received messages
			var msg = make([]byte, 512)
			for {
				select {
				case <-userCancelled:
					s.Logger.Infof("received interrupt, gracefully shutting down")
					s.ScraperClose(market, ws)
					os.Exit(0)
				default:
					m, err := ws.Read(msg)
					if err != nil {
						s.Logger.Errorf("[%s] %s", market, err)
						// an error reading means we may have lost the connection
						// return out and just try again
						return
					}
					newmsg := msg[:m]
					unzipmsg, err := parseGzip(newmsg)
					if err != nil {
						s.Logger.Errorf("[%s] problem saving to %s, err: %s", market, s.Writer.GetWriteFileName("huobi", market), err)
						return
					}
					s.Logger.Debugf("[%s] byteLen:%d, unzipLen:%d %s", market, m, len(unzipmsg), unzipmsg)
					if len(unzipmsg) == pingMsgLengthHuobi {
						if "ping" == string(unzipmsg[2:6]) {
							_, err := s.pong(string(unzipmsg[8:21]), market, ws)
							if err != nil {
								s.Logger.Errorf("[%s] problem ponging the websocket server, err: %s", market, err)
								return
							}
						}
					} else {
						// ensure that scrapeDataSaveLocation exists
						_, err := s.Writer.Write(string(unzipmsg)+"\n", scrapeDataSaveLocationHuobi+s.Writer.GetWriteFileName("huobi", market))
						if err != nil {
							s.Logger.Errorf("[%s] problem saving to %s, err: %s", market, s.Writer.GetWriteFileName("huobi", market), err)
							return
						}
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
		panic("incorrect market provided. should be of the form symbol_frequency")
	}
	tradeSymbol := parts[0]
	tradeFrequency := parts[1]
	containsSymbol := utils.Contains(&allowedMarketsHuobi, tradeSymbol)
	if !containsSymbol {
		panic("provided trade symbol is not supported")
	}
	containsFreq := utils.Contains(&allowedFrequenciesHuobi, tradeFrequency)
	if !containsFreq {
		panic("provided frequency is not supported")
	}
}

// Huobi websocket API sends back gzips, this will parse it.
func parseGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	defer r.Close()
	if err != nil {
		return nil, err
	}
	unzipped, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return unzipped, nil
}
