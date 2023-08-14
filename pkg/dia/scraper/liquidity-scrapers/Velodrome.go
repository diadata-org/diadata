package liquidityscrapers

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/velodrome"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type VelodromePoolScraper struct {
	RestClient  *ethclient.Client
	relDB       *models.RelDB
	datastore   *models.DB
	poolChannel chan dia.Pool
	doneChannel chan bool
	blockchain  string
	waitTime    int
	exchange    dia.Exchange
}

var (
	velodromeWaitMilliseconds = "500"
	velodromeRestDial         = ""
)

func NewVelodromePoolScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) (us *VelodromePoolScraper) {

	switch exchange.Name {
	case dia.VelodromeExchange:
		us = makeVelodromePoolScraper(exchange, relDB, datastore, velodromeRestDial, velodromeWaitMilliseconds)
	}

	go func() {
		us.fetchPools()
	}()
	return us

}

func makeVelodromePoolScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB, restDial string, waitMilliseconds string) *VelodromePoolScraper {
	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		us          *VelodromePoolScraper
		waitTime    int
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	us = &VelodromePoolScraper{
		RestClient:  restClient,
		relDB:       relDB,
		datastore:   datastore,
		poolChannel: poolChannel,
		doneChannel: doneChannel,
		blockchain:  exchange.BlockChain.Name,
		waitTime:    waitTime,
		exchange:    exchange,
	}
	return us
}

// fetchPools iterates through all (Velodrome) pools and sends them into the pool channel.
func (us *VelodromePoolScraper) fetchPools() {
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

// GetPoolByID returns the Velodrome Pool with the integer id @num.
func (us *VelodromePoolScraper) GetPoolByID(num int64) (dia.Pool, error) {
	var contract *velodrome.PoolFactoryCaller

	contract, err := velodrome.NewPoolFactoryCaller(common.HexToAddress(us.exchange.Contract), us.RestClient)
	if err != nil {
		log.Error(err)
		return dia.Pool{}, err
	}

	pairAddress, err := contract.AllPools(&bind.CallOpts{}, big.NewInt(num))
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

// Get a pool by its LP token address.
func (us *VelodromePoolScraper) GetPoolByAddress(pairAddress common.Address) (pool dia.Pool, err error) {
	var (
		pairContract *velodrome.IPoolCaller
		token0       dia.Asset
		token1       dia.Asset
	)

	connection := us.RestClient
	pairContract, err = velodrome.NewIPoolCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return dia.Pool{}, err
	}

	// Getting tokens from pair
	address0, _ := pairContract.Token0(&bind.CallOpts{})
	address1, _ := pairContract.Token1(&bind.CallOpts{})

	// Only fetch assets from on-chain in case they are not in our DB.
	token0, err = us.relDB.GetAsset(address0.Hex(), us.blockchain)
	if err != nil {
		token0, err = ethhelper.ETHAddressToAsset(address0, us.RestClient, us.blockchain)
		if err != nil {
			return
		}
	}
	token1, err = us.relDB.GetAsset(address1.Hex(), us.blockchain)
	if err != nil {
		token1, err = ethhelper.ETHAddressToAsset(address1, us.RestClient, us.blockchain)
		if err != nil {
			return
		}
	}

	// Getting liquidity
	liquidity, err := pairContract.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Error("get reserves: ", err)
		return
	}
	amount0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(liquidity.Reserve0), new(big.Float).SetFloat64(math.Pow10(int(token0.Decimals)))).Float64()
	amount1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(liquidity.Reserve1), new(big.Float).SetFloat64(math.Pow10(int(token1.Decimals)))).Float64()

	// TO DO: Fetch timestamp using block number?
	pool.Time = time.Now()

	// Fill Pool type with the above data
	pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{
		Asset:  token0,
		Volume: amount0,
		Index:  uint8(0),
	})
	pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{
		Asset:  token1,
		Volume: amount1,
		Index:  uint8(1),
	})

	// Determine USD liquidity
	if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
		us.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
	}

	pool.Address = pairAddress.Hex()
	pool.Blockchain = dia.BlockChain{Name: us.blockchain}
	pool.Exchange = dia.Exchange{Name: us.exchange.Name}

	return pool, nil
}

// GetDecimals returns the decimals of the token with address @tokenAddress
func (us *VelodromePoolScraper) GetDecimals(tokenAddress common.Address) (decimals uint8, err error) {

	var contract *velodrome.IERC20MetadataCaller
	contract, err = velodrome.NewIERC20MetadataCaller(tokenAddress, us.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	decimals, err = contract.Decimals(&bind.CallOpts{})

	return
}

func (us *VelodromePoolScraper) GetName(tokenAddress common.Address) (name string, err error) {

	var contract *velodrome.IERC20MetadataCaller
	contract, err = velodrome.NewIERC20MetadataCaller(tokenAddress, us.RestClient)
	if err != nil {
		log.Error(err)
		return
	}
	name, err = contract.Name(&bind.CallOpts{})

	return
}

func (us *VelodromePoolScraper) Pool() chan dia.Pool {
	return us.poolChannel
}

func (us *VelodromePoolScraper) Done() chan bool {
	return us.doneChannel
}

func (us *VelodromePoolScraper) getNumPairs() (int, error) {
	var contract *velodrome.IPoolFactoryCaller
	contract, err := velodrome.NewIPoolFactoryCaller(common.HexToAddress(us.exchange.Contract), us.RestClient)
	if err != nil {
		log.Error(err)
	}

	numPairs, err := contract.AllPoolsLength(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return int(numPairs.Int64()), err
}
