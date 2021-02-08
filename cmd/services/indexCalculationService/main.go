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
	indexSymbol := "SCIFI"
	indexTicker := time.NewTicker(2 * 60 * time.Second)
	go func() {
		for {
			select {
			case <-indexTicker.C:
				currentConstituents := getCurrentIndexCompositionForIndex(indexSymbol, ds)
				index := periodicIndexValueCalculation(currentConstituents, indexSymbol, ds)
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
		curr, err := ds.GetCryptoIndexConstituents(time.Now().Add(-5 * time.Hour), time.Now(), constituentSymbol)
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

func getCurrentIndexCompositionForIndex(indexSymbol string, ds *models.DB) []models.CryptoIndexConstituent {
	var constituents []models.CryptoIndexConstituent
	cryptoIndex, err := ds.GetCryptoIndex(time.Now().Add(-5 * time.Hour), time.Now(),indexSymbol)
	if err != nil {
		log.Error(err)
		return constituents
	}
	for _, constituent := range cryptoIndex[0].Constituents {
		curr, err := ds.GetCryptoIndexConstituents(time.Now().Add(-5 * time.Hour), time.Now(), constituent.Symbol)
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

func periodicIndexValueCalculation(currentConstituents []models.CryptoIndexConstituent, indexSymbol string, ds *models.DB) models.CryptoIndex {
	err := indexCalculationService.UpdateConstituentsMarketData(&currentConstituents)
	if err != nil {
		log.Error(err)
	}
	quotation := 0.0
	tradeObject, err := ds.GetTradeInflux(indexSymbol, "", time.Now())
	if err == nil {
		// Quotation does exist
		quotation = tradeObject.EstimatedUSDPrice
	}
	supply := 0.0
	supplyObject, err := ds.GetLatestSupply(indexSymbol)
	if err == nil {
		// Supply does exist
		supply = supplyObject.CirculatingSupply
	}
	indexValue := indexCalculationService.GetIndexValue(currentConstituents)
	currCryptoIndex, err := ds.GetCryptoIndex(time.Now().Add(-5 * time.Hour), time.Now(), indexSymbol)
	if err != nil {
		log.Error(err)
	}
	index := models.CryptoIndex{
		Name:              indexSymbol,
		Price:             quotation,
		CirculatingSupply: supply,
		Value:             indexValue,
		CalculationTime:   time.Now(),
		Constituents:      currentConstituents,
		Divisor:           currCryptoIndex[0].Divisor,
	}
	log.Info("Index: ", index)
	return index
}
