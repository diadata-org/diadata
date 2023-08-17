package tradescraper

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/diadata-org/diadata/pkg/dia"
	traderjoe "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/traderjoe2.1"
	models "github.com/diadata-org/diadata/pkg/model"
)

// NewTraderJoeTradeScraper creates and initializes a new TraderJoeScraper instance.
func NewTraderJoeTradeScraper(config ScraperConfig, relDB *models.RelDB, datastore *models.DB) *TraderJoeScraper {
	tjs := &TraderJoeScraper{
		relDB:        relDB,
		datastore:    datastore,
		poolChannel:  make(chan dia.Pool),
		tradeChannel: make(chan dia.Trade),
		doneChannel:  make(chan bool),
		config:       config,
	}

	// Initialize Ethereum REST client
	restClient, err := ethclient.Dial(config.RestDial)
	if err != nil {
		log.Fatal("Error initializing REST client: ", err)
	}
	tjs.RestClient = restClient

	// Initialize the Ethereum WebSocket client.
	wsClient, err := ethclient.Dial(config.WsDial)
	if err != nil {
		log.Fatal("Error initializing WebSocket client: ", err)
	}
	tjs.WsClient = wsClient

	go func() {
		tjs.fetchPools()
	}()

	go func() {
		tjs.fetchTrades()
	}()

	return tjs
}

// FetchPools is responsible for fetching pool data from TraderJoe.
func (tjs *TraderJoeScraper) FetchPools() {
	traderJoeContractAddress := tjs.config.FactoryContract
	traderJoeContract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(traderJoeContractAddress), tjs.RestClient)
	if err != nil {
		log.Printf("Error creating contract instance: %v\n", err)
		return
	}

	numLBPairs, err := traderJoeContract.GetNumberOfLBPairs(nil)
	if err != nil {
		log.Printf("Error calling GetPools function: %v\n", err)
		return
	}

	for i := big.NewInt(0); i.Cmp(numLBPairs) < 0; i.Add(i, big.NewInt(1)) {
		lbPairAddress, err := traderJoeContract.GetLBPairAtIndex(nil, i)
		if err != nil {
			log.Printf("Error calling GetLBPairAtIndex function: %v\n", err)
			continue
		}

		// Retrieve token addresses and bin step using the contract methods
		tokenX := traderJoeContract.TokenX(nil, lbPairAddress)
		tokenY := traderJoeContract.TokenY(nil, lbPairAddress)
		binStep := traderJoeContract.GetBinStep(nil, lbPairAddress)

		// Construct and populate the Pool object
		pool := dia.Pool{
			Address:    lbPairAddress.Hex(),
			Exchange:   dia.Exchange{Name: tjs.config.ExchangeName},
			Blockchain: dia.BlockChain{Name: tjs.config.BlockchainName},
			// Populate other fields with retrieved values and config as needed
			Time: time.Now(),
		}

		tjs.poolChannel <- pool
	}
	close(tjs.poolChannel)
}

func (tjs *TraderJoeScraper) FetchTrades() {
	traderJoeContractAddress := tjs.config.FactoryContract
	traderJoeContract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(traderJoeContractAddress), tjs.RestClient)
	if err != nil {
		log.Printf("Error creating contract instance: %v\n", err)
		return
	}

	logs := make(chan types.Log)
	sub, err := tjs.WsClient.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(traderJoeContractAddress)},
	}, logs)
	if err != nil {
		log.Printf("Error subscribing to logs: %v\n", err)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v\n", err)
			return
		case logger := <-logs:
			var trade dia.Trade
			err := traderJoeContract.UnpackLog(&trade, "TradeEvent", logger)
			if err != nil {
				log.Printf("Error decoding log data: %v\n", err)
				continue
			}

			// Set other fields of the trade object as needed
			trade.Exchange = dia.Exchange{Name: tjs.config.ExchangeName}
			trade.Blockchain = dia.BlockChain{Name: tjs.config.BlockchainName}
			trade.Time = time.Now()
			trade.PoolAddress = logger.Address.Hex()
			trade.VerifiedPair = true

			// Calculate Volume and EstimatedUSDPrice based on your logic
			// For example, if tokenX and tokenY prices are known:
			//tokenXPrice := ... // Retrieve tokenX price from an external source
			//tokenYPrice := ... // Retrieve tokenY price from an external source
			//
			//trade.Volume = ... // Calculate the trade volume based on trade amount
			//trade.EstimatedUSDPrice = (tokenXPrice + tokenYPrice) / 2 * trade.Price

			tjs.tradeChannel <- trade
		}
	}
}

