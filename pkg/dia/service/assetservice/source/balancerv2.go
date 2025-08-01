package source

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	balancervault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv2/vault"
	"go.uber.org/ratelimit"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	balancerV2RateLimitPerSec = 50
	balancerV2FilterPageSize  = 5000
)

type BalancerV2AssetSource struct {
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

func NewBalancerV2AssetSource(exchange dia.Exchange) (bas *BalancerV2AssetSource) {

	bas = makeBalancerV2AssetSource(exchange, "", uniswapWaitMilliseconds)
	bas.rl = ratelimit.New(balancerV2RateLimitPerSec)

	switch exchange.Name {
	case dia.BalancerV2Exchange:
		bas.startblockPoolRegister = 12272146
	case dia.BalancerV2ExchangeArbitrum:
		bas.startblockPoolRegister = 222832
	case dia.BeetsExchange:
		bas.startblockPoolRegister = 16896080
	case dia.BeetsV2ExchangeSonic:
		bas.startblockPoolRegister = 368312
	case dia.BalancerV2ExchangePolygon:
		bas.startblockPoolRegister = 15832990
	}

	go func() {
		bas.fetchAssets()
	}()
	return bas

}

func (bas *BalancerV2AssetSource) fetchAssets() {

	pools, err := bas.listPools()
	if err != nil {
		log.Fatal("list available pools: ", err)
	}

	bas.getAssetsFromPools(pools)

	bas.doneChannel <- true

}

// makeBalancerV2AssetSource returns an asset source as used in NewBalancerV2AssetSource.
func makeBalancerV2AssetSource(exchange dia.Exchange, restDial string, waitMilliseconds string) *BalancerV2AssetSource {
	var (
		restClient   *ethclient.Client
		err          error
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
		bas          *BalancerV2AssetSource
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

	bas = &BalancerV2AssetSource{
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
func (bas *BalancerV2AssetSource) allRegisteredPools() ([]*balancervault.BalancerVaultPoolRegistered, error) {
	var (
		offset     uint64 = balancerV2FilterPageSize
		startBlock uint64 = bas.startblockPoolRegister
		endBlock          = startBlock + offset
		events     []*balancervault.BalancerVaultPoolRegistered
	)

	filterer, err := balancervault.NewBalancerVaultFilterer(common.HexToAddress(bas.exchangeFactoryContractAddress), bas.RestClient)
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
func (bas *BalancerV2AssetSource) listPools() ([][]common.Address, error) {

	events, err := bas.allRegisteredPools()
	if err != nil {
		return nil, err
	}

	caller, err := balancervault.NewBalancerVaultCaller(common.HexToAddress(bas.exchangeFactoryContractAddress), bas.RestClient)
	if err != nil {
		return nil, err
	}

	pools := make([][]common.Address, len(events))
	for idx, evt := range events {
		pool, err := caller.GetPoolTokens(&bind.CallOpts{}, evt.PoolId)
		if err != nil {
			log.Errorf("fetch pool %s: %v", evt.PoolAddress.Hex(), err)
		}
		pools[idx] = pool.Tokens
	}

	return pools, nil
}

// getAssetsFromPools fetches all assets from @pools and sends them into the asset channel.
func (bas *BalancerV2AssetSource) getAssetsFromPools(pools [][]common.Address) {

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

func (bas *BalancerV2AssetSource) assetFromToken(token common.Address) (dia.Asset, error) {
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

func (bas *BalancerV2AssetSource) Asset() chan dia.Asset {
	return bas.assetChannel
}

func (bas *BalancerV2AssetSource) Done() chan bool {
	return bas.doneChannel
}
