package queryhelper

import (
	"github.com/diadata-org/diadata/pkg/dia"
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

func (bg *Blockgenerator) GenerateSize(blockSizeSeconds int64) (tradeBlocks []Block) {
	firstBlockStartTime := bg.trades[0].Time.UnixNano()
	currentBlockEndTime := firstBlockStartTime + (blockSizeSeconds * 1e9)
	tradeBlock := Block{
		TimeStamp: firstBlockStartTime,
	}

	for count, trade := range bg.trades {
		if trade.Time.UnixNano() >= firstBlockStartTime {
			if trade.Time.UnixNano() > currentBlockEndTime {
				currentBlockEndTime = trade.Time.UnixNano() + (blockSizeSeconds * 1e9)
				tradeBlocks = append(tradeBlocks, tradeBlock)
				tradeBlock = Block{
					TimeStamp: currentBlockEndTime - (blockSizeSeconds * 1e9),
				}
			}
			tradeBlock.Trades = append(tradeBlock.Trades, trade)

			// IF last block is not complete but trades are finished then add rest trades in block

			if count == len(bg.trades)-1 {
				tradeBlocks = append(tradeBlocks, tradeBlock)

			}

		} else {
			log.Infoln("Trade is out of initial block time Trdae time", trade.Time.UnixNano(), firstBlockStartTime)
		}

	}
	return
}

func (bg *Blockgenerator) GenerateShift(firstBlockStartTime, blockSizeSeconds, blockShiftSeconds int64) (tradeBlocks []Block) {
	var tradeBlock Block

	nextBlockStarttime := firstBlockStartTime
	tradeBlock.TimeStamp = firstBlockStartTime

	if len(bg.trades) == 1 {
		tradeBlock = Block{Trades: []dia.Trade{bg.trades[0]}, TimeStamp: nextBlockStarttime}
		tradeBlocks = append(tradeBlocks, tradeBlock)
		return
	}

	for _, trade := range bg.trades {
		if trade.Time.UnixNano() >= firstBlockStartTime {
			if trade.Time.UnixNano() > nextBlockStarttime {
				tradeBlocks = append(tradeBlocks, tradeBlock)
				lastTrades := removeTradesBlock(tradeBlock, int(blockShiftSeconds))
				nextBlockStarttime = nextBlockStarttime + (blockShiftSeconds * 1e9)
				tradeBlock = Block{Trades: lastTrades, TimeStamp: nextBlockStarttime}
			} else {
				tradeBlock.Trades = append(tradeBlock.Trades, trade)
			}

		} else {
			log.Infoln("Trade is out of initial block time Trade time", trade.Time.UnixNano(), firstBlockStartTime)
		}

	}
	return
}

func removeTradesBlock(tradeBlock Block, blockshift int) (trades []dia.Trade) {
	if len(tradeBlock.Trades) > 0 {
		var startTime = tradeBlock.Trades[0].Time
		for _, trade := range tradeBlock.Trades {
			if trade.Time.UnixNano()-startTime.UnixNano() >= int64(blockshift*1e9) {
				trades = append(trades, trade)
			}
		}

	}

	return

}
