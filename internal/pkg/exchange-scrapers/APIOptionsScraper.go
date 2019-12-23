package scrapers

import "time"

// OptionsScraper is an interface for all of the Options Contracts scrapers
type OptionsScraper interface {
	Scrape(market string) // a self-sustained goroutine that scrapes a single market
	ScrapeMarkets()       // will scrape the options markets defined during instantiation of the scraper
	ScraperClose(market string, websocketConnection interface{}) error
	Authenticate(market string, websocketConnection interface{}) error
}

// OptionType is used to signal whether the option is a call or a put
type OptionType int

// signals if the option is call or a put
const (
	CallOption OptionType = iota + 1
	PutOption
)

// OptionOrderbookDatum is a unit of option order book data; Other meta items like expiration time and strike price can be queried from the meta files / db.
type OptionOrderbookDatum struct {
	InstrumentName  string
	ObservationTime time.Time
	AskPrice        float64
	BidPrice        float64
	AskSize         float64
	BidSize         float64
}

// OptionMeta is option's meta information that can be directly extracted and formatted from the meta files (option scrapers)
type OptionMeta struct {
	ExpirationTime time.Time
	StrikePrice    float64
	OptionType     OptionType
}

// OptionMetaIndex is option's meta enriched with additional information OptionOrderbookDatum; This is the format used in the CVI calculation
type OptionMetaIndex struct {
	OptionMeta
	OptionOrderbookDatum
}

// OptionMetaForward is the type to be used in the calculation of the foward index level
type OptionMetaForward struct {
	StrikePrice float64
	CallPrice   float64
	PutPrice    float64 // this, as well as the above is defined as the bid price at a given strike price
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
