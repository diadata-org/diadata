package source

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/velodrome"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VelodromePair struct {
	Token0      dia.Asset
	Token1      dia.Asset
	ForeignName string
	Address     common.Address
}

const (
	velodromeWaitMilliseconds = "500"
)

var velodromeExchangeFactoryContractAddress string

type VelodromeAssetSource struct {
	RestClient   *ethclient.Client
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
	waitTime     int
}

func NewVelodromeAssetSource(exchange dia.Exchange) (uas *VelodromeAssetSource) {
	switch exchange.Name {
	case dia.VelodromeExchange:
		uas = makeVelodromeAssetSource(exchange, exchange.RestAPI, velodromeWaitMilliseconds)
	}

	velodromeExchangeFactoryContractAddress = exchange.Contract

	go func() {
		uas.fetchAssets()
	}()
	return uas

}

// makeVelodromeAssetSource returns an asset source as used in NewVelodromeAssetSource.
func makeVelodromeAssetSource(exchange dia.Exchange, restDial string, waitMilliseconds string) *VelodromeAssetSource {
	var (
		restClient   *ethclient.Client
		err          error
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		uas          *VelodromeAssetSource
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
	uas = &VelodromeAssetSource{
		RestClient:   restClient,
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		waitTime:     waitTime,
	}
	return uas
}

func (uas *VelodromeAssetSource) Asset() chan dia.Asset {
	return uas.assetChannel
}

func (uas *VelodromeAssetSource) Done() chan bool {
	return uas.doneChannel
}

func (uas *VelodromeAssetSource) getNumPairs() (int, error) {
	var contract *velodrome.PoolFactoryCaller
	contract, err := velodrome.NewPoolFactoryCaller(common.HexToAddress(velodromeExchangeFactoryContractAddress), uas.RestClient)
	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.AllPoolsLength(&bind.CallOpts{})
	return int(numPairs.Int64()), err
}

func (uas *VelodromeAssetSource) fetchAssets() {

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

// GetPairByID returns the VelodromePair with the integer id @num
func (uas *VelodromeAssetSource) GetPairByID(num int64) (VelodromePair, error) {
	var contract *velodrome.PoolFactoryCaller
	contract, err := velodrome.NewPoolFactoryCaller(common.HexToAddress(velodromeExchangeFactoryContractAddress), uas.RestClient)
	if err != nil {
		log.Error(err)
		return VelodromePair{}, err
	}
	numToken := big.NewInt(num)
	pairAddress, err := contract.AllPools(&bind.CallOpts{}, numToken)
	if err != nil {
		log.Error(err)
		return VelodromePair{}, err
	}

	pair, err := uas.GetPairByAddress(pairAddress)
	if err != nil {
		log.Error(err)
		return VelodromePair{}, err
	}
	// log.Infof("Get pair with ID %v: %v", num, pair)
	return pair, err
}

func (uas *VelodromeAssetSource) GetPairByAddress(pairAddress common.Address) (pair VelodromePair, err error) {
	connection := uas.RestClient
	var pairContract *velodrome.IPoolCaller
	pairContract, err = velodrome.NewIPoolCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return VelodromePair{}, err
	}

	// Getting tokens from pair ---------------------
	address0, _ := pairContract.Token0(&bind.CallOpts{})
	address1, _ := pairContract.Token1(&bind.CallOpts{})
	var token0Contract *velodrome.IERC20MetadataCaller
	var token1Contract *velodrome.IERC20MetadataCaller
	token0Contract, err = velodrome.NewIERC20MetadataCaller(address0, connection)
	if err != nil {
		log.Error(err)
	}
	token1Contract, err = velodrome.NewIERC20MetadataCaller(address1, connection)
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
		return VelodromePair{}, err
	}
	name1, err := uas.GetName(address1)
	if err != nil {
		log.Error(err)
		return VelodromePair{}, err
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
func (uas *VelodromeAssetSource) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *velodrome.IERC20MetadataCaller
	contract, err = velodrome.NewIERC20MetadataCaller(tokenAddress, uas.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (uas *VelodromeAssetSource) GetName(tokenAddress common.Address) (name string, err error) {

	var contract *velodrome.IERC20MetadataCaller
	contract, err = velodrome.NewIERC20MetadataCaller(tokenAddress, uas.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}
