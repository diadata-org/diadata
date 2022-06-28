package liquidityscrapers

import (
	"math"
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
	restDialEthereum  = ""
	restDialBSC       = ""
	restDialPolygon   = ""
	restDialCelo      = ""
	restDialFantom    = ""
	restDialMoonriver = ""
	restDialAurora    = ""
	restDialMetis     = ""
	restDialAvalanche = ""
	restDialTelos     = ""
	restDialEvmos     = ""
	restDialAstar     = ""

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
)

type UniswapScraper struct {
	RestClient   *ethclient.Client
	poolChannel  chan dia.Pool
	doneChannel  chan bool
	blockchain   string
	waitTime     int
	exchangeName string
}

var exchangeFactoryContractAddress string

func NewUniswapScraper(exchange dia.Exchange) (us *UniswapScraper) {

	switch exchange.Name {
	case dia.UniswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialEthereum, uniswapWaitMilliseconds)
	case dia.SushiSwapExchange:
		us = makeUniswapPoolScraper(exchange, restDialEthereum, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangePolygon:
		us = makeUniswapPoolScraper(exchange, restDialPolygon, sushiswapWaitMilliseconds)
	case dia.SushiSwapExchangeFantom:
		us = makeUniswapPoolScraper(exchange, restDialFantom, sushiswapWaitMilliseconds)
	case dia.PanCakeSwap:
		us = makeUniswapPoolScraper(exchange, restDialBSC, pancakeswapWaitMilliseconds)
	case dia.DfynNetwork:
		us = makeUniswapPoolScraper(exchange, restDialPolygon, dfynWaitMilliseconds)
	case dia.QuickswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialPolygon, dfynWaitMilliseconds)
	case dia.UbeswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialCelo, ubeswapWaitMilliseconds)
	case dia.SpookyswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialFantom, spookyswapWaitMilliseconds)
	case dia.SpiritswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialFantom, spiritswapWaitMilliseconds)
	case dia.SolarbeamExchange:
		us = makeUniswapPoolScraper(exchange, restDialMoonriver, solarbeamWaitMilliseconds)
	case dia.TrisolarisExchange:
		us = makeUniswapPoolScraper(exchange, restDialAurora, trisolarisWaitMilliseconds)
	case dia.NetswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialMetis, metisWaitMilliseconds)
	case dia.HuckleberryExchange:
		us = makeUniswapPoolScraper(exchange, restDialMoonriver, moonriverWaitMilliseconds)
	case dia.TraderJoeExchange:
		us = makeUniswapPoolScraper(exchange, restDialAvalanche, avalancheWaitMilliseconds)
	case dia.PangolinExchange:
		us = makeUniswapPoolScraper(exchange, restDialAvalanche, avalancheWaitMilliseconds)
	case dia.TethysExchange:
		us = makeUniswapPoolScraper(exchange, restDialMetis, metisWaitMilliseconds)
	case dia.HermesExchange:
		us = makeUniswapPoolScraper(exchange, restDialMetis, metisWaitMilliseconds)
	case dia.OmniDexExchange:
		us = makeUniswapPoolScraper(exchange, restDialTelos, telosWaitMilliseconds)
	case dia.DiffusionExchange:
		us = makeUniswapPoolScraper(exchange, restDialEvmos, evmosWaitMilliseconds)
	case dia.ArthswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialAstar, astarWaitMilliseconds)
	case dia.ApeswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialAstar, astarWaitMilliseconds)
	case dia.BiswapExchange:
		us = makeUniswapPoolScraper(exchange, restDialAstar, astarWaitMilliseconds)
	}

	exchangeFactoryContractAddress = exchange.Contract

	go func() {
		us.fetchPools()
	}()
	return us

}

// makeUniswapPoolScraper returns an asset source as used in NewUniswapAssetSource.
func makeUniswapPoolScraper(exchange dia.Exchange, restDial string, waitMilliseconds string) *UniswapScraper {
	var restClient *ethclient.Client
	var err error
	var poolChannel = make(chan dia.Pool)
	var doneChannel = make(chan bool)
	var us *UniswapScraper
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
	us = &UniswapScraper{
		RestClient:   restClient,
		poolChannel:  poolChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		waitTime:     waitTime,
		exchangeName: exchange.Name,
	}
	return us
}

