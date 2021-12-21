package source

import (
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	// Public Solana clients.
	anyswapAPIUrl = "https://bridgeapi.anyswap.exchange/v3/serverinfoV3?chainId=all&version=STABLEV3"
)

type anyToken struct {
	address  string
	name     string
	symbol   string
	decimals uint8
}

type AnyswapAssetSource struct {
	assetChannel chan dia.Asset
	doneChannel  chan bool
	URL          string
}

func NewAnyswapAssetSource(exchange dia.Exchange) *AnyswapAssetSource {

	var assetChannel = make(chan dia.Asset)
	var doneChannel = make(chan bool)

	sas := &AnyswapAssetSource{
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		URL:          anyswapAPIUrl,
	}

	sas.fetchEndpoint()

	go func() {
		sas.fetchAssets()
	}()
	return sas

}

func (sas *AnyswapAssetSource) Asset() chan dia.Asset {
	return sas.assetChannel
}

func (sas *AnyswapAssetSource) Done() chan bool {
	return sas.doneChannel
}

func (sas *AnyswapAssetSource) fetchAssets() {

}

func (sas *AnyswapAssetSource) fetchEndpoint() (string, error) {
	log.Info("fetch api ")
	data, _, err := utils.GetRequest(sas.URL)
	if err != nil {
		return "", err
	}
	var response map[string]map[string]interface{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return "", err
	}
	for key := range response {
		log.Info("key: ", response[key])
	}
	return "", nil
}
