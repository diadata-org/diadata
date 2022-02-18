package main

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

var (
	log        *logrus.Logger
	symbolsMap map[string]string
	indexMap   map[string]dia.Asset
)

func init() {
	log = logrus.New()
	symbolsMap = make(map[string]string)
	symbolsMap["1INCH"] = "0x111111111117dC0aa78b770fA6A738034120C302"
	symbolsMap["ALPHA"] = "0xa1faa113cbE53436Df28FF0aEe54275c13B40975"
	symbolsMap["AXS"] = "0xBB0E17EF65F82Ab018d8EDd776e8DD940327B28b"
	symbolsMap["BAL"] = "0xba100000625a3754423978a60c9317c58a424e3D"
	symbolsMap["COMP"] = "0xc00e94Cb662C3520282E6f5717214004A7f26888"
	symbolsMap["COVER"] = "0x5D8d9F5b96f4438195BE9b99eee6118Ed4304286"
	symbolsMap["CRV"] = "0xD533a949740bb3306d119CC777fa900bA034cd52"
	symbolsMap["DIA"] = "0x84cA8bc7997272c7CfB4D0Cd3D55cd942B3c9419"
	symbolsMap["EASY"] = "0x913D8ADf7CE6986a8CbFee5A54725D9Eea4F0729"
	symbolsMap["ETH"] = "0x0000000000000000000000000000000000000000"
	symbolsMap["FRM"] = "0xE5CAeF4Af8780E59Df925470b050Fb23C43CA68C"
	symbolsMap["FTM"] = "0x4E15361FD6b4BB609Fa63C81A2be19d873717870"
	symbolsMap["FTX Token"] = "0x50D1c9771902476076eCFc8B2A83Ad6b9355a4c9"
	symbolsMap["GEL"] = "0x15b7c0c907e4C6b9AdaAaabC300C08991D6CEA05"
	symbolsMap["GRT"] = "0xc944E90C64B2c07662A292be6244BDf05Cda44a7"
	symbolsMap["GTC"] = "0xDe30da39c46104798bB5aA3fe8B9e0e1F348163F"
	symbolsMap["HEGIC"] = "0x584bC13c7D411c00c01A62e8019472dE68768430"
	symbolsMap["HEX"] = "0x2b591e99afE9f32eAA6214f7B7629768c40Eeb39"
	symbolsMap["IDLE"] = "0x875773784Af8135eA0ef43b5a374AaD105c5D39e"
	symbolsMap["INJ"] = "0xe28b3B32B6c345A34Ff64674606124Dd5Aceca30"
	symbolsMap["KP3R"] = "0x1cEB5cB57C4D4E2b2433641b95Dd330A33185A44"
	symbolsMap["LINK"] = "0x514910771AF9Ca656af840dff83E8264EcF986CA"
	symbolsMap["MKR"] = "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2"
	symbolsMap["NMR"] = "0x1776e1F26f98b1A5dF9cD347953a26dd3Cb46671"
	symbolsMap["PERP"] = "0xbC396689893D065F41bc2C6EcbeE5e0085233447"
	symbolsMap["PICKLE"] = "0x429881672B9AE42b8EbA0E26cD9C73711b891Ca5"
	symbolsMap["POLS"] = "0x83e6f1E41cdd28eAcEB20Cb649155049Fac3D5Aa"
	symbolsMap["REN"] = "0x408e41876cCCDC0F92210600ef50372656052a38"
	symbolsMap["RFI"] = "0xA1AFFfE3F4D611d252010E3EAf6f4D77088b0cd7"
	symbolsMap["ROOK"] = "0xfA5047c9c78B8877af97BDcb85Db743fD7313d4a"
	symbolsMap["SNOW"] = "0xfe9A29aB92522D14Fc65880d817214261D8479AE"
	symbolsMap["SPICE"] = "0x1fDaB294EDA5112B7d066ED8F2E4E562D5bCc664"
	symbolsMap["STAKE"] = "0x0Ae055097C6d159879521C384F1D2123D1f195e6"
	symbolsMap["SUPER"] = "0xe53EC727dbDEB9E2d5456c3be40cFF031AB40A55"
	symbolsMap["SUSHI"] = "0x6B3595068778DD592e39A122f4f5a5cF09C90fE2"
	symbolsMap["UNI"] = "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"
	symbolsMap["UTK"] = "0xdc9Ac3C20D1ed0B540dF9b1feDC10039Df13F99c"
	symbolsMap["WBTC"] = "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599"
	symbolsMap["YF-DAI"] = "0xf4CD3d3Fda8d7Fd6C5a500203e38640A70Bf9577"
	symbolsMap["YFI"] = "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e"
	symbolsMap["Yf-DAI"] = "0xf4CD3d3Fda8d7Fd6C5a500203e38640A70Bf9577"
	symbolsMap["wNXM"] = "0x0d438F3b5175Bebc262bF23753C1E53d03432bDE"

	indexMap = make(map[string]dia.Asset)
	indexMap["SCIFI"] = dia.Asset{
		Address:    "0xfDC4a3FC36df16a78edCAf1B837d3ACAaeDB2CB4",
		Blockchain: dia.ETHEREUM,
		Symbol:     "SCIFI",
		Name:       "ScifiToken",
		Decimals:   uint8(18),
	}
	indexMap["GBI"] = dia.Asset{
		Address:    "0xCB67bE5c54eab9462967eE3C03C35bfFfeB801cD",
		Blockchain: dia.ETHEREUM,
		Symbol:     "GBI",
		Name:       "GalacticBlueIndex",
		Decimals:   uint8(18),
	}
}

