package configCollectors

import (
	"github.com/diadata-org/api-golang/dia"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"os"
	"os/user"
)

func (c *ConfigCollectors) Exchanges() []string {
	return []string{dia.BinanceExchange, dia.BitfinexExchange, dia.CoinBaseExchange, dia.KrakenExchange, dia.UnknownExchange}
}

type ConfigCollectors struct {
	Coins []dia.ConfigPair
}

func (c *ConfigCollectors) AllConfigPairsForExchange(exchange string) []dia.ConfigPair {
	founds := map[string]bool{}
	result := []dia.ConfigPair{}
	for _, configPair := range c.Coins {
		if configPair.Exchange == exchange {
			if _, ok := founds[configPair.ForeignName]; !ok {
				founds[configPair.ForeignName] = true
				result = append(result, configPair)
			}
		}
	}
	return result
}

func (c *ConfigCollectors) AllSymbols() []string {
	founds := map[string]bool{}
	result := []string{}

	for _, configPair := range c.Coins {
		if _, ok := founds[configPair.Symbol]; !ok {
			founds[configPair.Symbol] = true
			result = append(result, configPair.Symbol)
		}
	}
	return result
}

func (c *ConfigCollectors) IsSymbolInConfig(symbol string) bool {
	for _, configPair := range c.Coins {
		if configPair.Symbol == symbol {
			return true
		}
	}
	return false
}

func ConfigFileConnectors() string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if dir == "/home" {
		return "/config/exchange-scrapers.json" //hack for docker...
	}
	if dir == "/home/travis" {
		return "../exchange-scrapers/config/exchange-scrapers.json" //hack for travis
	}
	return os.Getenv("GOPATH") + "/src/github.com/diadata-org/api-golang/exchange-scrapers/config/exchange-scrapers.json"
}

func NewConfigCollectors() *ConfigCollectors {

	file := ConfigFileConnectors()

	var connectorConfig = ConfigCollectors{}

	err := gonfig.GetConf(file, &connectorConfig)
	if err != nil {
		log.Fatal("error loading <", file, "> ", err)
		return nil
	} else {
		log.Printf("loaded  <%v>", file)
	}
	return &connectorConfig
}
