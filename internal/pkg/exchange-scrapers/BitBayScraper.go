package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
)

// API base url
const apiURL string = "https://bitbay.net/API/Public/"

// API request
const apiRequest string = "/trades.json"

// Minimum delay between API calls
const apiDelay = time.Second

// Seconds to wait for scrappers to be ready
const waitForScrapers = 10

var BitBaySocketURL string = "wss://api.bitbay.net/websocket/"

type BitBaySubscribe struct {
	Action string `json:"action"`
	Module string `json:"module"`
	Path   string `json:"path"`
}

//{
//"action": "subscribe-public",
//"module": "trading",
//"path": "ticker/{market_code}"
//}

//TradeInfo as received from API response
type BitBayTrade struct {
	Date            int64   `json:"date"`
	Price           float64 `json:"price"`
	TransactionType string  `json:"type"`
	Amount          float64 `json:"amount"`
	Tid             string  `json:"tid"`
}

// BitBayScraper provides  methods needed to get Trade information from BitBay
type BitBayScraper struct {
	// control flag for main loop
	run      bool
	wsClient *ws.Conn

	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*BitBayPairScraper
	// exchange name
	exchangeName string
	// channel to send trades
	chanTrades chan *dia.Trade
}

//NewBitBayScraper get a scrapper for BitBay exchange
func NewBitBayScraper(exchange dia.Exchange) *BitBayScraper {
	s := &BitBayScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*BitBayPairScraper),

		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		closed:       false,
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(BitBaySocketURL, nil)

	if err != nil {
		println(err.Error())
	}

	s.wsClient = SwConn

	go s.mainLoop()

	return s
}

func (s *BitBayScraper) getMarkets() (markets []string) {
	var bbm BitBayMarkets
	b, err := utils.GetRequest("https://api.bitbay.net/rest/trading/ticker")
	if err != nil {
		log.Errorln("Error Getting markets", err)
	}
	json.Unmarshal(b, &bbm)

	for key, _ := range bbm.Items {
		markets = append(markets, key)
	}
	return
}

func (s *BitBayScraper) ping() {

	a := &BitBaySubscribe{
		Action: "ping",
	}

	log.Infoln("Ping", a)

	if err := s.wsClient.WriteJSON(a); err != nil {
		log.Println(err.Error())
	}
}

func (s *BitBayScraper) subscribe() {

	markets := s.getMarkets()

	for _, market := range markets {

		a := &BitBaySubscribe{
			Action: "subscribe-public",
			Module: "trading",
			Path:   "transactions/" + market,
		}

		log.Println("subscribing", a)

		if err := s.wsClient.WriteJSON(a); err != nil {
			log.Println(err.Error())
		}

	}
}

// runs in a goroutine until s is closed
func (s *BitBayScraper) mainLoop() {

	s.subscribe()

	pingTimer := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-pingTimer.C:
				go s.ping()
			}
		}
	}()

	for {

		var response BitBayWSResponse

		if s.error = s.wsClient.ReadJSON(&response); s.error != nil {
			log.Error(s.error.Error())
			break
		}

		//b,_ := json.Marshal(message)
		//
		//log.Infoln("Message",string(b[:]))

		if len(response.Message.Transactions) == 0 {
			log.Warn("empty message - continue")
			continue
		}

		timestamp, err := strconv.ParseInt(response.Timestamp, 10, 64)
		if err != nil {
			log.Error("Error Parsing time", err)
		}

		pair := strings.TrimPrefix(response.Topic, "trading/transactions/")
		if response.Topic == "" {
			log.Warn("empty response - continue.")
			continue
		}
		pair = strings.Replace(pair, "-", "", -1)
		pair = strings.ToUpper(pair)

		ps, ok := s.pairScrapers[pair]
		if !ok {
			log.Error("unknown pair: " + pair)
			continue
		}

		for _, trade := range response.Message.Transactions {

			f64Price, err := strconv.ParseFloat(trade.R, 64)
			if err != nil {
				log.Error("error parsing price: " + trade.R)
				continue
			}

			f64Volume, err := strconv.ParseFloat(trade.A, 64)
			if err != nil {
				log.Error("error parsing volume: " + trade.A)
				continue
			}

			if trade.Ty == "Sell" {
				f64Volume = -f64Volume
			}

			t := &dia.Trade{
				Symbol:         ps.Pair().Symbol,
				Pair:           pair,
				Price:          f64Price,
				Volume:         f64Volume,
				Time:           time.Unix(timestamp/1e3, 0),
				ForeignTradeID: trade.ID,
				Source:         s.exchangeName,
			}
			log.Info("got trade: ", t)
			ps.parent.chanTrades <- t
		}
	}
	if s.error == nil {
		s.error = errors.New(s.exchangeName + "Scraper: terminated by Close()")
	}
	s.cleanup(s.error)
}

