package scrapers

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
	utils "github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
	zap "go.uber.org/zap"
)

const scrapeDataSaveLocationCoinflex = ""

// CoinflexFuturesScraper - scrapes the futures from the Coinflex exchange
type CoinflexFuturesScraper struct {
	Markets   []string
	WaitGroup *sync.WaitGroup
	Writer    writers.Writer
	Logger    *zap.SugaredLogger
}

type tradeMessageCoinflex struct {
	Type string `json:"type"`
}

// the response of https://webapi.coinflex.com/markets/ is a list of marketCoinglex JSON objects
// This is used to validate that the market that you have selected to scrape acutally exists. This is done
// validateMarket function.
type marketCoinflex struct {
	Base     int64  `json:"base"`
	Counter  int64  `json:"string"`
	Name     string `json:"name"`
	SpotName string `json:"spot_name"`
	Tick     int64  `json:"tick"`
	Start    int64  `json:"start"`
	Expires  int64  `json:"expires"`
	Tenor    string `josn:"tenor"`
}

// the response of https://webapi.coinflex.com/assets/ is a list of [] of assetCoinflex JSON objects
// This is used to find the int id of the assets (base and quote <- coinflex call quote Counter, go figure)
// that you supply via market. So you would pass something like BTCDEC-USDDEC as the market, and the XXX method
// will then (1) check that such a market exists using the validateMarket function, it will then split on "-" and
// will find the int ids of the two assets. It then uses these int ids to make a websocket request to subscribe to
// the trade channel.
type assetCoinflex struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	SpotID   int64  `json:"spot_id"`
	SpotName string `json:"spot_name"`
	Scale    int64  `json:"scale"`
}

// this is how the message looks like that we receive from the trades channel
// the OrderModified, OrderClosed, OrderOpen have different structures. Also,
// notice that the below is OrdersMatched (Order**s** <- plural, not Order like the others)
type ordersMatchedCoinflex struct {
	Notice        string `json:"notice"`
	Bid           int64  `json:"bid"`
	BidTonce      int64  `json:"bid_tonce"`
	Ask           int64  `json:"ask"`
	AskTonce      int64  `json:"ask_tonce"`
	Base          int64  `json:"base"`
	Counter       int64  `json:"counter"`
	Quantity      int64  `json:"quantity"`
	Price         int64  `json:"price"`
	Total         int64  `json:"total"`
	BidRem        int64  `json:"bid_rem"` // if this is 0, then it was a market sell order
	AskRem        int64  `json:"ask_rem"` // if this is 0, then it was a market buy order
	Time          int64  `json:"time"`
	BidBaseFee    int64  `json:"bid_base_fee"`
	BidCounterFee int64  `json:"bid_counter_fee"`
	AskBaseFee    int64  `json:"ask_base_fee"`
	AskCounterFee int64  `json:"ask_counter_fee"`
}

// NewCoinflexFuturesScraper - returns an instance of the coinflex scraper
func NewCoinflexFuturesScraper(markets []string) FuturesScraper {
	wg := sync.WaitGroup{}
	writer := writers.FileWriter{}
	logger := zap.NewExample().Sugar() // or NewProduction, or NewDevelopment
	defer logger.Sync()

	var scraper FuturesScraper = &CoinflexFuturesScraper{
		WaitGroup: &wg,
		Markets:   markets,
		Writer:    &writer,
		Logger:    logger,
	}

	return scraper
}

func (s *CoinflexFuturesScraper) send(message *map[string]interface{}, market string, websocketConn *websocket.Conn) error {
	err := websocketConn.WriteJSON(*message)
	if err != nil {
		return err
	}
	s.Logger.Debugf("sent message [%s]: %s", market, message)
	return nil
}

// Authenticate - placeholder here, since we do not need to authneticate the connection.
func (s *CoinflexFuturesScraper) Authenticate(market string, connection interface{}) error { return nil }

// ScraperClose - safely closes the scraper; We pass the interface connection as the second argument
// primarily for the reason that Huobi scraper does not use the gorilla websocket; It uses golang's x websocket;
// If we did not define this method in our FuturesScraper interface, we could have easily used the pointer
// to gorilla websocket here; However, to make FuturesScraper more ubiquituous, we need an interface here.
func (s *CoinflexFuturesScraper) ScraperClose(market string, connection interface{}) error {
	switch c := connection.(type) {
	case *websocket.Conn:
		err := s.write(websocket.CloseMessage, []byte{}, c)
		if err != nil {
			return err
		}
		err = c.Close()
		if err != nil {
			return err
		}
		s.Logger.Infof("gracefully shutdown coinflex scraper on market: %s", market)
		time.Sleep(time.Duration(retryIn) * time.Second)
		return nil
	default:
		return fmt.Errorf("unknown connection type: %T", connection)
	}
}

