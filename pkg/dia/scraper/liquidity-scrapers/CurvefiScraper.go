package liquidityscrapers

import (
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefifactory"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefimeta"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	curveRestDial = ""
)

type CurveFIScraper struct {
	RestClient   *ethclient.Client
	relDB        *models.RelDB
	poolChannel  chan dia.Pool
	doneChannel  chan bool
	blockchain   string
	exchangeName string
}

type curveRegistry struct {
	Address common.Address
	Type    int
}

type CurveCoin struct {
	Symbol   string
	Decimals uint8
	Address  string
	Name     string
}

func NewCurveFIScraper(exchange dia.Exchange) *CurveFIScraper {
	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *CurveFIScraper
		registries  []curveRegistry
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", curveRestDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	switch exchange.Name {
	case dia.CurveFIExchange:
		basePools := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		cryptoswapPools := curveRegistry{Type: 1, Address: common.HexToAddress("0x8F942C20D02bEfc377D41445793068908E2250D0")}
		metaPools := curveRegistry{Type: 2, Address: common.HexToAddress("0xB9fC157394Af804a3578134A6585C0dc9cc990d4")}
		factoryPools := curveRegistry{Type: 3, Address: common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99")}
		registries = []curveRegistry{basePools, cryptoswapPools, metaPools, factoryPools}
	case dia.CurveFIExchangeFantom:
		exchange.Contract = ""
		// basePools := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x686d67265703D1f124c45E33d47d794c566889Ba")}
		registries = []curveRegistry{stableSwapFactory}
	case dia.CurveFIExchangeMoonbeam:
		// basePools := curveRegistry{Type: 1, Address: common.HexToAddress(exchange.Contract)}
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x4244eB811D6e0Ef302326675207A95113dB4E1F8")}
		registries = []curveRegistry{stableSwapFactory}
	case dia.CurveFIExchangePolygon:
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0x722272D36ef0Da72FF51c5A65Db7b870E2e8D4ee")}
		registries = []curveRegistry{stableSwapFactory}
	case dia.CurveFIExchangeArbitrum:
		stableSwapFactory := curveRegistry{Type: 2, Address: common.HexToAddress("0xb17b674D9c5CB2e441F8e196a2f048A81355d031")}
		registries = []curveRegistry{stableSwapFactory}

	}

	scraper = &CurveFIScraper{
		RestClient:   restClient,
		poolChannel:  poolChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		exchangeName: exchange.Name,
	}
	scraper.relDB, err = models.NewRelDataStore()
	if err != nil {
		log.Fatal("new postgres datastore: ", err)
	}

	go func() {
		for _, registry := range registries {
			poolAddresses, err := scraper.fetchPoolAddresses(registry)
			if err != nil {
				log.Error("fetch pool addresses: ", err)
			}
			for _, poolAddress := range poolAddresses {
				scraper.loadPoolData(poolAddress.Hex(), registry)
			}
		}
		scraper.doneChannel <- true
	}()

	return scraper
}

func (scraper *CurveFIScraper) fetchPoolAddresses(registry curveRegistry) (poolAddresses []common.Address, err error) {

	if registry.Type == 1 {
		log.Info("load base type pools..")
		var (
			contract  *curvefi.CurvefiCaller
			poolCount *big.Int
		)
		contract, err = curvefi.NewCurvefiCaller(registry.Address, scraper.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}

		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
		}
		log.Info("poolCount: ", int(poolCount.Int64()))
		for i := 0; i < int(poolCount.Int64()); i++ {
			poolAddress, errPool := contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if errPool != nil {
				log.Error("PoolList: ", err)
			}
			poolAddresses = append(poolAddresses, poolAddress)
		}
	}

	if registry.Type == 2 || registry.Type == 3 {
		log.Info("load meta / factory type pools...")
		var (
			contract  *curvefimeta.CurvefimetaCaller
			poolCount *big.Int
		)
		contract, err = curvefimeta.NewCurvefimetaCaller(registry.Address, scraper.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}
		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
		}
		log.Info("poolCount: ", int(poolCount.Int64()))
		for i := 0; i < int(poolCount.Int64()); i++ {
			poolAddress, err := contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Error("PoolList: ", err)
			}
			poolAddresses = append(poolAddresses, poolAddress)
		}
	}
	return
}

