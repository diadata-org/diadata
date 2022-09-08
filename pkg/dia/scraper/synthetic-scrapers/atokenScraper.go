package synthscrapers

import (
	"context"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/config/synthContracts/aavepool2"
	"github.com/diadata-org/diadata/config/synthContracts/aavepool3"

	ceth "github.com/diadata-org/diadata/config/synthContracts/cETH"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*

Polygon aave pool address 0x794a61358d6845594f94dc1db02a252b5b4814ad

1) get enabled assets getReservesList
2) getAtoken address getReserveData
3)


1) Available liquidity = balance of dai in adai contract https://polygonscan.com/address/0x82E64f49Ed5EC1bC6e43DAD4FC8Af9bb3A2312EE
2) Total supplied =  Total supply of adai contract
3 ) Total borrowed =  Total supplied -  Available liquidity



*/
type aTokenScraper struct {
	RestClient   *ethclient.Client
	synthChannel chan dia.SynthAssetSupply
	blockchain   string
	WsClient     *ethclient.Client
	atokens      map[string]string // atokenaddress -> underlyingtokenaddress
	pooladdress  string
	version      int
	start        uint64
	protocol     string
}

func NewaTokenScraper(blockchain, pooladdress string, version int) *aTokenScraper {

	scraper := &aTokenScraper{
		blockchain:   blockchain,
		pooladdress:  pooladdress,
		version:      version,
		protocol:     "Aave",
		synthChannel: make(chan dia.SynthAssetSupply),
	}

	startblock, err := strconv.ParseInt(utils.Getenv("START_BLOCK", "0"), 10, 64)

	if err != nil {
		startblock = 0
	}

	scraper.start = uint64(startblock)

	atokens := make(map[string]string)

	restclient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(scraper.blockchain)+"_URI_REST", ""))
	if err != nil {
		log.Error("Error connecting rest Client")
	}

	wsclient, err := ethclient.Dial(utils.Getenv(strings.ToUpper(scraper.blockchain)+"_URI_WS", ""))
	if err != nil {
		log.Error("Error connecting ws Client")
	}

	scraper.RestClient = restclient
	scraper.WsClient = wsclient

	scraper.atokens = atokens

	switch scraper.version {
	case 2:
		scraper.getPoolAssetsV2()

	case 3:
		scraper.getPoolAssetsV3()
	}

	log.Info(scraper.atokens)

	go scraper.mainLoop()

	return scraper
}

func (scraper *aTokenScraper) watchReservedataupdatedV2() chan *aavepool2.Aavepool2ReserveDataUpdated {
	filterer, err := aavepool2.NewAavepool2Filterer(common.HexToAddress(scraper.pooladdress), scraper.WsClient)

	if err != nil {
		log.Error("new NewAavepool3Filterer caller: ", err)
	}
	sink := make(chan *aavepool2.Aavepool2ReserveDataUpdated)
	// if start is zero start from latest
	if scraper.start == 0 {
		filterer.WatchReserveDataUpdated(&bind.WatchOpts{}, sink, []common.Address{})

	} else {
		filterer.WatchReserveDataUpdated(&bind.WatchOpts{Start: &scraper.start}, sink, []common.Address{})
	}
	return sink

}

func (scraper *aTokenScraper) mainLoop() {
	// runs whenever there is  ReserveDataUpdated occurs

	log.Info("looking for watchReservedataupdatedV2")
	sink := scraper.watchReservedataupdatedV2()

	go func() {
		for {
			dataupdated, ok := <-sink
			if ok {
				log.Infoln("Reserve Data updated fetch supply and reserver of token ", dataupdated.Raw.TxHash.Hex())

				atokenasset, underlyingasset, atokensupply, reserver, err := scraper.fetchsupplyandbalance(dataupdated.Reserve.Hex(), big.NewInt(int64(dataupdated.Raw.BlockNumber)))
				if err != nil {
					log.Errorln("err on fetchsupplyandbalance", err)
					continue
				}

				block, err := scraper.RestClient.BlockByNumber(context.Background(), big.NewInt(int64(dataupdated.Raw.BlockNumber)))
				if err != nil {
					log.Errorln("error getting block details", err)
				}
				collateralRatio := reserver / atokensupply
				sas := dia.SynthAssetSupply{Asset: atokenasset, AssetUnderlying: underlyingasset, Supply: atokensupply, LockedUnderlying: reserver, BlockNumber: dataupdated.Raw.BlockNumber, Time: time.Unix(int64(block.Time()), 0), ColleteralRatio: collateralRatio, Protocol: scraper.protocol}
				scraper.synthChannel <- sas
			}
		}
	}()

}

