package source

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi/token"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefifactory"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefimeta"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	curveRestDial = ""

	polygonWaitMilliseconds  = "500"
	fantomWaitMilliseconds   = "250"
	arbitrumWaitMilliseconds = "300"
)

type curveRegistry struct {
	Address common.Address
	Type    int
}

// CurvefiAssetSource is a curve finance scraper on a specific blockchain.
type CurvefiAssetSource struct {
	RestClient   *ethclient.Client
	assetChannel chan dia.Asset
	doneChannel  chan bool
	blockchain   string
	registries   []curveRegistry
	waitTime     int
}

func NewCurvefiAssetSource(exchange dia.Exchange) *CurvefiAssetSource {

	var cas *CurvefiAssetSource

	switch exchange.Name {
	case dia.CurveFIExchange:
		basePools := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		twocryptoNGPools := curveRegistry{Type: 3, Address: common.HexToAddress("0x98EE851a00abeE0d95D08cF4CA2BdCE32aeaAF7F")}
		cryptoswapPools := curveRegistry{Type: 1, Address: common.HexToAddress("0x8F942C20D02bEfc377D41445793068908E2250D0")}
		metaPools := curveRegistry{Type: 2, Address: common.HexToAddress("0xB9fC157394Af804a3578134A6585C0dc9cc990d4")}
		stableSwapRegistry := curveRegistry{Type: 3, Address: common.HexToAddress("0x6A8cbed756804B16E05E741eDaBd5cB544AE21bf")}
		factoryPools := curveRegistry{Type: 3, Address: common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99")}
		factory2Pools := curveRegistry{Type: 3, Address: common.HexToAddress("0x4F8846Ae9380B90d2E71D5e3D042dff3E7ebb40d")}
		registries := []curveRegistry{factoryPools, twocryptoNGPools, factory2Pools, basePools, cryptoswapPools, metaPools, stableSwapRegistry}
		cas = makeCurvefiAssetSource(exchange, registries, curveRestDial, uniswapWaitMilliseconds)
	case dia.CurveFIExchangeFantom:
		exchange.Contract = ""
		// basePools := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x686d67265703D1f124c45E33d47d794c566889Ba")}
		registries := []curveRegistry{stableSwapFactory}
		cas = makeCurvefiAssetSource(exchange, registries, curveRestDial, fantomWaitMilliseconds)

	case dia.CurveFIExchangeMoonbeam:
		exchange.Contract = ""
		// basePools := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x4244eB811D6e0Ef302326675207A95113dB4E1F8")}
		registries := []curveRegistry{stableSwapFactory}
		cas = makeCurvefiAssetSource(exchange, registries, curveRestDial, moonbeamWaitMilliseconds)

	case dia.CurveFIExchangePolygon:
		exchange.Contract = ""
		// basePools := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x722272D36ef0Da72FF51c5A65Db7b870E2e8D4ee")}
		registries := []curveRegistry{stableSwapFactory}
		cas = makeCurvefiAssetSource(exchange, registries, curveRestDial, polygonWaitMilliseconds)

	case dia.CurveFIExchangeArbitrum:
		exchange.Contract = ""
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0xb17b674D9c5CB2e441F8e196a2f048A81355d031")}
		stableSwapNGFactory := curveRegistry{Type: 3, Address: common.HexToAddress("0x9AF14D26075f142eb3F292D5065EB3faa646167b")}
		registries := []curveRegistry{stableSwapFactory, stableSwapNGFactory}
		cas = makeCurvefiAssetSource(exchange, registries, curveRestDial, arbitrumWaitMilliseconds)

	case dia.CurveFIExchangeSonic:
		stableSwapFactory := curveRegistry{Type: 3, Address: common.HexToAddress("0x7C2085419BE6a04f4ad88ea91bC9F5C6E6C463D8")}
		registries := []curveRegistry{stableSwapFactory}
		cas = makeCurvefiAssetSource(exchange, registries, curveRestDial, arbitrumWaitMilliseconds)

	}

	go func() {
		cas.fetchAssets()
	}()
	return cas
}

// makeCurvefiAssetSource returns a curve finance scraper as used in NewCurvefiScraper.
func makeCurvefiAssetSource(exchange dia.Exchange, registries []curveRegistry, restDial string, waitMilliseconds string) *CurvefiAssetSource {
	var (
		restClient   *ethclient.Client
		err          error
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		cas          *CurvefiAssetSource
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", restDial))
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
	cas = &CurvefiAssetSource{
		RestClient:   restClient,
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		waitTime:     waitTime,
		registries:   registries,
	}

	return cas
}

func (cas *CurvefiAssetSource) Asset() chan dia.Asset {
	return cas.assetChannel
}

func (cas *CurvefiAssetSource) Done() chan bool {
	return cas.doneChannel
}

func (cas *CurvefiAssetSource) fetchAssets() {
	// Load pools from registries.
	for _, registry := range cas.registries {
		err := cas.loadPoolsAndCoins(registry)
		if err != nil {
			log.Error("loadPoolsAndCoins: ", err)
		}
	}
	cas.doneChannel <- true
}

// contract.poolList.map(contract.GetPoolCoins(pool).)
func (cas *CurvefiAssetSource) loadPoolsAndCoins(registry curveRegistry) (err error) {

	if registry.Type == 1 {
		log.Info("load base type pools..")
		var contract *curvefi.CurvefiCaller
		var poolCount *big.Int
		contract, err = curvefi.NewCurvefiCaller(registry.Address, cas.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}
		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
		}
		log.Info("poolCount: ", int(poolCount.Int64()))
		for i := 0; i < int(poolCount.Int64()); i++ {
			var pool common.Address
			pool, err = contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Error("PoolList: ", err)
			}

			err = cas.loadPoolData(pool.Hex(), registry)
			if err != nil {
				log.Info("load pool data: ", err)
				// return err
			}
		}
	}

	if registry.Type == 2 || registry.Type == 3 {
		log.Info("load meta type pools...")
		var contract *curvefimeta.CurvefimetaCaller
		var poolCount *big.Int
		contract, err = curvefimeta.NewCurvefimetaCaller(registry.Address, cas.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}

		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
		}
		log.Info("poolCount: ", int(poolCount.Int64()))
		for i := 0; i < int(poolCount.Int64()); i++ {
			var pool common.Address
			pool, err = contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Error("PoolList: ", err)
			}

			err = cas.loadPoolData(pool.Hex(), registry)
			if err != nil {
				log.Error("load pool data: ", err)
				// return err
			}
		}
	}
	return err
}

