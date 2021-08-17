package queryhelper

import (
	"encoding/json"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia"
)

var trades []dia.Trade

func setupGenerator() []Block {
	json.Unmarshal([]byte(jsonTrades), &trades)

	// trades = trades[9:]
	bg := NewBlockGenerator(trades)
	blocks := bg.GenerateSize(8)

	return blocks
}

func setupGeneratorSizeShift() []Block {
	json.Unmarshal([]byte(jsonTrades), &trades)

	// trades = trades[9:]
	bg := NewBlockGenerator(trades)
	blocks := bg.GenerateShift(8, 7)

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

func TestFirstTrade(t *testing.T) {
	blocks := setupGenerator()
	tradeToCheck := blocks[0].Trades[0]
	if trades[0].ForeignTradeID != tradeToCheck.ForeignTradeID {
		t.Error("Firts Trade is not equa to block first trade")
	}
}

func TestIfTradesAreMissed(t *testing.T) {
	blocks := setupGenerator()
	totalTradesProvided := len(trades)
	tradesinblock := 0

	for _, block := range blocks {
		tradesinblock = tradesinblock + len(block.Trades)
	}

	if totalTradesProvided != tradesinblock {
		t.Errorf("Trades count Mimatched total %d in blocks its %d", totalTradesProvided, tradesinblock)
	}

}
