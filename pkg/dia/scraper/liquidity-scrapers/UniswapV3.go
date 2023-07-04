package liquidityscrapers

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	uniswapcontractv3 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv3"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapV3Scraper struct {
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

// NewUniswapV3Scraper returns a new UniswapV3Scraper.
func NewUniswapV3Scraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *UniswapV3Scraper {
	log.Info("NewUniswapScraper ", exchange.Name)
	log.Info("NewUniswapScraper Address ", exchange.Contract)

	var uls *UniswapV3Scraper

	switch exchange.Name {
	case dia.UniswapExchangeV3:
		uls = makeUniswapV3Scraper(exchange, "", "", relDB, datastore, "200", uint64(12369621))
	case dia.UniswapExchangeV3Polygon:
		uls = makeUniswapV3Scraper(exchange, "", "", relDB, datastore, "200", uint64(22757913))
	case dia.UniswapExchangeV3Arbitrum:
		uls = makeUniswapV3Scraper(exchange, "", "", relDB, datastore, "200", uint64(165))
	}

	go func() {
		uls.fetchPools()
	}()
	return uls
}

// makeUniswapV3Scraper returns a uniswap scraper as used in NewUniswapV3Scraper.
func makeUniswapV3Scraper(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB, datastore *models.DB, waitMilliseconds string, startBlock uint64) *UniswapV3Scraper {
	var (
		restClient  *ethclient.Client
		wsClient    *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		uls         *UniswapV3Scraper
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

	uls = &UniswapV3Scraper{
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
	return uls
}

// fetchPools fetches all registered pools from on-chain and sends them into the pool channel.
func (uls *UniswapV3Scraper) fetchPools() {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603

	log.Info("get pool creations from address: ", uls.factoryContract)
	poolsCount := 0
	contract, err := uniswapcontractv3.NewUniswapV3Filterer(common.HexToAddress(uls.factoryContract), uls.WsClient)
	if err != nil {
		log.Error(err)
	}

	poolCreated, err := contract.FilterPoolCreated(
		&bind.FilterOpts{Start: uls.startBlock},
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
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

		asset0, err = uls.relDB.GetAsset(poolCreated.Event.Token0.Hex(), uls.blockchain)
		if err != nil {
			asset0, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Token0, uls.RestClient, uls.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolCreated.Event.Token0.Hex())
				continue
			}
		}
		asset1, err = uls.relDB.GetAsset(poolCreated.Event.Token1.Hex(), uls.blockchain)
		if err != nil {
			asset1, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Token1, uls.RestClient, uls.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolCreated.Event.Token1.Hex())
				continue
			}
		}

		pool.Exchange = dia.Exchange{Name: uls.exchangeName}
		pool.Blockchain = dia.BlockChain{Name: uls.blockchain}
		pool.Address = poolCreated.Event.Pool.Hex()

		balance0Big, err := ethhelper.GetBalanceOf(common.HexToAddress(asset0.Address), common.HexToAddress(pool.Address), uls.RestClient)
		if err != nil {
			log.Error("GetBalanceOf: ", err)
		}
		balance1Big, err := ethhelper.GetBalanceOf(common.HexToAddress(asset1.Address), common.HexToAddress(pool.Address), uls.RestClient)
		if err != nil {
			log.Error("GetBalanceOf: ", err)
		}
		balance0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance0Big), new(big.Float).SetFloat64(math.Pow10(int(asset0.Decimals)))).Float64()
		balance1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance1Big), new(big.Float).SetFloat64(math.Pow10(int(asset1.Decimals)))).Float64()

		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset0, Volume: balance0, Index: uint8(0)})
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset1, Volume: balance1, Index: uint8(1)})

		// Determine USD liquidity
		if balance0 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD && balance1 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD {
			uls.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		pool.Time = time.Now()

		uls.poolChannel <- pool

	}
	uls.doneChannel <- true
}

func (uas *UniswapV3Scraper) Pool() chan dia.Pool {
	return uas.poolChannel
}

func (uas *UniswapV3Scraper) Done() chan bool {
	return uas.doneChannel
}
