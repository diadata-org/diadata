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

type BaseAssetInput struct {
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
	BaseAsset            *[]BaseAssetInput
}) (*[]*FilterPointResolver, error) {
	fpr, _ := r.GetChartMeta(ctx, args)

	return fpr.fpr, nil
}

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
	BaseAsset            *[]BaseAssetInput
}) (*FilterPointMetaResolver, error) {
	var (
		blockShiftSeconds int64
		tradeBlocks       []queryhelper.Block
		blockchain        string
		address           string
		sr                *FilterPointMetaResolver
	)
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
	var (
		asset      dia.Asset
		err        error
		baseAssets []dia.Asset
	)

	argsbaseasset := args.BaseAsset

	if argsbaseasset != nil {
		for _, baseasset := range *argsbaseasset {

			asset, err = r.RelDB.GetAsset(*baseasset.Address.Value, *baseasset.BlockChain.Value)
			if err != nil {
				log.Errorln("Asset not found with address %s and blockchain %s ", address, blockchain)
				continue
			}

			baseAssets = append(baseAssets, asset)
		}
	}

	log.Errorln("baseAssets", baseAssets)

	if address != "" && blockchain != "" {
		asset, err = r.RelDB.GetAsset(address, blockchain)
		if err != nil {
			log.Errorln("Asset not found with address %s and blockchain %s ", address, blockchain)
			return sr, err
		}

	} else {
		assets, err := r.RelDB.GetTopAssetByVolume(symbol)
		if err != nil {
			log.Errorln("Asset not found with symbol %s ", symbol)
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

		if blockShiftSeconds == 0 {
			if endtime.After(time.Now()) {
				endtime = time.Now()
			}
			maxStartTime := endtime.Add(time.Duration(-(blockSizeSeconds * 1000)) * time.Second)

			if starttime.Before(maxStartTime) {
				starttime = maxStartTime
			}

			trades, err := r.DS.GetTradesByExchangesAndBaseAssets(asset, baseAssets, exchangesString, starttime, endtime)
			if err != nil {
				return sr, err
			}
			tradeBlocks = queryhelper.NewBlockGenerator(trades).GenerateSize(blockSizeSeconds)
		} else {
			if endtime.After(time.Now()) {
				endtime = time.Now()
			}
			totalTimeDiff := (blockSizeSeconds + (blockShiftSeconds * 999)) * int64(time.Second)
			maxStartTime := endtime.Add(-time.Duration(totalTimeDiff))

			if starttime.Before(maxStartTime) {
				starttime = maxStartTime
			}

			var trades []dia.Trade
			if blockShiftSeconds <= blockSizeSeconds {
				trades, err = r.DS.GetTradesByExchangesAndBaseAssets(asset, baseAssets, exchangesString, starttime, endtime)
				if err != nil {
					log.Error("GetTradesByExchangesAndBaseAssets: ", err)
					return sr, err
				}
			} else {
				// In this case, fetch trades by time window and batch the Influx API requests.

				timeInit := starttime
				timeFinal := endtime

				// Make time windows for which trades from influx have to be fetched.
				var startTimes, endTimes []time.Time
				for timeFinal.After(timeInit) {
					startTimes = append(startTimes, timeInit)
					endTimes = append(endTimes, timeInit.Add(time.Duration(blockSizeSeconds*1e9)))
					timeInit = timeInit.Add(time.Duration(blockShiftSeconds * 1e9))
				}

				// Determine number of batches.
				var numBatches int
				batchSize := int(r.InfluxBatchSize)
				if len(startTimes) < int(r.InfluxBatchSize) {
					numBatches = 1
					batchSize = len(startTimes)
				} else {
					numBatches = len(startTimes) / int(r.InfluxBatchSize)
				}

				// Iterate over batches.
				for i := 0; i < numBatches; i++ {
					var tradesBatch []dia.Trade
					var err error
					lowerIndex := i * batchSize
					upperIndex := (i + 1) * batchSize
					tradesBatch, err = r.DS.GetTradesByExchangesBatched(asset, baseAssets, exchangesString, startTimes[lowerIndex:upperIndex], endTimes[lowerIndex:upperIndex])
					if err != nil {
						log.Error("fetch trades batch from influx: ", err)
					}
					trades = append(trades, tradesBatch...)
				}
				// Add remaining trades if existent.
				if len(startTimes)%int(r.InfluxBatchSize) > 0 {
					var tradesBatch []dia.Trade
					var err error
					lowerIndex := numBatches * (batchSize)
					upperIndex := len(startTimes)
					tradesBatch, err = r.DS.GetTradesByExchangesBatched(asset, baseAssets, exchangesString, startTimes[lowerIndex:upperIndex], endTimes[lowerIndex:upperIndex])
					if err != nil {
						log.Error("fetch trades batch from influx: ", err)
					}
					trades = append(trades, tradesBatch...)
				}

			}
			log.Println("Generating blocks, Total Trades", len(trades))
			tradeBlocks = queryhelper.NewBlockGenerator(trades).GenerateShift(trades[0].Time.UnixNano(), blockSizeSeconds, blockShiftSeconds)
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

func (r *DiaResolver) GetVWALP(ctx context.Context, args struct {
	Quotetokenblockchain graphql.NullString
	Quotetokenaddress    graphql.NullString
	BaseAssets           *[]BaseAssetInput
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
	trades, err := r.DS.GetTradesByExchanges(
		quoteAsset,
		baseAssets,
		exchanges,
		endtime.Add(-time.Duration(BlockDurationSeconds)*time.Second),
		endtime,
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
