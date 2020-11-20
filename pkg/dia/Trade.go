package dia

import (
	"strings"
)

// GetBaseToken returns the base token of a trading pair
// TO DO (20/11/2020): This method is no longer needed once we switch to new Token/Trade structs
func (t *Trade) GetBaseToken() string {

	pair := strings.ToUpper(t.Pair)
	if len(pair) > 3 {
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
	}

	second := strings.TrimPrefix(pair, t.Symbol+"_")
	if second != pair {
		return second
	}
	second = strings.TrimPrefix(pair, t.Symbol+"-")
	if second != pair {
		return second
	}
	second = strings.TrimPrefix(pair, t.Symbol+"/")
	if second != pair {
		return second
	}

	return strings.TrimPrefix(pair, t.Symbol)
}
