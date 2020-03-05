package scrapers

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type SOFRScraper struct {
	// signaling channels
	shutdown     chan nothing
	shutdownDone chan nothing

	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock        sync.RWMutex
	error            error
	closed           bool
	ticker           *time.Ticker
	datastore        models.Datastore
	chanInterestRate chan *dia.InterestRate
}

type Rss struct {
	Channel Channel `xml:"channel"`
	Item    Item    `xml:"item"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}
type Item struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Date        string     `xml:"date"`
	Statistics  Statistics `xml:"statistics"`
}

type Statistics struct {
	Country    string `xml:"country"`
	InstAbbrev string `xml:"institutionAbbrev"`
	Rate       Rate   `xml:"interestRate"`
}
type Rate struct {
	Value    string `xml:"value"`
	RateType string `xml:"rateType"`
}

// NewSOFRScraper returns a new SOFRScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewSOFRScraper(datastore models.Datastore) *SOFRScraper {
	s := &SOFRScraper{
		shutdown:         make(chan nothing),
		shutdownDone:     make(chan nothing),
		error:            nil,
		ticker:           time.NewTicker(refreshDelay),
		datastore:        datastore,
		chanInterestRate: make(chan *dia.InterestRate),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *SOFRScraper) mainLoop() {
	for {
		select {
		case <-s.ticker.C:
			s.Update()
		case <-s.shutdown: // user requested shutdown
			log.Println("SOFRScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *SOFRScraper) cleanup(err error) {

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
func (s *SOFRScraper) Close() error {
	if s.closed {
		return errors.New("SOFRScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns a channel that can be used to receive trades/pricing information
func (s *SOFRScraper) Channel() chan *dia.InterestRate {
	return s.chanInterestRate
}

// Update makes a GET request from an rss feed and sends updated value through
// Channel s.chanInterestRate
func (s *SOFRScraper) Update() error {
	log.Printf("SOFRScraper update")

	// Get rss from fed webpage
	response, err := http.Get("https://apps.newyorkfed.org/rss/feeds/sofr")

	// Check, whether request successful
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Close response body after function
	defer response.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if response.StatusCode != 200 {
		return fmt.Errorf("HTTP Response Error %d", response.StatusCode)
	}

	// Read the response body
	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Decode the body
	rss := new(Rss)
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Collext entries of InterestRate struct -----------------------------------
	symbol := rss.Item.Statistics.Rate.RateType

	// Convert interest rate from string to float64
	rate, err := strconv.ParseFloat(rss.Item.Statistics.Rate.Value, 64)
	if err != nil {
		fmt.Println(err)
	}

	// Convert time string to Time type in UTC
	dateTime, err := time.Parse(time.RFC3339, rss.Item.Date)
	if err != nil {
		fmt.Println(err)
	} else {
		dateTime = dateTime.UTC()
	}

	t := &dia.InterestRate{
		Symbol: symbol,
		Value:  rate,
		Time:   dateTime,
		Source: "FED",
	}

	// Send new data through channel chanInterestRate
	log.Printf("writing interestRate %#v in %v\n", t, s.chanInterestRate)
	s.chanInterestRate <- t

	return err
}