// fetchPools iterates through all (Uniswap) pools and sends them into the pool channel.
func (us *UniswapScraper) fetchPools() {

	numPairs, err := us.getNumPairs()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Found ", numPairs, " pools")

	for i := 0; i < numPairs; i++ {
		time.Sleep(time.Duration(us.waitTime) * time.Millisecond)
		pool, err := us.GetPoolByID(int64(numPairs - 1 - i))
		if err != nil {
			log.Errorln("Error getting pair with ID ", numPairs-1-i)
		}
		log.Info("found pool: ", pool)
		us.poolChannel <- pool
	}
	us.doneChannel <- true
}

// GetPoolByID returns the Uniswap Pool with the integer id @num.
func (us *UniswapScraper) GetPoolByID(num int64) (dia.Pool, error) {
	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), us.RestClient)
	if err != nil {
		log.Error(err)
		return dia.Pool{}, err
	}
	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPairs(&bind.CallOpts{}, numToken)
	if err != nil {
		log.Error(err)
		return dia.Pool{}, err
	}

	pool, err := us.GetPoolByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return dia.Pool{}, err
	}

	return pool, err
}

func (us *UniswapScraper) GetPoolByAddress(pairAddress common.Address) (pool dia.Pool, err error) {
	connection := us.RestClient
	var pairContract *uniswap.IUniswapV2PairCaller
	pairContract, err = uniswap.NewIUniswapV2PairCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return dia.Pool{}, err
	}

	// Getting tokens from pair
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

	name0, err := us.GetName(address0)
	if err != nil {
		log.Error("get name: ", err)
	}
	name1, err := us.GetName(address1)
	if err != nil {
		log.Error("get name: ", err)
	}
	token0 := dia.Asset{
		Address:    address0.Hex(),
		Blockchain: us.blockchain,
		Symbol:     symbol0,
		Name:       name0,
		Decimals:   decimals0,
	}
	token1 := dia.Asset{
		Address:    address1.Hex(),
		Blockchain: us.blockchain,
		Symbol:     symbol1,
		Name:       name1,
		Decimals:   decimals1,
	}

	// Getting liquidity
	liquidity, err := pairContract.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Error("get reserves: ", err)
	}
	amount0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(liquidity.Reserve0), new(big.Float).SetFloat64(math.Pow10(int(token0.Decimals)))).Float64()
	amount1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(liquidity.Reserve1), new(big.Float).SetFloat64(math.Pow10(int(token1.Decimals)))).Float64()

	// TO DO: Fetch timestamp using block number?
	pool.Time = time.Now()

	// Fill Pool type with the above data

	// assetvolumeMap[token0] = amount0
	// assetvolumeMap[token1] = amount1
	pool.Assetvolumes = append(pool.Assetvolumes, struct {
		Asset  dia.Asset
		Volume float64
	}{
		Asset:  token0,
		Volume: amount0,
	})
	pool.Assetvolumes = append(pool.Assetvolumes, struct {
		Asset  dia.Asset
		Volume float64
	}{
		Asset:  token1,
		Volume: amount1,
	})
	pool.Address = pairAddress.Hex()
	pool.Blockchain = dia.BlockChain{Name: us.blockchain}
	pool.Exchange = dia.Exchange{Name: us.exchangeName}

	return pool, nil
}

// GetDecimals returns the decimals of the token with address @tokenAddress
func (us *UniswapScraper) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, us.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (us *UniswapScraper) GetName(tokenAddress common.Address) (name string, err error) {

	var contract *uniswap.IERC20Caller
	contract, err = uniswap.NewIERC20Caller(tokenAddress, us.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}

func (us *UniswapScraper) Pool() chan dia.Pool {
	return us.poolChannel
}

func (us *UniswapScraper) Done() chan bool {
	return us.doneChannel
}

func (us *UniswapScraper) getNumPairs() (int, error) {
	var contract *uniswap.IUniswapV2FactoryCaller
	contract, err := uniswap.NewIUniswapV2FactoryCaller(common.HexToAddress(exchangeFactoryContractAddress), us.RestClient)
	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.AllPairsLength(&bind.CallOpts{})
	return int(numPairs.Int64()), err
}
