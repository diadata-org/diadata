package queryhelper

import (
	"encoding/json"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia"
)

func setupGenerator() []Block {
	var trades []dia.Trade
	json.Unmarshal([]byte(jsonTrades), &trades)
	bg := NewBlockGenerator(trades, 8)
	blocks := bg.generate()

	return blocks
}
func TestBlockGeneratorFirstAndLastTimeDiff(t *testing.T) {
	blocks := setupGenerator()
	for k := range blocks {
		diff := blocks[k].Trades[len(blocks[k].Trades)-1].Time.UnixNano() - blocks[k].Trades[0].Time.UnixNano()
		if diff > 8*1e9 {
			t.Error("diff is less than 8", diff)
		}

	}
}

func TestIfAnytradeisrepeadted(t *testing.T) {
	blocks := setupGenerator()
	tradeToCheck := blocks[0].Trades[0]
	for count, block := range blocks {
		if count == 0 {
			continue
		}
		for _, trade := range block.Trades {
			if trade.ForeignTradeID == tradeToCheck.ForeignTradeID {
				t.Error("Duplicate Trade")
			}

		}
	}
}
