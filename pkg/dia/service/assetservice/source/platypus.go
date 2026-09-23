package source

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance"
	platypusAssetABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/asset"
	platypusPoolABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/pool"
	platypusTokenABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/token"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	platypusRestDialAvalanche = "https://api.avax.network/ext/bc/C/rpc"
	platypusMasterRegV3Addr   = "0x7125B4211357d7C3a90F796c956c12c681146EbB"
	platypusMasterRegV2Addr   = "0x68c5f4374228BEEdFa078e77b5ed93C28a2f713E"
	platypusMasterRegV1Addr   = "0xB0523f9F473812FB195Ee49BC7d2ab9873a98044"
	waitMilliseconds          = ""
)

type platypusRegistry struct {
	Address common.Address
	Version int
}

type PlatypusCoin struct {
	Symbol   string
	Decimals uint8
	Address  string
	Name     string
}

// The scraper object for Platypus Finance.
type PlatypusAssetSource struct {
	RestClient   *ethclient.Client
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
	registries   []platypusRegistry
	waitTime     int
}

// Returns a new platypus asset scraper.
func NewPlatypusScraper(exchange dia.Exchange) *PlatypusAssetSource {

	var (
		restClient   *ethclient.Client
		err          error
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		pas          *PlatypusAssetSource
	)

	registries := []platypusRegistry{
		{Version: 3, Address: common.HexToAddress(platypusMasterRegV3Addr)},
		{Version: 2, Address: common.HexToAddress(platypusMasterRegV2Addr)},
		{Version: 1, Address: common.HexToAddress(platypusMasterRegV1Addr)},
	}

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", platypusRestDialAvalanche))
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
	pas = &PlatypusAssetSource{
		RestClient:   restClient,
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		waitTime:     waitTime,
		registries:   registries,
	}
	go func() {
		pas.fetchAssets()
	}()

	return pas
}

func (pas *PlatypusAssetSource) fetchAssets() {
	for _, registry := range pas.registries {
		err := pas.fetchPools(registry)
		if err != nil {
			log.Error("loadPoolsAndCoins: ", err)
		}
	}
	pas.doneChannel <- true
}

// Load pools and coins metadata from master registry
func (scraper *PlatypusAssetSource) fetchPools(registry platypusRegistry) error {
	log.Infof("loading master contract %s version %d and querying registry", registry.Address.Hex(), registry.Version)
	contractMaster, err := platypusfinance.NewBaseMasterPlatypusCaller(registry.Address, scraper.RestClient)
	if err != nil {
		log.Error("NewBaseMasterPlatypusCaller: ", err)
		return err
	}

	poolCount, err := contractMaster.PoolLength(&bind.CallOpts{})
	if err != nil {
		log.Error("PoolLength: ", err)
		return err
	}

	for i := 0; i < int(poolCount.Int64()); i++ {
		asset, errPoolInfo := contractMaster.PoolInfo(&bind.CallOpts{}, big.NewInt(int64(i)))
		if errPoolInfo != nil {
			log.Error("PoolInfo: ", errPoolInfo)
			return errPoolInfo
		}

		errPoolData := scraper.loadPoolData(asset.LpToken.Hex())
		if errPoolData != nil {
			log.Errorf("loadPoolData error at asset %s: %s", asset.LpToken.Hex(), errPoolData)
			return errPoolData
		}

	}

	return err
}

func (scraper *PlatypusAssetSource) loadPoolData(poolAddress string) (err error) {

	contractAsset, err := platypusAssetABI.NewAssetCaller(common.HexToAddress(poolAddress), scraper.RestClient)
	if err != nil {
		log.Error("NewAssetCaller: ", err)
	}

	poolContract, err := contractAsset.Pool(&bind.CallOpts{})
	if err != nil {
		log.Error("Pool: ", err)
	}

	poolCaller, err := platypusPoolABI.NewPoolCaller(poolContract, scraper.RestClient)
	if err != nil {
		log.Error("NewPoolCaller: ", err)
	}

	poolTokenAddresses, errGetTokens := poolCaller.GetTokenAddresses(&bind.CallOpts{})
	if errGetTokens != nil {
		symbol, err := contractAsset.Symbol(&bind.CallOpts{})
		if err != nil {
			log.Error("contractAsset.Symbol: ", err)
		}
		log.Warnf("error calling GetTokenAddresses for %s %s asset: %s", symbol, poolAddress, errGetTokens)
	}

	for _, c := range poolTokenAddresses {
		var (
			symbol      string
			decimalsBig *big.Int
			decimals    uint8
			name        string
		)

		if c == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			continue
		} else {
			contractToken, err := platypusTokenABI.NewTokenCaller(c, scraper.RestClient)
			if err != nil {
				log.Error("NewTokenCaller: ", err)
				continue
			}

			symbol, err = contractToken.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Error("Symbol: ", err, c.Hex())
				continue
			}

			decimalsBig, err = contractToken.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Error("Decimals: ", err)
				continue
			}
			decimals = uint8(decimalsBig.Uint64())

			name, err = contractToken.Name(&bind.CallOpts{})
			if err != nil {
				log.Error("Name: ", err)
				continue
			}

		}

		poolAsset := dia.Asset{
			Address:    c.Hex(),
			Blockchain: scraper.blockchain,
			Decimals:   decimals,
			Symbol:     symbol,
			Name:       name,
		}

		scraper.Asset() <- poolAsset

	}
	return
}

func (pas *PlatypusAssetSource) Asset() chan dia.Asset {
	return pas.assetChannel
}

func (pas *PlatypusAssetSource) Done() chan bool {
	return pas.doneChannel
}
