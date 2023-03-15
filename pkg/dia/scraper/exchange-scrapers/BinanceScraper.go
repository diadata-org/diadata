package scrapers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
	"github.com/zekroTJA/timedmap"
)

type binancePairScraperSet map[*BinancePairScraper]nothing

// BinanceScraper is a Scraper for collecting trades from the Binance websocket API
type BinanceScraper struct {
	client *binance.Client
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
	pairScrapers sync.Map // dia.ExchangePair -> binancePairScraperSet
	// pairSubscriptions sync.Map // dia.ExchangePair -> string (subscription ID)
	// pairLocks         sync.Map // dia.ExchangePair -> sync.Mutex
	exchangeName string
	scraperName  string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

// NewBinanceScraper returns a new BinanceScraper for the given pair
func NewBinanceScraper(apiKey string, secretKey string, exchange dia.Exchange, scraperName string, scrape bool, relDB *models.RelDB) *BinanceScraper {

	s := &BinanceScraper{
		client:       binance.NewClient(apiKey, secretKey),
		initDone:     make(chan nothing),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		exchangeName: exchange.Name,
		scraperName:  scraperName,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}

	var err error
	reverseBasetokens, err = getReverseTokensFromConfigFull("binance/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse basetokens: ", reverseBasetokens)

	// establish connection in the background
	if scrape {
		go s.mainLoop()
	}
	return s
}

func (up *BinanceScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	if pair.Symbol == "MIOTA" {
		pair.ForeignName = "M" + pair.ForeignName
	}
	if pair.Symbol == "YOYOW" {
		pair.ForeignName = "YOYOW" + pair.ForeignName[4:]
	}
	if pair.Symbol == "ETHOS" {
		pair.ForeignName = "ETHOS" + pair.ForeignName[3:]
	}
	if pair.Symbol == "WNXM" {
		pair.Symbol = "wNXM"
		pair.ForeignName = "wNXM" + pair.ForeignName[4:]
	}
	return pair, nil
}

// runs in a goroutine until s is closed
func (s *BinanceScraper) mainLoop() {
	close(s.initDone)
	for range s.shutdown { // user requested shutdown
		log.Println("BinanceScraper shutting down")
		s.cleanup()
		return
	}
	select {}
}

func (s *BinanceScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	// TO DO
	return dia.Asset{Symbol: symbol}, nil
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *BinanceScraper) cleanup() {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	// close all channels of PairScraper children
	s.pairScrapers.Range(func(k, v interface{}) bool {
		for ps := range v.(binancePairScraperSet) {
			ps.closed = true
		}
		s.pairScrapers.Delete(k)
		return true
	})

	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *BinanceScraper) Close() error {
	if s.closed {
		return errors.New("BinanceScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BinanceScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	<-s.initDone // wait until client is connected

	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	if s.closed {
		return nil, errors.New("BinanceScraper: Call ScrapePair on closed scraper")
	}

	ps := &BinancePairScraper{
		parent: s,
		pair:   pair,
	}

	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
		var exchangepair dia.ExchangePair

		volume, err := strconv.ParseFloat(event.Quantity, 64)
		price, err2 := strconv.ParseFloat(event.Price, 64)

		if err == nil && err2 == nil && event.Event == "aggTrade" {
			if !event.IsBuyerMaker {
				volume = -volume
			}
			pairNormalized, _ := s.NormalizePair(pair)
			exchangepair, err = s.db.GetExchangePairCache(s.scraperName, pair.ForeignName)
			if err != nil {
				log.Error(err)
			}
			t := &dia.Trade{
				Symbol:         pairNormalized.Symbol,
				Pair:           pairNormalized.ForeignName,
				Price:          price,
				Volume:         volume,
				Time:           time.Unix(event.TradeTime/1000, (event.TradeTime%1000)*int64(time.Millisecond)),
				ForeignTradeID: strconv.FormatInt(event.AggTradeID, 16),
				Source:         s.exchangeName,
				VerifiedPair:   exchangepair.Verified,
				BaseToken:      exchangepair.UnderlyingPair.BaseToken,
				QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
			}

			if utils.Contains(reverseBasetokens, t.BaseToken.Identifier()) {
				// If we need quotation of a base token, reverse pair
				tSwapped, errSwap := dia.SwapTrade(*t)
				if errSwap == nil {
					t = &tSwapped
				}
			}

			if exchangepair.Verified {
				log.Infoln("Got verified trade", t)
			}

			// Handle duplicate trades.
			discardTrade := t.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
			if !discardTrade {
				t.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
				ps.parent.chanTrades <- t
			}
		} else {
			log.Println("ignoring event ", event, err, err2)
		}
	}
	errHandler := func(err error) {
		log.Error(err)
	}

	_, _, err := binance.WsAggTradeServe(pair.ForeignName, wsAggTradeHandler, errHandler)
	if err != nil {
		log.Errorf("serving pair %s", pair.ForeignName)
	}

	return ps, err
}
func (s *BinanceScraper) normalizeSymbol(p dia.ExchangePair, foreignName string, params ...string) (pair dia.ExchangePair, err error) {
	// symbol := p.Symbol
	// status := params[0]
	// if status == "TRADING" {
	// 	if helpers.NameForSymbol(symbol) == symbol {
	// 		if !helpers.SymbolIsName(symbol) {
	// 			pair.Symbol = symbol
	// 			pair, _ = s.NormalizePair(pair)

	// 			return pair, errors.New("Foreign name can not be normalized:" + foreignName + " symbol:" + symbol)
	// 		}
	// 	}
	// 	if helpers.SymbolIsBlackListed(symbol) {
	// 		pair.Symbol = symbol
	// 		return pair, errors.New("Symbol is black listed:" + symbol)
	// 	}
	// } else {
	// 	return pair, errors.New("Symbol:" + symbol + " with foreign name:" + foreignName + " is:" + status)

	// }
	return pair, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BinanceScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	data, _, err := utils.GetRequest("https://api.binance.com/api/v1/exchangeInfo")

	if err != nil {
		return
	}
	var ar binance.ExchangeInfo
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar.Symbols {

			pairToNormalise := dia.ExchangePair{
				Symbol:      p.Symbol,
				ForeignName: p.Symbol,
				Exchange:    s.exchangeName,
			}

			pair, serr := s.normalizeSymbol(pairToNormalise, p.BaseAsset, p.Status)
			if serr == nil {
				pairs = append(pairs, pair)
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// BinancePairScraper implements PairScraper for Binance
type BinancePairScraper struct {
	parent *BinanceScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *BinancePairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("BinancePairScraper: Already closed")
	}

	// TODO stop collection for the pair

	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *BinanceScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BinancePairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BinancePairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

// getReverseTokensFromConfigFull returns a list of addresses from config file.
func getReverseTokensFromConfigFull(filename string) (*[]string, error) {

	var reverseTokens []string

	// Load file and read data
	filehandle := configCollectors.ConfigFileConnectors(filename, ".json")
	jsonFile, err := os.Open(filehandle)
	if err != nil {
		return &[]string{}, err
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &[]string{}, err
	}

	// // Unmarshal read data
	// type lockedAsset struct {
	// 	Address    string `json:"Address"`
	// 	Blockchain string `json:"Blockchain"`
	// 	Symbol     string `json:"Symbol"`
	// }
	type lockedAssetList struct {
		AllAssets []dia.Asset `json:"Tokens"`
	}
	var allAssets lockedAssetList
	err = json.Unmarshal(byteData, &allAssets)
	if err != nil {
		return &[]string{}, err
	}

	// Extract addresses
	for _, token := range allAssets.AllAssets {
		reverseTokens = append(reverseTokens, token.Identifier())
	}

	return &reverseTokens, nil
}
