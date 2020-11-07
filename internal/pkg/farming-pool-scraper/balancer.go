package pool

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"errors"

	"github.com/bep/debounce"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	balfactorycontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/balancer/factory"
	balancerpoolcontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/balancer/pool"
	baltokencontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/balancer/token"
	models "github.com/diadata-org/diadata/pkg/model"
)

var (
	// Balancer: Exchange Proxy 2
	// https://etherscan.io/address/0x3e66b66fd1d0b02fda6c811da9e0547970db2f21
	balancerExchangeAddress = "0x3E66B66Fd1d0b02fDa6C811Da9E0547970DB2f21"

	// Balancer: BAL Token
	// https://etherscan.io/address/0xba100000625a3754423978a60c9317c58a424e3d
	balancerTokenAddress = "0xba100000625a3754423978a60c9317c58a424e3D"

	// Balancer: BFactory
	// https://etherscan.io/address/0x9424b1412450d0f8fc2255faf6046b98213b76bd
	balFactoryAddress = common.HexToAddress("0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd")
	firstFactoryBlock = uint64(9562480)
)

var (
	newPoolTopic = "0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945"
)

var (
	balGoroutinesAmount = 4
)

type BalancerPool struct {
	scraper *PoolScraper

	restClient *ethclient.Client
	wsClient   *ethclient.Client
}

func NewBalancerPoolScrapper(scraper *PoolScraper) *BalancerPool {
	// create rest and ws eth clients
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}

	bp := &BalancerPool{
		restClient: restClient,
		wsClient:   wsClient,

		scraper: scraper,
	}

	go bp.mainLoop()

	return bp
}

func (bp *BalancerPool) watchFactory(newPoolChan chan *balfactorycontract.BALFactoryContractLOGNEWPOOL) {
	fr, _ := balfactorycontract.NewBALFactoryContractFilterer(balFactoryAddress, bp.wsClient)

	_, err := fr.WatchLOGNEWPOOL(&bind.WatchOpts{}, newPoolChan, nil, nil)
	if err != nil {
		log.Errorln("Error subscribing to New Pool events: ", err)
	} else {
		log.Info("Subscribed to New Pool events")
	}
}

func (bp *BalancerPool) ListAllPools(newPoolEventChan chan *balfactorycontract.BALFactoryContractLOGNEWPOOL) {
	fr, _ := balfactorycontract.NewBALFactoryContractFilterer(balFactoryAddress, bp.wsClient)

	it, err := fr.FilterLOGNEWPOOL(&bind.FilterOpts{
		Start:   firstFactoryBlock,
		End:     nil,
		Context: context.Background(),
	}, nil, nil)
	if err != nil {
		log.Errorln("creating New Pool iterator")
		return
	}
	log.Info("created historic new pool iterator")

	for it.Next() {
		ev := it.Event
		log.Info("new pool from historic new pool iterator")
		newPoolEventChan <- ev
	}

	if err := it.Error(); err != nil {
		log.Errorln("historic New Pool event iterator error: ", err)
		return
	}

	log.Info("finished iterating over historic pools")
}

func (bp *BalancerPool) poolHandler(poolChan chan common.Address) {
	for {
		pool := <-poolChan
		if err := bp.getPool(pool); err != nil {
			log.Errorln("error fetching balancer pool informations: ", err)
		}
	}
}

func (bp *BalancerPool) poolWatcher(perceive chan common.Address) {
	watchedPools := []common.Address{}

	for {
		pool := <-perceive

		// skip already watched pools
		found := false
		for _, watchedPool := range watchedPools {
			if watchedPool.Hex() == pool.Hex() {
				found = true
				break
			}
		}
		if found {
			continue
		}

		watchedPools = append(watchedPools, pool)

		// watch unknown pools for transaction events
		fr, err := balancerpoolcontract.NewBalancerPoolContractFilterer(pool, bp.wsClient)
		if err != nil {
			log.WithField("PoolAddress", pool).Debug("error creating an event filterer for pool")
			continue
		}

		sink := make(chan *balancerpoolcontract.BalancerPoolContractTransfer)
		_, err = fr.WatchTransfer(&bind.WatchOpts{}, sink, nil, nil)
		if err != nil {
			log.WithField("PoolAddress", pool).Error("error watching for transfers for pool")
			continue
		}

		go func() {
			fetch := func() {
				if err := bp.getPool(pool); err != nil {
					log.WithField("PoolAddress", pool).Error("error getting pool informations after transfer event")
				}
			}

			// debounce the function in order to avoid rate limit errors
			debounced := debounce.New(1 * time.Minute)

			for {
				tx := <-sink
				log.WithFields(logrus.Fields{
					"PoolAddress": pool,
					"Amount":      tx.Amt,
				}).Debug("new transaction on pool")

				debounced(fetch)
			}
		}()
	}
}

