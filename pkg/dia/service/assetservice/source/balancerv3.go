package source

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	vault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv3/vault"
	vaultextension "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv3/vaultextension"
	"go.uber.org/ratelimit"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	balancerV3RateLimitPerSec = 50
	balancerV3FilterPageSize  = 5000
)

type BalancerV3AssetSource struct {
	RestClient                     *ethclient.Client
	assetChannel                   chan dia.Asset
	doneChannel                    chan bool
	waitTime                       int
	blockchain                     string
	exchangeName                   string
	cachedAssets                   sync.Map // map[string]dia.Asset
	exchangeFactoryContractAddress string
	startblockPoolRegister         uint64
	rl                             ratelimit.Limiter
}

func NewBalancerV3AssetSource(exchange dia.Exchange) (bas *BalancerV3AssetSource) {

	bas = makeBalancerV3AssetSource(exchange, "", uniswapWaitMilliseconds)
	bas.rl = ratelimit.New(balancerV3RateLimitPerSec)

	switch exchange.Name {
	case dia.BalancerV3Exchange:
		bas.startblockPoolRegister = 21332121
	case dia.BalancerV3ExchangeBase:
		bas.startblockPoolRegister = 25343854
	case dia.BeetsV3ExchangeSonic:
		bas.startblockPoolRegister = 368135
	}

	// TO DO: implement postgres storage of last fetched block.

	go func() {
		bas.fetchAssets()
	}()
	return bas

}

func (bas *BalancerV3AssetSource) fetchAssets() {

	pools, err := bas.listPools()
	if err != nil {
		log.Fatal("list available pools: ", err)
	}

	bas.getAssetsFromPools(pools)

	bas.doneChannel <- true

}

// makeBalancerV3AssetSource returns an asset source as used in NewBalancerV3AssetSource.
func makeBalancerV3AssetSource(exchange dia.Exchange, restDial string, waitMilliseconds string) *BalancerV3AssetSource {
	var (
		restClient   *ethclient.Client
		err          error
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		bas          *BalancerV3AssetSource
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

	bas = &BalancerV3AssetSource{
		RestClient:                     restClient,
		assetChannel:                   assetChannel,
		doneChannel:                    doneChannel,
		blockchain:                     exchange.BlockChain.Name,
		waitTime:                       waitTime,
		exchangeName:                   exchange.Name,
		exchangeFactoryContractAddress: exchange.Contract,
	}
	return bas
}

// allRegisteredPools returns an iterator for all pools registered.
func (bas *BalancerV3AssetSource) allRegisteredPools() ([]*vault.VaultPoolRegistered, error) {
	var (
		offset     uint64 = balancerV3FilterPageSize
		startBlock uint64 = bas.startblockPoolRegister
		endBlock          = startBlock + offset
		events     []*vault.VaultPoolRegistered
	)

	filterer, err := vault.NewVaultFilterer(common.HexToAddress(bas.exchangeFactoryContractAddress), bas.RestClient)
	if err != nil {
		return nil, err
	}

	currBlock, err := bas.RestClient.BlockNumber(context.Background())
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

// listPools returns a list of pools given by the addresses of the underlying assets.
func (bas *BalancerV3AssetSource) listPools() ([][]common.Address, error) {

	events, err := bas.allRegisteredPools()
	if err != nil {
		return nil, err
	}

	vaultExtensionCaller, err := vaultextension.NewVaultextensionCaller(common.HexToAddress(bas.exchangeFactoryContractAddress), bas.RestClient)
	if err != nil {
		return [][]common.Address{}, err
	}

	pools := make([][]common.Address, len(events))
	for idx, evt := range events {

		poolAssets, err := vaultExtensionCaller.GetPoolTokens(&bind.CallOpts{}, evt.Pool)
		if err != nil {
			log.Error("GetPoolTokens: ", err)
		}
		log.Infof("fetch pool %s", evt.Pool.Hex())

		pools[idx] = poolAssets
	}

	return pools, nil
}

// getAssetsFromPools fetches all assets from @pools and sends them into the asset channel.
func (bas *BalancerV3AssetSource) getAssetsFromPools(pools [][]common.Address) {

	checkMap := make(map[string]struct{})

	for _, tokens := range pools {
		for i := 0; i < len(tokens); i++ {
			if _, ok := checkMap[tokens[i].Hex()]; ok {
				continue
			} else {
				checkMap[tokens[i].Hex()] = struct{}{}
			}
			asset, err := bas.assetFromToken(tokens[i])
			if err != nil {
				log.Error("get asset from token: ", err)
			}
			bas.assetChannel <- asset
		}
	}

}

func (bas *BalancerV3AssetSource) assetFromToken(token common.Address) (dia.Asset, error) {
	cached, ok := bas.cachedAssets.Load(token.Hex())
	if !ok {
		asset, err := ethhelper.ETHAddressToAsset(token, bas.RestClient, bas.blockchain)
		if err != nil {
			return dia.Asset{}, err
		}
		bas.cachedAssets.Store(token.Hex(), asset)
		return asset, nil
	}

	asset := cached.(dia.Asset)

	return asset, nil
}

func (bas *BalancerV3AssetSource) Asset() chan dia.Asset {
	return bas.assetChannel
}

func (bas *BalancerV3AssetSource) Done() chan bool {
	return bas.doneChannel
}
