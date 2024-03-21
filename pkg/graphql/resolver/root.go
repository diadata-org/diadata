package resolver

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	queryhelper "github.com/diadata-org/diadata/pkg/dia/helpers/queryHelper"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"

	models "github.com/diadata-org/diadata/pkg/model"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/sirupsen/logrus"
)

var (
	log                   = logrus.New()
	EXCHANGES             = scrapers.Exchanges
	BLOCKCHAINS           = scrapers.Blockchains
	lookbackTradesNumber  = 10
	errInvalidInputParams = errors.New("invalid input params")
	statusCodesMap        = make(map[string]int32)
	statusError           = "err"
)

const (
	PAIR_SEPARATOR                 = "-"
	LIST_SEPARATOR                 = ","
	TRADE_VOLUME_THRESHOLD_DEFAULT = float64(0.001)
)

func init() {
	statusCodesMap[""] = int32(0)
	statusCodesMap[statusError] = int32(1)
	statusCodesMap["pair not available on"] = int32(2)
	statusCodesMap["exchange name not valid"] = int32(3)
	statusCodesMap["not available on"] = int32(7)
}

// Resolver is the root resolver
type DiaResolver struct {
	DS              models.DB
	RelDB           models.RelDB
	InfluxBatchSize int64
}

func (r *DiaResolver) GetSupply(ctx context.Context, args struct{ Symbol graphql.NullString }) (*SupplyResolver, error) {
	if containsSpecialChars(*args.Symbol.Value) {
		return nil, errInvalidInputParams
	}
	q, err := r.DS.GetLatestSupply(*args.Symbol.Value, &r.RelDB)
	if err != nil {
		return nil, err
	}
	return &SupplyResolver{q: q}, nil
}

func (r *DiaResolver) GetSupplies(ctx context.Context, args struct{ Symbol graphql.NullString }) (*[]*SupplyResolver, error) {
	starttime := time.Unix(1, 0)
	endtime := time.Now()
	if containsSpecialChars(*args.Symbol.Value) {
		return nil, errInvalidInputParams
	}

	q, err := r.DS.GetSupply(*args.Symbol.Value, starttime, endtime, &r.RelDB)
	if err != nil {
		return nil, err
	}

	var sr []*SupplyResolver

	for _, supply := range q {
		sr = append(sr, &SupplyResolver{q: &supply})

	}
	return &sr, nil
}

type TradeBlock struct {
	Trades []dia.Trade
}

type Asset struct {
	Address    graphql.NullString
	BlockChain graphql.NullString
}

type FeedSelection struct {
	Address            graphql.NullString
	Blockchain         graphql.NullString
	LiquidityThreshold graphql.NullFloat
	Exchangepairs      *[]Exchangepair
}

type Exchangepair struct {
	Exchange           graphql.NullString
	Pairs              *[]graphql.NullString
	LiquidityThreshold graphql.NullFloat
}

func (r *DiaResolver) GetChart(ctx context.Context, args struct {
	Filter               graphql.NullString
	BlockDurationSeconds graphql.NullInt
	BlockShiftSeconds    graphql.NullInt
	Symbol               graphql.NullString
	StartTime            graphql.NullTime
	EndTime              graphql.NullTime
	Exchanges            *[]graphql.NullString
	Address              graphql.NullString
	BlockChain           graphql.NullString
	BaseAsset            *[]Asset
}) (*[]*FilterPointResolver, error) {
	fpr, _ := r.GetChartMeta(ctx, args)
	return fpr.fpr, nil
}

