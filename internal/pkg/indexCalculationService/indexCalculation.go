package indexCalculationService

import (
	"math"
	"sort"
	"strings"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

var (MAX_RELATIVE_CAP float64 = 0.3)

// Get supply and price information for the index constituents
func GetIndexBasket(symbolsList []string) ([]models.CryptoIndexConstituent, error) {
	db, err := models.NewDataStore()
	if err != nil {
		log.Error("Error connecting to datastore")
		return nil, err
	}

	var constituents []models.CryptoIndexConstituent

	for _, symbol := range symbolsList {
		currQuotation, err := db.GetQuotation(strings.ToUpper(symbol))
		if err != nil {
			log.Error("Error when retrieveing quotation for ", symbol)
			return nil, err
		}
		currSupply, err := db.GetLatestSupply(symbol)
		if err != nil {
			log.Error("Error when retrieveing supply for ", symbol)
			return nil, err
		}
		currLastTrade, err := db.GetLastTradesAllExchanges(strings.ToUpper(symbol), 1)
		if err != nil {
			log.Error("Error when retrieveing lst trades for ", symbol)
			return nil, err
		}
		newConstituent := models.CryptoIndexConstituent{
			Address:            "-",
			Name:								currQuotation.Name,
			Symbol:							currSupply.Symbol,
			Price:							currLastTrade[0].EstimatedUSDPrice,
			CirculatingSupply:	currSupply.CirculatingSupply,
			Weight:             0.0,
			CappingFactor:      0.0,
			NumBaseTokens:      0.0,
		}
		constituents = append(constituents, newConstituent)
	}
	return constituents, nil
}

func CalculateWeights(indexSymbol string, constituents *[]models.CryptoIndexConstituent) error {
	if indexSymbol == "SCIFI" {
		type MarketCap struct {
			Symbol        string
			RawMarketCap  float64
			RelativeCap   float64
			CappingFactor float64
		}

		var marketCaps []MarketCap
		// Get full market cap
		sumMarketCap := 0.0
		for _, constituent := range *constituents {
			marketCap := constituent.CirculatingSupply * constituent.Price
			marketCaps = append(marketCaps, MarketCap{
				constituent.Symbol,
				marketCap,
				0.0,
				1.0,
			})
			sumMarketCap += marketCap
		}

		// Cut off market cap at MAX_RELATIVE_CAP
		// 1. Sort constituents by market cap
		sort.Slice(marketCaps, func(i, j int) bool {
			return marketCaps[i].RawMarketCap > marketCaps[j].RawMarketCap
		})

		// 2. Determine number of offendors (i.e. bigger relative market cap than MAX_RELATIVE_CAP)
		// and set their relative cap to MAX_RELATIVE_CAP
		numOffendors := 0
		offendor := marketCaps[numOffendors]
		uncappedConstituentsMc := 0.0

		for (offendor.RawMarketCap * math.Pow((1 - MAX_RELATIVE_CAP), float64(numOffendors)) > MAX_RELATIVE_CAP * sumMarketCap) {
			marketCaps[numOffendors].RelativeCap = MAX_RELATIVE_CAP
			sumMarketCap -= offendor.RawMarketCap
			numOffendors += 1
			if numOffendors < len(marketCaps) {
				offendor = marketCaps[numOffendors]
			} else {
				break
			}
		}

		// 3. Go through all non-offending constitutes and fix their relative cap
		for i, constituent := range marketCaps[numOffendors:] {
			marketCaps[i + numOffendors].RelativeCap = constituent.RawMarketCap / sumMarketCap * (1 - MAX_RELATIVE_CAP * float64(numOffendors))
			marketCaps[i + numOffendors].CappingFactor = 1.0
			uncappedConstituentsMc += constituent.RawMarketCap
		}
		// 4. Go through all offending constitutes and set a capping factor (i.e. factor to multiply their MC)
		for i, constituent := range marketCaps[:numOffendors] {
			if uncappedConstituentsMc != 0 {
				marketCaps[i].CappingFactor = MAX_RELATIVE_CAP / (constituent.RawMarketCap * (1 - MAX_RELATIVE_CAP * float64(numOffendors))) * uncappedConstituentsMc
			} else {
				marketCaps[i].CappingFactor = MAX_RELATIVE_CAP / (constituent.RawMarketCap * (1 - MAX_RELATIVE_CAP * float64(numOffendors)))
			}
		}

		// 5. Go through everything again and hardcode SPICE to 2.5%
		spiceIndex := len(*constituents)
		for i, constituent := range marketCaps {
			if constituent.Symbol == "SPICE" {
				spiceIndex = i
				break
			}
		}

		initialSpiceWeight := marketCaps[spiceIndex].RelativeCap
		correctionFactor := 0.025 / initialSpiceWeight
		correctionDelta := 0.025 - initialSpiceWeight
		if correctionDelta > 0 {
			for i, constituent := range marketCaps {
				if constituent.Symbol == "SPICE" {
					marketCaps[i].RelativeCap = 0.025
					marketCaps[i].CappingFactor = constituent.CappingFactor * correctionFactor
					continue
				}
				// Determine constitute's relative share to "give up"
				subtractionShare := correctionDelta * constituent.RelativeCap
				marketCaps[i].RelativeCap = constituent.RelativeCap - subtractionShare
				marketCaps[i].CappingFactor = constituent.CappingFactor * (1 - subtractionShare)
			}
		}

		// 6. Final step! Set data in the output struct
		for i, mc := range marketCaps {
			for j, constituent := range *constituents {
				if mc.Symbol == constituent.Symbol {
					(*constituents)[j].CappingFactor = marketCaps[i].CappingFactor
					(*constituents)[j].Weight = marketCaps[i].RelativeCap
				}
			}
		}

		return nil
	} else {
		//GBI index

		numConstituents := float64(len(*constituents))
		for i, constituent := range *constituents {
			if "SPICE" == constituent.Symbol {
				(*constituents)[i].Weight = 0.025
			} else {
				(*constituents)[i].Weight = (1-0.025) / (numConstituents - 1)
			}
		}
		return nil
	}
}

func UpdateConstituentsMarketData(indexSymbol string, currentConstituents *[]models.CryptoIndexConstituent) error {
	db, err := models.NewDataStore()
	if err != nil {
		log.Error("Error connecting to datastore")
		return err
	}
	for i, c := range *currentConstituents {
		currSupply, err := db.GetLatestSupply(c.Symbol)
		if err != nil {
			log.Error("Error when retrieveing supply for ", c.Symbol)
			return err
		}
		currLastTrade, err := db.GetLastTradesAllExchanges(strings.ToUpper(c.Symbol), 1)
		if err != nil {
			log.Error("Error when retrieveing last trades for ", c.Symbol)
			return err
		}
		(*currentConstituents)[i].Price = currLastTrade[0].EstimatedUSDPrice
		(*currentConstituents)[i].CirculatingSupply = currSupply.CirculatingSupply
	}

	// Calculate current percentages: 1. get index value 2. Determine percentage of each asset
	currIndexValue := GetIndexValue(indexSymbol, *currentConstituents)
	if indexSymbol == "SCIFI" {
		for i, _ := range *currentConstituents {
			currPercentage := ((*currentConstituents)[i].Price * (*currentConstituents)[i].CirculatingSupply * (*currentConstituents)[i].CappingFactor) / currIndexValue
			(*currentConstituents)[i].Percentage = currPercentage
		}
	} else {
		// GBI
		for i, _ := range *currentConstituents {
			currPercentage := ((*currentConstituents)[i].Price * (*currentConstituents)[i].NumBaseTokens * 1e-16) / currIndexValue //1e-16 because index value is 100 at start
			(*currentConstituents)[i].Percentage = currPercentage
		}
	}
	return nil
}

func GetIndexValue(indexSymbol string, currentConstituents []models.CryptoIndexConstituent) float64 {
	indexValue := 0.0
	if indexSymbol == "SCIFI" {
		for _, constituent := range currentConstituents {
			indexValue += constituent.Price * constituent.CirculatingSupply * constituent.CappingFactor
		}
	} else {
		// GBI etc
		for _, constituent := range currentConstituents {
			indexValue += constituent.Price * constituent.NumBaseTokens * 1e-16 //1e-16 because index value is 100 at start
		}
	}
	return indexValue
}
