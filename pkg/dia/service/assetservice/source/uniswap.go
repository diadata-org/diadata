package source

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"
	models "github.com/diadata-org/diadata/pkg/model"

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
	restDial              = ""
	restDialBase          = ""
	restDialBSC           = ""
	restDialPolygon       = ""
	restDialCelo          = ""
	restDialFantom        = ""
	restDialMoonriver     = ""
	restDialAurora        = ""
	restDialArbitrum      = ""
	restDialMetis         = ""
	restDialAvalanche     = ""
	restDialTelos         = ""
	restDialEvmos         = "https://evmos-evm.publicnode.com"
	restDialAstar         = ""
	restDialMoonbeam      = ""
	restDialWanchain      = ""
	restDialUnrealTestnet = ""
	restDialUnreal        = ""
	restDialLinea         = ""
	restDialSonic         = ""

	uniswapWaitMilliseconds     = "25"
	baseWaitMilliseconds        = "200"
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
	relDB        *models.RelDB
	assetChannel chan dia.Asset
	doneChannel  chan bool
	exchange     dia.Exchange
	waitTime     int
}

var exchangeFactoryContractAddress string

func NewUniswapAssetSource(exchange dia.Exchange, relDB *models.RelDB) (uas *UniswapAssetSource) {

	switch exchange.Name {
	case dia.UniswapExchange:
		uas = makeUniswapAssetSource(exchange, restDial, relDB, uniswapWaitMilliseconds)
	case dia.UniswapExchangeBase:
		uas = makeUniswapAssetSource(exchange, restDialBase, relDB, baseWaitMilliseconds)
	case dia.UniswapExchangeArbitrum:
		uas = makeUniswapAssetSource(exchange, restDialArbitrum, relDB, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchange:
		uas = makeUniswapAssetSource(exchange, restDial, relDB, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangeArbitrum:
		uas = makeUniswapAssetSource(exchange, restDialArbitrum, relDB, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangeFantom:
		uas = makeUniswapAssetSource(exchange, restDialFantom, relDB, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangePolygon:
		uas = makeUniswapAssetSource(exchange, restDialPolygon, relDB, sushiswapWaitMilliseconds)
	case dia.ApeswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialBSC, relDB, sushiswapWaitMilliseconds)
	case dia.CamelotExchange:
		uas = makeUniswapAssetSource(exchange, restDialArbitrum, relDB, sushiswapWaitMilliseconds)
	case dia.PanCakeSwap:
		uas = makeUniswapAssetSource(exchange, restDialBSC, relDB, pancakeswapWaitMilliseconds)
	case dia.DfynNetwork:
		uas = makeUniswapAssetSource(exchange, restDialPolygon, relDB, dfynWaitMilliseconds)
	case dia.QuickswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialPolygon, relDB, dfynWaitMilliseconds)
	case dia.UbeswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialCelo, relDB, ubeswapWaitMilliseconds)
	case dia.SpookyswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialFantom, relDB, spookyswapWaitMilliseconds)
	case dia.SpiritswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialFantom, relDB, spiritswapWaitMilliseconds)
	case dia.SolarbeamExchange:
		uas = makeUniswapAssetSource(exchange, restDialMoonriver, relDB, solarbeamWaitMilliseconds)
	case dia.TrisolarisExchange:
		uas = makeUniswapAssetSource(exchange, restDialAurora, relDB, trisolarisWaitMilliseconds)
	case dia.NetswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialMetis, relDB, metisWaitMilliseconds)
	case dia.HuckleberryExchange:
		uas = makeUniswapAssetSource(exchange, restDialMoonriver, relDB, moonriverWaitMilliseconds)
	case dia.TraderJoeExchange:
		uas = makeUniswapAssetSource(exchange, restDialAvalanche, relDB, avalancheWaitMilliseconds)
	case dia.PangolinExchange:
		uas = makeUniswapAssetSource(exchange, restDialAvalanche, relDB, avalancheWaitMilliseconds)
	case dia.TethysExchange:
		uas = makeUniswapAssetSource(exchange, restDialMetis, relDB, metisWaitMilliseconds)
	case dia.HermesExchange:
		uas = makeUniswapAssetSource(exchange, restDialMetis, relDB, metisWaitMilliseconds)
	case dia.OmniDexExchange:
		uas = makeUniswapAssetSource(exchange, restDialTelos, relDB, telosWaitMilliseconds)
	case dia.DiffusionExchange:
		uas = makeUniswapAssetSource(exchange, restDialEvmos, relDB, evmosWaitMilliseconds)
	case dia.ArthswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialAstar, relDB, astarWaitMilliseconds)
	case dia.StellaswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialMoonbeam, relDB, moonbeamWaitMilliseconds)
	case dia.WanswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialWanchain, relDB, wanchainWaitMilliseconds)
	case dia.PearlfiExchangeTestnet:
		uas = makeUniswapAssetSource(exchange, restDialUnrealTestnet, relDB, wanchainWaitMilliseconds)
	case dia.PearlfiExchange:
		uas = makeUniswapAssetSource(exchange, restDialUnreal, relDB, wanchainWaitMilliseconds)
	case dia.PearlfiStableswapExchange:
		uas = makeUniswapAssetSource(exchange, restDialUnreal, relDB, wanchainWaitMilliseconds)
	case dia.RamsesV1Exchange:
		uas = makeUniswapAssetSource(exchange, restDialArbitrum, relDB, wanchainWaitMilliseconds)
	case dia.NileV1Exchange:
		uas = makeUniswapAssetSource(exchange, restDialLinea, relDB, wanchainWaitMilliseconds)
	case dia.ThenaExchange:
		uas = makeUniswapAssetSource(exchange, restDialBSC, relDB, sushiswapWaitMilliseconds)
	case dia.ShadowV2Exchange:
		uas = makeUniswapAssetSource(exchange, restDialSonic, relDB, baseWaitMilliseconds)
	}

	exchangeFactoryContractAddress = exchange.Contract

	go func() {
		uas.fetchAssets()
	}()
	return uas

}

// makeUniswapAssetSource returns an asset source as used in NewUniswapAssetSource.
func makeUniswapAssetSource(exchange dia.Exchange, restDial string, relDB *models.RelDB, waitMilliseconds string) *UniswapAssetSource {
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
		relDB:        relDB,
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		exchange:     exchange,
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

	_, index, err := uas.relDB.GetScraperIndex(uas.exchange.Name, dia.SCRAPER_TYPE_ASSETCOLLECTOR)
	if err != nil {
		log.Error("GetScraperIndex: ", err)
	}
	log.Info("Start at index ", index)

	for index < int64(numPairs) {
		log.Info("index: ", index)
		time.Sleep(time.Duration(uas.waitTime) * time.Millisecond)
		pair, err := uas.GetPairByID(index)
		if err != nil {
			log.Errorln("Error getting pair with ID ", index)
		}
		asset0 := pair.Token0
		asset1 := pair.Token1
		asset0.Blockchain = uas.exchange.BlockChain.Name
		asset1.Blockchain = uas.exchange.BlockChain.Name
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
		index++
	}
	err = uas.relDB.SetScraperIndex(uas.exchange.Name, dia.SCRAPER_TYPE_ASSETCOLLECTOR, dia.INDEX_TYPE_INDEX, index)
	if err != nil {
		log.Error("SetScraperIndex: ", err)
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
