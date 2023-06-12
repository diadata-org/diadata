package resolver

import (
	"context"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	queryhelper "github.com/diadata-org/diadata/pkg/dia/helpers/queryHelper"
	"github.com/diadata-org/diadata/pkg/utils"

	models "github.com/diadata-org/diadata/pkg/model"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Resolver is the root resolver
type DiaResolver struct {
	DS              models.DB
	RelDB           models.RelDB
	InfluxBatchSize int64
}

func (r *DiaResolver) GetSupply(ctx context.Context, args struct{ Symbol graphql.NullString }) (*SupplyResolver, error) {
	q, err := r.DS.GetLatestSupply(*args.Symbol.Value, &r.RelDB)
	if err != nil {
		return nil, err
	}
	return &SupplyResolver{q: q}, nil
}

func (r *DiaResolver) GetSupplies(ctx context.Context, args struct{ Symbol graphql.NullString }) (*[]*SupplyResolver, error) {
	starttime := time.Unix(1, 0)
	endtime := time.Now()
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

// func (r *DiaResolver) GetSymbols(ctx context.Context, args struct{ Exchange graphql.NullString }) (*[]*string, error) {
// 	exchange := args.Exchange.Value
// 	var allSymbols []string

// 	if *exchange == "" {
// 		allSymbols = r.DS.GetAllSymbols()
// 		if len(allSymbols) == 0 {
// 			return nil, errors.New("error No symbols")
// 		}
// 	} else {
// 		allSymbols = r.DS.GetSymbolsByExchange(*exchange)
// 		if len(allSymbols) == 0 {
// 			return nil, errors.New("error No Symbols for exchange " + *exchange)
// 		}
// 	}
// 	var sr []*string

// 	for _, symbol := range allSymbols {
// 		sr = append(sr, &symbol)
// 	}
// 	return &sr, nil
// }

type TradeBlock struct {
	Trades []dia.Trade
}

type AssetInput struct {
	Address    graphql.NullString
	BlockChain graphql.NullString
}

type BaseAsset struct {
	Address    string
	BlockChain string
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
	BaseAsset            *[]AssetInput
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
	BaseAsset            *[]AssetInput
}) (*FilterPointMetaResolver, error) {
	var (
		blockShiftSeconds int64
		tradeBlocks       []queryhelper.Block
		blockchain        string
		address           string
		sr                *FilterPointMetaResolver
		asset             dia.Asset
		err               error
		baseAssets        []dia.Asset
	)

	// Parsing input parameters.
	filter := args.Filter.Value
	blockSizeSeconds := int64(*args.BlockDurationSeconds.Value)
	if args.BlockShiftSeconds.Value != nil {
		blockShiftSeconds = int64(*args.BlockShiftSeconds.Value)
	}
	if args.Address.Value != nil {
		address = *args.Address.Value
	}
	if args.BlockChain.Value != nil {
		blockchain = *args.BlockChain.Value
	}
	symbol := *args.Symbol.Value
	starttime := args.StartTime.Value.Time
	endtime := args.EndTime.Value.Time
	starttimeimmutable := args.StartTime.Value.Time
	endtimeimmutable := args.EndTime.Value.Time
	exchanges := args.Exchanges
	var exchangesString []string
	if exchanges != nil {
		for _, v := range *exchanges {
			exchangesString = append(exchangesString, *v.Value)
		}
	}

	// Fetch base assets.
	argsbaseasset := args.BaseAsset
	if argsbaseasset != nil {
		for _, baseasset := range *argsbaseasset {
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
			return sr, err
		}
	} else {
		assets, err := r.RelDB.GetTopAssetByVolume(symbol)
		if err != nil {
			log.Errorf("Asset not found with symbol %s ", symbol)
			return sr, err
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
				return sr, err
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
				return sr, err
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
	QuoteAssets       *[]AssetInput
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

	var quoteAssets []dia.Asset
	if len(*args.QuoteAssets) > 0 {
		for i := range *args.QuoteAssets {
			quoteAssets = append(quoteAssets, dia.Asset{Address: *(*args.QuoteAssets)[i].Address.Value, Blockchain: *(*args.QuoteAssets)[i].BlockChain.Value})
		}
	}

	var exchanges []string
	if len(*args.Exchanges) > 0 {
		for i := range *args.Exchanges {
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

func (r *DiaResolver) GetVWALP(ctx context.Context, args struct {
	Quotetokenblockchain graphql.NullString
	Quotetokenaddress    graphql.NullString
	BaseAssets           *[]AssetInput
	Exchanges            *[]graphql.NullString
	BlockDurationSeconds graphql.NullInt
	EndTime              graphql.NullTime
	BasisPoints          graphql.NullInt
}) (*VWALPResolver, error) {

	// --- Parse input data ---
	var vr *VWALPResolver

	quoteAsset, err := r.RelDB.GetAsset(*args.Quotetokenaddress.Value, *args.Quotetokenblockchain.Value)
	if err != nil {
		log.Error("GetAsset: ", err)
	}

	var baseAssets []dia.Asset
	if len(*args.BaseAssets) > 0 {
		for i := range *args.BaseAssets {
			baseAssets = append(baseAssets, dia.Asset{Address: *(*args.BaseAssets)[i].Address.Value, Blockchain: *(*args.BaseAssets)[i].BlockChain.Value})
		}
	}

	var exchanges []string
	if len(*args.Exchanges) > 0 {
		for i := range *args.Exchanges {
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
