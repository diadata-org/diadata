package pool

import (
	"context"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/bep/debounce"
	"github.com/pkg/errors"
	logrus "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	balfactorycontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/balancer/factory"
	balancerpoolcontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/balancer/pool"
	baltokencontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/balancer/token"

	models "github.com/diadata-org/diadata/pkg/model"
)

var (
	// Balancer: BFactory
	// https://etherscan.io/address/0x9424b1412450d0f8fc2255faf6046b98213b76bd
	balFactoryAddress = common.HexToAddress("0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd")
	firstFactoryBlock = uint64(9562480)
)

var (
	// controls the number of pool handling goroutines to use.
	// more means faster pool information retrieval, but it may cause rate limitation errors.
	balPoolInfoFetchingGoroutines = 4
)

type BalancerPoolScrapper struct {
	scraper *PoolScraper

	restClient *ethclient.Client
	wsClient   *ethclient.Client
}

// NewBalancerPoolScrapper creates a new balancer scrapper.
func NewBalancerPoolScrapper(scraper *PoolScraper) *BalancerPoolScrapper {
	// create rest and ws eth clients
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}

	bp := &BalancerPoolScrapper{
		restClient: restClient,
		wsClient:   wsClient,

		scraper: scraper,
	}

	go bp.mainLoop()

	return bp
}

// watchFactory watches the BPFactory contract for New Pool creation events.
func (bp *BalancerPoolScrapper) watchFactory(newPoolChan chan *balfactorycontract.BALFactoryContractLOGNEWPOOL) {
	fr, _ := balfactorycontract.NewBALFactoryContractFilterer(balFactoryAddress, bp.wsClient)

	_, err := fr.WatchLOGNEWPOOL(&bind.WatchOpts{}, newPoolChan, nil, nil)
	if err != nil {
		log.Errorln("Error subscribing to New Pool events: ", err)
	} else {
		log.Debug("Subscribed to New Pool events")
	}
}

// fetchPreviousPools fetches all the past pool creation event on the BPFactory contract.
func (bp *BalancerPoolScrapper) fetchPreviousPools(newPoolEventChan chan *balfactorycontract.BALFactoryContractLOGNEWPOOL) {
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
	log.Trace("created historic new pool iterator")

	for it.Next() {
		ev := it.Event
		log.Trace("new pool from historic new pool iterator")
		newPoolEventChan <- ev
	}

	if err := it.Error(); err != nil {
		log.Errorln("historic New Pool event iterator error: ", err)
		return
	}

	log.Info("finished iterating over historic pools")
}

// poolFetcher executes bp.getPool on provided pools.
func (bp *BalancerPoolScrapper) poolFetcher(poolChan chan common.Address) {
	for {
		pool := <-poolChan
		if err := bp.getPool(pool); err != nil {
			log.
				WithField("poolAddress", pool.Hex()).
				WithField("error", err).
				Error("error fetching balancer pool informations")
		}
	}
}

// poolWatcher watches the given pools for transactions. Upon transaction, it asks for pool informations refreshal.
// which updates pool informations.
func (bp *BalancerPoolScrapper) poolWatcher(perceive chan common.Address, poolToFetch chan common.Address) {
	watchedPoolsMu := new(sync.Mutex)
	watchedPools := make(map[string]func(f func()))

	// watch new blocks for transaction for which destination is a balancer pool
	go func() {
		// subsribe for block headers
		headersChan := make(chan *types.Header)
		sus, err := bp.wsClient.SubscribeNewHead(context.Background(), headersChan)
		if err != nil {
			log.WithError(err).Error("error subscribing to blockchain head")
			return
		}

		for {
			select {
			case err := <-sus.Err():
				log.WithError(err).Error("blockchain subscription error")
			case header := <-headersChan:
				// on block header, fetch the block
				block, err := bp.restClient.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					log.WithError(err).WithField("block", block.Hash().Hex()).Error("error fetching block by hash")
					continue
				}

				// check if the destination of the transaction is a watchedPool.
				// if so, use it's debouncer and put the pool into the poolToFetch channel
				watchedPoolsMu.Lock()
				for _, tx := range block.Transactions() {
					if tx == nil {
						continue
					} else if tx.To() == nil {
						continue
					}

					debounced, ok := watchedPools[tx.To().Hex()]
					if !ok {
						continue
					}

					to := *tx.To()
					debounced(func() {
						log.Debug("fetching pool info because of transaction")
						poolToFetch <- to
					})

					log.WithFields(logrus.Fields{
						"PoolAddress": tx.To().Hex(),
						"TxHash":      tx.Hash().Hex(),
					}).Debug("new transaction on pool")
				}
				watchedPoolsMu.Unlock()
			}
		}
	}()

	for {
		pool := <-perceive

		// skip already watched pools
		if _, ok := watchedPools[pool.Hex()]; ok {
			log.WithField("poolAddress", pool.Hex()).Trace("pool already watched")
			continue
		}

		// create a debouncer for each pool in order to avoid hitting rate-limitations on pools with high volume
		watchedPoolsMu.Lock()
		watchedPools[pool.Hex()] = debounce.New(30 * time.Second)
		watchedPoolsMu.Unlock()

		log.WithField("poolAddress", pool.Hex()).Debug("watching pool for transactions...")
	}
}

func (bp *BalancerPoolScrapper) mainLoop() {
	// currently:
	// - watches for New Pool creation logs on the Balancer Factory contract
	// - retrieves and processes all the past pool creation events
	// - watches for transactions on every pool as it affects the pool rate

	// pools sent to newPoolChan will be handled by a bp.poolHandler routine and
	// watched for transactions by bp.poolWatcher
	newPoolChan := make(chan common.Address, 16)
	// pool creation events sent to newPoolEventChan will get their pool sent to newPoolChan
	newPoolEventChan := make(chan *balfactorycontract.BALFactoryContractLOGNEWPOOL, 16)
	// pools sent to poolHandlerChan are forwarded to a bp.poolHandler routine
	poolToFetchChan := make(chan common.Address, 16)
	// pools sent to poolWatcherPerceiver will be watched by bp.poolWatcher.
	// the same pool can be sent multiple times to this channel.
	poolWatcherPerceiver := make(chan common.Address, 1)

	go bp.watchFactory(newPoolEventChan)
	go bp.fetchPreviousPools(newPoolEventChan)
	go bp.poolWatcher(poolWatcherPerceiver, poolToFetchChan)
	go bp.bootstrapPools(newPoolChan)

	for i := 0; i < balPoolInfoFetchingGoroutines; i++ {
		go bp.poolFetcher(poolToFetchChan)
	}

	go func() {
		for {
			pool := <-newPoolChan
			log.Trace("got a new pool")
			poolWatcherPerceiver <- pool
			poolToFetchChan <- pool
		}
	}()

	// this loop extracts pools from pool events
	for {
		newPoolChan <- (<-newPoolEventChan).Pool
	}
}

