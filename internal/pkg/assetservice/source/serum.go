package source

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/streamingfast/solana-go"
	"github.com/streamingfast/solana-go/programs/serum"
	"github.com/streamingfast/solana-go/rpc"
)

type SerumPair struct {
	Token0      dia.Asset
	Token1      dia.Asset
	ForeignName string
	Address     string
}

type listPairsResponse struct {
	Success bool   `json:"success"`
	Pairs   []string `json:"data"`
}

type orderbookResponse struct {
	Success   bool   `json:"success"`
	Orderbook Data `json:"data"`
}

type Data struct {
	MarketAddress string `json:"marketAddress"`
}

const (
	// Public Solana clients.
	wsDialSolana   = "http://192.168.1.88:8900"
	restDialSolana = "http://192.168.1.88:8899"
	restEndpointSerum = "https://serum-api.bonfida.com"
	rpcEndpointSolana = "https://solana-api.projectserum.com"
)

type SerumAssetSource struct {
	serumRestClient *http.Client
	solanaRpcClient *rpc.Client
	assetChannel chan dia.Asset
	blockchain   string
}

func NewSerumAssetSource(exchange dia.Exchange) *SerumAssetSource {

	var assetChannel = make(chan dia.Asset)
	var sas *SerumAssetSource

	exchangeFactoryContractAddress = ""

	sas = &SerumAssetSource{
		serumRestClient: &http.Client{},
		solanaRpcClient: rpc.NewClient(rpcEndpointSolana),
		assetChannel: assetChannel,
		blockchain:   dia.SOLANA,
	}

	go func() {
		sas.fetchAssets()
	}()
	return sas

}

func (sas *SerumAssetSource) Asset() chan dia.Asset {
	return sas.assetChannel
}

func (sas *SerumAssetSource) fetchAssets() {
	resp, err := sas.serumRestClient.Get(restEndpointSerum + "/pairs")
	if err != nil {
		log.Error(err)
		return
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
	listPairsResponse := listPairsResponse{}
	err = json.Unmarshal(respBody, &listPairsResponse)
	if err != nil {
		log.Error(err)
		return
	}
	if !listPairsResponse.Success {
		log.Error(listPairsResponse.Pairs[0])
		return
	}
	log.Info("Found ", len(listPairsResponse.Pairs), " pairs")
	checkMap := make(map[string]struct{})
	for _, pair := range listPairsResponse.Pairs {
		assets := strings.Split(pair, "/")
		if len(assets) < 2 {
			continue
		}
		assetInformationAdded := true
		if _, ok := checkMap[assets[0]]; !ok {
			assetInformationAdded = false
		}
		if _, ok := checkMap[assets[1]]; !ok {
			assetInformationAdded = false
		}
		if assetInformationAdded {
			continue
		}
		log.Info("Collecting information for the pair: ", pair)
		resp, err = sas.serumRestClient.Get(restEndpointSerum + "/orderbooks/" + strings.Join(assets, ""))
		if err != nil {
			log.Error(err)
			continue
		}
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
			continue
		}
		orderbookResponse := orderbookResponse{}
		err = json.Unmarshal(respBody, &orderbookResponse)
		if err != nil {
			log.Error(err)
			continue
		}
		if orderbookResponse.Orderbook.MarketAddress == "" {
			log.Error("Invalid public key")
			continue
		}
		marketAddr := solana.MustPublicKeyFromBase58(orderbookResponse.Orderbook.MarketAddress)
		info, err := serum.FetchMarket(context.TODO(), sas.solanaRpcClient, marketAddr)
		if err != nil {
			log.Error(err)
			continue
		}
		if _, ok := checkMap[assets[0]]; !ok {
			checkMap[assets[0]] = struct{}{}
			sas.assetChannel <- dia.Asset{
				Symbol:     assets[0],
				Name:       assets[0],
				Address:    info.Market.GetBaseMint().String(),
				Decimals:   info.BaseMint.Decimals,
				Blockchain: sas.blockchain,
			}
		}
		if _, ok := checkMap[assets[1]]; !ok {
			checkMap[assets[1]] = struct{}{}
			sas.assetChannel <- dia.Asset{
				Symbol:     assets[1],
				Name:       assets[1],
				Address:    info.Market.GetQuoteMint().String(),
				Decimals:   info.QuoteMint.Decimals,
				Blockchain: sas.blockchain,
			}
		}
	}
}
