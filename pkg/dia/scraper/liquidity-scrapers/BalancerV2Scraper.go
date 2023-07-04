package liquidityscrapers

import (
	"context"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	balancervault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv2/vault"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	balancerV2FilterPageSize = 5000
	balancerV2RestDial       = ""
)

type BalancerV2Scraper struct {
	RestClient             *ethclient.Client
	relDB                  *models.RelDB
	datastore              *models.DB
	poolChannel            chan dia.Pool
	doneChannel            chan bool
	blockchain             string
	exchangeName           string
	vaultContract          string
	startblockPoolRegister uint64
	cachedAssets           map[string]dia.Asset
}

// NewBalancerV2Scraper returns a Balancer V2 scraper
func NewBalancerV2Scraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *BalancerV2Scraper {
	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *BalancerV2Scraper
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", balancerV2RestDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	scraper = &BalancerV2Scraper{
		RestClient:    restClient,
		relDB:         relDB,
		datastore:     datastore,
		poolChannel:   poolChannel,
		doneChannel:   doneChannel,
		blockchain:    exchange.BlockChain.Name,
		exchangeName:  exchange.Name,
		vaultContract: exchange.Contract,
		cachedAssets:  make(map[string]dia.Asset),
	}

	switch exchange.Name {
	case dia.BalancerV2Exchange:
		scraper.startblockPoolRegister = 12272146
	case dia.BalancerV2ExchangeArbitrum:
		scraper.startblockPoolRegister = 222832
	case dia.BalancerV2ExchangePolygon:
		scraper.startblockPoolRegister = 15832990
	case dia.BeetsExchange:
		scraper.startblockPoolRegister = 16896080
	}

	go func() {
		scraper.fetchPools()
	}()

	return scraper
}

// fetchPools collects all available pools and sends them into the pool channel.
func (scraper *BalancerV2Scraper) fetchPools() {
	events, err := scraper.allRegisteredPools()
	if err != nil {
		log.Fatal("fetch all registered pools: ", err)
	}

	caller, err := balancervault.NewBalancerVaultCaller(common.HexToAddress(scraper.vaultContract), scraper.RestClient)
	if err != nil {
		log.Fatal("make balancer vault caller: ", err)
	}

	for _, evt := range events {
		poolTokens, err := caller.GetPoolTokens(&bind.CallOpts{}, evt.PoolId)
		if err != nil {
			log.Warn("get pool tokens: ", err)
		}
		assetvolumes := scraper.extractPoolInfo(poolTokens)
		pool := dia.Pool{
			Exchange:     dia.Exchange{Name: scraper.exchangeName},
			Blockchain:   dia.BlockChain{Name: scraper.blockchain},
			Address:      evt.PoolAddress.String(),
			Assetvolumes: assetvolumes,
			Time:         time.Now(),
		}

		// Determine USD liquidity.
		if pool.SufficientNativeBalance(GLOBAL_NATIVE_LIQUIDITY_THRESHOLD) {
			scraper.datastore.GetPoolLiquiditiesUSD(&pool, priceCache)
		}

		scraper.poolChannel <- pool
	}
	scraper.doneChannel <- true
}

// allRegisteredPools returns a slice of all pool creation events.
func (scraper *BalancerV2Scraper) allRegisteredPools() ([]*balancervault.BalancerVaultPoolRegistered, error) {
	var (
		offset     uint64 = balancerV2FilterPageSize
		startBlock uint64 = scraper.startblockPoolRegister
		endBlock          = startBlock + offset
		events     []*balancervault.BalancerVaultPoolRegistered
	)

	filterer, err := balancervault.NewBalancerVaultFilterer(common.HexToAddress(scraper.vaultContract), scraper.RestClient)
	if err != nil {
		return nil, err
	}

	currBlock, err := scraper.RestClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	for {
		if endBlock > currBlock {
			endBlock = currBlock
		}
		log.Infof("startblock - endblock: %v --- %v ", startBlock, endBlock)

		it, err := filterer.FilterPoolRegistered(&bind.FilterOpts{
			Start: startBlock,
			End:   &endBlock,
		}, nil, nil)
		if err != nil {
			log.Warn("filterpoolregistered: ", err)
			continue
		}

		for it.Next() {
			events = append(events, it.Event)
			log.Info("pool address: ", it.Event.PoolAddress)
		}
		if err := it.Close(); err != nil {
			log.Warn("closing iterator: ", it)
		}

		if endBlock == currBlock {
			break
		}

		startBlock = endBlock + 1
		endBlock = endBlock + offset
	}

	return events, nil
}

// extractPoolInfo returns assetvolumes in the correct format for a dia.Pool.
func (scraper *BalancerV2Scraper) extractPoolInfo(poolTokens struct {
	Tokens          []common.Address
	Balances        []*big.Int
	LastChangeBlock *big.Int
}) (assetvolumes []dia.AssetVolume) {
	for i := range poolTokens.Tokens {

		asset, err := scraper.relDB.GetAsset(poolTokens.Tokens[i].Hex(), scraper.blockchain)
		if err != nil {
			asset, err = ethhelper.ETHAddressToAsset(poolTokens.Tokens[i], scraper.RestClient, scraper.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolTokens.Tokens[i].Hex())
				continue
			}
		}

		volume, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(poolTokens.Balances[i]), new(big.Float).SetFloat64(math.Pow10(int(asset.Decimals)))).Float64()

		assetvolumes = append(assetvolumes, dia.AssetVolume{Asset: asset, Volume: volume, Index: uint8(i)})
	}
	return
}

// assetFromToken returns the dia.Asset corresponding to an address on the underlying scraper's blockchain.
func (scraper *BalancerV2Scraper) assetFromToken(token common.Address) (dia.Asset, error) {
	cached, ok := scraper.cachedAssets[token.Hex()]
	if !ok {
		asset, err := ethhelper.ETHAddressToAsset(token, scraper.RestClient, scraper.blockchain)
		if err != nil {
			return dia.Asset{}, err
		}
		scraper.cachedAssets[token.Hex()] = asset
		return asset, nil
	} else {
		return cached, nil
	}
}

func (scraper *BalancerV2Scraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *BalancerV2Scraper) Done() chan bool {
	return scraper.doneChannel
}
