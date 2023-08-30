package liquidityscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
	traderjoe "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/traderjoe2.1"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"strings"
	"time"
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

	if exchange.Name == dia.TraderJoeExchange {
		tjls = makeTraderJoeScraper(exchange, "", "", relDB, datastore, "200", uint64(12369621))
		// TODO: startBlock value will need revisiting.
	}

	go func() {
		tjls.fetchPools()
	}()
	return tjls
}

// makeTraderJoeScraper initializes a Trader Joe liquidity scraper, creating an instance of the
// TraderJoeLiquidityScraper struct with the specified configuration and parameters.
func makeTraderJoeScraper(exchange dia.Exchange, restDial string, websocketDial string, relDB *models.RelDB, datastore *models.DB, waitMilliSeconds string, startBlock uint64) *TraderJoeLiquidityScraper {
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
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", websocketDial))
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
	log.Info("Getting pool creations from address: ", tjls.factoryContract)

	// Filtering setup initialization.
	contractFilter, err := traderjoe.NewTraderjoeFilterer(common.HexToAddress(tjls.factoryContract), tjls.WsClient)
	if err != nil {
		log.Error(err)
	}

	// Define filter options
	opts := &bind.FilterOpts{
		Start: tjls.startBlock,
	}

	// Define empty arrays for filter parameters
	var tokenXAddresses []common.Address
	var tokenYAddresses []common.Address
	var binSteps []*big.Int

	// Filter LBPairCreated events.
	iter, err := contractFilter.FilterLBPairCreated(opts, tokenXAddresses, tokenYAddresses, binSteps)
	if err != nil {
		log.Error("FilterLBPairCreated: ", err)
		return
	}
	defer func(iter *traderjoe.TraderjoeLBPairCreatedIterator) {
		err := iter.Close()
		if err != nil {
			log.Error(err)
		}
	}(iter)

	for iter.Next() {
		var pool dia.Pool

		// Parse event data from iter.Event() .
		event := iter.Event
		tokenX := event.TokenX
		tokenY := event.TokenY
		//binStep := event.BinStep

		// Fetch pool-specific data using Trader Joe functions.
		contractCaller, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(tjls.factoryContract), tjls.WsClient)
		if err != nil {
			log.Error(err)
			continue
		}

		// lbPairInfo for getting all LB pairs. rename _ to lbPairInfo
		_, err = contractCaller.GetAllLBPairs(nil, tokenX, tokenY)
		if err != nil {
			log.Error("GetAllLBPairs: ", err)
			continue
		}

		// TODO: Append 'pool' to 'tjs.PoolChannel' correctly.
		pool = dia.Pool{
			Exchange:     dia.Exchange{},
			Blockchain:   dia.BlockChain{},
			Address:      "",
			Assetvolumes: nil,
			Time:         time.Now(),
		}

		log.Info("Fetched pool data: ", pool)

		tjls.poolChannel <- pool
	}
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
