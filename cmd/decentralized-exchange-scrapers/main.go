package main

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/diadata-org/diadata/pkg/dia"
	dscrapers "github.com/diadata-org/diadata/pkg/dia/scraper/decentralized-exchange-scrapers"
)

func main() {
	log.Info("TO DO")
	scraper := dscrapers.NewBinanceScraper()
	centralTime, period, err := scraper.FetchClock()
	if err != nil {
		log.Fatal("Fetch clock: ", err)
	}
	currentTime := time.Now()
	nextClock := (currentTime.Unix() - centralTime.Unix())/


	var tradesblock dia.TradesBlock
	for {
		t := <-scraper.TradeChannel()
		tradesblock.TradesBlockData.Trades = append(tradesblock.TradesBlockData.Trades, *t)
		// TO DO: Every T seconds, Marshal tradesblock and store in IPFS.
	}
}
