package source

import (
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"
	"math/big"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapPair struct {
	Token0      dia.Asset
	Token1      dia.Asset
	ForeignName string
	Address     common.Address
}

const (
	restDial = "http://159.69.120.42:8545/"
	wsDial   = "ws://159.69.120.42:8546/"

	restDialBSC = ""
	wsDialBSC   = ""

	restDialPolygon = ""
	wsDialPolygon   = ""

	restDialCelo = ""
	wsDialCelo   = ""

	restDialFantom = ""
	wsDialFantom   = ""

	restDialMoonriver = ""
	wsDialMoonriver   = ""

	restDialAurora = ""
	wsDialAurora   = ""

	uniswapWaitMilliseconds     = "25"
	sushiswapWaitMilliseconds   = "100"
	pancakeswapWaitMilliseconds = "520"
	dfynWaitMilliseconds        = "100"
	ubeswapWaitMilliseconds     = "200"
	spookyswapWaitMilliseconds  = "200"
	spiritswapWaitMilliseconds  = "200"
	solarbeamWaitMilliseconds   = "200"
	trisolarisWaitMilliseconds  = "200"
)

type UniswapAssetSource struct {
	WsClient     *ethclient.Client
	RestClient   *ethclient.Client
	assetChannel chan dia.Asset
	blockchain   string
	waitTime     int
}

var exchangeFactoryContractAddress string

func NewUniswapAssetSource(exchange dia.Exchange) *UniswapAssetSource {

	var wsClient, restClient *ethclient.Client
	var err error
	var assetChannel = make(chan dia.Asset)
	var uas *UniswapAssetSource

	switch exchange.Name {
	case dia.UniswapExchange:
		log.Info(utils.Getenv("ETH_URI_REST", restDial))
		log.Info(utils.Getenv("ETH_URI_WS", wsDial))
		exchangeFactoryContractAddress = exchange.Contract.Hex()
		restClient, err = ethclient.Dial(utils.Getenv("ETH_URI_REST", restDial))
		if err != nil {
			log.Fatal(err)
		}

		wsClient, err = ethclient.Dial(utils.Getenv("ETH_URI_WS", wsDial))
		if err != nil {
			log.Fatal(err)
		}

		var waitTime int
		waitTimeString := utils.Getenv("UNISWAP_WAIT_TIME", uniswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}

		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.ETHEREUM,
			waitTime:     waitTime,
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
		var waitTime int
		waitTimeString := utils.Getenv("SUSHISWAP_WAIT_TIME", sushiswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.ETHEREUM,
			waitTime:     waitTime,
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
		var waitTime int
		waitTimeString := utils.Getenv("PANCAKESWAP_WAIT_TIME", pancakeswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 520
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.BINANCESMARTCHAIN,
			waitTime:     waitTime,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()

	case dia.DfynNetwork:
		log.Infoln("Init rest client for Polygon chain")

		restClient, err = ethclient.Dial(utils.Getenv("POLYGON_URI_REST", restDialPolygon))
		if err != nil {
			log.Fatal(err)
		}
		var waitTime int
		waitTimeString := utils.Getenv("DFYN_WAIT_TIME", dfynWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}
		uas = &UniswapAssetSource{
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.POLYGON,
			waitTime:     waitTime,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()

	case dia.QuickswapExchange:
		log.Infoln("Init rest client for Polygon chain")

		restClient, err = ethclient.Dial(utils.Getenv("POLYGON_URI_REST", restDialPolygon))
		if err != nil {
			log.Fatal(err)
		}
		var waitTime int
		waitTimeString := utils.Getenv("DFYN_WAIT_TIME", dfynWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}
		uas = &UniswapAssetSource{
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.POLYGON,
			waitTime:     waitTime,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()

	case dia.UbeswapExchange:
		log.Infoln("Init ws and rest client for Celo chain")
		wsClient, err = ethclient.Dial(utils.Getenv("CELO_URI_WS", wsDialCelo))
		if err != nil {
			log.Fatal(err)
		}
		restClient, err = ethclient.Dial(utils.Getenv("CELO_URI_REST", restDialCelo))
		if err != nil {
			log.Fatal(err)
		}
		var waitTime int
		waitTimeString := utils.Getenv("UBESWAP_WAIT_TIME", ubeswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 200
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.CELO,
			waitTime:     waitTime,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()

	case dia.SpookyswapExchange:
		log.Infoln("Init ws and rest client for Fantom chain")
		wsClient, err = ethclient.Dial(utils.Getenv("FANTOM_URI_WS", wsDialFantom))
		if err != nil {
			log.Fatal(err)
		}
		restClient, err = ethclient.Dial(utils.Getenv("FANTOM_URI_REST", restDialFantom))
		if err != nil {
			log.Fatal(err)
		}
		var waitTime int
		waitTimeString := utils.Getenv("FANTOM_WAIT_TIME", spookyswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 200
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.FANTOM,
			waitTime:     waitTime,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()

	case dia.SpiritswapExchange:
		log.Infoln("Init ws and rest client for Fantom chain")
		wsClient, err = ethclient.Dial(utils.Getenv("FANTOM_URI_WS", wsDialFantom))
		if err != nil {
			log.Fatal(err)
		}
		restClient, err = ethclient.Dial(utils.Getenv("FANTOM_URI_REST", restDialFantom))
		if err != nil {
			log.Fatal(err)
		}
		var waitTime int
		waitTimeString := utils.Getenv("FANTOM_WAIT_TIME", spiritswapWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 200
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.FANTOM,
			waitTime:     waitTime,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()

	case dia.SolarbeamExchange:
		log.Infoln("Init ws and rest client for Moonriver chain")
		wsClient, err = ethclient.Dial(utils.Getenv("MOONRIVER_URI_WS", wsDialMoonriver))
		if err != nil {
			log.Fatal(err)
		}
		restClient, err = ethclient.Dial(utils.Getenv("MOONRIVER_URI_REST", restDialMoonriver))
		if err != nil {
			log.Fatal(err)
		}
		var waitTime int
		waitTimeString := utils.Getenv("MOONRIVER_WAIT_TIME", solarbeamWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}
		uas = &UniswapAssetSource{
			WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.MOONRIVER,
			waitTime:     waitTime,
		}
		exchangeFactoryContractAddress = exchange.Contract.Hex()

	case dia.TrisolarisExchange:
		log.Infoln("Init ws and rest client for Near chain")
		// wsClient, err = ethclient.Dial(utils.Getenv("NEAR_URI_WS", wsDialNear))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		restClient, err = ethclient.Dial(utils.Getenv("AURORA_URI_REST", restDialAurora))
		if err != nil {
			log.Fatal("init rest client: ", err)
		}
		var waitTime int
		waitTimeString := utils.Getenv("AURORA_WAIT_TIME", trisolarisWaitMilliseconds)
		waitTime, err = strconv.Atoi(waitTimeString)
		if err != nil {
			log.Error("could not parse wait time: ", err)
			waitTime = 100
		}
		uas = &UniswapAssetSource{
			// WsClient:     wsClient,
			RestClient:   restClient,
			assetChannel: assetChannel,
			blockchain:   dia.AURORA,
			waitTime:     waitTime,
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
		time.Sleep(time.Duration(uas.waitTime) * time.Millisecond)
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
