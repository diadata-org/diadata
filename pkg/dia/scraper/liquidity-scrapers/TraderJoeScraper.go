package liquidityscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"strconv"
	"strings"
)

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

// NewTraderJoeLiquidityScraper is a constructor for the TraderJoeLiquidityScraper struct type.
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

// fetchPools fetches all registered pool data from on-chain and sends them to the pool channel.
func (tjls TraderJoeLiquidityScraper) fetchPools() {
	log.Info("Get pool creations from address: ", tjls.factoryContract)
	poolsCount := 0
	contract, err :- trad
}
