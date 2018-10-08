package scrapers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"strconv"

	"github.com/diadata-org/api-golang/dia"
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

type SimexScraper struct {
	// signaling channels for session initialization and finishing
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*SimexPairScraper
	pairIdTrade  map[string]*PairIdMap
}

func NewSimexScraper() *SimexScraper {

	s := &SimexScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*SimexPairScraper),
		error:        nil,
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
	go s.mainLoop()
	return s
}

// runs in a goroutine until s is closed
func (s *SimexScraper) mainLoop() {

	//wait for all pairscrapers have been created
	time.Sleep(7 * time.Second)

	layout := "2006-01-02 15:04:05"

	for true {

		for key, el := range s.pairScrapers {

			// more or less 60 calls per minute, the limit is 300
			time.Sleep(1 * time.Second)
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
						f64Price, _ := strconv.ParseFloat(f64Price_string, 64)
						f64Volume_string := tradeReturn["size"].(string)
						f64Volume, _ := strconv.ParseFloat(f64Volume_string, 64)

						if tradeReturn["side"] == "sell" {
							f64Volume = -f64Volume
						}

						timeStamp, _ := time.Parse(layout, tradeReturn["created_at"].(string))
						t := &dia.Trade{
							Symbol:         s.pairIdTrade[key].Symbol,
							Pair:           key,
							Price:          f64Price,
							Volume:         f64Volume,
							Time:           timeStamp,
							ForeignTradeID: strconv.FormatInt(int64(tradeReturn["id"].(float64)), 16),
							Source:         dia.SimexExchange,
						}
						el.chanTrades <- t
					}
				}
				if atLeastOneUpdate {
					s.pairIdTrade[key].LastIdTrade = newId
				}
			}
		}
	}
}

func getAPICall(params ...string) []interface{} {

	req, err := http.Get(_apiurl + params[0])
	if err != nil {
		fmt.Println(err)
	}

	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	confirm_temp := Confirm{}
	jsonErr := json.Unmarshal(body, &confirm_temp)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	data_temp := confirm_temp.Data.([]interface{})
	return data_temp

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

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *SimexScraper) Close() error {

	if s.closed {
		return errors.New("SimexScraper: Already closed")
	}

	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *SimexScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()

	if s.error != nil {
		return nil, s.error
	}

	if s.closed {
		return nil, errors.New("SimexScraper: Call ScrapePair on closed scraper")
	}

	ps := &SimexPairScraper{
		parent:     s,
		pair:       pair,
		chanTrades: make(chan *dia.Trade),
	}

	s.pairScrapers[pair.ForeignName] = ps

	return ps, nil
}

// SimexPairScraper implements PairScraper for Simex
type SimexPairScraper struct {
	parent     *SimexScraper
	pair       dia.Pair
	chanTrades chan *dia.Trade
	closed     bool
}

// Close stops listening for trades of the pair associated with s
func (ps *SimexPairScraper) Close() error {
	return nil
}

// Channel returns a channel that can be used to receive trades
func (ps *SimexPairScraper) Channel() chan *dia.Trade {
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
func (ps *SimexPairScraper) Pair() dia.Pair {
	return ps.pair
}
