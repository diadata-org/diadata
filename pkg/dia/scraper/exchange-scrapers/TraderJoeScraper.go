package scrapers

import (
	"encoding/json"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	TraderJoeExchangeFactoryContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

	//reverseBasetokens              *[]string
	//reverseQuotetokens             *[]string
	//mainBaseAssets                 = []string{
	//	"0xdAC17F958D2ee523a2206206994597C13D831ec7",
	//}

	MapOfPools = make(map[string]TraderJoePair)
)

// TODO: This needed?

type TraderJoeTokens struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
	Name     string
}

type TraderJoePair struct {
	Token0      TraderJoeTokens
	Token1      TraderJoeTokens
	ForeignName string
	Address     common.Address
}

type TraderJoeData struct {
	ID        string
	Timestamp int64
	Pair      TraderJoePair
	Amount0   float64
	Amount1   float64
	AmountIn  float64
	AmountOut float64
}

type TraderJoeScraper struct {
	WsClient   *ethclient.Client
	RestClient *ethclient.Client
	relDB      *models.RelDB
	// signaling channels for session initialization and finishing
	//initDone     chan nothing
	run          bool
	shutdown     chan nothing
	shutdownDone chan nothing
	// error handling; to read error or closed, first acquire read lock
	// only cleanup method should hold write lock
	errorLock sync.RWMutex
	error     error
	closed    bool
	// used to keep track of trading pairs that we subscribed to
	pairScrapers map[string]*TraderJoeScraper
	pairReceived chan *TraderJoePair

	exchangeName           string
	startBlock             uint64
	waitTime               int
	listenByAddress        bool
	chanTrades             chan *dia.Trade
	factoryContractAddress common.Address
}

func NewTraderJoeScraper(exchange dia.Exchange, scrape bool, relDB *models.RelDB) *TraderJoeScraper {
	log.Info("NewTraderJoeScraper ", exchange.Name)
	log.Info("NewTraderJoeScraper Address ", exchange.Contract)

	var (
		tjs             *TraderJoeScraper
		listenByAddress bool
		err             error
	)

	listenByAddress, err = strconv.ParseBool(utils.Getenv("LISTEN_BY_ADDRESS", ""))
	if err != nil {
		log.Fatal("parse LISTEN_BY_ADDRESS: ", err)
	}

	switch exchange.Name {
	case dia.TraderJoeExchange:
		tjs = makeTraderJoeScraper(exchange, listenByAddress, "", "", "200", uint64(12369621))
		// TODO: Are these uint64 values to be the same?
	}

	tjs.relDB = relDB

	// Only include pools with (minimum) liquidity bigger than given env var.
	liquidityThreshold, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD", "0"), 64)
	if err != nil {
		liquidityThreshold = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThreshold)
	}
	liquidityThresholdUSD, err := strconv.ParseFloat(utils.Getenv("LIQUIDITY_THRESHOLD_USD", "0"), 64)
	if err != nil {
		liquidityThresholdUSD = float64(0)
		log.Warnf("parse liquidity threshold:  %v. Set to default %v", err, liquidityThresholdUSD)
	}

	MapOfPools, err = tjs.makeTraderJoePoolMap(liquidityThreshold, liquidityThresholdUSD)
	if err != nil {
		log.Fatal("build poolMap: ", err)
	}

	if scrape {
		go tjs.mainLoop()
	}
	return tjs
}

