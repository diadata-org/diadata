package main

import (
	"sync"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
)

func collectMessages(chanTrades chan *dia.Trade) {
	// var t *dia.Trade
	for {
		select {
		// case t = <-chanTrades
		case <-chanTrades:
			continue
			// log.Infoln(*t)
		}
	}
}

func fetchAll() (c []dia.Pair, e error) {
	s := scrapers.NewCoinflexScraper()
	c, e = s.FetchAvailablePairs()
	return
}

func main() {
	log.SetLevel(log.DebugLevel)
	// c, _ := fetchAll()
	// fmt.Printf("%+v\n", c)
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := scrapers.NewCoinflexScraper()
	go c.ScrapePair(dia.Pair{
		Symbol:      "XBT/USDT",
		ForeignName: "10005/10053",
		Exchange:    dia.CoinflexExchange,
		Ignore:      false,
	})
	go c.ScrapePair(dia.Pair{
		Symbol:      "ETH/USDT",
		ForeignName: "10029/10053",
		Exchange:    dia.CoinflexExchange,
		Ignore:      false,
	})
	go c.ScrapePair(dia.Pair{
		Symbol:      "USDC/USDT",
		ForeignName: "10041/10053",
		Exchange:    dia.CoinflexExchange,
		Ignore:      false,
	})
	go c.ScrapePair(dia.Pair{
		Symbol:      "BCH/USDT",
		ForeignName: "10017/10053",
		Exchange:    dia.CoinflexExchange,
		Ignore:      false,
	})
	go c.ScrapePair(dia.Pair{
		Symbol:      "XBT/USDT",
		ForeignName: "63488/65283",
		Exchange:    dia.CoinflexExchange,
		Ignore:      false,
	})
	go c.ScrapePair(dia.Pair{
		Symbol:      "ETH/USDT",
		ForeignName: "63520/65283",
		Exchange:    dia.CoinflexExchange,
		Ignore:      false,
	})
	go collectMessages(c.Channel())
	wg.Wait()
}