// Close channels for shutdown
func (s *BitBayScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.chanTrades)
	close(s.shutdownDone)
}

// Close any existing API connections, as well as channels, and terminates main loop
func (s *BitBayScraper) Close() error {
	if s.closed {
		return errors.New(s.exchangeName + "Scraper: Already closed")
	}
	s.run = false
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BitBayScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	if s.closed {
		return nil, errors.New("s.exchangeName+Scraper: Call ScrapePair on closed scraper")
	}
	ps := &BitBayPairScraper{
		parent:      s,
		pair:        pair,
		apiEndPoint: pair.ForeignName,
		latestTrade: 0,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// set exchange base currency according to DIA standard
func (s *BitBayScraper) normalizeSymbol(baseCurrency string, name string) (symbol string, err error) {
	return "", errors.New(s.exchangeName + "Scraper:normalizeSymbol() not implemented.")
}

func (s *BitBayScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	return dia.Pair{}, nil
}

//Channel returns the channel to get trades
func (s *BitBayScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

//FetchAvailablePairs returns a list with all available trade pairs
func (s *BitBayScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	type items struct {
		Status  string                 `json:"status"`
		Markets map[string]interface{} `json:"items"`
	}
	var bitbayResponse items

	data, err := utils.GetRequest("https://api.bitbay.net/rest/trading/ticker")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &bitbayResponse)
	if err != nil {
		return
	}
	pairmap := bitbayResponse.Markets
	for key := range pairmap {
		pairslice := strings.Split(key, "-")
		pairs = append(pairs, dia.Pair{
			Symbol:      pairslice[0],
			ForeignName: pairslice[0] + pairslice[1],
			Exchange:    s.exchangeName,
			Ignore:      false,
		})
	}
	return pairs, err
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (s *BitBayScraper) Error() error {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// BitBayPairScraper implements PairScraper for BitBay
type BitBayPairScraper struct {
	apiEndPoint string
	parent      *BitBayScraper
	pair        dia.Pair
	closed      bool
	latestTrade int
}

// Close stops listening for trades of the pair associated
func (ps *BitBayPairScraper) Close() error {
	ps.closed = true
	return ps.Error()
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BitBayPairScraper) Error() error {
	ps.parent.errorLock.RLock()
	defer ps.parent.errorLock.RUnlock()
	return ps.parent.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BitBayPairScraper) Pair() dia.Pair {
	return ps.pair
}

type BitBayMarkets struct {
	Status string                  `json:"status"`
	Items  map[string]BitBayMarket `json:"items"`
}

type BitBayMarket struct {
	Code  string `json:"code"`
	First struct {
		Currency string `json:"currency"`
		MinOffer string `json:"minOffer"`
		Scale    int    `json:"scale"`
	} `json:"first"`
	Second struct {
		Currency string `json:"currency"`
		MinOffer string `json:"minOffer"`
		Scale    int    `json:"scale"`
	} `json:"second"`
}

type BitBayWSResponse struct {
	Action    string  `json:"action"`
	Message   Message `json:"message"`
	SeqNo     int     `json:"seqNo"`
	Timestamp string  `json:"timestamp"`
	Topic     string  `json:"topic"`
}
type BitBayTransaction struct {
	A  string `json:"a"`
	ID string `json:"id"`
	R  string `json:"r"`
	T  string `json:"t"`
	Ty string `json:"ty"`
}
type Message struct {
	Transactions []BitBayTransaction `json:"transactions"`
}