// TO DO: Use context?
func (r *DiaResolver) GetChartMeta(ctx context.Context, args struct {
	Filter               graphql.NullString
	BlockDurationSeconds graphql.NullInt
	BlockShiftSeconds    graphql.NullInt
	Symbol               graphql.NullString
	StartTime            graphql.NullTime
	EndTime              graphql.NullTime
	Exchanges            *[]graphql.NullString
	Address              graphql.NullString
	BlockChain           graphql.NullString
	BaseAsset            *[]Asset
}) (*FilterPointMetaResolver, error) {

	if containsSpecialChars(*args.Symbol.Value) || containsSpecialChars(*args.Address.Value) || containsSpecialChars(*args.BlockChain.Value) {
		return nil, errInvalidInputParams
	}
	var (
		blockShiftSeconds int64
		tradeBlocks       []queryhelper.Block
		blockchain        string
		address           string
		sr                FilterPointMetaResolver
		asset             dia.Asset
		err               error
		baseAssets        []dia.Asset
	)

	// Parsing input parameters.
	filter := args.Filter.Value
	if containsSpecialChars(*filter) {
		return nil, errInvalidInputParams
	}

	blockSizeSeconds := int64(*args.BlockDurationSeconds.Value)
	if args.BlockShiftSeconds.Value != nil {
		blockShiftSeconds = int64(*args.BlockShiftSeconds.Value)
	}
	if args.Address.Value != nil {
		if containsSpecialChars(*args.Address.Value) {
			return nil, errInvalidInputParams
		}
		address = *args.Address.Value

	}
	if args.BlockChain.Value != nil {
		if containsSpecialChars(*args.BlockChain.Value) {
			return nil, errInvalidInputParams
		}
		blockchain = *args.BlockChain.Value
	}
	symbol := *args.Symbol.Value
	starttime := args.StartTime.Value.Time
	endtime := args.EndTime.Value.Time
	starttimeimmutable := args.StartTime.Value.Time
	endtimeimmutable := args.EndTime.Value.Time

	if containsSpecialChars(symbol) {
		return nil, errInvalidInputParams
	}
	exchanges := args.Exchanges
	var exchangesString []string
	if exchanges != nil {
		for _, v := range *exchanges {
			if containsSpecialChars(*v.Value) {
				continue
			}
			exchangesString = append(exchangesString, *v.Value)
		}
	}

	// Fetch base assets.
	argsbaseasset := args.BaseAsset
	if argsbaseasset != nil {
		for _, baseasset := range *argsbaseasset {
			if containsSpecialChars(*baseasset.Address.Value) || containsSpecialChars(*baseasset.BlockChain.Value) {
				continue
			}
			asset, err = r.RelDB.GetAsset(*baseasset.Address.Value, *baseasset.BlockChain.Value)
			if err != nil {
				log.Errorf("Asset not found with address %s and blockchain %s ", address, blockchain)
				continue
			}
			baseAssets = append(baseAssets, asset)
		}
	}

	// Get quote asset either by blockchain&address or by topAsset method.
	if address != "" && blockchain != "" {
		asset, err = r.RelDB.GetAsset(address, blockchain)
		if err != nil {
			log.Errorf("Asset not found with address %s and blockchain %s ", address, blockchain)
			return &sr, err
		}
	} else {
		assets, err := r.RelDB.GetTopAssetByVolume(symbol)
		if err != nil {
			log.Errorf("Asset not found with symbol %s ", symbol)
			return &sr, err
		}

		log.Infoln("All assets having same symbol", assets)
		asset = assets[0]
	}

	log.Infoln("Asset Selected", asset)
	var (
		filterPoints, emaFilterPoints []dia.FilterPoint
		filterMetadata                *dia.FilterPointMetadata
	)

	if *filter != "ema" {

		if endtime.After(time.Now()) {
			endtime = time.Now()
		}

		// Limit the (time-)range of the query to 24 hours.
		maxStartTime := endtime.Add(-time.Duration(60*24) * time.Minute)
		if starttime.Before(maxStartTime) {
			starttime = maxStartTime
		}

		// Set blockShiftSeconds = blockSizeSeconds per default if not given.
		if blockShiftSeconds == 0 {
			blockShiftSeconds = blockSizeSeconds
		}

		// (Potentially) decrease starttime such that an integer number of bins fits into the whole range.
		if endtime.Unix()-starttime.Unix() < blockSizeSeconds {
			// Just one block.
			starttime = time.Unix(endtime.Unix()-blockSizeSeconds, 0)
		} else {
			starttime = time.Unix(starttime.Unix()-(blockShiftSeconds-((endtime.Unix()-starttime.Unix()-blockSizeSeconds)%blockShiftSeconds)), 0)
		}

		// Make time bins according to block size and block shift parameters.
		bins := utils.MakeBins(starttime, endtime, blockSizeSeconds, blockShiftSeconds)

		// Fetch trades.
		var trades []dia.Trade
		if blockShiftSeconds <= blockSizeSeconds {
			// Fetch all trades in time range.
			trades, err = r.DS.GetTradesByExchangesAndBaseAssets(asset, baseAssets, exchangesString, starttime, endtime, 0)
			if err != nil {
				log.Error("GetTradesByExchangesAndBaseAssets: ", err)
				return &sr, err
			}
		} else {
			// Fetch trades batched for disjoint bins.
			var starttimes, endtimes []time.Time
			for _, bin := range bins {
				starttimes = append(starttimes, bin.Starttime)
				endtimes = append(endtimes, bin.Endtime)
			}
			trades, err = r.DS.GetTradesByExchangesBatched(asset, baseAssets, exchangesString, starttimes, endtimes, 0)
			if err != nil {
				log.Error("GetTradesByExchangesAndBaseAssets: ", err)
				return &sr, err
			}
		}
		log.Println("Generating blocks, Total Trades", len(trades))
		log.Info("generating bins. Total bins: ", len(bins))

		if len(trades) > 0 && len(bins) > 0 {
			// In case the first bin is empty, look for the last trade before @starttime.
			if !utils.IsInBin(trades[0].Time, bins[0]) {
				previousTrade, baseAsseteErr := r.DS.GetTradesByExchangesAndBaseAssets(asset, baseAssets, exchangesString, endtime.AddDate(0, 0, -10), starttime, 1)
				if baseAsseteErr != nil || len(previousTrade) == 0 {
					log.Error("get initial trade: ", err, baseAsseteErr)
					// Fill with a zero trade so we can build blocks.
					auxTrade := trades[0]
					auxTrade.Volume = 0
					auxTrade.Price = 0
					auxTrade.EstimatedUSDPrice = 0
					trades = append([]dia.Trade{auxTrade}, trades...)
				} else {
					trades = append([]dia.Trade{previousTrade[0]}, trades...)
				}
			}
			tradeBlocks = queryhelper.NewBlockGenerator(trades).GenerateBlocks(blockSizeSeconds, blockShiftSeconds, bins)
			log.Println("Total TradeBlocks", len(tradeBlocks))
		}

	} else if *filter == "ema" {
		emaFilterPoints, err = r.DS.GetFilter("MA120", asset, "", starttimeimmutable, endtimeimmutable)
		if err != nil {
			log.Errorln("Error getting filter", err)
		}
	}

	switch *filter {
	case "ema":
		{
			filterPoints, filterMetadata = queryhelper.FilterEMA(emaFilterPoints, asset, int(blockSizeSeconds))
		}
	case "mair":
		{
			filterPoints, filterMetadata = queryhelper.FilterMAIR(tradeBlocks, asset, int(blockSizeSeconds))
		}
	case "ma":
		{
			filterPoints, filterMetadata = queryhelper.FilterMA(tradeBlocks, asset, int(blockSizeSeconds))
		}
	case "vwap":
		{
			filterPoints, filterMetadata = queryhelper.FilterVWAP(tradeBlocks, asset, int(blockSizeSeconds))
		}
	case "vwapir":
		{
			filterPoints, filterMetadata = queryhelper.FilterVWAPIR(tradeBlocks, asset, int(blockSizeSeconds))
		}
	case "medir":
		{
			filterPoints, filterMetadata = queryhelper.FilterMEDIR(tradeBlocks, asset, int(blockSizeSeconds))
		}
	case "vol":
		{
			filterPoints, filterMetadata = queryhelper.FilterVOL(tradeBlocks, asset, int(blockSizeSeconds))
		}

	}

	var fpr []*FilterPointResolver

	for _, fp := range filterPoints {
		fpr = append(fpr, &FilterPointResolver{q: fp})

	}

	return &FilterPointMetaResolver{fpr: &fpr, min: filterMetadata.Min, max: filterMetadata.Max}, nil
}

