package liquidityscrapers

import (
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefi/token"
	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/curvefimeta"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	curveFiMetaPoolsFactory = "0xB9fC157394Af804a3578134A6585C0dc9cc990d4"
	curveFiCryptoswapPools  = "0x8F942C20D02bEfc377D41445793068908E2250D0"
	curveRestDial           = ""
)

type CurveFIScraper struct {
	RestClient   *ethclient.Client
	poolChannel  chan dia.Pool
	doneChannel  chan bool
	blockchain   string
	exchangeName string
	poolAddrs    []string
}

func NewCurveFIScraper(exchange dia.Exchange) *CurveFIScraper {
	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *CurveFIScraper
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", curveRestDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	switch exchange.Name {
	case dia.CurveFIExchange:
		curveFiMetaPoolsFactory = "0xB9fC157394Af804a3578134A6585C0dc9cc990d4"
	case dia.CurveFIExchangeFantom:
		curveFiMetaPoolsFactory = "0x686d67265703D1f124c45E33d47d794c566889Ba"
	case dia.CurveFIExchangeMoonbeam:
		curveFiMetaPoolsFactory = "0x4244eB811D6e0Ef302326675207A95113dB4E1F8"
	case dia.CurveFIExchangePolygon:
		curveFiMetaPoolsFactory = "0x722272d36ef0da72ff51c5a65db7b870e2e8d4ee"
	case dia.CurveFIExchangeArbitrum:
		curveFiMetaPoolsFactory = "0xb17b674D9c5CB2e441F8e196a2f048A81355d031"

	}

	scraper = &CurveFIScraper{
		RestClient:   restClient,
		poolChannel:  poolChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		exchangeName: exchange.Name,
		poolAddrs:    []string{curveFiMetaPoolsFactory, curveFiCryptoswapPools, exchange.Contract},
	}

	go func() {
		for _, address := range scraper.poolAddrs {
			scraper.fetchPools(common.HexToAddress(address))
		}
		scraper.doneChannel <- true
	}()

	return scraper
}

// fetchPools collects all available pools and sends them into the pool channel.
func (scraper *CurveFIScraper) fetchPools(factoryAddress common.Address) {
	var pool dia.Pool

	if factoryAddress == common.HexToAddress(curveFiMetaPoolsFactory) || factoryAddress == common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99") {
		log.Info("load meta pools...")
		var (
			contract  *curvefimeta.CurvefimetaCaller
			poolCount *big.Int
			err       error
		)

		contract, err = curvefimeta.NewCurvefimetaCaller(factoryAddress, scraper.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}
		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
		}

		for i := 0; i < int(poolCount.Int64()); i++ {
			poolAddress, err := contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Error("PoolList: ", err)
			}

			pool = scraper.loadPoolData(poolAddress, factoryAddress)
			if err != nil {
				log.Info("load pool data: ", err)
			}
			scraper.poolChannel <- pool
		}
	} else {
		log.Info("load factory pools...")
		var (
			contract  *curvefi.CurvefiCaller
			poolCount *big.Int
			err       error
		)

		contract, err = curvefi.NewCurvefiCaller(factoryAddress, scraper.RestClient)
		if err != nil {
			log.Error("NewCurvefiCaller: ", err)
		}
		poolCount, err = contract.PoolCount(&bind.CallOpts{})
		if err != nil {
			log.Error("PoolCount: ", err)
			return
		}

		for i := 0; i < int(poolCount.Int64()); i++ {
			poolAddress, err := contract.PoolList(&bind.CallOpts{}, big.NewInt(int64(i)))
			if err != nil {
				log.Error("PoolList: ", err)
			}

			pool = scraper.loadPoolData(poolAddress, factoryAddress)
			if err != nil {
				log.Info("load pool data: ", err)
			}
			scraper.poolChannel <- pool
		}
	}
}

func (scraper *CurveFIScraper) loadPoolData(poolAddress common.Address, factoryContract common.Address) (pool dia.Pool) {
	var (
		poolCoinAddresses [8]common.Address
		poolAssets        []dia.Asset
		poolBalances      [8]*big.Int
	)

	if factoryContract == common.HexToAddress(curveFiMetaPoolsFactory) || factoryContract == common.HexToAddress("0xF18056Bbd320E96A48e3Fbf8bC061322531aac99") {
		contract, err := curvefimeta.NewCurvefimetaCaller(factoryContract, scraper.RestClient)

		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}

		aux, err := contract.GetCoins(&bind.CallOpts{}, poolAddress)
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}
		// GetCoins on meta pools returns [4]common.Address instead of [8]common.Address for standard pools.
		copy(poolCoinAddresses[:], aux[:])

		bal, err := contract.GetBalances(&bind.CallOpts{}, poolAddress)
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}
		copy(poolBalances[:], bal[:])

	} else {
		contract, err := curvefi.NewCurvefiCaller(factoryContract, scraper.RestClient)
		if err != nil {
			log.Error("loadPoolData - NewCurvefiCaller: ", err)
		}

		poolCoinAddresses, err = contract.GetCoins(&bind.CallOpts{}, poolAddress)
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}

		bal, err := contract.GetBalances(&bind.CallOpts{}, poolAddress)
		if err != nil {
			log.Error("loadPoolData - GetCoins: ", err)
		}

		copy(poolBalances[:], bal[:])

	}

	var err error
	for _, c := range poolCoinAddresses {
		var (
			coinCaller  *token.TokenCaller
			symbol      string
			decimals    uint8
			decimalsBig *big.Int
			name        string
		)
		if c == common.HexToAddress("0x0000000000000000000000000000000000000000") {
			continue
		} else if c == common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE") {
			symbol = "ETH"
			decimals = uint8(18)
			name = "Ether"
			c = common.HexToAddress("0x0000000000000000000000000000000000000000")
		} else {
			coinCaller, err = token.NewTokenCaller(c, scraper.RestClient)
			if err != nil {
				log.Error("NewTokenCaller: ", err)
				continue
			}
			symbol, err = coinCaller.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Error("Symbol: ", err, c.Hex())
				continue
			}
			decimalsBig, err = coinCaller.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Error("Decimals: ", err)
				continue
			}
			decimals = uint8(decimalsBig.Uint64())
			name, err = coinCaller.Name(&bind.CallOpts{})
			if err != nil {
				log.Error("Name: ", err)
				continue
			}
		}
		log.Info(symbol, " ", decimals, " ", "'", name, "'", " ", c)

		poolAssets = append(poolAssets, dia.Asset{
			Address:    c.Hex(),
			Blockchain: scraper.blockchain,
			Decimals:   decimals,
			Symbol:     symbol,
			Name:       name,
		})

	}

	for i := range poolAssets {
		var volume float64
		if poolBalances[i] != nil {
			volume, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(poolBalances[i]), new(big.Float).SetFloat64(math.Pow10(int(poolAssets[i].Decimals)))).Float64()
		}
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{
			Asset:  poolAssets[i],
			Volume: volume,
			Index:  uint8(i),
		})
	}
	pool.Exchange = dia.Exchange{Name: scraper.exchangeName}
	pool.Blockchain = dia.BlockChain{Name: scraper.blockchain}
	pool.Address = poolAddress.Hex()
	pool.Time = time.Now()

	return pool
}

func (scraper *CurveFIScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *CurveFIScraper) Done() chan bool {
	return scraper.doneChannel
}
