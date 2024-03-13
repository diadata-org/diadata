package dia

import (
	"strings"

	"github.com/diadata-org/diadata/pkg/utils"
)

// GetPairSymbols returns the two symbol tickers of @pair.
func GetPairSymbols(pair ExchangePair) ([]string, error) {
	if pair.Exchange == KrakenExchange {
		pair = normalizeKrakenPair(pair)
		foreignName := pair.ForeignName
		quoteToken := pair.Symbol
		baseToken := strings.TrimPrefix(foreignName, quoteToken)
		return []string{quoteToken, baseToken}, nil
	}
	foreignName := pair.ForeignName
	quoteToken := pair.Symbol

	baseToken := strings.TrimPrefix(foreignName, quoteToken+"_")
	if baseToken != foreignName {
		return []string{quoteToken, baseToken}, nil
	}
	baseToken = strings.TrimPrefix(foreignName, quoteToken+"-")
	if baseToken != foreignName {
		return []string{quoteToken, baseToken}, nil
	}
	baseToken = strings.TrimPrefix(foreignName, quoteToken+"/")
	if baseToken != foreignName {
		return []string{quoteToken, baseToken}, nil
	}
	baseToken = strings.TrimPrefix(foreignName, quoteToken)
	return []string{quoteToken, baseToken}, nil
}

// normalizeKrakenPair addresses the particular asset notation at Kraken.
func normalizeKrakenPair(pair ExchangePair) ExchangePair {
	if len(pair.ForeignName) == 7 {
		if pair.ForeignName[len(pair.ForeignName)-3:] == "XBT" {
			pair.ForeignName = pair.ForeignName[:len(pair.ForeignName)-3] + "BTC"
		}
		if pair.ForeignName[4:5] == "Z" || pair.ForeignName[4:5] == "X" {
			pair.ForeignName = pair.ForeignName[:4] + pair.ForeignName[5:]
		}
		if pair.ForeignName[:1] == "Z" || pair.ForeignName[:1] == "X" {
			pair.ForeignName = pair.ForeignName[1:]
		}
	}
	if len(pair.ForeignName) == 8 {
		if pair.ForeignName[4:5] == "Z" || pair.ForeignName[4:5] == "X" {
			pair.ForeignName = pair.ForeignName[:4] + pair.ForeignName[5:]
		}
		if pair.ForeignName[:1] == "Z" || pair.ForeignName[:1] == "X" {
			pair.ForeignName = pair.ForeignName[1:]
		}
	}
	if pair.ForeignName[len(pair.ForeignName)-3:] == "XBT" {
		pair.ForeignName = pair.ForeignName[:len(pair.ForeignName)-3] + "BTC"
	}
	if pair.ForeignName[:3] == "XBT" {
		pair.ForeignName = "BTC" + pair.ForeignName[len(pair.ForeignName)-3:]
	}
	return pair
}

// GetAllSymbolsFromPairs returns a unique list of symbols which constitute @pairs.
func GetAllSymbolsFromPairs(pairs []ExchangePair) ([]string, error) {
	var symbols []string
	for _, pair := range pairs {

		pairsymbols, err := GetPairSymbols(pair)
		if err != nil {
			return []string{}, err
		}
		symbols = append(symbols, pairsymbols[0], pairsymbols[1])
	}

	uniqueSymbols := utils.UniqueStrings(symbols)
	return uniqueSymbols, nil
}

// GetAllAssetsFromPairs returns the unique slice of assets underlying the exchange pairs @pairs.
func GetAllAssetsFromPairs(pairs []ExchangePair) (assets []Asset) {
	uniqueMap := make(map[Asset]struct{})

	for _, pair := range pairs {
		if _, ok := uniqueMap[pair.UnderlyingPair.BaseToken]; !ok {
			assets = append(assets, pair.UnderlyingPair.BaseToken)
			uniqueMap[pair.UnderlyingPair.BaseToken] = struct{}{}
		}
		if _, ok := uniqueMap[pair.UnderlyingPair.QuoteToken]; !ok {
			assets = append(assets, pair.UnderlyingPair.QuoteToken)
			uniqueMap[pair.UnderlyingPair.QuoteToken] = struct{}{}
		}
	}
	return
}

// ContainsExchangePair returns true iff @pair is contained in pairs.
// Here, equality refers to the unique identifier (exchange,foreignName).
func ContainsExchangePair(pairs []ExchangePair, pair ExchangePair) bool {
	for _, p := range pairs {
		if pair.Exchange == p.Exchange && pair.ForeignName == p.ForeignName {
			return true
		}
	}
	return false
}

// MergeExchangePairs appends @pairs2 to @pairs1 without repetition.
func MergeExchangePairs(pairs1, pairs2 []ExchangePair) []ExchangePair {
	for _, pair := range pairs2 {
		if ok := ContainsExchangePair(pairs1, pair); !ok {
			pairs1 = append(pairs1, pair)
		}
	}
	return pairs1
}
