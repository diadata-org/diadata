package scrapers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"
)

const (
	BINANCE_API_MAX_RETRIES = 5
)

var (
	binanceWSBaseString = "wss://stream.binance.com:9443/ws"
)

type binancePairScraperSet map[*BinancePairScraper]nothing

// BinanceScraper is a Scraper for collecting trades from the Binance websocket API
type BinanceScraper struct {
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	// use sync.Maps to concurrently handle multiple pairs
	pairScrapers      sync.Map // dia.ExchangePair -> binancePairScraperSet
	newPairScrapers   map[string]*BinancePairScraper
	proxyIndex        int
	exchangeName      string
	scraperName       string
	chanTrades        chan *dia.Trade
	db                *models.RelDB
	wsClient          *ws.Conn
	apiConnectRetries int
}

type binanceWSResponse struct {
	Timestamp      int64       `json:"T"`
	Price          string      `json:"p"`
	Volume         string      `json:"q"`
	ForeignTradeID int         `json:"t"`
	ForeignName    string      `json:"s"`
	Type           interface{} `json:"e"`
	Buy            bool        `json:"m"`
	Ignore         bool        `json:"M"`
}

// BinancePairScraper implements PairScraper for Binance
type BinancePairScraper struct {
	parent *BinanceScraper
	pair   dia.ExchangePair
	closed bool
}

// NewBinanceScraper returns a new BinanceScraper for the given pair
func NewBinanceScraper(apiKey string, secretKey string, exchange dia.Exchange, scraperName string, scrape bool, relDB *models.RelDB) *BinanceScraper {

	s := &BinanceScraper{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		exchangeName:    exchange.Name,
		scraperName:     scraperName,
		error:           nil,
		chanTrades:      make(chan *dia.Trade),
		db:              relDB,
		proxyIndex:      0,
		newPairScrapers: make(map[string]*BinancePairScraper),
	}

	var err error
	reverseBasetokens, err = getReverseTokensFromConfigFull("binance/reverse_tokens/" + s.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}
	log.Info("reverse basetokens: ", reverseBasetokens)

	err = s.connectToAPI()
	if err != nil {
		log.Error("getting an error while connecting to api: ", err)
	} else {
		log.Println("Successfully connect to websocket server.")
	}

	//establish connection in the background
	if scrape {
		go s.mainLoop()
	}

	return s
}

func (scraper *BinanceScraper) subscribe(pair dia.ExchangePair, subscribe bool) error {
	if scraper.closed {
		return errors.New("binance Scraper: Call ScrapePair on closed scraper")
	}

	// Validate WebSocket connection exists
	if scraper.wsClient == nil {
		return errors.New("WebSocket connection not initialized")
	}

	// Determine subscription type (SUBSCRIBE/UNSUBSCRIBE)
	subscribeType := "UNSUBSCRIBE"
	if subscribe {
		subscribeType = "SUBSCRIBE"
	}
	// Convert symbol+currency to lowercase (e.g., "btcusdt@trade")
	pairTicker := strings.ToLower(pair.ForeignName)

	subscribeMessage := map[string]interface{}{
		"method": subscribeType,
		"params": []string{pairTicker + "@trade"}, //btcusdt@trade
		"id":     time.Now().UnixNano(),
	}
	log.Info("Subscribe Message: ", subscribeMessage)

	if err := scraper.wsClient.WriteJSON(subscribeMessage); err != nil {
		log.Errorf("Failed to send subscription request: %v", err)
		return err
	}
	return nil
}

