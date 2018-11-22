package dia

import (
	"strings"
)

func (t *Trade) SecondPair() string {

	pair := strings.ToUpper(t.Pair)

	switch t.Source {
	case KrakenExchange:
		if pair[len(pair)-3:] == "XBT" {
			return "BTC"
		}
	case BitfinexExchange:
		if pair[len(pair)-3:] == "USD" {
			return "USDT"
		}
	case HitBTCExchange:
		if pair[len(pair)-3:] == "USD" {
			return "USDT"
		}
	}

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

	if pair[len(pair)-3:] == "BNB" {
		return "BNB"
	}

	second := strings.TrimPrefix(pair, t.Symbol+"_")
	if second != pair {
		return second
	}

	second = strings.TrimPrefix(pair, t.Symbol+"-")
	if second != pair {
		return second
	}

	if len(pair) == 6 {
		return pair[len(pair)-3:]
	}
	return strings.TrimPrefix(pair, t.Symbol)
}
