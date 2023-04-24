package dscrapers

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

type CEXScraper interface {
	// Read pairs from IPFS using a unique identifier (CID?).
	ReadPairs(id string) ([]dia.ExchangePair, error)
	// Channel that allows the scraper forwarding trades to the main.
	TradeChannel() chan *dia.Trade
	// Central clock allowing all decentralized scrapers to synchronize.
	// It consists of an (initial) timestamp and a block period given in seconds.
	FetchClock(id string) (time.Time, uint32, error)
}

type DEXScraper interface {
	ReadPools(id string) ([]dia.Pool, error)
	TradeChannel() chan *dia.Trade
	FetchClock(id string) (time.Time, uint32, error)
}
