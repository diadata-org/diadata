package dia

import (
	"testing"
)

func TestTrade(t *testing.T) {

	trade := &Trade{Pair: "BTCUSD", Source: KrakenExchange}
	r := trade.GetBaseToken()
	if r != "USD" {
		t.Errorf("error base token %v", r)
	}

	trade2 := &Trade{Symbol: "BTC", Pair: "BTC_USDT", Source: KrakenExchange}
	r = trade2.GetBaseToken()
	if r != "USDT" {
		t.Errorf("error base token %v", r)
	}

	trade3 := &Trade{Symbol: "BTC", Pair: "BTCUSDT", Source: KrakenExchange}
	r = trade3.GetBaseToken()
	if r != "USDT" {
		t.Errorf("error base token %v", r)
	}

	trade4 := &Trade{Symbol: "ZB", Pair: "zbusdt", Source: KrakenExchange}
	r = trade4.GetBaseToken()
	if r != "USDT" {
		t.Errorf("error base token %v", r)
	}
}

func TestBitfinexAndHitBTCSecondPairAreUSDT(t *testing.T) {
	trade := &Trade{Pair: "BTCUSD", Source: BitfinexExchange}
	r := trade.GetBaseToken()
	if r != "USDT" {
		t.Errorf("error base token %v", r)
	}
	trade2 := &Trade{Symbol: "BTC", Pair: "BTCUSD", Source: HitBTCExchange}
	r = trade2.GetBaseToken()
	if r != "USDT" {
		t.Errorf("error base token %v", r)
	}
}
