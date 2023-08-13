package tradescraper

import (
	"context"
	"github.com/diadata-org/diadata/pkg/dia"
	traderjoe "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/traderjoe2.1"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"strings"
)

var traderJoeABIJSON = `[
	{"type":"function","name":"getAllLBPairs","inputs":[],"outputs":[{"type":"tuple[]","name":"lbPairsAvailable","components":[{"type":"uint16","name":"binStep"},{"type":"address","name":"LBPair"},{"type":"bool","name":"createdByOwner"},{"type":"bool","name":"ignoredForRouting"}]}],"stateMutability":"view"},
		{"type":"event","name":"TradeEvent","inputs":[{"type":"address","name":"trader","indexed":true},{"type":"address","name":"src","indexed":true},{"type":"address","name":"dst","indexed":true},{"type":"uint256","name":"amount","indexed":false},{"type":"uint256","name":"tradeFee","indexed":false},{"type":"uint256","name":"srcBalance","indexed":false},{"type":"uint256","name":"dstBalance","indexed":false}],"anonymous":false}
	]`

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

func (tjs *TraderJoeScraper) fetchPools() {
	//TODO: Implement this:
	// logic to fetch pool data.

	traderJoeABI, err := abi.JSON(strings.NewReader(traderJoeABIJSON))
	if err != nil {
		log.Error("Error loading contract ABI: ", err)
		return
	}

	// Create the contract instance using the ABI and contract address.
	traderJoeContractAddress := "0x..."
	traderJoeContract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(traderJoeContractAddress), tjs.RestClient)
	if err != nil {
		log.Error("Error creating contract instance: ", err)
		return
	}

	// Call the getAllLBPairs function to fetch pool data.
	tokenX := common.HexToAddress("0x...") // Replace with the actual tokenX address
	tokenY := common.HexToAddress("0x...") // Replace with the actual tokenY address

	lbPairs, err := traderJoeContract.GetAllLBPairs(nil, tokenX, tokenY)
	if err != nil {
		log.Error("Error calllign getAllLBPairs function: ", err)
		return
	}

	// Process the returned lbPairs data.
	for _, lbPair := range lbPairs.LbPairsAvailable {
		pool := dia.Pool{}
	}
}

func (tjs *TraderJoeScraper) fetchTrades() {
	//TODO: Implement this:
	// logic to fetch trade data

	// Load the contract ABI
	traderJoeABI, err := abi.JSON(strings.NewReader(traderJoeABIJSON))
	if err != nil {
		log.Error("Error loading contract ABI: ", err)
		return
	}

	// Create a contract instance using the ABI and contract address.
	traderJoeContractAddress := "0x.."
	traderJoeContract, err := traderjoe.NewTraderjoeCaller(common.HexToAddress(traderJoeContractAddress), tjs.RestClient)
	if err != nil {
		log.Error("Error creating contract instance: ", err)
		return
	}

	// Setup subscription for the TradeEvent event.
	tradeEventSigHash := traderJoeABI.Events["TradeEvent"].ID
	logs := make(chan types.Log)
	sub, err := tjs.WsClient.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(traderJoeContractAddress)},
		Topics:    [][]common.Hash{{tradeEventSigHash}},
	}, logs)
	if err != nil {
		log.Error("Error subscribing to logs: ", err)
		return
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Error("Subscription error: ", err)
			return
		case log := <-logs:
			// Decode log data.
			var tradeEvent struct {
				Trader             common.Address
				Source             common.Address
				Destination        common.Address
				Amount             *big.Int
				TradeFee           *big.Int
				SourceBalance      *big.Int
				DestinationBalance *big.Int
			}
			err := traderJoeABI.UnpackIntoInterface(&tradeEvent, "TradeEvent", log.Data)
			if err != nil {
				log.Error("Error decoding log data: ", err)
				continue
			}

			// Create a dia.Trade object from the decoded data
			trade := dia.Trade{
				Trader:      tradeEvent.Trader.Hex(),
				Source:      tradeEvent.Source.Hex(),
				Destination: tradeEvent.Destination.Hex(),
				Amount:      tradeEvent.Amount.String(),
				TradeFee:    tradeEvent.TradeFee.String(),
				SrcBalance:  tradeEvent.SourceBalance.String(),
				DstBalance:  tradeEvent.DestinationBalance.String(),
			}

			// Send the trade object to the trade channel
			tjs.tradeChannel <- trade
		}
	}
}

func (tjs *TraderJoeScraper) Pool() chan dia.Pool {
	return tjs.poolChannel
}

func (tjs *TraderJoeScraper) Trade() chan dia.Trade {
	return tjs.tradeChannel
}

func (tjs *TraderJoeScraper) Done() chan bool {
	return tjs.doneChannel
}
