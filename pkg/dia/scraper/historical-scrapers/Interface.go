package historicalscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
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
		config, err := dia.GetConfig(source)
		if err != nil {
			log.Error("Get CG API key: ", err)
		}
		return NewCoingeckoScraper(rdb, datastore, config.ApiKey)
	default:
		return nil
	}
}
