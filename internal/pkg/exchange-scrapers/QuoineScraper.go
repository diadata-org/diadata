package scrapers

import (
	"encoding/json"
	"errors"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
)

var pingPeriod = 60*time.Second*2 - 1

var LiquidSocketURL string = "wss://tap.liquid.com/app/LiquidTapClient"
var LiquidSocketRestURL string = "http://api.liquid.com"

type LiquidSubscribe struct {
	Event string        `json:"event"`
	Data  LiquidChannel `json:"data"`
}

type LiquidChannel struct {
	Channel string `json:"channel"`
}

type LiquidProducts []LiquidProduct

type LiquidProduct struct {
	ID                      string      `json:"id"`
	ProductType             string      `json:"product_type"`
	Code                    string      `json:"code"`
	Name                    interface{} `json:"name"`
	MarketAsk               float64     `json:"market_ask"`
	MarketBid               float64     `json:"market_bid"`
	Indicator               interface{} `json:"indicator"`
	Currency                string      `json:"currency"`
	CurrencyPairCode        string      `json:"currency_pair_code"`
	Symbol                  interface{} `json:"symbol"`
	BtcMinimumWithdraw      interface{} `json:"btc_minimum_withdraw"`
	FiatMinimumWithdraw     interface{} `json:"fiat_minimum_withdraw"`
	PusherChannel           string      `json:"pusher_channel"`
	TakerFee                string      `json:"taker_fee"`
	MakerFee                string      `json:"maker_fee"`
	LowMarketBid            string      `json:"low_market_bid"`
	HighMarketAsk           string      `json:"high_market_ask"`
	Volume24H               string      `json:"volume_24h"`
	LastPrice24H            string      `json:"last_price_24h"`
	LastTradedPrice         string      `json:"last_traded_price"`
	LastTradedQuantity      string      `json:"last_traded_quantity"`
	AveragePrice            string      `json:"average_price"`
	QuotedCurrency          string      `json:"quoted_currency"`
	BaseCurrency            string      `json:"base_currency"`
	TickSize                string      `json:"tick_size"`
	Disabled                bool        `json:"disabled"`
	MarginEnabled           bool        `json:"margin_enabled"`
	CfdEnabled              bool        `json:"cfd_enabled"`
	PerpetualEnabled        bool        `json:"perpetual_enabled"`
	LastEventTimestamp      string      `json:"last_event_timestamp"`
	Timestamp               string      `json:"timestamp"`
	MultiplierUp            string      `json:"multiplier_up"`
	MultiplierDown          string      `json:"multiplier_down"`
	AverageTimeInterval     int         `json:"average_time_interval"`
	ProgressiveTierEligible bool        `json:"progressive_tier_eligible"`
}

const (
	API_DELAY       = 1*time.Second + 500*time.Millisecond
	EXECUTION_LIMIT = 200
)

type QuoineScraper struct {
	wsClient *ws.Conn

	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*QuoinePairScraper
	productPairIds map[string]string

	chanTrades chan *dia.Trade
}

func NewQuoineScraper(exchange dia.Exchange) *QuoineScraper {
	var err error

	scraper := &QuoineScraper{
		exchangeName:   exchange.Name,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]string),
		pairScrapers:   make(map[string]*QuoinePairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	err = scraper.readProductIds()

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(LiquidSocketURL, nil)

	if err != nil {
		println(err.Error())
	}

	scraper.wsClient = SwConn
	if err != nil {
		log.Error("Couldn't obtain Quoine product ids:", err)
	}
	go scraper.sendPing()

	go scraper.mainLoop()

	return scraper
}

func (scraper *QuoineScraper) sendPing() {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			ls := &LiquidSubscribe{
				Event: "pusher:ping",
			}
			scraper.wsClient.WriteJSON(ls)

		}
	}

}

type LiquidResponseTrade struct {
	CreatedAt int     `json:"created_at"`
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
	TakerSide string  `json:"taker_side"`
	Timestamp string  `json:"timestamp"`
}

type LiquidResponse struct {
	Channel string `json:"channel,omitempty"`
	Data    string `json:"data,omitempty"`
	Event   string `json:"event,omitempty"`
}

func (scraper *QuoineScraper) mainLoop() {
	for true {

		var message LiquidResponse

		err := scraper.wsClient.ReadJSON(&message)
		if err != nil {
			log.Errorln("Error reading JSON", err)
		}
		switch message.Event {

		case "created":
			var data LiquidResponseTrade
			e := json.Unmarshal([]byte(message.Data), &data)
			if err != nil {
				log.Errorln("Error Unmarshalling Trade", e)
				continue
			}

			pairScraper := scraper.pairScrapers[message.Channel]

			volume := data.Quantity

			if data.TakerSide == "sell" {
				volume = -volume
			}

			trade := &dia.Trade{
				Symbol:         pairScraper.pair.Symbol,
				Pair:           pairScraper.pair.ForeignName,
				Price:          data.Price,
				Volume:         volume,
				Time:           time.Unix(int64(data.CreatedAt), 0),
				ForeignTradeID: strconv.Itoa(int(data.ID)),
				Source:         scraper.exchangeName,
			}

			pairScraper.parent.chanTrades <- trade
		}

	}
	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close()")
	}
	scraper.cleanup(nil)
}

func (s *QuoineScraper) NormalizePair(pair dia.Pair) (dia.Pair, error) {
	symbol := pair.Symbol
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

func getLiquidProducts() (products LiquidProducts, err error) {
	var response []byte
	response, err = utils.GetRequest(LiquidSocketRestURL + "/products")
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &products)
	return

}

func (scraper *QuoineScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	var products LiquidProducts
	products, err = getLiquidProducts()
	if err != nil {
		return
	}

	pairs = make([]dia.Pair, len(products))

	for _, prod := range products {
		pairToNormalize := dia.Pair{
			Symbol:      prod.BaseCurrency,
			ForeignName: prod.CurrencyPairCode,
			Exchange:    scraper.exchangeName,
		}
		pair, serr := scraper.NormalizePair(pairToNormalize)
		if serr == nil {
			pairs = append(pairs, pair)
		} else {
			log.Error(serr)
		}
	}
	return
}

func (scraper *QuoineScraper) readProductIds() error {

	products, err := getLiquidProducts()
	if err != nil {
		return err
	}

	for _, prod := range products {
		// create a pair -> id mapping
		scraper.productPairIds[prod.CurrencyPairCode] = prod.PusherChannel

	}

	return nil
}

func (scraper *QuoineScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("Quoine scraper is closed")
	}

	pairScraper := &QuoinePairScraper{
		parent: scraper,
		pair:   pair,
	}

	channelName := "executions_cash_" + strings.ToLower(pair.ForeignName)

	channel := LiquidChannel{
		Channel: channelName,
	}

	a := &LiquidSubscribe{
		Event: "pusher:subscribe",
		Data:  channel,
	}

	if err := scraper.wsClient.WriteJSON(a); err != nil {
		log.Errorln(err.Error())
	}
	scraper.pairScrapers[channelName] = pairScraper

	return pairScraper, nil
}
func (s *QuoineScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (scraper *QuoineScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type QuoinePairScraper struct {
	parent *QuoineScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *QuoinePairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (pairScraper *QuoineScraper) Channel() chan *dia.Trade {
	return pairScraper.chanTrades
}

func (pairScraper *QuoinePairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *QuoinePairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
