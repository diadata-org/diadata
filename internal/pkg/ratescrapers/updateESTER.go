package ratescrapers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type (

	// Define the fields associated with the rss document.
	RssESTER struct {
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

// UpdateESTER makes a GET request from an rss feed and sends updated value through
// Channel s.chanInterestRate
func (s *RateScraper) UpdateESTER() error {
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
	rss := new(RssESTER)
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
