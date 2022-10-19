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

type atokendetail struct {
	address                  string
	stabledebttokenaddress   string
	variabledebttokenaddress string
}
type aTokenScraper struct {
	RestClient   *ethclient.Client
	synthChannel chan dia.SynthAssetSupply
	blockchain   string
	WsClient     *ethclient.Client
	atokens      map[string]atokendetail // atokenaddress -> underlyingtokenaddress
	pooladdress  string
	version      int
	start        uint64
	protocol     string
	endBlock     uint64
}

func NewaTokenScraper(blockchain, pooladdress string, version int) *aTokenScraper {

	scraper := &aTokenScraper{
		blockchain:   blockchain,
		pooladdress:  pooladdress,
		version:      version,
		protocol:     "Aave-V" + strconv.Itoa(version),
		synthChannel: make(chan dia.SynthAssetSupply),
	}

	startblock, err := strconv.ParseInt(utils.Getenv("START_BLOCK", "0"), 10, 64)
	endBlock, err := strconv.ParseInt(utils.Getenv("END_BLOCK", "0"), 10, 64)

	if err != nil {
		startblock = 0
	}

	scraper.start = uint64(startblock)

	atokens := make(map[string]atokendetail)

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
	scraper.endBlock = uint64(endBlock)

	switch scraper.version {
	case 2:
		scraper.getPoolAssetsV2()

	case 3:
		scraper.getPoolAssetsV3()
	}

	log.Info(scraper.atokens)

	if scraper.start == 0 {
		go scraper.mainLoop()

	} else {
		log.Info("running")
		go scraper.mainLoopHistory()

	}

	return scraper
}

func (scraper *aTokenScraper) filterReservedataupdatedV2History() chan *aavepool2.Aavepool2ReserveDataUpdated {
	filterer, err := aavepool2.NewAavepool2Filterer(common.HexToAddress(scraper.pooladdress), scraper.RestClient)

	if err != nil {
		log.Error("new NewAavepool3Filterer caller: ", err)
	}
	sink := make(chan *aavepool2.Aavepool2ReserveDataUpdated)
	// if start is zero start from latest

	log.Error("FilterReserveDataUpdated starting from block: ", scraper.start)
	endblock := scraper.start + 100

	go func() {
		for endblock <= scraper.endBlock {
			iter, err := filterer.FilterReserveDataUpdated(&bind.FilterOpts{Start: scraper.start, End: &endblock}, []common.Address{})
			if err != nil {
				log.Errorln("error getting filterdata", err)
			}
			for iter.Next() {
				sink <- iter.Event
			}
			scraper.start = scraper.start + 100
			endblock = scraper.start + 100
			log.Infof("increasing block  start: %d, End: %d ", scraper.start, endblock)

		}
	}()

	return sink

}

func (scraper *aTokenScraper) watchReservedataupdatedV2() chan *aavepool2.Aavepool2ReserveDataUpdated {
	filterer, err := aavepool2.NewAavepool2Filterer(common.HexToAddress(scraper.pooladdress), scraper.WsClient)

	if err != nil {
		log.Error("new NewAavepool3Filterer caller: ", err)
	}
	sink := make(chan *aavepool2.Aavepool2ReserveDataUpdated)
	// if start is zero start from latest
	if scraper.start == 0 {
		log.Error("WatchReserveDataUpdated: ", scraper.start)

		_, err := filterer.WatchReserveDataUpdated(&bind.WatchOpts{}, sink, []common.Address{})
		if err != nil {
			log.Error("Error on filterer.WatchReserveDataUpdated: ", err)

		}

	} else {
		log.Error("WatchReserveDataUpdated starting from block: ", scraper.start)

		filterer.WatchReserveDataUpdated(&bind.WatchOpts{Start: &scraper.start}, sink, []common.Address{})
	}
	return sink

}

func (scraper *aTokenScraper) mainLoop() {
	// runs whenever there is  ReserveDataUpdated occurs

	sink := scraper.watchReservedataupdatedV2()

	go func() {
		for {
			dataupdated, ok := <-sink
			if ok {
				log.Infoln("Reserve Data updated fetch supply and reserver of token ", dataupdated.Raw.TxHash.Hex())
				atokenasset, underlyingasset, atokensupply, reserver, totalDebt, err := scraper.fetchsupplyandbalance(dataupdated.Reserve.Hex(), big.NewInt(int64(dataupdated.Raw.BlockNumber)))
				if err != nil {
					log.Errorln("err on fetchsupplyandbalance", err)
					continue
				}

				block, err := scraper.RestClient.BlockByNumber(context.Background(), big.NewInt(int64(dataupdated.Raw.BlockNumber)))
				if err != nil {
					log.Errorln("error getting block details", err)
				}
				collateralRatio := (reserver + totalDebt) / atokensupply

				sas := dia.SynthAssetSupply{Asset: atokenasset, AssetUnderlying: underlyingasset, Supply: atokensupply, LockedUnderlying: reserver, BlockNumber: dataupdated.Raw.BlockNumber, Time: time.Unix(int64(block.Time()), 0), ColleteralRatio: collateralRatio, Protocol: scraper.protocol, TotalDebt: totalDebt}
				scraper.synthChannel <- sas
			}
		}
	}()

}

