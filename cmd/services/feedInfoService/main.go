package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/prometheus/common/log"
)

var (
	datastore *models.DB
	relDB     *models.RelDB
	err       error
	// local caches for pair-exchange and basetokens
	pairMap            = make(map[string]struct{})
	pairExchangeMap    = make(map[string]struct{})
	pairExchangeVolMap = make(map[dia.Pair]map[string]float64)
	basetokenMap       = make(map[string]dia.Asset)
	assets             []dia.Asset
)

const (
	LOOKBACK_SECONDS             = 86400
	BIN_DURATION_SECONDS         = 120
	BIN_THRESHOLD                = 2
	COLLECTION_FREQUENCY_SECONDS = 60 * 60 * 24
	NUM_RANGES                   = "24"
)

func main() {

	datastore, err = models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	relDB, err = models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore:", err)
	}

	numRangesString := utils.Getenv("NUM_RANGES", NUM_RANGES)
	numRanges, err := strconv.Atoi(numRangesString)
	if err != nil {
		log.Error("parse num ranges: ", err)
	}

	// Import assets for which we collect statistics from config file.
	assets, err = getAssetsFromConfig("feedInfoService/assets")
	if err != nil {
		log.Error("error getting tokens for which pairs should be reversed: ", err)
	}

	ticker := time.NewTicker(COLLECTION_FREQUENCY_SECONDS * time.Second)

	// Initial run.
	updateStats(assets, time.Now(), numRanges, datastore, relDB)
	log.Info("...update done.")

	for {
		select {
		case tFinal := <-ticker.C:
			updateStats(assets, tFinal, numRanges, datastore, relDB)
			log.Info("...update done.")
		}
	}

}

func updateStats(assets []dia.Asset, tFinal time.Time, numRanges int, datastore *models.DB, relDB *models.RelDB) {
	tInit := tFinal.Add(-time.Duration(LOOKBACK_SECONDS * time.Second))

	// Make time ranges for batching the trades getter.
	starttimes, endtimes := utils.MakeTimeRanges(tInit, tFinal, numRanges)

	for _, asset := range assets {
		// 1. Fetch trades for @asset in batches.
		log.Infof("collect trades for: %s - %s ...", asset.Address, asset.Blockchain)
		trades, err := datastore.GetTradesByExchangesBatchedFull(asset, []string{}, true, starttimes, endtimes)
		if err != nil {
			log.Fatal("GetTradesByExchangesBatched: ", err)
		}
		log.Info("...done collecting trades.")

		// 2. Get volumes per exchange and per pair.
		aggVolumes := computePairStats(asset, trades, tFinal)
		for _, pv := range aggVolumes {
			err = relDB.SetAggregatedVolume(pv)
			if err != nil {
				log.Errorf("SetAggregatedVolume for %s - %s: %v", asset.Address, asset.Blockchain, err)
			}
		}

		// 3. Get statistics on trades' frequency and distribution
		binDuration := time.Duration(BIN_DURATION_SECONDS * time.Second)
		starttime := tInit
		tradesFreq := computeTradesFrequency(asset, trades, binDuration, starttime, tFinal)
		err = relDB.SetTradesDistribution(tradesFreq)
		if err != nil {
			log.Error("set trades distributionfor %s - %s: %v", asset.Address, asset.Blockchain, err)
		}

	}
}

// computePairStats takes a slice of trades with @asset as quotetoken asset. It returns a slice of aggregated volumes.
func computePairStats(asset dia.Asset, trades []dia.Trade, timestamp time.Time) (pairVolumes []dia.AggregatedVolume) {
	quotetoken, err := relDB.GetAsset(asset.Address, asset.Blockchain)
	if err != nil {
		log.Fatal("get quotetoken: ", err)
	}

	for _, trade := range trades {
		basetoken, err := getBasetoken(basetokenMap, trade.BaseToken, relDB)
		if err != nil {
			log.Error("get basetoken: ", err)
		}
		pair := dia.Pair{
			QuoteToken: quotetoken,
			BaseToken:  basetoken,
		}

		pairID := getPairIdentifier(pair)
		pairExchangeID := getPairExchangeIdentifier(pair, trade.Source)
		if _, ok := pairMap[pairID]; !ok {
			// Pair is not registered yet.
			pairExchangeVolMap[pair] = make(map[string]float64)
			pairExchangeVolMap[pair][trade.Source] = trade.EstimatedUSDPrice * math.Abs(trade.Volume)
		} else {
			if _, ok := pairExchangeMap[pairExchangeID]; !ok {
				// Pair is registered, but not on this exchange yet
				pairExchangeVolMap[pair][trade.Source] = trade.EstimatedUSDPrice * math.Abs(trade.Volume)
			} else {
				// Pair is already registered for given exchange
				pairExchangeVolMap[pair][trade.Source] += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
			}
		}
	}

	// Fill pairVolumes slice, i.e. sort by exchange and pair.
	for pair, exchangeMap := range pairExchangeVolMap {
		for exchange, value := range exchangeMap {
			pairVolume := dia.AggregatedVolume{
				Pair:             pair,
				Exchange:         exchange,
				Volume:           value,
				Timestamp:        timestamp,
				TimeRangeSeconds: LOOKBACK_SECONDS,
			}
			pairVolumes = append(pairVolumes, pairVolume)
		}
	}
	return
}

