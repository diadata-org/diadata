package main

import (
	"time"

	"github.com/diadata-org/diadata/internal/pkg/indexCalculationService"
	"github.com/sirupsen/logrus"
	models "github.com/diadata-org/diadata/pkg/model"
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
	currentConstituents := periodicIndexRebalancingCalculation()
	indexTicker := time.NewTicker(20 * time.Second)
	rebalancingTicker := time.NewTicker(30 * 24 * time.Hour)
	go func() {
		for {
			select {
			case <-rebalancingTicker.C:
				currentConstituents = periodicIndexRebalancingCalculation()
			case <-indexTicker.C:
				index := periodicIndexValueCalculation(currentConstituents)
				err := ds.SetCryptoIndex(&index)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()
	select {}
}

func periodicIndexRebalancingCalculation() ([]models.CryptoIndexConstituent) {
	symbols := []string{"SUSHI", "REN", "KP3R", "COVER", "UTK", "AXS", "RFI", "Yf-DAI", "DIA", "STAKE", "POLS", "PICKLE", "EASY", "IDLE", "SPICE"}

	// Get constituents information
	constituents, err := indexCalculationService.GetIndexBasket(symbols)
	if err != nil {
		log.Error(err)
	}

	// Calculate relative weights
	err = indexCalculationService.CalculateWeights(&constituents)
	if err != nil {
		log.Error(err)
	}
	log.Info(constituents)
	return constituents
}

func periodicIndexValueCalculation(currentConstituents []models.CryptoIndexConstituent) (models.CryptoIndex) {
	err := indexCalculationService.UpdateConstituentsMarketData(&currentConstituents)
	if err != nil {
		log.Error(err)
	}
	indexValue := indexCalculationService.GetIndexValue(currentConstituents)
	index := models.CryptoIndex{
		Name:         "SCIFI",
		Value:        indexValue,
		CalculationTime: time.Now(),
		Constituents: currentConstituents,
	}
	log.Info("Index: ", index)
	return index
}
