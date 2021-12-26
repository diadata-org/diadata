package scrapers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

var ByBitSocketURL string = "wss://stream.bybit.com/realtime"

// ByBitScraper provides  methods needed to get Trade information from ByBit
type ByBitScraper struct {
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
	pairScrapers map[string]*ByBitPairScraper
	// exchange name
	exchangeName string
	// channel to send trades
	chanTrades chan *dia.Trade
	db         *models.RelDB
}

//NewByBitScraper get a scrapper for ByBit exchange
func NewByBitScraper(key string, secret string, exchange dia.Exchange, scrape bool, relDB *models.RelDB) *ByBitScraper {
	s := &ByBitScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*ByBitPairScraper),

		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		closed:       false,
		db:           relDB,
	}

	// Create HMAC instance from the secret key
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	apiSecretBytes := []byte(secret)
	// Generate expires.
	expires := int((time.Now().UnixNano() + 1) * 1000)
	expiresBytes := []byte(fmt.Sprintf("GET/realtime%d", expires))
	data := append(apiSecretBytes, expiresBytes...)
	h.Write([]byte(data))

	// Get the signature
	signature := hex.EncodeToString(h.Sum(nil))

	// Create the ws connection
	var wsDialer ws.Dialer

	// Generate the ws url.
	params := fmt.Sprintf("api_key=%s&expires=%d&signature=%s", secret, expires, signature)
	SwConn, _, err := wsDialer.Dial(ByBitSocketURL+"?"+params, nil)
	if err != nil {
		println(err.Error())
	}

	s.wsClient = SwConn

	if scrape {
		go s.mainLoop()
	}

	return s
}

// runs in a goroutine until s is closed
func (s *ByBitScraper) mainLoop() {
	// TODO
}

// Close channels for shutdown
func (s *ByBitScraper) cleanup(err error) {
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
func (s *ByBitScraper) Close() error {
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
func (s *ByBitScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if s.closed {
		return nil, errors.New("s.exchangeName+Scraper: Call ScrapePair on closed scraper")
	}
	ps := &ByBitPairScraper{
		parent:      s,
		pair:        pair,
		apiEndPoint: pair.ForeignName,
		latestTrade: 0,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (s *ByBitScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

//Channel returns the channel to get trades
func (s *ByBitScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *ByBitScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	// TO DO
	return dia.Asset{Symbol: symbol}, nil
}

//FetchAvailablePairs returns a list with all available trade pairs
func (s *ByBitScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return nil, nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (s *ByBitScraper) Error() error {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ByBitPairScraper implements PairScraper for ByBit
type ByBitPairScraper struct {
	apiEndPoint string
	parent      *ByBitScraper
	pair        dia.ExchangePair
	closed      bool
	latestTrade int
}

// Close stops listening for trades of the pair associated
func (ps *ByBitPairScraper) Close() error {
	ps.closed = true
	return ps.Error()
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *ByBitPairScraper) Error() error {
	ps.parent.errorLock.RLock()
	defer ps.parent.errorLock.RUnlock()
	return ps.parent.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *ByBitPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
