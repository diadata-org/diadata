package scrapers

// !! IF YOU EVER HAVE ISSUES JUST MAKE A COMMENT STRAIGHT AFTER READING A MESSAGE FROM WEBSOCKET
import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var exists nothing

// https://github.com/coinflex-exchange/API/blob/master/SCALE.md
const coinflexScale float64 = 10000.0

type coinflexNotice struct {
	Notice string `json:"notice"`
}

// CoinflexScraper - scrapes Coinflex. You need to call ScrapePair
// to commence scraping. You can add as many active / trading markets
// as you please. Supports Spot and Futures markets
type CoinflexScraper struct {
	// websocket related
	connectionURL url.URL

	pairScrapers sync.Map

	// ! note I have made this public
	// trades from all of the pairs go here
	chanTrades chan *dia.Trade

	err error

	// used to listen for shutdown signals
	shutdown chan nothing
	// signals scraper is done
	shutdownDone chan nothing
}

// CoinflexPairScraper - a scraper for a single pair. This is where
// the scraping actually happens. All of the scraped trades are sent
// to chanTrades in CoinflexScraper that spawned this pair scraper
type CoinflexPairScraper struct {
	CflxPair dia.Pair

	// Connections support one concurrent reader and one concurrent writer
	// need a separate wsClient for each pair.
	wsClient *websocket.Conn

	// this should be a pointer to CoinflexScraper's chanTrades
	chanTrades chan *dia.Trade

	// have we closed this pair scraper?
	Closed bool

	err error
}

