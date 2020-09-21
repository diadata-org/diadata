package dia

import (
	"strings"
)

// BaseToken returns the base token of a trading pair
func (t *Trade) BaseToken() string {

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

		// If one of these assets is base token, be sure to get price from redis
		topAssets := []string{"EUR", "USD", "BTC", "ETH", "USDC", "USDT", "BNB"}
		for _, asset := range topAssets {
			if pair[len(pair)-len(asset):] == asset {
				return asset
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

	if len(pair) == 6 {
		return pair[len(pair)-3:]
	}
	return strings.TrimPrefix(pair, t.Symbol)
}
