package tradescraper

import (
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TraderJoeScraper struct {
	RestClient   *ethclient.Client
	WsClient     *ethclient.Client
	relDB        *models.RelDB
	datastore    *models.DB
	poolChannel  chan dia.Pool
	tradeChannel chan dia.Trade
	doneChannel  chan bool
	config       ScraperConfig
}

type ScraperConfig struct {
	BlockchainName  string
	ExchangeName    string
	RestDial        string
	WsDial          string
	FactoryContract string
	StartBlock      uint64
	WaitTime        int
}
