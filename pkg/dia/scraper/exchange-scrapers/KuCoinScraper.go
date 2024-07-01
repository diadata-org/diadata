package scrapers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"
)

var (
	kucoinWSBaseString = "wss://ws-api-spot.kucoin.com/"
	kucoinTokenURL     = "https://api.kucoin.com/api/v1/bullet-public"
)

// A WebSocketSubscribeMessage represents a message to subscribe the public/private channel.
type kuCoinWSSubscribeMessage struct {
	Id             string `json:"id"`
	Type           string `json:"type"`
	Topic          string `json:"topic"`
	PrivateChannel bool   `json:"privateChannel"`
	Response       bool   `json:"response"`
}

type kuCoinWSResponse struct {
	Type    string       `json:"type"`
	Topic   string       `json:"topic"`
	Subject string       `json:"subject"`
	Data    kuCoinWSData `json:"data"`
}

type kuCoinWSData struct {
	Sequence string `json:"sequence"`
	Type     string `json:"type"`
	Symbol   string `json:"symbol"`
	Side     string `json:"side"`
	Price    string `json:"price"`
	Size     string `json:"size"`
	TradeID  string `json:"tradeId"`
	Time     string `json:"time"`
}

type KuCoinScraper struct {
	// signaling channels for session initialization and finishing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	// use sync.Maps to concurrently handle multiple pairs
	pairScrapers map[string]*KuCoinPairScraper // dia.ExchangePair -> KuCoinPairScraper
	// pairSubscriptions sync.Map                      // dia.ExchangePair -> string (subscription ID)
	// pairLocks         sync.Map                      // dia.ExchangePair -> sync.Mutex
	exchangeName string
	chanTrades   chan *dia.Trade
	publicToken  string
	pingInterval int64
	db           *models.RelDB
}

func NewKuCoinScraper(apiKey string, secretKey string, exchange dia.Exchange, scrape bool, relDB *models.RelDB) *KuCoinScraper {
	token, pingInterval, err := getPublicKuCoinToken(kucoinTokenURL)
	if err != nil {
		log.Fatal("getPublicKuCoinToken: ", err)
	}

	s := &KuCoinScraper{
		exchangeName: exchange.Name,
		pairScrapers: make(map[string]*KuCoinPairScraper),
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		publicToken:  token,
		pingInterval: pingInterval,
		db:           relDB,
	}

	// establish connection in the background
	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *KuCoinScraper) mainLoop() {

	time.Sleep(2 * time.Second)
	var pairTickers []string
	for p := range s.pairScrapers {
		pairTickers = append(pairTickers, p)
	}
	lotSize := 100

	var wsDialer ws.Dialer
	subscriptions := kuCoinWSSubscribeMessage{
		Type:  "subscribe",
		Topic: "/market/match:",
	}

	for i, pairTicker := range pairTickers {

		if (i%lotSize == 0 && i > 0) || i == len(pairTickers)-1 {
			// Subscribe for lots of 100 pairs.
			client, _, err := wsDialer.Dial(kucoinWSBaseString+"?token="+s.publicToken, nil)
			if err != nil {
				log.Fatal("Dial KuCoin ws base string: ", err)
			}
			if err := client.WriteJSON(subscriptions); err != nil {
				log.Fatal(err.Error())
			}
			go s.listenToTrades(client)

			// Initialize client and subscriptions for the next log.
			subscriptions = kuCoinWSSubscribeMessage{
				Type:  "subscribe",
				Topic: "/market/match:",
			}
		}

		if i%lotSize == 0 {
			subscriptions.Topic += pairTicker
		} else {
			subscriptions.Topic += "," + pairTicker
		}

	}

}

