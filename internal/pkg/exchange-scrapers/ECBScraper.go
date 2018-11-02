package scrapers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	refreshDelay = time.Second * 60 * 1
)

type ECBScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock    sync.RWMutex
	error        error
	closed       bool
	pairScrapers map[string]*ECBPairScraper // dia.Pair -> pairScraperSet
	ticker       *time.Ticker
	datastore    models.Datastore
}

type (
	// rssDocument defines the fields associated with the rss document.

	XMLCube struct {
		XMLName  xml.Name `xml:"Cube"`
		Currency string   `xml:"currency,attr"`
		Rate     string   `xml:"rate,attr"`
	}

	XMLCubeTime struct {
		XMLName xml.Name  `xml:"Cube"`
		Time    string    `xml:"time,attr"`
		Cube    []XMLCube `xml:"Cube"`
	}

	XMLEnvelope struct {
		XMLName  xml.Name      `xml:"Envelope"`
		CubeTime []XMLCubeTime `xml:"Cube>Cube"`
	}
)

// NewECBScraper returns a new ECBScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewECBScraper(datastore models.Datastore) *ECBScraper {
	s := &ECBScraper{
		shutdown:     make(chan nothing),
		shutdownDone: make(chan nothing),
		pairScrapers: make(map[string]*ECBPairScraper),
		error:        nil,
		ticker:       time.NewTicker(refreshDelay),
		datastore:    datastore,
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *ECBScraper) mainLoop() {
	for {
		select {
		case <-s.ticker.C:
			s.Update()
		case <-s.shutdown: // user requested shutdown
			log.Println("ECBScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *ECBScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections, as well as channels of
// PairScrapers from calls to ScrapePair
func (s *ECBScraper) Close() error {
	if s.closed {
		return errors.New("ECBScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// ECBPairScraper implements PairScraper for ECB
type ECBPairScraper struct {
	parent     *ECBScraper
	pair       dia.Pair
	chanTrades chan *dia.Trade
	closed     bool
}

// ScrapePair returns a PairScraper that can be used to get trades for a single pair from
// this APIScraper
func (s *ECBScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {

	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	if s.error != nil {
		return nil, s.error
	}
	if s.closed {
		return nil, errors.New("ECBScraper: Call ScrapePair on closed scraper")
	}
	ps := &ECBPairScraper{
		parent:     s,
		pair:       pair,
		chanTrades: make(chan *dia.Trade),
	}

	s.pairScrapers[pair.Symbol] = ps

	return ps, nil
}

// Channel returns a channel that can be used to receive trades/pricing information
func (ps *ECBPairScraper) Channel() chan *dia.Trade {
	return ps.chanTrades
}

func (ps *ECBPairScraper) Close() error {
	return nil
}

// Error returns an error when the channel Channel() is closed
// and nil otherwise
func (ps *ECBPairScraper) Error() error {
	s := ps.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Pair returns the pair this scraper is subscribed to
func (ps *ECBPairScraper) Pair() dia.Pair {
	return ps.pair
}

// retrieve performs a HTTP Get request for the rss feed and decodes the results.
func (s *ECBScraper) Update() error {
	log.Printf("ECBScraper update")

	// Retrieve the rss feed document from the web.
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return err
	}

	// Close the response once we return from the function.
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var document XMLEnvelope
	err = xml.NewDecoder(resp.Body).Decode(&document)

	if err != nil {
		fmt.Println(err)
	}

	for _, valueCubeTime := range document.CubeTime {
		change := &models.Change{
			[]models.CurrencyChange{},
		}

		euroDollar := 1.0
		for _, valueCube := range valueCubeTime.Cube {
			if valueCube.Currency == "USD" {
				euroDollar, err = strconv.ParseFloat(valueCube.Rate, 64)
				if err != nil {
					return fmt.Errorf("error parsing rate %v %v", valueCube.Rate, err)
				}
			}
		}

		for _, valueCube := range valueCubeTime.Cube {
			pair := string("EUR" + valueCube.Currency)
			ps := s.pairScrapers[pair]
			if ps != nil {
				rate, err := strconv.ParseFloat(valueCube.Rate, 64)
				if err != nil {
					return fmt.Errorf("error parsing rate %v %v", valueCube.Rate, err)
				}
				time, err := time.Parse("2006-01-02", valueCubeTime.Time)
				if err != nil {
					return fmt.Errorf("error parsing time %v %v", valueCubeTime.Time, err)
				}

				t := &dia.Trade{
					Pair:   pair,
					Symbol: pair,
					Price:  rate,
					Volume: 0,
					Time:   time,
					Source: "ECB",
				}
				log.Printf("writing trade %#v in %v\n", t, ps.chanTrades)
				ps.chanTrades <- t
				c := valueCube.Currency
				if c == "USD" {
					change.USD = append(change.USD, models.CurrencyChange{
						Symbol:        "EUR",
						Rate:          1.0 / euroDollar,
						RateYesterday: 1.0 / euroDollar, // TOFIX
					})
				} else {
					// list for coinhub
					if (c == "JPY") || c == "GBP" || c == "SEK" || c == "CHF" || c == "NOK" || c == "AUD" || c == "CAD" || c == "CNY" || c == "KRW" {
						change.USD = append(change.USD, models.CurrencyChange{
							Symbol:        c,
							Rate:          rate / euroDollar,
							RateYesterday: rate / euroDollar, // TOFIX
						})
					}
				}
			}
		}
		s.datastore.SetCurrencyChange(change)
	}
	return err
}
