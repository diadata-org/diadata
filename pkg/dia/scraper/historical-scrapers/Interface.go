package historicalscrapers

import (
	models "github.com/diadata-org/diadata/pkg/model"
)

type HistoricalScraper interface {
	FetchQuotations()
	QuoteChannel() chan models.AssetQuotation
	Done() chan bool
}

func NewHistoricalScraper(source string, rdb *models.RelDB, datastore *models.DB) HistoricalScraper {
	switch source {
	case "Coinmarketcap":
		return NewCoinmarketcapScraper(rdb, datastore)
	case "Coingecko":
		return NewCoingeckoScraper(rdb, datastore)
	default:
		return nil
	}
}
