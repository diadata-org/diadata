package configCollectors

import (
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/diadata-org/diadata/pkg/dia"
	logrus "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var log = logrus.New()

func (c *ConfigCollectors) Exchanges() []string {
	return []string{dia.BinanceExchange, dia.BitfinexExchange, dia.CoinBaseExchange, dia.KrakenExchange, dia.UnknownExchange}
}

type ConfigCollectors struct {
	Coins []dia.ExchangePair
}

func (c *ConfigCollectors) AllPairs() []dia.ExchangePair {
	founds := map[string]bool{}
	result := []dia.ExchangePair{}
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

// ConfigFileConnectors returns a path to folder @exchange in config folder if @filteype is empty.
// If @filteype is a filetype it returns the path to file @exchange as a @filteype file.
func ConfigFileConnectors(exchange string, filetype string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if dir == "/root" || dir == "/home" {
		return "/config/" + exchange + filetype //hack for docker...
	}
	if dir == "/home/travis" {
		return "../config/" + exchange + filetype //hack for travis
	}
	if configDir := os.Getenv("DIA_CONFIG_DIR"); configDir != "" { //hack for local tests.
		return path.Join(configDir, exchange+filetype)
	}
	return os.Getenv("GOPATH") + "/src/github.com/diadata-org/diadata/config/" + exchange + filetype
}

func NewConfigCollectorsIfExists(exchange string, filetype string) *ConfigCollectors {
	var connectorConfig = ConfigCollectors{
		Coins: []dia.ExchangePair{},
	}

	file := ConfigFileConnectors(exchange, filetype)
	err := gonfig.GetConf(file, &connectorConfig)
	if err != nil {
		log.Error("error loading <", file, "> ", err)
		return nil
	} else {
		log.Printf("loaded  <%v>", file)
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

// ReadJSONFromConfig reads a json file from the config folder and returns the byte slice of items.
func ReadJSONFromConfig(filename string) (content []byte, err error) {
	var (
		jsonFile *os.File
	)
	jsonFile, err = os.Open(ConfigFileConnectors(filename, ".json"))
	if err != nil {
		return
	}
	defer func() {
		cerr := jsonFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	content, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	return
}
