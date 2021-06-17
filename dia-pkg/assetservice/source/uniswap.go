package source

import (
	"github.com/diadata-org/diadata/dia-pkg/exchange-scrapers/uniswap"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type UniswapPair struct {
	Token0      dia.Asset
	Token1      dia.Asset
	ForeignName string
	Address     common.Address
}

const (
	// wsDial   = "wss://eth-mainnet.ws.alchemyapi.io/v2/CP4k5FRH3BZdqr_ANmGJFr0iI076CxR8"
	// restDial = "https://eth-mainnet.alchemyapi.io/v2/CP4k5FRH3BZdqr_ANmGJFr0iI076CxR8"
	wsDial   =  "ws://159.69.120.42:8546/" // ETH_URI_WS
	restDial = "http://159.69.120.42:8545/" // ETH_URI_REST

	wsDialBSC   = "wss://bsc-ws-node.nariox.org:443" // ETH_URI_WS_BSC
	restDialBSC = "https://bsc-dataseed2.defibit.io/" // ETH_URI_REST_BSC
)

type UniswapAssetSource struct {
	WsClient     *ethclient.Client
	RestClient   *ethclient.Client
	assetChannel chan dia.Asset
	blockchain   string
}

var exchangeFactoryContractAddress string

func NewUniswapAssetSource(exchange dia.Exchange) *UniswapAssetSource {

	var wsClient, restClient *ethclient.Client
	var err error
	var assetChannel = make(chan dia.Asset)
	var uas *UniswapAssetSource

	switch exchange.Name {
	case dia.UniswapExchange:
		log.Info(utils.Getenv("ETH_URI_REST", restDial)
		log.Info(utils.Getenv("ETH_URI_WS", wsDial)
		exchangeFactoryContractAddress = exchange.Contract.Hex()
		restClient, err = ethclient.Dial(utils.Getenv("ETH_URI_REST", restDial))
		if err != nil {
			log.Fatal(err)
		}

		wsClient, err = ethclient.Dial(utils.Getenv("ETH_URI_WS", wsDial))
		if err != nil {
			log.Fatal(err)
		}

		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.ETHEREUM,
		}
	case dia.SushiSwapExchange:
		exchangeFactoryContractAddress = exchange.Contract.Hex()
		wsClient, err = ethclient.Dial(utils.Getenv("ETH_URI_WS", wsDial))
		if err != nil {
			log.Fatal(err)
		}

		restClient, err = ethclient.Dial(utils.Getenv("ETH_URI_REST", restDial))
		if err != nil {
			log.Fatal(err)
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.ETHEREUM,
		}
	case dia.PanCakeSwap:
		log.Infoln("Init ws and rest client for BSC chain")
		wsClient, err = ethclient.Dial(utils.Getenv("ETH_URI_WS_BSC", wsDialBSC))
		if err != nil {
			log.Fatal(err)
		}
		restClient, err = ethclient.Dial(utils.Getenv("ETH_URI_REST_BSC", restDialBSC))
		if err != nil {
			log.Fatal(err)
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.BINANCESMARTCHAIN,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()
	}

	go func() {
		uas.fetchAssets()
	}()
	return uas

}

func (uas *UniswapAssetSource) Asset() chan dia.Asset {
	return uas.assetChannel
}

func (uas *UniswapAssetSource) getNumPairs() (int, error) {
	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), uas.RestClient)
	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.AllPairsLength(&bind.CallOpts{})
	return int(numPairs.Int64()), err
}

func (uas *UniswapAssetSource) fetchAssets() {

	numPairs, err := uas.getNumPairs()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Found ", numPairs, " pairs")
	checkMap := make(map[string]struct{})
	for i := 0; i < numPairs; i++ {
		pair, err := uas.GetPairByID(int64(i))
		if err != nil {
			log.Errorln("Error getting pair with ID ", i)
		}
		asset0 := pair.Token0
		asset1 := pair.Token1
		asset0.Blockchain = uas.blockchain
		asset1.Blockchain = uas.blockchain
		// Don't repeat sending already sent assets
		if _, ok := checkMap[asset0.Address]; !ok {
			if asset0.Symbol != "" {
				checkMap[asset0.Address] = struct{}{}
				uas.assetChannel <- asset0
			}
		}
		if _, ok := checkMap[asset1.Address]; !ok {
			if asset1.Symbol != "" {
				checkMap[asset1.Address] = struct{}{}
				uas.assetChannel <- asset1
			}
		}
	}
}

// GetPairByID returns the UniswapPair with the integer id @num
func (uas *UniswapAssetSource) GetPairByID(num int64) (UniswapPair, error) {
	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), uas.RestClient)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPairs(&bind.CallOpts{}, numToken)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	pair, err := uas.GetPairByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	// log.Infof("Get pair with ID %v: %v", num, pair)
	return pair, err
}

func (uas *UniswapAssetSource) GetPairByAddress(pairAddress common.Address) (pair UniswapPair, err error) {
	connection := uas.RestClient
	var pairContract *uniswap.IUniswapV2PairCaller
	pairContract, err = uniswap.NewIUniswapV2PairCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}

	// Getting tokens from pair ---------------------
	address0, _ := pairContract.Token0(&bind.CallOpts{})
	address1, _ := pairContract.Token1(&bind.CallOpts{})
	var token0Contract *uniswap.IERC20Caller
	var token1Contract *uniswap.IERC20Caller
	token0Contract, err = uniswap.NewIERC20Caller(address0, connection)
	if err != nil {
		log.Error(err)
	}
	token1Contract, err = uniswap.NewIERC20Caller(address1, connection)
	if err != nil {
		log.Error(err)
	}
	symbol0, err := token0Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	symbol1, err := token1Contract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	decimals0, err := token0Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	decimals1, err := token1Contract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}

	name0, err := uas.GetName(address0)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	name1, err := uas.GetName(address1)
	if err != nil {
		log.Error(err)
		return UniswapPair{}, err
	}
	token0 := dia.Asset{
		Address:  address0.Hex(),
		Symbol:   symbol0,
		Name:     name0,
		Decimals: decimals0,
	}
	token1 := dia.Asset{
		Address:  address1.Hex(),
		Symbol:   symbol1,
		Name:     name1,
		Decimals: decimals1,
	}
	pair.Token0 = token0
	pair.Token1 = token1

	return pair, nil
}

// GetDecimals returns the decimals of the token with address @tokenAddress
func (uas *UniswapAssetSource) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, uas.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (uas *UniswapAssetSource) GetName(tokenAddress common.Address) (name string, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, uas.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}
