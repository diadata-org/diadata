package queryhelper

import (
	"time"

	filters "github.com/diadata-org/diadata/internal/pkg/filtersBlockService"
	"github.com/diadata-org/diadata/pkg/dia"
)

func FilterMAextended(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPointsExtended []dia.FilterPointExtended) {

	lastfp := &dia.FilterPoint{}

	for _, block := range tradeBlocks {

		fpe := &dia.FilterPointExtended{}
		pairs, pools := block.GetBlockSources()
		fpe.Pairs = pairs
		fpe.Pools = pools
		fpe.TradesCount = int32(len(block.Trades))

		if len(block.Trades) > 0 {
			maFilter := filters.NewFilterMA(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)

			for _, trade := range block.Trades {
				maFilter.Compute(trade)
			}

			maFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := maFilter.FilterPointForBlock()

			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}

			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *fp
				filterPointsExtended = append(filterPointsExtended, *fpe)
				lastfp = fp
			} else if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *lastfp
				filterPointsExtended = append(filterPointsExtended, *fpe)
			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			fpe.FilterPoint = *lastfp
			filterPointsExtended = append(filterPointsExtended, *fpe)
		}
	}

	return
}

func FilterMAIRextended(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPointsExtended []dia.FilterPointExtended) {
	lastfp := &dia.FilterPoint{}

	for i, block := range tradeBlocks {
		log.Infof("block number: %v", i)
		fpe := &dia.FilterPointExtended{}
		pairs, pools := block.GetBlockSources()
		fpe.Pairs = pairs
		fpe.Pools = pools
		fpe.TradesCount = int32(len(block.Trades))

		if len(block.Trades) > 0 {
			mairFilter := filters.NewFilterMAIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)
			firstBlock := block.Trades[0]
			for _, trade := range block.Trades {
				mairFilter.Compute(trade)
			}

			mairFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := mairFilter.FilterPointForBlock()

			fp.FirstTrade = firstBlock
			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}

			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *fp
				filterPointsExtended = append(filterPointsExtended, *fpe)
				lastfp = fp
			} else if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *lastfp
				filterPointsExtended = append(filterPointsExtended, *fpe)
			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			fpe.FilterPoint = *lastfp
			filterPointsExtended = append(filterPointsExtended, *fpe)
		}
	}

	return
}

func FilterVWAPextended(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPointsExtended []dia.FilterPointExtended) {
	var lastfp *dia.FilterPoint

	for _, block := range tradeBlocks {
		fpe := &dia.FilterPointExtended{}
		pairs, pools := block.GetBlockSources()
		fpe.Pairs = pairs
		fpe.Pools = pools
		fpe.TradesCount = int32(len(block.Trades))

		if len(block.Trades) > 0 {

			vwapFilter := filters.NewFilterVWAP(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)
			for _, trade := range block.Trades {
				vwapFilter.Compute(trade)
			}

			vwapFilter.FinalCompute(block.Trades[0].Time)
			fp := vwapFilter.FilterPointForBlock()

			fp.FirstTrade = block.Trades[0]

			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = fp.FirstTrade
			}

			fp.Time = time.Unix(block.TimeStamp/1e9, 0)
			fpe.FilterPoint = *fp
			filterPointsExtended = append(filterPointsExtended, *fpe)
			lastfp = fp

		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			fpe.FilterPoint = *lastfp
			filterPointsExtended = append(filterPointsExtended, *fpe)
		}
	}
	return
}

func FilterVWAPIRextended(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPointsExtended []dia.FilterPointExtended) {
	var lastfp *dia.FilterPoint

	for _, block := range tradeBlocks {

		fpe := &dia.FilterPointExtended{}
		pairs, pools := block.GetBlockSources()
		fpe.Pairs = pairs
		fpe.Pools = pools
		fpe.TradesCount = int32(len(block.Trades))

		if len(block.Trades) > 0 {

			vwapirFilter := filters.NewFilterVWAPIR(asset, "", time.Unix(block.TimeStamp/1e9, 0), blockSize)
			for _, trade := range block.Trades {
				vwapirFilter.Compute(trade)
			}
			vwapirFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := vwapirFilter.FilterPointForBlock()

			fp.FirstTrade = block.Trades[0]

			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = fp.FirstTrade
			}

			if fp != nil && fp.Value > 0 {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *fp
				filterPointsExtended = append(filterPointsExtended, *fpe)
				lastfp = fp
			} else if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *lastfp
				filterPointsExtended = append(filterPointsExtended, *fpe)
			}
		} else {
			if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *lastfp
				filterPointsExtended = append(filterPointsExtended, *fpe)
			}
		}
	}
	return
}

func FilterMEDIRextended(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPointsExtended []dia.FilterPointExtended) {
	var lastfp *dia.FilterPoint

	for _, block := range tradeBlocks {

		fpe := &dia.FilterPointExtended{}
		pairs, pools := block.GetBlockSources()
		fpe.Pairs = pairs
		fpe.Pools = pools
		fpe.TradesCount = int32(len(block.Trades))

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

			fp.FirstTrade = block.Trades[0]

			if len(block.Trades) > 0 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}
			if fp != nil && fp.Value > 0 {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *fp
				filterPointsExtended = append(filterPointsExtended, *fpe)
				lastfp = fp
			} else if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *lastfp
				filterPointsExtended = append(filterPointsExtended, *fpe)
			}
		} else {
			if lastfp != nil {
				lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *lastfp
				filterPointsExtended = append(filterPointsExtended, *fpe)
			}
		}
	}
	return
}

func FilterEMAextended(points []dia.FilterPoint, asset dia.Asset, blockSize int) (filterPoints []dia.FilterPoint, metadata *dia.FilterPointMetadata) {
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

func FilterVOLextended(tradeBlocks []Block, asset dia.Asset, blockSize int) (filterPointsExtended []dia.FilterPointExtended) {

	lastfp := &dia.FilterPoint{}

	for _, block := range tradeBlocks {

		fpe := &dia.FilterPointExtended{}
		pairs, pools := block.GetBlockSources()
		fpe.Pairs = pairs
		fpe.Pools = pools
		fpe.TradesCount = int32(len(block.Trades))

		if len(block.Trades) > 0 {

			volFilter := filters.NewFilterVOL(asset, "", blockSize)
			for _, trade := range block.Trades {
				volFilter.Compute(trade)
			}
			volFilter.FinalCompute(time.Unix(block.TimeStamp/1e9, 0))
			fp := volFilter.FilterPointForBlock()

			fp.FirstTrade = block.Trades[0]
			if len(block.Trades) > 1 {
				fp.LastTrade = block.Trades[len(block.Trades)-1]
			} else {
				fp.LastTrade = block.Trades[0]
			}
			if fp != nil {
				fp.Time = time.Unix(block.TimeStamp/1e9, 0)
				fpe.FilterPoint = *fp
				filterPointsExtended = append(filterPointsExtended, *fpe)
				lastfp = fp
			}
		} else {
			lastfp.Time = time.Unix(block.TimeStamp/1e9, 0)
			fpe.FilterPoint = *lastfp
			filterPointsExtended = append(filterPointsExtended, *fpe)
		}
	}
	return
}
