package scrapers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	ws "github.com/gorilla/websocket"
)

const ()

type BitstampScraper struct {
	wsClient     *ws.Conn
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers map[string]*BitstampPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB

	urlSymbols map[string]string
}

type BitstampPairScraper struct {
	parent *BitstampScraper
	pair   dia.ExchangePair
	closed bool
}

type BitstampPairsInfo []struct {
	Name                        string `json:"name"`
	UrlSymbol                   string `json:"url_symbol"`
	BaseDecimal                 uint8  `json:"base_decimal"`
	CounterDecimals             uint8  `json:"counter_decimals"`
	InstantOrderCounterDecimals uint8  `json:"instant_order_counter_decimals"`
	MinimumOrder                string `json:"minimum_order"`
	Trading                     string `json:"trading"`
	InstantAndMarketOrders      string `json:"instant_and_market_orders"`
	Description                 string `json:"description"`
}

type BitstampWsResponse struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	Data    []byte `json:"data"`
}

type BitstampPingData struct {
	Status string `json:"status"`
}

type BitstampTradeData struct {
	Id             string  `json:"id"`
	Amount         float64 `json:"amount"`
	AmountStr      string  `json:"amount_str"`
	Price          float64 `json:"price"`
	PriceStr       string  `json:"price_str"`
	Type           uint8   `json:"type"`
	Timestamp      string  `json:"timestamp"`
	Microtimestamp string  `json:"microtimestamp"`
	BuyOrderId     uint64  `json:"buy_order_id"`
	SellOrderId    uint64  `json:"sell_order_id"`
}

func NewBitstampScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BitstampScraper {
	s := &BitstampScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		error:        nil,
		pairScrapers: make(map[string]*BitstampPairScraper),
		exchangeName: exchange.Name,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
		urlSymbols:   make(map[string]string),
	}
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws.bitstamp.net", nil)
	if err != nil {
		log.Error("Websocket connect error", err)
	}
	s.wsClient = wsConn
	if scrape {
		go s.mainLoop()
	}
	return s
}

func extractUrlSymbolFromChannel(channel string) string {
	// Channel is live_trades_{pair}
	// Remove the prefix live_trades_
	after, _ := strings.CutPrefix(channel, "live_trades_")
	return after
}

func (s *BitstampScraper) urlSymbolToForeignName(urlSymbol string) string {
	return s.urlSymbols[urlSymbol]
}

func (s *BitstampScraper) foreignNameToUrlSymbol(foreignName string) string {
	return strings.ToLower(strings.ReplaceAll(foreignName, "-", ""))
}

func (s *BitstampScraper) receive() {
	var resp BitstampWsResponse
	if err := s.wsClient.ReadJSON(&resp); err != nil {
		log.Error("Receive message error:", err)
		return
	}
	urlSymbol := extractUrlSymbolFromChannel(resp.Channel)
	foreignName := s.urlSymbolToForeignName(urlSymbol)
	switch event := resp.Event; event {
	case "bts:subscription_succeeded":
		log.Info("Subsription succeeded:", urlSymbol)
	// TODO: add subscription failed...
	// Probably "bts:subscription_failed"
	// Just a guess
	case "bts:request_reconnect":
		// TODO: handle reconnect
		log.Warn("Server request for a reconnect")
	case "bts:heartbeat":
		var data BitstampPingData
		if err := json.Unmarshal(resp.Data, &data); err != nil {
			log.Warn("Unmarshal ping error:", err, resp.Data)
		} else if data.Status == "success" {
			log.Info("Heart is beating")
		} else {
			log.Warning("Check Heart", data)
		}
	default:
		if strings.HasPrefix(event, "live_trades_") {
			var data BitstampTradeData
			if err := json.Unmarshal(resp.Data, &data); err != nil {
				log.Warn("Unmarshal live trades error:", err, resp.Data)
				return
			}
			ps, ok := s.pairScrapers[foreignName]
			print(ps, ok)
			if ok {
				timestamp, _ := strconv.ParseInt(data.Timestamp, 10, 64)
				// TODO: add nanosecond to time.Unix
				// microtimestamp, _ := strconv.ParseInt(data.Microtimestamp, 10, 64)
				t := &dia.Trade{
					Symbol:     ps.pair.Symbol,
					Pair:       foreignName,
					QuoteToken: ps.pair.UnderlyingPair.QuoteToken,
					BaseToken:  ps.pair.UnderlyingPair.BaseToken,
					Price:      data.Price,
					Volume:     data.Amount,
					Time:       time.Unix(timestamp, 0),
					// ForeignTradeID: ,
					// EstimatedUSDPrice: ,
					Source:       s.exchangeName,
					VerifiedPair: ps.pair.Verified,
				}
				log.Info("Found trade:", t)
			}
		} else {
			log.Warn("Unidentified response event", event, resp)
		}
	}
}

