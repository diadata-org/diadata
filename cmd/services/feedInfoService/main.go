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
	log "github.com/sirupsen/logrus"
)

var (
	basetokenMap = make(map[string]dia.Asset)
	assets       []dia.Asset
)

const (
	LOOKBACK_SECONDS             = 86400
	BIN_DURATION_SECONDS         = 120
	BIN_THRESHOLD                = 2
	COLLECTION_FREQUENCY_SECONDS = 60 * 60 * 24
	NUM_RANGES                   = "24"
)

type timeRange struct {
	tLeft  time.Time
	tRight time.Time
}

func main() {

	datastore, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	relDB, err := models.NewRelDataStore()
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
	for _, asset := range assets {
		updateStatsPerAsset(asset, time.Now(), numRanges, datastore, relDB)
	}
	log.Info("...update done.")

	for {
		select {
		case tFinal := <-ticker.C:
			for _, asset := range assets {
				updateStatsPerAsset(asset, tFinal, numRanges, datastore, relDB)
			}
			log.Info("...update done.")
		}
	}

}

func updateStatsPerAsset(asset dia.Asset, tFinal time.Time, numRanges int, datastore *models.DB, relDB *models.RelDB) {

	log.Infof("start processing data for %s -- %s....", asset.Blockchain, asset.Address)

	tInit := tFinal.Add(-time.Duration(LOOKBACK_SECONDS * time.Second))
	// Make time ranges for batching the trades getter.
	starttimes, endtimes := utils.MakeTimeRanges(tInit, tFinal, numRanges)

	var (
		pairMap            = make(map[string]struct{})
		pairExchangeMap    = make(map[string]struct{})
		pairExchangeVolMap = make(map[dia.Pair]map[string]float64)
		aggVolumes         []dia.AggregatedVolume
		tradesDistribution dia.TradesDistribution
		tradesCount        []int
		binMap             = make(map[timeRange]int)
		binDuration        = time.Duration(BIN_DURATION_SECONDS * time.Second)
	)

	// Constant trades distribution params.
	tradesDistribution.Asset = asset
	tradesDistribution.Timestamp = tFinal
	tradesDistribution.Threshold = BIN_THRESHOLD
	tradesDistribution.SizeBinSeconds = int64(BIN_DURATION_SECONDS)
	tradesDistribution.TimeRangeSeconds = int64(LOOKBACK_SECONDS)

	// Make bin map, mapping a time bin of size BIN_DURATION_SECONDS to the number of trades in this bin.
	leftTime := tInit
	for leftTime.Before(tFinal) {
		tr := timeRange{
			tLeft:  leftTime,
			tRight: leftTime.Add(binDuration),
		}
		binMap[tr] = 0
		leftTime = leftTime.Add(binDuration)
	}

	for i := range starttimes {

		// 1. Fetch trades for @asset in batches.
		trades, err := datastore.GetTradesByExchangesFull(asset, []dia.Asset{}, []string{}, true, starttimes[i], endtimes[i], 0)
		if err != nil {
			log.Warnf("GetTradesByExchangesFull for asset %s in time range %v -- %v: %v", asset.Address, starttimes[i], endtimes[i], err)
		}

		// 2. Get volumes per exchange and per pair.
		pairMap, pairExchangeMap, pairExchangeVolMap = computePairStats(asset, trades, pairMap, pairExchangeMap, pairExchangeVolMap, relDB)

		// 3. Get statistics on trades' frequency and distribution
		binMap = computeBinMap(trades, binMap)

	}

	// Get trades statistics from binMap.
	for _, value := range binMap {
		tradesCount = append(tradesCount, value)
		tradesDistribution.NumTradesTotal += value
		if value < BIN_THRESHOLD {
			tradesDistribution.NumLowBins += 1
		}
	}
	tradesDistribution.AvgNumPerBin = average(tradesCount)
	tradesDistribution.StdDeviation = math.Sqrt(variance(tradesCount))

	err := relDB.SetTradesDistribution(tradesDistribution)
	if err != nil {
		log.Errorf("set trades distributionfor %s - %s: %v", asset.Address, asset.Blockchain, err)
	}

	// Fill pairVolumes slice, i.e. sort by exchange and pair.
	for pair, exchangeMap := range pairExchangeVolMap {
		for exchange, value := range exchangeMap {
			pairVolume := dia.AggregatedVolume{
				Pair:             pair,
				Exchange:         exchange,
				Volume:           value,
				Timestamp:        tFinal,
				TimeRangeSeconds: LOOKBACK_SECONDS,
			}
			aggVolumes = append(aggVolumes, pairVolume)
		}
	}

	for _, pv := range aggVolumes {
		err = relDB.SetAggregatedVolume(pv)
		if err != nil {
			log.Errorf("SetAggregatedVolume for %s - %s: %v", asset.Address, asset.Blockchain, err)
		}
	}
	log.Info("...done processing.")

}

