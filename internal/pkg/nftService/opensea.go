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
	openseaAPIurl  = "https://api.opensea.io/api/v1/"
	openseaAPIWait = 250
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
	url := openseaAPIurl + "assets?order_direction=" + order_direction + "&offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit) + "&order_by=sale_count"
	resp, statusCode, err := utils.GetRequestWithStatus(url)
	if err != nil {
		if statusCode != 429 {
			return
		}
		// Retry get request once
		time.Sleep(time.Millisecond * openseaAPIWait)
		resp, _, err = utils.GetRequestWithStatus(url)
		if err != nil {
			return
		}
	}

	count := 0
	for statusCode == 429 && count < 20 {
		// Retry get request
		log.Info("sleep")
		time.Sleep(time.Millisecond * time.Duration(openseaAPIWait*count))
		resp, statusCode, err = utils.GetRequestWithStatus(url)
		log.Info("statusCode, err in second try: ", statusCode, err)
		if err != nil {
			return
		}
		count++
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
	checkmap := make(map[string]struct{})
	for k := 0; k < 200; k++ {
		assetContracts, err := fetchClasses(k*50, 50, "desc")
		if err != nil {
			log.Error(err)
		}
		for _, contract := range assetContracts {
			nftClass := dia.NFTClass{
				Address:      common.HexToAddress(contract.Address).Hex(),
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
