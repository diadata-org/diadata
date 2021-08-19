package queryhelper

import (
	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"

	"github.com/sirupsen/logrus"
)

var logruslog = logrus.New()

func FilterMA(tradeBlocks []Block, symbol string) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterMA(dia.Asset{Symbol: "symbol"}, "Binance", block.Trades[len(block.Trades)-1].Time, dia.BlockSizeSeconds)

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

func FilterMAIR(tradeBlocks []Block, symbol string) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterMAIR(dia.Asset{Symbol: "symbol"}, "Binance", block.Trades[len(block.Trades)-1].Time, dia.BlockSizeSeconds)

		for _, trade := range block.Trades {
			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}

func FilterVWAP(tradeBlocks []Block, symbol string) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterVWAP(dia.Asset{Symbol: "symbol"}, "Binance", block.Trades[len(block.Trades)-1].Time, dia.BlockSizeSeconds)

		for _, trade := range block.Trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}

func FilterVWAPIR(tradeBlocks []Block, symbol string) (filterPoints []dia.FilterPoint) {
	for _, block := range tradeBlocks {
		maFilter := filters.NewFilterVWAPIR(dia.Asset{Symbol: "symbol"}, "Binance", block.Trades[len(block.Trades)-1].Time, dia.BlockSizeSeconds)

		for _, trade := range block.Trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(block.Trades[0].Time)
		fp := maFilter.FilterPointForBlock()
		filterPoints = append(filterPoints, *fp)
	}
	return filterPoints
}
