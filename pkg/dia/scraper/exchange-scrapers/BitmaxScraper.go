package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	ws "github.com/gorilla/websocket"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	bitmaxMaxNumSubscriptionsPerConn = 200
)

type BitMaxPairResponse struct {
	Code int          `json:"code"`
	Data []BitMaxPair `json:"data"`
}

type BitMaxAssets struct {
	Code int           `json:"code"`
	Data []BitMaxAsset `json:"data"`
}

type BitMaxAsset struct {
	AssetCode        string `json:"assetCode"`
	AssetName        string `json:"assetName"`
	PrecisionScale   int    `json:"precisionScale"`
	NativeScale      int    `json:"nativeScale"`
	WithdrawalFee    string `json:"withdrawalFee"`
	MinWithdrawalAmt string `json:"minWithdrawalAmt"`
	Status           string `json:"status"`
}

type BitMaxPair struct {
	Symbol                string `json:"symbol"`
	DisplayName           string `json:"displayName"`
	BaseAsset             string `json:"baseAsset"`
	QuoteAsset            string `json:"quoteAsset"`
	Status                string `json:"status"`
	MinNotional           string `json:"minNotional"`
	MaxNotional           string `json:"maxNotional"`
	MarginTradable        bool   `json:"marginTradable"`
	CommissionType        string `json:"commissionType"`
	CommissionReserveRate string `json:"commissionReserveRate"`
	TickSize              string `json:"tickSize"`
	LotSize               string `json:"lotSize"`
}

type BitMaxScraper struct {
	// signaling channels for session initialization and finishing
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	// use sync.Maps to concurrently handle multiple pairs
	pairScrapers           map[string]*BitMaxPairScraper // dia.Pair -> BitMaxPairScraper
	exchangeName           string
	chanTrades             chan *dia.Trade
	wsClient1              *ws.Conn
	wsClient2              *ws.Conn
	numPairsClient1        int
	numPairsClient2        int
	currencySymbolName     map[string]string
	isTickerMapInitialised bool
	db                     *models.RelDB
}

func NewBitMaxScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *BitMaxScraper {
	var bitmaxSocketURL = "wss://ascendex.com/0/api/pro/v1/stream"
	s := &BitMaxScraper{
		initDone:               make(chan nothing),
		shutdown:               make(chan nothing),
		shutdownDone:           make(chan nothing),
		exchangeName:           exchange.Name,
		pairScrapers:           make(map[string]*BitMaxPairScraper),
		error:                  nil,
		chanTrades:             make(chan *dia.Trade),
		currencySymbolName:     make(map[string]string),
		isTickerMapInitialised: false,
		db:                     relDB,
	}

	// establish connection in the background
	var wsDialer ws.Dialer
	SwConn1, _, err := wsDialer.Dial(bitmaxSocketURL, nil)
	if err != nil {
		log.Fatal("connect to websocket server: ", err)
	}
	s.wsClient1 = SwConn1
	SwConn2, _, err := wsDialer.Dial(bitmaxSocketURL, nil)
	if err != nil {
		log.Fatal("connect to websocket server: ", err)
	}
	s.wsClient2 = SwConn2

	if scrape {
		go s.mainLoop(s.wsClient1)
		go s.mainLoop(s.wsClient2)
	}
	return s
}

type BitMaxTradeResponse struct {
	M      string `json:"m"`
	Symbol string `json:"symbol"`
	Data   []struct {
		P      string `json:"p"`
		Q      string `json:"q"`
		Ts     int64  `json:"ts"`
		Bm     bool   `json:"bm"`
		Seqnum int64  `json:"seqnum"`
	} `json:"data"`
}

