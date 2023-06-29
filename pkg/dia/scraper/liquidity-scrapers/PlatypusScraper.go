package liquidityscrapers

import (
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance"
	platypusAssetABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/asset"
	platypusPoolABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/pool"
	platypusTokenABI "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/platypusfinance/token"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	platypusRestDialEth     = "https://api.avax.network/ext/bc/C/rpc"
	platypusMasterRegV3Addr = "0x7125B4211357d7C3a90F796c956c12c681146EbB"
	platypusMasterRegV2Addr = "0x68c5f4374228BEEdFa078e77b5ed93C28a2f713E"
	platypusMasterRegV1Addr = "0xB0523f9F473812FB195Ee49BC7d2ab9873a98044"
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

// The scraper object for Platypus Finance
type PlatypusScraper struct {
	RestClient   *ethclient.Client
	datastore    *models.DB
	poolChannel  chan dia.Pool
	doneChannel  chan bool
	blockchain   string
	exchangeName string

	platypusCoins map[string]*PlatypusCoin

	basePoolRegistry platypusRegistry
}

// Returns a new exchange scraper
func NewPlatypusScraper(exchange dia.Exchange, datastore *models.DB) *PlatypusScraper {

	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *PlatypusScraper
	)
	registries := []platypusRegistry{
		{Version: 3, Address: common.HexToAddress(platypusMasterRegV3Addr)},
		{Version: 2, Address: common.HexToAddress(platypusMasterRegV2Addr)},
		{Version: 1, Address: common.HexToAddress(platypusMasterRegV1Addr)},
	}

	log.Infof("init rest and ws client for %s", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", platypusRestDialEth))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	scraper = &PlatypusScraper{
		blockchain:   exchange.BlockChain.Name,
		exchangeName: exchange.Name,
		RestClient:   restClient,
		datastore:    datastore,
		poolChannel:  poolChannel,
		doneChannel:  doneChannel,

		platypusCoins: make(map[string]*PlatypusCoin),
	}

	// // Load metadata from master registries
	// for _, registry := range registries {
	// 	err := scraper.loadPoolsAndCoins(registry)
	// 	if err != nil {
	// 		log.Errorf("loadPoolsAndCoins error w %s registry (v%d): %s", registry.Address.Hex(), registry.Version, err)
	// 	}
	// 	log.Infof("metadata loaded, now scraper have %d pools data and %d coins", len(scraper.pools.pools), len(scraper.platypusCoins))
	// }

	scraper.basePoolRegistry = platypusRegistry{Version: 3, Address: common.HexToAddress(platypusMasterRegV3Addr)}

	go func() {
		for _, registry := range registries {
			scraper.fetchPools(registry)
		}
		scraper.doneChannel <- true
	}()

	return scraper
}

// Load pools and coins metadata from master registry
func (scraper *PlatypusScraper) fetchPools(registry platypusRegistry) (err error) {
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

		pool := scraper.loadPoolData(asset.LpToken.Hex())
		scraper.Pool() <- pool
	}

	return err
}

func (scraper *PlatypusScraper) loadPoolData(poolAddress string) (pool dia.Pool) {
	var err error

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
			balance     *big.Int
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

			balance, err = contractToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
			if err != nil {
				log.Error("get balance of: ", err)
			}

		}

		poolAsset := dia.Asset{
			Address:    c.Hex(),
			Blockchain: scraper.blockchain,
			Decimals:   decimals,
			Symbol:     symbol,
			Name:       name,
		}
		log.Info("pool asset: ", poolAsset)

		volume, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance), new(big.Float).SetFloat64(math.Pow10(int(decimals)))).Float64()
		pool.Assetvolumes = append(pool.Assetvolumes, dia.AssetVolume{
			Asset:  poolAsset,
			Volume: volume,
		})

	}

	// Determine USD liquidity.
	if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
		scraper.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
	}

	pool.Exchange = dia.Exchange{Name: scraper.exchangeName}
	pool.Blockchain = dia.BlockChain{Name: scraper.blockchain}
	pool.Address = poolAddress
	pool.Time = time.Now()

	return pool
}

func (scraper *PlatypusScraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *PlatypusScraper) Done() chan bool {
	return scraper.doneChannel
}
