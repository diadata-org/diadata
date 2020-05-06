package main

import (
	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	allScrapers := scrapers.NewAllOKExOptionsScrapers(
			&wg,
			[]string{
				"BTC-USD-200508-9000-C",
				"BTC-USD-200508-9000-P",
				"BTC-USD-200508-9250-C",
				"BTC-USD-200508-10500-C",
				"BTC-USD-200626-9500-C",
				"BTC-USD-200508-11000-P",
				"BTC-USD-200508-11500-P",
				"BTC-USD-200626-10500-P",
				"BTC-USD-200626-10500-C",
				"BTC-USD-200515-7000-C",
			},
		)
	allScrapers.ScrapeMarkets()
}
