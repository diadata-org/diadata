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

// TODO: Check if we can delete this function
func (t *Trade) NormalizedSymbol() string {
	symbol := strings.ToUpper(t.Symbol)

	switch t.Source {
	case KrakenExchange:
		if (len(symbol) < 4) {
			return t.Symbol
		}
		if (symbol[:3] == "XBT" || symbol[:4] == "XXBT") {
			return "BTC"
		}
		if (symbol[:4] == "XREP") {
			return "REP"
		}
		if (symbol[:4] == "XETH") {
			return "ETH"
		}
		if (symbol[:4] == "XETC") {
			return "ETC"
		}
		if (symbol[:4] == "XLTC") {
			return "LTC"
		}
		if (symbol[:4] == "XXMR") {
			return "XMR"
		}
		if (symbol[:4] == "XXRP") {
			return "XRP"
		}
		if (symbol[:4] == "XXLM") {
			return "XLM"
		}
		if (symbol[:4] == "XZEC") {
			return "ZEC"
		}
	}
	return t.Symbol
}