func (scraper *aTokenScraper) getPoolAssetsV2() {
	log.Info("Getting V2 Pool assets and atokens")

	filterer, err := aavepool2.NewAavepool2Caller(common.HexToAddress(scraper.pooladdress), scraper.RestClient)
	if err != nil {
		log.Error("err on  NewAavepool2Caller: ", err)
	}

	assets, err := filterer.GetReservesList(&bind.CallOpts{})
	if err != nil {
		log.Error("err on GetReservesList: ", err)
	}

	log.Info("total assets found ", len(assets))

	log.Info("getting respective atokens for assets ...")

	for _, asset := range assets {

		data, err := filterer.GetReserveData(&bind.CallOpts{}, asset)
		if err != nil {
			log.Error("GetReserveData: ", err)
		}

		scraper.atokens[asset.Hex()] = data.ATokenAddress.Hex()

	}

	log.Info("total assets in pool ", len(scraper.atokens))

}
func (scraper *aTokenScraper) getPoolAssetsV3() {
	log.Info("Getting V3 Pool assets and atokens")

	filterer, err := aavepool3.NewAavepool3Caller(common.HexToAddress(scraper.pooladdress), scraper.RestClient)
	if err != nil {
		log.Error("error on  NewAavepool3Caller: ", err)
	}

	assets, err := filterer.GetReservesList(&bind.CallOpts{})
	if err != nil {
		log.Error("err on  GetReservesList: ", err)
	}

	for _, asset := range assets {

		data, err := filterer.GetReserveData(&bind.CallOpts{}, asset)
		if err != nil {
			log.Error("err on GetReserveData: ", err)
		}

		scraper.atokens[asset.Hex()] = data.ATokenAddress.Hex()

	}

	log.Info("total assets in pool ", len(scraper.atokens))

}

func (scraper *aTokenScraper) GetSynthSupplyChannel() chan dia.SynthAssetSupply {
	return scraper.synthChannel
}

func (scraper *aTokenScraper) FetchSynthSupply() error {

	return nil
}

func (scraper *aTokenScraper) fetchsupplyandbalance(underlyingtokenaddress string, blocknumber *big.Int) (atokenasset, underlyingasset dia.Asset, atokensupply, underlyingreserve float64, err error) {

	atokenaddress := scraper.atokens[underlyingtokenaddress]

	log.Info("atokenaddress", atokenaddress)
	log.Info("underlyingtokenaddress", underlyingtokenaddress)

	filterer, err := ceth.NewERC20Caller(common.HexToAddress(atokenaddress), scraper.RestClient)
	if err != nil {
		log.Error("new erc20 caller: ", err)
		return
	}
	supply, err := filterer.TotalSupply(&bind.CallOpts{BlockNumber: blocknumber})
	if err != nil {
		log.Error("get supply: ", err)
		return
	}

	atokendecimal, err := filterer.Decimals(&bind.CallOpts{BlockNumber: blocknumber})
	if err != nil {
		log.Error("get Decimals: ", err)
		return
	}
	atokensymbol, err := filterer.Symbol(&bind.CallOpts{BlockNumber: blocknumber})
	if err != nil {
		log.Error("get Decimals: ", err)
		return
	}

	atokenasset.Symbol = atokensymbol
	atokenasset.Address = atokenaddress
	atokenasset.Decimals = uint8(atokendecimal.Int64())
	atokenasset.Blockchain = scraper.blockchain

	atokensupply, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(supply), new(big.Float).SetFloat64(math.Pow10(int(atokendecimal.Int64())))).Float64()

	underlyingfilterer, err := ceth.NewERC20Caller(common.HexToAddress(underlyingtokenaddress), scraper.RestClient)
	if err != nil {
		log.Error("new erc20 caller: ", err)
		return
	}
	balanceof, err := underlyingfilterer.BalanceOf(&bind.CallOpts{BlockNumber: blocknumber}, common.HexToAddress(atokenaddress))
	if err != nil {
		log.Error("get balanceof: ", err)
		return
	}
	log.Info("balanceof: ", balanceof)

	underlyingdecimals, err := underlyingfilterer.Decimals(&bind.CallOpts{BlockNumber: blocknumber})
	if err != nil {
		log.Error("get Decimals: ", err)
		return
	}
	underlyingsymbol, err := underlyingfilterer.Symbol(&bind.CallOpts{BlockNumber: blocknumber})
	if err != nil {
		log.Error("get Symbol: ", err)
		return
	}

	underlyingreserve, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(balanceof), new(big.Float).SetFloat64(math.Pow10(int(underlyingdecimals.Int64())))).Float64()
	underlyingasset.Symbol = underlyingsymbol

	underlyingasset.Address = underlyingtokenaddress
	underlyingasset.Decimals = uint8(underlyingdecimals.Int64())
	underlyingasset.Blockchain = scraper.blockchain
	return

}
