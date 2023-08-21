package scrapers

import (
	"encoding/json"
	"errors"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/diadata-org/diadata/pkg/dia/helpers/configCollectors"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

// NewTraderJoeScraper returns a new TraderJoeScraper.
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

	if exchange.Name == dia.TraderJoeExchange {
		tjs = makeTraderJoeScraper(exchange, listenByAddress, "", "", "200", uint64(12369621))
		// TODO: startBlock value will need revisiting.
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

func (tjs *TraderJoeScraper) makeTraderJoePoolMap(liquidityThreshold, liquidityThresholdUSD float64) (map[string]TraderJoePair, error) {
	poolMap := make(map[string]TraderJoePair)
	var (
		pools []dia.Pool
		err   error
	)

	if tjs.listenByAddress {
		// Only load pool info for addresses from json file.
		poolAddresses, errAddr := getTradeAddressesFromConfig("traderjoe/subscribe_pools/" + tjs.exchangeName)
		// TODO: is this address correct?
		if errAddr != nil {
			log.Error("fetch pool addresses from config file: ", errAddr)
		}
		for _, address := range poolAddresses {
			pool, errPool := tjs.relDB.GetPoolByAddress(Exchanges[tjs.exchangeName].BlockChain.Name, address.Hex())
			if errPool != nil {
				log.Fatalf("Get pool with address %tjs: %v", address.Hex(), errPool)
			}
			pools = append(pools, pool)
		}
	} else {
		// Load all pools above liquidity threshold.
		pools, err = tjs.relDB.GetAllPoolsExchange(tjs.exchangeName, liquidityThreshold)
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

func (tjs *TraderJoeScraper) mainLoop() {
	var err error
	reverseBasetokens, err = getReverseTokensFromConfig("traderjoe/reverse_tokens/" + tjs.exchangeName + "Basetoken")
	if err != nil {
		log.Error("error getting base tokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following basetokens on %s: %v", tjs.exchangeName, reverseBasetokens)
	reverseQuotetokens, err = getReverseTokensFromConfig("traderjoe/reverse_tokens/" + tjs.exchangeName + "Quotetoken")
	if err != nil {
		log.Error("error getting quote tokens for which pairs should be reversed: ", err)
	}
	log.Infof("reverse the following quotetokens on %s: %v", tjs.exchangeName, reverseQuotetokens)

	time.Sleep(4 * time.Second)
	tjs.run = true

	go func() {
		pools := tjs.feedPoolsToSubscriptions()
		log.Info("Found ", len(pools), " pairs")
		log.Info("Found ", len(tjs.pairScrapers), " pairScrapers")
	}()

	if len(tjs.pairScrapers) == 0 {
		tjs.error = errors.New("uniswap: No pairs to scrape provided")
		log.Error(tjs.error.Error())
	}

	count := 0
	for {
		pool := <-tjs.pairReceived
		log.Infoln("Subscribing for pair: ", pool)

		if len(pool.Token0.Symbol) < 2 || len(pool.Token1.Symbol) < 2 {
			log.Info("skip pair: ", pool.ForeignName)
			continue
		}
		if helpers.AddressIsBlacklisted(pool.Token0.Address) || helpers.AddressIsBlacklisted(pool.Token1.Address) {
			log.Info("skip pair ", pool.ForeignName, ", address is blacklisted")
			continue
		}
		if helpers.PoolIsBlacklisted(pool.Address) {
			log.Info("skip blacklisted pool ", pool.Address)
			continue
		}
		log.Infof("%v found pair scraper for: %s with address %s", count, pool.ForeignName, pool.Address.Hex())
		count++
		// TODO: Paused here.
	}
}

func (tjs *TraderJoeScraper) sendTrade(tradeData TraderJoeData, pool *TraderJoePair) {
	price, volume := tjs.getTradeData(tradeData)
	token0 := dia.Asset{
		Address:    pool.Token0.Address.Hex(),
		Symbol:     pool.Token0.Symbol,
		Name:       pool.Token0.Name,
		Decimals:   pool.Token0.Decimals,
		Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
	}
	token1 := dia.Asset{
		Address:    pool.Token1.Address.Hex(),
		Symbol:     pool.Token1.Symbol,
		Name:       pool.Token1.Name,
		Decimals:   pool.Token1.Decimals,
		Blockchain: Exchanges[tjs.exchangeName].BlockChain.Name,
	}

	t := &dia.Trade{
		Symbol:         pool.Token0.Symbol,
		Pair:           pool.ForeignName,
		QuoteToken:     token1,
		BaseToken:      token0,
		Price:          price,
		Volume:         volume,
		Time:           time.Unix(tradeData.Timestamp, 0),
		PoolAddress:    pool.Address.Hex(),
		ForeignTradeID: tradeData.ID,
		//EstimatedUSDPrice: 0,
		Source:       tjs.exchangeName,
		VerifiedPair: true,
	}

	switch {
	case utils.Contains(reverseBasetokens, pool.Token1.Address.Hex()):
		// If we need quotation of a base token, reverse pair
		tSwapped, err := dia.SwapTrade(*t)
		if err == nil {
			t = &tSwapped
		}
	case utils.Contains(reverseQuotetokens, pool.Token0.Address.Hex()):
		// If we need quotation of a base token, reverse pair
		tSwapped, err := dia.SwapTrade(*t)
		if err == nil {
			t = &tSwapped
		}
	}
	if price > 0 {
		log.Infof("Got trade on pool %s: %v", pool.Address.Hex(), t)
		tjs.chanTrades <- t
	}
}

// TODO: Is GetSwapChannel necessary here?

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

func (tjs *TraderJoeScraper) feedPoolsToSubscriptions() (pairs []TraderJoePair) {
	for i := range MapOfPools {
		up := MapOfPools[i]
		pairs = append(pairs, up)
		tjs.pairReceived <- &up
	}
	return
}

func (tjs *TraderJoeScraper) getTradeData(swap TraderJoeData) (price float64, volume float64) {
	volume = swap.Amount0
	price = math.Abs(swap.Amount1 / swap.Amount0)
	return
}

func (tjs *TraderJoeScraper) FetchAvailablePairs() (pairs []dia.ExchangePair, err error) {
	return
}

func (tjs *TraderJoeScraper) GetPairData() {}
