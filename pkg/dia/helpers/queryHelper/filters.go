package queryhelper

import (
	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
)

func FilterMA(tradeBlocks []Block, symbol string) (filterPoints []dia.FilterPoint) {
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

func FilterMAIR(tradeBlocks []Block, symbol string) (filterPoints []dia.FilterPoint) {
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
