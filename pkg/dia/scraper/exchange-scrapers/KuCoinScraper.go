package scrapers

import (
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/zekroTJA/timedmap"
)

type KuExchangePairs []KuExchangePair

type KucoinMarketMatch struct {
	Symbol       string `json:"symbol"`
	Sequence     string `json:"sequence"`
	Side         string `json:"side"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TakerOrderID string `json:"takerOrderId"`
	Time         string `json:"time"`
	Type         string `json:"type"`
	MakerOrderID string `json:"makerOrderId"`
	TradeID      string `json:"tradeId"`
}

type KucoinCurrency struct {
	Symbol  string `json:"currency"`
	Name    string `json:"fullName"`
	Address string `json:"contractAddress"`
}

type KuExchangePair struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	BaseCurrency    string `json:"baseCurrency"`
	QuoteCurrency   string `json:"quoteCurrency"`
	FeeCurrency     string `json:"feeCurrency"`
	Market          string `json:"market"`
	BaseMinSize     string `json:"baseMinSize"`
	QuoteMinSize    string `json:"quoteMinSize"`
	BaseMaxSize     string `json:"baseMaxSize"`
	QuoteMaxSize    string `json:"quoteMaxSize"`
	BaseIncrement   string `json:"baseIncrement"`
	QuoteIncrement  string `json:"quoteIncrement"`
	PriceIncrement  string `json:"priceIncrement"`
	PriceLimitRate  string `json:"priceLimitRate"`
	IsMarginEnabled bool   `json:"isMarginEnabled"`
	EnableTrading   bool   `json:"enableTrading"`
}

type KuCoinScraper struct {
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
	pairScrapers map[string]*KuCoinPairScraper // dia.ExchangePair -> KuCoinPairScraper
	// pairSubscriptions sync.Map                      // dia.ExchangePair -> string (subscription ID)
	// pairLocks         sync.Map                      // dia.ExchangePair -> sync.Mutex
	exchangeName string
	chanTrades   chan *dia.Trade
	apiService   *kucoin.ApiService
	db           *models.RelDB
}

func NewKuCoinScraper(apiKey string, secretKey string, exchange dia.Exchange, scrape bool, relDB *models.RelDB) *KuCoinScraper {
	apiService := kucoin.NewApiService()

	s := &KuCoinScraper{
		initDone:     make(chan nothing),
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		exchangeName: exchange.Name,
		pairScrapers: make(map[string]*KuCoinPairScraper),
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		apiService:   apiService,
		db:           relDB,
	}

	// establish connection in the background
	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *KuCoinScraper) mainLoop() {
	var channelsForClient1, channelsForClient2, channelsForClient3 []*kucoin.WebSocketSubscribeMessage

	close(s.initDone)

	lastTradeMap := make(map[dia.Pair]time.Time)
	countMap := make(map[dia.Pair]int)

	rsp, err := s.apiService.WebSocketPublicToken()
	if err != nil {
		// Handle error
		log.Error("Error WebSocketPublicToken", err)
	}

	tk := &kucoin.WebSocketTokenModel{}
	if err = rsp.ReadData(tk); err != nil {
		log.Error("Error Reading data", err)
	}

	client1 := s.apiService.NewWebSocketClient(tk)
	client2 := s.apiService.NewWebSocketClient(tk)
	client3 := s.apiService.NewWebSocketClient(tk)

	client1DownStream, _, err := client1.Connect()
	if err != nil {
		log.Error("Error Reading data", err)
	}
	client2DownStream, _, err := client2.Connect()
	if err != nil {
		log.Error("Error Reading data", err)
	}
	client3DownStream, _, err := client3.Connect()
	if err != nil {
		log.Error("Error Reading data", err)
	}

	count := 0
	for pair := range s.pairScrapers {
		ch := kucoin.NewSubscribeMessage("/market/match:"+pair, false)
		if count >= 598 {
			channelsForClient3 = append(channelsForClient3, ch)
			count++
			continue
		}
		if count >= 299 {
			channelsForClient2 = append(channelsForClient2, ch)
			count++
			continue
		} else {
			channelsForClient1 = append(channelsForClient1, ch)
			count++
		}
	}

	log.Info("number of pairs: ", count)

	if err := client1.Subscribe(channelsForClient1...); err != nil {
		log.Fatal("Error while subscribing client1 ", err)
	}
	if err := client2.Subscribe(channelsForClient2...); err != nil {
		log.Fatal("Error while subscribing client2 ", err)
	}
	if err := client3.Subscribe(channelsForClient3...); err != nil {
		log.Fatal("Error while subscribing client3 ", err)
	}

	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	go func() {
		var msg *kucoin.WebSocketDownstreamMessage
		for {
			select {
			case msg = <-client1DownStream:
				if msg == nil {
					continue
				}
				t := &KucoinMarketMatch{}
				if err := msg.ReadData(t); err != nil {
					log.Printf("Failure to read: %s", err.Error())
					return
				}
				asset := strings.Split(t.Symbol, "-")
				f64Price, _ := strconv.ParseFloat(t.Price, 64)
				f64Volume, _ := strconv.ParseFloat(t.Size, 64)
				timeOrder, err := strconv.ParseInt(t.Time, 10, 64)
				if err != nil {
					log.Error("parse trade time: ", err)
				}
				// WS returns different lengths of Unix timestamps. Adjust to nanoseconds if returns milliseconds.
				if len(t.Time) == 13 {
					timeOrder *= 1e6
				}

				if t.Side == "sell" {
					f64Volume = -f64Volume
				}

				exchangepair, err := s.db.GetExchangePairCache(s.exchangeName, t.Symbol)
				if err != nil {
					log.Error(err)
				}

				// Make trade times unique
				tradeTime := time.Unix(0, timeOrder)
				pair := dia.Pair{QuoteToken: exchangepair.UnderlyingPair.QuoteToken, BaseToken: exchangepair.UnderlyingPair.BaseToken}
				if _, ok := lastTradeMap[pair]; ok {
					if lastTradeMap[pair] != tradeTime {
						lastTradeMap[pair] = tradeTime
						countMap[pair] = 0
					} else {
						tradeTime = tradeTime.Add(time.Duration((countMap[pair] + 1)) * time.Nanosecond)
						countMap[pair] += 1
					}
				} else {
					lastTradeMap[pair] = tradeTime
				}

				trade := &dia.Trade{
					Symbol:         asset[0],
					Pair:           t.Symbol,
					Price:          f64Price,
					Time:           tradeTime,
					Volume:         f64Volume,
					Source:         s.exchangeName,
					VerifiedPair:   exchangepair.Verified,
					BaseToken:      exchangepair.UnderlyingPair.BaseToken,
					QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
					ForeignTradeID: t.TradeID,
				}
				if exchangepair.Verified {
					log.Info("Got verified trade from stream 1: ", trade)
				}

				// Handle duplicate trades.
				discardTrade := trade.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
				if !discardTrade {
					trade.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
					s.chanTrades <- trade
				}

			case msg = <-client2DownStream:
				if msg == nil {
					continue
				}
				t := &KucoinMarketMatch{}
				if err := msg.ReadData(t); err != nil {
					log.Errorf("Failure to read: %v", err)
					return
				}
				asset := strings.Split(t.Symbol, "-")
				f64Price, _ := strconv.ParseFloat(t.Price, 64)
				f64Volume, _ := strconv.ParseFloat(t.Size, 64)
				timeOrder, err := strconv.ParseInt(t.Time, 10, 64)
				if err != nil {
					log.Error("parse trade time: ", err)
				}
				// WS returns different lengths of Unix timestamps. Adjust to nanoseconds if returns milliseconds.
				if len(t.Time) == 13 {
					timeOrder *= 1e6
				}

				if t.Side == "sell" {
					f64Volume = -f64Volume
				}

				exchangepair, err := s.db.GetExchangePairCache(s.exchangeName, t.Symbol)
				if err != nil {
					log.Error(err)
				}

				// Make trade times unique
				tradeTime := time.Unix(0, timeOrder)
				pair := dia.Pair{QuoteToken: exchangepair.UnderlyingPair.QuoteToken, BaseToken: exchangepair.UnderlyingPair.BaseToken}
				if _, ok := lastTradeMap[pair]; ok {
					if lastTradeMap[pair] != tradeTime {
						lastTradeMap[pair] = tradeTime
						countMap[pair] = 0
					} else {
						//nolint
						tradeTime.Add(time.Duration(countMap[pair]+1) * time.Nanosecond)

						countMap[pair] += 1
					}
				} else {
					lastTradeMap[pair] = tradeTime
				}

				trade := &dia.Trade{
					Symbol:         asset[0],
					Pair:           t.Symbol,
					Price:          f64Price,
					Time:           time.Unix(0, timeOrder),
					Volume:         f64Volume,
					Source:         s.exchangeName,
					VerifiedPair:   exchangepair.Verified,
					BaseToken:      exchangepair.UnderlyingPair.BaseToken,
					QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
					ForeignTradeID: t.TradeID,
				}
				if exchangepair.Verified {
					log.Info("Got verified trade from stream 2: ", trade)
				}
				s.chanTrades <- trade

			case msg = <-client3DownStream:
				if msg == nil {
					continue
				}
				t := &KucoinMarketMatch{}
				if err := msg.ReadData(t); err != nil {
					log.Errorf("Failure to read: %v", err)
					return
				}
				asset := strings.Split(t.Symbol, "-")
				f64Price, _ := strconv.ParseFloat(t.Price, 64)
				f64Volume, _ := strconv.ParseFloat(t.Size, 64)
				timeOrder, err := strconv.ParseInt(t.Time, 10, 64)
				if err != nil {
					log.Error("parse trade time: ", err)
				}
				// WS returns different lengths of Unix timestamps. Adjust to nanoseconds if returns milliseconds.
				if len(t.Time) == 13 {
					timeOrder *= 1e6
				}

				if t.Side == "sell" {
					f64Volume = -f64Volume
				}

				exchangepair, err := s.db.GetExchangePairCache(s.exchangeName, t.Symbol)
				if err != nil {
					log.Error(err)
				}

				// Make trade times unique
				tradeTime := time.Unix(0, timeOrder)
				pair := dia.Pair{QuoteToken: exchangepair.UnderlyingPair.QuoteToken, BaseToken: exchangepair.UnderlyingPair.BaseToken}
				if _, ok := lastTradeMap[pair]; ok {
					if lastTradeMap[pair] != tradeTime {
						lastTradeMap[pair] = tradeTime
						countMap[pair] = 0
					} else {
						//nolint
						tradeTime.Add(time.Duration(countMap[pair]+1) * time.Nanosecond)
						countMap[pair] += 1
					}
				} else {
					lastTradeMap[pair] = tradeTime
				}

				trade := &dia.Trade{
					Symbol:         asset[0],
					Pair:           t.Symbol,
					Price:          f64Price,
					Time:           time.Unix(0, timeOrder),
					Volume:         f64Volume,
					Source:         s.exchangeName,
					VerifiedPair:   exchangepair.Verified,
					BaseToken:      exchangepair.UnderlyingPair.BaseToken,
					QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
					ForeignTradeID: t.TradeID,
				}
				if exchangepair.Verified {
					log.Info("Got verified trade from stream 3: ", trade)
				}
				s.chanTrades <- trade

			case <-s.shutdown: // user requested shutdown
				log.Println("KuCoin shutting down")
				s.cleanup(nil)
				return
			}
		}
	}()
}

func (s *KuCoinScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	return dia.ExchangePair{}, nil
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *KuCoinScraper) cleanup(err error) {
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
func (s *KuCoinScraper) Close() error {
	if s.closed {
		return errors.New("KuCoinScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *KuCoinScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("KucoinScraper: Call ScrapePair on closed scraper")
	}

	ps := &KuCoinPairScraper{
		parent: s,
		pair:   pair,
	}
	s.pairScrapers[pair.ForeignName] = ps
	return ps, nil
}

// FetchAvailablePairs returns all traded pairs on kucoin.
func (s *KuCoinScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	response, err := s.apiService.Symbols("")
	if err != nil {
		log.Println("Error Getting  Symbols for KuCoin Exchange", err)
	}

	var kep KuExchangePairs
	err = response.ReadData(&kep)
	if err != nil {
		log.Println("Error Reading  Symbols for KuCoin Exchange", err)
	}
	for _, p := range kep {
		pairs = append(pairs, dia.ExchangePair{
			Symbol:      p.BaseCurrency,
			ForeignName: p.Symbol,
			Exchange:    s.exchangeName,
		})
	}
	return
}

// FillSymbolData adds the name to the asset underlying @symbol on kucoin.
func (s *KuCoinScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	// Comment Philipp:
	// Kucoin's notations for symbols differ too often from the ones used in the underlying contracts.

	// resp, err := s.apiService.Currency(symbol, "")
	// if err != nil {
	// 	log.Errorf("error fetching %s from kucoin api: %v", symbol, err)
	// }
	// var kc KucoinCurrency
	// err = resp.ReadData(&kc)
	// if err != nil {
	// 	log.Errorf("error reading data for %s: %v", symbol, err)
	// }
	asset.Symbol = symbol
	// asset.Name = kc.Name
	// asset.Address = kc.Address
	return
}

// KuCoinPairScraper implements PairScraper for kuCoin
type KuCoinPairScraper struct {
	parent *KuCoinScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *KuCoinPairScraper) Close() error {
	var err error
	s := ps.parent
	// if parent already errored, return early
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return s.error
	}
	if ps.closed {
		return errors.New("KuCoinPairScraper: Already closed")
	}

	// TODO stop collection for the pair

	ps.closed = true
	return err
}

// Channel returns a channel that can be used to receive trades
func (ps *KuCoinScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *KuCoinPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *KuCoinPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