// GetxcFeed returns filter points for a (possibly) cross chain feed given by @QuoteAssets.
func (r *DiaResolver) GetxcFeed(ctx context.Context, args struct {
	Filter            graphql.NullString
	QuoteAssets       *[]Asset
	Exchanges         *[]graphql.NullString
	BlockSizeSeconds  graphql.NullInt
	BlockShiftSeconds graphql.NullInt
	StartTime         graphql.NullTime
	EndTime           graphql.NullTime
}) (*[]*FilterPointResolver, error) {

	// --- Parse input data ---
	var (
		vr           *[]*FilterPointResolver
		tradeBlocks  []queryhelper.Block
		filterPoints []dia.FilterPoint
	)

	filter := args.Filter.Value
	if containsSpecialChars(*filter) {
		return nil, errInvalidInputParams
	}

	var quoteAssets []dia.Asset
	if len(*args.QuoteAssets) > 0 {
		for i := range *args.QuoteAssets {
			quoteAssets = append(quoteAssets, dia.Asset{Address: *(*args.QuoteAssets)[i].Address.Value, Blockchain: *(*args.QuoteAssets)[i].BlockChain.Value})
		}
	}

	var exchanges []string
	if len(*args.Exchanges) > 0 {
		for i := range *args.Exchanges {
			if containsSpecialChars(*(*args.Exchanges)[i].Value) {
				continue
			}
			exchanges = append(exchanges, *(*args.Exchanges)[i].Value)
		}
	}

	blockSizeSeconds := int64(*args.BlockSizeSeconds.Value)
	var blockShiftSeconds int64
	if args.BlockShiftSeconds.Value != nil {
		blockShiftSeconds = int64(*args.BlockShiftSeconds.Value)
	} else {
		blockShiftSeconds = blockSizeSeconds
	}

	// --- Make time bins according to block size and block shift parameters ---
	starttime := args.StartTime.Value.Time
	endtime := args.EndTime.Value.Time
	if args.EndTime.Set {
		endtime = args.EndTime.Value.Time
	}
	if endtime.After(time.Now()) {
		endtime = time.Now()
	}

	// --- Limit the (time-)range of the query to 24 hours ---
	maxStartTime := endtime.Add(-time.Duration(60*24) * time.Minute)
	if starttime.Before(maxStartTime) {
		starttime = maxStartTime
	}

	// (Potentially) decrease starttime such that an integer number of bins fits into the whole range.
	if endtime.Unix()-starttime.Unix() < blockSizeSeconds {
		// Just one block.
		starttime = time.Unix(endtime.Unix()-blockSizeSeconds, 0)
	} else {
		starttime = time.Unix(starttime.Unix()-(blockShiftSeconds-((endtime.Unix()-starttime.Unix()-blockSizeSeconds)%blockShiftSeconds)), 0)
	}

	// --- Make time bins according to block size and block shift parameters ---
	bins := utils.MakeBins(starttime, endtime, blockSizeSeconds, blockShiftSeconds)
	var (
		starttimes []time.Time
		endtimes   []time.Time
	)
	if blockShiftSeconds <= blockSizeSeconds {
		// Bins overlap and hence all trades in total time range are needed.
		starttimes = []time.Time{starttime}
		endtimes = []time.Time{endtime}
	} else {
		// Bins are disjoint. Hence, only fetch trades per bin.
		for _, bin := range bins {
			starttimes = append(starttimes, bin.Starttime)
			endtimes = append(endtimes, bin.Endtime)
		}
	}

	// --- Fetch trades ---
	trades, err := r.DS.GetxcTradesByExchangesBatched(quoteAssets, exchanges, starttimes, endtimes)
	if err != nil {
		return vr, err
	}

	if len(trades) > 0 && len(bins) > 0 {
		// In case the first bin is empty, look for the last trade before @starttime.
		if !utils.IsInBin(trades[0].Time, bins[0]) {
			trades = r.fillFirstBin(quoteAssets, exchanges, trades, starttime.AddDate(0, 0, -10), starttime)
		}
		tradeBlocks = queryhelper.NewBlockGenerator(trades).GenerateBlocks(blockSizeSeconds, blockShiftSeconds, bins)
		log.Println("Total TradeBlocks", len(tradeBlocks))
	}

	switch *filter {
	case "mair":
		{
			filterPoints, _ = queryhelper.FilterMAIR(tradeBlocks, quoteAssets[0], int(blockSizeSeconds))
		}
	case "ma":
		{
			filterPoints, _ = queryhelper.FilterMA(tradeBlocks, quoteAssets[0], int(blockSizeSeconds))
		}
	case "vwap":
		{
			filterPoints, _ = queryhelper.FilterVWAP(tradeBlocks, quoteAssets[0], int(blockSizeSeconds))
		}
	case "vwapir":
		{
			filterPoints, _ = queryhelper.FilterVWAPIR(tradeBlocks, quoteAssets[0], int(blockSizeSeconds))
		}
	case "medir":
		{
			filterPoints, _ = queryhelper.FilterMEDIR(tradeBlocks, quoteAssets[0], int(blockSizeSeconds))
		}
	case "vol":
		{
			filterPoints, _ = queryhelper.FilterVOL(tradeBlocks, quoteAssets[0], int(blockSizeSeconds))
		}
	}

	var fpr []*FilterPointResolver
	for _, fp := range filterPoints {
		fpr = append(fpr, &FilterPointResolver{q: fp})
	}

	return &fpr, nil
}

