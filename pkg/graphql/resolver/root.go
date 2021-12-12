package resolver

import (
	"context"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	queryhelper "github.com/diadata-org/diadata/pkg/dia/helpers/queryHelper"

	models "github.com/diadata-org/diadata/pkg/model"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Resolver is the root resolver
type DiaResolver struct {
	DS    models.DB
	RelDB models.RelDB
}

// GetQuotation Get quotation
func (r *DiaResolver) GetQuotation(ctx context.Context, args struct{ Symbol graphql.NullString }) (*QuotationResolver, error) {
	q, err := r.DS.GetQuotation(*args.Symbol.Value)
	if err != nil {
		return nil, err
	}
	return &QuotationResolver{q: *q}, nil
}

func (r *DiaResolver) GetSupply(ctx context.Context, args struct{ Symbol graphql.NullString }) (*SupplyResolver, error) {
	q, err := r.DS.GetLatestSupply(*args.Symbol.Value)
	if err != nil {
		return nil, err
	}
	return &SupplyResolver{q: q}, nil
}

func (r *DiaResolver) GetSupplies(ctx context.Context, args struct{ Symbol graphql.NullString }) (*[]*SupplyResolver, error) {
	starttime := time.Unix(1, 0)
	endtime := time.Now()
	q, err := r.DS.GetSupply(*args.Symbol.Value, starttime, endtime)
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
}) (*[]*FilterPointResolver, error) {
	var (
		blockShiftSeconds int64
		tradeBlocks       []queryhelper.Block
		blockchain        string
		address           string
		sr                []*FilterPointResolver
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
	exchanges := args.Exchanges
	var exchangesString []string
	if exchanges != nil {
		for _, v := range *exchanges {
			exchangesString = append(exchangesString, *v.Value)
		}
	}

	var (
		asset dia.Asset
		err   error
	)

	if address != "" && blockchain != "" {
		asset, err = r.RelDB.GetAsset(address, blockchain)
		if err != nil {
			log.Errorln("Asset not found with address %s and blockchain %s ", address, blockchain)
			return &sr, err
		}

	} else {
		assets, err := r.RelDB.GetTopAssetByVolume(symbol)
		if err != nil {
			log.Errorln("Asset not found with symbol %s ", symbol)
			return &sr, err
		}

		log.Println("All assets", assets)
		asset = assets[0]

	}

	log.Println("asset", asset)

	if blockShiftSeconds == 0 {
		if endtime.After(time.Now()) {
			endtime = time.Now()
		}
		maxStartTime := endtime.Add(time.Duration(-(blockSizeSeconds * 1000)) * time.Second)

		if starttime.Before(maxStartTime) {
			starttime = maxStartTime
		}

		trades, err := r.DS.GetTradesByExchanges(asset, exchangesString, starttime, endtime)
		if err != nil {
			return &sr, err
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

		trades, err := r.DS.GetTradesByExchanges(asset, exchangesString, starttime, endtime)
		if err != nil {
			return &sr, err
		}
		log.Println("Generating blocks, Total Trades", len(trades))
		tradeBlocks = queryhelper.NewBlockGenerator(trades).GenerateShift(trades[0].Time.UnixNano(), blockSizeSeconds, blockShiftSeconds)
		log.Println("Total TradeBlocks", len(tradeBlocks))

	}

	var filterPoints []dia.FilterPoint

	switch *filter {

	case "mair":
		{
			filterPoints = queryhelper.FilterMAIR(tradeBlocks, asset, int(blockSizeSeconds))
		}
	case "ma":
		{

			filterPoints = queryhelper.FilterMA(tradeBlocks, asset, int(blockSizeSeconds))

		}
	case "vwap":
		{
			filterPoints = queryhelper.FilterVWAP(tradeBlocks, asset, int(blockSizeSeconds))
		}
	case "vwapir":
		{
			filterPoints = queryhelper.FilterVWAPIR(tradeBlocks, asset, int(blockSizeSeconds))
		}
	}

	for _, fp := range filterPoints {
		sr = append(sr, &FilterPointResolver{q: fp})

	}

	// log.Println("Filter point", fp)
	// log.Println("Start Time", trades[len(trades)-1].Time)
	// log.Println("End Time", trades[0].Time)

	return &sr, nil
}