// runs in a goroutine until s is closed
func (s *BitMaxScraper) mainLoop(client *ws.Conn) {
	var err error
	for {
		message := &BitMaxTradeResponse{}
		if err = client.ReadJSON(&message); err != nil {
			log.Error("read message: ", err.Error())
			// break
		}
		switch message.M {
		case "trades":
			{
				for _, trade := range message.Data {
					var exchangepair dia.ExchangePair
					priceFloat, _ := strconv.ParseFloat(trade.P, 64)
					volumeFloat, _ := strconv.ParseFloat(trade.Q, 64)
					exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, message.Symbol)
					if err != nil {
						log.Error("get exchange pair from cache: ", err)
					}
					t := &dia.Trade{
						Symbol:         strings.Split(message.Symbol, "/")[0],
						Pair:           message.Symbol,
						Price:          priceFloat,
						Volume:         volumeFloat,
						Time:           time.Unix(0, trade.Ts*int64(time.Millisecond)),
						ForeignTradeID: strconv.FormatInt(trade.Seqnum, 10),
						Source:         s.exchangeName,
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
		case "ping":
			{
				a := &BitMaxRequest{
					Op: "pong",
				}
				err := client.WriteJSON(a)
				if err != nil {
					log.Warn("send pong to server: ", err)
				}
				log.Infoln("Send Pong to keep connection alive")

			}
		}
	}

}

// FillSymbolData collects all available information on an asset traded on Bitmax
func (s *BitMaxScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {

	// Fetch Data
	if !s.isTickerMapInitialised {
		var (
			response BitMaxAssets
			data     []byte
		)
		data, _, err = utils.GetRequest("https://ascendex.com/api/pro/v1/assets")
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &response)
		if err != nil {
			return
		}

		for _, asset := range response.Data {
			s.currencySymbolName[asset.AssetCode] = asset.AssetName
		}
		s.isTickerMapInitialised = true

	}

	asset.Symbol = symbol
	asset.Name = s.currencySymbolName[symbol]
	return asset, nil
}

func (s *BitMaxScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BitMaxScraper) Close() error {
	if s.closed {
		return errors.New("BitMaxScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

type BitMaxRequest struct {
	Op string `json:"op"`
	ID string `json:"id"`
	Ch string `json:"ch"`
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BitMaxScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("LoopringScraper: Call ScrapePair on closed scraper")
	}
	ps := &BitMaxPairScraper{
		parent: s,
		pair:   pair,
	}
	a := &BitMaxRequest{
		Op: "sub",
		Ch: "trades:" + pair.ForeignName,
		ID: fmt.Sprint(time.Now().Unix()),
	}
	if s.numPairsClient1 < bitmaxMaxNumSubscriptionsPerConn {
		if err := s.wsClient1.WriteJSON(a); err != nil {
			log.Error("write pair sub: ", err.Error())
		}
		s.numPairsClient1++
	} else {
		if err := s.wsClient2.WriteJSON(a); err != nil {
			log.Error("write pair sub: ", err.Error())
		}
		s.numPairsClient2++
	}
	log.Info("Subscribed to get trades for ", pair.ForeignName)
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

func (s *BitMaxScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	var bitmaxResponse BitMaxPairResponse
	response, err := http.Get("https://ascendex.com/api/pro/v1/products")
	if err != nil {
		log.Error("get symbols: ", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("read symbols: ", err)
	}

	err = json.Unmarshal(body, &bitmaxResponse)
	if err != nil {
		log.Error("unmarshal symbols: ", err)
	}

	for _, p := range bitmaxResponse.Data {
		pairToNormalize := dia.ExchangePair{
			Symbol:      strings.Split(p.Symbol, "/")[0],
			ForeignName: p.Symbol,
			Exchange:    s.exchangeName,
		}
		pairs = append(pairs, pairToNormalize)
	}
	return
}

// BitMax implements PairScraper for BitMax
type BitMaxPairScraper struct {
	parent *BitMaxScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *BitMaxPairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("BitMaxPairScraper: Already closed")
	}

	// TODO stop collection for the pair

	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *BitMaxScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BitMaxPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BitMaxPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
