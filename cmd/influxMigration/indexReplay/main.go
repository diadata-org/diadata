package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
)

var (
	log              *logrus.Logger
	symbolsMap       map[string]string
	cappingMap       map[string]float64
	numBaseTokensMap map[string]float64
	indexMap         map[string]dia.Asset
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

	cappingMap = make(map[string]float64)
	cappingMap["0x6B3595068778DD592e39A122f4f5a5cF09C90fE2"] = float64(0.8481712413)
	cappingMap["0x84cA8bc7997272c7CfB4D0Cd3D55cd942B3c9419"] = float64(1.829351993)
	cappingMap["0x1fDaB294EDA5112B7d066ED8F2E4E562D5bCc664"] = float64(1111.932789)
	cappingMap["0x1776e1F26f98b1A5dF9cD347953a26dd3Cb46671"] = float64(0.8501919544)
	cappingMap["0x1cEB5cB57C4D4E2b2433641b95Dd330A33185A44"] = float64(0.8474172955)
	cappingMap["0xD533a949740bb3306d119CC777fa900bA034cd52"] = float64(0.8452712665)
	cappingMap["0xbC396689893D065F41bc2C6EcbeE5e0085233447"] = float64(0.8496875939)
	cappingMap["0x4E15361FD6b4BB609Fa63C81A2be19d873717870"] = float64(0.5305555209)
	cappingMap["0xba100000625a3754423978a60c9317c58a424e3D"] = float64(0.8494870174)
	cappingMap["0x50D1c9771902476076eCFc8B2A83Ad6b9355a4c9"] = float64(0.6908255324)

	numBaseTokensMap = make(map[string]float64)
	numBaseTokensMap["0x6B3595068778DD592e39A122f4f5a5cF09C90fE2"] = float64(34969303170127981)
	numBaseTokensMap["0x84cA8bc7997272c7CfB4D0Cd3D55cd942B3c9419"] = float64(8067221838047583)
	numBaseTokensMap["0x1fDaB294EDA5112B7d066ED8F2E4E562D5bCc664"] = float64(382922797228701943)
	numBaseTokensMap["0x1776e1F26f98b1A5dF9cD347953a26dd3Cb46671"] = float64(1610484614974134)
	numBaseTokensMap["0x1cEB5cB57C4D4E2b2433641b95Dd330A33185A44"] = float64(47725897426893)
	numBaseTokensMap["0xD533a949740bb3306d119CC777fa900bA034cd52"] = float64(127432811325197740)
	numBaseTokensMap["0xbC396689893D065F41bc2C6EcbeE5e0085233447"] = float64(4920853921311107)
	numBaseTokensMap["0x4E15361FD6b4BB609Fa63C81A2be19d873717870"] = float64(234454930204332481)
	numBaseTokensMap["0xba100000625a3754423978a60c9317c58a424e3D"] = float64(3941681400720949)
	numBaseTokensMap["0x50D1c9771902476076eCFc8B2A83Ad6b9355a4c9"] = float64(16635301231735492)

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

func main() {

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore: ", err)
	}

	indexName := utils.Getenv("INDEX_NAME", "SCIFI")

	// timeFinalString is the last timestamp for which trades are read from influx in Unix time seconds.
	timeInitString := utils.Getenv("TIME_INIT", "1645138800")
	timeInitInt, err := strconv.ParseInt(timeInitString, 10, 64)
	if err != nil {
		log.Error("parse timeInit: ", err)
	}
	timeInit := time.Unix(timeInitInt, 0)

	timeFinalString := utils.Getenv("TIME_FINAL", "1645182000")
	timeFinalInt, err := strconv.ParseInt(timeFinalString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}
	timeFinal := time.Unix(timeFinalInt, 0)
	log.Info("timeInit: ", timeInit)
	log.Info("timeFinal: ", timeFinal)

	numRangesString := utils.Getenv("NUM_RANGES", "6")
	numRanges, err := strconv.Atoi(numRangesString)
	if err != nil {
		log.Error("parse num ranges: ", err)
	}

	log.Infof("make %d time-ranges..", numRanges)
	starttimes, endtimes := makeTimeRanges(timeInit, timeFinal, numRanges)
	log.Info("..done making time-ranges.")

	for i := 0; i < len(starttimes); i++ {
		log.Infof("time-range: [%v, %v].", starttimes[i], endtimes[i])
	}

	stepSizeString := utils.Getenv("STEP_SIZE_MINUTES", "10")
	stepSize, err := strconv.ParseInt(stepSizeString, 10, 64)
	if err != nil {
		log.Error("parse timeFinal: ", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < len(starttimes); i++ {

		wg.Add(1)
		go func(starttime, endtime time.Time, step int64, name string, datastore *models.DB, w *sync.WaitGroup) {
			defer w.Done()
			processIndexVals(starttime, endtime, step, name, datastore)
		}(starttimes[i], endtimes[i], stepSize, indexName, ds, &wg)

	}

	wg.Wait()

}

func processIndexVals(timeInit, timeFinal time.Time, stepSize int64, indexSymbol string, ds *models.DB) {

	index := indexMap[indexSymbol]

	for timeFinal.After(timeInit) {

		endTime := timeInit.Add(time.Duration(stepSize) * time.Minute)

		// Fetch old index values.
		log.Infof("fetch index values in time range %v -- %v ...", timeInit, endTime)
		oldIndexVals, err := ds.GetCryptoIndex(timeInit, endTime, indexSymbol, 0)
		if err != nil {
			log.Errorf("fetch index values in time range %v -- %v: %v", timeInit, endTime, err)
		}
		log.Infof("...fetch done. Found %d values.", len(oldIndexVals))

		for i := range oldIndexVals {

			constituents := oldIndexVals[i].Constituents
			for j := range constituents {
				// 1. Substitute cappingFactor
				constituents[j].CappingFactor = cappingMap[constituents[j].Asset.Address]
				// 2. Substitute numBaseTokens
				if indexSymbol == "SCIFI" {
					constituents[j].NumBaseTokens = numBaseTokensMap[constituents[j].Asset.Address]
				}
				// 3. Substitute price
				if constituents[j].Price == float64(0) {
					price, err := ds.GetAssetPriceUSD(constituents[j].Asset, oldIndexVals[i].CalculationTime)
					if err != nil {
						log.Errorf("get asset price for address %s", constituents[j].Asset.Address)
					}
					constituents[j].Price = price
				}
			}

			currIndexValue := models.GetIndexValue(indexSymbol, constituents)
			computePercentages(indexSymbol, &constituents, currIndexValue)
			for j := range constituents {
				err = ds.SetCryptoIndexConstituent(&constituents[j], oldIndexVals[i].Asset, oldIndexVals[i].CalculationTime)
				if err != nil {
					log.Error("set crypto index constituent: ", err)
				}
				log.Infof("set constituent at time %v: %v", oldIndexVals[i].CalculationTime, constituents[j])
			}

			// Compute new index value.
			indexValue := models.GetIndexValue(index.Symbol, constituents)
			index := indexValueCalculation(constituents, oldIndexVals[i], indexValue, ds)

			// Save index.
			err = ds.SetCryptoIndex(&index)
			if err != nil {
				log.Error(err)
			}
			log.Infof("set crypto index at calculation time %v: value, divisor, price -- %v, %v, %v ", index.CalculationTime, index.Value, index.Divisor, index.Price)

			// log.Infof("successfully set index %s at time %v with value %v.", index.Asset.Symbol, index.CalculationTime, index.Value)

		}

		timeInit = endTime
	}
	log.Info("...done copying.")
	time.Sleep(24 * 60 * time.Minute)
}

func indexValueCalculation(currentConstituents []models.CryptoIndexConstituent, cryptoIndex models.CryptoIndex, indexValue float64, datastore *models.DB) models.CryptoIndex {

	var price float64
	indexQuotation, err := datastore.GetAssetQuotation(cryptoIndex.Asset, cryptoIndex.CalculationTime)
	if err == nil {
		price = indexQuotation.Price
	}

	index := models.CryptoIndex{
		Asset:           cryptoIndex.Asset,
		Price:           price,
		Value:           indexValue,
		CalculationTime: cryptoIndex.CalculationTime,
		Constituents:    currentConstituents,
		Divisor:         cryptoIndex.Divisor,
	}
	log.Info("Index: ", index)
	return index
}

func computePercentages(indexSymbol string, constituents *[]models.CryptoIndexConstituent, currentIndexValue float64) {
	if indexSymbol == "SCIFI" {
		for i := range *constituents {
			currPercentage := ((*constituents)[i].Price * (*constituents)[i].CirculatingSupply * (*constituents)[i].CappingFactor) / currentIndexValue
			(*constituents)[i].Percentage = currPercentage
		}
	} else {
		// GBI
		for i := range *constituents {
			currPercentage := ((*constituents)[i].Price * (*constituents)[i].NumBaseTokens * 1e-16) / currentIndexValue //1e-16 because index value is 100 at start
			(*constituents)[i].Percentage = currPercentage
		}
	}
}

// makeTimeRanges returns @numRanges start- and endtimes partitioning [@timeInit, @timeFinal] in intervals of identical size.
func makeTimeRanges(timeInit, timeFinal time.Time, numRanges int) (starttimes, endtimes []time.Time) {
	a := timeInit
	b := timeFinal
	totalSize := b.Sub(a)
	sizeRange := totalSize / time.Duration(numRanges)
	starttime := timeInit
	for k := 0; k < numRanges; k++ {
		starttimes = append(starttimes, starttime)
		endtimes = append(endtimes, starttime.Add(sizeRange))
		starttime = starttime.Add(sizeRange)
	}
	return
}
