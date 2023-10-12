package liquidityscrapers

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	camelotv3 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/camelotv3"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CamelotV3Scraper struct {
	RestClient      *ethclient.Client
	WsClient        *ethclient.Client
	relDB           *models.RelDB
	datastore       *models.DB
	poolChannel     chan dia.Pool
	doneChannel     chan bool
	blockchain      string
	startBlock      uint64
	factoryContract string
	exchangeName    string
	waitTime        int
}

// NewCamelotV3Scraper returns a new CamelotV3Scraper.
func NewCamelotV3Scraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *CamelotV3Scraper {
	log.Info("NewCamelotScraper ", exchange.Name)
	log.Info("NewCamelotScraper Address ", exchange.Contract)

	var cls *CamelotV3Scraper

	switch exchange.Name {
	case dia.CamelotExchangeV3:
		cls = makeCamelotV3Scraper(exchange, "", "", relDB, datastore, "200", uint64(101163738))
	}

	go func() {
		cls.fetchPools()
	}()
	return cls
}

// makeCamelotV3Scraper returns a camelot scraper as used in NewCamelotV3Scraper.
func makeCamelotV3Scraper(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB, datastore *models.DB, waitMilliseconds string, startBlock uint64) *CamelotV3Scraper {
	var (
		restClient  *ethclient.Client
		wsClient    *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		cls         *CamelotV3Scraper
	)

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init ws client: ", err)
	}

	var waitTime int
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	cls = &CamelotV3Scraper{
		WsClient:        wsClient,
		RestClient:      restClient,
		relDB:           relDB,
		datastore:       datastore,
		poolChannel:     poolChannel,
		doneChannel:     doneChannel,
		blockchain:      exchange.BlockChain.Name,
		startBlock:      startBlock,
		factoryContract: exchange.Contract,
		exchangeName:    exchange.Name,
		waitTime:        waitTime,
	}
	return cls
}

// fetchPools fetches all registered pools from on-chain and sends them into the pool channel.
func (cls *CamelotV3Scraper) fetchPools() {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603

	log.Info("get pool creations from address: ", cls.factoryContract)
	poolsCount := 0
	contract, err := camelotv3.NewCamelotv3Filterer(common.HexToAddress(cls.factoryContract), cls.WsClient)
	if err != nil {
		log.Error(err)
	}

	poolCreated, err := contract.FilterPool(
		&bind.FilterOpts{Start: cls.startBlock},
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		log.Error("filter pool created: ", err)
	}

	for poolCreated.Next() {
		poolsCount++
		var (
			pool   dia.Pool
			asset0 dia.Asset
			asset1 dia.Asset
		)
		log.Info("pools count: ", poolsCount)

		asset0, err = cls.relDB.GetAsset(poolCreated.Event.Token0.Hex(), cls.blockchain)
		if err != nil {
			asset0, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Token0, cls.RestClient, cls.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolCreated.Event.Token0.Hex())
				continue
			}
		}
		asset1, err = cls.relDB.GetAsset(poolCreated.Event.Token1.Hex(), cls.blockchain)
		if err != nil {
			asset1, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Token1, cls.RestClient, cls.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolCreated.Event.Token1.Hex())
				continue
			}
		}

		pool.Exchange = dia.Exchange{Name: cls.exchangeName}
		pool.Blockchain = dia.BlockChain{Name: cls.blockchain}
		pool.Address = poolCreated.Event.Pool.Hex()

		balance0Big, err := ethhelper.GetBalanceOf(common.HexToAddress(asset0.Address), common.HexToAddress(pool.Address), cls.RestClient)
		if err != nil {
			log.Error("GetBalanceOf: ", err)
		}
		balance1Big, err := ethhelper.GetBalanceOf(common.HexToAddress(asset1.Address), common.HexToAddress(pool.Address), cls.RestClient)
		if err != nil {
			log.Error("GetBalanceOf: ", err)
		}
		balance0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0Big), new(big.Float).SetFloat64(math.Pow10(int(asset0.Decimals)))).Float64()
		balance1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1Big), new(big.Float).SetFloat64(math.Pow10(int(asset1.Decimals)))).Float64()

		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset0, Volume: balance0, Index: uint8(0)})
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset1, Volume: balance1, Index: uint8(1)})

		// Determine USD liquidity
		if balance0 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD && balance1 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD {
			cls.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		pool.Time = time.Now()

		cls.poolChannel <- pool

	}
	cls.doneChannel <- true
}

func (cls *CamelotV3Scraper) Pool() chan dia.Pool {
	return cls.poolChannel
}

func (cls *CamelotV3Scraper) Done() chan bool {
	return cls.doneChannel
}