func (s *KuCoinScraper) listenToTrades(client *ws.Conn) {
	lastTradeMap := make(map[dia.Pair]time.Time)
	countMap := make(map[dia.Pair]int)
	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	go ping(client, 20)
	// go ping(client, s.pingInterval)

	for {
		var message kuCoinWSResponse

		_, data, err := client.ReadMessage()
		if err != nil {
			log.Errorf("ReadMessage: %v", err)
			continue
		}

		err = json.Unmarshal(data, &message)
		if err != nil {
			log.Errorf("ReadMessage: %v", err)
			continue
		}
		// err := client.ReadJSON(&message)
		// if err != nil {
		// 	log.Errorf("ReadMessage: %v", err)
		// 	continue
		// }

		if message.Type == "pong" {
			log.Info("Successful ping: received pong.")
		} else if message.Type == "message" {

			// Parse trade quantities.
			price, volume, timestamp, foreignTradeID, pairTicker, err := parseKuCoinTradeMessage(message)
			if err != nil {
				log.Error("parseTradeMessage: ", err)
			}

			exchangepair, err := s.db.GetExchangePairCache(s.exchangeName, pairTicker)
			if err != nil {
				log.Error(err)
			}

			// Make trade times unique
			pair := dia.Pair{QuoteToken: exchangepair.UnderlyingPair.QuoteToken, BaseToken: exchangepair.UnderlyingPair.BaseToken}
			if _, ok := lastTradeMap[pair]; ok {
				if lastTradeMap[pair] != timestamp {
					lastTradeMap[pair] = timestamp
					countMap[pair] = 0
				} else {
					//nolint
					timestamp.Add(time.Duration(countMap[pair]+1) * time.Nanosecond)
					countMap[pair] += 1
				}
			} else {
				lastTradeMap[pair] = timestamp
			}

			trade := &dia.Trade{
				QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
				BaseToken:      exchangepair.UnderlyingPair.BaseToken,
				Symbol:         exchangepair.UnderlyingPair.QuoteToken.Symbol,
				Pair:           pairTicker,
				Price:          price,
				Volume:         volume,
				Time:           timestamp,
				Source:         s.exchangeName,
				ForeignTradeID: foreignTradeID,
				VerifiedPair:   exchangepair.Verified,
			}

			discardTrade := trade.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
			if !discardTrade {
				trade.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
				s.chanTrades <- trade
			}
			log.Info("Got trade: ", trade)
			s.chanTrades <- trade
		}
	}
}

func parseKuCoinTradeMessage(message kuCoinWSResponse) (price float64, volume float64, timestamp time.Time, foreignTradeID string, pairTicker string, err error) {
	price, err = strconv.ParseFloat(message.Data.Price, 64)
	if err != nil {
		return
	}
	volume, err = strconv.ParseFloat(message.Data.Size, 64)
	if err != nil {
		return
	}
	if message.Data.Side == "sell" {
		volume -= 1
	}
	timeMilliseconds, err := strconv.Atoi(message.Data.Time)
	if err != nil {
		return
	}
	pairTicker = message.Data.Symbol
	timestamp = time.Unix(0, int64(timeMilliseconds))
	foreignTradeID = message.Data.TradeID
	return
}

func (s *KuCoinScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *KuCoinScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("KucoinScraper: Call ScrapePair on closed scraper")
	}

	ps := &KuCoinPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// FetchAvailablePairs returns all traded pairs on kucoin.
func (s *KuCoinScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

// FillSymbolData adds the name to the asset underlying @symbol on kucoin.
func (s *KuCoinScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	asset.Symbol = symbol
	return
}

// KuCoinPairScraper implements PairScraper for kuCoin
type KuCoinPairScraper struct {
	parent *KuCoinScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *KuCoinPairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("KuCoinPairScraper: Already closed")
	}

	// TODO stop collection for the pair

	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *KuCoinScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *KuCoinPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *KuCoinPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

// ------------------------------------------------------------
// Connection utilities
// ------------------------------------------------------------

// A WebSocketMessage represents a message between the WebSocket client and server.
type kuCoinWSMessage struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type kuCoinPostResponse struct {
	Code string `json:"code"`
	Data struct {
		Token           string            `json:"token"`
		InstanceServers []instanceServers `json:"instanceServers"`
	} `json:"data"`
}

type instanceServers struct {
	PingInterval int64 `json:"pingInterval"`
}

// Send ping to server.
func ping(wsClient *ws.Conn, pingInterval int64) {
	var ping kuCoinWSMessage
	ping.Type = "ping"
	tick := time.NewTicker(time.Duration(pingInterval/2) * time.Second)

	for range tick.C {
		log.Infof("send ping.")
		if err := wsClient.WriteJSON(ping); err != nil {
			log.Error(err.Error())
		}
	}
}

// getPublicKuCoinToken returns a token for public market data along with the pingInterval in seconds.
func getPublicKuCoinToken(url string) (token string, pingInterval int64, err error) {
	postBody, _ := json.Marshal(map[string]string{})
	responseBody := bytes.NewBuffer(postBody)
	data, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return
	}
	defer data.Body.Close()
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return
	}

	var postResp kuCoinPostResponse
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return
	}
	if len(postResp.Data.InstanceServers) > 0 {
		pingInterval = postResp.Data.InstanceServers[0].PingInterval
	}
	token = postResp.Data.Token
	return
}
