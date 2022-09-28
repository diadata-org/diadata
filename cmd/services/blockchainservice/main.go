package main

import (
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

const (
	configFileBlockchains  = "blockchains/blockchains"
	configFileExchanges    = "exchanges/exchanges"
	configFileChains       = "chainconfig/chainconfig"
	configFileNFTExchanges = "nftExchanges/nftexchanges"
)

type chainConfigs struct {
	ChainConfigs []dia.ChainConfig `json:"ChainConfigs"`
}

type blockchainsConfig struct {
	Blockchains []dia.BlockChain `json:"Blockchains"`
}

type exchangesConfig struct {
	Exchanges []dia.Exchange `json:"Exchanges"`
}

type nftExchangesConfig struct {
	Exchanges []dia.NFTExchange `json:"Exchanges"`
}

func main() {

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("new relational datastore: ", err)
	}

	blockchains, err := fetchBlockchainsFromConfig()
	if err != nil {
		log.Fatal("fetch blockchains from config file: ", err)
	}
	for _, blockchain := range blockchains {
		err = rdb.SetBlockchain(blockchain)
		if err != nil {
			log.Error("set blockchain to postgres: ", err)
		}
	}

	exchanges, err := fetchExchangesFromConfig()
	if err != nil {
		log.Errorln("fetch exchanges from config file: ", err)
	} else {
		for _, exchange := range exchanges {
			err = rdb.SetExchange(exchange)
			if err != nil {
				log.Error("set blockchain to postgres: ", err)
			}
		}

	}

	nftexchanges, err := fetchNFTExchangesFromConfig()
	if err != nil {
		log.Fatal("fetch nftexchanges from config file: ", err)
	}
	for _, nftexchange := range nftexchanges {
		err = rdb.SetNFTExchange(nftexchange)
		if err != nil {
			log.Error("set SetNFTExchange to postgres: ", err)
		}
	}

	chanconfigs, err := fetchChainConfigFromConfig()
	log.Infoln(chanconfigs)

	if err != nil {
		log.Fatal("fetch chainconfigs from config file: ", err)
	}

	for _, chainconfig := range chanconfigs {
		err = rdb.SetChainConfig(chainconfig)
		log.Infoln(chainconfig)
		if err != nil {
			log.Error("set chainconfig to postgres: ", err)
		}
	}

}

func fetchBlockchainsFromConfig() (blockchains []dia.BlockChain, err error) {
	content, err := configCollectors.ReadJSONFromConfig(configFileBlockchains)
	if err != nil {
		return
	}
	var blockchainList blockchainsConfig
	err = json.Unmarshal(content, &blockchainList)
	blockchains = blockchainList.Blockchains
	return
}

func fetchChainConfigFromConfig() (chainconfigs []dia.ChainConfig, err error) {
	content, err := configCollectors.ReadJSONFromConfig(configFileChains)
	if err != nil {
		return
	}
	var chainconfigList chainConfigs
	err = json.Unmarshal(content, &chainconfigList)

	chainconfigs = chainconfigList.ChainConfigs
	return
}

func fetchExchangesFromConfig() (exchanges []dia.Exchange, err error) {
	content, err := configCollectors.ReadJSONFromConfig(configFileExchanges)
	if err != nil {
		return
	}
	var exchangeList exchangesConfig
	err = json.Unmarshal(content, &exchangeList)
	exchanges = exchangeList.Exchanges
	return
}

func fetchNFTExchangesFromConfig() (exchanges []dia.NFTExchange, err error) {
	content, err := configCollectors.ReadJSONFromConfig(configFileNFTExchanges)
	if err != nil {
		return
	}
	var exchangeList nftExchangesConfig
	err = json.Unmarshal(content, &exchangeList)
	exchanges = exchangeList.Exchanges
	return
}
