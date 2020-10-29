package pool

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

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

func (bp *BalancerPool) mainLoop() {
	// currently:
	// - watches for new pools
	// - downloads all pools and gets informations

	newPoolEventChan := make(chan *balfactorycontract.BALFactoryContractLOGNEWPOOL, 16)
	newPoolChan := make(chan common.Address, 16)

	go bp.watchFactory(newPoolEventChan)
	go bp.ListAllPools(newPoolEventChan)

	for i := 0; i < balGoroutinesAmount; i++ {
		go bp.poolHandler(newPoolChan)
	}

	for {
		select {
		case event := <-newPoolEventChan:
			log.Info("got a New Pool event")
			newPoolChan <- event.Pool
		}
	}
}

func (bp *BalancerPool) getPool(poolAddress common.Address) (err error) {
	pool, err := balancerpoolcontract.NewBalancerPoolContractCaller(poolAddress, bp.restClient)
	if err != nil {
		return
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
	balances := []float64{}

	for _, token := range tokens {
		tokCaller, err := baltokencontract.NewBALTokenContractCaller(token, bp.restClient)
		if err != nil {
			return err
		}

		symbol, err := tokCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			return err
		}

		balanceInt, err := pool.GetBalance(&bind.CallOpts{}, token)
		if err != nil {
			return err
		}

		balance := float64(balanceInt.Int64()) / 10e17

		symbols = append(symbols, symbol)
		balances = append(balances, balance)
	}

	swapFee, err := pool.GetSwapFee(&bind.CallOpts{})
	if err != nil {
		return
	}

	swapFeeFloat := float64(swapFee.Int64()) / float64(10e15)

	// TODO: merge token balances into a single balance
	pr := models.FarmingPool{
		// Balance      float64

		TimeStamp:    time.Now(),
		PoolID:       poolAddress.String(),
		ProtocolName: bp.scraper.poolName,

		OutputAsset: symbols,
		InputAsset:  symbols,

		Rate: swapFeeFloat,

		BlockNumber: header.Number.Int64(),
	}

	bp.scraper.chanPoolInfo <- &pr
	return nil
}
