package source

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	traderjoe "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/traderjoe2.1"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TraderJoeAssetSource manages the scraping of assets for the Trader Joe exchange.
type TraderJoeAssetSource struct {
	RestClient             *ethclient.Client
	WsClient               *ethclient.Client
	relDB                  *models.RelDB
	assetChannel           chan dia.Asset
	doneChannel            chan bool
	exchange               dia.Exchange
	startBlock             uint64
	factoryContractAddress common.Address
	waitTime               int
}

// NewTraderJoeAssetSource initializes a Trader Joe asset sourcer, creating an instance of the
// NewTraderJoeAssetSource struct. It configures necessary parameters, initiates asset fetching, and returns
// the initialized scraper.
func NewTraderJoeAssetSource(exchange dia.Exchange, relDB *models.RelDB) *TraderJoeAssetSource {
	log.Info("NewTraderJoeLiquidityScraper ", exchange.Name)
	log.Info("NewTraderJoeLiquidityScraper Address ", exchange.Contract)

	var tjas *TraderJoeAssetSource

	switch exchange.Name {
	case dia.TraderJoeExchangeV2_1:
		tjas = makeTraderJoeAssetSource(exchange, "", "", relDB, "200", uint64(17821282))
	case dia.TraderJoeExchangeV2_1Arbitrum:
		tjas = makeTraderJoeAssetSource(exchange, "", "", relDB, "200", uint64(77473199))
	case dia.TraderJoeExchangeV2_1Avalanche:
		tjas = makeTraderJoeAssetSource(exchange, "", "", relDB, "200", uint64(28371397))
	case dia.TraderJoeExchangeV2_1BNB:
		tjas = makeTraderJoeAssetSource(exchange, "", "", relDB, "200", uint64(27099340))
	case dia.TraderJoeExchangeV2_2Avalanche:
		tjas = makeTraderJoeAssetSource(exchange, "", "", relDB, "200", uint64(46536129))
	}

	go func() {
		tjas.fetchAssets()
	}()
	return tjas
}

// makeTraderJoeAssetSource initializes a Trader Joe asset source, creating an instance of the
// TraderJoeAssetSource struct with the specified configuration and parameters.
func makeTraderJoeAssetSource(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB, waitMilliSeconds string, startBlock uint64) *TraderJoeAssetSource {
	var (
		restClient   *ethclient.Client
		wsClient     *ethclient.Client
		err          error
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		tjas         *TraderJoeAssetSource
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

	tjas = &TraderJoeAssetSource{
		RestClient:             restClient,
		WsClient:               wsClient,
		relDB:                  relDB,
		assetChannel:           assetChannel,
		doneChannel:            doneChannel,
		exchange:               exchange,
		waitTime:               waitTime,
		startBlock:             startBlock,
		factoryContractAddress: common.HexToAddress(exchange.Contract),
	}
	return tjas
}

// fetchAssets retrieves pool creation events from the Trader Joe factory contract address and processes them.
func (tjas *TraderJoeAssetSource) fetchAssets() {

	log.Info("Fetching Trader Joe LBPairCreated events...")
	log.Info("Getting lb pairs creations from address: ", tjas.factoryContractAddress.Hex())

	var blocknumber int64
	_, startblock, err := tjas.relDB.GetScraperIndex(tjas.exchange.Name, dia.SCRAPER_TYPE_ASSETCOLLECTOR)
	if err != nil {
		log.Error("GetScraperIndex: ", err)
	} else {
		tjas.startBlock = uint64(startblock)
	}

	// Initialize an Ethereum event filter for the Trader Joe factory contract.
	contractFilter, err := traderjoe.NewTraderjoeFilterer(tjas.factoryContractAddress, tjas.WsClient)
	if err != nil {
		log.Error(err)
	}

	// Retrieve LBPairCreated events using the event filter.
	lbPairCreated, err := contractFilter.FilterLBPairCreated(
		&bind.FilterOpts{Start: tjas.startBlock},
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Error("filter pool created: ", err)
	}

	// Don't send duplicate assets to the main.
	checkMap := make(map[string]struct{})

	// Iterate through the LBPairCreated events.
	for lbPairCreated.Next() {
		blocknumber = int64(lbPairCreated.Event.Raw.BlockNumber)

		asset0, err := ethhelper.ETHAddressToAsset(lbPairCreated.Event.TokenX, tjas.RestClient, tjas.exchange.BlockChain.Name)
		if err != nil {
			log.Errorf("ETHAddressToAsset for address %s: %v", lbPairCreated.Event.TokenX.Hex(), err)
		}
		asset1, err := ethhelper.ETHAddressToAsset(lbPairCreated.Event.TokenY, tjas.RestClient, tjas.exchange.BlockChain.Name)
		if err != nil {
			log.Errorf("ETHAddressToAsset for address %s: %v", lbPairCreated.Event.TokenX.Hex(), err)
		}

		if _, ok := checkMap[asset0.Address]; !ok {
			if asset0.Symbol != "" {
				checkMap[asset0.Address] = struct{}{}
				tjas.assetChannel <- asset0
			}
		}
		if _, ok := checkMap[asset1.Address]; !ok {
			if asset1.Symbol != "" {
				checkMap[asset1.Address] = struct{}{}
				tjas.assetChannel <- asset1
			}
		}

	}
	err = tjas.relDB.SetScraperIndex(tjas.exchange.Name, dia.SCRAPER_TYPE_ASSETCOLLECTOR, dia.INDEX_TYPE_BLOCKNUMBER, blocknumber)
	if err != nil {
		log.Error("SetScraperIndex: ", err)
	}
	// Signal that pool retrieval and processing is complete.
	tjas.doneChannel <- true
}

// Asset returns a channel for receiving dia.Asset instances scraped by the Trader Joe asset source.
func (tjas *TraderJoeAssetSource) Asset() chan dia.Asset {
	return tjas.assetChannel
}

// Done returns a channel for signaling the completion of Trader Joe asset scraping.
func (tjas *TraderJoeAssetSource) Done() chan bool {
	return tjas.doneChannel
}
