package source

import (
	"strconv"
	"strings"

	camelotv3 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/camelotv3"
	uniswapcontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CamelotV3AssetSource struct {
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
	// signaling channels for session initialization and finishing
	assetChannel    chan dia.Asset
	doneChannel     chan bool
	blockchain      string
	startBlock      uint64
	factoryContract string
	waitTime        int
}

// NewCamelotV3AssetSource returns a new CamelotV3AssetSource
func NewCamelotV3AssetSource(exchange dia.Exchange) *CamelotV3AssetSource {
	log.Info("NewCamelotV3Scraper ", exchange.Name)
	log.Info("NewCamelotV3Scraper Address ", exchange.Contract)

	var uas *CamelotV3AssetSource

	switch exchange.Name {
	case dia.CamelotExchangeV3:
		uas = makeCamelotV3AssetSource(exchange, "", "", "200", uint64(101163738))
	}

	go func() {
		uas.fetchAssets()
	}()
	return uas

}

// makeCamelotV3AssetSource returns a Camelot asset source.
func makeCamelotV3AssetSource(exchange dia.Exchange, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *CamelotV3AssetSource {
	var (
		restClient   *ethclient.Client
		wsClient     *ethclient.Client
		err          error
		uas          *CamelotV3AssetSource
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	var waitTime int
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	uas = &CamelotV3AssetSource{
		RestClient:      restClient,
		WsClient:        wsClient,
		assetChannel:    assetChannel,
		doneChannel:     doneChannel,
		blockchain:      exchange.BlockChain.Name,
		startBlock:      startBlock,
		factoryContract: exchange.Contract,
		waitTime:        waitTime,
	}
	return uas
}

// getNumPairs returns the number of available pairs on Camelot
func (uas *CamelotV3AssetSource) fetchAssets() {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603

	log.Info("get pool creations from address: ", uas.factoryContract)
	poolsCount := 0
	checkMap := make(map[string]struct{})
	contract, err := camelotv3.NewCamelotv3Filterer(common.HexToAddress(uas.factoryContract), uas.WsClient)
	if err != nil {
		log.Error(err)
	}

	poolCreated, err := contract.FilterPool(
		&bind.FilterOpts{Start: uas.startBlock},
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Error("filter pool created: ", err)
	}
	for poolCreated.Next() {
		poolsCount++
		log.Info("pools count: ", poolsCount)
		// Don't repeat sending already sent assets
		if _, ok := checkMap[poolCreated.Event.Token0.Hex()]; !ok {
			checkMap[poolCreated.Event.Token0.Hex()] = struct{}{}
			asset, err := uas.GetAssetFromAddress(poolCreated.Event.Token0)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolCreated.Event.Token0.Hex())
			}
			uas.assetChannel <- asset
		}
		if _, ok := checkMap[poolCreated.Event.Token1.Hex()]; !ok {
			checkMap[poolCreated.Event.Token1.Hex()] = struct{}{}
			asset, err := uas.GetAssetFromAddress(poolCreated.Event.Token1)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolCreated.Event.Token1.Hex())
			}
			uas.assetChannel <- asset
		}
	}
	uas.doneChannel <- true
}

func (uas *CamelotV3AssetSource) GetAssetFromAddress(address common.Address) (asset dia.Asset, err error) {
	connection := uas.RestClient

	var tokenContract *uniswapcontract.IERC20Caller

	tokenContract, err = uniswapcontract.NewIERC20Caller(address, connection)
	if err != nil {
		log.Error(err)
	}

	symbol, err := tokenContract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	name, err := tokenContract.Name(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	decimals, err := tokenContract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}

	asset = dia.Asset{
		Symbol:     symbol,
		Name:       name,
		Address:    address.Hex(),
		Blockchain: uas.blockchain,
		Decimals:   decimals,
	}

	return
}

func (uas *CamelotV3AssetSource) Asset() chan dia.Asset {
	return uas.assetChannel
}

func (uas *CamelotV3AssetSource) Done() chan bool {
	return uas.doneChannel
}
