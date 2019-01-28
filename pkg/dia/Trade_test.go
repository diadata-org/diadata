package dia

import (
	"testing"
)

func TestTrade(t *testing.T) {

	trade := &Trade{Pair: "BTCUSD", Source: KrakenExchange}
	r := trade.SecondPair()
	if r != "USD" {
		t.Errorf("error secondpair %v", r)
	}

	trade2 := &Trade{Symbol: "BTC", Pair: "BTC_USDT", Source: KrakenExchange}
	r = trade2.SecondPair()
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}

	trade3 := &Trade{Symbol: "BTC", Pair: "BTCUSDT", Source: KrakenExchange}
	r = trade3.SecondPair()
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}

	trade4 := &Trade{Symbol: "ZB", Pair: "zbusdt", Source: KrakenExchange}
	r = trade4.SecondPair()
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}
}

func TestBitfinexAndHitBTCSecondPairAreUSDT(t *testing.T) {
	trade := &Trade{Pair: "BTCUSD", Source: BitfinexExchange}
	r := trade.SecondPair()
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}
	trade2 := &Trade{Symbol: "BTC", Pair: "BTCUSD", Source: HitBTCExchange}
	r = trade2.SecondPair()
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}
}
