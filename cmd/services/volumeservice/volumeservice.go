package main

import (
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"
)

var (
	datastore *models.DB
	relDB     *models.RelDB
	err       error
)

const (
	lookbackDays = 2
)

func main() {

	datastore, err = models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	relDB, err = models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore:", err)
	}

	// initial run.
	fetchAndUpdateVolume()

	s := gocron.NewScheduler()
	err := s.Every(6).Hour().Do(fetchAndUpdateVolume)
	if err != nil {
		log.Error("schedule cron job: ", err)
	}
	<-s.Start()

}

func fetchAndUpdateVolume() {

	totalAssets, err := datastore.GetAssetsWithVOLInflux(time.Now().AddDate(0, 0, -lookbackDays))
	if err != nil {
		log.Errorln("Get assets with volume: ", err)
	}
	log.Infoln("Total Assets: ", totalAssets)

	for _, asset := range totalAssets {
		volume, err := datastore.Get24HoursAssetVolume(asset)
		if err != nil {
			log.Errorf("get volume of asset %s, %s", asset.Symbol, err.Error())
		} else {
			err = relDB.SetAssetVolume24H(asset, *volume, time.Now())
			if err != nil {
				log.Error("set asset volume: ", err)
			}
		}
	}

}