// TO DO: Use context?
func (r *DiaResolver) GetFeed(ctx context.Context, args struct {
	Filter               graphql.NullString
	BlockSizeSeconds     graphql.NullInt
	BlockShiftSeconds    graphql.NullInt
	StartTime            graphql.NullTime
	EndTime              graphql.NullTime
	TradeVolumeThreshold graphql.NullFloat
	NativeDenomination   graphql.NullBool
	FeedSelection        *[]FeedSelection
}) (*[]*FilterPointExtendedResolver, error) {
	var (
		tradeBlocks          []queryhelper.Block
		sr                   *[]*FilterPointExtendedResolver
		starttime            time.Time
		endtime              time.Time
		blockShiftSeconds    int64
		tradeVolumeThreshold float64
		err                  error
		nativeDenomination   bool
	)

	// Parsing input parameters.
	if args.Filter.Value == nil {
		return sr, errors.New("filter must be set")
	}
	filter := *args.Filter.Value
	if containsSpecialChars(filter) {
		return nil, errInvalidInputParams
	}
	blockSizeSeconds := int64(*args.BlockSizeSeconds.Value)
	if args.BlockShiftSeconds.Value != nil {
		blockShiftSeconds = int64(*args.BlockShiftSeconds.Value)
	}
	// Set blockShiftSeconds = blockSizeSeconds per default if not given.
	if blockShiftSeconds == 0 {
		blockShiftSeconds = blockSizeSeconds
	}

	// Handle timestamps.
	if args.EndTime.Value != nil {
		endtime = args.EndTime.Value.Time
	} else {
		endtime = time.Now()
	}
	if args.StartTime.Value != nil {
		starttime = args.StartTime.Value.Time
	} else {
		starttime = endtime.Add(-time.Duration(1 * time.Hour))
	}
	if endtime.Before(starttime) {
		return sr, errors.New("startTime must be before EndTime")
	}
	if endtime.After(time.Now()) {
		endtime = time.Now()
	}
	// --- Limit the (time-)range of the query to 24 hours ---
	maxStartTime := endtime.Add(-time.Duration(60*24) * time.Minute)
	if starttime.Before(maxStartTime) {
		starttime = maxStartTime
	}
	// (Potentially) decrease starttime such that an integer number of bins fits into the whole range.
	if endtime.Unix()-starttime.Unix() < blockSizeSeconds {
		// Just one block.
		starttime = time.Unix(endtime.Unix()-blockSizeSeconds, 0)
	} else {
		starttime = time.Unix(starttime.Unix()-(blockShiftSeconds-((endtime.Unix()-starttime.Unix()-blockSizeSeconds)%blockShiftSeconds)), 0)
	}

	if args.TradeVolumeThreshold.Value != nil {
		tradeVolumeThreshold = *args.TradeVolumeThreshold.Value
	} else {
		tradeVolumeThreshold = TRADE_VOLUME_THRESHOLD_DEFAULT
	}

	// If nativeDenomination is true, price is returned in terms of the respective base asset.
	// Default is false, i.e. price is returned in USD denomination.
	if args.NativeDenomination.Value != nil {
		nativeDenomination = *args.NativeDenomination.Value
	}

	if args.FeedSelection == nil {
		return sr, errors.New("at least 1 asset must be selected")
	}
	feedselection, statusMessage, statusCode, err := r.castLocalFeedSelection(*args.FeedSelection)
	if err != nil {
		return sr, err
	}

	var (
		// filterPoints, emaFilterPoints []dia.FilterPoint
		filterPoints []dia.FilterPointExtended
	)

	// Make time bins according to block size and block shift parameters.
	bins := utils.MakeBins(starttime, endtime, blockSizeSeconds, blockShiftSeconds)

	// Fetch trades.
	var (
		trades     []dia.Trade
		starttimes []time.Time
		endtimes   []time.Time
	)

	if blockShiftSeconds <= blockSizeSeconds {
		// Fetch all trades in time range.
		starttimes = []time.Time{starttime}
		endtimes = []time.Time{endtime}
	} else {
		// Fetch trades batched for disjoint bins.
		for _, bin := range bins {
			starttimes = append(starttimes, bin.Starttime)
			endtimes = append(endtimes, bin.Endtime)
		}
	}
	trades, err = r.DS.GetTradesByFeedselectionRedis(feedselection, starttimes, endtimes, -1)
	if err != nil {
		return sr, err
	}
	log.Println("Generating blocks, Total Trades", len(trades))
	log.Info("generating bins. Total bins: ", len(bins))

	if len(bins) > 0 {
		// In case the first bin is empty, look for the last trades before @starttime
		// in order to select the most recent one with sufficient volume.
		if len(trades) == 0 || !utils.IsInBin(trades[0].Time, bins[0]) || trades[0].VolumeUSD() < tradeVolumeThreshold {
			previousTrade, err := r.DS.GetTradesByFeedselectionRedis(feedselection, []time.Time{endtime.AddDate(0, 0, -10)}, []time.Time{starttime}, int64(lookbackTradesNumber))
			if len(previousTrade) == 0 {
				log.Error("get initial trade: ", err)
				// Fill with a zero trade so we can build blocks.
				auxTrade := trades[0]
				auxTrade.Volume = 0
				auxTrade.Price = 0
				auxTrade.EstimatedUSDPrice = 0
				trades = append([]dia.Trade{auxTrade}, trades...)
			} else {
				for _, t := range previousTrade {
					if t.VolumeUSD() > tradeVolumeThreshold {
						trades = append([]dia.Trade{t}, trades...)
						break
					}
				}
			}
		}
		tradeBlocks = queryhelper.NewBlockGenerator(trades).GenerateBlocks(blockSizeSeconds, blockShiftSeconds, bins)
		log.Println("Total TradeBlocks", len(tradeBlocks))
	}

	switch filter {
	case "mair":
		{
			filterPoints = queryhelper.FilterMAIRextended(tradeBlocks, feedselection[0].Asset, int(blockSizeSeconds), tradeVolumeThreshold, nativeDenomination)
		}
	case "ma":
		{
			filterPoints = queryhelper.FilterMAextended(tradeBlocks, feedselection[0].Asset, int(blockSizeSeconds), tradeVolumeThreshold, nativeDenomination)
		}
	case "vwap":
		{
			filterPoints = queryhelper.FilterVWAPextended(tradeBlocks, feedselection[0].Asset, int(blockSizeSeconds), tradeVolumeThreshold, nativeDenomination)
		}
	case "vwapir":
		{
			filterPoints = queryhelper.FilterVWAPIRextended(tradeBlocks, feedselection[0].Asset, int(blockSizeSeconds), tradeVolumeThreshold, nativeDenomination)
		}
	case "medir":
		{
			filterPoints = queryhelper.FilterMEDIRextended(tradeBlocks, feedselection[0].Asset, int(blockSizeSeconds), tradeVolumeThreshold, nativeDenomination)
		}
	case "vol":
		{
			filterPoints = queryhelper.FilterVOLextended(tradeBlocks, feedselection[0].Asset, int(blockSizeSeconds))
		}

	}

	var fper []*FilterPointExtendedResolver

	for _, fpe := range filterPoints {
		fpe.StatusMessage = statusMessage
		fpe.StatusCode = statusCode
		fper = append(fper, &FilterPointExtendedResolver{q: fpe})
	}

	return &fper, nil
}

