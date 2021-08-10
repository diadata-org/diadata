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
	CsafrRatesSecondaryFindByDateResponseInd struct {
		XMLName               xml.Name                 `xml:"safrRatesSecondaryFindByDateResponse,omitempty" json:"safrRatesSecondaryFindByDateResponse,omitempty"`
		CsafrRatesFindItemInd []*CsafrRatesFindItemInd `xml:"safrRatesFindItem,omitempty" json:"safrRatesFindItem,omitempty"`
	}

	CsafrRatesFindItemInd struct {
		XMLName           xml.Name           `xml:"safrRatesFindItem,omitempty" json:"safrRatesFindItem,omitempty"`
		CrateOperationInd *CrateOperationInd `xml:"rateOperation,omitempty" json:"rateOperation,omitempty"`
	}

	CrateOperationInd struct {
		XMLName             xml.Name             `xml:"rateOperation,omitempty" json:"rateOperation,omitempty"`
		CeffectiveDateInd   *CeffectiveDateInd   `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
		CinsertTimestampInd *CinsertTimestampInd `xml:"insertTimestamp,omitempty" json:"insertTimestamp,omitempty"`
		CrateIndexInd       *CrateIndexInd       `xml:"rateIndex,omitempty" json:"rateIndex,omitempty"`
		CrateTypeInd        *CrateTypeInd        `xml:"rateType,omitempty" json:"rateType,omitempty"`
	}

	CeffectiveDateInd struct {
		XMLName     xml.Name `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
		CEffDateInd string   `xml:",chardata" json:",omitempty"`
	}

	CinsertTimestampInd struct {
		XMLName       xml.Name `xml:"insertTimestamp,omitempty" json:"insertTimestamp,omitempty"`
		CTimestampInd string   `xml:",chardata" json:",omitempty"`
	}

	CrateIndexInd struct {
		XMLName   xml.Name `xml:"rateIndex,omitempty" json:"rateIndex,omitempty"`
		CValueInd string   `xml:",chardata" json:",omitempty"`
	}

	CrateTypeInd struct {
		XMLName  xml.Name `xml:"rateType,omitempty" json:"rateType,omitempty"`
		CTypeInd string   `xml:",chardata" json:",omitempty"`
	}
)

// WriteHistoricSAFR makes a GET request to fetch the historic data of the SOFR
// average index and writes it into the redis database.
func WriteHistoricSAFR(ds models.Datastore) error {
	log.Printf("Writing historic SAFR values")

	// Get rss from fed webpage
	XMLdata, _, err := utils.GetRequest("https://apps.newyorkfed.org/api/safrate/r1")

	if err != nil {
		return err
	}

	// Decode the body
	rss := new(CsafrRatesSecondaryFindByDateResponseInd)
	buffer := bytes.NewBuffer(XMLdata)
	decoded := xml.NewDecoder(buffer)
	err = decoded.Decode(rss)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// A slice containing all historic data
	histDataSlice := rss.CsafrRatesFindItemInd
	numData := len(histDataSlice)

	for i := 0; i < numData; i++ {
		var rate float64
		var dateTime time.Time
		var effDate time.Time

		// Convert interest rate from string to float64
		rate, err = strconv.ParseFloat(histDataSlice[i].CrateOperationInd.CrateIndexInd.CValueInd, 64)
		if err != nil {
			fmt.Println(err)
		}

		// Convert time string to Time type in UTC and pass date (without daytime)
		dateTime, err = time.Parse(time.RFC3339, histDataSlice[i].CrateOperationInd.CinsertTimestampInd.CTimestampInd)
		if err != nil {
			log.Error("error parsing publishing time of historic SAFR: ", err)
		} else {
			dateTime = dateTime.Round(time.Second).UTC()
		}

		effDate, err = time.Parse("2006-01-02", histDataSlice[i].CrateOperationInd.CeffectiveDateInd.CEffDateInd)
		if err != nil {
			log.Error("error parsing effective date of historic SAFR: ", err)
		}

		t := models.InterestRate{
			Symbol:          "SAFR",
			Value:           rate,
			PublicationTime: dateTime,
			EffectiveDate:   effDate,
			Source:          "FED",
		}

		err = ds.SetInterestRate(&t)
		if err != nil {
			log.Error(err)
		}

	}

	log.Info("Writing historic SAFR data complete.")

	return err
}
