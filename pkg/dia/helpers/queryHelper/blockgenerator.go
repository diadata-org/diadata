package queryhelper

import (
	"github.com/diadata-org/diadata/pkg/dia"
	scrapers "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/utils"
)

var (
	EXCHANGES = scrapers.Exchanges
)

const (
	PAIR_SPLITTER = "-"
)

type Block struct {
	Trades    []dia.Trade
	TimeStamp int64
}

type Blockgenerator struct {
	trades []dia.Trade
}

// GetBlockSources returns a unique list of all pairs and pools contained in the block's @b trades.
func (b *Block) GetBlockSources() (pairs []dia.ExchangePair, pools []dia.Pool) {

	pairsCheckMap := make(map[string]struct{})
	poolsCheckMap := make(map[string]struct{})

	for _, t := range b.Trades {
		if EXCHANGES[t.Source].Centralized {
			err := t.NormalizeSymbols(true, PAIR_SPLITTER)
			if err != nil {
				log.Errorf("NormalizeSymbols on pair %s and exchange %s: %v", t.Pair, t.Source, err)
			}
			if _, ok := pairsCheckMap[t.Source+t.Pair]; !ok {
				pairs = append(pairs, dia.ExchangePair{Exchange: t.Source, ForeignName: t.Pair})
				pairsCheckMap[t.Source+t.Pair] = struct{}{}
			}
		} else {
			if _, ok := poolsCheckMap[t.Source+t.PoolAddress]; !ok {
				pools = append(pools, dia.Pool{Exchange: EXCHANGES[t.Source], Address: t.PoolAddress})
				poolsCheckMap[t.Source+t.PoolAddress] = struct{}{}
			}
		}
	}
	return
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
