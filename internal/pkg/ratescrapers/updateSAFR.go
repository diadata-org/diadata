package ratescrapers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
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
	XMLdata, err := utils.GetRequest("https://apps.newyorkfed.org/rss/feeds/sofr-avg-ind")
	if err != nil {
		return err
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

	effDate, err := time.Parse("2006-01-02", rss.ItemInd.DateInd[:10])
	if err != nil {
		log.Error("Error parsing effective date for SAFR: ", err)
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