// Scrape starts a websocket scraper for market
func (s *CoinflexFuturesScraper) Scrape(market string) {
	validated, err := s.validateMarket(market)
	if !validated || err != nil {
		s.Logger.Errorf("could not validate %s market", market)
		if err != nil {
			s.Logger.Errorf("issue with validating, err: %s", err)
		}
		return
	}
	baseID, quoteID, err := s.getBaseAndCounterID(market)
	// splits the string market into the base and the counter and then finds the int id of them.
	// coinflex expects that we provide an int for the assets when we make the websocket requests.
	if err != nil {
		s.Logger.Errorf("issue with getting an id for base and quote: %s", err)
		return
	}

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
			u := url.URL{Scheme: "wss", Host: "api.coinflex.com", Path: "/v1"}
			s.Logger.Debugf("connecting to [%s], market: [%s]", u.String(), market)
			ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				s.Logger.Errorf("dial: %s", err)
				time.Sleep(time.Duration(retryIn) * time.Second)
				return
			}
			defer s.ScraperClose(market, ws)
			// to let you know that the websocket connection is alive. Coinflex do not have the heartbeat channel
			// and they send you frame pong messages. Thus, this handler.
			ws.SetPongHandler(func(appData string) error {
				s.Logger.Debugf("received a pong frame")
				return nil
			})
			err = s.send(&map[string]interface{}{"base": baseID, "counter": quoteID, "watch": true, "method": "WatchOrders"}, market, ws)
			if err != nil {
				s.Logger.Errorf("could not send a channel subscription message. retrying")
				return
			}
			if err != nil {
				s.Logger.Errorf("could not send an initial ping message. retrying")
				return
			}
			tick := time.NewTicker(30 * time.Second) // every 45 seconds we have to ping Coinflex. we also have a 15 second write limit of the ping frame (thus, 30 seconds here)
			defer tick.Stop()
			// we require a separate goroutine for ticker, because ReadMessage is blocking
			// and we may fail sending ping before we get any message on a market, thus
			// forcing Coinflex to close our websocket out.
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
					msg := ordersMatchedCoinflex{}
					if err != nil {
						s.Logger.Errorf("problem reading coinflex on [%s], err: %s", market, err)
						return
					}
					err = json.Unmarshal(message, &msg)
					if err != nil {
						s.Logger.Errorf("could not unmarshal coinflex message on [%s], err: %s", market, err)
						return
					}
					s.Logger.Debugf("received a message: %s", message)
					if msg.Notice == "OrdersMatched" {
						s.Logger.Debugf("received new match message on [%s]: %s", market, message)
						_, err = s.Writer.Write(string(message)+"\n", scrapeDataSaveLocationCoinflex+s.Writer.GetWriteFileName("coinflex", market))
						if err != nil {
							s.Logger.Errorf("could not save to file: %s, on market: [%s], err: %s", scrapeDataSaveLocationCoinflex+s.Writer.GetWriteFileName("coinflex", market), market, err)
							return
						}
					}
				}
			}
		}()
	}
}

// write's primary purpose is to write a ping frame op code to keep the websocket connection alive
func (s *CoinflexFuturesScraper) write(mt int, payload []byte, ws *websocket.Conn) error {
	ws.SetWriteDeadline(time.Now().Add(15 * time.Second))
	return ws.WriteMessage(mt, payload)
}

// ScrapeMarkets - will scrape the markets specified during instantiation
func (s *CoinflexFuturesScraper) ScrapeMarkets() {
	for _, market := range s.Markets {
		s.WaitGroup.Add(1)
		go s.Scrape(market)
	}
	s.WaitGroup.Wait()
}

func (s *CoinflexFuturesScraper) getBaseAndCounterID(market string) (int64, int64, error) {
	assets := strings.Split(market, "/")
	var baseID int64 = 0
	var quoteID int64 = 0
	base := assets[0]
	quote := assets[1] // coinflex call this "counter"
	baseID, err := s.assetID(base)
	if err != nil {
		return baseID, quoteID, err
	}
	quoteID, err = s.assetID(quote)
	if err != nil {
		return baseID, quoteID, err
	}
	return baseID, quoteID, nil
}

// ensures that market available to trade
func (s *CoinflexFuturesScraper) validateMarket(market string) (bool, error) {
	// should validate that there is an available market
	marketAvailable := false
	marketsCoinflex, err := s.availableMarketsCoinflex()
	fmt.Printf("[DEBUG] all coinflex's available markets are: %v\n", marketsCoinflex)
	if err != nil {
		return marketAvailable, err
	}
	for _, availableMarket := range marketsCoinflex {
		if availableMarket.Name == market {
			marketAvailable = true
		}
	}
	return marketAvailable, nil
}

func (s *CoinflexFuturesScraper) availableMarketsCoinflex() ([]marketCoinflex, error) {
	body, err := utils.GetRequest("https://webapi.coinflex.com/markets/")
	if err != nil {
		return []marketCoinflex{}, err
	}

	markets := []marketCoinflex{}
	err = json.Unmarshal(body, &markets)
	if err != nil {
		return []marketCoinflex{}, err
	}
	return markets, nil
}

// uses /assets/ GET endpoint to obtain all the Coinflex's assets
func (s *CoinflexFuturesScraper) getAllAssets() ([]assetCoinflex, error) {
	body, err := utils.GetRequest("https://webapi.coinflex.com/assets/")
	if err != nil {
		return nil, err
	}
	s.Logger.Debugf("retrieved all of the Coinflex assets: %s", string(body))
	assets := []assetCoinflex{}
	err = json.Unmarshal(body, &assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

// gives you the id of the asset. Asset can be, not limited to, ETH, XBTJUL, BTCDEC, etc.
func (s *CoinflexFuturesScraper) assetID(asset string) (int64, error) {
	var assetsID int64 = 0
	assets, err := s.getAllAssets()
	if err != nil {
		return assetsID, fmt.Errorf("could not retrieve all Coinflex's assets, err: %s", err)
	}
	for _, assetObj := range assets {
		if assetObj.Name == asset {
			assetsID = assetObj.ID
		}
	}
	return assetsID, nil
}
