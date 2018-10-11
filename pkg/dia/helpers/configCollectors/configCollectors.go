package configCollectors

import (
	"github.com/diadata-org/diadata/pkg/dia"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
	"os"
	"os/user"
)

func (c *ConfigCollectors) Exchanges() []string {
	return []string{dia.BinanceExchange, dia.BitfinexExchange, dia.CoinBaseExchange, dia.KrakenExchange, dia.UnknownExchange}
}

type ConfigCollectors struct {
	Coins []dia.Pair
}

func (c *ConfigCollectors) AllPairs() []dia.Pair {
	founds := map[string]bool{}
	result := []dia.Pair{}
	for _, configPair := range c.Coins {
		if _, ok := founds[configPair.ForeignName]; !ok {
			founds[configPair.ForeignName] = true
			result = append(result, configPair)
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

func configFileConnectors(exchange string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if dir == "/home" {
		return "/config/" + exchange + ".json" //hack for docker...
	}
	if dir == "/home/travis" {
		return "../config/" + exchange + ".json" //hack for travis
	}
	return os.Getenv("GOPATH") + "/src/github.com/diadata-org/diadata/config/" + exchange + ".json"
}

func NewConfigCollectors(exchange string) *ConfigCollectors {
	var connectorConfig = ConfigCollectors{
		Coins: []dia.Pair{},
	}
	if exchange == "" {
		for _, e := range dia.Exchanges() {
			var c = ConfigCollectors{}
			file := configFileConnectors(e)
			err := gonfig.GetConf(file, &c)
			if err != nil {
				log.Error("error loading <", file, "> ", err)
			} else {
				log.Printf("loaded  <%v>", file)
				connectorConfig.Coins = append(connectorConfig.Coins, c.Coins...)
			}
		}
	} else {
		file := configFileConnectors(exchange)
		err := gonfig.GetConf(file, &connectorConfig)
		if err != nil {
			log.Fatal("error loading <", file, "> ", err)
			return nil
		} else {
			log.Printf("loaded  <%v>", file)
		}
	}

	return &connectorConfig
}
