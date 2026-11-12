package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"
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
	Event   string      `json:"event"`
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
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

func extractUrlSymbolFromChannel(channel string) (after string) {
	// Channel is live_trades_{pair}
	// Remove the prefix live_trades_
	if strings.HasPrefix(channel, "live_trades_") {
		after = strings.Split(channel, "live_trades_")[1]
	}
	return after
}

func (s *BitstampScraper) foreignNameToUrlSymbol(foreignName string) string {
	return strings.ToLower(strings.ReplaceAll(foreignName, "-", ""))
}

func (s *BitstampScraper) receive() {
	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	var resp BitstampWsResponse
	if err := s.wsClient.ReadJSON(&resp); err != nil {
		log.Error("Receive message error:", err)
		return
	}
	foreignName := extractUrlSymbolFromChannel(resp.Channel)

	switch event := resp.Event; event {
	case "bts:subscription_succeeded":
		log.Info("Subsription succeeded:", foreignName)
	// TODO: add subscription failed...
	// Probably "bts:subscription_failed"
	// Just a guess
	case "bts:request_reconnect":
		// TODO: handle reconnect
		log.Warn("Server request for a reconnect")
	case "bts:heartbeat":
		var data BitstampPingData
		if err := json.Unmarshal([]byte(resp.Data.(string)), &data); err != nil {
			log.Warn("Unmarshal ping error:", err, resp.Data)
		} else if data.Status == "success" {
			log.Info("Heart is beating")
		} else {
			log.Warning("Check Heart", data)
		}
	default:

		if strings.HasPrefix(event, "trade") {

			data := resp.Data.(map[string]interface{})
			ps, ok := s.pairScrapers[foreignName]

			if ok {
				timestamp, _ := strconv.ParseInt(data["microtimestamp"].(string), 10, 64)
				volume := data["amount"].(float64)
				side := data["type"].(float64)
				if side == 1 {
					volume *= -1
				}

				pair, err := s.db.GetExchangePairCache(s.exchangeName, foreignName)
				if err != nil {
					log.Error("get exchange pair from cache: ", err)
				}

				t := &dia.Trade{
					Symbol:         ps.pair.Symbol,
					Pair:           foreignName,
					Price:          data["price"].(float64),
					Volume:         volume,
					Time:           time.Unix(0, 1000*timestamp),
					Source:         s.exchangeName,
					ForeignTradeID: fmt.Sprintf("%d", int(data["id"].(float64))),
					VerifiedPair:   pair.Verified,
					QuoteToken:     pair.UnderlyingPair.QuoteToken,
					BaseToken:      pair.UnderlyingPair.BaseToken,
				}

				// Handle duplicate trades.
				discardTrade := t.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
				if !discardTrade {
					t.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
					s.chanTrades <- t
				}
				log.Info("Found trade:", t)
			}
		} else {
			log.Warnf("Unidentified response event %s: -- %v", event, resp)
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
		s.receive()
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

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Warn("Error closing file: ", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Read pair body:", err)
	}

	err = json.Unmarshal(body, &bitstampPairsInfo)
	if err != nil {
		log.Error("Unmarshal pairs:", err)
	}

	for _, p := range bitstampPairsInfo {
		pairToNormalized := dia.ExchangePair{
			Symbol:      strings.Split(p.Name, "/")[0],
			ForeignName: p.UrlSymbol,
			Exchange:    s.exchangeName,
			UnderlyingPair: dia.Pair{
				QuoteToken: dia.Asset{
					Symbol: strings.Split(p.Name, "/")[0],
				},
				BaseToken: dia.Asset{
					Symbol: strings.Split(p.Name, "/")[1],
				},
			},
		}
		pairs = append(pairs, pairToNormalized)
	}
	return
}

func (s *BitstampScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (s *BitstampScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

func (s *BitstampScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

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

func (ps *BitstampPairScraper) Close() error {
	ps.closed = true
	return nil
}

func (ps *BitstampPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (ps *BitstampPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
