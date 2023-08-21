package liquidityscrapers

import (
	"context"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	pairfactorycontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/maverick/pairfactory"
	poolcontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/maverick/pool"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	factoryContractAddressEth                = "0xEb6625D65a0553c9dBc64449e56abFe519bd9c9B"
	factoryContractAddressDeploymentBlockEth = uint64(17210221)
	defaultWaitMillis                        = "25"
)

type MaverickPool struct {
	Token0      dia.Asset
	Token1      dia.Asset
	fee         float64
	tickSpacing int64
	lookback    *big.Int
	ForeignName string
	Address     common.Address
}

type MaverickScraper struct {
	RestClient                       *ethclient.Client
	poolFactoryContractAddress       string
	poolFactoryContractCreationBlock uint64
	relDB                            *models.RelDB
	datastore                        *models.DB
	poolChannel                      chan dia.Pool
	doneChannel                      chan bool
	blockchain                       string
	waitTime                         int
	exchangeName                     string
	pathToPools                      string
}

func NewMaverickScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) (s *MaverickScraper) {

	pathToPools := utils.Getenv("PATH_TO_POOLS", "")

	switch exchange.Name {
	case dia.MaverickExchange:
		s = makeMaverickPoolScraper(exchange, pathToPools, exchange.RestAPI, relDB, datastore, defaultWaitMillis, factoryContractAddressEth, factoryContractAddressDeploymentBlockEth)
		//case dia.MaverickExchangeBNB:
		//	s = makeMaverickScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEth, wsDialEth, maverickWaitMilliseconds)
		//case dia.MaverickExchangeZKSync:
		//	s = makeMaverickScraper(exchange, listenByAddress, fetchPoolsFromDB, restDialEth, wsDialEth, maverickWaitMilliseconds)
	}

	exchangeFactoryContractAddress = exchange.Contract

	go func() {
		s.fetchPools()
	}()
	return s

}

func makeMaverickPoolScraper(exchange dia.Exchange, pathToPools string, restDial string, relDB *models.RelDB, datastore *models.DB, waitMilliseconds string, factoryContractAddess string, factoryContractDeploymentBlock uint64) *MaverickScraper {
	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		s           *MaverickScraper
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

	s = &MaverickScraper{
		RestClient:                       restClient,
		poolFactoryContractAddress:       factoryContractAddess,
		poolFactoryContractCreationBlock: factoryContractDeploymentBlock,
		relDB:                            relDB,
		datastore:                        datastore,
		poolChannel:                      poolChannel,
		doneChannel:                      doneChannel,
		blockchain:                       exchange.BlockChain.Name,
		waitTime:                         waitTime,
		exchangeName:                     exchange.Name,
		pathToPools:                      pathToPools,
	}
	return s
}

func (s *MaverickScraper) fetchPools() {
	if s.pathToPools != "" {

		// Collect all pool addresses from json file.
		poolAddresses, err := getAddressesFromConfig("liquidity-scrapers/maverick/" + s.pathToPools)
		if err != nil {
			log.Error("fetch pool addresses from config file: ", err)
		}
		numPairs := len(poolAddresses)
		log.Infof("listening to %d pools: %v", numPairs, poolAddresses)

		for _, pool := range poolAddresses {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			pool, err := s.getPoolByAddress(pool)
			if err != nil {
				log.Errorln("Error getting pool ", pool)
			}
			log.Info("found pool: ", pool)
			s.poolChannel <- pool
		}

	} else {

		pools, err := s.getAllPools()
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Found ", len(pools), " pools")

		for _, pool := range pools {
			time.Sleep(time.Duration(s.waitTime) * time.Millisecond)
			log.Info("found pool: ", pool)
			s.poolChannel <- pool
		}
	}
	s.doneChannel <- true
}

func (s *MaverickScraper) getPoolByAddress(pairAddress common.Address) (pool dia.Pool, err error) {
	var (
		poolContractInstance *poolcontract.PoolCaller
		token0               dia.Asset
		token1               dia.Asset
	)

	connection := s.RestClient
	poolContractInstance, err = poolcontract.NewPoolCaller(pairAddress, connection)
	if err != nil {
		log.Error(err)
		return dia.Pool{}, err
	}

	// Getting tokens from pair
	address0, _ := poolContractInstance.TokenA(&bind.CallOpts{})
	address1, _ := poolContractInstance.TokenB(&bind.CallOpts{})

	//log.Info(address0)
	//log.Info(address1)
	// Only fetch assets from on-chain in case they are not in our DB.
	token0, err = s.relDB.GetAsset(address0.Hex(), s.blockchain)
	if err != nil {
		token0, err = ethhelper.ETHAddressToAsset(address0, s.RestClient, s.blockchain)
		if err != nil {
			return
		}
	}
	token1, err = s.relDB.GetAsset(address1.Hex(), s.blockchain)
	if err != nil {
		token1, err = ethhelper.ETHAddressToAsset(address1, s.RestClient, s.blockchain)
		if err != nil {
			return
		}
	}

	// Getting liquidity

	balanceA, err := poolContractInstance.BinBalanceA(&bind.CallOpts{})
	balanceB, err := poolContractInstance.BinBalanceB(&bind.CallOpts{})

	if err != nil {
		log.Error("get reserves: ", err)
	}
	amount0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balanceA), new(big.Float).SetFloat64(math.Pow10(int(token0.Decimals)))).Float64()
	amount1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balanceB), new(big.Float).SetFloat64(math.Pow10(int(token1.Decimals)))).Float64()

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
		s.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
	}

	pool.Address = pairAddress.Hex()
	pool.Blockchain = dia.BlockChain{Name: s.blockchain}
	pool.Exchange = dia.Exchange{Name: s.exchangeName}

	return pool, nil
}

func (s *MaverickScraper) getAllPools() ([]dia.Pool, error) {
	pools := make([]dia.Pool, 0)

	var factoryContractInstance *pairfactorycontract.PairfactoryFilterer
	factoryContractInstance, err := pairfactorycontract.NewPairfactoryFilterer(common.HexToAddress(s.poolFactoryContractAddress), s.RestClient)
	if err != nil {
		log.Error(err)
		return pools, err
	}

	currBlock, err := s.RestClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	var offset uint64 = 2500
	startBlock := s.poolFactoryContractCreationBlock
	var endBlock = startBlock + offset

	for {
		if endBlock > currBlock {
			endBlock = currBlock
		}
		log.Infof("startblock - endblock: %v --- %v ", startBlock, endBlock)

		it, err := factoryContractInstance.FilterPoolCreated(
			&bind.FilterOpts{
				Start: startBlock,
				End:   &endBlock,
			})
		if err != nil {
			log.Error(err)
			//return pools, err
			if endBlock == currBlock {
				break
			}

			startBlock = endBlock + 1
			endBlock = endBlock + offset
			continue
		}

		for it.Next() {
			pool, err := s.getPoolByAddress(it.Event.PoolAddress)
			if err != nil {
				log.Warn(err)
			} else {
				pools = append(pools, pool)
			}
		}
		if err := it.Close(); err != nil {
			log.Warn("closing iterator: ", it)
		}

		if endBlock == currBlock {
			break
		}

		startBlock = endBlock + 1
		endBlock = endBlock + offset
	}

	return pools, err
}

func (s *MaverickScraper) Pool() chan dia.Pool {
	return s.poolChannel
}

func (s *MaverickScraper) Done() chan bool {
	return s.doneChannel
}
