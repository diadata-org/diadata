package scrapers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

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

type CREX4Asset []struct {
	Symbol                   string      `json:"symbol"`
	Name                     string      `json:"name"`
	IsFiat                   bool        `json:"isFiat"`
	DepositsAllowed          bool        `json:"depositsAllowed"`
	DepositConfirmationCount int         `json:"depositConfirmationCount"`
	MinDeposit               float64     `json:"minDeposit"`
	WithdrawalsAllowed       bool        `json:"withdrawalsAllowed"`
	WithdrawalPrecision      int         `json:"withdrawalPrecision"`
	MinWithdrawal            float64     `json:"minWithdrawal"`
	MaxWithdrawal            float64     `json:"maxWithdrawal"`
	IsDelisted               bool        `json:"isDelisted"`
	Transports               interface{} `json:"transports"`
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
	db           *models.RelDB
}

func NewCREX24Scraper(exchange dia.Exchange, relDB *models.RelDB) *CREX24Scraper {
	s := &CREX24Scraper{
		pairScrapers: make(map[string]*CREX24PairScraper),
		exchangeName: exchange.Name,
		chanTrades:   make(chan *dia.Trade),
		connected:    false,
		closed:       false,
		msgId:        1,
		db:           relDB,
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

func (s *CREX24Scraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
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
				err := json.NewDecoder(strings.NewReader(payload)).Decode(&parsedUpdate)
				if err != nil {
					log.Error(err)
				}
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
		if trade.Side == "s" {
			volume *= -1
		}
		if pok == nil && vok == nil {
			exchangepair, err := s.db.GetExchangePairCache(s.exchangeName, pair.ForeignName)
			if err != nil {
				log.Error(err)
			}
			t := &dia.Trade{
				Pair:           pair.ForeignName,
				Price:          price,
				Symbol:         pair.Symbol,
				Volume:         volume,
				Time:           time.Unix(trade.Timestamp, 0),
				ForeignTradeID: "",
				Source:         dia.CREX24Exchange,
				VerifiedPair:   exchangepair.Verified,
				BaseToken:      exchangepair.UnderlyingPair.BaseToken,
				QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
			}
			if exchangepair.Verified {
				log.Infoln("Got verified trade", t)
			}
			s.chanTrades <- t
		}
	}
}

func (s *CREX24Scraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	resp, err := http.Get("https://api.crex24.com/v2/public/instruments")
	if err != nil {
		return nil, err
	}
	defer utils.CloseHTTPResp(resp)

	var parsedPairs []CREX24ApiInstrument
	err = json.NewDecoder(resp.Body).Decode(&parsedPairs)
	if err != nil {
		return nil, err
	}

	var results = make([]dia.ExchangePair, len(parsedPairs))
	for i := 0; i < len(parsedPairs); i++ {
		results[i] = dia.ExchangePair{
			Symbol:      parsedPairs[i].BaseCurrency,
			ForeignName: parsedPairs[i].Symbol,
			Exchange:    s.exchangeName,
		}
	}

	return results, nil
}

// FillSymbolData collects all available information on an asset traded on CREX24
func (s *CREX24Scraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	var response CREX4Asset
	data, err := utils.GetRequest("https://api.crex24.com/v2/public/currencies?filter=" + symbol)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return
	}
	asset.Symbol = response[0].Symbol
	asset.Name = response[0].Name
	return asset, nil
}

func (s *CREX24Scraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
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
	pair   dia.ExchangePair
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
		return errors.New("scraper has been closed")
	} else {
		return nil
	}
}

func (ps *CREX24PairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
