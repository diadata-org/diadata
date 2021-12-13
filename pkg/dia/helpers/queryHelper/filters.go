package queryhelper

import (
	"time"

	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func FilterMA(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	var lastfp *dia.FilterPoint
	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			maFilter := filters.NewFilterMA(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {

				maFilter.Compute(trade)

			}

			maFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := maFilter.FilterPointForBlock()

			if fp != nil {
				filterPoints = append(filterPoints, *fp)
				lastfp = fp
			} else if lastfp != nil {
				log.Println("block.TimeStamp", block.TimeStamp)
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *lastfp)

			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *lastfp)

		}
	}
	return filterPoints
}

func FilterMAIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	var lastfp *dia.FilterPoint
	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			mairFilter := filters.NewFilterMAIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {
				mairFilter.Compute(trade)
			}

			mairFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := mairFilter.FilterPointForBlock()
			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *fp)
				lastfp = fp
			} else if lastfp != nil {
				log.Println("block.TimeStamp", block.TimeStamp)
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *lastfp)
			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *lastfp)
		}
	}
	return filterPoints
}

func FilterVWAP(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	var lastfp *dia.FilterPoint
	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			maFilter := filters.NewFilterVWAP(asset, "", block.Trades[len(block.Trades)-1].Time, blockSize)

			for _, trade := range block.Trades {

				maFilter.Compute(trade)
			}

			maFilter.FinalCompute(block.Trades[0].Time)
			fp := maFilter.FilterPointForBlock()
			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *fp)
				lastfp = fp
			} else {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *lastfp)
			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *lastfp)
		}
	}
	return filterPoints
}

func FilterVWAPIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	var lastfp *dia.FilterPoint

	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			maFilter := filters.NewFilterVWAPIR(asset, "Binance", block.Trades[len(block.Trades)-1].Time, blockSize)

			for _, trade := range block.Trades {

				maFilter.Compute(trade)
			}

			maFilter.FinalCompute(block.Trades[0].Time)
			fp := maFilter.FilterPointForBlock()
			if fp != nil && fp.Value > 0 {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *fp)
				lastfp = fp
			} else {
				if lastfp != nil {
					lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
					filterPoints = append(filterPoints, *lastfp)
				}
			}
		} else {
			if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *lastfp)
			}
		}
	}
	return filterPoints
}
