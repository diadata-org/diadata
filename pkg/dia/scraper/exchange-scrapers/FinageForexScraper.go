package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gorilla/websocket"
)

type FinageWSMessage struct {
	Action  string `json:"action"`
	Symbols string `json:"symbols"`
}

type FinageTrade struct {
	Symbol    string  `json:"s"`
	PriceAsk  float64 `json:"a"`
	PriceBid  float64 `json:"b"`
	VolumeAsk string  `json:"dc"`
	VolumeBid string  `json:"dd"`
	Timestamp int64   `json:"t"`
}

// var pairs = "INR/USD,USD/INR,AED/INR, INR/AED,INR/CAD,CAD/INR"

type FinageForexScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*FinageForexPairScraper // dia.ExchangePair -> pairScraperSet
	ticker       *time.Ticker
	datastore    *models.RelDB
	chanTrades   chan *dia.Trade
	wsConn       *websocket.Conn
	exchangeName string
	apiKey       string
}

// SpawnECBScraper returns a new ECBScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewFinageForexScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB, finageAPIkey string, finageWebsocketKey string) *FinageForexScraper {
	var finage = "wss://w29hxx2ndd.finage.ws:8001/?token=" + finageWebsocketKey

	c, _, err := websocket.DefaultDialer.Dial(finage, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	scraper := &FinageForexScraper{
		wsConn:       c,
		shutdown:     make(chan nothing),
		exchangeName: exchange.Name,
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*FinageForexPairScraper),
		error:        nil,
		ticker:       time.NewTicker(refreshDelay),
		chanTrades:   make(chan *dia.Trade),
		datastore:    relDB,
		apiKey:       finageAPIkey,
	}

	log.Info("Scraper is built and initiated")
	if scrape {
		go scraper.mainLoop()
	}
	return scraper
}

func (scraper *FinageForexScraper) subscribe() error {

	pairTosubscribe := ""

	pairs, err := scraper.datastore.GetExchangePairSymbols(scraper.exchangeName)
	if err != nil {
		log.Errorln("Error getting pairs", err)
		return err
	}

	log.Println("Pairs", pairs)

	for _, ps := range pairs {
		pairTosubscribe = pairTosubscribe + "," + ps.ForeignName
	}
	log.Infoln("pairTosubscribe", pairTosubscribe)
	return scraper.wsConn.WriteJSON(FinageWSMessage{Action: "subscribe", Symbols: pairTosubscribe})

}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *FinageForexScraper) mainLoop() {
	subscribeErr := scraper.subscribe()
	if subscribeErr != nil {
		log.Error("got error subscribing to scraper ", subscribeErr)
	}
	log.Infoln("Sunbscribed to all asset pairs")
	err := scraper.Update()
	if err != nil {
		log.Error(err)
	}
	for {
		select {
		case <-scraper.shutdown: // user requested shutdown
			log.Println("FinageScraper shutting down")
			scraper.cleanup(nil)
			return
		}
	}
}

