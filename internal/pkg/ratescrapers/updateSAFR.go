package ratescrapers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

type (

	// Define the fields associated with the rss document.
	RssSAFR struct {
		ItemInd ItemInd `xml:"item"`
	}
	ItemInd struct {
		DescriptionInd string        `xml:"description"`
		DateInd        string        `xml:"date"`
		StatisticsInd  StatisticsInd `xml:"statistics"`
	}
	StatisticsInd struct {
		RateInd RateInd `xml:"interestRate"`
	}
	RateInd struct {
		ValueInd    string `xml:"value"`
		RateTypeInd string `xml:"rateType"`
	}
)

// UpdateSAFR makes a GET request from an rss feed and sends updated value through
// Channel s.chanInterestRate
func (s *RateScraper) UpdateSAFR() error {
	log.Printf("SAFRScraper update")

	// Get rss from fed webpage
	response, err := http.Get("https://apps.newyorkfed.org/rss/feeds/sofr-avg-ind")

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
	rss := new(RssSAFR)
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Collext entries of InterestRate struct -----------------------------------
	symbol := rss.ItemInd.StatisticsInd.RateInd.RateTypeInd

	// Convert interest rate from string to float64
	rate, err := strconv.ParseFloat(rss.ItemInd.StatisticsInd.RateInd.ValueInd, 64)
	if err != nil {
		fmt.Println(err)
	}

	// Convert time string to Time type in UTC and pass date (without daytime)
	dateTime, err := time.Parse(time.RFC3339, rss.ItemInd.DateInd)
	if err != nil {
		fmt.Println(err)
	} else {
		dateTime = dateTime.UTC()
	}

	t := &models.InterestRate{
		Symbol: symbol,
		Value:  rate,
		Time:   dateTime,
		Source: "FED",
	}

	// Send new data through channel chanInterestRate
	log.Printf("Write interestRate %#v in %v\n", t, s.chanInterestRate)
	s.chanInterestRate <- t

	log.Info("Update complete")

	return err
}