type coinflexWsWatchOrders struct {
	Ask       int64   `json:"ask"`
	AskRem    int64   `json:"ask_rem"`
	Base      int64   `json:"base"`
	Bid       int64   `json:"bid"`
	BidRem    int64   `json:"bid_rem"`
	Counter   int64   `json:"counter"`
	Notice    string  `json:"notice"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
	TakerSide string  `json:"taker_side"`
	Time      int64   `json:"time"`
	Total     int64   `json:"total"`
}

// NewCoinflexScraper - get an instance of CoinflexScraper
// use ScraperPair to add a markte to be scraped. It will
// send the trades from this market to chanTrades.
func NewCoinflexScraper() *CoinflexScraper {
	s := &CoinflexScraper{
		connectionURL: url.URL{Scheme: "wss",
			Host: "api.coinflex.com",
			Path: "/v1"},
		chanTrades:   make(chan *dia.Trade),
		err:          nil,
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		// pairScrapers: make(map[CoinflexPairScraper]nothing),
	}
	go s.mainLoop()
	return s
}

// in the websocket case, it is responsible for initiating a for-select
// loop that listens for a shutdown
func (s *CoinflexScraper) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Info("[%s] closing pair scrapers\n", dia.CoinflexExchange)
			s.cleanup(nil)
			return
		}
	}
}

// ScrapePair - creates a gorilla websocket connection (one per pair)
// and listens on the subscribed channel, OrdersMatched here
func (s *CoinflexScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	// on successfull connection, add the pair to pairScrapers
	ws, _, err := websocket.DefaultDialer.Dial(s.connectionURL.String(), nil)
	if err != nil {
		log.Warningln(err)
		return nil, err
	}
	// consider removing pong handler
	ws.SetPongHandler(func(d string) error {
		log.Debugln("pong")
		return nil
	})
	ps := CoinflexPairScraper{
		CflxPair:   pair,
		wsClient:   ws,
		chanTrades: s.Channel(),
		Closed:     false,
	}
	// add pair scraper to pairScraprs. ! Single goroutine should call
	s.pairScrapers.Store(ps, exists)
	// s.pairScrapers[ps] = exists
	// run scraping in a go routine
	go ps.scrape(ws)
	return &ps, nil
}

// Channel - gives you the trades channel
func (s *CoinflexScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// internal close
// ! note that I am returning a slice of errors here! since we can get an
// ! error from any scraper
func (s *CoinflexScraper) cleanup(err error) (errors []error) {
	// * this err error argument is a bit awkward. do we need it?
	if err != nil {
		errors = append(errors, err)
	}

	// for sp := range s.pairScrapers. {
	s.pairScrapers.Range(func(k interface{}, v interface{}) bool {
		err = k.(*CoinflexPairScraper).Close()
		if err != nil {
			log.Warnln(err)
		}
		s.pairScrapers.Delete(k)
		return true
	})

	if len(errors) > 0 {
		log.Warnf("%s unsuccessful cleanup\n", dia.CoinflexExchange)
		return
	}

	close(s.shutdown)
	close(s.shutdownDone)
	return nil // to make it explicit that everything is OK at this point
}

// Close - APIScraper io.Close implementation
// exported Close that lib users will call
func (s *CoinflexScraper) Close() (err error) {
	errors := s.cleanup(nil)

	// at this point, we have alreade logged the errors in the cleanup
	if len(errors) > 0 {
		return errors[0]
	}

	s.shutdown <- struct{}{}

	return
}

func (cp *CoinflexPairScraper) scrape(w *websocket.Conn) {
	var err error
	log.Infof("connected %v\n", cp.CflxPair.ForeignName)
	defer cp.Close() // ! errors will never be caught. Invokes wsClient.Close inside

	// sends ping messages to Coinflex to keep the connection alive
	go func() {
		tick := time.NewTicker(45 * time.Second)
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				err = w.WriteMessage(websocket.PingMessage, []byte{})
				log.Debugln("ping")
				if err != nil {
					cp.err = err
					log.Warnln(err)
					return
				}
			}
		}
	}()

	parts := strings.Split(cp.CflxPair.ForeignName, "/")
	if len(parts) != 2 {
		err = fmt.Errorf("parts has too many parts")
		cp.err = err
		return
	}
	base, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		cp.err = err
		return
	}
	quote, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		cp.err = err
		return
	}

	err = w.WriteJSON(&map[string]interface{}{
		"base":    base,
		"counter": quote,
		"watch":   true,
		"method":  "WatchOrders",
	})

	if err != nil {
		cp.err = err
		return
	}

	func() {
		var m []byte
		var rm coinflexWsWatchOrders
		var n coinflexNotice

		for {
			_, m, err = w.ReadMessage() // m stands for message
			// log.Debugf("[%s] message: %s", cp.Pair.ForeignName, string(m)) // ! consider removing
			if err != nil {
				cp.err = err
				log.Warnln(err)
				return
			}

			err = json.Unmarshal(m, &n)
			if err != nil {
				cp.err = err
				log.Warnln(err)
				return
			}

			if n.Notice != "OrdersMatched" {
				continue
			}

			err = json.Unmarshal(m, &rm)
			if err != nil {
				cp.err = err
				log.Warningln(err)
				return
			}

			// identify if it was a sell
			isSell := rm.TakerSide == "ask"
			qty := rm.Quantity
			if isSell {
				qty = -qty
			}

			dt := dia.Trade{
				Symbol:         cp.CflxPair.Symbol, // ?
				Pair:           cp.CflxPair.Symbol,
				Price:          rm.Price / coinflexScale,
				Volume:         qty, // ! negative if a sell
				Time:           time.Unix(rm.Time/1000000, 0),
				ForeignTradeID: cp.CflxPair.ForeignName,
				Source:         dia.CoinflexExchange,
			}

			log.Debugf("[%s] [%s] %+v\n", dia.CoinflexExchange, cp.CflxPair.ForeignName, dt)
			cp.chanTrades <- &dt // will block until the receiver collects it
		}
	}()
}

// Closed - tells you if this scraper is scraping any pairs
// if closed, all of the pair scrapers have been terminated.
// Create a new instance of CoinflexSctaper and use
// ScrapePair on it to add the markets to scrape
func (s *CoinflexScraper) Closed() bool {
	_, closed := <-s.shutdownDone
	return closed
}

// Close - closes websocket connection by sending
// an appropriate message, if fails, gives the error
// sets the closed flag to true if successful
func (cp *CoinflexPairScraper) Close() (err error) {
	if cp.Closed {
		return
	}

	if err != nil {
		cp.err = err
		log.Warnln(err)
	}

	defer cp.wsClient.Close() // ! no closing errors are caught

	cp.Closed = true
	return
}

func (cp *CoinflexPairScraper) Error() (err error) {
	return cp.err
}

// Pair - returns the dia.Pair that is being scraped
func (cp *CoinflexPairScraper) Pair() dia.Pair {
	return dia.Pair{
		Symbol:      cp.CflxPair.Symbol,
		ForeignName: cp.CflxPair.ForeignName,
		Exchange:    dia.CoinflexExchange,
		Ignore:      false, // * always false
	}
}

// json names have to be exportable!!!
// stands for CoinflexPair
type clxPr struct {
	Base     int64  `json:"base"`
	Quote    int64  `json:"counter"`
	Name     string `json:"name"`
	SpotName string `json:"spot_name"`
	// Start    int    `json:"start'`
	Expires int64 `json:"expires"`
	// tenor    string `json:"tenor"`
	// tick     int    `json:"tick"`
}

// FetchAvailablePairs - fetches all of the tradable pairs
// ! dia.Pair lacks information like baseID and quoteID that we need to instantiate
// ! websocket in (c *CoinflexPairScraper) scrape function
func (s *CoinflexScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	data, err := utils.GetRequest("https://webapi.coinflex.com/markets/")

	if err != nil {
		return
	}

	// stands for resolvedData
	var rd []clxPr
	err = json.Unmarshal(data, &rd)

	if err != nil {
		return
	}

	for _, pair := range rd {
		// this checks that the pair is still tradable
		tm := time.Unix(pair.Expires/1000, 0)
		currentTime := time.Now()
		ignore := currentTime.After(tm)
		if pair.Expires == 0 {
			// ! comment this conditional block if you want to ignore the spot market
			ignore = false
		}

		// rp := CoinflexPair
		rp := dia.Pair{
			Symbol:      pair.SpotName,
			ForeignName: fmt.Sprintf("%s/%s", strconv.FormatInt(pair.Base, 10), strconv.FormatInt(pair.Quote, 10)),
			Exchange:    dia.CoinflexExchange,
			Ignore:      ignore,
			// BaseID:      pair.Base,
			// QuoteID:     pair.Quote,
		}

		pairs = append(pairs, rp)
	}

	return
}
