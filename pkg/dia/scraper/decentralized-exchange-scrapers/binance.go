package dscrapers

import (
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Info("TO DO")
}

type BinanceScraper struct {
	tradeChannel chan *dia.Trade
}

func NewBinanceScraper() BinanceScraper {
	scraper := BinanceScraper{
		tradeChannel: make(chan *dia.Trade),
	}
	go scraper.mainLoop()
	return scraper
}

func (scraper *BinanceScraper) mainLoop() {
	// TO DO: Fetch trades from ws API and send into tradeChannel.
}

func (scraper *BinanceScraper) ReadPairs(id string) {
	// TO DO: Read pairs from IPFS using id (as CID?).
}

func (scraper *BinanceScraper) TradeChannel() chan *dia.Trade {
	return scraper.tradeChannel
}

func (scraper *BinanceScraper) FetchClock() (time.Time, int32, error) {
	return time.Now(), 120, nil
}
