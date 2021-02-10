package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	utils "github.com/diadata-org/diadata/pkg/utils"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var _socketURL string = "socket.stex.com"

const (
	apiBaseURL = "https://api3.stex.com/public"
)

type Trade struct {
	Amount         string      `json:"amount"`
	Amount2        float64     `json:"amount2"`
	BuyOrderID     int64       `json:"buy_order_id"`
	CurrencyPairID int64       `json:"currency_pair_id"`
	Date           string      `json:"date"`
	ID             int64       `json:"id"`
	OrderType      string      `json:"order_type"`
	Price          string      `json:"price"`
	SellOrderID    int64       `json:"sell_order_id"`
	Socket         interface{} `json:"socket"`
	Timestamp      int64       `json:"timestamp"`
}

type STEXTrade struct {
	ID        int        `json:"id"`
	Price     *big.Float `json:"price"`
	Amount    *big.Float `json:"amount"`
	Type      string     `json:"type"`
	TimeStamp string     `json:"timestamp"`
}
type Channel struct {
	Channel string `json:"channel"`
}

type STEXScraper struct {
	c *gosocketio.Client
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers      map[string]*STEXPairScraper
	pairSymbolToID    map[string]int
	pairLastTimeStamp map[string]time.Time
	pairIDToSymbol    map[int]string
	exchangeName      string
	chanTrades        chan *dia.Trade
}

// NewSTEXScraper returns a new STEXScraper for the given pair
func NewSTEXScraper(exchange dia.Exchange) *STEXScraper {
	s := &STEXScraper{
		shutdown:          make(chan nothing),
		shutdownDone:      make(chan nothing),
		pairScrapers:      make(map[string]*STEXPairScraper),
		pairSymbolToID:    make(map[string]int),
		pairIDToSymbol:    make(map[int]string),
		pairLastTimeStamp: make(map[string]time.Time),
		exchangeName:      exchange.Name,
		error:             nil,
		chanTrades:        make(chan *dia.Trade),
	}

	c, err := gosocketio.Dial(
		gosocketio.GetUrl(_socketURL, 443, true),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Printf("dial: %v", err)
	}
	s.c = c
	go s.mainLoop()
	return s
}

// Reconnect to socketIO when the connection is down.
func (s *STEXScraper) reconnectToSocketIO() {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(_socketURL, 443, true),
		transport.GetDefaultWebsocketTransport())
	if err != nil {
		log.Printf("dial: %v", err)
	} else {
		log.Info("successfully reconnected.")
	}
	s.c = c
}

type StexTradeResponse struct {
	SETXTrades []STEXTrade `json:"data"`
	Success    bool        `json:"success"`
}

// runs in a goroutine until s is closed
func (s *STEXScraper) mainLoop() {
	log.Info("mainLoop() waiting for pairs to be added...")
	time.Sleep(10 * time.Second)
	for {
		s.scrapeTrades()
	}
}

func (s *STEXScraper) scrapeTrades() {
	var numRequests int
	s.FetchAvailablePairs()
	for _, pairScraper := range s.pairScrapers {
		if numRequests > 180 {
			// API limit is 180 requests per min.
			log.Info("sleep for a minute due to STEX API rate limit")
			time.Sleep(1 * time.Minute)
			numRequests = 0
		} else {
			s.scrapePair(pairScraper.pair)
			numRequests++
		}
	}
	// Sleep after getting trades for all pairs
	log.Info("Scraped all pairs. Wait for next iteration.")
	time.Sleep(4 * time.Second)
}

// scrapePair scrapes the @pair associated to s.pairScraper
func (s *STEXScraper) scrapePair(pair dia.Pair) {

	if (s.pairLastTimeStamp[pair.Symbol] == time.Time{}) {
		// Set last trade time to 10 mins ago for initial run
		s.pairLastTimeStamp[pair.Symbol] = time.Now().Add(-10 * time.Minute)
	}
	trades, _ := s.GetNewTrades(strconv.Itoa(s.pairSymbolToID[pair.Symbol]), s.pairLastTimeStamp[pair.Symbol])
	for _, trade := range trades {

		f64Price, _ := trade.Price.Float64()
		f64Volume, _ := trade.Amount.Float64()
		timeStamp, _ := strconv.ParseInt(trade.TimeStamp, 10, 32)

		if trade.Type == "SELL" {
			f64Volume = -f64Volume
		}

		t := &dia.Trade{
			Symbol:         strings.Split(pair.Symbol, "-")[0],
			Pair:           pair.ForeignName,
			Price:          f64Price,
			Volume:         f64Volume,
			Time:           time.Unix(timeStamp, 0),
			ForeignTradeID: strconv.Itoa(trade.ID),
			Source:         s.exchangeName,
		}
		s.chanTrades <- t
		log.Info("got trade: ", t)
	}
}

