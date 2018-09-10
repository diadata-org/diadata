package scrapers

import (
	"io"

	"github.com/diadata-org/api-golang/dia"
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

func NewAPIScraper(exchange string, key string, secret string) APIScraper {
	switch exchange {
	case dia.BinanceExchange:
		return NewBinanceScraper(key, secret)
	case dia.BitfinexExchange:
		return NewBitfinexScraper(key, secret)
	case dia.CoinBaseExchange:
		return NewCoinBaseScraper()
	case dia.KrakenExchange:
		return NewKrakenScraper(key, secret)
	case dia.HitBTCExchange:
		return NewHitBTCScraper()
	case dia.SimexExchange:
		return NewSimexScraper()
	case dia.OKExExchange:
		return NewOKExScraper()
	case dia.HuobiExchange:
		return NewHuobiScraper()
	default:
		return nil
	}
}
