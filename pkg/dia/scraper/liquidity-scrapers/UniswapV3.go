package liquidityscrapers

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	uniswapcontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"
	uniswapcontractv3 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv3"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UniswapV3Scraper struct {
	RestClient      *ethclient.Client
	WsClient        *ethclient.Client
	poolChannel     chan dia.Pool
	doneChannel     chan bool
	blockchain      string
	startBlock      uint64
	factoryContract string
	exchangeName    string
	waitTime        int
}

// NewUniswapV3Scraper returns a new UniswapV3Scraper.
func NewUniswapV3Scraper(exchange dia.Exchange) *UniswapV3Scraper {
	log.Info("NewUniswapScraper ", exchange.Name)
	log.Info("NewUniswapScraper Address ", exchange.Contract)

	var uls *UniswapV3Scraper

	switch exchange.Name {
	case dia.UniswapExchangeV3:
		uls = makeUniswapV3Scraper(exchange, "", "", "200", uint64(12369621))
	case dia.UniswapExchangeV3Polygon:
		uls = makeUniswapV3Scraper(exchange, "", "", "200", uint64(22757913))
	case dia.UniswapExchangeV3Arbitrum:
		uls = makeUniswapV3Scraper(exchange, "", "", "200", uint64(165))
	}

	go func() {
		uls.fetchPools()
	}()
	return uls
}

// makeUniswapV3Scraper returns a uniswap scraper as used in NewUniswapV3Scraper.
func makeUniswapV3Scraper(exchange dia.Exchange, restDial string, wsDial string, waitMilliseconds string, startBlock uint64) *UniswapV3Scraper {
	var (
		restClient  *ethclient.Client
		wsClient    *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		uls         *UniswapV3Scraper
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

	uls = &UniswapV3Scraper{
		WsClient:        wsClient,
		RestClient:      restClient,
		poolChannel:     poolChannel,
		doneChannel:     doneChannel,
		blockchain:      exchange.BlockChain.Name,
		startBlock:      startBlock,
		factoryContract: exchange.Contract,
		exchangeName:    exchange.Name,
		waitTime:        waitTime,
	}
	return uls
}

// fetchPools fetches all registered pools from on-chain and sends them into the pool channel.
func (uls *UniswapV3Scraper) fetchPools() {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603

	log.Info("get pool creations from address: ", uls.factoryContract)
	poolsCount := 0
	contract, err := uniswapcontractv3.NewUniswapV3Filterer(common.HexToAddress(uls.factoryContract), uls.WsClient)
	if err != nil {
		log.Error(err)
	}

	poolCreated, err := contract.FilterPoolCreated(
		&bind.FilterOpts{Start: uls.startBlock},
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		log.Error("filter pool created: ", err)
	}

	for poolCreated.Next() {
		poolsCount++
		var pool dia.Pool
		log.Info("pools count: ", poolsCount)

		asset0, err := uls.GetAssetFromAddress(poolCreated.Event.Token0)
		if err != nil {
			log.Warn("cannot fetch asset from address ", poolCreated.Event.Token0.Hex())
		}
		asset1, err := uls.GetAssetFromAddress(poolCreated.Event.Token1)
		if err != nil {
			log.Warn("cannot fetch asset from address ", poolCreated.Event.Token1.Hex())
		}

		pool.Exchange = dia.Exchange{Name: uls.exchangeName}
		pool.Blockchain = dia.BlockChain{Name: uls.blockchain}
		pool.Address = poolCreated.Event.Pool.Hex()
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset0, Index: uint8(0)})
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{Asset: asset1, Index: uint8(1)})
		pool.Time = time.Now()

		uls.poolChannel <- pool

	}
	uls.doneChannel <- true
}

func (uls *UniswapV3Scraper) GetAssetFromAddress(address common.Address) (asset dia.Asset, err error) {
	connection := uls.RestClient

	var tokenContract *uniswapcontract.IERC20Caller

	tokenContract, err = uniswapcontract.NewIERC20Caller(address, connection)
	if err != nil {
		log.Error(err)
	}

	symbol, err := tokenContract.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	name, err := tokenContract.Name(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}
	decimals, err := tokenContract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error(err)
	}

	asset = dia.Asset{
		Symbol:     symbol,
		Name:       name,
		Address:    address.Hex(),
		Blockchain: uls.blockchain,
		Decimals:   decimals,
	}

	return
}

func (uas *UniswapV3Scraper) Pool() chan dia.Pool {
	return uas.poolChannel
}

func (uas *UniswapV3Scraper) Done() chan bool {
	return uas.doneChannel
}
