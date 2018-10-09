package dia

import (
	"strings"
)

func (t *Trade) SecondPair() string {

	pair := strings.ToUpper(t.Pair)

	if pair[len(pair)-4:] == "USDT" {
		return "USDT"
	}

	if pair[len(pair)-3:] == "EUR" {
		return "EUR"
	}

	if pair[len(pair)-3:] == "USD" {
		return "USD"
	}

	if pair[len(pair)-3:] == "ETH" {
		return "ETH"
	}

	if pair[len(pair)-3:] == "BTC" {
		return "BTC"
	}

	if len(pair) == 6 {
		return pair[len(pair)-3:]
	}

	second := strings.TrimPrefix(pair, t.Symbol+"_")
	if second != pair {
		return second
	}

	return strings.TrimPrefix(pair, t.Symbol)
}