func (scraper *CurveFIScraper) loadPoolData(poolAddress string, registry curveRegistry) {
	// We need to handle ETH with address 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE because of pools such as:
	// https://etherscan.io/address/0xDC24316b9AE028F1497c275EB9192a3Ea0f67022#readContract

	var (
		poolCoins [8]common.Address
		pool      dia.Pool
	)
	pool.Blockchain = dia.BlockChain{Name: scraper.blockchain}
	pool.Exchange = dia.Exchange{
		Name:        scraper.exchangeName,
		Centralized: false,
		Bridge:      false,
		BlockChain:  pool.Blockchain,
	}
	pool.Address = poolAddress

	if registry.Type == 1 {
		contract, err := curvefi.NewCurvefiCaller(registry.Address, scraper.RestClient)
		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
			return
		}

		poolCoins, err = contract.GetCoins(&bind.CallOpts{}, common.HexToAddress(poolAddress))
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
			return
		}

		liquidities, err := contract.GetBalances(&bind.CallOpts{}, common.HexToAddress(poolAddress))
		if err != nil {
			log.Error("loadPoolData - GetBalances: ", err)
			return
		}

		var i int
		for i < 8 && poolCoins[i] != (common.Address{}) {
			if poolCoins[i].Hex() == "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE" {
				poolCoins[i] = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
			}
			asset, err := scraper.relDB.GetAsset(poolCoins[i].Hex(), scraper.blockchain)
			if err != nil {
				asset, err = ethhelper.ETHAddressToAsset(poolCoins[i], scraper.RestClient, scraper.blockchain)
				if err != nil {
					return
				}
			}
			liquidity, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(liquidities[i]), new(big.Float).SetFloat64(math.Pow10(int(asset.Decimals)))).Float64()
			av := dia.AssetVolume{Asset: asset, Volume: liquidity, Index: uint8(i)}
			pool.Assetvolumes = append(pool.Assetvolumes, av)
			pool.Time = time.Now()
			i++
		}

		scraper.poolChannel <- pool

	}

	if registry.Type == 2 {
		contract, err := curvefimeta.NewCurvefimetaCaller(registry.Address, scraper.RestClient)
		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
			return
		}

		poolCoins, err := contract.GetCoins(&bind.CallOpts{}, common.HexToAddress(poolAddress))
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
			return
		}

		liquidities, err := contract.GetBalances(&bind.CallOpts{}, common.HexToAddress(poolAddress))
		if err != nil {
			log.Error("loadPoolData - GetBalances: ", err)
			return
		}

		var i int
		for i < 4 && poolCoins[i] != (common.Address{}) {
			if poolCoins[i].Hex() == "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE" {
				poolCoins[i] = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
			}
			asset, err := scraper.relDB.GetAsset(poolCoins[i].Hex(), scraper.blockchain)
			if err != nil {
				asset, err = ethhelper.ETHAddressToAsset(poolCoins[i], scraper.RestClient, scraper.blockchain)
				if err != nil {
					return
				}
			}
			liquidity, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(liquidities[i]), new(big.Float).SetFloat64(math.Pow10(int(asset.Decimals)))).Float64()
			av := dia.AssetVolume{Asset: asset, Volume: liquidity, Index: uint8(i)}
			pool.Assetvolumes = append(pool.Assetvolumes, av)
			pool.Time = time.Now()
			i++
		}

		scraper.poolChannel <- pool

	}
	if registry.Type == 3 {
		contract, err := curvefifactory.NewCurvefifactoryCaller(common.HexToAddress(poolAddress), scraper.RestClient)
		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}

		var i int64
		for i < 8 {
			poolAssetAddress, err := contract.Coins(&bind.CallOpts{}, big.NewInt(i))
			if err != nil {
				i++
				continue
			}
			if poolAssetAddress == (common.Address{}) {
				i++
				continue
			}
			if poolAssetAddress.Hex() == "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE" {
				poolAssetAddress = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
			}

			liquidityBig, err := contract.Balances(&bind.CallOpts{}, big.NewInt(i))
			if err != nil {
				log.Error("Get Balances: ", err)
			}

			asset, err := scraper.relDB.GetAsset(poolCoins[i].Hex(), scraper.blockchain)
			if err != nil {
				asset, err = ethhelper.ETHAddressToAsset(poolCoins[i], scraper.RestClient, scraper.blockchain)
				if err != nil {
					return
				}
			}
			liquidity, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(liquidityBig), new(big.Float).SetFloat64(math.Pow10(int(asset.Decimals)))).Float64()
			av := dia.AssetVolume{Asset: asset, Volume: liquidity, Index: uint8(i)}
			pool.Assetvolumes = append(pool.Assetvolumes, av)
			pool.Time = time.Now()
			i++
		}

		scraper.poolChannel <- pool

	}

}

func (scraper *CurveFIScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *CurveFIScraper) Done() chan bool {
	return scraper.doneChannel
}

// getAssetFromCache returns the full asset given by blockchain and address, either from the map @localCache
// or from Postgres, in which latter case it also adds the asset to the local cache.
// Remember that maps are always passed by reference.
func (scraper *CurveFIScraper) getAssetFromCache(localCache map[string]dia.Asset, blockchain string, address string) dia.Asset {
	if asset, ok := localCache[assetIdentifier(blockchain, address)]; ok {
		return asset
	}
	fullAsset, err := scraper.relDB.GetAsset(address, blockchain)
	if err != nil {
		log.Warnf("could not find asset with address %s on blockchain %s in postgres: ", address, blockchain)
	}
	localCache[assetIdentifier(blockchain, address)] = fullAsset
	return fullAsset
}

func assetIdentifier(blockchain string, address string) string {
	return blockchain + "-" + address
}