// GetFeedAggregation computes stats such as volume and number of trades fo a given FeedSelection.
func (r *DiaResolver) GetFeedAggregation(ctx context.Context, args struct {
	StartTime            graphql.NullTime
	EndTime              graphql.NullTime
	TradeVolumeThreshold graphql.NullFloat
	FeedSelection        *[]FeedSelection
}) (*[]*FeedSelectionAggregatedResolver, error) {
	var (
		sr                   []*FeedSelectionAggregatedResolver
		starttime            time.Time
		endtime              time.Time
		tradeVolumeThreshold float64
		err                  error
	)

	// Handle timestamps.
	if args.EndTime.Value != nil {
		endtime = args.EndTime.Value.Time
	} else {
		endtime = time.Now()
	}
	if args.StartTime.Value != nil {
		starttime = args.StartTime.Value.Time
	} else {
		starttime = endtime.Add(-time.Duration(1 * time.Hour))
	}
	if endtime.Before(starttime) {
		return &sr, errors.New("startTime must be before EndTime")
	}
	if endtime.After(time.Now()) {
		endtime = time.Now()
	}
	// --- Limit the (time-)range of the query to 24 hours ---
	maxStartTime := endtime.Add(-time.Duration(7*24*60) * time.Minute)
	if starttime.Before(maxStartTime) {
		starttime = maxStartTime
	}

	if args.TradeVolumeThreshold.Value != nil {
		tradeVolumeThreshold = *args.TradeVolumeThreshold.Value
	} else {
		tradeVolumeThreshold = TRADE_VOLUME_THRESHOLD_DEFAULT
	}

	if args.FeedSelection == nil {
		return &sr, errors.New("at least 1 asset must be selected")
	}
	feedselection, statusMessage, statusCode, err := r.castLocalFeedSelection(*args.FeedSelection)
	if err != nil {
		return &sr, err
	}

	// Get aggregated data in given time-range.
	fsa, err := r.DS.GetAggregatedFeedSelection(feedselection, starttime, endtime, tradeVolumeThreshold)
	if err != nil {
		log.Error("GetAggregatedFeedSelection: ", err)
		return &sr, err
	}

	// Fill response slice.
	var fsar []*FeedSelectionAggregatedResolver
	for _, fs := range fsa {
		fs.StatusMessage = statusMessage
		fs.StatusCode = statusCode
		if fs.Pooladdress != "" {
			pool, err := r.RelDB.GetPoolByAddress(fs.Basetoken.Blockchain, fs.Pooladdress)
			if err != nil {
				log.Error("GetPoolByAddress: ", err)
			}
			if pool.Time.After(time.Now().AddDate(0, 0, -7)) {
				fs.PoolLiquidityUSD, _ = pool.GetPoolLiquidityUSD()
			}
		}
		fsar = append(fsar, &FeedSelectionAggregatedResolver{q: fs})
	}

	return &fsar, nil
}

