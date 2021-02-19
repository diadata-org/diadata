package dia

import (
	"strings"

	"github.com/diadata-org/diadata/pkg/utils"
)

// GetPairSymbols returns the two symbol tickers of @pair.
func GetPairSymbols(pair ExchangePair) ([]string, error) {
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
