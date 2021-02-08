package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance"
	"github.com/diadata-org/diadata/pkg/dia"
	utils "github.com/diadata-org/diadata/pkg/utils"
)

type binancePairScraperSet map[*BinancePairScraper]nothing

// BinanceScraper is a Scraper for collecting trades from the Binance websocket API
type BinanceScraper struct {
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
	pairScrapers      sync.Map // dia.Pair -> binancePairScraperSet
	pairSubscriptions sync.Map // dia.Pair -> string (subscription ID)
	pairLocks         sync.Map // dia.Pair -> sync.Mutex
	exchangeName      string
	chanTrades        chan *dia.Trade
}

// NewBinanceScraper returns a new BinanceScraper for the given pair
func NewBinanceScraper(apiKey string, secretKey string, exchange dia.Exchange) *BinanceScraper {

	s := &BinanceScraper{
		client:       binance.NewClient(apiKey, secretKey),
		initDone:     make(chan nothing),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	// establish connection in the background
	go s.mainLoop()
	return s
}

func eventHandler(event *binance.WsAggTradeEvent) {
	fmt.Println(event)

}

func errorHandler(err error) {
	fmt.Println(err)

}

// runs in a goroutine until s is closed
func (s *BinanceScraper) mainLoop() {
	close(s.initDone)
	for {
		select {
		case <-s.shutdown: // user requested shutdown
			log.Println("BinanceScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *BinanceScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	// close all channels of PairScraper children
	s.pairScrapers.Range(func(k, v interface{}) bool {
		for ps := range v.(binancePairScraperSet) {
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
func (s *BinanceScraper) Close() error {
	if s.closed {
		return errors.New("BinanceScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BinanceScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	<-s.initDone // wait until client is connected

	if s.closed {
		return nil, errors.New("BinanceScraper: Call ScrapePair on closed scraper")
	}

	ps := &BinancePairScraper{
		parent: s,
		pair:   pair,
	}

	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {

		volume, err := strconv.ParseFloat(event.Quantity, 64)
		price, err2 := strconv.ParseFloat(event.Price, 64)

		if err == nil && err2 == nil && event.Event == "aggTrade" {
			if event.IsBuyerMaker == false {
				volume = -volume
			}
			pairNormalized, _ := s.NormalizePair(pair)
			t := &dia.Trade{
				Symbol:         pairNormalized.Symbol,
				Pair:           pairNormalized.ForeignName,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(event.TradeTime/1000, (event.TradeTime%1000)*int64(time.Millisecond)),
				ForeignTradeID: strconv.FormatInt(event.AggTradeID, 16),
				Source:         s.exchangeName,
			}
			ps.parent.chanTrades <- t
			log.Info("got trade: ", t)
		} else {
			log.Println("ignoring event ", event, err, err2)
		}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}

	_, _, err := binance.WsAggTradeServe(pair.ForeignName, wsAggTradeHandler, errHandler)

	return ps, err
}

// NormalizePair description
func (s *BinanceScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	if pair.Symbol == "RENBTC" {
		pair.Symbol = "renBTC"
		pair.ForeignName = "renBTC" + pair.ForeignName[6:]
		return pair, nil
	}
	if pair.Symbol == "WNXM" {
		pair.Symbol = "wNXM"
		pair.ForeignName = "wNXM" + pair.ForeignName[4:]
		return pair, nil
	}
	if pair.Symbol == "IOTA" {
		pair.Symbol = "MIOTA"
		pair.ForeignName = "M" + pair.ForeignName
		return pair, nil
	}
	if pair.Symbol == "YOYO" {
		pair.Symbol = "YOYOW"
		pair.ForeignName = "YOYOW" + pair.ForeignName[4:]
		return pair, nil
	}
	if pair.Symbol == "BQX" {
		pair.Symbol = "ETHOS"
		pair.ForeignName = "ETHOS" + pair.ForeignName[3:]
		return pair, nil
	}
	return pair, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BinanceScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	data, err := utils.GetRequest("https://api.binance.com/api/v1/exchangeInfo")

	if err != nil {
		return
	}
	var ar binance.ExchangeInfo
	var pair dia.Pair
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar.Symbols {
			if p.Status == "TRADING" {
				pair = dia.Pair{
					Symbol:      p.BaseAsset,
					ForeignName: p.Symbol,
					Exchange:    s.exchangeName,
					Ignore:      false,
				}
				pairs = append(pairs, pair)
			}
		}
	}
	return
}

// BinancePairScraper implements PairScraper for Binance
type BinancePairScraper struct {
	parent *BinanceScraper
	pair   dia.Pair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *BinancePairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("BinancePairScraper: Already closed")
	}

	// TODO stop collection for the pair

	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *BinanceScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BinancePairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BinancePairScraper) Pair() dia.Pair {
	return ps.pair
}
