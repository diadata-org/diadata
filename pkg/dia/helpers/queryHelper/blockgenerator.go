package queryhelper

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
)

type Block struct {
	Trades    []dia.Trade
	TimeStamp int64
}

type Blockgenerator struct {
	trades []dia.Trade
}

func NewBlockGenerator(trades []dia.Trade) *Blockgenerator {
	return &Blockgenerator{trades: trades}
}

func (bg *Blockgenerator) GenerateBlocks(blockSizeSeconds int64, blockShiftSeconds int64, bins []utils.TimeBin) (tradeBlocks []Block) {

	if len(bg.trades) == 0 {
		return
	}

	tradeIndex := 0
	for _, bin := range bins {

		tradeBlock := Block{TimeStamp: bin.Endtime.UnixNano()}
		// Fill bin as long as trade's timestamp is smaller than or equal to the right border of the bin.
		for tradeIndex < len(bg.trades) && (bg.trades[tradeIndex].Time.Before(bin.Endtime) || bg.trades[tradeIndex].Time == bin.Endtime) {
			tradeBlock.Trades = append(tradeBlock.Trades, bg.trades[tradeIndex])
			tradeIndex++
		}

		tradeBlocks = append(tradeBlocks, tradeBlock)
	}

	return

}