func (scraper *aTokenScraper) mainLoopHistory() {
	// runs whenever there is  ReserveDataUpdated occurs

	log.Info("looking for filterReservedataupdatedV2History")
	sink := scraper.filterReservedataupdatedV2History()

	go func() {
		for {
			dataupdated, ok := <-sink
			if ok {
				log.Infoln("Reserve Data updated fetch supply and reserver of token ", dataupdated.Raw.TxHash.Hex())
				atokenasset, underlyingasset, atokensupply, reserver, totalDebt, err := scraper.fetchsupplyandbalance(dataupdated.Reserve.Hex(), big.NewInt(int64(dataupdated.Raw.BlockNumber)))
				if err != nil {
					log.Errorln("err on fetchsupplyandbalance", err)
					continue
				}

				block, err := scraper.RestClient.BlockByNumber(context.Background(), big.NewInt(int64(dataupdated.Raw.BlockNumber)))
				if err != nil {
					log.Errorln("error getting block details", err)
				}
				collateralRatio := (reserver + totalDebt) / atokensupply

				sas := dia.SynthAssetSupply{Asset: atokenasset, AssetUnderlying: underlyingasset, Supply: atokensupply, LockedUnderlying: reserver, BlockNumber: dataupdated.Raw.BlockNumber, Time: time.Unix(int64(block.Time()), 0), ColleteralRatio: collateralRatio, Protocol: scraper.protocol, TotalDebt: totalDebt}
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

		scraper.atokens[asset.Hex()] = atokendetail{address: data.ATokenAddress.Hex(), stabledebttokenaddress: data.StableDebtTokenAddress.Hex(), variabledebttokenaddress: data.VariableDebtTokenAddress.Hex()}

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

		scraper.atokens[asset.Hex()] = atokendetail{address: data.ATokenAddress.Hex(), stabledebttokenaddress: data.StableDebtTokenAddress.Hex(), variabledebttokenaddress: data.VariableDebtTokenAddress.Hex()}

	}

	log.Info("total assets in pool ", len(scraper.atokens))

}

func (scraper *aTokenScraper) GetSynthSupplyChannel() chan dia.SynthAssetSupply {
	return scraper.synthChannel
}

func (scraper *aTokenScraper) FetchSynthSupply() error {

	return nil
}

func (scraper *aTokenScraper) fetchsupplyandbalance(underlyingtokenaddress string, blocknumber *big.Int) (atokenasset, underlyingasset dia.Asset, atokensupply, underlyingreserve, totalDebt float64, err error) {

	atokendetail := scraper.atokens[underlyingtokenaddress]

	log.Info("atokenaddress", atokendetail.address)
	log.Info("underlyingtokenaddress", underlyingtokenaddress)

	filterer, err := ceth.NewERC20Caller(common.HexToAddress(atokendetail.address), scraper.WsClient)
	if err != nil {
		log.Error("new erc20 caller: ", err)
		return
	}
	supply, err := filterer.TotalSupply(&bind.CallOpts{BlockNumber: blocknumber})
	if err != nil {
		log.Error("get supply: ", err)
		return
	}

	atokendecimal, err := filterer.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error("get Decimals: ", err)
		return
	}
	atokensymbol, err := filterer.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error("get Decimals: ", err)
		return
	}

	atokenasset.Symbol = atokensymbol
	atokenasset.Address = atokendetail.address
	atokenasset.Decimals = uint8(atokendecimal.Int64())
	atokenasset.Blockchain = scraper.blockchain

	atokensupply, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(supply), new(big.Float).SetFloat64(math.Pow10(int(atokendecimal.Int64())))).Float64()

	underlyingfilterer, err := ceth.NewERC20Caller(common.HexToAddress(underlyingtokenaddress), scraper.RestClient)
	if err != nil {
		log.Error("new erc20 caller: ", err)
		return
	}
	balanceof, err := underlyingfilterer.BalanceOf(&bind.CallOpts{BlockNumber: blocknumber}, common.HexToAddress(atokendetail.address))
	if err != nil {
		log.Error("get balanceof: ", err)
		return
	}
	log.Info("balanceof: ", balanceof)

	underlyingdecimals, err := underlyingfilterer.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error("get Decimals: ", err)
		return
	}
	underlyingsymbol, err := underlyingfilterer.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Error("get Symbol: ", err)
		return
	}

	underlyingreserve, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(balanceof), new(big.Float).SetFloat64(math.Pow10(int(underlyingdecimals.Int64())))).Float64()
	underlyingasset.Symbol = underlyingsymbol

	underlyingasset.Address = underlyingtokenaddress
	underlyingasset.Decimals = uint8(underlyingdecimals.Int64())
	underlyingasset.Blockchain = scraper.blockchain

	stabledebtTokenSupply, err := scraper.getTokenSupply(atokendetail.stabledebttokenaddress, blocknumber)
	if err != nil {
		log.Error("getTokenSupply: ", err)
		return
	}

	log.Infof("total supply of %v is %v ", atokendetail.stabledebttokenaddress, stabledebtTokenSupply)

	variabledebtTokenSupply, err := scraper.getTokenSupply(atokendetail.variabledebttokenaddress, blocknumber)
	if err != nil {
		log.Error("getTokenSupply: ", err)
		return
	}

	log.Infof("total supply of %v is %v ", atokendetail.variabledebttokenaddress, variabledebtTokenSupply)

	totalDebt = stabledebtTokenSupply + variabledebtTokenSupply

	return

}

func (scraper *aTokenScraper) getTokenSupply(address string, blocknumber *big.Int) (float64, error) {

	tokenfilterer, err := ceth.NewERC20Caller(common.HexToAddress(address), scraper.WsClient)
	if err != nil {
		log.Error("new erc20 caller: ", err)
		return 0.0, err

	}
	decimal, err := tokenfilterer.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Error("get Decimals: ", err)
		return 0.0, err

	}
	supply, err := tokenfilterer.TotalSupply(&bind.CallOpts{BlockNumber: blocknumber})
	if err != nil {
		log.Error("get Decimals: ", err)
		return 0.0, err
	}
	atokensupply, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(supply), new(big.Float).SetFloat64(math.Pow10(int(decimal.Int64())))).Float64()
	return atokensupply, nil
}
