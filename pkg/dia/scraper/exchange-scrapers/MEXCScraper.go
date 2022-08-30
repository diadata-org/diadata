package scrapers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	ws "github.com/gorilla/websocket"
)

const mexc_socketurl string = "wss://wbs.mexc.com/ws"

const api_url = "https://api.mexc.com"

type MEXCExchangeSymbol struct {
	Symbol                     string   `json:"symbol"`
	Status                     string   `json:"status"`
	BaseAsset                  string   `json:"baseAsset"`
	BaseAssetPrecision         int      `json:"baseAssetPrecision"`
	QuoteAsset                 string   `json:"quoteAsset"`
	QuotePrecision             int      `json:"quotePrecision"`
	QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
	OrderTypes                 []string `json:"orderTypes"` // [LIMIT, LIMIT_MAKER]
	QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
	QuoteAmountPrecision       string   `json:"quoteAmountPrecision"`
	BaseSizePrecision          string   `json:"baseSizePrecision"`
	Permissions                []string `json:"permissions"`
	Filters                    []string `json:"filters"`
	MaxQuoteAmount             string   `json:"maxQuoteAmount"`
	MakerCommission            string   `json:"makerCommission"`
	TakerCommission            string   `json:"takerCommission"`
}

type MEXCExchangeInfo struct {
	Timezone        string               `json:"timezone"`
	ServerTime      int                  `json:"serverTime"`
	RateLimits      string               `json:"rateLimits"`
	ExchangeFilters string               `json:"exchangeFilters"`
	Symbols         []MEXCExchangeSymbol `json:"symbols"`
}

// MEXCScraper is a scraper for MEXC
type MEXCScraper struct {
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
	pairScrapers map[string]*MEXCPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

func NewMEXCScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *MEXCScraper {
	s := &MEXCScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*MEXCPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(mexc_socketurl, nil)

	if err != nil {
		println(err.Error())
	}
	s.wsClient = SwConn

	if scrape {
		go s.mainLoop()
	}

	return s
}

type MEXCTradeResponse struct {
	C string `json:"c"`
	D struct {
		Deals []struct {
			T  int    `json:"T"`
			P  string `json:"p"`
			Q  string `json:"q"`
			TS int64  `json:"t"`
		} `json:"deals"`
	} `json:"d"`
	S string `json:"s"`
}

func (s *MEXCScraper) mainLoop() {
	var err error
	for {
		message := &MEXCTradeResponse{}
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error("read message: ", err.Error())
			continue
			// deal it
		}
		log.Info(message)
		for _, trade := range message.D.Deals {
			var exchangePair dia.ExchangePair
			priceFloat, _ := strconv.ParseFloat(trade.P, 64)
			volumeFloat, _ := strconv.ParseFloat(trade.Q, 64)
			exchangePair, err = s.db.GetExchangePairCache(s.exchangeName, message.S)
			if err != nil {
				log.Error(err)
			}
			t := &dia.Trade{
				Symbol: strings.Split(message.S, "_")[0],
				Pair:   message.S,
				Price:  priceFloat,
				Volume: volumeFloat,
				Time:   time.Unix(0, trade.TS*int64(time.Millisecond)),
				// ForeignTradeID: ,
				Source:       s.exchangeName,
				VerifiedPair: exchangePair.Verified,
				BaseToken:    exchangePair.UnderlyingPair.BaseToken,
				QuoteToken:   exchangePair.UnderlyingPair.QuoteToken,
			}

			if exchangePair.Verified {
				log.Infoln("Got verified trade", t)
			}
			s.chanTrades <- t
		}
	}
}

// FillSymbolData from MEXCScraper
// @todo more update
func (s *MEXCScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	asset.Symbol = symbol
	return
}

func (s *MEXCScraper) Close() error {
	if s.closed {
		return errors.New("MEXCScraper: Already closed")
	}
	close(s.shutdown)
	err := s.wsClient.Close()
	if err != nil {
		return err
	}

	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *MEXCScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

type MEXCRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int64    `json:"id"`
}

func (s *MEXCScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("MEXCScraper: Call ScrapePair on closed scraper")
	}

	ps := &MEXCPairScraper{
		parent: s,
		pair:   pair,
	}

	a := &MEXCRequest{
		Method: "SUBSCRIPTION",
		Params: []string{"spot@public.aggre.deals@" + pair.ForeignName},
		ID:     time.Now().Unix(),
	}

	log.Info("spot@public.aggre.deals@" + pair.ForeignName)

	if err := s.wsClient.WriteJSON(a); err != nil {
		log.Error("write pair sub: ", err.Error())
	}
	log.Info("Subscribed to get trades for ", pair.ForeignName)
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (s *MEXCScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	var mexcExchangeInfo MEXCExchangeInfo
	response, err := http.Get(api_url + "/api/v3/exchangeInfo")
	if err != nil {
		log.Error("get symbols: ", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Error("read symbols: ", err)
	}

	err = json.Unmarshal(body, &mexcExchangeInfo)

	if err != nil {
		log.Error("unmarshal symbols: ", err)
	}

	for _, p := range mexcExchangeInfo.Symbols {
		pairToNormalized := dia.ExchangePair{
			Symbol:      p.BaseAsset,
			ForeignName: p.BaseAsset + "_" + p.QuoteAsset,
			Exchange:    s.exchangeName,
		}
		pairs = append(pairs, pairToNormalized)
	}
	return
}

// Channel returns a channel that can be used to receive trades
func (s *MEXCScraper) Channel() chan *dia.Trade {
	return s.chanTrades
}

// MEXCPairScraper implements PairScraper for MEXC
type MEXCPairScraper struct {
	parent *MEXCScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *MEXCPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *MEXCPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *MEXCPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
