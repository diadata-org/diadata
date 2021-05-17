package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	utils "github.com/diadata-org/diadata/pkg/utils"
)

type PairIdMap struct {
	Id          float64
	LastIdTrade int
	Symbol      string
}

type Confirm struct {
	Data interface{} `json:"data"`
}

var _apiurl string = "https://simex.global/api"

type SimexTicker struct {
	Data []SimexAsset `json:"data"`
}

type SimexAsset struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SimexScraper struct {
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
	pairScrapers           map[string]*SimexPairScraper
	pairIdTrade            map[string]*PairIdMap
	exchangeName           string
	chanTrades             chan *dia.Trade
	currencySymbolName     map[string]string
	isTickerMapInitialised bool
}

func NewSimexScraper(exchange dia.Exchange, scrape bool) *SimexScraper {
	s := &SimexScraper{
		shutdown:               make(chan nothing),
		shutdownDone:           make(chan nothing),
		pairScrapers:           make(map[string]*SimexPairScraper),
		exchangeName:           exchange.Name,
		error:                  nil,
		chanTrades:             make(chan *dia.Trade),
		currencySymbolName:     make(map[string]string),
		isTickerMapInitialised: false,
	}
	pairMap := map[string]*PairIdMap{}
	//API call used for retrievi all pairs
	//necessary to obtain the id used in next API calls
	data_temp := getAPICall("/pairs")
	//loop over each pair
	for _, el := range data_temp {
		md_element := el.(map[string]interface{})
		base_map := md_element["base"].(map[string]interface{})
		base := base_map["name"].(string)
		quote_map := md_element["quote"].(map[string]interface{})
		quote := quote_map["name"].(string)
		pim := &PairIdMap{
			Id:          md_element["id"].(float64),
			LastIdTrade: 0,
			Symbol:      base,
		}
		pairMap[base+quote] = pim
	}
	s.pairIdTrade = pairMap
	if scrape {
		go s.mainLoop()
	}
	return s
}

// runs in a goroutine until s is closed
func (s *SimexScraper) mainLoop() {
	relDB, err := models.NewRelDataStore()
	if err != nil {
		panic("Couldn't initialize relDB, error: " + err.Error())
	}

	//wait for all pairscrapers have been created
	time.Sleep(7 * time.Second)
	layout := "2006-01-02 15:04:05"
	s.run = true
	for s.run {
		if len(s.pairScrapers) == 0 {
			s.error = errors.New("SimexScraper: No pairs to scrap provided")
			log.Error(s.error.Error())
			break
		}
		for key, el := range s.pairScrapers {
			// more or less 60 calls per minute, the limit is 300
			time.Sleep(1 * time.Second)

			if s.pairIdTrade[key] == nil {
				log.Error(key, "s.pairIdTrade[key] == nil")
				continue
			}
			pairTrade := getAPICall("/trades/?pair_id=" + strconv.Itoa(int(s.pairIdTrade[key].Id)))
			if len(pairTrade) > 0 {
				newId := 0
				atLeastOneUpdate := false
				for _, elTrade := range pairTrade {

					tradeReturn := elTrade.(map[string]interface{})
					idInt := int(tradeReturn["id"].(float64))

					if s.pairIdTrade[key].LastIdTrade < idInt {
						if idInt > newId {
							newId = idInt
							atLeastOneUpdate = true
						}

						f64Price_string := tradeReturn["price"].(string)
						f64Price, err := strconv.ParseFloat(f64Price_string, 64)
						if err != nil {
							log.Errorln("ParseFloat:", err)
							continue
						}

						f64Volume_string := tradeReturn["size"].(string)
						f64Volume, err := strconv.ParseFloat(f64Volume_string, 64)
						if err != nil {
							log.Errorln("ParseFloat2:", err)
							continue
						}

						if tradeReturn["side"] == "sell" {
							f64Volume = -f64Volume
						}

						timeStamp, _ := time.Parse(layout, tradeReturn["created_at"].(string))
						exchangepair, err := relDB.GetExchangePairCache(s.exchangeName, tradeReturn["name"].(string))
						if err != nil {
							log.Error("Error Getting ExchangePair from cache", err)
						}

						t := &dia.Trade{
							Symbol:         s.pairIdTrade[key].Symbol,
							Pair:           key,
							Price:          f64Price,
							Volume:         f64Volume,
							Time:           timeStamp,
							ForeignTradeID: strconv.FormatInt(int64(tradeReturn["id"].(float64)), 16),
							Source:         s.exchangeName,
							VerifiedPair:   exchangepair.Verified,
							BaseToken:      exchangepair.UnderlyingPair.BaseToken,
							QuoteToken:     exchangepair.UnderlyingPair.QuoteToken,
						}
						if exchangepair.Verified {
							log.Infoln("Got verified trade", t)
						}

						el.parent.chanTrades <- t
					}
				}
				if atLeastOneUpdate {
					s.pairIdTrade[key].LastIdTrade = newId
				}
			}
		}
	}
	if s.error == nil {
		s.error = errors.New("SimexScraper: terminated by Close()")
	}
	s.cleanup(s.error)
}