//// fetchPools is responsible for fetching pool data from TraderJoe.
//func (tjs *TraderJoeScraper) fetchPools() {
//	// Create the contract instance using the ABI and contract address.
//	traderJoeContractAddress := tjs.config.FactoryContract
//	traderJoeContract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(traderJoeContractAddress), tjs.RestClient)
//	if err != nil {
//		log.Printf("Error creating contract instance: %v\n", err)
//		return
//	}
//
//	// Retrieve number of liquidity pairs
//	numLBPairs, err := traderJoeContract.GetNumberOfLBPairs(nil)
//	if err != nil {
//		log.Printf("Error calling GetPools function:%v\n ", err)
//		return
//	}
//	// Fetch pool address and process each pool
//	for i := big.NewInt(0); i.Cmp(numLBPairs) < 0; i.Add(i, big.NewInt(1)) {
//		lbPairAddress, err := traderJoeContract.GetLBPairAtIndex(nil, i)
//		if err != nil {
//			log.Printf("Error calling GetLBPairAtIndex function: %v\n", err)
//			continue
//		}
//
//		// TraderJoe token addresses and binStep
//		tokenA := common.HexToAddress("0x8e42f2F4101563bF679975178e880FD87d3eFd4e") // Replace with actual token address
//		tokenB := common.HexToAddress("0xb4315e873dBcf96Ffd0acd8EA43f689D8c20fB30") // Replace with actual token address
//		binStep := big.NewInt(1)                                                    // Replace with appropriate binStep value
//
//		// TODO: Sort out poolInfo
//		// Fetch pool information using GetLBPairInformation
//		poolInfo, err := traderJoeContract.GetLBPairInformation(nil, tokenA, tokenB, binStep)
//		if err != nil {
//			log.Printf("Error calling GetLBPairInformation function: %v\n", err)
//			continue
//		}
//
//		// TODO: Fill AssetVolumes:
//		pool := dia.Pool{
//			Address:      lbPairAddress.Hex(),
//			Exchange:     dia.Exchange{Name: tjs.config.ExchangeName},
//			Blockchain:   dia.BlockChain{Name: tjs.config.BlockchainName},
//			Assetvolumes: []dia.AssetVolume{},
//			Time:         time.Now(),
//		}
//
//		tjs.poolChannel <- pool
//	}
//	close(tjs.poolChannel)
//}
//
//// fetchTrades is responsible for fetching trade data from TraderJoe.
//func (tjs *TraderJoeScraper) fetchTrades() {
//	// Create a contract instance using the ABI and contract address.
//	traderJoeContractAddress := tjs.config.FactoryContract
//	traderJoeContract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(traderJoeContractAddress), tjs.RestClient)
//	if err != nil {
//		log.Printf("Error creating contract instance: %v\n", err)
//		return
//	}
//
//	// Setup subscription for the TradeEvent event.
//	tradeEventSigHash := traderJoeContract.
//	logs := make(chan types.Log)
//	sub, err := tjs.WsClient.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
//		Addresses: []common.Address{common.HexToAddress(traderJoeContractAddress)},
//		Topics:    [][]common.Hash{{tradeEventSigHash}},
//	}, logs)
//	if err != nil {
//		log.Error("Error subscribing to logs: ", err)
//		return
//	}
//	defer sub.Unsubscribe()
//
//	for {
//		select {
//		case err := <-sub.Err():
//			log.Error("Subscription error: ", err)
//			return
//		case log := <-logs:
//			// Decode log data.
//			var tradeEvent struct {
//				Trader             common.Address
//				Source             common.Address
//				Destination        common.Address
//				Amount             *big.Int
//				TradeFee           *big.Int
//				SourceBalance      *big.Int
//				DestinationBalance *big.Int
//			}
//			err := traderJoeABI.UnpackIntoInterface(&tradeEvent, "TradeEvent", log.Data)
//			if err != nil {
//				log.Error("Error decoding log data: ", err)
//				continue
//			}
//
//			// Create a dia.Trade object from the decoded data
//			trade := dia.Trade{
//				Trader:      tradeEvent.Trader.Hex(),
//				Source:      tradeEvent.Source.Hex(),
//				Destination: tradeEvent.Destination.Hex(),
//				Amount:      tradeEvent.Amount.String(),
//				TradeFee:    tradeEvent.TradeFee.String(),
//				SrcBalance:  tradeEvent.SourceBalance.String(),
//				DstBalance:  tradeEvent.DestinationBalance.String(),
//			}
//
//			// Send the trade object to the trade channel
//			tjs.tradeChannel <- trade
//		}
//	}
//}

// Pool returns the channel for sending pool data.
func (tjs *TraderJoeScraper) Pool() chan dia.Pool {
	return tjs.poolChannel
}

// Trade returns the channel for sending trade data.
func (tjs *TraderJoeScraper) Trade() chan dia.Trade {
	return tjs.tradeChannel
}

// Done returns the channel to signal the scraper is done.
func (tjs *TraderJoeScraper) Done() chan bool {
	return tjs.doneChannel
}
