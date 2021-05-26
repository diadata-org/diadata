package main

import (
	"fmt"
	filters "github.com/diadata-org/diadata/internal/pkg/filtersOptionService"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"
	"math"
	"sort"
	"time"
)

// Get Near Term options
// Get NExt Term Options

var log = logrus.New()

type UnderlyingAsset struct {
	Symbol       string
	InterestRate float64
}

var underlyingAsset map[string]UnderlyingAsset

func init() {

	underlyingAsset = make(map[string]UnderlyingAsset)
	underlyingAsset["ETH"] = UnderlyingAsset{Symbol: "ETH", InterestRate: 0.01}
	underlyingAsset["BTC"] = UnderlyingAsset{Symbol: "BTC", InterestRate: 0.0054}

}

type OptionsTable struct {
	StrikePrice          float64
	PutAsk               float64
	PutBid               float64
	CallAsk              float64
	CallBid              float64
	CallMid              float64
	PutMid               float64
	Difference           float64
	ContributionByStrike float64
	Type                 dia.OptionType
}

func main() {
	asset := underlyingAsset["ETH"]

	// Get Data source
	ds, err := models.NewDataStore()
	if err != nil {
		log.Errorln("Error Getting Datastore", err)
	}
	for {
		CalculateVIX(asset, ds)
		time.Sleep(5 * time.Minute)
	}

}

