package scrapers

import (
	//"log"
	"sync"

	//writers "github.com/diadata-org/diadata/internal/pkg/scraper-writers"
	"github.com/diadata-org/diadata/pkg/model"
	zap "go.uber.org/zap"
)

// DeribitScraperKind - used to distinguish between the futures and options scrapers
type DeribitScraperKind int

const (
	// DeribitFuture - constant to signal the futures scraper
	DeribitFuture DeribitScraperKind = iota + 1
	// DeribitOption - constant to signal the options scraper
	DeribitOption
)

// DeribitScraper - used in conjunction with the DeribitScraperKind in a new struct to define futures and options scrapers
type DeribitScraper struct {
	Markets					[]string
	WaitGroup				*sync.WaitGroup
	//Writer					writers.Writer
	Logger					*zap.SugaredLogger
	DataStore  *models.DB

	// required for deribit to:
	// 1. authenticate (trades is a private channel)
	// 2. referesh the token from step 1., so that the channel isn't closed
	AccessKey    string
	AccessSecret string

	RefreshTokenEvery int16 // how often we refresh the token (in seconds)
	MarketKind        DeribitScraperKind
}
