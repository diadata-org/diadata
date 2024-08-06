package liquidityscrapers

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	traderjoe "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/traderjoe2.1"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/traderjoe2.1/traderjoeILBPair"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TraderJoeLiquidityScraper manages the scraping of liquidity data for the Trader Joe exchange.
type TraderJoeLiquidityScraper struct {
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

// NewTraderJoeLiquidityScraper initializes a Trader Joe liquidity scraper, creating an instance of the
// TraderJoeLiquidityScraper struct. It configures necessary parameters, initiates pool fetching, and returns
// the initialized scraper.
func NewTraderJoeLiquidityScraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *TraderJoeLiquidityScraper {
	log.Info("NewTraderJoeLiquidityScraper ", exchange.Name)
	log.Info("NewTraderJoeLiquidityScraper Address ", exchange.Contract)

	var tjls *TraderJoeLiquidityScraper

	switch exchange.Name {
	case dia.TraderJoeExchangeV2_1:
		tjls = makeTraderJoeScraper(exchange, "", "", relDB, datastore, "200", uint64(17821282))
	case dia.TraderJoeExchangeV2_1Arbitrum:
		tjls = makeTraderJoeScraper(exchange, "", "", relDB, datastore, "200", uint64(77473199))
	case dia.TraderJoeExchangeV2_1Avalanche:
		tjls = makeTraderJoeScraper(exchange, "", "", relDB, datastore, "200", uint64(28371397))
	case dia.TraderJoeExchangeV2_1BNB:
		tjls = makeTraderJoeScraper(exchange, "", "", relDB, datastore, "200", uint64(27099340))
	case dia.TraderJoeExchangeV2_2Avalanche:
		tjls = makeTraderJoeScraper(exchange, "", "", relDB, datastore, "200", uint64(46536129))
	}

	go func() {
		tjls.fetchPools()
	}()
	return tjls
}

// makeTraderJoeScraper initializes a Trader Joe liquidity scraper, creating an instance of the
// TraderJoeLiquidityScraper struct with the specified configuration and parameters.
func makeTraderJoeScraper(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB, datastore *models.DB, waitMilliSeconds string, startBlock uint64) *TraderJoeLiquidityScraper {
	var (
		restClient  *ethclient.Client
		wsClient    *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		tjls        *TraderJoeLiquidityScraper
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
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliSeconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	tjls = &TraderJoeLiquidityScraper{
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
	return tjls
}

// fetchPools retrieves pool creation events from the Trader Joe factory contract address and processes them.
func (tjls *TraderJoeLiquidityScraper) fetchPools() {

	log.Info("Fetching Trader Joe LBPairCreated events...")
	log.Info("Getting lb pairs creations from address: ", tjls.factoryContract)

	// Initialize a count for the number of pairs processed.
	pairCount := 0

	// Initialize an Ethereum event filter for the Trader Joe factory contract.
	contractFilter, err := traderjoe.NewTraderjoeFilterer(common.HexToAddress(tjls.factoryContract), tjls.WsClient)
	if err != nil {
		log.Error(err)
	}

	// Retrieve LBPairCreated events using the event filter.
	lbPairCreated, err := contractFilter.FilterLBPairCreated(
		&bind.FilterOpts{Start: tjls.startBlock},
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Error("filter pool created: ", err)
	}

	// Iterate through the LBPairCreated events.
	for lbPairCreated.Next() {
		pairCount++
		var (
			pool   dia.Pool
			asset0 dia.Asset
			asset1 dia.Asset
		)
		log.Info("pairs count: ", pairCount)
		time.Sleep(time.Duration(tjls.waitTime) * time.Millisecond)

		// Retrieve information about the first token of the liquidity pool.
		asset0, err = tjls.relDB.GetAsset(lbPairCreated.Event.TokenX.Hex(), tjls.blockchain)
		if err != nil {
			// If asset information cannot be retrieved from the database, try fetching it from the Ethereum network.
			asset0, err = ethhelper.ETHAddressToAsset(lbPairCreated.Event.TokenX, tjls.RestClient, tjls.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", lbPairCreated.Event.TokenX.Hex())
				continue
			}
		}

		// Retrieve information about the second token of the liquidity pool.
		asset1, err = tjls.relDB.GetAsset(lbPairCreated.Event.TokenY.Hex(), tjls.blockchain)
		if err != nil {
			// If asset information cannot be retrieved from the database, try fetching it from the Ethereum network.
			asset1, err = ethhelper.ETHAddressToAsset(lbPairCreated.Event.TokenY, tjls.RestClient, tjls.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", lbPairCreated.Event.TokenY.Hex())
				continue
			}
		}

		// Populate pool information.
		pool.Exchange = dia.Exchange{Name: tjls.exchangeName}
		pool.Blockchain = dia.BlockChain{Name: tjls.blockchain}
		pool.Address = lbPairCreated.Event.LBPair.Hex()

		pairFiltererContract, err := traderjoeILBPair.NewILBPairCaller(lbPairCreated.Event.LBPair, tjls.RestClient)
		if err != nil {
			log.Fatal(err)
		}
		reserves, err := pairFiltererContract.GetReserves(&bind.CallOpts{})
		if err != nil {
			log.Fatal("get reserves on pool ", lbPairCreated.Event.LBPair.Hex())
		}

		// Calculate token balances in floating-point format.
		balance0, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(reserves.ReserveX), new(big.Float).SetFloat64(math.Pow10(int(asset0.Decimals)))).Float64()
		balance1, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(reserves.ReserveY), new(big.Float).SetFloat64(math.Pow10(int(asset1.Decimals)))).Float64()

		// Append asset volumes to the pool.
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset0, Volume: balance0, Index: uint8(0)})
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset1, Volume: balance1, Index: uint8(1)})

		// Determine USD liquidity for the pool if both token balances meet the threshold.
		if balance0 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD && balance1 > GLOBAL_NATIVE_LIQUIDITY_THRESHOLD {
			tjls.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		// Set the timestamp for the pool.
		pool.Time = time.Now()

		// Send the processed pool information to the channel.
		tjls.poolChannel <- pool
	}

	// Signal that pool retrieval and processing is complete.
	tjls.doneChannel <- true
}

// Pool returns a channel for receiving dia.Pool instances scraped by the Trader Joe liquidity scraper.
func (tjls *TraderJoeLiquidityScraper) Pool() chan dia.Pool {
	return tjls.poolChannel
}

// Done returns a channel for signaling the completion of Trader Joe liquidity scraping.
func (tjls *TraderJoeLiquidityScraper) Done() chan bool {
	return tjls.doneChannel
}
