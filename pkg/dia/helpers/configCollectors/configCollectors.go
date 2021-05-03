package configCollectors

import (
	"os"
	"os/user"

	"github.com/diadata-org/diadata/pkg/dia"
	logrus "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var log = logrus.New()

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

func (c *ConfigCollectors) IsSymbolInConfig(symbol string) bool {
	for _, configPair := range c.Coins {
		if configPair.Symbol == symbol {
			return true
		}
	}
	return false
}

func ConfigFileConnectors(exchange string, filetype string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if dir == "/root" || dir == "/home" {
		return "/config/" + exchange + filetype //hack for docker...
	}
	if dir == "/home/travis" {
		return "../config/" + exchange + filetype //hack for travis
	}
	return os.Getenv("GOPATH") + "/src/github.com/diadata-org/diadata/config/" + exchange + ".json"
}

func NewConfigCollectorsIfExists(exchange string, filetype string) *ConfigCollectors {
	var connectorConfig = ConfigCollectors{
		Coins: []dia.Pair{},
	}
	if exchange == "" {
		for _, e := range dia.Exchanges() {
			var c = ConfigCollectors{}
			file := ConfigFileConnectors(e, filetype)
			err := gonfig.GetConf(file, &c)
			if err != nil {
				log.Error("error loading <", file, "> ", err)
			} else {
				log.Printf("loaded  <%v>", file)
				connectorConfig.Coins = append(connectorConfig.Coins, c.Coins...)
			}
		}
	} else {
		file := ConfigFileConnectors(exchange, filetype)
		err := gonfig.GetConf(file, &connectorConfig)
		if err != nil {
			log.Error("error loading <", file, "> ", err)
			return nil
		} else {
			log.Printf("loaded  <%v>", file)
		}
	}
	return &connectorConfig
}

func NewConfigCollectors(exchange string, filetype string) *ConfigCollectors {
	cc := NewConfigCollectorsIfExists(exchange, filetype)
	if cc == nil {
		log.Fatal("error in NewConfigCollectors")
	}
	return cc
}
