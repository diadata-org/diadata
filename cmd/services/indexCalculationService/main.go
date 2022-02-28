package main

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func main() {
	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}
	indexAssets := []dia.Asset{
		{
			Symbol:     "SCIFI",
			Name:       "ScifiToken",
			Address:    "0xfDC4a3FC36df16a78edCAf1B837d3ACAaeDB2CB4",
			Decimals:   18,
			Blockchain: dia.ETHEREUM,
		},
		{
			Symbol:     "GBI",
			Name:       "GalacticBlueIndex",
			Address:    "0xCB67bE5c54eab9462967eE3C03C35bfFfeB801cD",
			Decimals:   18,
			Blockchain: dia.ETHEREUM,
		},
	}
	indexTicker := time.NewTicker(300 * time.Second)
	go func() {
		for range indexTicker.C {
			for _, index := range indexAssets {

				// Get current constituents.
				log.Infof("get current index composition for %s...", index.Symbol)
				currentConstituents := ds.GetCurrentIndexCompositionForIndex(index)
				log.Info("current composition: ", currentConstituents)

				// Compute new index.
				err := ds.UpdateConstituentsMarketData(index.Symbol, &currentConstituents)
				if err != nil {
					log.Error(err)
				}
				indexValue := models.GetIndexValue(index.Symbol, currentConstituents)
				index := ds.IndexValueCalculation(currentConstituents, index, indexValue)

				// Save index.
				err = ds.SetCryptoIndex(&index)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()
	select {}
}
