package tradescraper

import (
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

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
}

func (tjs *TraderJoeScraper) fetchTrades() {
	//TODO: Implement this:
	// logic to fetch trade data
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
