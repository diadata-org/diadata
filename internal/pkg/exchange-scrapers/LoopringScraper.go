package scrapers

import (
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

var _LoopringSocketurl string = "wss://ws.api3.loopring.io/v3/ws"

type WebSocketRequest struct {
	Op       string          `json:"op"`
	Sequence int             `json:"sequence"`
	Topics   []LoopringTopic `json:"topics"`
	Result   struct {
		Status string `json:"status"`
	} `json:"result"`
}

type WebSocketResponse struct {
	Topic struct {
		Topic  string `json:"topic"`
		Market string `json:"market"`
	} `json:"topic"`
	Ts   int64      `json:"ts"`
	Data [][]string `json:"data"`
}

type LoopringTopic struct {
	Topic    string `json:"topic"`
	Market   string `json:"market"`
	Count    int64  `json:"count"`
	Snapshot bool   `json:"snapshot"`
}

type Topic struct {
	Topic    string `json:"topic"`
	Market   string `json:"market"`
	Count    int64  `json:"count"`
	Snapshot bool   `json:"snapshot"`
}

type LoopringMarket struct {
	Data []struct {
		Market             string `json:"market"`
		BaseTokenID        int    `json:"baseTokenId"`
		QuoteTokenID       int    `json:"quoteTokenId"`
		PrecisionForPrice  int    `json:"precisionForPrice"`
		OrderbookAggLevels int    `json:"orderbookAggLevels"`
		Enabled            bool   `json:"enabled"`
	} `json:"markets"`
}

type LoopringScraper struct {
	wsClient      *ws.Conn
	decimalsAsset map[string]float64
	// signaling channels for session initialization and finishing
	//TODO: Channel not used. Consider removing or refactoring
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*LoopringPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	wsURL        string
}

type LoopringKey struct {
	Key string `json:"key"`
}

// NewLoopringScraper returns a new LoopringScraper for the given pair
func NewLoopringScraper(exchange dia.Exchange) *LoopringScraper {

	decimalAsset := make(map[string]float64)
	decimalAsset["ETH"] = 18
	decimalAsset["WETH"] = 18
	decimalAsset["LRC"] = 18
	decimalAsset["USDT"] = 6
	decimalAsset["DAI"] = 18
	decimalAsset["LINK"] = 18
	decimalAsset["KEEP"] = 18
	decimalAsset["USDC"] = 6
	decimalAsset["DXD"] = 18
	decimalAsset["TRB"] = 18
	decimalAsset["AUC"] = 18
	decimalAsset["RPL"] = 18
	decimalAsset["WBTC"] = 8
	decimalAsset["RENBTC"] = 8
	decimalAsset["PAX"] = 18
	decimalAsset["MKR"] = 18
	decimalAsset["BUSD"] = 18
	decimalAsset["SNX"] = 18
	decimalAsset["GNO"] = 18
	decimalAsset["LEND"] = 18
	decimalAsset["REN"] = 18
	decimalAsset["REP"] = 18
	decimalAsset["BNT"] = 18
	decimalAsset["PBTC"] = 18
	decimalAsset["COMP"] = 18
	decimalAsset["PNT"] = 18
	decimalAsset["PNK"] = 18
	decimalAsset["NEST"] = 18
	decimalAsset["BTU"] = 18
	decimalAsset["BZRX"] = 18
	decimalAsset["VBZRX"] = 18
	decimalAsset["GRID"] = 12

	s := &LoopringScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		pairScrapers:  make(map[string]*LoopringPairScraper),
		exchangeName:  exchange.Name,
		error:         nil,
		chanTrades:    make(chan *dia.Trade),
		decimalsAsset: decimalAsset,
	}

	key, err := getAPIKey()
	if err != nil {
		log.Fatal("get api key: ", err)
	}
	s.wsURL = _LoopringSocketurl + "?wsApiKey=" + key
	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(s.wsURL, nil)
	if err != nil {
		log.Error("Error connecting to ws: ", err.Error())
	}

	s.wsClient = SwConn

	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *LoopringScraper) mainLoop() {

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)
	s.subscribeToALL()

	for {

		var makemap WebSocketResponse
		messageType, message, err := s.wsClient.ReadMessage()
		if err != nil {
			log.Error("reading websocket message: ", err)
			s.reconnectToWS()
			s.subscribeToALL()
		}

		err = json.Unmarshal(message, &makemap)

		if err != nil {
			message := string(message)
			if message == "ping" {
				e := s.Pong(messageType)
				if e != nil {
					log.Error("send pong: ", err)
				} else {
					log.Info("sent pong")
				}
			}
		} else {
			if makemap.Topic.Topic == "trade" {
				asset := strings.Split(makemap.Topic.Market, "-")
				f64Price, _ := strconv.ParseFloat(makemap.Data[0][4], 64)
				timestamp, err := strconv.ParseInt(makemap.Data[0][0], 10, 64)
				if err != nil {
					log.Error("Error Parsing time", err)
				}
				volume, err := strconv.ParseFloat(makemap.Data[0][3], 64)
				if err != nil {
					log.Error("Error Parsing time", err)
				}
				volume = volume / math.Pow(10, s.decimalsAsset[asset[0]])
				if makemap.Data[0][2] == "SELL" {
					volume = -volume
				}
				t := &dia.Trade{
					Symbol: asset[0],
					Pair:   makemap.Topic.Market,
					Price:  f64Price,
					Time:   time.Unix(timestamp/1000, 0),
					Volume: volume,
					Source: s.exchangeName,
				}
				s.chanTrades <- t
				log.Info("Got trade: ", t)
			}
		}
	}
}

