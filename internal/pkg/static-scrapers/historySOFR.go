package staticscrapers

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
	CmktRatesSecondaryFindByDateResponse struct {
		XMLName           xml.Name             `xml:"mktRatesSecondaryFindByDateResponse,omitempty" json:"mktRatesSecondaryFindByDateResponse,omitempty"`
		CmktRatesFindItem []*CmktRatesFindItem `xml:"mktRatesFindItem,omitempty" json:"mktRatesFindItem,omitempty"`
	}

	CmktRatesFindItem struct {
		XMLName        xml.Name        `xml:"mktRatesFindItem,omitempty" json:"mktRatesFindItem,omitempty"`
		CrateOperation *CrateOperation `xml:"rateOperation,omitempty" json:"rateOperation,omitempty"`
	}

	CrateOperation struct {
		XMLName           xml.Name           `xml:"rateOperation,omitempty" json:"rateOperation,omitempty"`
		CeffectiveDate    *CeffectiveDate    `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
		CinsertTimestamp  *CinsertTimestamp  `xml:"insertTimestamp,omitempty" json:"insertTimestamp,omitempty"`
		Crate             *Crate             `xml:"rate,omitempty" json:"rate,omitempty"`
		CrateType         *CrateType         `xml:"rateType,omitempty" json:"rateType,omitempty"`
		CtradingVolume    *CtradingVolume    `xml:"tradingVolume,omitempty" json:"tradingVolume,omitempty"`
		CbusinessEventLog *CbusinessEventLog `xml:"businessEventLog,omitempty" json:"businessEventLog,omitempty"`
	}

	CbusinessEventLog struct {
		XMLName    xml.Name    `xml:"businessEventLog,omitempty" json:"businessEventLog,omitempty"`
		CeventDate *CeventDate `xml:"eventDate,omitempty" json:"eventDate,omitempty"`
	}

	CeventDate struct {
		XMLName xml.Name `xml:"eventDate,omitempty" json:"eventDate,omitempty"`
		CEvDate string   `xml:",chardata" json:",omitempty"`
	}

	CeffectiveDate struct {
		XMLName  xml.Name `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
		CEffDate string   `xml:",chardata" json:",omitempty"`
	}

	Crate struct {
		XMLName xml.Name `xml:"rate,omitempty" json:"rate,omitempty"`
		CValue  string   `xml:",chardata" json:",omitempty"`
	}

	CrateType struct {
		XMLName xml.Name `xml:"rateType,omitempty" json:"rateType,omitempty"`
		CType   string   `xml:",chardata" json:",omitempty"`
	}

	CtradingVolume struct {
		XMLName xml.Name `xml:"tradingVolume,omitempty" json:"tradingVolume,omitempty"`
		CVolume string   `xml:",chardata" json:",omitempty"`
	}

	CinsertTimestamp struct {
		XMLName    xml.Name `xml:"insertTimestamp,omitempty" json:"insertTimestamp,omitempty"`
		CTimestamp string   `xml:",chardata" json:",omitempty"`
	}
)

// WriteHistoricSOFR makes a GET request to fetch the historic data of the SOFR index
// and writes it into the redis database.
func WriteHistoricSOFR(ds models.Datastore) error {
	log.Printf("Writing historic SOFR data")

	// Get rss from fed webpage
	XMLdata, err := utils.GetRequest("https://apps.newyorkfed.org/api/mktrates/r3")

	if err != nil {
		return err
	}

	// Decode the body
	rss := new(CmktRatesSecondaryFindByDateResponse)
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// A slice containing all historic data
	histDataSlice := rss.CmktRatesFindItem
	numData := len(histDataSlice)

	for i := 0; i < numData; i++ {

		// Collect entries of InterestRate struct -----------------------------------
		symbol := histDataSlice[i].CrateOperation.CrateType.CType

		// Convert interest rate from string to float64
		rate, err := strconv.ParseFloat(histDataSlice[i].CrateOperation.Crate.CValue, 64)
		if err != nil {
			fmt.Println(err)
		}

		// Convert time string to Time type in UTC and pass date (without daytime)
		dateTime, err := time.Parse(time.RFC3339, histDataSlice[i].CrateOperation.CbusinessEventLog.CeventDate.CEvDate)

		if err != nil {
			fmt.Println(err)
		} else {
			dateTime = dateTime.Round(time.Second).UTC()
		}

		t := models.InterestRate{
			Symbol: symbol,
			Value:  rate,
			Time:   dateTime,
			Source: "FED",
		}

		ds.SetInterestRate(&t)

	}

	log.Info("Writing historic SOFR data complete.")

	return err
}
