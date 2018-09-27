package dia

import (
	"testing"
)

func TestTrade(t *testing.T) {

	trade := &Trade{Pair: "BTCUSD"}
	if trade.SecondPair() != "USD" {
		t.Errorf("error secondpair %v", trade.SecondPair())
	}

	trade2 := &Trade{Symbol: "BTC", Pair: "BTC_USDT"}
	if trade2.SecondPair() != "USDT" {
		t.Errorf("error secondpair %v", trade.SecondPair())
	}

	trade3 := &Trade{Symbol: "BTC", Pair: "BTCUSDT"}
	if trade3.SecondPair() != "USDT" {
		t.Errorf("error secondpair %v", trade.SecondPair())
	}
}