type oldCryptoIndexConstituent struct {
	Name              string
	Symbol            string
	Address           string
	Price             float64
	PriceYesterday    float64
	PriceYesterweek   float64
	CirculatingSupply float64
	Weight            float64
	Percentage        float64
	CappingFactor     float64
	NumBaseTokens     float64
}

type oldCryptoIndex struct {
	Name              string
	Value             float64
	Price             float64
	Price1h           float64
	Price24h          float64
	Price7d           float64
	Price14d          float64
	Price30d          float64
	Volume24hUSD      float64
	CirculatingSupply float64
	Divisor           float64
	CalculationTime   time.Time
	Constituents      []oldCryptoIndexConstituent
}

// func pain() {

// 	// timeFinalString is the last timestamp for which trades are read from influx in Unix time seconds.
// 	timeInitString := utils.Getenv("TIME_INIT", "1608338769")
// 	timeInitInt, err := strconv.ParseInt(timeInitString, 10, 64)
// 	if err != nil {
// 		log.Error("parse timeInit: ", err)
// 	}
// 	timeInit := time.Unix(timeInitInt, 0)

// 	timeFinalString := utils.Getenv("TIME_FINAL", "1645093427")
// 	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
// 	if err != nil {
// 		log.Error("parse timeFinal: ", err)
// 	}
// 	timeFinal := time.Unix(timeFinalInt, 0)
// 	log.Info("timeInit: ", timeInit)
// 	log.Info("timeFinal: ", timeFinal)

// 	stepSizeString := utils.Getenv("STEP_SIZE_MINUTES", "30")
// 	stepSize, err := strconv.ParseInt(stepSizeString, 10, 64)
// 	if err != nil {
// 		log.Error("parse timeFinal: ", err)
// 	}

// 	indexName := utils.Getenv("INDEX_NAME", "GBI")

// 	ds, err := models.NewDataStore()
// 	if err != nil {
// 		log.Fatal("datastore: ", err)
// 	}

// 	rdb, err := models.NewRelDataStore()
// 	if err != nil {
// 		log.Fatal("relational datastore: ", err)
// 	}

// 	assetMap, err := getAssetMap(symbolsMap, rdb)
// 	if err != nil {
// 		log.Error("get asset map: ", err)
// 	}

// 	for timeFinal.After(timeInit) {
// 		endTime := timeInit.Add(time.Duration(stepSize) * time.Minute)

// 		// Fetch old index values.
// 		log.Infof("fetch index values in time range %v -- %v ...", timeInit, endTime)
// 		oldIndexVals, err := getOldIndexFromAPI(indexName, timeInit, endTime)
// 		if err != nil {
// 			log.Errorf("fetch index values in time range %v -- %v: %v", timeInit, endTime, err)
// 		}
// 		log.Infof("...fetch done. Found %d values.", len(oldIndexVals))

// 		for i := range oldIndexVals {

// 			// Fix GBI values.
// 			if indexName == "GBI" {
// 				log.Infof("amend GBI values for i=%d", i)
// 				oldIndexVals[i] = amendGBI(oldIndexVals[i], assetMap, ds)
// 			}

// 			// Map old index to new index values.
// 			newIndexVal := assignIndexVal(oldIndexVals[i], assetMap)

