package dia

// ContainsPair returns true iff @pair is contained in pairs
func ContainsPair(pairs []ExchangePair, pair ExchangePair) bool {
	for _, p := range pairs {
		if pair == p {
			return true
		}
	}
	return false
}

// MergePairs appends @pairs2 to @pairs1 without repetition
func MergePairs(pairs1, pairs2 []ExchangePair) []ExchangePair {
	for _, pair := range pairs2 {
		if ok := ContainsPair(pairs1, pair); !ok {
			pairs1 = append(pairs1, pair)
		}
	}
	return pairs1
}
