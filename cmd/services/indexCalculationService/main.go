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
	currentConstituents := periodicIndexRebalancingCalculation()
	indexTicker := time.NewTicker(2 * time.Minute)
	rebalancingTicker := time.NewTicker(30 * 24 * time.Hour)
	go func() {
		for {
			select {
			case <-rebalancingTicker.C:
				currentConstituents = periodicIndexRebalancingCalculation()
			case <-indexTicker.C:
				periodicIndexValueCalculation(currentConstituents)
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
	indexValue := indexCalculationService.GetIndexValue(currentConstituents)
	index := models.CryptoIndex{
		Value:        indexValue,
		Time:         time.Now(),
		Constituents: currentConstituents,
	}
	log.Info("Index: ", index)
	return index
}