func (bp *BalancerPool) mainLoop() {
	// currently:
	// - watches for new pools
	// - downloads all pools and gets informations
	// todo:
	// - watching for transactions as it updates the pool rate

	newPoolEventChan := make(chan *balfactorycontract.BALFactoryContractLOGNEWPOOL, 16)
	newPoolChan := make(chan common.Address, 16)
	poolWatcherPerceiver := make(chan common.Address, 1)

	// go bp.watchFactory(newPoolEventChan)
	// go bp.ListAllPools(newPoolEventChan)
	go bp.poolWatcher(poolWatcherPerceiver)

	// https://pools.balancer.exchange/#/pool/0x59a19d8c652fa0284f44113d0ff9aba70bd46fb4/
	newPoolChan <- common.HexToAddress("0x59a19d8c652fa0284f44113d0ff9aba70bd46fb4")
	// newPoolChan <- common.HexToAddress("0xd8e9690eff99e21a2de25e0b148ffaf47f47c972")
	// newPoolChan <- common.HexToAddress("0x9866772a9bdb4dc9d2c5a4753e8658b8b0ca1fc3")

	for i := 0; i < balGoroutinesAmount; i++ {
		go bp.poolHandler(newPoolChan)
	}

	for {
		select {
		case event := <-newPoolEventChan:
			log.Info("got a New Pool event")

			newPoolChan <- event.Pool
			poolWatcherPerceiver <- event.Pool
		}
	}
}

// https://steemit.com/tutorial/@gopher23/power-and-root-functions-using-big-float-in-golang
func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}

func Pow(a *big.Float, e uint64) *big.Float {
	result := Zero().Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = Mul(result, a)
	}
	return result
}

func (bp *BalancerPool) getPool(poolAddress common.Address) (err error) {
	pool, err := balancerpoolcontract.NewBalancerPoolContractCaller(poolAddress, bp.restClient)
	if err != nil {
		return
	}

	// ignore unfinalized pools
	finalized, err := pool.IsFinalized(&bind.CallOpts{})
	if err != nil {
		return
	}

	if !finalized {
		return errors.New("pool is not finalized")
	}

	// note(gravi): this represents the latest known block
	header, err := bp.restClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return
	}

	tokens, err := pool.GetCurrentTokens(&bind.CallOpts{})
	if err != nil {
		return
	}

	symbols := []string{}

	// V = (b1^w1)(b2^w2)(b3^w3)
	// V is the product of each token of the pool.
	// b1/b2/b3 are the balances of the token
	// w1/w2/w3 are their respective weights
	V := big.NewFloat(1)

	for _, token := range tokens {
		tokCaller, err := baltokencontract.NewBALTokenContractCaller(token, bp.restClient)
		if err != nil {
			return err
		}

		symbol, err := tokCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			return err
		}

		// get token weight in pool
		weight, err := pool.GetNormalizedWeight(&bind.CallOpts{}, token)
		if err != nil {
			return err
		}
		weightFloat := new(big.Float).SetInt(weight)
		weightFloat = weightFloat.Quo(weightFloat, big.NewFloat(10e17))

		// get token balance in pool
		balance, err := pool.GetBalance(&bind.CallOpts{}, token)
		if err != nil {
			return err
		}
		balanceFloat := new(big.Float).SetInt(balance)
		balanceFloat = balanceFloat.Quo(balanceFloat, big.NewFloat(10e17))

		fmt.Println(symbol, "balance:", balanceFloat)

		balanceFloatFloat, _ := balanceFloat.Float64()
		weightFloatFloat, _ := weightFloat.Float64()

		fmt.Println(symbol, "weight:", weightFloatFloat)

		thing := math.Pow(balanceFloatFloat, weightFloatFloat)
		V.Mul(V, new(big.Float).SetFloat64(thing))

		symbols = append(symbols, symbol)
	}

	balance, err := pool.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return
	}
	balanceFloat := new(big.Float).SetInt(balance)
	balanceFloatFloat, _ := balanceFloat.Float64()
	balanceFloatFloat /= 10e17

	VFloat, _ := V.Float64()

	pr := models.FarmingPool{
		// Balance is the total supply of pool token.
		Balance: balanceFloatFloat,

		TimeStamp:    time.Now(),
		PoolID:       poolAddress.String(),
		ProtocolName: bp.scraper.poolName,

		OutputAsset: symbols,
		InputAsset:  symbols,

		Rate: VFloat / balanceFloatFloat, // v = V / num_bpt

		BlockNumber: header.Number.Int64(),
	}

	bp.scraper.chanPoolInfo <- &pr
	return nil
}