func (cas *CurvefiAssetSource) loadPoolData(pool string, registry curveRegistry) error {
	var poolCoins [8]common.Address

	if registry.Type == 1 {
		contract, err := curvefi.NewCurvefiCaller(registry.Address, cas.RestClient)
		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}
		poolCoins, err = contract.GetCoins(&bind.CallOpts{}, common.HexToAddress(pool))
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}

		poolName, err := contract.GetPoolName(&bind.CallOpts{}, common.HexToAddress(pool))
		if err != nil {
			log.Error("loadPoolData - GetPoolName: ", err)
		}
		log.Info("pool name: ", poolName)
	}

	if registry.Type == 2 {
		// common.HexToAddress(curveFiMetaPoolsFactory) || factoryContract == common.HexToAddress(curveFiCryptoPoolFactory) {
		contract, err := curvefimeta.NewCurvefimetaCaller(registry.Address, cas.RestClient)

		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}

		aux, err := contract.GetCoins(&bind.CallOpts{}, common.HexToAddress(pool))
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}
		// GetCoins on meta pools returns [4]common.Address instead of [8]common.Address for standard pools.
		//nolint
		copy(aux[:], poolCoins[:])

	}
	if registry.Type == 3 {
		log.Info("pool type 3...")
		contract, err := curvefifactory.NewCurvefifactoryCaller(common.HexToAddress(pool), cas.RestClient)
		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}

		var i int64
		poolAssetAddress, err := contract.Coins(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}
		for poolAssetAddress != (common.Address{}) && i < 8 {
			poolCoins[i] = poolAssetAddress
			i++
			poolAssetAddress, err = contract.Coins(&bind.CallOpts{}, big.NewInt(i))
			if err != nil {
				log.Warn("loadPoolData - GetCoins: ", err)
			}
		}

	}

	var err error
	checkMap := make(map[string]struct{})

	for _, c := range poolCoins {
		time.Sleep(time.Duration(cas.waitTime) * time.Millisecond)
		var (
			coinCaller *token.TokenCaller
			decimals   *big.Int
			asset      dia.Asset
		)
		if c == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			continue
		} else {
			asset.Address = c.Hex()
			coinCaller, err = token.NewTokenCaller(c, cas.RestClient)
			if err != nil {
				log.Error("NewTokenCaller: ", err)
				continue
			}
			asset.Symbol, err = coinCaller.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Error("Symbol: ", err, c.Hex())
				continue
			}
			decimals, err = coinCaller.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Error("Decimals: ", err)
				continue
			}
			asset.Decimals = uint8(decimals.Int64())

			asset.Name, err = coinCaller.Name(&bind.CallOpts{})
			if err != nil {
				log.Error("Name: ", err)
				continue
			}
		}
		asset.Blockchain = cas.blockchain

		if _, ok := checkMap[asset.Address]; !ok {
			if asset.Symbol != "" {
				checkMap[asset.Address] = struct{}{}
				cas.assetChannel <- asset
			}
		}

	}
	return err
}
