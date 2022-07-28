package queryhelper

import (
	"encoding/json"
	"fmt"
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
	err := json.Unmarshal([]byte(jsonTrades), &trades)

	fmt.Printf("t: %v\n", err)

	// trades = trades[9:]
	bg := NewBlockGenerator(trades)
	fmt.Printf("t: %v\n", trades)

	blocks := bg.GenerateShift(trades[0].Time.UnixNano(), 120, 240)

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

func TestVwap(t *testing.T) {

	blocks := setupGeneratorSizeShift()
	fmt.Printf("t: %v\n", blocks)

	r, _ := FilterVWAPIR(blocks, dia.Asset{Symbol: "DIA"}, 120)
	for _, v := range r {
		fmt.Printf("%v\n", v)
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
