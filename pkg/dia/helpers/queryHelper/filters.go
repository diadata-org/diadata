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
			maFilter := filters.NewFilterVWAP(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

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
			maFilter := filters.NewFilterVWAPIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {

				maFilter.Compute(trade)
			}

			maFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
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

func FilterMEDIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	var lastfp *dia.FilterPoint

	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			medirFilter := filters.NewFilterMEDIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {

				medirFilter.Compute(trade)
			}

			medirFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := medirFilter.FilterPointForBlock()
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

func FilterEMA(points []dia.FilterPoint, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint) {
	for index, point := range points {
		emaFilter := filters.NewFilterEMA(asset, "", time.Unix(point.Time.UnixNano()/1e9, 10), blockSize)
		if index%5 == 0 {
			emaFilter.FinalCompute(time.Unix(point.Time.UnixNano()/1e9, 0))
			fp := emaFilter.FilterPointForBlock()
			if fp != nil {
				filterPoints = append(filterPoints, *fp)
				log.Println("append index%5  %v  points %v filterPoints %v", index%5, point.Value, fp.Value)

			}
			log.Println("index%5  %v  points %v filterPoints %v", index%5, point.Value, fp.Value)

		} else {
			log.Println("Compute index%5  %v  points %v ", index%5, point.Value)
			emaFilter.Compute(point)
		}
	}
	return filterPoints
}
