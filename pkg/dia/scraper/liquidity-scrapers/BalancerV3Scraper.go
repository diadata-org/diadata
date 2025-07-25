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
	vault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv3/vault"
	vaultextension "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv3/vaultextension"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	balancerV3FilterPageSize = 5000
	balancerV3RestDial       = ""
)

type BalancerV3Scraper struct {
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

// NewBalancerV3Scraper returns a Balancer V3 scraper
func NewBalancerV3Scraper(exchange dia.Exchange, relDB *models.RelDB, datastore *models.DB) *BalancerV3Scraper {
	var (
		restClient  *ethclient.Client
		err         error
		poolChannel = make(chan dia.Pool)
		doneChannel = make(chan bool)
		scraper     *BalancerV3Scraper
	)

	log.Infof("Init rest client for %s.", exchange.BlockChain.Name)
	restClient, err = ethclient.Dial(utils.Getenv(strings.ToUpper(exchange.BlockChain.Name)+"_URI_REST", balancerV3RestDial))
	if err != nil {
		log.Fatal("init rest client: ", err)
	}

	scraper = &BalancerV3Scraper{
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
	case dia.BalancerV3Exchange:
		scraper.startblockPoolRegister = 21332121

	}

	go func() {
		scraper.fetchPools()
	}()

	return scraper
}

// fetchPools collects all available pools and sends them into the pool channel.
func (scraper *BalancerV3Scraper) fetchPools() {
	events, err := scraper.allRegisteredPools()
	if err != nil {
		log.Fatal("fetch all registered pools: ", err)
	}

	vaultExtensionCaller, err := vaultextension.NewVaultextensionCaller(common.HexToAddress(scraper.vaultContract), scraper.RestClient)
	if err != nil {
		log.Fatal("vault extension caller: ", err)
	}

	for _, evt := range events {
		poolTokens, err := vaultExtensionCaller.GetPoolData(&bind.CallOpts{}, evt.Pool)
		if err != nil {
			log.Warn("get pool tokens: ", err)
		}
		assetvolumes := scraper.extractPoolInfo(poolTokens)
		pool := dia.Pool{
			Exchange:     dia.Exchange{Name: scraper.exchangeName},
			Blockchain:   dia.BlockChain{Name: scraper.blockchain},
			Address:      evt.Pool.String(),
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
func (scraper *BalancerV3Scraper) allRegisteredPools() ([]*vault.VaultPoolRegistered, error) {
	var (
		offset     uint64 = balancerV3FilterPageSize
		startBlock uint64 = scraper.startblockPoolRegister
		endBlock          = startBlock + offset
		events     []*vault.VaultPoolRegistered
	)

	filterer, err := vault.NewVaultFilterer(common.HexToAddress(scraper.vaultContract), scraper.RestClient)
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
func (scraper *BalancerV3Scraper) extractPoolInfo(poolData vaultextension.PoolData) (assetvolumes []dia.AssetVolume) {
	for i := range poolData.Tokens {

		asset, err := scraper.relDB.GetAsset(poolData.Tokens[i].Hex(), scraper.blockchain)
		if err != nil {
			asset, err = ethhelper.ETHAddressToAsset(poolData.Tokens[i], scraper.RestClient, scraper.blockchain)
			if err != nil {
				log.Warn("cannot fetch asset from address ", poolData.Tokens[i].Hex())
				continue
			}
		}

		volume, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(poolData.BalancesRaw[i]), new(big.Float).SetFloat64(math.Pow10(int(asset.Decimals)))).Float64()

		assetvolumes = append(assetvolumes, dia.AssetVolume{Asset: asset, Volume: volume, Index: uint8(i)})
	}
	return
}

func (scraper *BalancerV3Scraper) Pool() chan dia.Pool {
	return scraper.poolChannel
}

func (scraper *BalancerV3Scraper) Done() chan bool {
	return scraper.doneChannel
}
