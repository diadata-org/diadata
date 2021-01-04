package main

import (
	"time"

	"github.com/diadata-org/diadata/internal/pkg/indexCalculationService"
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
	//currentConstituents := periodicIndexRebalancingCalculation()
	symbols := []string{"SUSHI", "REN", "KP3R", "UTK", "AXS", "Yf-DAI", "DIA", "STAKE", "POLS", "PICKLE", "EASY", "IDLE", "SPICE"}
	currentConstituents := getCurrentIndexComposition(symbols, ds)
	indexTicker := time.NewTicker(2 * 60 * time.Second)
	rebalancingTicker := time.NewTicker(30 * 24 * time.Hour)
	go func() {
		for {
			select {
			case <-rebalancingTicker.C:
				currentConstituents = periodicIndexRebalancingCalculation()
			case <-indexTicker.C:
				index := periodicIndexValueCalculation(currentConstituents, ds)
				err := ds.SetCryptoIndex(&index)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()
	select {}
}

func getCurrentIndexComposition(constituentsSymbols []string, ds *models.DB) []models.CryptoIndexConstituent {
	var constituents []models.CryptoIndexConstituent
	for _, constituentSymbol := range constituentsSymbols {
		curr, err := ds.GetCryptoIndexConstituents(time.Now().Add(-5*time.Hour), time.Now(), constituentSymbol)
		if err != nil {
			log.Error(err)
			return constituents
		}
		if len(curr) > 0 {
			constituents = append(constituents, curr[0])
		}
	}
	return constituents
}

func periodicIndexRebalancingCalculation() []models.CryptoIndexConstituent {
	symbols := []string{"SUSHI", "REN", "KP3R", "UTK", "AXS", "Yf-DAI", "DIA", "STAKE", "POLS", "PICKLE", "EASY", "IDLE", "SPICE"}

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

func periodicIndexValueCalculation(currentConstituents []models.CryptoIndexConstituent, ds *models.DB) models.CryptoIndex {
	symbol := "SCIFI"
	err := indexCalculationService.UpdateConstituentsMarketData(&currentConstituents)
	if err != nil {
		log.Error(err)
	}
	quotation := 0.0
	quotationObject, err := ds.GetQuotation(symbol)
	if err == nil {
		// Quotation does exist
		quotation = quotationObject.Price
	}
	supply := 0.0
	supplyObject, err := ds.GetLatestSupply(symbol)
	if err == nil {
		// Supply does exist
		supply = supplyObject.CirculatingSupply
	}
	indexValue := indexCalculationService.GetIndexValue(currentConstituents)
	index := models.CryptoIndex{
		Name:              symbol,
		Price:             quotation,
		CirculatingSupply: supply,
		Value:             indexValue,
		CalculationTime:   time.Now(),
		Constituents:      currentConstituents,
	}
	log.Info("Index: ", index)
	return index
}
