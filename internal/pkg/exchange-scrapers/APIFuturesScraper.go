package scrapers

// FuturesScraper is an interface for all of the Futures Contracts scrapers
type FuturesScraper interface {
	Scrape(market string) // a self-sustained goroutine that scrapes a single market
	ScrapeMarkets()       // will scrape the futures markets defined during instantiation of the scraper
	ScraperClose(market string, websocketConnection interface{}) error
	Authenticate(market string, websocketConnection interface{}) error
}

const retryIn uint8 = 5 // how long to wait in seconds before restarting a failed websocket
