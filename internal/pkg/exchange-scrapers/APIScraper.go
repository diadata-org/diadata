package scrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
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
	// FetchAvailablePairs returns a list with all available trade pairs
	FetchAvailablePairs() (pairs []dia.Pair, err error)
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
		return NewBinanceScraper(key, secret, dia.BinanceExchange)
	case dia.BitfinexExchange:
		return NewBitfinexScraper(key, secret, dia.BitfinexExchange)
	case dia.CoinBaseExchange:
		return NewCoinBaseScraper(dia.CoinBaseExchange)
	case dia.KrakenExchange:
		return NewKrakenScraper(key, secret, dia.KrakenExchange)
	case dia.HitBTCExchange:
		return NewHitBTCScraper(dia.HitBTCExchange)
	case dia.SimexExchange:
		return NewSimexScraper(dia.SimexExchange)
	case dia.OKExExchange:
		return NewOKExScraper(dia.OKExExchange)
	case dia.HuobiExchange:
		return NewHuobiScraper(dia.HuobiExchange)
	case dia.LBankExchange:
		return NewLBankScraper(dia.LBankExchange)
	case dia.GateIOExchange:
		return NewGateIOScraper(dia.GateIOExchange)
	case dia.ZBExchange:
		return NewZBScraper(dia.ZBExchange)
	default:
		return nil
	}
}