func (scraper *BinanceScraper) connectToAPI() error {
	log.Info("Starting connect to API")

	// Switch to alternative Proxy whenever too many retries on the first.
	if scraper.apiConnectRetries > BINANCE_API_MAX_RETRIES {
		log.Errorf("too many timeouts for Binance api connection with proxy %v. Switch to alternative proxy.", scraper.proxyIndex)
		scraper.apiConnectRetries = 0
		scraper.proxyIndex = (scraper.proxyIndex + 1) % 2
	}

	username := utils.Getenv("BINANCE_PROXY"+strconv.Itoa(scraper.proxyIndex)+"_USERNAME", "")
	password := utils.Getenv("BINANCE_PROXY"+strconv.Itoa(scraper.proxyIndex)+"_PASSWORD", "")
	user := url.UserPassword(username, password)
	host := utils.Getenv("BINANCE_PROXY"+strconv.Itoa(scraper.proxyIndex)+"_HOST", "")

	log.Infof("User: {%s}, Host: {%s}", user, host)
	var d ws.Dialer
	if host != "" {
		d = ws.Dialer{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http", // or "https" depending on your proxy
				User:   user,
				Host:   host,
				Path:   "/",
			},
			),
		}
	}

	// Connect to Binance API.
	conn, _, err := d.Dial(binanceWSBaseString, nil)
	if err != nil {
		log.Errorf("Binance - Connect to API: %s.", err.Error())
		return err
	}
	scraper.wsClient = conn

	return nil
}

func (up *BinanceScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}

func (s *BinanceScraper) mainLoop() {

	defer func() {
		log.Println("BinanceScraper main loop exiting")
		s.cleanup()
	}()

	for {
		var message binanceWSResponse
		err := s.wsClient.ReadJSON(&message)
		if err != nil {
			log.Error("JSON decode error: ", err)
			continue
		} else {
			log.Info("Successfully received the trade response: ", message)
		}

		if message.ForeignName == "" || message.Price == "" || message.Volume == "" {
			log.Warn("Skipping invalid trade data:", message)
		} else {
			log.Info("Valid trade received: ", message)
			s.parseWSResponse(message)
		}
	}
}

func (s *BinanceScraper) parseWSResponse(message binanceWSResponse) {
	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	var exchangepair dia.ExchangePair
	var err error

	ps := s.newPairScrapers[message.ForeignName]
	pair := ps.pair

	exchangepair, err = s.db.GetExchangePairCache(s.scraperName, message.ForeignName)
	if err != nil {
		log.Error(err)
	}

	tradeTime := time.Unix(0, message.Timestamp*1000000)
	tradePrice, err := strconv.ParseFloat(message.Price, 64)
	if err != nil {
		log.Errorf("Binance - Parse price: %v.", err)
	}
	tradeVolume, err := strconv.ParseFloat(message.Volume, 64)
	if err != nil {
		log.Errorf("Binance - Parse volume: %v.", err)
	}

	if !message.Buy {
		tradeVolume = -tradeVolume
	}
	tradeForeignTradeID := strconv.Itoa(message.ForeignTradeID)

	t := &dia.Trade{
		Symbol:         pair.Symbol,
		Pair:           message.ForeignName,
		Price:          tradePrice,
		Volume:         tradeVolume,
		Time:           tradeTime,
		ForeignTradeID: tradeForeignTradeID,
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
	ps := &BinancePairScraper{
		parent: s,
		pair:   pair,
	}

	s.newPairScrapers[pair.ForeignName] = ps

	if err := s.subscribe(pair, true); err != nil {
		log.Error("Subscription failed:", err)
	}

	//ensure that no more than 5 requests are sent per second(as required by Binance).
	time.Sleep(400 * time.Millisecond)

	return ps, nil
}
func (s *BinanceScraper) normalizeSymbol(p dia.ExchangePair, foreignName string, params ...string) (pair dia.ExchangePair, err error) {
	return pair, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BinanceScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	// data, _, err := utils.GetRequest("https://api.binance.com/api/v1/exchangeInfo")

	// if err != nil {
	// 	return
	// }
	// var ar binance.ExchangeInfo
	// err = json.Unmarshal(data, &ar)
	// if err == nil {
	// 	for _, p := range ar.Symbols {

	// 		pairToNormalise := dia.ExchangePair{
	// 			Symbol:      p.Symbol,
	// 			ForeignName: p.Symbol,
	// 			Exchange:    s.exchangeName,
	// 		}

	// 		pair, serr := s.normalizeSymbol(pairToNormalise, p.BaseAsset, p.Status)
	// 		if serr == nil {
	// 			pairs = append(pairs, pair)
	// 		} else {
	// 			log.Error(serr)
	// 		}
	// 	}
	// }
	return
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
