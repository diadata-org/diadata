package source

import (
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// Public Solana clients.
	anyswapAPIUrl = "https://bridgeapi.anyswap.exchange/v3/serverinfoV3?chainId=all&version=STABLEV3"
)

var (
	chainMap map[string]string
)

func init() {

	chainMap = make(map[string]string)
	chainMap["1"] = dia.ETHEREUM
	chainMap["56"] = dia.BINANCESMARTCHAIN
	chainMap["137"] = dia.POLYGON
	chainMap["250"] = dia.FANTOM
	chainMap["1285"] = dia.MOONRIVER

	// chainMap["66"] = chainInfo{Name: "OKExChain", Client: ""}
	// chainMap["128"] = chainInfo{Name: "HuobiECOChain", Client: ""}
	// chainMap["288"] = chainInfo{Name: "Boba", Client: ""}
	// chainMap["42161"] = chainInfo{Name: "Arbitrum", Client: ""}
	// chainMap["43114"] = "Avalanche"

}

type assetInfo struct {
}

type AnyswapAssetSource struct {
	assetChannel chan dia.Asset
	doneChannel  chan bool
	URL          string
	relDB        *models.RelDB
}

func NewAnyswapAssetSource(exchange dia.Exchange) *AnyswapAssetSource {

	var assetChannel = make(chan dia.Asset)
	var doneChannel = make(chan bool)

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("make new relational datastore: ", err)
	}

	sas := &AnyswapAssetSource{
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		URL:          anyswapAPIUrl,
		relDB:        relDB,
	}

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
	allAssetsAllChains, err := sas.fetchEndpoint()
	if err != nil {
		log.Error("fetch endpoint: ", err)
		return
	}

	clientMap, err := getClientMap()
	if err != nil {
		log.Error("get client map: ", err)
	}

	// @chainMap can be substituted with allAssetsAllChains as soon as we have clients for all chains.
	for chainID := range clientMap {

		allAssets := allAssetsAllChains[chainID]
		for address := range allAssets {
			asset, err := ethhelper.ETHAddressToAsset(common.HexToAddress(address), clientMap[chainID], chainMap[chainID])
			if err != nil {
				log.Error("fetch asset from on-chain: ", err)
			}
			sas.assetChannel <- asset
		}

	}
	sas.doneChannel <- true

}

func getClientMap() (map[string]*ethclient.Client, error) {

	clientMap := make(map[string]*ethclient.Client)
	for key := range chainMap {
		restClient, err := ethclient.Dial(utils.Getenv("ETH_URI_"+key, ""))
		if err != nil {
			return clientMap, err
		}
		clientMap[key] = restClient

	}

	return clientMap, nil
}

// fetchEndpoint returns all assets available in the Anyswap bridge obtained through an API endpoint.
func (sas *AnyswapAssetSource) fetchEndpoint() (response map[string]map[string]interface{}, err error) {
	// @response is of type map[chainID]map[assetAddress]interface{}
	data, _, err := utils.GetRequest(sas.URL)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return
	}
	return
}
