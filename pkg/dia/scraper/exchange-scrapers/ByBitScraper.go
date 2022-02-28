package scrapers

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

var ByBitSocketURL string = "wss://stream.bybit.com/realtime"

type ByBitMarket struct {
	Name           string `json:"name"`
	Alias          string `json:"alias"`
	Status         string `json:"status"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	PriceScale     int    `json:"price_scale"`
	TakerFee       string `json:"taker_fee"`
	MakerFee       string `json:"maker_fee"`
	LeverageFilter struct {
		MinLeverage  int    `json:"min_leverage"`
		MaxLeverage  int    `json:"max_leverage"`
		LeverageStep string `json:"leverage_step"`
	} `json:"leverage_filter"`
	PriceFilter struct {
		MinPrice string `json:"min_price"`
		MaxPrice string `json:"max_price"`
		TickSize string `json:"tick_size"`
	} `json:"price_filter"`
	LotSizeFilter struct {
		MaxTradingQty float64 `json:"max_trading_qty"`
		MinTradingQty float64 `json:"min_trading_qty"`
		QtyStep       float64 `json:"qty_step"`
	} `json:"lot_size_filter"`
}

type ByBitMarketsResponse struct {
	RetCode int           `json:"ret_code"`
	RetMsg  string        `json:"ret_msg"`
	ExtCode string        `json:"ext_code"`
	ExtInfo string        `json:"ext_info"`
	Result  []ByBitMarket `json:"result"`
	TimeNow string        `json:"time_now"`
}

type ByBitTradeResponse struct {
	Topic string `json:"topic"`
	Data  []struct {
		Timestamp     string  `json:"timestamp"`
		TradeTimeMs   int64   `json:"trade_time_ms"`
		Symbol        string  `json:"symbol"`
		Side          string  `json:"side"`
		Size          float64 `json:"size"`
		Price         float64 `json:"price"`
		TickDirection string  `json:"tick_direction"`
		TradeID       string  `json:"trade_id"`
		CrossSeq      int     `json:"cross_seq"`
	} `json:"data"`
}

type ByBitSubscribe struct {
	OP   string   `json:"op"`
	Args []string `json:"args"`
}

// ByBitScraper provides  methods needed to get Trade information from ByBit
type ByBitScraper struct {
	// control flag for main loop
	run      bool
	wsClient *ws.Conn

	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*ByBitPairScraper
	// exchange name
	exchangeName string
	// channel to send trades
	chanTrades chan *dia.Trade
	db         *models.RelDB
}

//NewByBitScraper get a scrapper for ByBit exchange
func NewByBitScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *ByBitScraper {
	s := &ByBitScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*ByBitPairScraper),

		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		closed:       false,
		db:           relDB,
	}

	/*
	    In the case of needing access to private urls.
	    // Create HMAC instance from the secret key
	   	h := hmac.New(sha256.New, []byte(secret))

	   	// Write Data to it
	   	apiSecretBytes := []byte(secret)
	   	// Generate expires.
	   	expires := int((time.Now().UnixNano() + 1) * 1000)
	   	expiresBytes := []byte(fmt.Sprintf("GET/realtime%d", expires))
	   	data := append(apiSecretBytes, expiresBytes...)
	   	h.Write([]byte(data))

	   	// Get the signature
	   	signature := hex.EncodeToString(h.Sum(nil))

	   	// Generate the ws url.
	   	params := fmt.Sprintf("api_key=%s&expires=%d&signature=%s", secret, expires, signature)
	*/

	// Create the ws connection
	var wsDialer ws.Dialer

	SwConn, _, err := wsDialer.Dial(ByBitSocketURL, nil)
	if err != nil {
		println(err.Error())
	}

	s.wsClient = SwConn

	if scrape {
		go s.mainLoop()
	}

	return s
}

func (s *ByBitScraper) getMarkets() (markets []string) {
	var bbm ByBitMarketsResponse
	b, _, err := utils.GetRequest("https://api.bybit.com/v2/public/symbols")
	if err != nil {
		log.Errorln("Error Getting markets", err)
	}
	err = json.Unmarshal(b, &bbm)
	if err != nil {
		log.Error("getting markets: ", err)
	}

	for _, m := range bbm.Result {
		markets = append(markets, m.Name)
	}
	return
}

func (s *ByBitScraper) ping() {
	a := &ByBitSubscribe{
		OP: "ping",
	}
	log.Infoln("Ping: ", a.OP)
	if err := s.wsClient.WriteJSON(a); err != nil {
		log.Println(err.Error())
	}
}

func (s *ByBitScraper) subscribe() {
	// Subscribing to the all markets at once.
	a := &ByBitSubscribe{
		OP:   "subscribe",
		Args: []string{"trade.*"},
	}
	log.Println("subscribing", a)
	if err := s.wsClient.WriteJSON(a); err != nil {
		log.Println(err.Error())
	}
}

// runs in a goroutine until s is closed
func (s *ByBitScraper) mainLoop() {
	var err error

	pingTimer := time.NewTicker(10 * time.Second)
	go func() {
		for range pingTimer.C {
			go s.ping()
		}
	}()
	s.subscribe()
	for {
		message := &ByBitTradeResponse{}
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error("ws connection error: ", err.Error())
			s.subscribe()
		}
		// the topic format is something like trade.BTCUSD
		log.Info("got topic: ", message.Topic)
		topic := strings.Split(message.Topic, ".")

		if len(topic) == 2 && topic[0] == "trade" {
			ps, ok := s.pairScrapers[topic[1]]
			if ok {
				mdData := message.Data
				for _, v := range mdData {
					var f64Price float64
					var f64Volume float64
					var exchangepair dia.ExchangePair
					f64Price = v.Price
					f64Volume = v.Size

					timeStamp, _ := time.Parse(time.RFC3339, v.Timestamp)
					if v.TradeID != "" {
						if v.Side == "Sell" {
							f64Volume = -f64Volume
						}
						// Volume is given in USD on API.
						if f64Price != 0 {
							f64Volume = f64Volume / f64Price
						} else {
							continue
						}

						exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, v.Symbol)
						if err != nil {
							log.Error(err)
						}
						t := &dia.Trade{
							Symbol:         ps.pair.Symbol,
							Pair:           v.Symbol,
							Price:          f64Price,
							Volume:         f64Volume,
							Time:           timeStamp,
							ForeignTradeID: v.TradeID,
							Source:         s.exchangeName,
							VerifiedPair:   exchangepair.Verified,
							BaseToken:      exchangepair.UnderlyingPair.BaseToken,
							QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
						}
						if exchangepair.Verified {
							log.Infoln("Got verified trade: ", t)
						}
						ps.parent.chanTrades <- t
					}

				}
			} else {
				log.Error("Unknown Pair " + topic[1])
			}
		}
	}
	s.cleanup(err)
}

// Close channels for shutdown
func (s *ByBitScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.chanTrades)
	close(s.shutdownDone)
}

// Close any existing API connections, as well as channels, and terminates main loop
func (s *ByBitScraper) Close() error {
	if s.closed {
		return errors.New(s.exchangeName + "Scraper: Already closed")
	}
	s.run = false
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *ByBitScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	if s.closed {
		return nil, errors.New("s.exchangeName+Scraper: Call ScrapePair on closed scraper")
	}
	ps := &ByBitPairScraper{
		parent:      s,
		pair:        pair,
		apiEndPoint: pair.ForeignName,
		latestTrade: 0,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (s *ByBitScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	symbol := strings.ToUpper(pair.Symbol)
	pair.Symbol = symbol

	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			return pair, errors.New("Foreign name can not be normalized:" + pair.ForeignName + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return pair, errors.New("Symbol is black listed:" + symbol)
	}
	return pair, nil
}

//Channel returns the channel to get trades
func (s *ByBitScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *ByBitScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	// TO DO
	return dia.Asset{Symbol: symbol}, nil
}

//FetchAvailablePairs returns a list with all available trade pairs
func (s *ByBitScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	data, _, err := utils.GetRequest("https://api.bybit.com/v2/public/symbols")
	if err != nil {
		return
	}
	var ar ByBitMarketsResponse
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar.Result {
			if p.Status != "Trading" {
				continue
			}
			pairToNormalize := dia.ExchangePair{
				Symbol:      p.BaseCurrency,
				ForeignName: p.Name,
				Exchange:    s.exchangeName,
			}
			pair, serr := s.NormalizePair(pairToNormalize)
			if serr == nil {
				pairs = append(pairs, pair)
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (s *ByBitScraper) Error() error {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ByBitPairScraper implements PairScraper for ByBit
type ByBitPairScraper struct {
	apiEndPoint string
	parent      *ByBitScraper
	pair        dia.ExchangePair
	closed      bool
	latestTrade int
}

// Close stops listening for trades of the pair associated
func (ps *ByBitPairScraper) Close() error {
	ps.closed = true
	return ps.Error()
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *ByBitPairScraper) Error() error {
	ps.parent.errorLock.RLock()
	defer ps.parent.errorLock.RUnlock()
	return ps.parent.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *ByBitPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
