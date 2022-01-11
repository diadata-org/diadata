package main

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

const (
	influxDbOldTradesTable = "trades"
	influxDbTestTable      = "tradestest"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

type VerifiableAsset struct {
	Asset    dia.Asset
	Verified bool
}

func main() {

	influxURL := utils.Getenv("INFLUX_URL", "")
	log.Info("influxURL: ", influxURL)
	fromDB := utils.Getenv("INFLUX_DB_ORIGIN", "")
	toDB := utils.Getenv("INFLUX_DB_DESTINATION", "")
	fromTable := utils.Getenv("INFLUX_TABLE_ORIGIN", "")
	toTable := utils.Getenv("INFLUX_TABLE_DESTINATION", "")
	log.Info("fromDB: ", fromDB)
	log.Info("toDB: ", toDB)
	log.Info("fromTable: ", fromTable)
	log.Info("toTable: ", toTable)

	// The oldest record we have in influx is younger than 3000 days old.
	timeInitString := utils.Getenv("TIME_INIT", "1609455600")
	timeInitInt, err := strconv.ParseInt(timeInitString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	timeInit := time.Unix(timeInitInt, 0)
	log.Info("timeInit: ", timeInit)
	// timeFinalString is the last timestamp for which trades are read from influx in Unix time seconds.
	timeFinalString := utils.Getenv("TIME_FINAL", "1636618800")
	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	timeFinal := time.Unix(timeFinalInt, 0)
	log.Info("timeFinal: ", timeFinal)

	stepSizeString := utils.Getenv("STEP_SIZE_MINUTES", "720")
	stepSizeInt, err := strconv.ParseInt(stepSizeString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}
	ds.SetInfluxClient(influxURL)

	for timeFinal.After(timeInit) {
		endTime := timeInit.Add(time.Duration(stepSizeInt) * time.Minute)
		// err := ds.DeleteInfluxMeasurement(toDB, toTable, timeInit, endTime)
		// if err != nil {
		// 	log.Fatal("delete influx: ", err)
		// }
		// log.Infof("deleted from %v -- %v..", timeInit, endTime)
		numCopied, err := ds.CopyInfluxMeasurements(fromDB, toDB, fromTable, toTable, timeInit, endTime)
		if err != nil {
			log.Errorf("copying in time range %v -- %v: %v", timeInit, endTime, err)
		} else {
			log.Infof("copied %v rows in time range %v -- %v.", numCopied, timeInit, endTime)
		}
		timeInit = endTime
	}

}
