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
	utils "github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type ConfirmData struct {
	Result  []interface{} `json:"result"`
	Success bool          `json:"success"`
	Message string        `json:"message"`
}

var _bittrexapiurl string = "https://api.bittrex.com/api/v1.1/public"

type BittrexScraper struct {
	// signaling channels for session initialization and finishing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*BittrexPairScraper
	exchangeName string
	chanTrades   chan *dia.Trade
}

func NewBittrexScraper(exchangeName string) *BittrexScraper {
	s := &BittrexScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*BittrexPairScraper),
		exchangeName: exchangeName,
		error:        nil,
		chanTrades:   make(chan *dia.Trade),
	}

	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *BittrexScraper) mainLoop() {

	//wait for all pairscrapers have been created
	time.Sleep(7 * time.Second)
	layout := "2006-01-02T15:04:05"
	s.run = true
	for s.run {
		if len(s.pairScrapers) == 0 {
			s.error = errors.New("BittrexScraper: No pairs to scrap provided")
			log.Error(s.error.Error())
			break
		}
		for key, el := range s.pairScrapers {

			// more or less 60 calls per minute, the limit is 60
			time.Sleep(1 * time.Second)

			//swap the pairs name, necessary for api call
			sPairs := strings.Split(key, "-")
			sLeft, sRight := sPairs[0], sPairs[1]

			pairTrade := getAPICallBittrex("/getmarkethistory?market=" + sRight + "-" + sLeft)

			if len(pairTrade) > 0 {
				newId := 0
				atLeastOneUpdate := false
				for _, elTrade := range pairTrade {

					tradeReturn := elTrade.(map[string]interface{})
					idInt := int(tradeReturn["Id"].(float64))

					if el.lastIdTrade < idInt {
						if idInt > newId {
							newId = idInt
							atLeastOneUpdate = true
						}

						f64Price := tradeReturn["Price"].(float64)
						f64Volume := tradeReturn["Quantity"].(float64)

						if tradeReturn["OrderType"] == "SELL" {
							f64Volume = -f64Volume
						}

						timeStamp, _ := time.Parse(layout, tradeReturn["TimeStamp"].(string))
						t := &dia.Trade{
							Symbol:         el.pair.Symbol,
							Pair:           key,
							Price:          f64Price,
							Volume:         f64Volume,
							Time:           timeStamp,
							ForeignTradeID: strconv.FormatInt(int64(tradeReturn["Id"].(float64)), 16),
							Source:         s.exchangeName,
						}
						el.parent.chanTrades <- t
					}
				}
				if atLeastOneUpdate {
					el.lastIdTrade = newId
				}
			}
		}
	}
	if s.error == nil {
		s.error = errors.New("BittrexScraper: terminated by Close()")
	}
	s.cleanup(s.error)
}

func getAPICallBittrex(params ...string) []interface{} {

	body, err := utils.GetRequest(_bittrexapiurl + params[0])
	if err != nil {
		fmt.Println(err)
	}
	confirmTemp := ConfirmData{}
	jsonErr := json.Unmarshal(body, &confirmTemp)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return confirmTemp.Result
}

func (s *BittrexScraper) cleanup(err error) {
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
func (s *BittrexScraper) Close() error {

	if s.closed {
		return errors.New("BittrexScraper: Already closed")
	}
	s.run = false
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *BittrexScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("BittrexScraper: Call ScrapePair on closed scraper")
	}

	ps := &BittrexPairScraper{
		parent:      s,
		pair:        pair,
		lastIdTrade: 0,
	}

	s.pairScrapers[pair.ForeignName] = ps

	return ps, nil
}

func (s *BittrexScraper) normalizeSymbol(baseCurrency string, name string) (symbol string, err error) {
	symbol = strings.ToUpper(baseCurrency)
	if helpers.NameForSymbol(symbol) == symbol {
		if !helpers.SymbolIsName(symbol) {
			return symbol, errors.New("Foreign name can not be normalized:" + name + " symbol:" + symbol)
		}
	}
	if helpers.SymbolIsBlackListed(symbol) {
		return symbol, errors.New("Symbol is black listed:" + symbol)
	}
	return symbol, nil
}

// FetchAvailablePairs returns a list with all available trade pairs
func (s *BittrexScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {

	allPairs := getAPICallBittrex("/getmarkets")
	if len(allPairs) > 0 {
		for _, p := range allPairs {
			pairReturn := p.(map[string]interface{})
			symbol, serr := s.normalizeSymbol(pairReturn["MarketCurrency"].(string), pairReturn["MarketCurrencyLong"].(string))
			if serr == nil {
				pairs = append(pairs, dia.Pair{
					Symbol:      symbol,
					ForeignName: symbol + "-" + pairReturn["BaseCurrency"].(string),
					Exchange:    s.exchangeName,
				})
			} else {
				log.Error(serr)
			}
		}
	}
	return
}

// BittrexPairScraper implements PairScraper for Bittrex
type BittrexPairScraper struct {
	parent      *BittrexScraper
	pair        dia.Pair
	closed      bool
	lastIdTrade int
}

// Close stops listening for trades of the pair associated with s
func (ps *BittrexPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *BittrexScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *BittrexPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *BittrexPairScraper) Pair() dia.Pair {
	return ps.pair
}
