package staticscrapers

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
	CsafrRatesSecondaryFindByDateResponseAvg struct {
		XMLName               xml.Name                 `xml:"safrRatesSecondaryFindByDateResponse,omitempty" json:"safrRatesSecondaryFindByDateResponse,omitempty"`
		CsafrRatesFindItemAvg []*CsafrRatesFindItemAvg `xml:"safrRatesFindItem,omitempty" json:"safrRatesFindItem,omitempty"`
	}

	CsafrRatesFindItemAvg struct {
		XMLName           xml.Name           `xml:"safrRatesFindItem,omitempty" json:"safrRatesFindItem,omitempty"`
		CrateOperationAvg *CrateOperationAvg `xml:"rateOperation,omitempty" json:"rateOperation,omitempty"`
	}

	CrateOperationAvg struct {
		XMLName             xml.Name             `xml:"rateOperation,omitempty" json:"rateOperation,omitempty"`
		CeffectiveDateAvg   *CeffectiveDateAvg   `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
		CinsertTimestampAvg *CinsertTimestampAvg `xml:"insertTimestamp,omitempty" json:"insertTimestamp,omitempty"`
		CrateIndexAvg       *CrateIndexAvg       `xml:"rateIndex,omitempty" json:"rateIndex,omitempty"`
		CrateTypeAvg        *CrateTypeAvg        `xml:"rateType,omitempty" json:"rateType,omitempty"`
	}

	CeffectiveDateAvg struct {
		XMLName     xml.Name `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
		CEffDateAvg string   `xml:",chardata" json:",omitempty"`
	}

	CinsertTimestampAvg struct {
		XMLName       xml.Name `xml:"insertTimestamp,omitempty" json:"insertTimestamp,omitempty"`
		CTimestampAvg string   `xml:",chardata" json:",omitempty"`
	}

	CrateIndexAvg struct {
		XMLName   xml.Name `xml:"rateIndex,omitempty" json:"rateIndex,omitempty"`
		CValueAvg string   `xml:",chardata" json:",omitempty"`
	}

	CrateTypeAvg struct {
		XMLName  xml.Name `xml:"rateType,omitempty" json:"rateType,omitempty"`
		CTypeAvg string   `xml:",chardata" json:",omitempty"`
	}
)

// WriteHistoricSOFRAvg makes a GET request to fetch the historic data of the SOFR
// average index and writes it into the redis database.
func WriteHistoricSOFRAvg(ds models.Datastore) error {
	log.Printf("Writing historic SOFR average values")

	// Get rss from fed webpage
	response, err := http.Get("https://apps.newyorkfed.org/api/safrate/r1")

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
	rss := new(CsafrRatesSecondaryFindByDateResponseAvg)
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// A slice containing all historic data
	histDataSlice := rss.CsafrRatesFindItemAvg
	numData := len(histDataSlice)

	for i := 0; i < numData; i++ {

		// Convert interest rate from string to float64
		rate, err := strconv.ParseFloat(histDataSlice[i].CrateOperationAvg.CrateIndexAvg.CValueAvg, 64)
		if err != nil {
			fmt.Println(err)
		}

		// Convert time string to Time type in UTC and pass date (without daytime)
		dateTime, err := time.Parse(time.RFC3339, histDataSlice[i].CrateOperationAvg.CinsertTimestampAvg.CTimestampAvg)

		if err != nil {
			fmt.Println(err)
		} else {
			dateTime = dateTime.Round(time.Second).UTC()
		}

		t := models.InterestRate{
			Symbol: "SOFR-AVG",
			Value:  rate,
			Time:   dateTime,
			Source: "FED",
		}

		ds.SetInterestRate(&t)

	}

	log.Info("Writing historic SOFR average data complete.")

	return err
}