func CalculateVIX(asset UnderlyingAsset, ds *models.DB) {

	var (
		nearTerm         map[float64]OptionsTable
		nextTerm         map[float64]OptionsTable
		nearTermCombined map[string]OptionsTable
		nextTermCombined map[string]OptionsTable
	)
	nearTerm = make(map[float64]OptionsTable)
	nextTerm = make(map[float64]OptionsTable)
	nearTermCombined = make(map[string]OptionsTable)
	nextTermCombined = make(map[string]OptionsTable)
	options, _ := filters.GetOptionComponents(asset.Symbol)

	// Calculate T1
	//  T = (Minutes Left in Current Day + Minutes in Settlement Day + Minutes in Other Days)/Minutes in Year
	t1 := filters.CalculateT(510.0, 34560.0)
	t2 := filters.CalculateT(900.0, 44640.0)

	nextTermOption, _ := filters.GetNext(options)
	nearTermOptions, _ := filters.GetNear(options)

	// Array for strike Prices, helps in sorting
	var nearstrikePrices []float64
	var nextstrikePrices []float64

	log.Infoln("Total Options selected", len(options))
	log.Infoln("Total Next Term Options", len(nextTermOption))
	log.Infoln("Total Near Term Options", len(nearTermOptions))

	// Create map of strike Price and call, bid, difference
	for _, option := range nextTermOption {
		ot := OptionsTable{}
		v, ok := nextTerm[option.StrikePrice]
		if ok {
			ot = v
		}

		isExist := false
		for _, strikePrice := range nextstrikePrices {
			if strikePrice == option.StrikePrice {
				isExist = true
			}
		}
		if !isExist {
			nextstrikePrices = append(nextstrikePrices, option.StrikePrice)
		}

		ot.StrikePrice = option.StrikePrice

		// Get orderbook for option
		orderbook, _ := ds.GetOptionOrderbookDataInflux(option)
		ot.Type = option.OptionType

		if option.OptionType == dia.CallOption {
			ot.CallAsk = orderbook.AskPrice
			ot.CallBid = orderbook.BidPrice
			strike := fmt.Sprintf("%f", option.StrikePrice)
			nextTermCombined[strike+"-C"] = ot

		}
		if option.OptionType == dia.PutOption {
			ot.PutAsk = orderbook.AskPrice
			ot.PutBid = orderbook.BidPrice
			strike := fmt.Sprintf("%f", option.StrikePrice)
			nextTermCombined[strike+"-P"] = ot
		}

		nextTerm[option.StrikePrice] = ot

	}

	for _, option := range nearTermOptions {
		ot := OptionsTable{}
		v, ok := nearTerm[option.StrikePrice]
		if ok {
			ot = v
		}

		isExist := false
		for _, strikePrice := range nearstrikePrices {
			if strikePrice == option.StrikePrice {
				isExist = true
			}
		}
		if !isExist {
			nearstrikePrices = append(nearstrikePrices, option.StrikePrice)
		}

		ot.StrikePrice = option.StrikePrice

		orderbook, _ := ds.GetOptionOrderbookDataInflux(option)
		ot.Type = option.OptionType

		if option.OptionType == dia.CallOption {
			ot.CallAsk = orderbook.AskPrice
			ot.CallBid = orderbook.BidPrice
			strike := fmt.Sprintf("%f", option.StrikePrice)
			nearTermCombined[strike+"-C"] = ot
		}
		if option.OptionType == dia.PutOption {
			ot.PutAsk = orderbook.AskPrice
			ot.PutBid = orderbook.BidPrice

			strike := fmt.Sprintf("%f", option.StrikePrice)
			nearTermCombined[strike+"-P"] = ot
		}
		nearTerm[option.StrikePrice] = ot

	}

	nearTerm = CalculateMidAndDifference(nearTerm)
	nextTerm = CalculateMidAndDifference(nextTerm)

	// Minimum Mid is the strike price used to Select  Option which is used for Forward Index calculation
	minimumDiffStrikePriceNearterm := FindMinimumMid(nearTerm)
	log.Infoln("Option Selected for calculating Forward Index for near", nearTerm[minimumDiffStrikePriceNearterm])
	minimumDiffStrikePriceNextterm := FindMinimumMid(nextTerm)
	log.Infoln("Option Selected for calculating Forward Index for next", nextTerm[minimumDiffStrikePriceNextterm])

	f1Data := nearTerm[minimumDiffStrikePriceNearterm]
	f2Data := nextTerm[minimumDiffStrikePriceNextterm]
	roi := asset.InterestRate

	f1 := filters.CalculateForwardIndex(f1Data.StrikePrice, roi, t1, f1Data.CallMid, f1Data.PutMid)
	f2 := filters.CalculateForwardIndex(f2Data.StrikePrice, roi, t2, f2Data.CallMid, f2Data.PutMid)

	log.Infoln("Forward Index F1", f1)
	log.Infoln("Forward Index F2", f2)

	// Get strike Price which is less or equal to the strike price of Forward Index

	k01 := 0.0
	for _, v := range nearstrikePrices {
		if v < f1 && v > k01 {
			k01 = v
		}
	}

	k02 := 0.0
	for _, v := range nextstrikePrices {
		if v < f2 && v > k02 {
			k02 = v
		}
	}
	log.Infoln("Strike Price k01", k01)
	log.Infoln("Strike Price k02", k02)

	sort.Float64s(nextstrikePrices)
	sort.Float64s(nearstrikePrices)

	log.Infoln("Near Strike Prices", nearstrikePrices)
	log.Infoln("Next Strike Prices", nextstrikePrices)

	//for _,strikePrice := range nearstrikePrices{
	//	log.Println(nearTerm[strikePrice])
	//}

	/*Select OTM options
	For Near Term
	1) For call option select options whose strike Price is Greater than k0
	2) For Put option select options whose strike Price is Less than k0

	For Near Term
	1) For call option select options whose strike Price is Greater than k1
	2) For Put option select options whose strike Price is Less than k1

	*/

	var (
		nearTermOTM map[float64]OptionsTable
		nextTermOTM map[float64]OptionsTable
	)
	nearTermOTM = make(map[float64]OptionsTable)
	nextTermOTM = make(map[float64]OptionsTable)
	//nextOTMOption := []float64{}

	// OTM for Near Term Call
	nearTermOTM = calculateOTMCall(nearstrikePrices, nearTermCombined, k01)

	nearTermOTMPut := calculateOTMPut(nearstrikePrices, nearTermCombined, k01)
	for k, v := range nearTermOTMPut {
		nearTermOTM[k] = v
	}

	// OTM for Next Term

	nextTermOTM = calculateOTMCall(nextstrikePrices, nextTermCombined, k02)

	nextTermOTMPut := calculateOTMPut(nextstrikePrices, nextTermCombined, k02)

	for k, v := range nextTermOTMPut {
		nextTermOTM[k] = v
	}

	maxDiff := 0.0
	selectedStrikePrice := 0.0
	var (
		nearTermOTMStrikePrices []float64
		nextTermOTMStrikePrices []float64
	)
	for k := range nearTermOTM {
		nearTermOTMStrikePrices = append(nearTermOTMStrikePrices, k)
		if nearTerm[k].Difference > maxDiff {
			maxDiff = nearTerm[k].Difference
			selectedStrikePrice = k
		}
	}

	log.Infoln("Selected Near", nearTermOTM[selectedStrikePrice])

	maxDiff = 0.0
	selectedStrikePrice = 0.0
	for k := range nextTermOTM {
		nextTermOTMStrikePrices = append(nextTermOTMStrikePrices, k)

		if nextTerm[k].Difference > maxDiff {
			maxDiff = nextTerm[k].Difference
			selectedStrikePrice = k
		}
	}

	sort.Float64s(nearTermOTMStrikePrices)
	log.Infoln("Total OTM Near Strike Price", len(nearTermOTMStrikePrices))
	sort.Float64s(nextTermOTMStrikePrices)
	log.Infoln("Total OTM Next Strike Price", len(nextTermOTMStrikePrices))

	//Calculate Sigma for ALL OTM

	//previousStrike := 0.0
	for _, strikePrice := range nearTermOTMStrikePrices {
		sigma := CalculateSigma(nearTermOTM[strikePrice], roi, t1)
		ot := nearTermOTM[strikePrice]
		ot.ContributionByStrike = sigma
		nearTermOTM[strikePrice] = ot
		//previousStrike = strikePrice
	}

	//previousStrike = 0.0
	for _, strikePrice := range nextTermOTMStrikePrices {
		sigma := CalculateSigma(nextTermOTM[strikePrice], roi, t2)
		ot := nextTermOTM[strikePrice]
		ot.ContributionByStrike = sigma
		nextTermOTM[strikePrice] = ot
		//previousStrike = strikePrice
	}

	// Sum of sigma

	nearSigma := 0.0
	nextSigma := 0.0

	for _, v := range nearTermOTM {
		nearSigma = nearSigma + v.ContributionByStrike
	}

	for _, v := range nextTermOTM {
		nextSigma = nextSigma + v.ContributionByStrike
	}

	//adjustment https://youtu.be/qToj8UiPBdk?t=613
	nearSigma = nearSigma * (2 / t1)
	nextSigma = nextSigma * (2 / t2)

	log.Infoln("nearSigma", nearSigma)
	log.Infoln("nextSigma", nextSigma)

	//Now calculate σ2 1 and σ2 2:

	cvinear := nearSigma - math.Pow((f1/k01)-1, 2)/t1

	cvinext := nextSigma - math.Pow((f2/k02)-1, 2)/t2

	log.Infoln("cvinear", cvinear)
	log.Infoln("cvinext", cvinext)

	vix := 100 * math.Sqrt((t1*cvinear*(46394-43200/46394-35942))+(t2*cvinext*(43200-35924/46394-35924))*525600/43200)

	log.Infoln("Saving CVI", vix)
	err := filters.ETHCVIToDatastore(vix)
	if err != nil {
		log.Error(err)
	}
}