// 			// Save new index values.
// 			err = ds.SetCryptoIndex(&newIndexVal)
// 			if err != nil {
// 				log.Error("set crypto index: ", err)
// 			}

// 		}

// 		timeInit = endTime
// 	}
// 	log.Info("...done copying.")
// 	time.Sleep(24 * 60 * time.Minute)

// }

func main() {

	// timeFinalString is the last timestamp for which trades are read from influx in Unix time seconds.
	timeInitString := utils.Getenv("TIME_INIT", "1608338769")
	timeInitInt, err := strconv.ParseInt(timeInitString, 10, 64)
	if err != nil {
		log.Error("parse timeInit: ", err)
	}
	timeInit := time.Unix(timeInitInt, 0)

	timeFinalString := utils.Getenv("TIME_FINAL", "1645093427")
	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	timeFinal := time.Unix(timeFinalInt, 0)
	log.Info("timeInit: ", timeInit)
	log.Info("timeFinal: ", timeFinal)

	stepSizeString := utils.Getenv("STEP_SIZE_MINUTES", "30")
	stepSize, err := strconv.ParseInt(stepSizeString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}

	indexName := utils.Getenv("INDEX_NAME", "SCIFI")

	numRangesString := utils.Getenv("NUM_RANGES", "12")
	numRanges, err := strconv.Atoi(numRangesString)
	if err != nil {
		log.Error("parse num ranges: ", err)
	}

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore: ", err)
	}

	assetMap, err := getAssetMap(symbolsMap, rdb)
	if err != nil {
		log.Error("get asset map: ", err)
	}

	a := timeInit
	b := timeFinal
	totalSize := b.Sub(a)
	sizeRange := totalSize / time.Duration(numRanges)
	starttimes := []time.Time{}
	endtimes := []time.Time{}
	starttime := timeInit
	for k := 0; k < numRanges; k++ {
		starttimes = append(starttimes, starttime)
		endtimes = append(endtimes, starttime.Add(sizeRange))
		starttime = starttime.Add(sizeRange)
	}

	var wg sync.WaitGroup
	for i := 0; i < len(starttimes); i++ {

		wg.Add(1)
		go func(starttime, endtime time.Time, step int64, name string, aMap map[string]dia.Asset, datastore *models.DB, w *sync.WaitGroup) {
			defer w.Done()
			processIndexVals(starttime, endtime, step, name, aMap, datastore)
		}(starttimes[i], endtimes[i], stepSize, indexName, assetMap, ds, &wg)

	}

	wg.Wait()

}

func processIndexVals(timeInit, timeFinal time.Time, stepSize int64, indexName string, assetMap map[string]dia.Asset, ds *models.DB) {
	for timeFinal.After(timeInit) {
		endTime := timeInit.Add(time.Duration(stepSize) * time.Minute)

		// Fetch old index values.
		log.Infof("fetch index values in time range %v -- %v ...", timeInit, endTime)
		oldIndexVals, err := getOldIndexFromAPI(indexName, timeInit, endTime)
		if err != nil {
			log.Errorf("fetch index values in time range %v -- %v: %v", timeInit, endTime, err)
		}
		log.Infof("...fetch done. Found %d values.", len(oldIndexVals))

		for i := range oldIndexVals {

			// Fix GBI values.
			if indexName == "GBI" {
				log.Infof("amend GBI values for i=%d", i)
				oldIndexVals[i] = amendGBI(oldIndexVals[i], assetMap, ds)
			}

			// Map old index to new index values.
			newIndexVal := assignIndexVal(oldIndexVals[i], assetMap)

			// Save new index values.
			err = ds.SetCryptoIndex(&newIndexVal)
			if err != nil {
				log.Error("set crypto index: ", err)
			}
			log.Infof("successfully set index %s at time %v with value %v.", newIndexVal.Asset.Symbol, newIndexVal.CalculationTime, newIndexVal.Value)

		}

		timeInit = endTime
	}
	log.Info("...done copying.")
	time.Sleep(24 * 60 * time.Minute)
}

