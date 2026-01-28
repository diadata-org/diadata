package source

import (
	"context"
	"strconv"
	"strings"
	"time"

	uniswapcontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"
	uniswapcontractv4 "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswapv4"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var numBlocksQueryUniswapV4 = uint64(1000)

type UniswapV4AssetSource struct {
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
	relDB      *models.RelDB
	// signaling channels for session initialization and finishing
	assetChannel    chan dia.Asset
	doneChannel     chan bool
	exchange        dia.Exchange
	startBlock      uint64
	factoryContract string
	waitTime        int
}

// NewUniswapV4AssetSource returns a new UniswapV4AssetSource
func NewUniswapV4AssetSource(exchange dia.Exchange, relDB *models.RelDB) *UniswapV4AssetSource {
	log.Info("NewUniswapV4Scraper ", exchange.Name)
	log.Info("NewUniswapV4Scraper Address ", exchange.Contract)

	var uas *UniswapV4AssetSource

	switch exchange.Name {
	case dia.UniswapExchangeV4:
		uas = makeUniswapV4AssetSource(exchange, "", "", relDB, "200", uint64(21688329))
	case dia.UniswapExchangeV4Base:
		uas = makeUniswapV4AssetSource(exchange, "", "", relDB, "200", uint64(25350988))
	}

	go func() {
		uas.fetchAssets()
	}()
	return uas

}

// makeUniswapV4AssetSource returns a uniswap asset source.
func makeUniswapV4AssetSource(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB, waitMilliseconds string, startBlock uint64) *UniswapV4AssetSource {
	var (
		restClient   *ethclient.Client
		wsClient     *ethclient.Client
		err          error
		uas          *UniswapV4AssetSource
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)

	log.Infof("Init rest and ws client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}
	wsClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_WS", wsDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	var waitTime int
	waitTimeString := utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_WAIT_TIME", waitMilliseconds)
	waitTime, err = strconv.Atoi(waitTimeString)
	if err != nil {
		log.Error("could not parse wait time: ", err)
		waitTime = 500
	}

	uas = &UniswapV4AssetSource{
		RestClient:      restClient,
		WsClient:        wsClient,
		relDB:           relDB,
		assetChannel:    assetChannel,
		doneChannel:     doneChannel,
		exchange:        exchange,
		startBlock:      startBlock,
		factoryContract: exchange.Contract,
		waitTime:        waitTime,
	}
	return uas
}

// getNumPairs returns the number of available pairs on Uniswap
func (uas *UniswapV4AssetSource) fetchAssets() {

	// filter from contract created https://etherscan.io/tx/0x1e20cd6d47d7021ae7e437792823517eeadd835df09dde17ab45afd7a5df4603

	log.Info("get pool creations from address: ", uas.factoryContract)
	poolsCount := 0
	var blocknumber int64
	checkMap := make(map[string]struct{})
	_, startblock, err := uas.relDB.GetScraperIndex(uas.exchange.Name, dia.SCRAPER_TYPE_ASSETCOLLECTOR)
	if err != nil {
		log.Error("GetScraperIndex: ", err)
	} else {
		uas.startBlock = uint64(startblock)
	}

	contract, err := uniswapcontractv4.NewPoolmanagerFilterer(common.HexToAddress(uas.factoryContract), uas.WsClient)
	if err != nil {
		log.Error(err)
	}

	currentBlockNumber, err := uas.RestClient.BlockNumber(context.Background())
	if err != nil {
		log.Error("GetBlockNumber: ", err)
	}

	endblock := utils.Min(uint64(uas.startBlock)+numBlocksQueryUniswapV4, currentBlockNumber)
	log.Infof("startblock -- endblock: %v -- %v", uas.startBlock, endblock)

	for uas.startBlock <= currentBlockNumber {
		poolCreated, err := contract.FilterInitialize(
			&bind.FilterOpts{
				Start: uas.startBlock,
				End:   &endblock,
			},
			[][32]byte{},
			[]common.Address{},
			[]common.Address{},
		)
		if err != nil {
			log.Error("filter pool created: ", err)
		}
		for poolCreated.Next() {
			time.Sleep(time.Duration(uas.waitTime) * time.Millisecond)
			poolsCount++
			log.Info("pools count: ", poolsCount)

			blocknumber = int64(poolCreated.Event.Raw.BlockNumber)
			address0 := poolCreated.Event.Currency0
			address1 := poolCreated.Event.Currency1
			// Don't repeat sending already sent assets.
			// Take into account that UniswapV4 allows for trading unwrapped ETH.
			if (address0 != common.Address{}) {
				if _, ok := checkMap[address0.Hex()]; !ok {
					checkMap[address0.Hex()] = struct{}{}
					asset, err := uas.GetAssetFromAddress(address0)
					if err != nil {
						log.Warnf("cannot fetch asset from address %s: %v", address0.Hex(), err)
					}
					uas.assetChannel <- asset
				}
			}
			if (address1 != common.Address{}) {
				if _, ok := checkMap[address1.Hex()]; !ok {
					checkMap[address1.Hex()] = struct{}{}
					asset, err := uas.GetAssetFromAddress(address1)
					if err != nil {
						log.Warnf("cannot fetch asset from address %s: %v", address1.Hex(), err)
					}
					uas.assetChannel <- asset
				}
			}
		}
		err = uas.relDB.SetScraperIndex(uas.exchange.Name, dia.SCRAPER_TYPE_ASSETCOLLECTOR, dia.INDEX_TYPE_BLOCKNUMBER, blocknumber)
		if err != nil {
			log.Error("SetScraperIndex: ", err)
		}
		endblock += numBlocksQueryUniswapV4
		uas.startBlock += numBlocksQueryUniswapV4
	}

	uas.doneChannel <- true
}

func (uas *UniswapV4AssetSource) GetAssetFromAddress(address common.Address) (asset dia.Asset, err error) {
	connection := uas.RestClient

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
		Blockchain: uas.exchange.BlockChain.Name,
		Decimals:   decimals,
	}

	return
}

func (uas *UniswapV4AssetSource) Asset() chan dia.Asset {
	return uas.assetChannel
}

func (uas *UniswapV4AssetSource) Done() chan bool {
	return uas.doneChannel
}
