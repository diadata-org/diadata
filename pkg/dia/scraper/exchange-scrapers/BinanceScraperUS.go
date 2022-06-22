package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/cryptwire/go-binance/v2"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
)

const (
	BinanceUSWsURL = "wss://stream.binance.us:9443/ws"
)

type BinanceUSPairScraperSet map[*BinanceUSPairScraper]nothing

// BinanceScraperUS is a Scraper for collecting trades from the Binance websocket API
type BinanceScraperUS struct {
	client *binance.Client
	// signaling channels for session initialization and finishing
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	// use sync.Maps to concurrently handle multiple pairs
	pairScrapers sync.Map // dia.ExchangePair -> BinanceUSPairScraperSet
	// pairSubscriptions sync.Map // dia.ExchangePair -> string (subscription ID)
	// pairLocks         sync.Map // dia.ExchangePair -> sync.Mutex
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewBinanceScraperUS returns a new BinanceScraperUS for the given pair
func NewBinanceScraperUS(apiKey string, secretKey string, exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BinanceScraperUS {
	binance.BaseWsMainURL = BinanceUSWsURL

	s := &BinanceScraperUS{
		client:       binance.NewClient(apiKey, secretKey),
		initDone:     make(chan nothing),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	// establish connection in the background
	if scrape {
		go s.mainLoop()
	}
	return s
}

func (up *BinanceScraperUS) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	if pair.Symbol == "MIOTA" {
		pair.ForeignName = "M" + pair.ForeignName
	}
	if pair.Symbol == "YOYOW" {
		pair.ForeignName = "YOYOW" + pair.ForeignName[4:]
	}
	if pair.Symbol == "ETHOS" {
		pair.ForeignName = "ETHOS" + pair.ForeignName[3:]
	}
	if pair.Symbol == "WNXM" {
		pair.Symbol = "wNXM"
		pair.ForeignName = "wNXM" + pair.ForeignName[4:]
	}
	return pair, nil
}

// runs in a goroutine until s is closed
func (s *BinanceScraperUS) mainLoop() {
	close(s.initDone)
	for range s.shutdown { // user requested shutdown
		log.Println("BinanceScraperUS shutting down")
		s.cleanup()
		return
	}
	select {}
}

func (s *BinanceScraperUS) FillSymbolData(symbol string) (dia.Asset, error) {
	// TO DO
	return dia.Asset{Symbol: symbol}, nil
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *BinanceScraperUS) cleanup() {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	// close all channels of PairScraper children
	s.pairScrapers.Range(func(k, v interface{}) bool {
		for ps := range v.(BinanceUSPairScraperSet) {
			ps.closed = true
		}
		s.pairScrapers.Delete(k)
		return true
	})

	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BinanceScraperUS) Close() error {
	if s.closed {
		return errors.New("BinanceScraperUS: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BinanceScraperUS) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	<-s.initDone // wait until client is connected

	if s.closed {
		return nil, errors.New("BinanceScraperUS: Call ScrapePair on closed scraper")
	}

	ps := &BinanceUSPairScraper{
		parent: s,
		pair:   pair,
	}

	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
		var exchangepair dia.ExchangePair

		volume, err := strconv.ParseFloat(event.Quantity, 64)
		price, err2 := strconv.ParseFloat(event.Price, 64)

		if err == nil && err2 == nil && event.Event == "aggTrade" {
			if !event.IsBuyerMaker {
				volume = -volume
			}
			pairNormalized, _ := s.NormalizePair(pair)
			exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, pair.ForeignName)
			if err != nil {
				log.Error(err)
			}
			t := &dia.Trade{
				Symbol:         pairNormalized.Symbol,
				Pair:           pairNormalized.ForeignName,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(event.TradeTime/1000, (event.TradeTime%1000)*int64(time.Millisecond)),
				ForeignTradeID: strconv.FormatInt(event.AggTradeID, 16),
				Source:         s.exchangeName,
				VerifiedPair:   exchangepair.Verified,
				BaseToken:      exchangepair.UnderlyingPair.BaseToken,
				QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
			}
			if exchangepair.Verified {
				log.Infoln("Got verified trade", t)
			}
			ps.parent.chanTrades <- t
		} else {
			log.Println("ignoring event ", event, err, err2)
		}
	}
	errHandler := func(err error) {
		log.Error(err)
	}

	_, _, err := binance.WsAggTradeServe(pair.ForeignName, wsAggTradeHandler, errHandler)
	if err != nil {
		log.Errorf("serving pair %s", pair.ForeignName)
	}

	return ps, err
}
func (s *BinanceScraperUS) normalizeSymbol(p dia.ExchangePair, foreignName string, params ...string) (pair dia.ExchangePair, err error) {
	pair = p
	return pair, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BinanceScraperUS) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	data, _, err := utils.GetRequest("https://api.binance.us/api/v1/exchangeInfo")

	if err != nil {
		return
	}
	var ar binance.ExchangeInfo
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar.Symbols {
			pairToNormalise := dia.ExchangePair{
				Symbol:      p.BaseAsset,
				ForeignName: p.Symbol,
				Exchange:    s.exchangeName,
			}
			pair, serr := s.normalizeSymbol(pairToNormalise, p.BaseAsset, p.Status)
			if serr == nil {
				pairs = append(pairs, pair)
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// BinanceUSPairScraper implements PairScraper for Binance
type BinanceUSPairScraper struct {
	parent *BinanceScraperUS
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *BinanceUSPairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("BinanceUSPairScraper: Already closed")
	}

	// TODO stop collection for the pair

	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *BinanceScraperUS) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BinanceUSPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BinanceUSPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
