package scrapers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	"github.com/zekroTJA/timedmap"
)

var (
	CoinExWSBaseString = "wss://socket.coinex.com/v2/spot"
)

type coinExPairScraperSet map[*BinancePairScraper]nothing

type CoinExScraper struct {
	// signaling channels for session finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	// use sync.Maps to concurrently handle multiple pairs
	pairScrapers    sync.Map // dia.ExchangePair -> coinexPairScraperSet
	newPairScrapers map[string]*CoinExPairScraper
	exchangeName    string
	chanTrades      chan *dia.Trade
	db              *models.RelDB
	wsClient        *ws.Conn
}

type CoinExPairScraper struct {
	parent *CoinExScraper
	pair   dia.ExchangePair
	closed bool
}

type SubscribeRequest struct {
	Method string     `json:"method"`
	Params MarketList `json:"params"`
	ID     int        `json:"id"`
}

type MarketList struct {
	MarketList []string `json:"market_list"`
}

type coinexWSResponse struct {
	Data struct {
		ForeignName string `json:"market"`
		DealList    []Deal `json:"deal_list"`
	} `json:"data"`
}

type Deal struct {
	DealID    int64  `json:"deal_id"`
	CreatedAt int64  `json:"created_at"`
	Side      string `json:"side"`
	Price     string `json:"price"`
	Amount    string `json:"amount"`
}

func NewCoinExScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *CoinExScraper {

	s := &CoinExScraper{
		shutdown:        make(chan nothing),
		shutdownDone:    make(chan nothing),
		exchangeName:    exchange.Name,
		error:           nil,
		chanTrades:      make(chan *dia.Trade),
		db:              relDB,
		newPairScrapers: make(map[string]*CoinExPairScraper),
	}

	err := s.connectToAPI()
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

func (s *CoinExScraper) mainLoop() {

	defer func() {
		log.Println("CoinExScraper main loop exiting")
		s.cleanup()
	}()

	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	for {
		_, message, err := s.wsClient.ReadMessage()
		if err != nil {
			log.Error("Receive error:", err)
			continue
		}

		// Decompress GZIP data
		reader, err := gzip.NewReader(bytes.NewReader(message))
		if err != nil {
			log.Error("Decompression failed:", err)
			continue
		}
		defer reader.Close()

		// Read decompressed content
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(reader); err != nil {
			log.Error("Read failed:", err)
			continue
		}

		var response coinexWSResponse
		if err := json.Unmarshal(buf.Bytes(), &response); err != nil {
			log.Errorf("JSON parsing failed:%v | raw data:%s", err, buf.String())
			continue
		}

		if len(response.Data.DealList) == 0 {
			log.Warnf("Empty trade list | raw data:%s", buf.String())
			continue
		} else {
			s.parseWSResponse(response, tmFalseDuplicateTrades, tmDuplicateTrades)
		}
	}
}

func (s *CoinExScraper) parseWSResponse(
	message coinexWSResponse,
	tmFalseDuplicateTrades *timedmap.TimedMap,
	tmDuplicateTrades *timedmap.TimedMap,
) {
	if len(message.Data.DealList) == 0 {
		log.Warn("Empty Trade Message:", message)
		return
	}

	var exchangepair dia.ExchangePair
	var err error

	ps := s.newPairScrapers[message.Data.ForeignName]
	pair := ps.pair

	exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, message.Data.ForeignName)
	if err != nil {
		log.Error(err)
	}

	for _, deal := range message.Data.DealList {
		tradeTime := time.Unix(deal.CreatedAt/1000, (deal.CreatedAt%1000)*1e6)

		tradePrice, err := strconv.ParseFloat(deal.Price, 64)
		if err != nil {
			log.Errorf("Price parsing failed:%v | raw value:%s", err, deal.Price)
		}

		tradeVolume, err := strconv.ParseFloat(deal.Amount, 64)
		if err != nil {
			log.Errorf("Volume parsing failed:%v | raw value:%s", err, deal.Amount)
		}

		if strings.ToLower(deal.Side) == "sell" {
			tradeVolume = -math.Abs(tradeVolume)
		} else {
			tradeVolume = math.Abs(tradeVolume)
		}

		tradeForeignTradeID := strconv.FormatInt(deal.DealID, 10)

		t := &dia.Trade{
			Symbol:         pair.Symbol,
			Pair:           message.Data.ForeignName,
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
}

func (s *CoinExScraper) cleanup() {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	// close all channels of PairScraper children
	s.pairScrapers.Range(func(k, v interface{}) bool {
		for ps := range v.(coinExPairScraperSet) {
			ps.closed = true
		}
		s.pairScrapers.Delete(k)
		return true
	})

	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

func (scraper *CoinExScraper) connectToAPI() error {
	log.Info("Starting connect to API")

	dialer := ws.Dialer{
		EnableCompression: true, // Enable compression support
	}
	// Connect to CoinEx API.
	conn, _, err := dialer.Dial(CoinExWSBaseString, nil)
	if err != nil {
		log.Errorf("CoinEx - Connect to API: %s.", err.Error())
		return err
	}
	scraper.wsClient = conn

	return nil
}

func (scraper *CoinExScraper) subscribe(pair dia.ExchangePair, subscribe bool) error {
	if scraper.closed {
		return errors.New("CoinEx Scraper: Call ScrapePair on closed scraper")
	}

	// Validate WebSocket connection exists
	if scraper.wsClient == nil {
		return errors.New("WebSocket connection not initialized")
	}

	// Determine subscription type (SUBSCRIBE/UNSUBSCRIBE)
	subscribeType := "deals.unsubscribe"
	if subscribe {
		subscribeType = "deals.subscribe"
	}
	// Convert symbol+currency to uppercase (e.g., "btcusdt@trade")
	pairTicker := strings.ToUpper(pair.ForeignName)

	subscribeMessage := SubscribeRequest{
		Method: subscribeType,
		Params: MarketList{
			MarketList: []string{pairTicker},
		},
		ID: int(time.Now().UnixNano()),
	}
	log.Info("Subscribe Message: ", subscribeMessage)

	if err := scraper.wsClient.WriteJSON(subscribeMessage); err != nil {
		log.Errorf("Failed to send subscription request: %v", err)
		return err
	}
	return nil
}

func (s *CoinExScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	ps := &CoinExPairScraper{
		parent: s,
		pair:   pair,
	}

	s.newPairScrapers[pair.ForeignName] = ps

	s.subscribe(pair, true)

	return ps, nil
}

func (s *CoinExScraper) normalizeSymbol(p dia.ExchangePair, foreignName string, params ...string) (pair dia.ExchangePair, err error) {
	return pair, nil
}

func (s *CoinExScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return pairs, err
}

func (ps *CoinExPairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("CoinExPairScraper: Already closed")
	}

	ps.closed = true
	return err
}

func (ps *CoinExScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *CoinExPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (ps *CoinExPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}

func (s *CoinExScraper) Close() error {
	if s.closed {
		return errors.New("CoinExScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *CoinExScraper) FillSymbolData(symbol string) (dia.Asset, error) {
	return dia.Asset{Symbol: symbol}, nil
}

func (up *CoinExScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return pair, nil
}
