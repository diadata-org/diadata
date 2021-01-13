package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	utils "github.com/diadata-org/diadata/pkg/utils"
)

// API base url
const apiURL string = "https://bitbay.net/API/Public/"

// API request
const apiRequest string = "/trades.json"

// Minimum delay between API calls
const apiDelay = time.Second

// Seconds to wait for scrappers to be ready
const waitForScrapers = 10

//TradeInfo as received from API response
type TradeInfo struct {
	Date            int64   `json:"date"`
	Price           float64 `json:"price"`
	TransactionType string  `json:"type"`
	Amount          float64 `json:"amount"`
	Tid             string  `json:"tid"`
}

// BitBayScraper provides  methods needed to get Trade information from BitBay
type BitBayScraper struct {
	// control flag for main loop
	run bool
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers []*BitBayPairScraper
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
		pairScrapers: make([]*BitBayPairScraper, 0),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		closed:       false,
	}
	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *BitBayScraper) mainLoop() {
	// wait up to 10 [s] for scrapers
	waitForScrapers := 10
	s.run = false
	s.error = errors.New(s.exchangeName + "Scraper: No pair to scrap provided")
	log.Info("Waiting for pairs")
	for waitForScrapers > 0 {
		log.Info(".")
		if len(s.pairScrapers) > 0 {
			// If at least one scraper is ready we can start querying the API
			log.Info("Pair availables")
			s.error = nil
			s.run = true
			break
		}
		time.Sleep(time.Second)
		waitForScrapers--
	}
	for s.run {
		for _, pair := range s.pairScrapers {
			// time.Sleep(apiDelay)
			trades, err := fetchTrades(pair.apiEndPoint, strconv.Itoa(pair.latestTrade))
			if err == nil {
				for _, trade := range trades {
					i, err := strconv.Atoi(trade.Tid)
					if err == nil {
						if i > pair.latestTrade {
							pair.latestTrade = i
							t := &dia.Trade{
								Symbol:         pair.pair.Symbol,
								Pair:           pair.apiEndPoint,
								Price:          trade.Price,
								Volume:         trade.Amount,
								Time:           time.Unix(trade.Date, 0),
								ForeignTradeID: trade.Tid,
								Source:         s.exchangeName,
							}
							log.Info("got trade: ", t)
							pair.parent.chanTrades <- t
						}
					} else {
						log.Error(s.exchangeName + "Scrapper:Error converting trade id:" + trade.Tid + " into int")
					}
				}
			} else {
				log.Error("Error fetching pairs:", err)
			}
		}
		time.Sleep(1 * time.Second)
	}
	if s.error == nil {
		s.error = errors.New(s.exchangeName + "Scraper: terminated by Close()")
	}
	s.cleanup(s.error)
}

// retrieve pair data
func fetchTrades(pairs string, latestTrade string) ([]TradeInfo, error) {

	log.Info("requesting:" + apiURL + pairs + apiRequest + "?since=" + latestTrade)

	body, err := utils.GetRequest(apiURL + pairs + apiRequest + "?since=" + latestTrade)
	type APIResponse []TradeInfo
	var ar APIResponse

	if err != nil {
		return ar, err
	}

	jsonErr := json.Unmarshal(body, &ar)
	if jsonErr != nil {
		log.Error("Error unmarshalling:", jsonErr)
		return nil, jsonErr
	}
	return ar, err
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
	s.pairScrapers = append(s.pairScrapers, ps)
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
	return []dia.Pair{}, errors.New(s.exchangeName + "Scraper:FetchAvailablePairs() not implemented. No public API available")
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
