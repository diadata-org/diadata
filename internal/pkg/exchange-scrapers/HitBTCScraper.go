package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
)

var _socketurl string = "wss://api.hitbtc.com/api/2/ws"

const WS_TIMEOUT = 10 * time.Second

type Event struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	Id     int         `json:"id"`
}

type HitBTCAsset struct {
	ID                 string `json:"id"`
	FullName           string `json:"fullName"`
	Crypto             bool   `json:"crypto"`
	PayinEnabled       bool   `json:"payinEnabled"`
	PayinPaymentID     bool   `json:"payinPaymentId"`
	PayinConfirmations int    `json:"payinConfirmations"`
	PayoutEnabled      bool   `json:"payoutEnabled"`
	PayoutIsPaymentID  bool   `json:"payoutIsPaymentId"`
	TransferEnabled    bool   `json:"transferEnabled"`
	Delisted           bool   `json:"delisted"`
	PayoutFee          string `json:"payoutFee"`
	PrecisionPayout    int    `json:"precisionPayout"`
	PrecisionTransfer  int    `json:"precisionTransfer"`
}

type HitBTCScraper struct {
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
	pairScrapers map[string]*HitBTCPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewHitBTCScraper returns a new HitBTCScraper for the given pair
func NewHitBTCScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *HitBTCScraper {

	s := &HitBTCScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*HitBTCPairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(_socketurl, nil)
	if err != nil {
		println(err.Error())
	}
	s.wsClient = SwConn
	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *HitBTCScraper) mainLoop() {
	var err error
	for {
		message := &Event{}
		if err = s.wsClient.ReadJSON(&message); err != nil {
			log.Error(err.Error())
			break
		}
		if message.Method == "updateTrades" {
			md := message.Params.(map[string]interface{})
			ps, ok := s.pairScrapers[md["symbol"].(string)]
			if ok {
				mdData := md["data"].([]interface{})
				for _, v := range mdData {
					var f64Price float64
					var f64Volume float64
					mdElement := v.(map[string]interface{})
					f64PriceString := mdElement["price"].(string)
					f64Price, err = strconv.ParseFloat(f64PriceString, 64)
					if err == nil {
						f64VolumeString := mdElement["quantity"].(string)
						f64Volume, err = strconv.ParseFloat(f64VolumeString, 64)
						if err == nil {
							timeStamp, _ := time.Parse(time.RFC3339, mdElement["timestamp"].(string))
							if mdElement["id"] != 0 {
								if mdElement["side"] == "sell" {
									f64Volume = -f64Volume
								}

								exchangepair, err := s.db.GetExchangePairCache(s.exchangeName, md["symbol"].(string))
								if err != nil {
									log.Error(err)
								}
								t := &dia.Trade{
									Symbol:         ps.pair.Symbol,
									Pair:           md["symbol"].(string),
									Price:          f64Price,
									Volume:         f64Volume,
									Time:           timeStamp,
									ForeignTradeID: strconv.FormatInt(int64(mdElement["id"].(float64)), 16),
									Source:         s.exchangeName,
									VerifiedPair:   exchangepair.Verified,
									BaseToken:      exchangepair.UnderlyingPair.BaseToken,
									QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
								}
								if exchangepair.Verified {
									log.Infoln("Got verified trade", t)
								}
								ps.parent.chanTrades <- t
							}
						} else {
							log.Error("error parsing volume " + mdElement["quantity"].(string))
						}
					} else {
						log.Error("error parsing price " + mdElement["price"].(string))
					}
				}
			} else {
				log.Error("Unknown Pair " + md["symbol"].(string))
			}
		}
	}
	s.cleanup(err)
}

func (s *HitBTCScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone)
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *HitBTCScraper) Close() error {
	if s.closed {
		return errors.New("HitBTCScraper: Already closed")
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

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *HitBTCScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("HitBTCScraper: Call ScrapePair on closed scraper")
	}

	ps := &HitBTCPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	a := &Event{
		Method: "subscribeTrades",
		Params: map[string]interface{}{
			"symbol": pair.ForeignName,
		},
		Id: int(time.Now().Unix()) * 1000,
	}

	if err := s.wsClient.WriteJSON(a); err != nil {
		fmt.Println(err.Error())
	}

	return ps, nil
}

// func (s *HitBTCScraper) normalizeSymbol(foreignName string, baseCurrency string) (symbol string, err error) {
// 	symbol = strings.ToUpper(baseCurrency)
// 	if helpers.NameForSymbol(symbol) == symbol {
// 		if !helpers.SymbolIsName(symbol) {
// 			return symbol, errors.New("Foreign name can not be normalized:" + foreignName + " symbol:" + symbol)
// 		}
// 	}
// 	if helpers.SymbolIsBlackListed(symbol) {
// 		return symbol, errors.New("Symbol is black listed:" + symbol)
// 	}
// 	return symbol, nil
// }

func (s *HitBTCScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
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

// FetchAvailablePairs returns a list with all available trade pairs
func (s *HitBTCScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	type APIResponse struct {
		Id                   string  `json:"id"`
		BaseCurrency         string  `json:"baseCurrency"`
		QuoteCurrency        string  `json:"quoteCurrency"`
		QuantityIncrement    float64 `json:"quantityIncrement,string"`
		TickSize             float64 `json:"tickSize,string"`
		TakeLiquidityRate    float64 `json:"takeLiquidityRate,string"`
		ProvideLiquidityRate float64 `json:"provideLiquidityRate,string"`
		FeeCurrency          string  `json:"feeCurrency"`
	}
	data, _, err := utils.GetRequest("https://api.hitbtc.com/api/2/public/symbol")
	if err != nil {
		return
	}
	var ar []APIResponse
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar {
			pairToNormalize := dia.ExchangePair{
				Symbol:      p.BaseCurrency,
				ForeignName: p.Id,
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

// HitBTCPairScraper implements PairScraper for HitBTC
type HitBTCPairScraper struct {
	parent *HitBTCScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *HitBTCPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *HitBTCScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// FillSymbolData collects all available information on an asset traded on HitBTC
func (ps *HitBTCScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	var response HitBTCAsset
	data, _, err := utils.GetRequest("https://api.hitbtc.com/api/2/public/currency/" + symbol)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return
	}
	asset.Symbol = response.ID
	asset.Name = response.FullName
	return asset, nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *HitBTCPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *HitBTCPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