// GetNewTrades fetches new trades from the STEX restAPI dating back until @fromTimestamp
func (s *STEXScraper) GetNewTrades(pairID string, fromTimestamp time.Time) ([]STEXTrade, error) {
	var (
		response StexTradeResponse
		err      error
		bytes    []byte
		url      string
	)
	if (fromTimestamp == time.Time{}) {
		url = apiBaseURL + "/trades/" + pairID + "?sort=DESC&limit=100"
	} else {
		unixTime := strconv.Itoa(int(fromTimestamp.Unix()))
		url = apiBaseURL + "/trades/" + pairID + "?sort=DESC&from=" + unixTime + "&limit=100"
	}

	bytes, err = utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &response)
	// Update timestamp
	pairIDInt, _ := strconv.ParseInt(pairID, 10, 32)
	symbol := s.pairIDToSymbol[int(pairIDInt)]

	if len(response.SETXTrades) > 0 {
		var lastTimestamp int64
		lastTimestamp, err = strconv.ParseInt(response.SETXTrades[0].TimeStamp, 10, 64)
		if err != nil {
			log.Error("error parsing trade's timestamp")
		}
		s.pairLastTimeStamp[symbol] = time.Unix(lastTimestamp+1, 0)
	}

	return response.SETXTrades, nil
}

func (s *STEXScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *STEXScraper) Close() error {
	if s.closed {
		return errors.New("STEXScraper: Already closed")
	}
	close(s.shutdown)
	s.c.Close()
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *STEXScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("STEXScraper: Call ScrapePair on closed scraper")
	}

	ps := &STEXPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps
	a := &Channel{
		Channel: fmt.Sprintf("trade_c%s", pair.ForeignName),
	}
	if err := s.c.Emit("subscribe", a); err != nil {
		fmt.Println(err.Error())
	}

	return ps, nil
}

func (s *STEXScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	symbol := strings.ToUpper(pair.Symbol)
	pair.Symbol = symbol
	if helpers.SymbolIsBlackListed(symbol) {
		return pair, errors.New("Symbol is black listed:" + symbol)
	}
	return pair, nil

}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *STEXScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	type CurrencyPairs struct {
		Success bool `json:"success"`
		Data    []struct {
			ID                int    `json:"id"`
			CurrencyID        int    `json:"currency_id"`
			CurrencyCode      string `json:"currency_code"`
			CurrencyName      string `json:"currency_name"`
			MarketCurrencyID  int    `json:"market_currency_id"`
			MarketCode        string `json:"market_code"`
			MarketName        string `json:"market_name"`
			MinOrderAmount    string `json:"min_order_amount"`
			MinBuyPrice       string `json:"min_buy_price"`
			MinSellPrice      string `json:"min_sell_price"`
			BuyFeePercent     string `json:"buy_fee_percent"`
			SellFeePercent    string `json:"sell_fee_percent"`
			Active            bool   `json:"active"`
			Delisted          bool   `json:"delisted"`
			PairMessage       string `json:"pair_message"`
			CurrencyPrecision int    `json:"currency_precision"`
			MarketPrecision   int    `json:"market_precision"`
			Symbol            string `json:"symbol"`
			GroupName         string `json:"group_name"`
			GroupID           int    `json:"group_id"`
			AmountMultiplier  int    `json:"amount_multiplier"`
		} `json:"data"`
	}
	data, err := utils.GetRequest("https://api3.stex.com/public/currency_pairs/list/ALL")
	if err != nil {
		return
	}
	var response CurrencyPairs
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	if !response.Success {
		return nil, errors.New("unsuccessful attempt to get currency pairs on STEX exchange")
	}
	for _, p := range response.Data {
		pairToNormalize := dia.Pair{
			Symbol:      p.CurrencyCode,
			ForeignName: p.Symbol,
			Exchange:    s.exchangeName,
		}
		pair, serr := s.NormalizePair(pairToNormalize)
		if serr == nil {
			pairs = append(pairs, pair)
			s.pairSymbolToID[pair.Symbol] = p.ID
			s.pairIDToSymbol[p.ID] = pair.Symbol
		} else {
			log.Println(serr)
		}
	}
	return
}

// STEXPairScraper implements PairScraper for STEX
type STEXPairScraper struct {
	parent *STEXScraper
	pair   dia.Pair
	id     int
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *STEXPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *STEXScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *STEXPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *STEXPairScraper) Pair() dia.Pair {
	return ps.pair
}