func amendGBI(oldIndex oldCryptoIndex, assetMap map[string]dia.Asset, ds *models.DB) (oldIndexUpdated oldCryptoIndex) {
	// 1. look for wrong numBaseTokens in SPICE
	// 2. Correct numBaseTokens
	// 3. Fetch correct price for SPICE
	// 4. Calculate Index (numBaseTokens, price)
	for i := range oldIndex.Constituents {
		if oldIndex.Constituents[i].Symbol == "SPICE" {
			if oldIndex.Constituents[i].NumBaseTokens > float64(1390505109261745600) {
				oldIndex.Constituents[i].NumBaseTokens = float64(102974575097533460)
				newPrice, err := ds.GetLastPriceBefore(assetMap["SPICE"], "MAIR120", "", oldIndex.CalculationTime)
				if err != nil {
					log.Error("fetch price for SPICE: ", err)
				}
				oldIndex.Constituents[i].Price = newPrice.Price
			}
		}
	}

	oldIndex.Value = GetOldIndexValue("GBI", oldIndex.Constituents)

	for i := range oldIndex.Constituents {
		currPercentage := (oldIndex.Constituents[i].Price * oldIndex.Constituents[i].NumBaseTokens * 1e-16) / oldIndex.Value //1e-16 because index value is 100 at start
		oldIndex.Constituents[i].Percentage = currPercentage
	}

	return
}

// assignIndexVal assigns a crypto index in the new format to a crypto index in the old format.
func assignIndexVal(oldIndex oldCryptoIndex, assetMap map[string]dia.Asset) (newIndex models.CryptoIndex) {

	newIndex.Asset = indexMap[oldIndex.Name]
	newIndex.Value = oldIndex.Value
	newIndex.Price = oldIndex.Price
	newIndex.Price1h = oldIndex.Price1h
	newIndex.Price24h = oldIndex.Price24h
	newIndex.Price7d = oldIndex.Price7d
	newIndex.Price14d = oldIndex.Price14d
	newIndex.Price30d = oldIndex.Price30d
	newIndex.Volume24hUSD = oldIndex.Volume24hUSD
	newIndex.CirculatingSupply = oldIndex.CirculatingSupply
	newIndex.Divisor = oldIndex.Divisor
	newIndex.CalculationTime = oldIndex.CalculationTime

	var constituents []models.CryptoIndexConstituent
	for _, constituent := range oldIndex.Constituents {
		constituents = append(constituents, assignConstituent(constituent, assetMap))
	}
	newIndex.Constituents = constituents

	return
}

// assignConstituent assigns a constituent in the new format to a constituent in the old format.
func assignConstituent(oldConstituent oldCryptoIndexConstituent, assetMap map[string]dia.Asset) (newConstituent models.CryptoIndexConstituent) {
	newConstituent.Asset = assetMap[oldConstituent.Symbol]
	newConstituent.Price = oldConstituent.Price
	newConstituent.PriceYesterday = oldConstituent.PriceYesterday
	newConstituent.PriceYesterweek = oldConstituent.PriceYesterweek
	newConstituent.CirculatingSupply = oldConstituent.CirculatingSupply
	newConstituent.Weight = oldConstituent.Weight
	newConstituent.Percentage = oldConstituent.Percentage
	newConstituent.CappingFactor = oldConstituent.CappingFactor
	newConstituent.NumBaseTokens = oldConstituent.NumBaseTokens
	return
}

func GetOldIndexValue(indexSymbol string, currentConstituents []oldCryptoIndexConstituent) float64 {
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

// getOldIndexFromAPI returns index values from api.diadata in the given time range.
func getOldIndexFromAPI(name string, timeInit, timeFinal time.Time) (indexVals []oldCryptoIndex, err error) {
	var resp []byte
	timeInitString := strconv.Itoa(int(timeInit.Unix()))
	timeFinalString := strconv.Itoa(int(timeFinal.Unix()))
	resp, _, err = utils.GetRequest("https://api.diadata.org/v1/index/" + name + "?starttime=" + timeInitString + "&endtime=" + timeFinalString + "&maxResults=0")
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &indexVals)
	if err != nil {
		return
	}
	return
}

// getAssetMap returns a map that maps the symbol of a crypto index constituent to its underlying dia.Asset.
func getAssetMap(symbolsMap map[string]string, rdb *models.RelDB) (map[string]dia.Asset, error) {
	assetMap := make(map[string]dia.Asset)
	for key := range symbolsMap {
		asset, err := rdb.GetAsset(symbolsMap[key], dia.ETHEREUM)
		if err != nil {
			log.Error("get asset: ", symbolsMap[key])
			return assetMap, err
		}
		assetMap[key] = asset
	}
	return assetMap, nil
}
