package main

import (
	"time"

	"github.com/diadata-org/diadata/dia-pkg/indexCalculationService"
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
				currentConstituents := getCurrentIndexCompositionForIndex(index, ds)
				log.Info(currentConstituents)
				index := periodicIndexValueCalculation(currentConstituents, index, ds)
				err := ds.SetCryptoIndex(&index)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()
	select {}
}

func getCurrentIndexCompositionForIndex(index dia.Asset, ds *models.DB) []models.CryptoIndexConstituent {
	var constituents []models.CryptoIndexConstituent
	cryptoIndex, err := ds.GetCryptoIndex(time.Now().Add(-5*time.Hour), time.Now(), index.Symbol)
	if err != nil {
		log.Error(err)
		return constituents
	}
	for _, constituent := range cryptoIndex[0].Constituents {
		curr, err := ds.GetCryptoIndexConstituents(time.Now().Add(-5*time.Hour), time.Now(), constituent.Asset, index.Symbol)
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

func periodicIndexValueCalculation(currentConstituents []models.CryptoIndexConstituent, indexAsset dia.Asset, ds *models.DB) models.CryptoIndex {
	err := indexCalculationService.UpdateConstituentsMarketData(indexAsset.Symbol, &currentConstituents)
	if err != nil {
		log.Error(err)
	}

	var price float64
	tradeObject, err := ds.GetIndexPrice(indexAsset, time.Now())
	if err == nil {
		// Quotation does exist
		price = tradeObject.EstimatedUSDPrice
	}
	var circSupply float64
	supplyObject, err := ds.GetSupplyInflux(indexAsset, time.Time{}, time.Time{})
	if err == nil && len(supplyObject) > 0 {
		// Supply does exist
		circSupply = supplyObject[0].CirculatingSupply
	}
	indexValue := indexCalculationService.GetIndexValue(indexAsset.Symbol, currentConstituents)
	currCryptoIndex, err := ds.GetCryptoIndex(time.Now().Add(-5*time.Hour), time.Now(), indexAsset.Symbol)
	if err != nil {
		log.Error(err)
	}
	index := models.CryptoIndex{
		Asset:             indexAsset,
		Price:             price,
		CirculatingSupply: circSupply,
		Value:             indexValue,
		CalculationTime:   time.Now(),
		Constituents:      currentConstituents,
		Divisor:           currCryptoIndex[0].Divisor,
	}
	log.Info("Index: ", index)
	return index
}