// getPool gets informations of the pool at the given pool address. These informations are emited as scrapped data.
// To avoid rate limitation errors, do not call this function directly. It is preferred to use a fixed amount of draining goroutines.
func (bp *BalancerPoolScrapper) getPool(poolAddress common.Address) (err error) {
	pool, err := balancerpoolcontract.NewBalancerPoolContractCaller(poolAddress, bp.restClient)
	if err != nil {
		return errors.Wrap(err, "loading pool contract")
	}

	// ignore unfinalized pools
	// note(gravi): should we log unfinalized pools ?
	//				Currently, they are logged upon their first private transaction,
	// 				so that every public pool gets logged.
	finalized, err := pool.IsFinalized(&bind.CallOpts{})
	if err != nil {
		errors.Wrap(err, "checking if pool is finalized")
	} else if !finalized {
		return errors.New("pool is not finalized")
	}

	// note(gravi): this represents the latest known block
	header, err := bp.restClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "getting the current heading block")
	}

	tokens, err := pool.GetCurrentTokens(&bind.CallOpts{})
	if err != nil {
		return errors.Wrap(err, "getting pool's tokens")
	}

	symbols := []string{}

	// V = (b1^w1)(b2^w2)(b3^w3)
	// V is the product of each token of the pool.
	// b1/b2/b3 are the balances of the token
	// w1/w2/w3 are their respective weights
	V := big.NewFloat(1)

	for _, token := range tokens {
		// tokCaller, err := erctoken.NewERC20Token(token, bp.restClient)
		// tokCaller, err := erctoken.NewERC20TokenCaller(token, bp.restClient)
		tokCaller, err := baltokencontract.NewBALTokenContractCaller(token, bp.restClient)
		if err != nil {
			return errors.Wrapf(err, "creating bal token contract caller for %s", token.Hex())
		}

		// ERC-20 specification does not includes symbol informations, so fetching the symbol won't work on some contracts
		// in case it fails, we'll use a map to associate token addresses with symbol
		symbol, err := tokCaller.Symbol(&bind.CallOpts{})
		if err != nil {
			// add tokens with buggy symbols here
			symbolMap := map[string]string{
				"0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2": "MKR",
				"0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359": "SAI",
				"0xC011A72400E58ecD99Ee497CF89E3775d4bd732F": "SNX",
				"0xF1290473E210b2108A85237fbCd7b6eb42Cc654F": "HEDG",
				"0xeb269732ab75A6fD61Ea60b06fE994cD32a83549": "DF-USDx",
				"0x431ad2ff6a9C365805eBaD47Ee021148d6f7DBe0": "DF",
			}

			var ok bool
			symbol, ok = symbolMap[token.Hex()]
			if !ok {
				log.WithField("tokenAddrs", token.Hex()).Error("unknown symbol for token")
				symbol = "UNKNOWN_SYMBOL"
			}
		}

		// get token weight in pool
		weight, err := pool.GetNormalizedWeight(&bind.CallOpts{}, token)
		if err != nil {
			return errors.Wrapf(err, "getting normalized weight for token %s in pool", symbol)
		}
		weightFloat := new(big.Float).SetInt(weight)
		weightFloat = weightFloat.Quo(weightFloat, big.NewFloat(10e17))

		// get token balance in pool
		balance, err := pool.GetBalance(&bind.CallOpts{}, token)
		if err != nil {
			return errors.Wrapf(err, "getting balance for token %s in pool", symbol)
		}
		balanceFloat := new(big.Float).SetInt(balance)
		balanceFloat = balanceFloat.Quo(balanceFloat, big.NewFloat(10e17))

		balanceFloatFloat, _ := balanceFloat.Float64()
		weightFloatFloat, _ := weightFloat.Float64()

		thing := math.Pow(balanceFloatFloat, weightFloatFloat)
		V.Mul(V, new(big.Float).SetFloat64(thing))

		symbols = append(symbols, string(symbol))
	}

	balance, err := pool.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return errors.Wrap(err, "getting pool's total supply of BPT")
	}
	balanceFloat := new(big.Float).SetInt(balance)
	balanceFloatFloat, _ := balanceFloat.Float64()
	balanceFloatFloat /= 10e17

	VFloat, _ := V.Float64()
	rate := VFloat / balanceFloatFloat // v = V / num_bpt
	if math.IsNaN(rate) {
		rate = 0
	}

	pr := models.FarmingPool{
		// Balance is the total supply of pool token.
		Balance: balanceFloatFloat,

		TimeStamp:    time.Now(),
		PoolID:       poolAddress.String(),
		ProtocolName: bp.scraper.poolName,

		OutputAsset: []string{"BPT"},
		InputAsset:  symbols,

		Rate: rate,

		BlockNumber: header.Number.Int64(),

		// note(gravi): there are private pools for which only the creator can add liquidity and swap tokens
		//              currently, transactions for these pools are watched so that they are logged when they
		//              get their first public transaction. Also, pools that are discovered but that
		//              are private are ignored. Should we log private pools anyway ? Adding a "isPrivate"
		//				property to models.FarmingPool would allow to log them.
	}

	bp.scraper.chanPoolInfo <- &pr
	return nil
}

