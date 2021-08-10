package staticscrapers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	utils "github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type (
	CMessageGroup struct {
		XMLName  xml.Name  `xml:"MessageGroup,omitempty" json:"MessageGroup,omitempty"`
		CDataSet *CDataSet `xml:"http://www.SDMX.org/resources/SDMXML/schemas/v2_0/message DataSet,omitempty" json:"DataSet,omitempty"`
	}

	CDataSet struct {
		XMLName    xml.Name   `xml:"DataSet,omitempty" json:"DataSet,omitempty"`
		Attraction string     `xml:"action,attr"  json:",omitempty"`
		CSeries    []*CSeries `xml:"http://www.SDMX.org/resources/SDMXML/schemas/v2_0/message Series,omitempty" json:"Series,omitempty"`
	}

	CSeries struct {
		XMLName xml.Name `xml:"Series,omitempty" json:"Series,omitempty"`
		CObs    []*CObs  `xml:"http://www.SDMX.org/resources/SDMXML/schemas/v2_0/message Obs,omitempty" json:"Obs,omitempty"`
	}

	CObs struct {
		XMLName         xml.Name `xml:"Obs,omitempty" json:"Obs,omitempty"`
		AttrOBS_VALUE   string   `xml:"OBS_VALUE,attr"  json:",omitempty"`
		AttrTIME_PERIOD string   `xml:"TIME_PERIOD,attr"  json:",omitempty"`
	}
)

const pathESTER = "../../internal/pkg/ratescraperData/ESTER_Historic_data.xml"

// GetHistoricESTER downloads historic ESTER data from the ECB Statistical Data
// Warehouse and stores it in an xml file. Here, historic means up until two days before now.
func GetHistoricESTER() error {
	linkESTER := "https://sdw.ecb.europa.eu/export.do?org.apache.struts.taglib.html.TOKEN=" +
		"8e79ff477e3157edfac21bf72dfb31b8&df=true&ec=&dc=&oc=&pb=&rc=&DATASET=0&r" +
		"emoveItem=&removedItemList=&mergeFilter=&activeTab=EST&showHide=&MAX_DOW" +
		"NLOAD_SERIES=500&SERIES_MAX_NUM=50&node=9698150&legendRef=reference&legendNor=&exportType=sdmx&ajaxTab=true"
	pathESTERAbs, _ := filepath.Abs(pathESTER)
	err := os.MkdirAll(filepath.Dir(pathESTERAbs), os.ModePerm)
	if err != nil {
		fmt.Println("here")
		return err
	}

	err = utils.DownloadResource(pathESTERAbs, linkESTER)
	if err != nil {
		fmt.Println("there")
		return err
	}
	return err
}

// WriteHistoricESTER makes a GET request to fetch the historic data of the SOFR index
// and writes it into the redis database.
func WriteHistoricESTER(ds models.Datastore) (err error) {

	log.Printf("Writing historic ESTER data")

	pathESTERAbs, _ := filepath.Abs(pathESTER)
	xmlFile, err := os.Open(pathESTERAbs) //nolint:gosec
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		cerr := xmlFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var myVar CMessageGroup
	err = xml.Unmarshal(byteValue, &myVar)
	if err != nil {
		return
	}

	// A slice containing all historic data. Value data is in the 8th row of series
	histDataSlice := myVar.CDataSet.CSeries[8].CObs
	numData := len(histDataSlice)

	for i := 0; i < numData; i++ {

		// Convert interest rate from string to float64
		rate, err := strconv.ParseFloat(histDataSlice[i].AttrOBS_VALUE, 64)
		if err != nil {
			fmt.Println(err)
		}

		// Convert time string to Time type in UTC and pass date (without daytime)
		// ESTR is published at around 08:00 CET. This timestamp lacks in historic tables.
		myTime := histDataSlice[i].AttrTIME_PERIOD
		layout := "2006-01-02"
		// loc, _ := time.LoadLocation("CET")
		// Parse time to UTC time
		// dateTime, err := time.ParseInLocation(layout, myTime, loc)
		dateTime, err := time.Parse(layout, myTime)

		if err != nil {
			fmt.Println(err)
		} else {
			dateTime = dateTime.UTC()
		}

		t := models.InterestRate{
			Symbol:          "ESTER",
			Value:           rate,
			PublicationTime: dateTime,
			EffectiveDate:   dateTime,
			Source:          "ECB",
		}

		err = ds.SetInterestRate(&t)
		if err != nil {
			log.Error(err)
		}

	}

	log.Info("Writing historic ESTER data complete.")

	return
}