func (r *DiaResolver) GetVWALP(ctx context.Context, args struct {
	Quotetokenblockchain graphql.NullString
	Quotetokenaddress    graphql.NullString
	BaseAssets           *[]Asset
	Exchanges            *[]graphql.NullString
	BlockDurationSeconds graphql.NullInt
	EndTime              graphql.NullTime
	BasisPoints          graphql.NullInt
}) (*VWALPResolver, error) {

	// --- Parse input data ---
	var vr *VWALPResolver
	if containsSpecialChars(*args.Quotetokenaddress.Value) || containsSpecialChars(*args.Quotetokenblockchain.Value) {
		return nil, errInvalidInputParams
	}

	quoteAsset, err := r.RelDB.GetAsset(*args.Quotetokenaddress.Value, *args.Quotetokenblockchain.Value)
	if err != nil {
		log.Error("GetAsset: ", err)
	}

	var baseAssets []dia.Asset
	if len(*args.BaseAssets) > 0 {
		for i := range *args.BaseAssets {
			if containsSpecialChars(*(*args.BaseAssets)[i].Address.Value) || containsSpecialChars(*(*args.BaseAssets)[i].BlockChain.Value) {
				continue
			}
			baseAssets = append(baseAssets, dia.Asset{Address: *(*args.BaseAssets)[i].Address.Value, Blockchain: *(*args.BaseAssets)[i].BlockChain.Value})
		}
	}

	var exchanges []string
	if len(*args.Exchanges) > 0 {
		for i := range *args.Exchanges {
			if containsSpecialChars(*(*args.Exchanges)[i].Value) {
				continue
			}
			exchanges = append(exchanges, *(*args.Exchanges)[i].Value)
		}
	}

	BlockDurationSeconds := *args.BlockDurationSeconds.Value
	basisPoints := *args.BasisPoints.Value
	endtime := time.Now()
	if args.EndTime.Set {
		endtime = args.EndTime.Value.Time
	}
	//  -----------------------

	// Fetch trades from Influx.
	trades, err := r.DS.GetTradesByExchangesAndBaseAssets(
		quoteAsset,
		baseAssets,
		exchanges,
		endtime.Add(-time.Duration(BlockDurationSeconds)*time.Second),
		endtime,
		0,
	)
	if err != nil {
		return vr, err
	}

	tradesByExchange := make(map[string][]dia.Trade)
	for _, trade := range trades {
		tradesByExchange[trade.Source] = append(tradesByExchange[trade.Source], trade)
	}

	// Get last trades and volumes.
	var lastPrices []float64
	var volumes []float64
	for exchange := range tradesByExchange {
		block := queryhelper.Block{Trades: tradesByExchange[exchange], TimeStamp: endtime.UnixNano()}
		filterPoints, _ := queryhelper.FilterVOL([]queryhelper.Block{block}, quoteAsset, int(BlockDurationSeconds))
		if len(filterPoints) > 0 {
			lastPrices = append(lastPrices, filterPoints[0].LastTrade.EstimatedUSDPrice)
			volumes = append(volumes, filterPoints[0].Value)
		}
	}

	// Outlier detection.
	prices, volumes, _, err := utils.DiscardOutliers(lastPrices, volumes, float64(basisPoints))
	if err != nil {
		log.Error("DiscardOutliers: ", err)
	}

	// Build vwap.
	var vwap float64
	var volTotal float64
	for i := range prices {
		vwap += prices[i] * volumes[i]
		volTotal += volumes[i]
	}
	if volTotal > 0 {
		vwap /= volTotal
	}

	var response vwalp
	response.Symbol = quoteAsset.Symbol
	response.Value = vwap
	response.Time = endtime

	return &VWALPResolver{q: response}, nil
}

// GetNFT returns an NFT by address, blockchain and token_id.
func (r *DiaResolver) GetNFT(ctx context.Context, args struct {
	Address    graphql.NullString
	Blockchain graphql.NullString
	TokenID    graphql.NullString
}) (*NFTResolver, error) {

	if containsSpecialChars(*args.Address.Value) || containsSpecialChars(*args.Blockchain.Value) || containsSpecialChars(*args.TokenID.Value) {
		return nil, errInvalidInputParams
	}

	n, err := r.RelDB.GetNFT(*args.Address.Value, *args.Blockchain.Value, *args.TokenID.Value)
	if err != nil {
		return nil, err
	}

	return &NFTResolver{n: n}, nil
}

// GetNFTTrades returns trades of an NFT by address, blockchain, token_id and time range
func (r *DiaResolver) GetNFTTrades(ctx context.Context, args struct {
	Address    graphql.NullString
	Blockchain graphql.NullString
	TokenID    graphql.NullString
}) (*[]*NFTTradeResolver, error) {

	if containsSpecialChars(*args.Address.Value) || containsSpecialChars(*args.Blockchain.Value) || containsSpecialChars(*args.TokenID.Value) {
		return nil, errInvalidInputParams
	}

	var tr []*NFTTradeResolver
	trades, err := r.RelDB.GetNFTTrades(*args.Address.Value, *args.Blockchain.Value, *args.TokenID.Value, time.Time{}, time.Now())
	if err != nil {
		return nil, err
	}

	for _, trade := range trades {
		trade.NFT.NFTClass.Address = *args.Address.Value
		trade.NFT.NFTClass.Blockchain = *args.Blockchain.Value
		trade.NFT.TokenID = *args.TokenID.Value
		tr = append(tr, &NFTTradeResolver{trade: trade})
	}

	return &tr, nil
}

// GetNFTOffers returns offers of an NFT by address, blockchain, token_id and time range
func (r *DiaResolver) GetNFTOffers(ctx context.Context, args struct {
	Address    graphql.NullString
	Blockchain graphql.NullString
	TokenID    graphql.NullString
}) (*[]*NFTOfferResolver, error) {

	if containsSpecialChars(*args.Address.Value) || containsSpecialChars(*args.Blockchain.Value) || containsSpecialChars(*args.TokenID.Value) {
		return nil, errInvalidInputParams
	}

	var or []*NFTOfferResolver
	offers, err := r.RelDB.GetNFTOffers(*args.Address.Value, *args.Blockchain.Value, *args.TokenID.Value)
	if err != nil {
		return nil, err
	}

	for _, offer := range offers {
		offer.NFT.NFTClass.Address = *args.Address.Value
		offer.NFT.NFTClass.Blockchain = *args.Blockchain.Value
		offer.NFT.TokenID = *args.TokenID.Value
		or = append(or, &NFTOfferResolver{offer: offer})
	}

	return &or, nil
}