// computeTradesFrequency returns statistics on the trades of @asset in the given time-range, assuming a @binDuration.
func computeTradesFrequency(asset dia.Asset, trades []dia.Trade, binDuration time.Duration, starttime time.Time, endtime time.Time) (tradesDistribution dia.TradesDistribution) {

	var tradesCount []int
	tradesCount = append(tradesCount, 0)
	var binCount int
	var lowBinCount int
	for _, trade := range trades {
		if trade.Time.Before(starttime.Add(binDuration)) {
			tradesCount[binCount] += 1
		} else {
			if tradesCount[binCount] < int(BIN_THRESHOLD) {
				lowBinCount++
			}
			tradesCount = append(tradesCount, 0)
			binCount++
			starttime = starttime.Add(binDuration)
		}
	}
	tradesDistribution.Asset = asset
	tradesDistribution.NumTradesTotal = len(trades)
	tradesDistribution.NumLowBins = lowBinCount
	tradesDistribution.Threshold = BIN_THRESHOLD
	tradesDistribution.SizeBinSeconds = int64(BIN_DURATION_SECONDS)
	tradesDistribution.AvgNumPerBin = average(tradesCount)
	tradesDistribution.StdDeviation = math.Sqrt(variance(tradesCount))
	tradesDistribution.TimeRangeSeconds = int64(LOOKBACK_SECONDS)
	tradesDistribution.Timestamp = endtime
	return
}

// getBasetoken fetches basetoken info either from local cache if in there or from postgres if not.
func getBasetoken(basetokenMap map[string]dia.Asset, token dia.Asset, relDB *models.RelDB) (basetoken dia.Asset, err error) {
	if val, ok := basetokenMap[token.Blockchain+"-"+token.Address]; !ok {
		basetoken, err = relDB.GetAsset(token.Address, token.Blockchain)
		if err != nil {
			return
		}
		basetokenMap[token.Blockchain+"-"+token.Address] = basetoken
	} else {
		basetoken = val
	}
	return
}

func getPairIdentifier(pair dia.Pair) string {
	return pair.QuoteToken.Blockchain + "-" + pair.QuoteToken.Address + "-" + pair.BaseToken.Blockchain + "-" + pair.BaseToken.Address
}

func getPairExchangeIdentifier(pair dia.Pair, exchange string) string {
	return getPairIdentifier(pair) + "-" + exchange
}

func average(series []int) (avg float64) {
	if len(series) == 0 {
		return
	}

	var sum int
	for _, item := range series {
		sum += item
	}
	avg = float64(int64(sum)) / float64(len(series))
	return avg
}

func variance(series []int) (sd float64) {
	if len(series) == 0 {
		return
	}

	var variance float64
	for _, item := range series {
		variance += math.Pow(float64(int64(item))-average(series), float64(2))
	}
	return variance / float64(int64(len(series)))
}

// getAssetsFromConfig returns a list of address-blockchain pairs from the config file with @path.
func getAssetsFromConfig(path string) (assets []dia.Asset, err error) {

	// Load file and read data
	filehandle := configCollectors.ConfigFileConnectors(path, ".json")
	jsonFile, err := os.Open(filehandle)
	if err != nil {
		return
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	type assetList struct {
		AllAssets []dia.Asset `json:"Tokens"`
	}
	var allAssets assetList
	err = json.Unmarshal(byteData, &allAssets)
	if err != nil {
		return
	}
	assets = allAssets.AllAssets
	return
}
