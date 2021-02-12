package dia

import (
	"errors"
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

	second := strings.TrimPrefix(pair, strings.ToUpper(t.Symbol)+"_")
	if second != pair {
		return second
	}
	second = strings.TrimPrefix(pair, strings.ToUpper(t.Symbol)+"-")
	if second != pair {
		return second
	}
	second = strings.TrimPrefix(pair, strings.ToUpper(t.Symbol)+"/")
	if second != pair {
		return second
	}

	return strings.TrimPrefix(pair, strings.ToUpper(t.Symbol))
}

// SwapTrade swaps base and quote token of a trade and inverts the price accordingly
func SwapTrade(t Trade) (Trade, error) {
	if t.Price == 0 {
		return t, errors.New("zero price. cannot swap trade")
	}
	symbol := t.Symbol
	baseToken := (&t).GetBaseToken()
	t.Symbol = baseToken
	t.Pair = baseToken + "-" + symbol
	t.Volume = -t.Price * t.Volume
	t.Price = 1 / t.Price

	return t, nil
}