// GetNFTOffers returns offers of an NFT by address, blockchain, token_id and time range
func (r *DiaResolver) GetNFTBids(ctx context.Context, args struct {
	Address    graphql.NullString
	Blockchain graphql.NullString
	TokenID    graphql.NullString
}) (*[]*NFTBidResolver, error) {

	if containsSpecialChars(*args.Address.Value) || containsSpecialChars(*args.Blockchain.Value) || containsSpecialChars(*args.TokenID.Value) {
		return nil, errInvalidInputParams
	}

	var br []*NFTBidResolver
	bids, err := r.RelDB.GetNFTBids(*args.Address.Value, *args.Blockchain.Value, *args.TokenID.Value)
	if err != nil {
		return nil, err
	}

	for _, bid := range bids {
		bid.NFT.NFTClass.Address = *args.Address.Value
		bid.NFT.NFTClass.Blockchain = *args.Blockchain.Value
		bid.NFT.TokenID = *args.TokenID.Value
		br = append(br, &NFTBidResolver{bid: bid})
	}

	return &br, nil
}

func (r *DiaResolver) fillFirstBin(assets []dia.Asset, exchanges []string, trades []dia.Trade, starttime time.Time, endtime time.Time) []dia.Trade {
	previousTrade, quoteAssetsErr := r.DS.GetxcTradesByExchangesBatched(assets, exchanges, []time.Time{starttime}, []time.Time{endtime})
	if quoteAssetsErr != nil || len(previousTrade) == 0 {
		log.Error("get initial trade: ", quoteAssetsErr)
		// Fill with a zero trade so we can build blocks.
		auxTrade := trades[0]
		auxTrade.Volume = 0
		auxTrade.Price = 0
		auxTrade.EstimatedUSDPrice = 0
		trades = append([]dia.Trade{auxTrade}, trades...)
	} else {
		trades = append([]dia.Trade{previousTrade[len(previousTrade)-1]}, trades...)
		log.Warn("previous trade: ", previousTrade[len(previousTrade)-1])
	}
	return trades
}

