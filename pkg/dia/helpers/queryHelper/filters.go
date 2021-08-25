package queryhelper

import (
	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
)

func FilterMA(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterMA(asset, "", block.Trades[len(block.Trades)-1].Time, blockSize)

		for _, trade := range block.Trades {

			maFilter.Compute(trade)

		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()

		if fp != nil {
			filterPoints = append(filterPoints, *fp)
		}
	}
	return filterPoints
}

func FilterMAIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterMAIR(asset, "", block.Trades[len(block.Trades)-1].Time, blockSize)

		for _, trade := range block.Trades {
			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}

func FilterVWAP(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterVWAP(asset, "", block.Trades[len(block.Trades)-1].Time, blockSize)

		for _, trade := range block.Trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}

func FilterVWAPIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterVWAPIR(asset, "Binance", block.Trades[len(block.Trades)-1].Time, blockSize)

		for _, trade := range block.Trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}
