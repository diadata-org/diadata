package tradescraper

import (
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

func NewTraderJoeDEXScraper(config ScraperConfig, relDB *models.RelDB, datastore *models.DB) *TraderJoeScraper {
	tjs := &TraderJoeScraper{
		relDB:        relDB,
		datastore:    datastore,
		poolChannel:  make(chan dia.Pool),
		tradeChannel: make(chan dia.Trade),
		doneChannel:  make(chan bool),
		config:       config,
	}

	go func() {
		tjs.fetchPools()
	}()

	go func() {
		tjs.fetchTrades()
	}()

	return tjs
}

func (tjs *TraderJoeScraper) fetchPools() {
	//TODO: Implement this:
	// logic to fetch pool data.
}

func (tjs *TraderJoeScraper) fetchTrades() {
	//TODO: Implement this:
	// logic to fetch trade data
}

func (tjs *TraderJoeScraper) Pool() chan dia.Pool {
	return tjs.poolChannel
}

func (tjs *TraderJoeScraper) Trade() chan dia.Trade {
	return tjs.tradeChannel
}

func (tjs *TraderJoeScraper) Done() chan bool {
	return tjs.doneChannel
}
