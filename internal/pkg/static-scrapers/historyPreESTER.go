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
	CMessageGroupPre struct {
		XMLName     xml.Name     `xml:"MessageGroup,omitempty" json:"MessageGroup,omitempty"`
		CDataSetPre *CDataSetPre `xml:"http://www.SDMX.org/resources/SDMXML/schemas/v2_0/message DataSet,omitempty" json:"DataSet,omitempty"`
	}

	CDataSetPre struct {
		XMLName    xml.Name      `xml:"DataSet,omitempty" json:"DataSet,omitempty"`
		Attraction string        `xml:"action,attr"  json:",omitempty"`
		CSeriesPre []*CSeriesPre `xml:"http://www.SDMX.org/resources/SDMXML/schemas/v2_0/message Series,omitempty" json:"Series,omitempty"`
	}

	CSeriesPre struct {
		XMLName xml.Name   `xml:"Series,omitempty" json:"Series,omitempty"`
		CTitle  string     `xml:"TITLE,attr"  json:",omitempty"`
		CObsPre []*CObsPre `xml:"http://www.SDMX.org/resources/SDMXML/schemas/v2_0/message Obs,omitempty" json:"Obs,omitempty"`
	}

	CObsPre struct {
		XMLName         xml.Name `xml:"Obs,omitempty" json:"Obs,omitempty"`
		AttrCONF_STATUS string   `xml:"CONF_STATUS,attr"  json:",omitempty"`
		AttrOBS_STATUS  string   `xml:"OBS_STATUS,attr"  json:",omitempty"`
		AttrOBS_VALUE   string   `xml:"OBS_VALUE,attr"  json:",omitempty"`
		AttrTIME_PERIOD string   `xml:"TIME_PERIOD,attr"  json:",omitempty"`
	}
)

const pathPreESTER = "../../internal/pkg/ratescraperData/PRE-ESTER_Historic_data.xml"

// GetHistoricPreESTER downloads historic ESTER data from the ECB Statistical Data
// Warehouse and stores it in an xml file. Here, historic means up until two days before now.
func GetHistoricPreESTER() error {
	linkPreESTER := "https://sdw.ecb.europa.eu/export.do?org.apache.struts.taglib.html.TOKEN" +
		"=f27da7a7aef6da13e1141c5c6932a48e&df=true&ec=&dc=&oc=&pb=&rc=&DATASET=0&removeItem" +
		"=&removedItemList=&mergeFilter=&activeTab=MMSR&showHide=&MAX_DOWNLOAD_SERIES=500&" +
		"SERIES_MAX_NUM=50&node=9693657&legendPub=published&legendNor=&exportType=sdmx&ajaxTab=true"
	pathPreESTERAbs, _ := filepath.Abs(pathPreESTER)

	err := utils.DownloadResource(pathPreESTERAbs, linkPreESTER)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// WriteHistoricPreESTER makes a GET request to fetch the historic data of the SOFR index
// and writes it into the redis database.
func WriteHistoricPreESTER(ds models.Datastore) error {
	log.Printf("Writing historic Pre-ESTER data")

	// The path relative to the calling main / executable
	absPath, _ := filepath.Abs(pathPreESTER)
	xmlFile, err := os.Open(absPath)

	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var myVar CMessageGroupPre
	xml.Unmarshal(byteValue, &myVar)

	// A slice containing all historic data. Value data is in the 8th row of series
	histDataSlice := myVar.CDataSetPre.CSeriesPre[7].CObsPre
	numData := len(histDataSlice)

	for i := 0; i < numData; i++ {

		// Convert interest rate from string to float64
		rate, err := strconv.ParseFloat(histDataSlice[i].AttrOBS_VALUE, 64)
		if err != nil {
			fmt.Println(err)
		}

		// Convert time string to Time type in UTC and pass date (without daytime)
		// Pre-ESTER is published at around 08:00 CET. This timestamp lacks in historic tables.
		myTime := histDataSlice[i].AttrTIME_PERIOD + "T08:00:00"
		layout := "2006-01-02T15:04:05"
		loc, _ := time.LoadLocation("CET")
		// Parse time to UTC time
		dateTime, err := time.ParseInLocation(layout, myTime, loc)

		if err != nil {
			fmt.Println(err)
		} else {
			dateTime = dateTime.UTC()
		}

		t := models.InterestRate{
			Symbol: "PRE-ESTER",
			Value:  rate,
			Time:   dateTime,
			Source: "ECB",
		}

		ds.SetInterestRate(&t)

	}

	log.Info("Writing historic Pre-ESTER data complete.")

	return err
}
