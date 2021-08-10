package ratescrapers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type (

	// Define the fields associated with the rss document.
	RssSOFR struct {
		Channel Channel `xml:"channel"`
		Item    Item    `xml:"item"`
	}
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
	}
	Item struct {
		Title       string     `xml:"title"`
		Link        string     `xml:"link"`
		Description string     `xml:"description"`
		Date        string     `xml:"date"`
		Statistics  Statistics `xml:"statistics"`
	}
	Statistics struct {
		Country    string `xml:"country"`
		InstAbbrev string `xml:"institutionAbbrev"`
		Rate       Rate   `xml:"interestRate"`
	}
	Rate struct {
		Value    string `xml:"value"`
		RateType string `xml:"rateType"`
	}
)

// UpdateSOFR makes a GET request from an rss feed and sends updated value through
// Channel s.chanInterestRate
func (s *RateScraper) UpdateSOFR() error {
	log.Printf("SOFRScraper update")

	XMLdata, _, err := utils.GetRequest("https://apps.newyorkfed.org/rss/feeds/sofr")

	if err != nil {
		return err
	}

	// Decode the body
	rss := new(RssSOFR)
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

	// Convert time string to Time type in UTC and pass date (without daytime)
	dateTime, err := time.Parse(time.RFC3339, rss.Item.Date)
	if err != nil {
		fmt.Println(err)
	} else {
		dateTime = dateTime.UTC()
	}

	// Extract effective date from title
	titleString := rss.Item.Title
	dateRegexp, _ := regexp.Compile(`\d{4}-\d{2}-\d{2}`)
	foundDate := dateRegexp.FindString(titleString)
	effDate, err := time.Parse("2006-01-02", foundDate)
	if err != nil {
		log.Errorf("Error parsing the effective date of %v", symbol)
	}

	t := &models.InterestRate{
		Symbol:          symbol,
		Value:           rate,
		PublicationTime: dateTime,
		EffectiveDate:   effDate,
		Source:          "FED",
	}

	// Send new data through channel chanInterestRate
	log.Printf("Write interestRate %#v in %v\n", t, s.chanInterestRate)
	s.chanInterestRate <- t

	log.Info("Update complete")

	return err
}