func (s *LoopringScraper) subscribeToALL() {
	log.Info("Subscribing To all pairs")

	count := 0
	var topics []LoopringTopic
	for key := range s.pairScrapers {
		lptopic := LoopringTopic{Market: key, Topic: "trade", Count: 20, Snapshot: true}
		topics = append(topics, lptopic)
		count++
		if count > 19 {
			break
		}
	}
	log.Info("topics for sub: ", topics)
	wr := &WebSocketRequest{
		Op:       "sub",
		Sequence: 1000,
		Topics:   topics,
	}
	if err := s.wsClient.WriteJSON(wr); err != nil {
		log.Error(err)
	}

}

// Pong sends the string "pong" to the server.
func (s *LoopringScraper) Pong(messageType int) error {
	err := s.wsClient.WriteMessage(messageType, []byte("pong"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoopringScraper) reconnectToWS() {

	log.Info("Reconnecting ws")
	key, err := getAPIKey()
	if err != nil {
		log.Fatal("fetching api key: ", err)
	}

	s.wsURL = _LoopringSocketurl + "?wsApiKey=" + key

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(s.wsURL, nil)

	if err != nil {
		println(err.Error())
	}
	s.wsClient = SwConn
}

func getAPIKey() (string, error) {
	resp, err := utils.GetRequest("https://api3.loopring.io/v3/ws/key")
	if err != nil {
		return "", err
	}
	var lkResponse LoopringKey
	err = json.Unmarshal(resp, &lkResponse)
	if err != nil {
		return "", err
	}
	return lkResponse.Key, nil
}

func (s *LoopringScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	return dia.Pair{}, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *LoopringScraper) Close() error {

	if s.closed {
		return errors.New("LoopringScraper: Already closed")
	}
	s.wsClient.Close()
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *LoopringScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("LoopringScraper: Call ScrapePair on closed scraper")
	}

	ps := &LoopringPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *LoopringScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	data, err := utils.GetRequest("https://api3.loopring.io/api/v3/exchange/markets")

	if err != nil {
		return
	}

	var ar LoopringMarket
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar.Data {
			symbols := strings.Split(p.Market, "-")
			pairs = append(pairs, dia.Pair{
				Symbol:      symbols[0],
				ForeignName: p.Market,
				Exchange:    s.exchangeName,
			})
		}
	}
	return
}

// LoopringPairScraper implements PairScraper for Loopring exchange
type LoopringPairScraper struct {
	parent *LoopringScraper
	pair   dia.Pair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *LoopringPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *LoopringScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *LoopringPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *LoopringPairScraper) Pair() dia.Pair {
	return ps.pair
}
