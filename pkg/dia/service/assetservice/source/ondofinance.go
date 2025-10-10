package source

import (
	"strings"
	"time"

	ondo "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/ondofinance"
	uniswapcontract "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/uniswap"
	models "github.com/diadata-org/diadata/pkg/model"

	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type OndoFinanceAssetSource struct {
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
	relDB      *models.RelDB
	// signaling channels for session initialization and finishing
	assetChannel    chan dia.Asset
	doneChannel     chan bool
	exchange        dia.Exchange
	factoryContract string
}

// NewOndoFinanceAssetSource returns a new OndoFinanceAssetSource
func NewOndoFinanceAssetSource(exchange dia.Exchange, relDB *models.RelDB) *OndoFinanceAssetSource {
	log.Infof("New %s scraper with address %s", exchange.Name, exchange.Contract)

	uas := makeOndoFinanceAssetSource(exchange, "", "", relDB)

	go func() {
		uas.fetchAssets()
	}()

	return uas

}

// makeOndoFinanceAssetSource returns a uniswap asset source.
func makeOndoFinanceAssetSource(exchange dia.Exchange, restDial string, wsDial string, relDB *models.RelDB) *OndoFinanceAssetSource {
	var (
		restClient   *ethclient.Client
		wsClient     *ethclient.Client
		err          error
		uas          *OndoFinanceAssetSource
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

	uas = &OndoFinanceAssetSource{
		RestClient:      restClient,
		WsClient:        wsClient,
		relDB:           relDB,
		assetChannel:    assetChannel,
		doneChannel:     doneChannel,
		exchange:        exchange,
		factoryContract: exchange.Contract,
	}
	return uas
}

// getNumPairs returns the number of available pairs on Uniswap
func (uas *OndoFinanceAssetSource) fetchAssets() {

	time.Sleep(4 * time.Second)

	sink, err := uas.getTradesChannel()
	if err != nil {
		log.Error("error fetching trades channel: ", err)
	}

	go func() {
		for {
			rawTrade, ok := <-sink
			if ok {

				asset, err := uas.GetAssetFromAddress(rawTrade.Asset)
				if err != nil {
					log.Error("GetAssetFromAddress: ", err)
					continue
				}
				uas.assetChannel <- asset

			}
		}
	}()

}

func (uas *OndoFinanceAssetSource) getTradesChannel() (chan *ondo.OndotokenmanagerTradeExecuted, error) {

	// TO DO: Contract address from exchange struct.
	contract, err := ondo.NewOndotokenmanagerFilterer(common.HexToAddress(uas.factoryContract), uas.WsClient)
	if err != nil {
		log.Error("NewOndotokenmanagerFilterer: ", err)
	}
	tradesSink := make(chan *ondo.OndotokenmanagerTradeExecuted)
	_, err = contract.WatchTradeExecuted(&bind.WatchOpts{}, tradesSink)
	if err != nil {
		log.Error("WatchTradeExecuted: ", err)
	}

	return tradesSink, nil
}

func (uas *OndoFinanceAssetSource) GetAssetFromAddress(address common.Address) (asset dia.Asset, err error) {

	var tokenContract *uniswapcontract.IERC20Caller

	tokenContract, err = uniswapcontract.NewIERC20Caller(address, uas.RestClient)
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

func (uas *OndoFinanceAssetSource) Asset() chan dia.Asset {
	return uas.assetChannel
}

func (uas *OndoFinanceAssetSource) Done() chan bool {
	return uas.doneChannel
}