// computeBinMap returns the mapping of a time bin of size BIN_DURATION_SECONDS to the number of trades in this bin.
func computeBinMap(trades []dia.Trade, binMap map[timeRange]int) map[timeRange]int {
	for _, trade := range trades {
		binMap = mapTradeToBin(trade, binMap)
	}
	return binMap
}

// mapTradeToBin maps a trade to the corresponding time bin and increments the counter.
func mapTradeToBin(trade dia.Trade, binMap map[timeRange]int) map[timeRange]int {
	for key := range binMap {
		if inBin(trade, key) {
			binMap[key] += 1
			return binMap
		}
	}
	return binMap
}

func inBin(trade dia.Trade, tr timeRange) bool {
	if !trade.Time.Before(tr.tLeft) && !trade.Time.After(tr.tRight) {
		return true
	}
	return false
}

func computePairStats(
	asset dia.Asset,
	trades []dia.Trade,
	pairMap map[string]struct{},
	pairExchangeMap map[string]struct{},
	pairExchangeVolMap map[dia.Pair]map[string]float64,
	relDB *models.RelDB,
) (
	map[string]struct{},
	map[string]struct{},
	map[dia.Pair]map[string]float64,
) {

	quotetoken, err := relDB.GetAsset(asset.Address, asset.Blockchain)
	if err != nil {
		log.Fatal("get quotetoken: ", err)
	}

	for _, trade := range trades {
		basetoken, err := getBasetoken(basetokenMap, trade.BaseToken, relDB)
		if err != nil {
			log.Errorf("get basetoken %v: %v", trade.BaseToken, err)
			basetoken.Address = trade.BaseToken.Address
			basetoken.Blockchain = trade.BaseToken.Blockchain
			basetoken.Symbol = trade.BaseToken.Symbol
			basetoken.Name = trade.BaseToken.Name
			basetoken.Decimals = trade.BaseToken.Decimals
		}
		pair := dia.Pair{
			QuoteToken: quotetoken,
			BaseToken:  basetoken,
		}

		pairID := pair.Identifier()
		pairExchangeID := pair.PairExchangeIdentifier(trade.Source)
		if _, ok := pairMap[pairID]; !ok {
			// Pair is not registered yet.
			pairExchangeVolMap[pair] = make(map[string]float64)
			pairExchangeVolMap[pair][trade.Source] = trade.EstimatedUSDPrice * math.Abs(trade.Volume)
			pairMap[pairID] = struct{}{}
		} else {
			if _, ok := pairExchangeMap[pairExchangeID]; !ok {
				// Pair is registered, but not on this exchange yet.
				pairExchangeVolMap[pair][trade.Source] = trade.EstimatedUSDPrice * math.Abs(trade.Volume)
				pairExchangeMap[pairExchangeID] = struct{}{}
			} else {
				// Pair is already registered for given exchange.
				pairExchangeVolMap[pair][trade.Source] += trade.EstimatedUSDPrice * math.Abs(trade.Volume)
			}
		}
	}

	return pairMap, pairExchangeMap, pairExchangeVolMap
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