// Update performs a HTTP Get request for the rss feed and decodes the results.
func (scraper *FinageForexScraper) Update() error {

	go func() {
		for {
			_, message, err := scraper.wsConn.ReadMessage()
			if err != nil {
				//s.subscribe()
				log.Println("err", err)
			}

			var ftrade FinageTrade
			err = json.Unmarshal(message, &ftrade)
			log.Info("Symbol: ", ftrade.Symbol)
			log.Info("PriceAsk: ", ftrade.PriceAsk)
			log.Info("PriceBid: ", ftrade.PriceBid)
			log.Info("VolumeAsk: ", ftrade.VolumeAsk)
			log.Info("VolumeBid: ", ftrade.VolumeBid)
			log.Info("Timestamp: ", ftrade.Timestamp)

			if err != nil {
				log.Errorln("Not a Trade", err)
				break
			} else {
				tradePair, _ := scraper.datastore.GetExchangePairCache(scraper.exchangeName, strings.Replace(ftrade.Symbol, "/", "-", 1))
				if ftrade.Symbol != "" {
					t := &dia.Trade{
						Symbol:       strings.Split(ftrade.Symbol, "/")[0],
						Pair:         strings.Replace(ftrade.Symbol, "/", "-", 1),
						Price:        ftrade.PriceAsk,
						Volume:       1,
						BaseToken:    tradePair.UnderlyingPair.BaseToken,
						QuoteToken:   tradePair.UnderlyingPair.QuoteToken,
						VerifiedPair: tradePair.Verified,
						Time:         time.Unix(ftrade.Timestamp/1e3, 0),
						Source:       scraper.exchangeName,
					}
					if t.VerifiedPair {
						log.Info("got verified trade: ", t)
					}
					scraper.chanTrades <- t
				}

			}
		}
	}()

	return nil
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (scraper *FinageForexScraper) cleanup(err error) {

	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()

	scraper.ticker.Stop()

	if err != nil {
		scraper.error = err
	}
	scraper.closed = true

	close(scraper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (scraper *FinageForexScraper) Close() error {
	if scraper.closed {
		return errors.New("FinageForexScraper: Already closed")
	}
	close(scraper.shutdown)
	<-scraper.shutdownDone
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()
	return scraper.error
}

// ECBPairScraper implements PairScraper for ECB
type FinageForexPairScraper struct {
	parent *FinageForexScraper
	pair   dia.ExchangePair
	closed bool
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (scraper *FinageForexScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()
	if scraper.error != nil {
		return nil, scraper.error
	}
	if scraper.closed {
		return nil, errors.New("ECBScraper: Call ScrapePair on closed scraper")
	}
	ps := &FinageForexPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (scraper *FinageForexScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *FinageForexPairScraper) Close() error {
	pairScraper.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (pairScraper *FinageForexPairScraper) Error() error {
	scraper := pairScraper.parent
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()
	return scraper.error
}

// Pair returns the pair this scraper is subscribed to
func (pairScraper *FinageForexPairScraper) Pair() dia.ExchangePair {
	return pairScraper.pair
}

type FinageSymbolResponse struct {
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	Symbols   []struct {
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	} `json:"symbols"`
}

func (scraper *FinageForexScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	log.Infoln("Fetching Available Pairs")
	var finageurl = "https://api.finage.co.uk/symbol-list/forex?apikey=" + scraper.apiKey + "&page="
	var finageSymbolResponse FinageSymbolResponse
	data, _, err := utils.GetRequest(finageurl + strconv.Itoa(1))
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &finageSymbolResponse)
	if err != nil {
		return
	}

	for i := 1; i <= finageSymbolResponse.TotalPage; i++ {

		data, _, err := utils.GetRequest(finageurl + strconv.Itoa(i))
		if err != nil {
			continue
		}

		err = json.Unmarshal(data, &finageSymbolResponse)
		if err != nil {
			log.Error("unmarshal pair: ", err)
			continue
		}

		for _, symbol := range finageSymbolResponse.Symbols {

			pairToNormalize := dia.ExchangePair{
				Symbol:      "",
				ForeignName: symbol.Symbol,
				Exchange:    scraper.exchangeName,
			}
			pair, serr := scraper.NormalizePair(pairToNormalize)
			if serr == nil {
				pairs = append(pairs, pair)
			} else {
				log.Error(serr)
			}

		}

	}

	return
}

func (scraper *FinageForexScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	asset.Symbol = symbol
	asset.Blockchain = dia.FIAT
	return asset, nil
}

func (scraper *FinageForexScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	runes := []rune(pair.ForeignName)
	asset0 := runes[0:3]
	asset1 := runes[3:6]
	pair.ForeignName = string(asset0) + "/" + string(asset1)
	pair.Symbol = string(asset0)
	pair.Verified = true
	return pair, nil
}
