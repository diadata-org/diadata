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
		CTenor1Avg          *CTenor1Avg          `xml:"tenor1,omitempty" json:"tenor1,omitempty"`
		CTenor2Avg          *CTenor2Avg          `xml:"tenor2,omitempty" json:"tenor2,omitempty"`
		CTenor3Avg          *CTenor3Avg          `xml:"tenor3,omitempty" json:"tenor3,omitempty"`
	}

	CeffectiveDateAvg struct {
		XMLName     xml.Name `xml:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
		CEffDateAvg string   `xml:",chardata" json:",omitempty"`
	}

	CinsertTimestampAvg struct {
		XMLName       xml.Name `xml:"insertTimestamp,omitempty" json:"insertTimestamp,omitempty"`
		CTimestampAvg string   `xml:",chardata" json:",omitempty"`
	}

	CTenor1Avg struct {
		XMLName    xml.Name `xml:"tenor1,omitempty" json:"tenor1,omitempty"`
		CValue1Avg string   `xml:",chardata" json:",omitempty"`
	}

	CTenor2Avg struct {
		XMLName    xml.Name `xml:"tenor2,omitempty" json:"tenor1,omitempty"`
		CValue2Avg string   `xml:",chardata" json:",omitempty"`
	}

	CTenor3Avg struct {
		XMLName    xml.Name `xml:"tenor3,omitempty" json:"tenor1,omitempty"`
		CValue3Avg string   `xml:",chardata" json:",omitempty"`
	}
)

// WriteHistoricSAFRAvgs makes a GET request to fetch the historic data of the SOFR
// average index and writes it into the redis database.
func WriteHistoricSAFRAvgs(ds models.Datastore) error {
	log.Printf("Writing historic SAFR average values")

	// Get rss from fed webpage
	XMLdata, err := utils.GetRequest("https://apps.newyorkfed.org/api/safrate/r1")

	if err != nil {
		return err
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

		// Convert rates from string to float64
		rate1, err := strconv.ParseFloat(histDataSlice[i].CrateOperationAvg.CTenor1Avg.CValue1Avg, 64)
		if err != nil {
			fmt.Println(err)
		}
		rate2, err := strconv.ParseFloat(histDataSlice[i].CrateOperationAvg.CTenor2Avg.CValue2Avg, 64)
		if err != nil {
			fmt.Println(err)
		}

		rate3, err := strconv.ParseFloat(histDataSlice[i].CrateOperationAvg.CTenor3Avg.CValue3Avg, 64)
		if err != nil {
			fmt.Println(err)
		}

		// Convert time string to Time type in UTC and pass date (without daytime)
		dateTime, err := time.Parse(time.RFC3339, histDataSlice[i].CrateOperationAvg.CinsertTimestampAvg.CTimestampAvg)
		if err != nil {
			log.Error("Error parsing publishing time for SOFRXXX: ", err)
		} else {
			dateTime = dateTime.Round(time.Second).UTC()
		}

		effDate, err := time.Parse("2006-01-02", histDataSlice[i].CrateOperationAvg.CeffectiveDateAvg.CEffDateAvg)
		if err != nil {
			log.Error("Error parsing effective date for SOFRXXX: ", err)
		}

		t1 := models.InterestRate{
			Symbol:          "SOFR30",
			Value:           rate1,
			PublicationTime: dateTime,
			EffectiveDate:   effDate,
			Source:          "FED",
		}

		t2 := models.InterestRate{
			Symbol:          "SOFR90",
			Value:           rate2,
			PublicationTime: dateTime,
			EffectiveDate:   effDate,
			Source:          "FED",
		}

		t3 := models.InterestRate{
			Symbol:          "SOFR180",
			Value:           rate3,
			PublicationTime: dateTime,
			EffectiveDate:   effDate,
			Source:          "FED",
		}

		ds.SetInterestRate(&t1)
		ds.SetInterestRate(&t2)
		ds.SetInterestRate(&t3)

	}

	log.Info("Writing historic SAFR average data complete.")

	return err
}
