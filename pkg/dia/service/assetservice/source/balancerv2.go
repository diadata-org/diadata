package source

import (
	"context"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	balancervault "github.com/diadata-org/diadata/pkg/dia/scraper/exchange-scrapers/balancerv2/vault"
	"go.uber.org/ratelimit"
	"golang.org/x/sync/errgroup"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	balancerV2RateLimitPerSec = 50
	balancerV2FilterPageSize  = 50000
	balancerV2RestDial        = "https://rpc.ftm.tools/"
)

var (
	balancerV2StartBlockPoolRegister = 16896080
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
	rl                             ratelimit.Limiter
}

func NewBalancerV2AssetSource(exchange dia.Exchange) (bas *BalancerV2AssetSource) {

	switch exchange.Name {
	case dia.BeetsExchange:
		bas = makeBalancerV2AssetSource(exchange, restDial, uniswapWaitMilliseconds)
		balancerV2StartBlockPoolRegister = 16896080
		rest, err := ethclient.Dial(utils.Getenv("ETH_URI_REST", balancerV2RestDial))
		if err != nil {
			log.Error("dial rest client: ", err)
			return nil
		}
		bas.RestClient = rest
	}

	bas.exchangeFactoryContractAddress = exchange.Contract.Hex()
	bas.rl = ratelimit.New(balancerV2RateLimitPerSec)
	bas.blockchain = exchange.BlockChain.Name
	bas.exchangeName = exchange.Name

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
	pairs, err := bas.listPairs(pools)
	if err != nil {
		log.Fatal("list available pairs: ", err)
	}
	log.Info("Found ", len(pairs), " pairs")

	checkMap := make(map[string]struct{})
	for _, pair := range pairs {
		time.Sleep(time.Duration(bas.waitTime) * time.Millisecond)

		asset0 := pair.QuoteToken
		asset1 := pair.BaseToken
		asset0.Blockchain = bas.blockchain
		asset1.Blockchain = bas.blockchain
		// Don't repeat sending already sent assets
		if _, ok := checkMap[asset0.Address]; !ok {
			if asset0.Symbol != "" {
				checkMap[asset0.Address] = struct{}{}
				bas.assetChannel <- asset0
			}
		}
		if _, ok := checkMap[asset1.Address]; !ok {
			if asset1.Symbol != "" {
				checkMap[asset1.Address] = struct{}{}
				bas.assetChannel <- asset1
			}
		}
	}
	bas.doneChannel <- true
}

// makeBalancerV2AssetSource returns an asset source as used in NewBalancerV2AssetSource.
func makeBalancerV2AssetSource(exchange dia.Exchange, restDial string, waitMilliseconds string) *BalancerV2AssetSource {
	var restClient *ethclient.Client
	var err error
	var assetChannel = make(chan dia.Asset)
	var doneChannel = make(chan bool)
	var bas *BalancerV2AssetSource
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
		RestClient:   restClient,
		assetChannel: assetChannel,
		doneChannel:  doneChannel,
		blockchain:   exchange.BlockChain.Name,
		waitTime:     waitTime,
	}
	return bas
}

func (bas *BalancerV2AssetSource) Asset() chan dia.Asset {
	return bas.assetChannel
}

func (bas *BalancerV2AssetSource) Done() chan bool {
	return bas.doneChannel
}

func (bas *BalancerV2AssetSource) allRegisteredPools() ([]*balancervault.BalancerVaultPoolRegistered, error) {
	filterer, err := balancervault.NewBalancerVaultFilterer(common.HexToAddress(bas.exchangeFactoryContractAddress), bas.RestClient)
	if err != nil {
		return nil, err
	}

	currBlock, err := bas.RestClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	var offset uint64 = balancerV2FilterPageSize
	var startBlock uint64 = uint64(balancerV2StartBlockPoolRegister)
	var endBlock = startBlock + offset
	var events []*balancervault.BalancerVaultPoolRegistered
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

func (bas *BalancerV2AssetSource) listPools() ([][]common.Address, error) {
	events, err := bas.allRegisteredPools()
	if err != nil {
		return nil, err
	}

	caller, err := balancervault.NewBalancerVaultCaller(common.HexToAddress(bas.exchangeFactoryContractAddress), bas.RestClient)
	if err != nil {
		return nil, err
	}

	var g errgroup.Group
	var mu sync.Mutex
	pools := make([][]common.Address, len(events))
	for idx, evt := range events {
		idx := idx
		evt := evt
		g.Go(func() error {
			bas.rl.Take()
			pool, err := caller.GetPoolTokens(&bind.CallOpts{}, evt.PoolId)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()

			pools[idx] = pool.Tokens

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return pools, nil
}

func (bas *BalancerV2AssetSource) listPairs(pools [][]common.Address) (pairs []dia.Pair, err error) {
	pairCount := 0
	pairMap := make(map[int]dia.Pair)
	var g errgroup.Group
	var mu sync.Mutex
	for _, tokens := range pools {
		if len(tokens) < 2 {
			continue
		}

		for i := 0; i < len(tokens); i++ {
			for j := i + 1; j < len(tokens); j++ {
				pairCount++
				i := i
				j := j
				pairCount := pairCount
				tokens := tokens
				g.Go(func() error {
					bas.rl.Take()
					pair, err := bas.makePair(tokens[i], tokens[j])
					if err != nil {
						log.Warn(err)

						return nil
					}

					mu.Lock()
					defer mu.Unlock()

					pairMap[pairCount] = pair

					return nil
				})
			}
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	keys := make([]int, 0, len(pairMap))
	for k := range pairMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		pairs = append(pairs, pairMap[k])
	}

	return
}

func (bas *BalancerV2AssetSource) makePair(token0, token1 common.Address) (dia.Pair, error) {
	asset0, err := bas.assetFromToken(token0)
	if err != nil {
		return dia.Pair{}, err
	}
	asset1, err := bas.assetFromToken(token1)
	if err != nil {
		return dia.Pair{}, err
	}

	var pair dia.Pair
	pair.QuoteToken = asset0
	pair.BaseToken = asset1
	return pair, nil
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
