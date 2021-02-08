package pool

import (
	"context"
	"math"
	"math/big"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	protocolfeevault "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/loopring/feeVault"
	stakingpool "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/loopring/stakingpool"
	lrctoken "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/loopring/token"
)

type LRCPool struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

const (
	lrcStakingPoolAddress = "0xF4662bB1C4831fD411a95b8050B3A5998d8A4A5b"
	lrcFeeVaultAddress    = "0x4b89f8996892d137c3dE1312d1dD4E4F4fFcA171"
)

func NewLRCPoolScraper(scraper *PoolScraper) *LRCPool {
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	lrc := &LRCPool{scraper: scraper, RestClient: restClient, WsClient: wsClient}
	go lrc.mainLoop()

	return lrc

}

// runs in a goroutine until s is closed
func (cv *LRCPool) mainLoop() {

	go func() {
		// Pool rates change per deposit and withdraw
		for {
			select {
			case <-cv.scraper.tickerRate.C:
				err := cv.scrapePools()
				if err != nil {
					log.Errorln("Error while Scraping", err)
				}
			}

		}
	}()

}

func (cv *LRCPool) scrapePools() (err error) {
	header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	// Get contract methods
	stakingPool, err := stakingpool.NewUserStakingPoolCaller(common.HexToAddress(lrcStakingPoolAddress), cv.RestClient)
	if err != nil {
		return err
	}
	// 02/02/2021: The stakingPool contract emits zero address as fee vault address, therefore hard-coded above
	// lrcFeeVaultAddress, err := stakingPool.ProtocolFeeVaultAddress(&bind.CallOpts{})
	// if err != nil {
	// 	return err
	// }
	feevault, err := protocolfeevault.NewProtocolFeeVaultCaller(common.HexToAddress(lrcFeeVaultAddress), cv.RestClient)
	if err != nil {
		return err
	}
	tokenaddress, err := stakingPool.LrcAddress(&bind.CallOpts{})
	if err != nil {
		return err
	}
	token, err := lrctoken.NewLRCV2Caller(tokenaddress, cv.RestClient)
	if err != nil {
		return err
	}

	// Determine feeVault and staking values
	symbol, err := token.Symbol(&bind.CallOpts{})
	if err != nil {
		return err
	}
	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		return err
	}

	totalStaking, err := stakingPool.Total(&bind.CallOpts{BlockNumber: header.Number})
	if err != nil {
		return err
	}
	balance := totalStaking.Balance

	protocolFeeStats, err := feevault.GetProtocolFeeStats(&bind.CallOpts{})
	if err != nil {
		return err
	}
	remainingReward := protocolFeeStats.RemainingReward

	poolBalance, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(balance), new(big.Float).SetFloat64(math.Pow10(int(decimals)))).Float64()
	poolReward, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(remainingReward), new(big.Float).SetFloat64(math.Pow10(int(decimals)))).Float64()

	var pr models.FarmingPool
	pr.TimeStamp = time.Now()
	pr.Rate = poolReward
	pr.Balance = poolBalance
	pr.ProtocolName = cv.scraper.poolName
	pr.PoolID = symbol
	pr.OutputAsset = []string{symbol}
	pr.BlockNumber = header.Number.Int64()
	pr.InputAsset = []string{symbol}
	cv.scraper.chanPoolInfo <- &pr

	return

}