func (s *BitstampScraper) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Warn("Shutting down BitstampScraper")
			s.cleanup(nil)
			return
		default:
		}
		go s.receive()
	}
}

func (s *BitstampScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("BitstampScraper: Call ScrapePair on closed scraper")
	}
	ps := &BitstampPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps

	// Subscribe to pair
	urlSymbol := s.foreignNameToUrlSymbol(pair.ForeignName)
	message := map[string]interface{}{
		"event": "bts:subscribe",
		"data": map[string]interface{}{
			"channel": "live_trades_" + urlSymbol,
		},
	}
	err := s.wsClient.WriteJSON(message)
	if err != nil {
		log.Error("Error sending subscription for", ps.pair.ForeignName, err)
	}
	log.Info("Sent subscription for:", ps.pair.ForeignName)

	return ps, nil
}

func (s *BitstampScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	var bitstampPairsInfo BitstampPairsInfo
	resp, err := http.Get("https://www.bitstamp.net/api/v2/trading-pairs-info/")
	if err != nil {
		log.Error("Get Pairs:", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Read pair body:", err)
	}

	err = json.Unmarshal(body, &bitstampPairsInfo)
	if err != nil {
		log.Error("Unmarshal pairs:", err)
	}

	for _, p := range bitstampPairsInfo {
		foreignName := strings.ReplaceAll(p.Name, "/", "-")
		s.urlSymbols[p.UrlSymbol] = foreignName
		pairToNormalized := dia.ExchangePair{
			Symbol:      strings.Split(p.Name, "/")[0],
			ForeignName: foreignName,
			Exchange:    s.exchangeName,
			UnderlyingPair: dia.Pair{
				QuoteToken: dia.Asset{
					Symbol: strings.Split(p.Name, "/")[0],
					Name:   strings.Split(p.Description, " / ")[0],
				},
				BaseToken: dia.Asset{
					Symbol: strings.Split(p.Name, "/")[1],
					Name:   strings.Split(p.Description, " / ")[1],
				},
			},
		}
		pairs = append(pairs, pairToNormalized)
	}
	return
}

// TODO: don't why it's necessary
func (s *BitstampScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

// TODO: don't why it's necessary
func (s *BitstampScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// TODO: don't why it's necessary
func (s *BitstampScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (s *BitstampScraper) Close() error {
	if s.closed {
		return errors.New("BitstampScraper: Already closed")
	}
	if err := s.wsClient.Close(); err != nil {
		log.Error("Error closing Bitstamp.wsClient", err)
	}
	close(s.shutdown)
	<-s.shutdownDone
	defer s.errorLock.RUnlock()
	return s.error
}

// TODO: don't why it's necessary
func (ps *BitstampPairScraper) Close() error {
	// TODO: probably have to unsubscribe in ws
	ps.closed = true
	return nil
}

// TODO: don't why it's necessary
func (ps *BitstampPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// TODO: don't why it's necessary
func (ps *BitstampPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