func makeTraderJoeScraper(exchange dia.Exchange, listenByAddress bool, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *TraderJoeScraper {
	var restClient, wsClient *ethclient.Client
	var err error
	var s *TraderJoeScraper

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

	s = &TraderJoeScraper{
		WsClient:               wsClient,
		RestClient:             restClient,
		shutdown:               make(chan nothing),
		shutdownDone:           make(chan nothing),
		pairScrapers:           make(map[string]*TraderJoeScraper),
		exchangeName:           exchange.Name,
		pairReceived:           make(chan *TraderJoePair),
		error:                  nil,
		chanTrades:             make(chan *dia.Trade),
		waitTime:               waitTime,
		listenByAddress:        listenByAddress,
		startBlock:             startBlock,
		factoryContractAddress: common.HexToAddress(exchange.Contract),
	}

	return s
}

func (s *TraderJoeScraper) makeTraderJoePoolMap(liquidityThreshold, liquidityThresholdUSD float64) (map[string]TraderJoePair, error) {
	poolMap := make(map[string]TraderJoePair)
	var (
		pools []dia.Pool
		err   error
	)

	if s.listenByAddress {
		// Only load pool info for addresses from json file.
		poolAddresses, errAddr := getTradeAddressesFromConfig("traderjoe/subscribe_pools/" + s.exchangeName)
		// TODO: is this address correct?
		if errAddr != nil {
			log.Error("fetch pool addresses from config file: ", errAddr)
		}
		for _, address := range poolAddresses {
			pool, errPool := s.relDB.GetPoolByAddress(Exchanges[s.exchangeName].BlockChain.Name, address.Hex())
			if errPool != nil {
				log.Fatalf("Get pool with address %s: %v", address.Hex(), errPool)
			}
			pools = append(pools, pool)
		}
	} else {
		// Load all pools above liquidity threshold.
		pools, err = s.relDB.GetAllPoolsExchange(s.exchangeName, liquidityThreshold)
		if err != nil {
			return poolMap, err
		}
	}

	log.Info("Found ", len(pools), " pools.")
	log.Info("make pool map...")
	lowerBoundCount := 0
	for _, pool := range pools {
		if len(pool.Assetvolumes) != 2 {
			continue
		}
		liquidity, lowerBound := pool.GetPoolLiquidityUSD()
		// Discard pool if complete USD liquidity is below threshold.
		if !lowerBound && liquidity < liquidityThresholdUSD {
			continue
		}
		if lowerBound {
			lowerBoundCount++
		}

		up := TraderJoePair{
			Address: common.HexToAddress(pool.Address),
		}
		if pool.Assetvolumes[0].Index == 0 {
			up.Token0 = asset2TraderJoeAsset(pool.Assetvolumes[0].Asset)
			up.Token1 = asset2TraderJoeAsset(pool.Assetvolumes[1].Asset)
		} else {
			up.Token0 = asset2TraderJoeAsset(pool.Assetvolumes[1].Asset)
			up.Token1 = asset2TraderJoeAsset(pool.Assetvolumes[0].Asset)
		}
		up.ForeignName = up.Token0.Symbol + "-" + up.Token1.Symbol
		poolMap[pool.Address] = up
	}

	log.Infof("found %v subscribable pools.", len(poolMap))
	log.Infof("%v pools with lowerBound=true.", lowerBoundCount)

	return poolMap, err
}

func (s *TraderJoeScraper) mainLoop() {

}

func asset2TraderJoeAsset(asset dia.Asset) TraderJoeTokens {
	return TraderJoeTokens{
		Address:  common.HexToAddress(asset.Address),
		Decimals: asset.Decimals,
		Symbol:   asset.Symbol,
		Name:     asset.Name,
	}
}

func getTradeAddressesFromConfig(filename string) (pairAddresses []common.Address, err error) {

	// Load file and read data
	fileHandle := configCollectors.ConfigFileConnectors(filename, ".json")
	jsonFile, err := os.Open(fileHandle)
	if err != nil {
		return
	}
	defer func() {
		err = jsonFile.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		return
	}

	// Unmarshal read data
	type scrapedPair struct {
		Address     string `json:"Address"`
		ForeignName string `json:"ForeignName"`
	}
	type scrapedPairList struct {
		AllPairs []scrapedPair `json:"Pools"`
	}
	var allPairs scrapedPairList
	err = json.Unmarshal(byteData, &allPairs)
	if err != nil {
		return
	}

	// Extract addresses
	for _, token := range allPairs.AllPairs {
		pairAddresses = append(pairAddresses, common.HexToAddress(token.Address))
	}

	return
}
