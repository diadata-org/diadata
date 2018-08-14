package scrapers

import (
	"github.com/diadata-org/api-golang/dia"
	"io"
)

// empty type used for signaling
type nothing struct{}

// APIScraper provides common methods needed to get Trade information from
// exchange APIs.
type APIScraper interface {
	io.Closer

	// ScrapePair returns a PairScraper that continuously scrapes trades for a
	// single pair from this APIScraper
	ScrapePair(pair dia.Pair) (PairScraper, error)
}

// PairScraper receives trades for a single pc.Pair from a single exchange.
type PairScraper interface {
	io.Closer

	// Channel returns a channel that can be used to receive trades
	Channel() chan *dia.Trade

	// Error returns an error when the channel Channel() is closed
	// and nil otherwise
	Error() error

	// Pair returns the pair this scraper is subscribed to
	Pair() dia.Pair
}
