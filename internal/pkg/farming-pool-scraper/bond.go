package pool

import (
	"context"
	"math/big"
	"time"

	bondstaking "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/barnbridge/bondstaking"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	daiTokenAddress   = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	usdTokenAddress   = common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	susdTokenAddress  = common.HexToAddress("0x57ab1ec28d129707052df4df418d58a2d46d5f51")
	bondTokenAddress  = common.HexToAddress("0x0391D2021f89DC339F60Fff84546EA23E337750f")
	uniLPTokenAddress = common.HexToAddress("0x6591c4bcd6d7a1eb4e537da8b78676c1576ba244")
	stakingAddress    = common.HexToAddress("0xb0fa2beee3cf36a7ac7e99b885b48538ab364853")
)

type BONDScraper struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

func NewBONDScraper(scraper *PoolScraper) *BONDScraper {
	restClient, err := ethclient.Dial("https://mainnet.infura.io/v3/251a25bd10b8460fa040bb7202e22571")
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/251a25bd10b8460fa040bb7202e22571")
	if err != nil {
		log.Fatal(err)
	}
	bd := &BONDScraper{scraper: scraper, RestClient: restClient, WsClient: wsClient}
	go bd.mainLoop()

	return bd

}

// runs in a goroutine until s is closed
func (bd *BONDScraper) mainLoop() {

	go func() {
		// Pool rates change per deposit and withdraw
		for {
			select {
			case <-bd.scraper.tickerRate.C:
				err := bd.scrapePools()
				if err != nil {
					log.Errorln("Error while Scrapping", err)
				}
			}
		}
	}()

}

func (bd *BONDScraper) scrapePools() (err error) {
	log.Info("begin scraping BOND pools")
	for _, poolDetail := range bd.getBONDPools() {

		// TO DO:
		// 1. Get vault contracts (please add to the commit) and make go bindings.
		// 2. Get tvl as commented out below
		// 3. Figure out how to get rate

		staking, err := bondstaking.NewBondstakingCaller(common.HexToAddress(poolDetail.VaultAddress), bd.RestClient)
		if err != nil {
			return err
		}
		header, err := bd.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err

		}

		currentEpoch, err := staking.GetCurrentEpoch(&bind.CallOpts{})
		if err != nil {
			return err
		}

		log.Infof("current epoch for pool %s: %v \n", poolDetail.PoolID, currentEpoch)

		var priceStablePool = big.NewInt(0)

		// //
		// bal, err := staking.GetPoolSize(&bind.CallOpts{})
		// if err != nil {
		// 	log.Error(err)
		// 	return err
		// }
		// poolBalance, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(bal), new(big.Float).SetFloat64(math.Pow10(int(decimals)))).Float64()

		pricePerFullShare := new(big.Float).SetInt(priceStablePool)
		rate, _ := big.NewFloat(0).Sub(pricePerFullShare.Quo(pricePerFullShare, new(big.Float).SetFloat64(1e18)), big.NewFloat(1)).Float64()

		var pr models.FarmingPool
		pr.TimeStamp = time.Now()
		pr.Rate = rate
		// pr.Balance = poolBalance
		pr.ProtocolName = bd.scraper.poolName
		pr.PoolID = poolDetail.PoolID
		pr.OutputAsset = []string{poolDetail.TokenName}
		pr.BlockNumber = header.Number.Int64()
		pr.InputAsset = []string{poolDetail.TokenName}
		bd.scraper.chanPoolInfo <- &pr
	}
	return

}

func (bd *BONDScraper) getBONDPools() (pools []*BONDPoolDetail) {

	pools = append(pools, &BONDPoolDetail{TokenName: "USDC/DAI/sUSD", VaultAddress: "0xB3F7abF8FA1Df0fF61C5AC38d35e20490419f4bb", PoolID: "USDC/DAI/sUSD"})
	pools = append(pools, &BONDPoolDetail{TokenName: "USDC_BOND_UNI_LP", VaultAddress: "0xC25c37c387C5C909a94055F4f16184ca325D3a76", PoolID: "USDC_BOND_UNI_LP"})
	pools = append(pools, &BONDPoolDetail{TokenName: "BOND", VaultAddress: "0x3FdFb07472ea4771E1aD66FD3b87b265Cd4ec112", PoolID: "BOND"})

	return

}

type BONDPoolDetail struct {
	TokenName    string
	VaultAddress string
	PoolID       string
}
