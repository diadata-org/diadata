package queryhelper

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/prometheus/common/log"
)

type Block struct {
	Trades []dia.Trade
}

type Blockgenerator struct {
	trades           []dia.Trade
	blockSizeSeconds int64
}

func NewBlockGenerator(trades []dia.Trade, blockSizeSeconds int64) *Blockgenerator {
	return &Blockgenerator{trades: trades, blockSizeSeconds: blockSizeSeconds}
}

func (bg *Blockgenerator) Generate() (tradeBlocks []Block) {
	var tradeBlock Block

	firstBlockStartTime := bg.trades[0].Time.UnixNano()
	currentBlockStartTime := firstBlockStartTime + (bg.blockSizeSeconds * 1e9)

	for _, trade := range bg.trades {
		if trade.Time.UnixNano() >= firstBlockStartTime {
			if trade.Time.UnixNano() > currentBlockStartTime {
				currentBlockStartTime = trade.Time.UnixNano() + (bg.blockSizeSeconds * 1e9)
				tradeBlocks = append(tradeBlocks, tradeBlock)
				tradeBlock = Block{}
			} else {
				tradeBlock.Trades = append(tradeBlock.Trades, trade)
			}

		} else {
			log.Infoln("Trade is out of initial block time Trdae time", trade.Time.UnixNano(), firstBlockStartTime)
		}

	}
	return
}
