package liquidityscrapers

import (
	"context"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	uniswapcontractv4 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv4"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapV4Scraper struct {
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
	chunksBlockSize uint64
	waitTime        int
}

const univ4_start_block = uint64(21688329)

// NewUniswapV4Scraper returns a new UniswapV4Scraper.
func NewUniswapV4Scraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *UniswapV4Scraper {
	log.Info("NewUniswapV4Scraper ", exchange.Name)
	log.Info("NewUniswapV4Scraper Address ", exchange.Contract)

	var uls *UniswapV4Scraper

	switch exchange.Name {
	case dia.UniswapExchangeV4:
		uls = makeUniswapV4Scraper(exchange, "", "", relDB, datastore, "200", univ4_start_block)
	case dia.UniswapExchangeV4Base:
		uls = makeUniswapV4Scraper(exchange, "", "", relDB, datastore, "200", uint64(25350988))
	}

	go func() {
		uls.fetchPools()
	}()
	return uls
}

// makeUniswapV4Scraper returns a uniswap scraper as used in NewUniswapV4Scraper.
func makeUniswapV4Scraper(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB, datastore *models.DB, waitMilliseconds string, startBlock uint64) *UniswapV4Scraper {
	var (
		restClient  *ethclient.Client
		wsClient    *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		uls         *UniswapV4Scraper
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
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	uls = &UniswapV4Scraper{
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
	blockSize := utils.Getenv("CHUNKS_BLOCK_SIZE", "10000")
	uls.chunksBlockSize, err = strconv.ParseUint(blockSize, 10, 64)
	if err != nil {
		log.Error("Parse CHUNKS_BLOCK_SIZE: ", err)
		uls.chunksBlockSize = uint64(10000)
	}

	return uls
}

// fetchPools fetches all registered pools from on-chain and sends them into the pool channel.
func (uls *UniswapV4Scraper) fetchPools() {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603
	log.Info("get pool creations from address: ", uls.factoryContract)
	poolsCount := 0
	// contract, err := uniswapcontractv4.NewUniswapV4Filterer(common.HexToAddress(uls.factoryContract), uls.WsClient)
	contract, err := uniswapcontractv4.NewPoolmanagerFilterer(common.HexToAddress(uls.factoryContract), uls.WsClient)
	if err != nil {
		log.Error(err)
	}

	// Iterate over chunks of blocks.
	currentBlockNumber, err := uls.RestClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Get current block number: ", err)
	}
	var startblock, endblock uint64
	startblock = uls.startBlock
	endblock = uls.startBlock + uls.chunksBlockSize

	for endblock < currentBlockNumber+uls.chunksBlockSize {

		time.Sleep(time.Duration(uls.waitTime) * time.Millisecond)

		poolCreated, err := contract.FilterInitialize(
			&bind.FilterOpts{
				Start: startblock,
				End:   &endblock,
			},
			[][32]byte{},
			[]common.Address{},
			[]common.Address{},
		)
		if err != nil {
			log.Error("filter pool created: ", err)
			startblock = endblock
			endblock = startblock + uls.chunksBlockSize
			continue
		}

		for poolCreated.Next() {
			poolsCount++
			var (
				pool   dia.Pool
				asset0 dia.Asset
				asset1 dia.Asset
			)
			log.Info("pools count: ", poolsCount)

			asset0, err = uls.relDB.GetAsset(poolCreated.Event.Currency0.Hex(), uls.blockchain)
			if err != nil {
				asset0, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Currency0, uls.RestClient, uls.blockchain)
				if err != nil {
					log.Warn("Currency0: cannot fetch asset with address ", poolCreated.Event.Currency0.Hex())
					continue
				}
			}
			asset1, err = uls.relDB.GetAsset(poolCreated.Event.Currency1.Hex(), uls.blockchain)
			if err != nil {
				asset1, err = ethhelper.ETHAddressToAsset(poolCreated.Event.Currency1, uls.RestClient, uls.blockchain)
				if err != nil {
					log.Warn("Currency1: cannot fetch asset with address ", poolCreated.Event.Currency1.Hex())
					continue
				}
			}

			log.Infof("%s-%s : %s -- %s", asset0.Symbol, asset1.Symbol, poolCreated.Event.Currency0.Hex(), poolCreated.Event.Currency1.Hex())

			pool.Exchange = dia.Exchange{Name: uls.exchangeName}
			pool.Blockchain = dia.BlockChain{Name: uls.blockchain}
			pool.Address = hex.EncodeToString(poolCreated.Event.Id[:])

			pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset0, Volume: float64(0), Index: uint8(0)})
			pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset1, Volume: float64(0), Index: uint8(1)})
			pool.Time = time.Now()

			uls.poolChannel <- pool

		}
		startblock = endblock
		endblock = startblock + uls.chunksBlockSize

	}
	uls.doneChannel <- true
}

func (uls *UniswapV4Scraper) Pool() chan dia.Pool {
	return uls.poolChannel
}

func (uls *UniswapV4Scraper) Done() chan bool {
	return uls.doneChannel
}
