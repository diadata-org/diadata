package main

import (
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

const (
	configFileBlockchains = "blockchains/blockchains"
)

type blockchainsConfig struct {
	Blockchains []dia.BlockChain `json:"Blockchains"`
}

func main() {

	rdb, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("new relational datastore: ", err)
	}

	blockchains, err := fetchBlockchainsfromConfig()
	if err != nil {
		log.Fatal("fetch blockchains from config file: ", err)
	}

	for _, blockchain := range blockchains {
		err = rdb.SetBlockchain(blockchain)
		if err != nil {
			log.Error("set blockchain to postgres: ", err)
		}
	}

}

func fetchBlockchainsfromConfig() (blockchains []dia.BlockChain, err error) {
	content, err := configCollectors.ReadJSONFromConfig(configFileBlockchains)
	if err != nil {
		return
	}
	var blockchainList blockchainsConfig
	err = json.Unmarshal(content, &blockchainList)
	blockchains = blockchainList.Blockchains
	return
}
