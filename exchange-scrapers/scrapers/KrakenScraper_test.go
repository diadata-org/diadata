package scrapers_test

import (
	"github.com/beldur/kraken-go-api-client"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/exchange-scrapers/scrapers"
	"github.com/tkanos/gonfig"
	"testing"
)

// to avoid that we add a crypto to trade in the config only and forget to update CurrentBalance
func TestThatCurrentBalanceMapperMatchesConfig(t *testing.T) {
	var configConnector dia.ConfigConnector

	configFile := "../config/exchange-scrapers.json"
	err := gonfig.GetConf(configFile, &configConnector)

	if err != nil {
		t.Errorf("error loading configFile")
	}

	for _, configPair := range configConnector.Coins {
		if configPair.Exchange == dia.KrakenExchange {
			_, err = scrapers.CurrentBalance(configPair.Symbol, &krakenapi.BalanceResponse{})
			if err != nil {
				t.Errorf("please update CurrentBalance for the new trading crypto added in the config")
			}
		}
	}

	_, err = scrapers.CurrentBalance("blabla", &krakenapi.BalanceResponse{})
	if err == nil {
		t.Errorf("blabla")
	}
}
