package resolver

import (
	"context"
	"errors"
	"time"

	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/sirupsen/logrus"
)

// Resolver is the root resolver
type DiaResolver struct {
	DS models.DB
}

var log = logrus.New()

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

func (r *DiaResolver) GetSymbols(ctx context.Context, args struct{ Exchange graphql.NullString }) (*[]*string, error) {
	exchange := args.Exchange.Value
	var allSymbols []string

	if *exchange == "" {
		allSymbols = r.DS.GetAllSymbols()
		if len(allSymbols) == 0 {
			return nil, errors.New("error No symbols")
		}
	} else {
		allSymbols = r.DS.GetSymbolsByExchange(*exchange)
		if len(allSymbols) == 0 {
			return nil, errors.New("error No Symbols for exchange " + *exchange)
		}
	}
	var sr []*string

	for _, symbol := range allSymbols {
		sr = append(sr, &symbol)
	}
	return &sr, nil
}

type TradeBlock struct {
	Trades []dia.Trade
}

func (r *DiaResolver) GetChart(ctx context.Context, args struct {
	Filter               graphql.NullString
	BlockDurationSeconds graphql.NullInt
	Symbol               graphql.NullString
	StartTime            graphql.NullTime
	EndTime              graphql.NullTime
	Exchanges            *[]graphql.NullString
}) (*[]*FilterPointResolver, error) {
	filter := args.Filter.Value
	blockSizeSeconds := int64(*args.BlockDurationSeconds.Value)
	symbol := string(*args.Symbol.Value)
	starttime := args.StartTime.Value.Time
	endtime := args.EndTime.Value.Time
	exchanges := args.Exchanges
	var exchangesString []string
	for _, v := range *exchanges {
		exchangesString = append(exchangesString, *v.Value)
	}

	trades, err := r.DS.GetTradesByExchange(symbol, exchangesString, starttime, endtime, 0)
	if err != nil {
		return nil, nil
	}

	var tradeBlocks []TradeBlock
	var tradeBlock TradeBlock

	firstBlockStartTime := trades[0].Time.UnixNano()
	currentBlockStartTime := firstBlockStartTime + (blockSizeSeconds * 1e9)

	for _, trade := range trades {

		if trade.Time.UnixNano() >= firstBlockStartTime {
			if trade.Time.UnixNano() > currentBlockStartTime {
				currentBlockStartTime = trade.Time.UnixNano() + (blockSizeSeconds * 1e9)
				tradeBlocks = append(tradeBlocks, tradeBlock)
				tradeBlock = TradeBlock{}
			} else {
				tradeBlock.Trades = append(tradeBlock.Trades, trade)
			}

		} else {
			log.Infoln("Trade is out of initial block time Trdae time", trade.Time.UnixNano(), firstBlockStartTime)
		}

	}
	var filterPoints []dia.FilterPoint
	switch *filter {

	case "mair":
		{
			filterPoints = filterMAIR(tradeBlocks, symbol)
		}

	case "ma":
		{
			filterPoints = filterMA(tradeBlocks, symbol)
		}
	}

	var sr []*FilterPointResolver

	for _, fp := range filterPoints {
		sr = append(sr, &FilterPointResolver{q: fp})

	}

	// log.Println("Filter point", fp)
	// log.Println("Start Time", trades[len(trades)-1].Time)
	// log.Println("End Time", trades[0].Time)

	return &sr, nil
}

func filterMA(tradeBlocks []TradeBlock, symbol string) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterMA(symbol, "Binance", block.Trades[len(block.Trades)-1].Time, dia.BlockSizeSeconds)

		for _, trade := range block.Trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}

func filterMAIR(tradeBlocks []TradeBlock, symbol string) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterMAIR(symbol, "Binance", block.Trades[len(block.Trades)-1].Time, dia.BlockSizeSeconds)

		for _, trade := range block.Trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}