// castLocalFeedSelection casts the input selection into a @[]dia.FeedSelection type
// so that we can call the corresponding influx trades getter.
// It also checks for admissibility of the underlying input data.
func (r *DiaResolver) castLocalFeedSelection(fs []FeedSelection) ([]dia.FeedSelection, string, int32, error) {
	var dfs []dia.FeedSelection
	var warning string
	var warnings []string
	for _, localFeedSelection := range fs {
		var diaFeedSelection dia.FeedSelection

		// Parse asset.
		if localFeedSelection.Address.Value == nil || localFeedSelection.Blockchain.Value == nil {
			err := errors.New("missing asset's address and blockchain")
			return dfs, warning, 1, err
		}
		diaFeedSelection.Asset = dia.Asset{
			Address:    normalizeAddress(*localFeedSelection.Address.Value, *localFeedSelection.Blockchain.Value),
			Blockchain: *localFeedSelection.Blockchain.Value,
		}

		// Parse liquidity threshold.
		if localFeedSelection.LiquidityThreshold.Value != nil {
			diaFeedSelection.LiquidityThreshold = *localFeedSelection.LiquidityThreshold.Value
		}

		// Check for exchange input. Continue if neither of exchange and liquidty threshold are set.
		if localFeedSelection.Exchangepairs == nil && diaFeedSelection.LiquidityThreshold == 0 {
			dfs = append(dfs, diaFeedSelection)
			continue
		}

		// Fetch admissible pools in case liquidity threshold is set.
		if diaFeedSelection.LiquidityThreshold > 0 {
			pools, errPools := r.RelDB.GetPoolsByAsset(diaFeedSelection.Asset, 0, diaFeedSelection.LiquidityThreshold)
			if errPools != nil {
				return dfs, warning, 1, errPools
			}
			for _, pool := range pools {

				// In order to append pool to corresponding entry of diaFeedSelection first check
				// if the exchange was already added.
				var poolAdded bool
				for i := range diaFeedSelection.Exchangepairs {
					if diaFeedSelection.Exchangepairs[i].Exchange == pool.Exchange {
						diaFeedSelection.Exchangepairs[i].Pools = append(diaFeedSelection.Exchangepairs[i].Pools, pool)
						poolAdded = true
						break
					}
				}
				if !poolAdded {
					var eps dia.ExchangepairSelection
					eps.Exchange = pool.Exchange
					eps.Pools = append(eps.Pools, pool)
					diaFeedSelection.Exchangepairs = append(diaFeedSelection.Exchangepairs, eps)
				}

			}
			if len(pools) > 0 {
				dfs = append(dfs, diaFeedSelection)
			}
			// Continue, i.e. ignore (possibly) given exchangepairs.
			continue
		}

		// Parse exchanges.
		for _, ep := range *localFeedSelection.Exchangepairs {
			var diaExchangepairselection dia.ExchangepairSelection
			if ep.Exchange.Value == nil {
				err := errors.New("missing exchange name")
				return dfs, warning, statusCodesMap[statusError], err
			}

			// Check if exchange is valid.
			var exchangeOk bool
			var exchange dia.Exchange
			for e := range EXCHANGES {
				if strings.EqualFold(*ep.Exchange.Value, e) {
					exchangeOk = true
					exchange = EXCHANGES[e]
					break
				}
			}
			if !exchangeOk && *ep.Exchange.Value != "" {
				warnings = append(warnings, fmt.Sprintf("exchange name not valid: %s. ", *ep.Exchange.Value))
				continue
			}
			diaExchangepairselection.Exchange = exchange

			// Select all pairs on given exchange if not specified.
			if ep.Pairs == nil {
				// Check if asset is traded as quotetoken on exchange and emit warning if not.
				eps, err := r.RelDB.GetExchangepairsByAsset(
					dia.Asset{
						Address:    diaFeedSelection.Asset.Address,
						Blockchain: diaFeedSelection.Asset.Blockchain,
					},
					exchange.Name,
					false,
				)
				log.Infof("Number of pairs on exchange %s: %v", exchange.Name, len(eps))
				for _, e := range eps {
					log.Info("ForeignName: ", e.ForeignName)
				}
				if err != nil {
					log.Errorf("GetExchangepairsByAsset for %s on %s: %v", diaFeedSelection.Asset.Address, diaFeedSelection.Asset.Blockchain, err)
				}
				if len(eps) == 0 {
					warnings = append(warnings, fmt.Sprintf("%s-%s: asset not available on %s. ", diaFeedSelection.Asset.Blockchain, diaFeedSelection.Asset.Address, exchange.Name))
				}

				diaFeedSelection.Exchangepairs = append(diaFeedSelection.Exchangepairs, diaExchangepairselection)
				continue
			}

			// Check whether pairs or pools are given.
			var pairs bool
			var pairsCount int
			for _, p := range *ep.Pairs {
				if p.Value == nil {
					return dfs, warning, statusCodesMap[statusError], errors.New("missing pair identifier")
				}
				if len(strings.Split(*p.Value, PAIR_SEPARATOR)) == 2 {
					pairs = true
					pairsCount++
				}
			}
			if !(pairsCount == len(*ep.Pairs) || pairsCount == 0) {
				return dfs, warning, statusCodesMap[statusError], errors.New("pair/pool notation not consistent")
			}

			if exchange.Name != "" && exchange.Centralized {
				if !pairs {
					return dfs, warning, statusCodesMap[statusError], errors.New("wrong notation for pairs")
				}
				for _, p := range *ep.Pairs {
					if len(strings.Split(*p.Value, PAIR_SEPARATOR)) < 2 {
						return dfs, warning, statusCodesMap[statusError], errors.New("pair not in requested format TokenA-TokenB")
					}

					exchangepair, err := r.RelDB.GetExchangePair(exchange.Name, *p.Value, false)
					if err != nil {
						warnings = append(warnings, fmt.Sprintf("%s: pair not available on %s. ", *p.Value, exchange.Name))
					}
					pair := dia.Pair{
						QuoteToken: exchangepair.UnderlyingPair.QuoteToken,
						BaseToken:  exchangepair.UnderlyingPair.BaseToken,
					}
					diaExchangepairselection.Pairs = append(diaExchangepairselection.Pairs, pair)
				}
			} else if exchange.Name == "" {
				diaExchangepairselection.Exchange.Centralized = true
				for _, p := range *ep.Pairs {

					// Check admissibility of input.
					pairString := strings.Split(*p.Value, LIST_SEPARATOR)
					if len(pairString) < 2 {
						return dfs, warning, statusCodesMap[statusError], errors.New("exactly 2 assets must be given for each pair")
					}
					var pair dia.Pair

					if len(strings.Split(pairString[0], PAIR_SEPARATOR)) < 2 {
						return dfs, warning, statusCodesMap[statusError], errors.New("asset not in requested format Blockchain-Address")
					}
					blockchainQuotetoken := strings.Title(strings.ToLower(strings.Split(pairString[0], PAIR_SEPARATOR)[0]))
					pair.QuoteToken.Address = normalizeAddress(strings.Split(pairString[0], PAIR_SEPARATOR)[1], blockchainQuotetoken)
					pair.QuoteToken.Blockchain = blockchainQuotetoken

					blockchainBasetoken := strings.Title(strings.ToLower(strings.Split(pairString[1], PAIR_SEPARATOR)[0]))
					pair.BaseToken.Address = normalizeAddress(strings.Split(pairString[1], PAIR_SEPARATOR)[1], blockchainBasetoken)
					pair.BaseToken.Blockchain = blockchainBasetoken

					diaExchangepairselection.Pairs = append(diaExchangepairselection.Pairs, pair)
				}
			} else {
				for _, p := range *ep.Pairs {
					pool, err := r.RelDB.GetPoolByAddress(
						diaFeedSelection.Asset.Blockchain,
						normalizeAddress(*p.Value, diaFeedSelection.Asset.Blockchain),
					)
					if err != nil {
						return dfs, warning, statusCodesMap[statusError], err
					}
					diaExchangepairselection.Pools = append(diaExchangepairselection.Pools, pool)
				}
			}
			diaFeedSelection.Exchangepairs = append(diaFeedSelection.Exchangepairs, diaExchangepairselection)
		}

		dfs = append(dfs, diaFeedSelection)
	}
	uniqueWarnings := ""
	for _, w := range utils.UniqueStrings(warnings) {
		uniqueWarnings += w
	}
	return dfs, uniqueWarnings, statusCodes(uniqueWarnings), nil
}

// Returns the EIP55 compliant address in case @blockchain has an Ethereum ChainID.
func makeAddressEIP55Compliant(address string, blockchain string) string {
	if strings.Contains(BLOCKCHAINS[blockchain].ChainID, "Ethereum") {
		return common.HexToAddress(address).Hex()
	}
	return address
}

// Normalize address depending on the blockchain.
func normalizeAddress(address string, blockchain string) string {
	if strings.Contains(BLOCKCHAINS[blockchain].ChainID, "Ethereum") {
		return makeAddressEIP55Compliant(address, blockchain)
	}
	if BLOCKCHAINS[blockchain].Name == dia.OSMOSIS {
		if strings.Contains(address, "ibc-") && len(strings.Split(address, "-")[1]) > 1 {
			return "ibc/" + strings.Split(address, "-")[1]
		}
	}
	return address
}

func containsSpecialChars(s string) bool {
	return strings.ContainsAny(s, "!@#$%^&*()'\"|{}[];><?/`~,")
}

func statusCodes(statusMessage string) int32 {
	var statusCode int32
	if strings.Contains(statusMessage, statusError) {
		return statusCodesMap[statusError]
	}
	for status := range statusCodesMap {
		if strings.Contains(statusMessage, status) {
			statusCode += statusCodesMap[status]
		}
	}
	return statusCode
}