func getAPICall(params ...string) []interface{} {

	body, err := utils.GetRequest(_apiurl + params[0])
	if err != nil {
		fmt.Println(err)
	}

	confirmTemp := Confirm{}
	jsonErr := json.Unmarshal(body, &confirmTemp)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	dataTemp := confirmTemp.Data.([]interface{})
	return dataTemp
}

func (s *SimexScraper) cleanup(err error) {
	s.errorLock.Lock()
	defer s.errorLock.Unlock()
	if err != nil {
		s.error = err
	}
	s.closed = true
	close(s.shutdownDone)
}

func (s *SimexScraper) FetchTickerData(symbol string) (asset dia.Asset, err error) {

	// Fetch Data
	if !s.isTickerMapInitialised {
		var (
			response SimexTicker
			data     []byte
		)
		data, err = utils.GetRequest("https://simex.global/api/currencies")
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &response)
		if err != nil {
			return
		}

		for _, asset := range response.Data {
			s.currencySymbolName[asset.Name] = asset.Description
		}
		s.isTickerMapInitialised = true

	}

	asset.Symbol = symbol
	asset.Name = s.currencySymbolName[symbol]
	return asset, nil
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *SimexScraper) Close() error {

	if s.closed {
		return errors.New("SimexScraper: Already closed")
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
func (s *SimexScraper) ScrapePair(pair dia.ExchangePair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("SimexScraper: Call ScrapePair on closed scraper")
	}

	ps := &SimexPairScraper{
		parent: s,
		pair:   pair,
	}

	s.pairScrapers[pair.ForeignName] = ps

	return ps, nil
}

func (s *SimexScraper) NormalizePair(pair dia.ExchangePair) (dia.ExchangePair, error) {
	symbol := strings.ToUpper(pair.Symbol)
	pair.Symbol = symbol
	pair.ForeignName = symbol + pair.ForeignName

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
func (s *SimexScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	type NameT struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	type DataT struct {
		Base  NameT `json:"base"`
		Quote NameT `json:"quote"`
	}
	type APIResponse struct {
		Data []DataT `json:"data"`
	}

	data, err := utils.GetRequest("https://simex.global/api/pairs")

	if err != nil {
		return
	}

	var ar APIResponse
	err = json.Unmarshal(data, &ar)
	if err == nil {
		for _, p := range ar.Data {

			pairToNormalize := dia.ExchangePair{
				Symbol:      p.Base.Name,
				ForeignName: p.Base.Name,
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

// SimexPairScraper implements PairScraper for Simex
type SimexPairScraper struct {
	parent *SimexScraper
	pair   dia.ExchangePair
	closed bool
}

// Close stops listening for trades of the pair associated with s
func (ps *SimexPairScraper) Close() error {
	ps.closed = true
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *SimexScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *SimexPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *SimexPairScraper) Pair() dia.ExchangePair {
	return ps.pair
}
