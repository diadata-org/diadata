package scrapers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/carterjones/signalr"
	"github.com/carterjones/signalr/hubs"
	"github.com/diadata-org/diadata/pkg/dia"
)

type CREX24ApiInstrument struct {
	Symbol       string `json:"symbol"`
	BaseCurrency string `json:"baseCurrency"`
}

type CREX24ApiTrade struct {
	Price     string `json:"P"`
	Volume    string `json:"V"`
	Side      string `json:"S"`
	Timestamp int64  `json:"T"`
}

type CREX24ApiTradeUpdate struct {
	I  string
	NT []CREX24ApiTrade
}

// CREX24 Scraper is a scraper for collecting trades from the CREX24 signalR api
type CREX24Scraper struct {
	client *signalr.Client
	msgId  int
	// signalR Connection
	connectLock sync.RWMutex
	connected   bool
	// closed
	closed bool
	// pair scrapers
	pairScrapers map[string]*CREX24PairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
}

func NewCREX24Scraper(exchange dia.Exchange) *CREX24Scraper {
	s := &CREX24Scraper{
		pairScrapers: make(map[string]*CREX24PairScraper),
		exchangeName: exchange.Name,
		chanTrades:   make(chan *dia.Trade),
		connected:    false,
		closed:       false,
		msgId:        1,
	}
	return s
}

// connect establishes a signalR connection to CREX24
func (s *CREX24Scraper) connect() error {
	s.client = signalr.New(
		"api.crex24.com",
		"1.5",
		"/signalr",
		`[{"name": "publicCryptoHub"}]`,
		nil,
	)
	msgHandler := s.handleMessage
	panicIfErr := func(err error) {
		if err != nil {
			log.Panic(err)
		}
	}
	err := s.client.Run(msgHandler, panicIfErr)
	if err != nil {
		return errors.New("CREX24Scraper: Could not establish signalR connection")
	}
	s.connected = true
	return nil
}

func (s *CREX24Scraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *CREX24Scraper) Close() error {
	if s.closed {
		return errors.New("CREX24Scraper: Already closed")
	}
	s.closed = true
	s.connectLock.Lock()
	defer s.connectLock.Unlock()
	if s.connected {
		s.client.Close()
		s.connected = false
	}
	return nil
}

func (s *CREX24Scraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	return dia.Pair{}, nil
}

func (s *CREX24Scraper) handleMessage(msg signalr.Message) {
	if s.closed {
		return
	}
	for _, data := range msg.M {
		if data.M == "UpdateSource" {
			arguments := data.A
			payload, ok := arguments[1].(string)
			var parsedUpdate CREX24ApiTradeUpdate
			if ok {
				json.NewDecoder(strings.NewReader(payload)).Decode(&parsedUpdate)
				s.sendTradesToChannel(&parsedUpdate)
			}
		}
	}
}

func (s *CREX24Scraper) sendTradesToChannel(update *CREX24ApiTradeUpdate) {
	ps := s.pairScrapers[update.I]
	pair := ps.Pair()
	for _, trade := range update.NT {
		price, pok := strconv.ParseFloat(trade.Price, 64)
		volume, vok := strconv.ParseFloat(trade.Volume, 64)
		if pok == nil && vok == nil {
			t := &dia.Trade{
				Pair:           pair.ForeignName,
				Price:          price,
				Symbol:         pair.Symbol,
				Volume:         volume,
				Time:           time.Unix(trade.Timestamp, 0),
				ForeignTradeID: "",
				Source:         dia.CREX24Exchange,
			}
			s.chanTrades <- t
			log.Info("got trade: ", t)
		}
	}
}

func (s *CREX24Scraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	resp, err := http.Get("https://api.crex24.com/v2/public/instruments")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var parsedPairs []CREX24ApiInstrument
	err = json.NewDecoder(resp.Body).Decode(&parsedPairs)
	if err != nil {
		return nil, err
	}

	var results = make([]dia.Pair, len(parsedPairs))
	for i := 0; i < len(parsedPairs); i++ {
		results[i] = dia.Pair{
			Symbol:      parsedPairs[i].BaseCurrency,
			ForeignName: parsedPairs[i].Symbol,
			Ignore:      false,
			Exchange:    s.exchangeName,
		}
	}

	return results, nil
}

func (s *CREX24Scraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	if s.closed {
		return nil, errors.New("CREX24Scraper: Call ScrapePair on closed scraper")
	}

	s.connectLock.Lock()
	defer s.connectLock.Unlock()
	if !s.connected {
		// create the signalR connection lazily for the first pair scraper
		if err := s.connect(); err != nil {
			return nil, errors.New("CREX24Scraper: Could not establish signalR connection")
		}
	}

	msg := hubs.ClientMsg{
		H: "publicCryptoHub",
		M: "joinTradeHistory",
		A: []interface{}{pair.ForeignName},
		I: s.msgId,
	}

	err := s.client.Send(msg)
	s.msgId += 1

	if err != nil {
		return nil, errors.New("CREX24Scraper: could not send subscription signalR message")
	}

	ps := &CREX24PairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	return ps, nil
}

type CREX24PairScraper struct {
	parent *CREX24Scraper
	pair   dia.Pair
}

func (ps *CREX24PairScraper) Close() error {
	s := ps.parent
	if s.closed {
		return errors.New("CREX24Scraper: Scraper already closed")
	}
	msg := hubs.ClientMsg{
		H: "publicCryptoHub",
		M: "leaveTradeHistory",
		A: []interface{}{ps.pair.ForeignName},
		I: ps.parent.msgId,
	}
	err := ps.parent.client.Send(msg)
	ps.parent.msgId += 1
	if err != nil {
		return errors.New("CREX24Scraper: Could not send unsubscribe signalR message")
	}
	return nil
}

func (ps *CREX24PairScraper) Error() error {
	if ps.parent.closed {
		return errors.New("Scraper has been closed")
	} else {
		return nil
	}
}

func (ps *CREX24PairScraper) Pair() dia.Pair {
	return ps.pair
}