// delta K/ k^2 * e^rt * q()
func CalculateSigma(ot OptionsTable, roi float64, t float64) float64 {
	return (5 / math.Pow(ot.StrikePrice, 2)) * math.Exp(roi*t) * math.Abs((ot.PutBid+ot.PutAsk)/2)

}

func FindMinimumMid(m map[float64]OptionsTable) (minimumStrikePrice float64) {
	var minimum float64
	for strikePrice, table := range m {
		if minimum < table.Difference {
			minimumStrikePrice = strikePrice
			minimum = table.Difference
		}

	}
	return
}

func CalculateMidAndDifference(m map[float64]OptionsTable) (calculated map[float64]OptionsTable) {

	var key float64
	for key = range m {
		//calculate call and put Mid
		v := m[key]
		v.CallMid = (v.CallAsk + v.CallBid) / 2
		v.PutMid = (v.PutAsk + v.PutBid) / 2
		v.Difference = math.Abs(v.CallMid - v.PutMid)
		m[key] = v
		if v.Difference == 0 {
			delete(m, key)
		}
	}
	calculated = m
	return

}

func unique(floatSlice []float64) []float64 {
	keys := make(map[float64]bool)
	list := []float64{}
	for _, entry := range floatSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func calculateOTMCall(strikePrices []float64, optionTableCombined map[string]OptionsTable, k float64) map[float64]OptionsTable {
	var otm (map[float64]OptionsTable)
	otm = make(map[float64]OptionsTable)
	lastStrikePrice := 0.0
	for _, strikePrice := range strikePrices {
		if lastStrikePrice != 0.0 {
			// OTM for Call Option
			strikePriceStr := fmt.Sprintf("%f", strikePrice)
			lastStrikePriceStr := fmt.Sprintf("%f", lastStrikePrice)

			optionTableCall := optionTableCombined[strikePriceStr+"-C"]
			lastoptionTableCall := optionTableCombined[lastStrikePriceStr+"-C"]

			if optionTableCall.Type == dia.CallOption {
				if lastoptionTableCall.CallBid == 0.0 {
					//skipping as last call price in 0.0
					log.Infoln("Skipping as last call Bid is 0.0")
					lastStrikePrice = strikePrice
					//Delete all
					for k := range otm {
						if otm[k].Type == dia.CallOption {
							delete(otm, k)
						}
					}
					continue
				}
				if optionTableCall.CallBid == 0.0 {
					continue
				}

				if strikePrice > k {
					otm[strikePrice] = optionTableCall
				}
			}

		}
		lastStrikePrice = strikePrice

	}
	return otm

}

func calculateOTMPut(strikePrices []float64, optionTableCombined map[string]OptionsTable, k float64) map[float64]OptionsTable {
	var otm (map[float64]OptionsTable)
	otm = make(map[float64]OptionsTable)
	lastStrikePrice := 0.0
	for _, strikePrice := range strikePrices {
		if lastStrikePrice != 0.0 {
			// OTM for Call Option
			strikePriceStr := fmt.Sprintf("%f", strikePrice)
			lastStrikePriceStr := fmt.Sprintf("%f", lastStrikePrice)

			optionTablePut := optionTableCombined[strikePriceStr+"-P"]
			lastoptionTablePut := optionTableCombined[lastStrikePriceStr+"-P"]

			if optionTablePut.Type == dia.PutOption {
				if lastoptionTablePut.PutBid == 0.0 {
					//skipping as last call price in 0.0
					log.Infoln("Skipping as last call Put is 0.0", lastoptionTablePut.PutBid)
					lastStrikePrice = strikePrice
					//Delete all
					for k := range otm {
						if otm[k].Type == dia.PutOption {
							delete(otm, k)
						}
					}
					continue
				}
				if optionTablePut.PutBid == 0.0 {
					continue
				}
				log.Infoln("strikePrice k", strikePrice, k)

				if strikePrice < k {

					otm[strikePrice] = optionTablePut
				}
			}

		}
		lastStrikePrice = strikePrice

	}
	return otm

}
