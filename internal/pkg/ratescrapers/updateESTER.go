package ratescrapers

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
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
		RefDate string `xml:"REF_DATE"`
		Rate    string `xml:"RATE"`
	}
)

func getRSS() (pubDate string, link string, err error) {
	// Scrapes the actual rss feed address for the ESTER index

	// First define the structure for decoding the xml
	type (
		// Define the fields associated with the rss document.
		Item struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		}
		Channel struct {
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			Item    []Item `xml:"item"`
		}
		RssMain struct {
			Channel Channel `xml:"channel"`
		}
	)

	XMLdata, err := utils.GetRequest("http://mid.ecb.europa.eu/rss/mid.xml")

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
			pubDate = rss.Channel.PubDate
			link = item.Link
			return
		}
	}
	return "", "", err
}

// UpdateESTER makes a GET request from an rss feed and sends updated value through
// Channel s.chanInterestRate
func (s *RateScraper) UpdateESTER() error {
	log.Printf("ESTERScraper update")

	// Get link to ESTER publication feed and publication time
	pubTime, address, err := getRSS()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get response from ESTER feed
	XMLdata, err := utils.GetRequest(address)
	if err != nil {
		return err
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

	// Convert publication time and effective date from string to Time type in UTC
	loc, _ := time.LoadLocation("CET") // Time is given as CET time
	publicationTime, err := time.ParseInLocation("2006/01/02 15:04:05", pubTime, loc)
	if err != nil {
		fmt.Println(err)
	} else {
		publicationTime = publicationTime.UTC()
	}
	effDate, err := time.Parse("2006-01-02", rss.RssBody.Content.EstrData.Results.RefDate)
	if err != nil {
		fmt.Println(err)
	}

	t := &models.InterestRate{
		Symbol:          "ESTER",
		Value:           rate,
		PublicationTime: publicationTime,
		EffectiveDate:   effDate,
		Source:          "ECB",
	}

	// Send new data through channel chanInterestRate
	log.Printf("writing interestRate %#v in %v\n", t, s.chanInterestRate)
	s.chanInterestRate <- t

	log.Info("Update complete")

	return err
}
