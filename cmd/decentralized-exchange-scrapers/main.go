package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/diadata-org/diadata/pkg/dia"
	dscrapers "github.com/diadata-org/diadata/pkg/dia/scraper/decentralized-exchange-scrapers"
)

func main() {
	log.Info("WIP")
	scraper := dscrapers.NewBinanceScraper()
	centralTime, periodSeconds, err := scraper.FetchClock()
	if err != nil {
		log.Fatal("Fetch clock: ", err)
	}

	var tradesblock dia.DTradesBlock
	tradesblock.SetTimeRange(centralTime, periodSeconds, time.Now())

	for {
		// TO DO: How to proceed with empty blocks?
		// 1. Don't send a block and let the processing layer handle.
		// 2. Send an empty block with start end endtime.
		// A: Solution 2 poses again the problem of synchronization as we would need a timer running and sending an empty
		// block to the processing layer as soon as the current block time is up.
		t := <-scraper.TradeChannel()
		log.Info("got trade: ", *t)

		// With regards to timestamps we fill tradesblocks as half-open intervals being closed at the right border.
		// i.e. the latest possible trade time is exactly @tradesblock.EndTime, while the earliest trade's time will
		// always be bigger than @tradesblock.StartTime.
		if !t.Time.After(tradesblock.EndTime) {
			tradesblock.Trades = append(tradesblock.Trades, *t)
		} else {
			if len(tradesblock.Trades) > 0 {
				// TO DO:
				// In case the block is not empty, marshal tradesblock and store in IPFS.
				// ...

				// Setup new tradesblock.
				startTimeNew := tradesblock.EndTime
				endTimeNew := time.Unix(startTimeNew.Unix()+int64(periodSeconds), 0)
				tradesblock = dia.DTradesBlock{StartTime: startTimeNew, EndTime: endTimeNew}
			} else {
				// Setup new tradesblock in case the previous was empty - and possibly not from the last time bin.
				tradesblock = dia.DTradesBlock{}
				tradesblock.SetTimeRange(centralTime, periodSeconds, time.Now())
			}
			tradesblock.Trades = append(tradesblock.Trades, *t)
		}
	}
}
