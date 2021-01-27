package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
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
	pairScrapers   map[string]*STEXPairScraper
	pairSymbolToID map[string]int
	pairIDToSymbol map[int]string
	exchangeName   string
	chanTrades     chan *dia.Trade
}

// NewSTEXScraper returns a new STEXScraper for the given pair
func NewSTEXScraper(exchange dia.Exchange) *STEXScraper {
	s := &STEXScraper{
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		pairScrapers:   make(map[string]*STEXPairScraper),
		pairSymbolToID: make(map[string]int),
		pairIDToSymbol: make(map[int]string),
		exchangeName:   exchange.Name,
		error:          nil,
		chanTrades:     make(chan *dia.Trade),
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

// Subscribe again to all channels
func (s *STEXScraper) subscribeToALL() {

	// FetchAvailablePairs fetches pair IDs needed for ws subscriptions
	s.FetchAvailablePairs()

	for _, pairScraper := range s.pairScrapers {

		log.Info("pair: ", pairScraper.pair)
		a := &Channel{
			Channel: fmt.Sprintf("trade_c%s", strconv.Itoa(s.pairSymbolToID[pairScraper.pair.Symbol])),
		}

		b,_ := json.Marshal(a)
		log.Info("subscribe: ", string(b))
		if err := s.c.Emit("subscribe", a); err != nil {
			fmt.Println(err.Error())
		}

		var err error

		err = s.c.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
			log.Println("Connected")
		})
		if err != nil {
			log.Println(err)
		}

		err = s.c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
			s.closed = true
			log.Info("socket is  closed, subscribe and reconnect")
			s.reconnectToSocketIO()
		})

		err = s.c.On("App\\Events\\OrderFillCreated", func(h *gosocketio.Channel, message Trade) {
			log.Info("Handle new trades...")
			var pairID = message.CurrencyPairID
			var forName = s.pairIDToSymbol[int(pairID)]
			ps, ok := s.pairScrapers[forName]

			if ok {
				go func() {
					f64Price, _ := strconv.ParseFloat(message.Price, 64)
					f64Volume := message.Amount2
					timeStamp := time.Now().UTC()

					if message.OrderType == "SELL" {
						f64Volume = -f64Volume
					}
					// element id is more than int64/uint64 in size
					// leave the id in float64 format
					t := &dia.Trade{
						Symbol:         ps.pair.Symbol,
						Pair:           forName,
						Price:          f64Price,
						Volume:         f64Volume,
						Time:           timeStamp,
						ForeignTradeID: fmt.Sprintf("%d", message.ID),
						Source:         s.exchangeName,
					}
					ps.parent.chanTrades <- t
					log.Info("got trade: ", t)
				}()
			} else {
				log.Printf("Unknown Pair %v", forName)
			}
		})
		if err != nil {
			log.Println(err)
		}
	}

}

// runs in a goroutine until s is closed
func (s *STEXScraper) mainLoop() {
	log.Info("mainLoop() waiting for pairs to be added...")
	time.Sleep(10 * time.Second)

	if s.closed {
		log.Info("s closed, subscribe and reconnect")
		s.reconnectToSocketIO()
	} else {
		s.subscribeToALL()
	}

	s.cleanup(errors.New("main loop terminated by Close()"))
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
	log.Println("sending req")
	data, err := utils.GetRequest("https://api3.stex.com/public/currency_pairs/list/ALL")
	if err != nil {
		return
	}
	log.Println("sending req done")
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
