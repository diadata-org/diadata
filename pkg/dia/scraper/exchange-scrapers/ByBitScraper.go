package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

var ByBitSocketURL string = utils.Getenv("BYBIT_WS_URL", "wss://stream.bybit.com/spot/public/v5")

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
	Data struct {
		TradeID   string `json:"v"`
		Timestamp int64  `json:"t"`
		Price     string `json:"p"`
		Size      string `json:"q"`
		Side      bool   `json:"m"`
		Type      string `json:"type"`
	} `json:"data"`
	Type      string `json:"type"`
	Topic     string `json:"topic"`
	Timestamp int64  `json:"ts"`
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

// NewByBitScraper get a scrapper for ByBit exchange
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
		log.Error("Connect to websocket server: ", err.Error())
	}

	s.wsClient = SwConn

	if scrape {
		go s.mainLoop()
	}

	return s
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

func (s *ByBitScraper) subscribe(foreignName string) {
	// Subscribing to the all markets at once.
	a := &ByBitSubscribe{
		OP:   "subscribe",
		Args: []string{"trade." + foreignName},
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

	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	for {
		message := &ByBitTradeResponse{}
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error("ws connection error: ", err.Error())
		}

		// the topic format is something like trade.BTCUSD
		topic := strings.Split(message.Topic, ".")

		if message.Data.Type == "" {
			continue
		}

		var tradeType int
		tradeType, err = strconv.Atoi(message.Data.Type)
		if err != nil {
			log.Error("parse type: ", err)
			continue
		}

		if len(topic) == 2 && topic[0] == "trade" && tradeType == 0 {
			ps, ok := s.pairScrapers[topic[1]]
			if ok {
				var (
					f64Price     float64
					f64Volume    float64
					exchangepair dia.ExchangePair
				)
				mdData := message.Data

				f64Price, err = strconv.ParseFloat(mdData.Price, 64)
				if err != nil {
					log.Error("parse price: ", err)
				}

				f64Volume, err = strconv.ParseFloat(mdData.Size, 64)
				if err != nil {
					log.Error("parse volume: ", err)
				}

				timeStamp := time.Unix(0, mdData.Timestamp*1e6)
				if mdData.TradeID != "" {
					if !mdData.Side {
						f64Volume = -f64Volume
					}

					exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, topic[1])
					if err != nil {
						log.Error(err)
					}
					t := &dia.Trade{
						Symbol:         ps.pair.Symbol,
						Pair:           topic[1],
						Price:          f64Price,
						Volume:         f64Volume,
						Time:           timeStamp,
						ForeignTradeID: mdData.TradeID,
						Source:         s.exchangeName,
						VerifiedPair:   exchangepair.Verified,
						BaseToken:      exchangepair.UnderlyingPair.BaseToken,
						QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
					}
					if exchangepair.Verified {
						log.Infoln("Got verified trade: ", t)
					}
					// Handle duplicate trades.
					discardTrade := t.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
					if !discardTrade {
						t.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
						ps.parent.chanTrades <- t
					}

				}

			} else {
				log.Error("Unknown Pair " + topic[1])
			}
		}
	}

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
	s.subscribe(pair.ForeignName)
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

// Channel returns the channel to get trades
func (s *ByBitScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

func (s *ByBitScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	// TO DO
	return dia.Asset{Symbol: symbol}, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
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
