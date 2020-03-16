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
	"strings"
	"sync"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type ESTERScraper struct {
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
	chanInterestRate chan *models.InterestRate
}

type (

	// Define the fields associated with the rss document.
	RSS struct {
		Header  Header  `xml:"header"`
		RssBody RssBody `xml:"body"`
	}
	Header struct {
		ReleaseTitle string `xml:"releaseTitle"`
		ReleaseTime  string `xml:"releaseDateTime"`
	}
	RssBody struct {
		Content Content `xml:"content"`
	}
	Content struct {
		EstrData EstrData `xml:"EURO-SHORT-TERM-RATE_MID_PUBLICATION_MESSAGE"`
	}
	EstrData struct {
		Results Results `xml:"CALCULATION_RESULTS"`
	}
	Results struct {
		Rate string `xml:"RATE"`
	}
)

func getRSS() (string, error) {
	// Scrapes the actual rss feed address for the ESTER index

	// First define the structure for decoding the xml
	type (
		// Define the fields associated with the rss document.
		Item struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		}
		Channel struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
			Item  []Item `xml:"item"`
		}
		RssMain struct {
			Channel Channel `xml:"channel"`
		}
	)

	response, err := http.Get("http://mid.ecb.europa.eu/rss/mid.xml")

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
		return "", fmt.Errorf("HTTP Response Error %d", response.StatusCode)
	}

	// Read the response body
	XMLdata, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Decode the body
	rss := new(RssMain)
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Determine the item containing the link to the ESTER feed
	for _, item := range rss.Channel.Item {
		if strings.Contains(item.Title, "EURO-SHORT-TERM-RATE") {
			return item.Link, nil
		}
	}

	return "", err
}

// SpawnESTERScraper returns a new ESTERScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func SpawnESTERScraper(datastore models.Datastore) *ESTERScraper {
	s := &ESTERScraper{
		shutdown:         make(chan nothing),
		shutdownDone:     make(chan nothing),
		error:            nil,
		ticker:           time.NewTicker(refreshDelay),
		datastore:        datastore,
		chanInterestRate: make(chan *models.InterestRate),
	}

	log.Info("ESTER scraper is built and triggered")
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *ESTERScraper) mainLoop() {
	for {
		select {
		case <-s.ticker.C:
			s.Update()
		case <-s.shutdown: // user requested shutdown
			log.Println("ESTERScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// closes all connected PairScrapers
// must only be called from mainLoop
func (s *ESTERScraper) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	s.ticker.Stop()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (s *ESTERScraper) Close() error {
	if s.closed {
		return errors.New("ESTERScraper: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

// Channel returns a channel that can be used to receive rate information
func (s *ESTERScraper) Channel() chan *models.InterestRate {
	return s.chanInterestRate
}

// Update makes a GET request from an rss feed and sends updated value through
// Channel s.chanInterestRate
func (s *ESTERScraper) Update() error {
	log.Printf("ESTERScraper update")

	// Get link to ESTER publication feed
	address, err := getRSS()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get response from ESTER feed
	response, err := http.Get(address)

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
	rss := new(RSS)
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Collect entries of InterestRate struct -----------------------------------

	// Convert interest rate from string to float64
	rate, err := strconv.ParseFloat(rss.RssBody.Content.EstrData.Results.Rate, 64)
	if err != nil {
		fmt.Println(err)
	}

	// Convert time string to Time type in UTC
	layout := "2006-01-02T15:04:05"
	loc, _ := time.LoadLocation("CET") // Time is given as CET time
	dateTime, err := time.ParseInLocation(layout, rss.Header.ReleaseTime, loc)

	if err != nil {
		fmt.Println(err)
	} else {
		dateTime = dateTime.UTC()
	}

	t := &models.InterestRate{
		Symbol: "ESTER",
		Value:  rate,
		Time:   dateTime,
		Source: "ECB",
	}

	// Send new data through channel chanInterestRate
	log.Printf("writing interestRate %#v in %v\n", t, s.chanInterestRate)
	s.chanInterestRate <- t

	log.Info("Update complete")

	return err
}
