package nftsource

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
)

const (
	openseaAPIurl = "https://api.opensea.io/api/v1/"
)

type openseaAPIResponse struct {
	AssetContracts []Asset `json:"assets"`
}

type Asset struct {
	AssetContract AssetContract `json:"asset_contract"`
}

type AssetContract struct {
	Address      string `json:"address"`
	ContractType string `json:"asset_contract_type"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
}

type OpenseaNFTSource struct {
	APIurl          string
	nftClassChannel chan dia.NFTClass
	nftChannel      chan dia.NFT
	closed          chan bool
	blockchain      string
}

func NewOpenseaNFTSource(secret string) *OpenseaNFTSource {

	var nftClassChannel = make(chan dia.NFTClass)
	var closed = make(chan bool)

	ons := &OpenseaNFTSource{
		APIurl:          openseaAPIurl,
		nftClassChannel: nftClassChannel,
		closed:          closed,
		blockchain:      dia.ETHEREUM,
	}

	go func() {
		ons.fetchAllNFTClasses()
	}()
	return ons

}

func (ons *OpenseaNFTSource) NFTClass() chan dia.NFTClass {
	return ons.nftClassChannel
}

func (ons *OpenseaNFTSource) Close() chan bool {
	return ons.closed
}

// retrieve nft classes from opensea api. Ordered by number of sales in descending order.
func fetchClasses(offset, limit int, order_direction string) (acs []AssetContract, err error) {
	resp, err := utils.GetRequest(openseaAPIurl + "assets?order_direction=" + order_direction + "&offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit) + "&order_by=sale_count")
	if err != nil {
		return
	}
	var openseaResponse openseaAPIResponse
	err = json.Unmarshal(resp, &openseaResponse)
	if err != nil {
		return
	}
	// make assetContract slice unique
	checkmap := make(map[string]struct{})
	for _, contract := range openseaResponse.AssetContracts {
		if _, ok := checkmap[contract.AssetContract.Address]; !ok {
			acs = append(acs, contract.AssetContract)
			checkmap[contract.AssetContract.Address] = struct{}{}
		}
	}
	return
}

func (ons *OpenseaNFTSource) fetchAllNFTClasses() {
	// totalPages := 50 * 200 - this is the limit on offset parameter in the api endpoint
	checkmap := make(map[common.Address]struct{})
	for k := 0; k < 200; k++ {
		assetContracts, err := fetchClasses(k*50, 50, "desc")
		if err != nil {
			log.Error(err)
		}
		// assetContractsAsc, err := fetchClasses(k*50, 50, "asc")
		// assetContracts = append(assetContracts, assetContractsAsc...)
		// if err != nil {
		// 	log.Error(err)
		// }
		for _, contract := range assetContracts {
			nftClass := dia.NFTClass{
				Address:      common.HexToAddress(contract.Address),
				Symbol:       contract.Symbol,
				Name:         contract.Name,
				Blockchain:   dia.ETHEREUM,
				ContractType: contract.ContractType,
			}
			if _, ok := checkmap[nftClass.Address]; !ok {
				ons.nftClassChannel <- nftClass
				checkmap[nftClass.Address] = struct{}{}
			}
		}
		time.Sleep(2 * time.Second)
	}

	// Gracefully close channel after iterating through all classes
	ons.closed <- true
}
