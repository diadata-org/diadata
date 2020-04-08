package scrapers

import (
	"time"
)

// OptionsScraper is an interface for all of the Options Contracts scrapers
type OptionsScraper interface {
	Scrape(market string) // a self-sustained goroutine that scrapes a single market
//	ScrapeMarkets()       // will scrape the options markets defined during instantiation of the scraper
	ScraperClose(market string, websocketConnection interface{}) error
	Authenticate(market string, websocketConnection interface{}) error
}

// ComputedCVI is a struct representing our CVI value at a point in time
type ComputedCVI struct {
	CVI             float64
	CalculationTime time.Time
}

// ComputedCVIs is the channel type that will communicate the cvis
type ComputedCVIs chan ComputedCVI

// OptionSettlement - is an enum, signalling if the settlement is regular or weekly
type OptionSettlement int

// OptionSettlement enums
const (
	RegularOptionSettlement OptionSettlement = iota + 1
	WeeklyOptionSettlement
)
