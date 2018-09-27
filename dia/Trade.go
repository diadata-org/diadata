package dia

import (
	"strings"
)

func (t *Trade) SecondPair() string {

	if len(t.Pair) == 6 {
		return t.Pair[len(t.Pair)-3:]
	}

	if t.Pair[len(t.Pair)-4:] == "USDT" {
		return "USDT"
	}

	if t.Pair[len(t.Pair)-3:] == "EUR" {
		return "EUR"
	}

	if t.Pair[len(t.Pair)-3:] == "USD" {
		return "USD"
	}

	if t.Pair[len(t.Pair)-3:] == "ETH" {
		return "ETH"
	}

	if t.Pair[len(t.Pair)-3:] == "BTC" {
		return "BTC"
	}

	second := strings.TrimPrefix(t.Pair, t.Symbol+"_")
	if second != t.Pair {
		return second
	}

	return strings.TrimPrefix(t.Pair, t.Symbol)
}
