package dia

import (
	"testing"
)

func TestTrade(t *testing.T) {

	trade := &Trade{Pair: "BTCUSD"}
	r := trade.SecondPair(KrakenExchange)
	if r != "USD" {
		t.Errorf("error secondpair %v", r)
	}

	trade2 := &Trade{Symbol: "BTC", Pair: "BTC_USDT"}
	r = trade2.SecondPair(KrakenExchange)
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}

	trade3 := &Trade{Symbol: "BTC", Pair: "BTCUSDT"}
	r = trade3.SecondPair(KrakenExchange)
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}

	trade4 := &Trade{Symbol: "ZB", Pair: "zbusdt"}
	r = trade4.SecondPair(KrakenExchange)
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}
}

func TestBitfinexAndHitBTCSecondPairAreUSDT(t *testing.T) {
	trade := &Trade{Pair: "BTCUSD"}
	r := trade.SecondPair(BitfinexExchange)
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}
	trade2 := &Trade{Symbol: "BTC", Pair: "BTCUSD"}
	r = trade2.SecondPair(HitBTCExchange)
	if r != "USDT" {
		t.Errorf("error secondpair %v", r)
	}
}
