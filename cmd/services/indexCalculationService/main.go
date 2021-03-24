package main

import (
	"time"

	"github.com/diadata-org/diadata/internal/pkg/indexCalculationService"
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
	firstRun := true
	go func() {
		for {
			select {
			case <-indexTicker.C:
				for _, index := range indexAssets {
					var currentConstituents []models.CryptoIndexConstituent
					if index.Symbol == "GBI" && firstRun {
						firstRun = false
						// symbols := []string{"WBTC", "ETH", "YFI", "UNI", "COMP", "MKR", "LINK", "SPICE"}
						addresses := []string{
							"0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
							"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
							"0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e",
							"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
							"0xc00e94Cb662C3520282E6f5717214004A7f26888",
							"0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2",
							"0x514910771AF9Ca656af840dff83E8264EcF986CA",
							"0x1fDaB294EDA5112B7d066ED8F2E4E562D5bCc664",
						}
						// Get constituents information
						currentConstituents, err = indexCalculationService.GetIndexBasket(addresses)
						if err != nil {
							log.Error(err)
						}

						// Calculate relative weights
						err = indexCalculationService.CalculateWeights(index.Symbol, &currentConstituents)
						if err != nil {
							log.Error(err)
						}
						for i, constituent := range currentConstituents {
							currentConstituents[i].NumBaseTokens = ((constituent.Weight * 100) / constituent.Price) * 1e16 //((Weight * IndexPrice) / TokenPrice) * 1e18  (divided by 100 because index level is 100 = 1 usd)
						}
						log.Info(currentConstituents)
					} else {
						currentConstituents = getCurrentIndexCompositionForIndex(index, ds)
					}
					log.Info(currentConstituents)
					index := periodicIndexValueCalculation(currentConstituents, index, ds)
					err := ds.SetCryptoIndex(&index)
					if err != nil {
						log.Error(err)
					}
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
