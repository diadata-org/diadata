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
	tradeChannel chan *dia.DTrade
}

func NewBinanceScraper() BinanceScraper {
	scraper := BinanceScraper{
		tradeChannel: make(chan *dia.DTrade),
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

func (scraper *BinanceScraper) TradeChannel() chan *dia.DTrade {
	return scraper.tradeChannel
}

func (scraper *BinanceScraper) FetchClock() (time.Time, uint32, error) {
	// TO DO: Fetch clock from central control unit.
	return time.Now(), 120, nil
}
