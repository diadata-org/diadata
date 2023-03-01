package source

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"

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
	restDial          = "http://159.69.120.42:8545/"
	restDialBSC       = ""
	restDialPolygon   = ""
	restDialCelo      = ""
	restDialFantom    = ""
	restDialMoonriver = ""
	restDialAurora    = ""
	restDialArbitrum  = ""
	restDialMetis     = ""
	restDialAvalanche = ""
	restDialTelos     = ""
	restDialEvmos     = ""
	restDialAstar     = ""
	restDialMoonbeam  = ""
	restDialWanchain  = ""

	uniswapWaitMilliseconds     = "25"
	sushiswapWaitMilliseconds   = "100"
	pancakeswapWaitMilliseconds = "100"
	dfynWaitMilliseconds        = "100"
	ubeswapWaitMilliseconds     = "200"
	spookyswapWaitMilliseconds  = "200"
	spiritswapWaitMilliseconds  = "200"
	solarbeamWaitMilliseconds   = "200"
	trisolarisWaitMilliseconds  = "200"
	metisWaitMilliseconds       = "200"
	moonriverWaitMilliseconds   = "500"
	avalancheWaitMilliseconds   = "200"
	telosWaitMilliseconds       = "400"
	evmosWaitMilliseconds       = "400"
	astarWaitMilliseconds       = "1000"
	moonbeamWaitMilliseconds    = "1000"
	wanchainWaitMilliseconds    = "1000"
)

type UniswapAssetSource struct {
	RestClient   *ethclient.Client
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
	waitTime     int
}

var exchangeFactoryContractAddress string

func NewUniswapAssetSource(exchange dia.Exchange) (uas *UniswapAssetSource) {

	switch exchange.Name {
	case dia.UniswapExchange:
		uas = makeUniswapAssetSource(exchange, restDial, uniswapWaitMilliseconds)
	case dia.SushiSwapExchange:
		uas = makeUniswapAssetSource(exchange, restDial, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangeArbitrum:
		uas = makeUniswapAssetSource(exchange, restDialArbitrum, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangeFantom:
		uas = makeUniswapAssetSource(exchange, restDialFantom, sushiswapWaitMilliseconds)
	case dia.CamelotExchange:
		uas = makeUniswapAssetSource(exchange, restDialArbitrum, sushiswapWaitMilliseconds)
	case dia.PanCakeSwap:
		uas = makeUniswapAssetSource(exchange, restDialBSC, pancakeswapWaitMilliseconds)
	case dia.DfynNetwork:
		uas = makeUniswapAssetSource(exchange, restDialPolygon, dfynWaitMilliseconds)
	case dia.QuickswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialPolygon, dfynWaitMilliseconds)
	case dia.UbeswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialCelo, ubeswapWaitMilliseconds)
	case dia.SpookyswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialFantom, spookyswapWaitMilliseconds)
	case dia.SpiritswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialFantom, spiritswapWaitMilliseconds)
	case dia.SolarbeamExchange:
		uas = makeUniswapAssetSource(exchange, restDialMoonriver, solarbeamWaitMilliseconds)
	case dia.TrisolarisExchange:
		uas = makeUniswapAssetSource(exchange, restDialAurora, trisolarisWaitMilliseconds)
	case dia.NetswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialMetis, metisWaitMilliseconds)
	case dia.HuckleberryExchange:
		uas = makeUniswapAssetSource(exchange, restDialMoonriver, moonriverWaitMilliseconds)
	case dia.TraderJoeExchange:
		uas = makeUniswapAssetSource(exchange, restDialAvalanche, avalancheWaitMilliseconds)
	case dia.PangolinExchange:
		uas = makeUniswapAssetSource(exchange, restDialAvalanche, avalancheWaitMilliseconds)
	case dia.TethysExchange:
		uas = makeUniswapAssetSource(exchange, restDialMetis, metisWaitMilliseconds)
	case dia.HermesExchange:
		uas = makeUniswapAssetSource(exchange, restDialMetis, metisWaitMilliseconds)
	case dia.OmniDexExchange:
		uas = makeUniswapAssetSource(exchange, restDialTelos, telosWaitMilliseconds)
	case dia.DiffusionExchange:
		uas = makeUniswapAssetSource(exchange, restDialEvmos, evmosWaitMilliseconds)
	case dia.ArthswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialAstar, astarWaitMilliseconds)
	case dia.StellaswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialMoonbeam, moonbeamWaitMilliseconds)
	case dia.WanswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialWanchain, wanchainWaitMilliseconds)
	}

	exchangeFactoryContractAddress = exchange.Contract

	go func() {
		uas.fetchAssets()
	}()
	return uas

}

// makeUniswapAssetSource returns an asset source as used in NewUniswapAssetSource.
func makeUniswapAssetSource(exchange dia.Exchange, restDial string, waitMilliseconds string) *UniswapAssetSource {
	var (
		restClient   *ethclient.Client
		err          error
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		uas          *UniswapAssetSource
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
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
	uas = &UniswapAssetSource{
		RestClient:   restClient,
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		waitTime:     waitTime,
	}
	return uas
}

func (uas *UniswapAssetSource) Asset() chan dia.Asset {
	return uas.assetChannel
}

func (uas *UniswapAssetSource) Done() chan bool {
	return uas.doneChannel
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
		pair, err := uas.GetPairByID(int64(numPairs - 1 - i))
		if err != nil {
			log.Errorln("Error getting pair with ID ", numPairs-1-i)
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
	uas.doneChannel <- true
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
