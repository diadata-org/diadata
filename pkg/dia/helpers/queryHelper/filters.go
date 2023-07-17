package queryhelper

import (
	"time"

	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func FilterMA(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {

	lastfp := &dia.FilterPoint{}
	metadata = dia.NewFilterPointMetadata()
	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			maFilter := filters.NewFilterMA(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {
				maFilter.Compute(trade)
			}

			maFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := maFilter.FilterPointForBlock()

			metadata.AddPoint(fp.Value)

			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}

			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *fp)
				lastfp = fp
			} else if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *lastfp)

			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *lastfp)
		}
	}

	return
}

func FilterMAIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {
	lastfp := &dia.FilterPoint{}
	metadata = dia.NewFilterPointMetadata()

	for i, block := range tradeBlocks {
		log.Infof("block number: %v", i)

		if len(block.Trades) > 0 {
			mairFilter := filters.NewFilterMAIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)
			firstBlock := block.Trades[0]
			for _, trade := range block.Trades {
				mairFilter.Compute(trade)
			}

			mairFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := mairFilter.FilterPointForBlock()

			metadata.AddPoint(fp.Value)

			fp.FirstTrade = firstBlock
			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}

			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *fp)
				lastfp = fp
			} else if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *lastfp)
			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *lastfp)
		}
	}

	return
}

func FilterVWAP(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {
	var lastfp *dia.FilterPoint
	metadata = dia.NewFilterPointMetadata()
	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			vwapFilter := filters.NewFilterVWAP(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {
				vwapFilter.Compute(trade)
			}

			vwapFilter.FinalCompute(block.Trades[0].Time)
			fp := vwapFilter.FilterPointForBlock()
			metadata.AddPoint(fp.Value)
			fp.FirstTrade = block.Trades[0]

			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = fp.FirstTrade
			}

			fp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *fp)
			lastfp = fp

		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *lastfp)
		}
	}
	return
}

func FilterVWAPIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {
	var lastfp *dia.FilterPoint
	metadata = dia.NewFilterPointMetadata()

	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {

			vwapirFilter := filters.NewFilterVWAPIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)
			for _, trade := range block.Trades {
				vwapirFilter.Compute(trade)
			}
			vwapirFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := vwapirFilter.FilterPointForBlock()

			metadata.AddPoint(fp.Value)

			fp.FirstTrade = block.Trades[0]

			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = fp.FirstTrade
			}
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
	return
}

func FilterMEDIR(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {
	var lastfp *dia.FilterPoint
	metadata = dia.NewFilterPointMetadata()

	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			medirFilter := filters.NewFilterMEDIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {

				medirFilter.Compute(trade)
			}

			medirFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := medirFilter.FilterPointForBlock()
			if fp == nil {
				log.Error("failed getting FilterPointForBlock")
				return
			}
			metadata.AddPoint(fp.Value)
			fp.FirstTrade = block.Trades[0]
			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}
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
	return
}

func FilterEMA(points []dia.FilterPoint, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {
	emaFilter := filters.NewFilterEMA(asset, "", points[0].Time, blockSize)
	metadata = dia.NewFilterPointMetadata()

	for index, point := range points {
		if index%5 == 0 {
			emaFilter.FinalCompute(point.Time)
			fp := emaFilter.FilterPointForBlock()
			metadata.AddPoint(fp.Value)
			if fp.Value > 0 {
				filterPoints = append(filterPoints, *fp)
				log.Println("append index%5  %v  points %v filterPoints %v filterPoints size %v", index%5, point.Value, fp.Value, len(filterPoints))

			}
			log.Println("index%5  %v  points %v filterPoints %v", index%5, point.Value, fp.Value)

		} else {
			log.Println("Compute index%5  %v  points %v ", index%5, point.Value)
			emaFilter.Compute(point)
		}
	}

	return
}

func FilterVOL(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {
	metadata = dia.NewFilterPointMetadata()

	lastFp := &dia.FilterPoint{}
	for _, block := range tradeBlocks {
		if len(block.Trades) > 0 {
			volFilter := filters.NewFilterVOL(asset, "", blockSize)

			for _, trade := range block.Trades {
				volFilter.Compute(trade)
			}

			volFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := volFilter.FilterPointForBlock()
			metadata.AddPoint(fp.Value)
			fp.FirstTrade = block.Trades[0]
			if len(block.Trades) > 1 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}
			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				filterPoints = append(filterPoints, *fp)
				lastFp = fp
			}
		} else {
			lastFp.Value = 0
			lastFp.Time = time.Unix(block.TimeStamp/1e9, 0)
			filterPoints = append(filterPoints, *lastFp)
		}
	}
	return
}
