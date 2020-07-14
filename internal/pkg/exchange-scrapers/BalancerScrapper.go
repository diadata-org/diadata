package scrapers

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	zap "go.uber.org/zap"

	writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
	"github.com/gorilla/websocket"
)

const scrapeDataSaveLocationBalancer = ""

// BalancerScraper - use the NewBalancerScraper function to create an instance
type BalancerScraper struct {
	Markets   []string
	WaitGroup *sync.WaitGroup
	Writer    writers.Writer
	Logger    *zap.SugaredLogger
}

// NewBalancerScraper - returns an instance of an options scraper.
func NewBalancerScraper(markets []string) FuturesScraper {
	wg := sync.WaitGroup{}
	writer := writers.FileWriter{}
	logger := zap.NewExample().Sugar() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	var scraper FuturesScraper = &BalancerScraper{
		WaitGroup: &wg,
		Markets:   markets,
		Writer:    &writer,
		Logger:    logger,
	}

	return scraper
}

func (s *BalancerScraper) send(message *map[string]interface{}, market string, websocketConn *websocket.Conn) error {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		return err
	}
	s.Logger.Debugf("sent message [%s]: %s", market, message)
	return nil
}

// Authenticate - placeholder here, since we do not need to authneticate the connection.
func (s *BalancerScraper) Authenticate(market string, connection interface{}) error {
	return nil
}

// ScraperClose - safely closes the scraper; We pass the interface connection as the second argument
// primarily for the reason that Huobi scraper does not use the gorilla websocket; It uses golang's x websocket;
// If we did not define this method in our FuturesScraper interface, we could have easily used the pointer
// to gorilla websocket here; However, to make FuturesScraper more ubiquituous, we need an interface here.
func (s *BalancerScraper) ScraperClose(market string, connection interface{}) error {
	switch c := connection.(type) {
	case *websocket.Conn:
		// unsubscribe from the channel
		err := s.send(&map[string]interface{}{"jsonrpc": "2.0", "method": "unsubscribe", "params": &map[string]interface{}{"channel": "lightning_ticker_" + market}}, market, c)
		if err != nil {
			s.Logger.Errorf("could not send a channel unsubscription message, err: %s", err)
			return err
		}
		// close the websocket connection
		err = s.write(websocket.CloseMessage, []byte{}, c)
		if err != nil {
			return err
		}
		err = c.Close()
		if err != nil {
			return err
		}
		s.Logger.Infof("gracefully shutdown bitflyer scraper on market: %s", market)
		time.Sleep(time.Duration(retryIn) * time.Second)
		return nil
	default:
		return fmt.Errorf("unknown connection type: %T", connection)
	}
}

// Scrape starts a websocket scraper for market
func (s *BalancerScraper) Scrape(market string) {
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
		// immediately invoked function expression for easy clenup with defer
		func() {
			u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
			s.Logger.Debugf("connecting to [%s], market: [%s]", u.String(), market)
			ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				s.Logger.Errorf("could not dial Bitflyer websocket: %s", err)
				time.Sleep(time.Duration(retryIn) * time.Second)
				return
			}
			defer s.ScraperClose(market, ws)
			ws.SetPongHandler(func(appData string) error {
				s.Logger.Debugf("received a pong frame")
				return nil
			})
			err = s.send(&map[string]interface{}{"jsonrpc": "2.0", "method": "subscribe", "params": &map[string]interface{}{"channel": "lightning_ticker_" + market}}, market, ws)
			if err != nil {
				s.Logger.Errorf("could not send a channel subscription message. retrying, err: %s", err)
				return
			}
			tick := time.NewTicker(15 * time.Second)
			defer tick.Stop()
			go func() {
				for {
					select {
					case <-tick.C:
						err := s.write(websocket.PingMessage, []byte{}, ws)
						if err != nil {
							s.Logger.Errorf("error experienced pinging coinflex, err: %s", err)
							return
						}
						s.Logger.Debugf("pinged the coinflex server. market: [%s]", market)
					}
				}
			}()
			for {
				select {
				case <-userCancelled:
					s.Logger.Infof("received interrupt, gracefully shutting down")
					s.ScraperClose(market, ws)
					os.Exit(0)
				default:
					_, message, err := ws.ReadMessage()
					if err != nil {
						s.Logger.Errorf("repeated read error, restarting")
						return
					}
					s.Logger.Debugf("received new message: %s, saving new message", message)
					_, err = s.Writer.Write(string(message)+"\n", scrapeDataSaveLocationBalancer+s.Writer.GetWriteFileName("Bitflyer", market))
					if err != nil {
						s.Logger.Errorf("could not write to file, err: %s", err)
						return
					}
				}
			}
		}()
	}
}

// write's primary purpose is to write a ping frame op code to keep the websocket connection alive
func (s *BalancerScraper) write(mt int, payload []byte, ws *websocket.Conn) error {
	ws.SetWriteDeadline(time.Now().Add(15 * time.Second))
	return ws.WriteMessage(mt, payload)
}

// ScrapeMarkets - will scrape the markets specified during instantiation
func (s *BalancerScraper) ScrapeMarkets() {
	for _, market := range s.Markets {
		s.WaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.WaitGroup.Wait()
}
