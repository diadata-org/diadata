package scrapers

import (
	"errors"
	"math"
	"strconv"
	"sync"
	"time"

	krakenapi "github.com/beldur/kraken-go-api-client"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

const (
	krakenRefreshDelay = time.Second * 30 * 1
)

type KrakenScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*KrakenPairScraper // pc.ExchangePair -> pairScraperSet
	api          *krakenapi.KrakenApi
	ticker       *time.Ticker
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewKrakenScraper returns a new KrakenScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewKrakenScraper(key string, secret string, exchange dia.Exchange, scrape bool, relDB *models.RelDB) *KrakenScraper {
	s := &KrakenScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*KrakenPairScraper),
		api:          krakenapi.New(key, secret),
		ticker:       time.NewTicker(krakenRefreshDelay),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}
	if scrape {
		go s.mainLoop()
	}
	return s
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// func neededBalanceAdjustement(current float64, minChange float64, desired float64) (float64, string) {
// 	obj := desired - current
// 	roundedObj := Round(obj, minChange)
// 	message := fmt.Sprintf("current position: %v, min change: %v, desired position: %v, delta current/desired: %v, rounded delta: %v", current, minChange, desired, obj, roundedObj)
// 	return roundedObj, message
// }

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', -1, 64)
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *KrakenScraper) mainLoop() {
	for {
		select {
		case <-s.ticker.C:
			s.Update()
		case <-s.shutdown: // user requested shutdown
			log.Printf("KrakenScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *KrakenScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *KrakenScraper) Close() error {
	if s.closed {
		return errors.New("KrakenScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// KrakenPairScraper implements PairScraper for Kraken
type KrakenPairScraper struct {
	parent     *KrakenScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *KrakenScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("KrakenScraper: Call ScrapePair on closed scraper")
	}
	ps := &KrakenPairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0, //TODO FIX to figure out the last we got...
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

func (s *KrakenScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *KrakenScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return []dia.ExchangePair{}, errors.New("FetchAvailablePairs() not implemented")
}

// NormalizePair accounts for the par
func (ps *KrakenScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	if len(pair.ForeignName) == 7 {
		if pair.ForeignName[4:5] == "Z" || pair.ForeignName[4:5] == "X" {
			pair.ForeignName = pair.ForeignName[:4] + pair.ForeignName[5:]
			return pair, nil
		}
		if pair.ForeignName[:1] == "Z" || pair.ForeignName[:1] == "X" {
			pair.ForeignName = pair.ForeignName[1:]
		}
	}
	if len(pair.ForeignName) == 8 {
		if pair.ForeignName[4:5] == "Z" || pair.ForeignName[4:5] == "X" {
			pair.ForeignName = pair.ForeignName[:4] + pair.ForeignName[5:]
		}
		if pair.ForeignName[:1] == "Z" || pair.ForeignName[:1] == "X" {
			pair.ForeignName = pair.ForeignName[1:]
		}
	}
	if pair.ForeignName[len(pair.ForeignName)-3:] == "XBT" {
		pair.ForeignName = pair.ForeignName[:len(pair.ForeignName)-3] + "BTC"
	}
	if pair.ForeignName[:3] == "XBT" {
		pair.ForeignName = "BTC" + pair.ForeignName[len(pair.ForeignName)-3:]
	}
	return pair, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *KrakenScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *KrakenPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *KrakenPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *KrakenPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func NewTrade(pair dia.ExchangePair, info krakenapi.TradeInfo, foreignTradeID string, relDB *models.RelDB) *dia.Trade {
	volume := info.VolumeFloat
	if info.Sell {
		volume = -volume
	}
	exchangepair, err := relDB.GetExchangePairCache(dia.KrakenExchange, pair.ForeignName)
	if err != nil {
		log.Error("get exchangepair from cache: ", err)
	}
	t := &dia.Trade{
		Pair:           pair.ForeignName,
		Price:          info.PriceFloat,
		Symbol:         pair.Symbol,
		Volume:         volume,
		Time:           time.Unix(info.Time, 0),
		ForeignTradeID: foreignTradeID,
		Source:         dia.KrakenExchange,
		VerifiedPair:   exchangepair.Verified,
		BaseToken:      exchangepair.UnderlyingPair.BaseToken,
		QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
	}
	if exchangepair.Verified {
		log.Infoln("Got verified trade", t)
	}
	return t
}

func (s *KrakenScraper) Update() {

	for _, ps := range s.pairScrapers {

		r, err := s.api.Trades(ps.pair.ForeignName, ps.lastRecord)

		if err != nil {
			log.Printf("err on collect trades %v %v", err, ps.pair.ForeignName)
			time.Sleep(1 * time.Minute)
		} else {
			if r != nil {
				ps.lastRecord = r.Last
				for _, ti := range r.Trades {
					// p, _ := s.NormalizePair(ps.pair)
					t := NewTrade(ps.pair, ti, strconv.FormatInt(r.Last, 16), s.db)
					ps.parent.chanTrades <- t
				}
			} else {
				log.Printf("r nil")
			}
		}
	}
}
