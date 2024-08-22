package scrapers

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	ws "github.com/gorilla/websocket"
	gdax "github.com/preichenberger/go-coinbasepro/v2"
	"github.com/zekroTJA/timedmap"
)

type CoinBaseScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*CoinBasePairScraper // pc.ExchangePair -> pairScraperSet
	wsConn       *ws.Conn
	exchangeName string
	chanTrades   chan *dia.Trade
	db           *models.RelDB
}

const (
	ChannelHeartbeat = "heartbeat"
	ChannelTicker    = "ticker"
	ChannelLevel2    = "level2"
	ChannelUser      = "user"
	ChannelMatches   = "matches"
	ChannelFull      = "full"
)

// NewCoinBaseScraper returns a new CoinBaseScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewCoinBaseScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *CoinBaseScraper {
	s := &CoinBaseScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*CoinBasePairScraper),
		exchangeName: exchange.Name,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
		db:           relDB,
	}
	var wsDialer ws.Dialer
	SwConn, _, err := wsDialer.Dial(utils.Getenv("WEBSOCKET_API_URL", "wss://ws-feed.exchange.coinbase.com"), nil)
	if err != nil {
		println(err.Error())
	}
	s.wsConn = SwConn
	if scrape {
		go s.mainLoop()
	}
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *CoinBaseScraper) mainLoop() {
	var err error
	tmFalseDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)
	tmDuplicateTrades := timedmap.New(duplicateTradesScanFrequency)

	for {
		message := gdax.Message{}
		if err = s.wsConn.ReadJSON(&message); err != nil {
			println(err.Error())
			break
		}
		if message.Type == ChannelTicker {
			ps, ok := s.pairScrapers[message.ProductID]
			if ok {
				var f64Price float64
				var f64Volume float64
				var exchangepair dia.ExchangePair
				f64Price, err = strconv.ParseFloat(message.Price, 64)
				if err == nil {
					f64Volume, err = strconv.ParseFloat(message.LastSize, 64)
					if err == nil {
						if message.TradeID != 0 {
							if message.Side == "sell" {
								f64Volume = -f64Volume
							}

							exchangepair, err = s.db.GetExchangePairCache(s.exchangeName, message.ProductID)
							if err != nil {
								log.Error("get exchangepair from cache: ", err)
							}
							t := &dia.Trade{
								Symbol:         ps.pair.Symbol,
								Pair:           message.ProductID,
								Price:          f64Price,
								Volume:         f64Volume,
								Time:           message.Time.Time(),
								ForeignTradeID: strconv.FormatInt(int64(message.TradeID), 16),
								Source:         s.exchangeName,
								VerifiedPair:   exchangepair.Verified,
								BaseToken:      exchangepair.UnderlyingPair.BaseToken,
								QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
							}
							if t.VerifiedPair {
								log.Info("got verified trade: ", t)
							}
							// Handle duplicate trades.
							discardTrade := t.IdentifyDuplicateFull(tmFalseDuplicateTrades, duplicateTradesMemory)
							if !discardTrade {
								t.IdentifyDuplicateTagset(tmDuplicateTrades, duplicateTradesMemory)
								ps.parent.chanTrades <- t
							}

						}
					} else {
						log.Error("error parsing LastSize " + message.LastSize)
					}
				} else {
					log.Error("error parsing price " + message.Price)
				}
			} else {
				log.Error("unknown productError" + message.ProductID)
			}
		}
	}
	s.cleanup(err)
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *CoinBaseScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *CoinBaseScraper) Close() error {
	if s.closed {
		return errors.New("CoinBaseScraper: Already closed")
	}
	err := s.wsConn.Close()
	if err != nil {
		log.Error(err)
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (s *CoinBaseScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	// str := strings.Split(pair.ForeignName, "-")
	// symbol := str[0]
	// pair.Symbol = symbol
	// if helpers.NameForSymbol(symbol) == symbol {
	// 	return pair, errors.New("Foreign name can not be normalized:" + pair.ForeignName + " symbol:" + symbol)
	// }
	// if helpers.SymbolIsBlackListed(symbol) {
	// 	return pair, errors.New("Symbol is black listed:" + symbol)
	// }
	return pair, nil

}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *CoinBaseScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {

	data, _, err := utils.GetRequest("https://api.pro.coinbase.com/products")
	if err != nil {
		return
	}
	var ar []gdax.Product
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar {
			pairToNormalise := dia.ExchangePair{
				Symbol:      p.BaseCurrency,
				ForeignName: p.ID,
				Exchange:    s.exchangeName,
			}
			pair, serr := s.NormalizePair(pairToNormalise)
			if serr == nil {
				pairs = append(pairs, pair)
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// FillSymbolData collects all available information on an asset traded on CoinBase
func (s *CoinBaseScraper) FillSymbolData(symbol string) (asset dia.Asset, err error) {
	var response gdax.Currency
	data, _, err := utils.GetRequest("https://api.pro.coinbase.com/currencies/" + symbol)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return
	}
	asset.Symbol = response.ID
	asset.Name = response.Name
	return asset, nil
}

// CoinBasePairScraper implements PairScraper for GDax
type CoinBasePairScraper struct {
	parent     *CoinBaseScraper
	pair       dia.ExchangePair
	closed     bool
	lastRecord int64
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *CoinBaseScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("CoinBaseScraper: Call ScrapePair on closed scraper")
	}
	ps := &CoinBasePairScraper{
		parent:     s,
		pair:       pair,
		lastRecord: 0, //TODO FIX to figure out the last we got...
	}

	s.pairScrapers[pair.ForeignName] = ps

	subscribe := gdax.Message{
		Type: "subscribe",
		Channels: []gdax.MessageChannel{
			{
				Name: ChannelHeartbeat,
				ProductIds: []string{
					pair.ForeignName,
				},
			},
			{
				Name: ChannelTicker,
				ProductIds: []string{
					pair.ForeignName,
				},
			},
		},
	}
	if err := s.wsConn.WriteJSON(subscribe); err != nil {
		println(err.Error())
	}

	return ps, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *CoinBaseScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *CoinBasePairScraper) Close() error {
	ps.closed = true
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *CoinBasePairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *CoinBasePairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
