package scrapers

import "time"

// OptionsScraper is an interface for all of the Options Contracts scrapers
type OptionsScraper interface {
	Scrape(market string) // a self-sustained goroutine that scrapes a single market
	ScrapeMarkets()       // will scrape the options markets defined during instantiation of the scraper
	ScraperClose(market string, websocketConnection interface{}) error
	Authenticate(market string, websocketConnection interface{}) error
}

// OptionOrderbookDatum is a unit of option order book data. Other meta thing like expiration time and strike price can be queried from the meta files / db.
type OptionOrderbookDatum struct {
	InstrumentName  string
	ObservationTime time.Time
	AskPrice        float64
	BidPrice        float64
	AskSize         float64
	BidSize         float64
}