func (bp *BalancerPoolScrapper) bootstrapPools(newPoolChan chan common.Address) {
	// first shared pools from pools.balancer.exchange
	// not necessary: will be fetched anyway from historical New Pool log events.
	// useful for testing
	newPoolChan <- common.HexToAddress("0x548b5c92a5e0006628f69c60e0085373fd5b63d9")
	newPoolChan <- common.HexToAddress("0x1eff8af5d577060ba4ac8a29a13525bb0ee2a3d5")
	newPoolChan <- common.HexToAddress("0x59a19d8c652fa0284f44113d0ff9aba70bd46fb4")
	newPoolChan <- common.HexToAddress("0xcce41676a4624f4a1e33a787a59d6bf96e5067bc")
	newPoolChan <- common.HexToAddress("0x8a649274e4d777ffc6851f13d23a86bbfa2f2fbf")
	newPoolChan <- common.HexToAddress("0x2dd7255b487a62d738110bd10f8bc4b4ea989778")
	newPoolChan <- common.HexToAddress("0xf54025af2dc86809be1153c1f20d77adb7e8ecf4")
	newPoolChan <- common.HexToAddress("0xee9a6009b926645d33e10ee5577e9c8d3c95c165")
	newPoolChan <- common.HexToAddress("0xd8e9690eff99e21a2de25e0b148ffaf47f47c972")
	newPoolChan <- common.HexToAddress("0x454c1d458f9082252750ba42d60fae0887868a3b")
	newPoolChan <- common.HexToAddress("0xe036cce08cf4e23d33bc6b18e53caf532afa8513")
	newPoolChan <- common.HexToAddress("0x5b2dc8c02728e8fb6aea03a622c3849875a48801")
	newPoolChan <- common.HexToAddress("0xb1f9ec02480dd9e16053b010dfc6e6c4b72ecad5")
	newPoolChan <- common.HexToAddress("0x72cd8f4504941bf8c5a21d1fd83a96499fd71d2c")
	newPoolChan <- common.HexToAddress("0x9dde0b1d39d0d2c6589cde1bfed3542d2a3c5b11")
	newPoolChan <- common.HexToAddress("0xe0e6b25b22173849668c85e06bc2ce1f69baff8c")
	newPoolChan <- common.HexToAddress("0xe867be952ee17d2d294f2de62b13b9f4af521e9a")
	newPoolChan <- common.HexToAddress("0xfae2809935233d4bfe8a56c2355c4a2e7d1fff1a")
	newPoolChan <- common.HexToAddress("0xd3c8dcfcf2a5203f6a3210591dabea14e181db2d")
	newPoolChan <- common.HexToAddress("0x7c90a3cd7ec80dd2f633ed562480abbeed3be546")
	newPoolChan <- common.HexToAddress("0x9762c600a39bda3359f50a1fb1f90945cead379d")
	newPoolChan <- common.HexToAddress("0x9866772a9bdb4dc9d2c5a4753e8658b8b0ca1fc3")
	newPoolChan <- common.HexToAddress("0xf9ab7776c7baeed1d65f0492fe2bb3951a1787ef")
	newPoolChan <- common.HexToAddress("0xa5da8cc7167070b62fdcb332ef097a55a68d8824")
	newPoolChan <- common.HexToAddress("0xe42237f32708bd5c04d69cc77e1e36c8f911a016")
	newPoolChan <- common.HexToAddress("0xa795600590a7da0057469049ab8f1284baed977e")
	newPoolChan <- common.HexToAddress("0x834fb8276b4e8a24010e2108fdd7f8417c8922bd")
	newPoolChan <- common.HexToAddress("0x415900c6e18b89531e3e24c902b05c031c71a925")
	newPoolChan <- common.HexToAddress("0x6b9887422e2a4ae11577f59ea9c01a6c998752e2")
	newPoolChan <- common.HexToAddress("0xa0313dea2ed259edd2d95986cc9046d1a65f649b")
	newPoolChan <- common.HexToAddress("0x30cb859317e171832b064c97cc03ccb081954d1b")
	newPoolChan <- common.HexToAddress("0x1d408746abe6956488c2f6c3b05c21dad67b9785")
	newPoolChan <- common.HexToAddress("0x5e37910cfb8de1b14ec4e4bac0bec27c35dc07d5")
	newPoolChan <- common.HexToAddress("0x1b8874baceaafba9ea194a625d12e8b270d77016")
	newPoolChan <- common.HexToAddress("0x5b18df96d3c8b9f1d1b9e38752498f92d1e2d490")
	newPoolChan <- common.HexToAddress("0x1b2cf5418e8982ab13a90820594dc6c870f95730")
	newPoolChan <- common.HexToAddress("0x35333cf3db8e334384ec6d2ea446da6e445701df")
	newPoolChan <- common.HexToAddress("0xd44082f25f8002c5d03165c5d74b520fbc6d342d")
	newPoolChan <- common.HexToAddress("0xe010fcda8894c16a8acfef7b37741a760faeddc4")
	newPoolChan <- common.HexToAddress("0x57755f7dec33320bca83159c26e93751bfd30fbe")
	newPoolChan <- common.HexToAddress("0x917fbcdc6e7077c55c9ea1cf9442809dd55fcb91")
	newPoolChan <- common.HexToAddress("0x41284a88d970d3552a26fae680692ed40b34010c")
	newPoolChan <- common.HexToAddress("0xdbe29107464d469c64a02afe631aba2e6fabedce")
	newPoolChan <- common.HexToAddress("0x68a241796628ecf44e48f0533fb00d07dd3419d2")
	newPoolChan <- common.HexToAddress("0xc81d50c17754b379f1088574cf723be4fb00307d")
	newPoolChan <- common.HexToAddress("0xe93e8aa4e88359dacf33c491cf5bd56eb6c110c1")
	newPoolChan <- common.HexToAddress("0xde0999ee4e4bea6fecb03bf4ebef2626942ec6f5")
	newPoolChan <- common.HexToAddress("0x9fa4eb27c2a859a76caa1da9fe4ae2e48c2307c3")
	newPoolChan <- common.HexToAddress("0xae95d3198d602acfb18f9188d733d710e14a27dd")
	newPoolChan <- common.HexToAddress("0x62713503753e179cde1066d9efc854088a4debf0")
	newPoolChan <- common.HexToAddress("0x4118f3935231c517a71e154ab945249e5b9ed684")
	newPoolChan <- common.HexToAddress("0x987d7cc04652710b74fff380403f5c02f82e290a")
	newPoolChan <- common.HexToAddress("0x01d5314ca775a0ec2ed11b19ff745a08d9d3c7f9")
	newPoolChan <- common.HexToAddress("0x80cba5ba9259c08851d94d6bf45e248541fb3e86")
	newPoolChan <- common.HexToAddress("0x003a70265a3662342010823bea15dc84c6f7ed54")
	newPoolChan <- common.HexToAddress("0xbed340a301b4f07fa7b61ab9e0767faa172f530d")
	newPoolChan <- common.HexToAddress("0xe2eb726bce7790e57d978c6a2649186c4d481658")
	newPoolChan <- common.HexToAddress("0xa33a45749d4a4118ad119f7e6c1474ab568f5928")
	newPoolChan <- common.HexToAddress("0xb38b0c480a451db976837a1a464af95bb0f3f5e2")
	newPoolChan <- common.HexToAddress("0x4bcbc51fbabe5ed04b4cd93d361674fc3fb519b8")
	newPoolChan <- common.HexToAddress("0x39ee0f29253017d69909b7286caf8c080e9d6da2")
	newPoolChan <- common.HexToAddress("0x77f5f3dfd95d0b061c9ef47c4f4ca4681d807776")
	newPoolChan <- common.HexToAddress("0x9b6d305147931afc298e08303d03fd16fb583431")
	newPoolChan <- common.HexToAddress("0x032d9bc3f3c1042b431f29df63aaa547f5ed6ee6")
	newPoolChan <- common.HexToAddress("0x9b208194acc0a8ccb2a8dcafeacfbb7dcc093f81")
	newPoolChan <- common.HexToAddress("0xd7067192df19c415ee35cc86708020eaf924cd75")
	newPoolChan <- common.HexToAddress("0x98a48026355cffaacf1dade7928442aa856a5e63")
	newPoolChan <- common.HexToAddress("0xc3eb65c18902357f0c3577ced27b1fa914cd0e57")
	newPoolChan <- common.HexToAddress("0x73eba399fbbea50852359ff8b8d0e3eba1f22500")
	newPoolChan <- common.HexToAddress("0x75e751509ef6aa0325452a8c039003565b630f89")
	newPoolChan <- common.HexToAddress("0x02b77ad79e27ad80a0a0868cc7528520e8442c9b")
	newPoolChan <- common.HexToAddress("0xc8782bd1604ab96712818f1f55e11df0e80a0873")
	newPoolChan <- common.HexToAddress("0x1af23b311f203844108137d6ee399109e4981401")
	newPoolChan <- common.HexToAddress("0x5311b0bfd1eb656522ef75ea4d22ce8bcfa6fd42")
	newPoolChan <- common.HexToAddress("0x2f49eea1efc1b04e9ecd3b81321060e29db26a19")
	newPoolChan <- common.HexToAddress("0x5390f43ef8b8fe0b103e89f1ca74bfb985669f7b")
	newPoolChan <- common.HexToAddress("0xed5ad5f258eef6a9745042bde7d46e8a5254c183")
	newPoolChan <- common.HexToAddress("0xd06726afe91a423afec003fbb2f075ef3adc3cd7")
	newPoolChan <- common.HexToAddress("0x70f1c87b22d7371c68e3ea64870833a2a64cf8b3")
	newPoolChan <- common.HexToAddress("0x35f6d69f83773d92bbfa326dae2f4e337ca577a5")
	newPoolChan <- common.HexToAddress("0xc5eed59f8692aa9ba14ed1e3b3ddff91cc1d2d34")
	newPoolChan <- common.HexToAddress("0x653911da49db4cdb0b7c3e4d929cfb56024cd4e6")
	newPoolChan <- common.HexToAddress("0xbfee82ff0402b1a24f895ccc2545f6187369b4c3")
	newPoolChan <- common.HexToAddress("0x280267901c175565c64acbd9a3c8f60705a72639")
	newPoolChan <- common.HexToAddress("0xe26a220a341eaca116bda64cf9d5638a935ae629")
	newPoolChan <- common.HexToAddress("0x8c7769f9f1e5042c0809b8702e4b9947b1bcb3f3")
	newPoolChan <- common.HexToAddress("0x239c53350be2b399fe104f0e207e26dfb24164f3")
	newPoolChan <- common.HexToAddress("0x16cac1403377978644e78769daa49d8f6b6cf565")
	newPoolChan <- common.HexToAddress("0x6cb5c5cb789fae62ce5ce280e1fbc5dd3bbdad81")
	newPoolChan <- common.HexToAddress("0x930ae255053e40f430a4fda533eae0de5b131924")
	newPoolChan <- common.HexToAddress("0x79c44cd4291a72aa26a2f04a3559daf67e3ff761")
	newPoolChan <- common.HexToAddress("0xea63de726a757202242b253361a6d5f42f982094")
	newPoolChan <- common.HexToAddress("0x95f0aa355f4251291e6413dd0174488f0ca8d2db")
	newPoolChan <- common.HexToAddress("0x68198e4ed6975204b3467dc217166d9ff1cbb57a")
	newPoolChan <- common.HexToAddress("0x2cf9106faf2c5c8713035d40df655fb1b9b0f9b9")
	newPoolChan <- common.HexToAddress("0x535442c3127656f389041f6986991e0128e4ff06")
	newPoolChan <- common.HexToAddress("0xb21e53d8bd2c81629dd916eead08d338e7fcc201")
	newPoolChan <- common.HexToAddress("0xc409d34accb279620b1acdc05e408e287d543d17")
	newPoolChan <- common.HexToAddress("0xd4dbf96db2fdf8ed40296d8d104b371adf7dee12")
	newPoolChan <- common.HexToAddress("0x7fc95945eaa14e7a2954052a4c9bfbaa79d170ae")
	newPoolChan <- common.HexToAddress("0x5aec4cff7fc3880ade1582e5e37cf89152e70ace")
	newPoolChan <- common.HexToAddress("0x07626319b650c26b5d81dbfbe4f9d1c0cd24781a")
	newPoolChan <- common.HexToAddress("0xe5d505055a4c607cf9049b0629c7f069abaa8dbf")
	newPoolChan <- common.HexToAddress("0x4304ae5fd14cec2299caee4e9a4afbedd046d612")
	newPoolChan <- common.HexToAddress("0xd3466427f7ee5d0d9cd68b36df677e51bda26321")
	newPoolChan <- common.HexToAddress("0xc2b070ebab858b2e638b21ceae937146f9cb6d2b")
	newPoolChan <- common.HexToAddress("0x0ffc0ac3139a7cde29e6858c369b32cbbaf41cf4")
	newPoolChan <- common.HexToAddress("0x2cdfd9e12e482d0238a2e1db8d1ae11cb47c4515")
	newPoolChan <- common.HexToAddress("0xb189802d50c45515369bc5373d5ca71aa9b70d88")
	newPoolChan <- common.HexToAddress("0x9dcc2104f25d40db2f455b458fcbe538c6812bbc")
	newPoolChan <- common.HexToAddress("0x17996cbddd23c2a912de8477c37d43a1b79770b8")
	newPoolChan <- common.HexToAddress("0xf7dde178fabe8386ade62d584019326a203b5394")
	newPoolChan <- common.HexToAddress("0x4a77b41ec5708bfa5ca0dccead8a4145e31c85a0")
	newPoolChan <- common.HexToAddress("0xaf71d6c242a00e8364ea0ef3c007f3413e975011")
	newPoolChan <- common.HexToAddress("0x76730720bc5ef699e8e6f691065948be2fb3632e")
	newPoolChan <- common.HexToAddress("0x56efe58653d94f3b9c87e32cc93cd20bd0e27e14")
	newPoolChan <- common.HexToAddress("0x24f598f7df68f663bcac6f94bec005a48570d7a4")
	newPoolChan <- common.HexToAddress("0x9011f62a83c8a43a55e52b555fd3bd93313f4372")
	newPoolChan <- common.HexToAddress("0x24d97eed6e171e70c82bc60afd37c7d1e549a0ad")
	newPoolChan <- common.HexToAddress("0x10d9b57f769fbb355cdc2f3c076a65a288ddc78e")
	newPoolChan <- common.HexToAddress("0x9588dcdc8b4fe73f1312469818405f02465bcfbc")
	newPoolChan <- common.HexToAddress("0xba20d4f41121b997a1eaca6d938ac40b67dad226")
	newPoolChan <- common.HexToAddress("0xe5d1fab0c5596ef846dcc0958d6d0b20e1ec4498")
	newPoolChan <- common.HexToAddress("0x01c1adb4ad6678d2bf2b772301f111d257b7b109")
	newPoolChan <- common.HexToAddress("0x2b36d183be387ca2cf81b63efdddb030f3a643eb")
	newPoolChan <- common.HexToAddress("0xe3f9cf7d44488715361581dd8b3a15379953eb4c")
	newPoolChan <- common.HexToAddress("0x96d99093f22719dd06fb8db8e93779979a2acab3")
	newPoolChan <- common.HexToAddress("0xb53acbbe6db0f80f9dab5b9dc9e95e2e2393f000")
	newPoolChan <- common.HexToAddress("0x294de1cde8b04bf6d25f98f1d547052f8080a177")
	newPoolChan <- common.HexToAddress("0x1eeadcfd9b4889256f9ec49f9ca06b3291e0c7c4")
	newPoolChan <- common.HexToAddress("0x814c01ecb99b090cae15adfe662d949b3244c51c")
	newPoolChan <- common.HexToAddress("0xb62357d2a05f6bc6e4f05078de9f5fc5c9c004cb")
	newPoolChan <- common.HexToAddress("0x4c34a687906092ec11cc04ddf30b71e29747ed76")
	newPoolChan <- common.HexToAddress("0x70985e557ae0cd6dc88189a532e54fbc61927bad")
	newPoolChan <- common.HexToAddress("0xf28bf5bcd8d245d4c3247f7e78a5c0dc64d4cf4d")
	newPoolChan <- common.HexToAddress("0x01d7e6a79f6e6dc6f0884743078f76ac1239520a")
	newPoolChan <- common.HexToAddress("0x10dd17ecfc86101eab956e0a443cab3e9c62d9b4")
	newPoolChan <- common.HexToAddress("0xcd461b73d5fc8ea1d69a600f44618bdfac98364d")
	newPoolChan <- common.HexToAddress("0x2eb6cfbffc8785cd0d9f2d233d0a617bf4269eef")
	newPoolChan <- common.HexToAddress("0x05f661a1bd2da5e59c8e06b92acdcde2706110e5")
	newPoolChan <- common.HexToAddress("0xfb23a5806c7beeeb13a2519fae6253eec9239617")
	newPoolChan <- common.HexToAddress("0x330416c863f2acce7af9c9314b422d24c672534a")
	newPoolChan <- common.HexToAddress("0x36a57f5ea6cceb75523639aba3224036e5f8231d")
	newPoolChan <- common.HexToAddress("0x92e7eb99a38c8eb655b15467774c6d56fb810bc9")
	newPoolChan <- common.HexToAddress("0xf64f38ae6b1e6144f292ecfa0a7c9236885f032e")
	newPoolChan <- common.HexToAddress("0xc55458d09ae7b98f23c8d1b977e4ab003b8a3800")
	newPoolChan <- common.HexToAddress("0xf3ad3821a0036fd9caec0a4cf5d31515fad03f00")
	newPoolChan <- common.HexToAddress("0x23f2369dd17aecf5e719421296856d1188b8ca92")
	newPoolChan <- common.HexToAddress("0xee60ca08cca968dc1b36e7bbcc779efe34d7ef62")
	newPoolChan <- common.HexToAddress("0x877d2a28ba931964b3479c23191c72a63993c2e7")
	newPoolChan <- common.HexToAddress("0x47de4e96f2366ba4eab132da944ac8e49f59cfb9")
	newPoolChan <- common.HexToAddress("0x0badc63d5cc7a25007f7c35acc23369c00df6ee7")
	newPoolChan <- common.HexToAddress("0xf1880d81ce7d3b4b0d96ea0665cdb03cb7332c52")
	newPoolChan <- common.HexToAddress("0x7d014a7464c91f20da99a3e6f77bc5506ddf3c5e")
	newPoolChan <- common.HexToAddress("0xe5ac9548275787cd86df2350248614afab0088ee")
	newPoolChan <- common.HexToAddress("0xd9646c5173f2f4381da50f0b3efa85f9529f485b")
	newPoolChan <- common.HexToAddress("0x7f032763009b53c7dbbc72697724ca83c1c239b1")
	newPoolChan <- common.HexToAddress("0x74a5d106b18c86dc37be5c817093a873cdcff216")
	newPoolChan <- common.HexToAddress("0x8c6f57b7eb93991c2e90ea8f69265b78d2fa479a")
	newPoolChan <- common.HexToAddress("0x5bfdddef3d0bab839b797d23b35797b107923095")
	newPoolChan <- common.HexToAddress("0xa7d7d09484fa6e5f497b6b687f979509373c6530")
	newPoolChan <- common.HexToAddress("0x2471de1547296aadb02cc1af84afe369b6f67c87")
	newPoolChan <- common.HexToAddress("0x795dfdfd413c4a9492cef5b58723f9fb3c8af624")
	newPoolChan <- common.HexToAddress("0x57f07147a30927eb1d9aa1c7f89ac3c806fd0525")
	newPoolChan <- common.HexToAddress("0xdff4f867855fd7db4d240b60fd0a88f6a049427a")
	newPoolChan <- common.HexToAddress("0x4019ba88158daa468a063ac48171a3bfe8cd9f3b")
	newPoolChan <- common.HexToAddress("0xf859a6c44f18aa5f7fc71387b603c6b63a0dbe04")
	newPoolChan <- common.HexToAddress("0xf11da347012b9752931202a40ed1328982551f5c")
	newPoolChan <- common.HexToAddress("0x7cd40497c8284df1018673e6b29c3d97ef811685")
	newPoolChan <- common.HexToAddress("0x05c04d16da1a7efd6d6768a3f0e609c09f99710b")
	newPoolChan <- common.HexToAddress("0x2e22a35a29f0ccafe8c6fb935f7d0269f706ea72")
	newPoolChan <- common.HexToAddress("0x7a9d2c9d07c4e5ff4900e591526691ff7e7a1220")
	newPoolChan <- common.HexToAddress("0xe969991ce475bcf817e01e1aad4687da7e1d6f83")
	newPoolChan <- common.HexToAddress("0x57c9821179a4d94657161eeaad9dfdf5280f86db")
	newPoolChan <- common.HexToAddress("0x82374bf03dbfe35845cb5b90b6c9a660a49f69aa")
	newPoolChan <- common.HexToAddress("0xf01d9f08fdd1daa58904710f8d23c54f924e32e0")
	newPoolChan <- common.HexToAddress("0xae1019cfc59dc21a2395c8b38a6f6d0df61d2c22")
	newPoolChan <- common.HexToAddress("0xf41a12bf3e5dadd8bcc8494a7c7cb80101331098")
	newPoolChan <- common.HexToAddress("0x7ae275fafdbcded0a4cdc6e25884b849f00efc48")
	newPoolChan <- common.HexToAddress("0xa6cf0a736fbb441b3cd39a476699eeec4737f28e")
	newPoolChan <- common.HexToAddress("0x3031745e732dce8fecccc94aca13d5fa18f1012d")
	newPoolChan <- common.HexToAddress("0xf3168b50751173e150af543997abe0fec5d58b7f")
	newPoolChan <- common.HexToAddress("0x2e41132dab88a9bad80740a1392d322bf023d494")
	newPoolChan <- common.HexToAddress("0xb485fe9998285f23b86b3b99f41799c367894166")
	newPoolChan <- common.HexToAddress("0xfe5fa46409a8aacda43e3d32ea9b30d10279a459")
	newPoolChan <- common.HexToAddress("0x1373e57f764a7944bdd7a4bd5ca3007d496934da")
	newPoolChan <- common.HexToAddress("0xe02e8ddecbe302ba1a40107c726e3ee889b0ff10")
	newPoolChan <- common.HexToAddress("0x349388dafb463ba2bc5cedc8ebc55c25d70f61bb")
	newPoolChan <- common.HexToAddress("0x6b74fb4e4b3b177b8e95ba9fa4c3a3121d22fbfb")
	newPoolChan <- common.HexToAddress("0x6f7f69409f7fd631309649db32f6775a5ca834e6")
	newPoolChan <- common.HexToAddress("0xc351648240e9fc9213e8b6016e566f95c5fa7909")
	newPoolChan <- common.HexToAddress("0xee8337f497f234442160935f210c73dc47eb2676")
	newPoolChan <- common.HexToAddress("0x4b0b0bf60abbf79a2fd028e4d52ac393982488ce")
	newPoolChan <- common.HexToAddress("0xb537ae8eebe8a7bec9ee8720c936fd15c35940fe")
	newPoolChan <- common.HexToAddress("0xd8296132c6076ab18e22eeb8c7ce08a319215400")
	newPoolChan <- common.HexToAddress("0x92d179ee51a027a254feee0c5d91cc1f7058200b")
	newPoolChan <- common.HexToAddress("0x95c4b6c7cff608c0ca048df8b81a484aa377172b")
	newPoolChan <- common.HexToAddress("0x5a0d85166a20f9cd27be2cb293e4d10188f0c97d")
	newPoolChan <- common.HexToAddress("0x0b0dea284a07d60d026691476b130f160e48bd0a")
	newPoolChan <- common.HexToAddress("0x49e4cb1bebc5a060505d4571ec50bd2b65c3ed11")
	newPoolChan <- common.HexToAddress("0xeb85b2e12320a123d447ca0da26b49e666b799db")
	newPoolChan <- common.HexToAddress("0xf8a01bdab3c85edce962c2eb026a6283163ea9e7")
	newPoolChan <- common.HexToAddress("0x4833e8b56fc8e8a777fcc5e37cb6035c504c9478")
	newPoolChan <- common.HexToAddress("0xe763fdbe8dc96687ba050561f59608ca1e977e5b")
	newPoolChan <- common.HexToAddress("0x815f8ef4863451f4faf34fbc860034812e7377d9")
	newPoolChan <- common.HexToAddress("0x9e7cbc6e96993e7c5abb3ec4cb30aa8275618b2b")
	newPoolChan <- common.HexToAddress("0x7bee47379c01a8a2f73c6506d980ac628388ee00")
	newPoolChan <- common.HexToAddress("0xbefb01de17ee3c22b7a798f56ad8cf5c43284c12")
	newPoolChan <- common.HexToAddress("0xee4148ff794cf1d0976d2063a32a8f562abe6d75")
	newPoolChan <- common.HexToAddress("0x224e906efe965dbe8ceb71e3cf1d5536c9ce196d")
	newPoolChan <- common.HexToAddress("0x9a16bddf0e89541d6abea035b706968465241e4f")
	newPoolChan <- common.HexToAddress("0xf86ca32514a22fcd350a99121f34a0e70a059ec2")
	newPoolChan <- common.HexToAddress("0xdd62dc215524f5809e04478a51e5d7665a9f9047")
	newPoolChan <- common.HexToAddress("0xa738103affdbc04101842a5b7b04576e17924fb8")
	newPoolChan <- common.HexToAddress("0xa69e2d058d50394cc3b4f1a1aacb41e6809d4be9")
	newPoolChan <- common.HexToAddress("0x9157901e72ffaa25b558952b0ca35c296ec853d3")
	newPoolChan <- common.HexToAddress("0x99e582374015c1d2f3c0f98d0763b4b1145772b7")
	newPoolChan <- common.HexToAddress("0xc88fd90b4b410640b755d733e8ee69930e39522d")
	newPoolChan <- common.HexToAddress("0x516bd93d659a8e3b6028022bb07e49b435a1e74a")
	newPoolChan <- common.HexToAddress("0x4512330f4492672adf1f1bacb3f709df79ae62a9")
	newPoolChan <- common.HexToAddress("0xd9b1a242a7a067cce13b5418c0c4c1d307036ba8")
	newPoolChan <- common.HexToAddress("0x66840d5c7b99a0b133ea982915ad24e5ee78bd64")
	newPoolChan <- common.HexToAddress("0x9457facebbe9a8730684f511bdcca760879411f2")
	newPoolChan <- common.HexToAddress("0x2e933a4612c4fa03b258c2fd183c299d099975be")
	newPoolChan <- common.HexToAddress("0x50a6af8d1b16e9fb18c912562239d65c0ff6724a")
	newPoolChan <- common.HexToAddress("0x0d88b55317516b84e78fbba7cde3f910d5686901")
	newPoolChan <- common.HexToAddress("0xa5432a868622a4a75d5fac8d9e201fd7950ff28a")
	newPoolChan <- common.HexToAddress("0xbd7a8f648262b6cb29d38b575df9f27e6cdecde1")
	newPoolChan <- common.HexToAddress("0x4f0f422874e1525d88551d6a755116b9b7dcea7c")
	newPoolChan <- common.HexToAddress("0x1a3c63a48c806a813f35462727b158b08e4cd68c")
	newPoolChan <- common.HexToAddress("0xdc7d39628855b6013000c9af957e6cd484beda6c")
	newPoolChan <- common.HexToAddress("0xa1ec308f05bca8acc84eaf76bc9c92a52ac25415")
	newPoolChan <- common.HexToAddress("0x7e2105a44c1e316a87591c99bb8d65f9ecd64791")
	newPoolChan <- common.HexToAddress("0xc6c68cae4681d7b3b3deb54091d7e10d7ec7757d")
	newPoolChan <- common.HexToAddress("0x29d2e5957249e5c8c2b22a9101002183f7d35382")
	newPoolChan <- common.HexToAddress("0x0ce8f4e28b175eeec37ba7e28560852b61f72971")
	newPoolChan <- common.HexToAddress("0x23b2b580f6d733d87eeb0d23ab7f3a0128326636")
	newPoolChan <- common.HexToAddress("0xd961daa343c8bf4c945c4fc66fc4eebe9c6a8472")
	newPoolChan <- common.HexToAddress("0x09bc559189d309695804cfe0ed64d8fa23128206")
	newPoolChan <- common.HexToAddress("0x8b0460c23d4e177c3b22ab6ea872d64f101ae5a2")
	newPoolChan <- common.HexToAddress("0x58ef3abab72c6c365d4d0d8a70039752b9f32bc9")
	newPoolChan <- common.HexToAddress("0xe97f5d1d35bcca8dfcd290ccd1aa26c95995b77c")
	newPoolChan <- common.HexToAddress("0x65bdb4b10b381378050a4c0bf910930e82421946")
	newPoolChan <- common.HexToAddress("0xcf2e274fc7d55d3d2b1c50d5b937ec76cfb8d253")
	newPoolChan <- common.HexToAddress("0x6545773483142fd781023ec74ee008d93ad5466b")
	newPoolChan <- common.HexToAddress("0xbfdef139103033990082245c24ff4b23dafd88cf")
	newPoolChan <- common.HexToAddress("0x63b642819845f3edacedd4cf614a4ac0a7950cb0")
	newPoolChan <- common.HexToAddress("0x150170b0aaf7539bd6a2709efbaa5ba2765c7729")
	newPoolChan <- common.HexToAddress("0x4c0ebb723bcb72626dcf8387e642a9aa4fc69231")
	newPoolChan <- common.HexToAddress("0x5e44a81dbc615871968219698098fb050e58a0f3")
	newPoolChan <- common.HexToAddress("0x6153dfd06f58b929e4c1b184b10ac799223ca095")
	newPoolChan <- common.HexToAddress("0x7842792a8471d0f5ae645f513cc5999b1bb6b182")
	newPoolChan <- common.HexToAddress("0x7aa7305838d858f7089f86fc136d0816baacb0d8")
	newPoolChan <- common.HexToAddress("0x64bd63cc9d697a58c307791b43a34a08c722491d")
	newPoolChan <- common.HexToAddress("0xb454f3f2b982e56b962392d93d01e5895b532995")
	newPoolChan <- common.HexToAddress("0x1811a719a05d20b6447ca556a54f00f9c14578be")
	newPoolChan <- common.HexToAddress("0x01dfeaf2f6b648296b4c9dc4f0c5fddcb0984c0d")
	newPoolChan <- common.HexToAddress("0x2a8d66fbd0efb5f72e9b1d0dc9912d4456ac6b55")
	newPoolChan <- common.HexToAddress("0xd2f574637898526fcddfb3d487cc73c957fa0268")
	newPoolChan <- common.HexToAddress("0x2c70ea56c08ba4279d62f83ca14952ed28944072")
	newPoolChan <- common.HexToAddress("0xfdef5cfe60c3b97de7478a97185310f75e11bf8b")
	newPoolChan <- common.HexToAddress("0xa4095bf13ae0dbc9eb41084514d67846fbd55118")
	newPoolChan <- common.HexToAddress("0xc48b8329d47ae8fd504c0b81cf8435486380e858")
	newPoolChan <- common.HexToAddress("0x5a82503652d05b21780f33178fdf53d31c29b916")
	newPoolChan <- common.HexToAddress("0xef6b465a5ea69b502433d960248b6c99023fbda9")
	newPoolChan <- common.HexToAddress("0xe269e0151f6bc8e226211d517516fa808667c3cb")
	newPoolChan <- common.HexToAddress("0x3a17f93ef0baee1b9903449c92a350fe6b47bd2c")
	newPoolChan <- common.HexToAddress("0xde7e006a9583a4cdcbdd0e233155d88d80962b31")
	newPoolChan <- common.HexToAddress("0x60626db611a9957c1ae4ac5b7ede69e24a3b76c5")
	newPoolChan <- common.HexToAddress("0x594415978a756c5b02eabdff98d867cdda65e888")
	newPoolChan <- common.HexToAddress("0x55d291693f444d9ebd773b38ca130cce10ec25a5")
	newPoolChan <- common.HexToAddress("0x8bf495304697e420b3b766ca4d2fb6c9051f3b04")
	newPoolChan <- common.HexToAddress("0x19db8e1a8553955757e87bc2620b76cca4262577")
	newPoolChan <- common.HexToAddress("0x471eb7dcf6647abaf838a5aad94940ce6932198c")
	newPoolChan <- common.HexToAddress("0x78cea5ef876187e44ace4d8c3bbeb4bd45e91355")
	newPoolChan <- common.HexToAddress("0x3d67f5fa074f39d946b29fe86bd2539709771e1c")
	newPoolChan <- common.HexToAddress("0xf7f98cacc470dd82627424387c53ecb9d58ff098")
	newPoolChan <- common.HexToAddress("0x5cac4d8f47baf46102400c003b0f01adf3d58315")
	newPoolChan <- common.HexToAddress("0xaad387f18e015cf9be4bdf96854845be51ddc977")
	newPoolChan <- common.HexToAddress("0x05d7c706210c981778dce43de41ee08766378570")
	newPoolChan <- common.HexToAddress("0xa054ba2c35382c0520c66d9c32f00de8bd7330e8")
	newPoolChan <- common.HexToAddress("0xe75ce27c39cf46d086c75f23648a154fdcc97668")
	newPoolChan <- common.HexToAddress("0x7d2f4bcb767eb190aed0f10713fe4d9c07079ee8")
	newPoolChan <- common.HexToAddress("0x5f0fb6a75787be0d217edc11a0a9a6ce6a6695fa")
	newPoolChan <- common.HexToAddress("0x1159907045848ed563a7ab883d2abe75b02d4723")
	newPoolChan <- common.HexToAddress("0x53b89ce35928dda346c574d9105a5479cb87231c")
	newPoolChan <- common.HexToAddress("0x41523be0ab764a2aef588b2d49eb424b71bcb9b0")
	newPoolChan <- common.HexToAddress("0xbcf42073e917e6629eddd60246a433a5d0305bfe")
	newPoolChan <- common.HexToAddress("0x8d924a1287dfa27ff5f7111b31d7392674a3e928")
	newPoolChan <- common.HexToAddress("0x1dd2cb1745aaf061380a46d5dd3b1ee2f5d5e8ad")
	newPoolChan <- common.HexToAddress("0x8434019a3ac641d10f93f20efdf5a56456adf8af")
	newPoolChan <- common.HexToAddress("0xf24a132fb79352f80f44ff01eac4d80ccdf5390b")
	newPoolChan <- common.HexToAddress("0xc19e3035a4f6f69b981c7dc2f533e862aa3af496")
	newPoolChan <- common.HexToAddress("0xd7b6d14b76595ab7938c2c781079d787ef3bd5ee")
	newPoolChan <- common.HexToAddress("0xfaac045833b86bb7c8889d146733f09e4c427cfa")
	newPoolChan <- common.HexToAddress("0x742569fd5266486fd2a50171dbdc88b8ee893ee9")
	newPoolChan <- common.HexToAddress("0xed201fc1d50a49ecf8b6e99730012dfff4daf8d0")
	newPoolChan <- common.HexToAddress("0x6fe47a03a1c6d9f055e1d212f4840027f8a9008d")
	newPoolChan <- common.HexToAddress("0x17b91b5189ff7ebf9fdfc4f1dd194b6eeee4502a")
	newPoolChan <- common.HexToAddress("0x1be4c2afd3705cbe6dbbe45b1f87a637616683f5")
	newPoolChan <- common.HexToAddress("0x3bd8f51b08b0518d33c848d121286e14f52c5c70")
	newPoolChan <- common.HexToAddress("0xb711ccdfb3609ea144930c9fb5e6a36219b478de")
	newPoolChan <- common.HexToAddress("0x72717b53da57926bdf988f5645c23f7a3214ac9a")
	newPoolChan <- common.HexToAddress("0x97f50f649db045b76edebe7a947d101d9b1c5129")
	newPoolChan <- common.HexToAddress("0x2fd8cdcb9391876554dab81d09c439e8f1012ed2")
	newPoolChan <- common.HexToAddress("0x10a39605ae49bf6bee84dcd09551e3ede7401b63")
	newPoolChan <- common.HexToAddress("0xa28979ee7ac9d688e798cc5450da307ca5f51908")
	newPoolChan <- common.HexToAddress("0x4fa407ec9fad5b87a2c15ed579e021f8320017dd")
	newPoolChan <- common.HexToAddress("0x9e04b421149043c04b33865d5ecd8f6c87f174b6")
	newPoolChan <- common.HexToAddress("0xfad1ddc14c266868c57a1304863fd953eed57ecb")
	newPoolChan <- common.HexToAddress("0xf11df5baf7a9ecc6539a6e1f3e27316654a5595a")
	newPoolChan <- common.HexToAddress("0x57f604e1be07f2db3e19b58ea00d0005156889d2")
	newPoolChan <- common.HexToAddress("0x21d068e86f5537d77196d3f1223f24390f59ff3b")
	newPoolChan <- common.HexToAddress("0xc0f0ab9767ec5117cc640127255fad744ddc55b0")
	newPoolChan <- common.HexToAddress("0x3c71fad2a1fc9ea1c5f7b87491daae7f1669d21a")
	newPoolChan <- common.HexToAddress("0xfe793bc3d1ef8d38934896980254e81d0c5f6239")
	newPoolChan <- common.HexToAddress("0xe36d96440a8522844ef85a709b2e36fd255e580a")
	newPoolChan <- common.HexToAddress("0xd2bbac179784e875bde9cbbdcd3fd578c4d0851c")
	newPoolChan <- common.HexToAddress("0xbe9bf871c99015374e8aecfc711eca09cbb7ca5c")
	newPoolChan <- common.HexToAddress("0xba95fa2b47de4c971209aec5874c8204a3721e98")
	newPoolChan <- common.HexToAddress("0x7f8616a9f6916376f2ef066e79ddd93bea6b475b")
	newPoolChan <- common.HexToAddress("0xd3095c5de7b8f4476231c414c69d058f17154afe")
	newPoolChan <- common.HexToAddress("0xb135b31665bd67c4497f6eb3d5964e8897e5feb7")
	newPoolChan <- common.HexToAddress("0x9446625e3090d155fd7c74752edf4e87cf44de05")
	newPoolChan <- common.HexToAddress("0x9c9bb6d95abbdbadd6e3a421b050e653a64d9389")
	newPoolChan <- common.HexToAddress("0x5dcb5baa17a5af85aaec5d2b422cab13c0141bce")
	newPoolChan <- common.HexToAddress("0xa74485e5f668bba37b5c044c386b363f4cbd7c8c")
	newPoolChan <- common.HexToAddress("0xbddf5c35b183a30b946d0b54010b9d5c860ffae4")
	newPoolChan <- common.HexToAddress("0xe5871b2c1429059c5c3b21452a2dc38f69b4b824")
	newPoolChan <- common.HexToAddress("0x6bd73b07d48626aa4915709c57087cb098546ae6")
	newPoolChan <- common.HexToAddress("0x12a4395d92971c4f7bbf5f6c9b9ae1257b0a0e63")
	newPoolChan <- common.HexToAddress("0x2835bfa92f7ce41f17759ea8895172a94a610fef")
	newPoolChan <- common.HexToAddress("0x433d0c33288b985cf232a7e312bcfafd372460a8")
	newPoolChan <- common.HexToAddress("0x7470452303ce12947659818e1d44942899e37032")
	newPoolChan <- common.HexToAddress("0xf5ee7deafa315c2803b13adf64c620cf3681f22e")
	newPoolChan <- common.HexToAddress("0x5cf95fa2d491599d954c106b7152137002dacc9f")
	newPoolChan <- common.HexToAddress("0x50d3cac90d890ba0cc48c957bb9412b85794edb6")
	newPoolChan <- common.HexToAddress("0xa4a8438a40070bc5af16a0b6d9ee666d6390d11d")
	newPoolChan <- common.HexToAddress("0x2108f1d1328fce8c2cb92b2fac720a79195eff23")
	newPoolChan <- common.HexToAddress("0xc7c52b5d386b999ffa1e469eb723a299881e9df4")
	newPoolChan <- common.HexToAddress("0x8194efab90a290b987616f687bc380b041a2cc25")
	newPoolChan <- common.HexToAddress("0x789f2a5ebbade9b7b5bbc2fa1bd1ebd489d69c3c")
	newPoolChan <- common.HexToAddress("0x4e363b56c01f47a20c6879cc4909b18847f92446")
	newPoolChan <- common.HexToAddress("0xaf59044157ed99f3630c4c67f8870e09a5d1a6e2")
	newPoolChan <- common.HexToAddress("0x8b2e66c3b277b2086a976d053f1762119a92d280")
	newPoolChan <- common.HexToAddress("0xc6ecc714b8e50a34314244a8165a4aa0f7d16308")
	newPoolChan <- common.HexToAddress("0xc16ede1d32fbacf3b61efaa6d4f16d390fef64bd")
	newPoolChan <- common.HexToAddress("0xcbdc3bdf78c1df509f94f1f52cc5c5c093d4322b")
	newPoolChan <- common.HexToAddress("0x5a5d4a1fe3114ef5635ce740ff4a9c4d1ba8b463")
	newPoolChan <- common.HexToAddress("0x6fc73b0ac200924f94677e1994453ab1a37ab93b")
	newPoolChan <- common.HexToAddress("0xf4db33341239b33664c2c523876c07b4748ec9a7")
	newPoolChan <- common.HexToAddress("0x66c03a9d8c7df62a05f043caa4e33629780eaf7a")
	newPoolChan <- common.HexToAddress("0xc2fccdb698862ceb2d54e22f6d1eb50cbea51960")
	newPoolChan <- common.HexToAddress("0xd386bb106e6fb44f91e180228edeca24ef73c812")
	newPoolChan <- common.HexToAddress("0xe65de360befdba16426f8d1166b03fce01535a93")
	newPoolChan <- common.HexToAddress("0x285d6228ca5234edcb127077cd46b492c879760c")
	newPoolChan <- common.HexToAddress("0x660f30293df5e2506b6a2e09ce64c1454eebd9a4")
}
