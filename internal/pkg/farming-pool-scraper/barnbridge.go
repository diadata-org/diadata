package pool

import (
	"context"
	"math/big"
	"time"

	bondvault "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/barnbridge/vaults/bond"
	lpvault "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/barnbridge/vaults/lp"
	stablecoinvault "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/barnbridge/vaults/stablecoin"
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

type BARNBRIDGEScraper struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

func NewBARNBRIDGEScraper(scraper *PoolScraper) *BARNBRIDGEScraper {
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	bd := &BARNBRIDGEScraper{scraper: scraper, RestClient: restClient, WsClient: wsClient}
	go bd.mainLoop()

	return bd

}

// runs in a goroutine until s is closed
func (bd *BARNBRIDGEScraper) mainLoop() {

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

func (bd *BARNBRIDGEScraper) scrapePools() (err error) {
	log.Info("begin scraping BARNBRIDGE pools")
	pools := bd.getBARNBRIDGEPools()
	header, err := bd.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return err
	}

	var poolBalance float64
	var totalReward float64
	var numEpochs float64
	var currEpoch *big.Int
	var balanceBig *big.Int
	var reward *big.Int
	for _, pool := range pools {
		switch pool.PoolID {
		case "STABLECOIN":
			vault, err := stablecoinvault.NewVaultCaller(common.HexToAddress(pools[0].VaultAddress), bd.RestClient)
			if err != nil {
				return err
			}
			currEpoch, err = vault.GetCurrentEpoch(&bind.CallOpts{})
			balanceBig, err = vault.GetPoolSize(&bind.CallOpts{}, currEpoch)
			if err != nil {
				return err
			}
			reward, err = vault.TOTALDISTRIBUTEDAMOUNT(&bind.CallOpts{})
			numEpochsBig, err := vault.NROFEPOCHS(&bind.CallOpts{})
			numEpochs = float64(numEpochsBig.Int64())
		case "LP":
			vault, err := lpvault.NewVaultCaller(common.HexToAddress(pools[1].VaultAddress), bd.RestClient)
			if err != nil {
				return err
			}
			currEpoch, err = vault.GetCurrentEpoch(&bind.CallOpts{})
			balanceBig, err = vault.GetPoolSize(&bind.CallOpts{}, currEpoch)
			if err != nil {
				return err
			}
			reward, err = vault.TOTALDISTRIBUTEDAMOUNT(&bind.CallOpts{})
			numEpochsBig, err := vault.NROFEPOCHS(&bind.CallOpts{})
			numEpochs = float64(numEpochsBig.Int64())
		case "BOND":
			vault, err := bondvault.NewVaultCaller(common.HexToAddress(pools[2].VaultAddress), bd.RestClient)
			if err != nil {
				return err
			}
			currEpoch, err = vault.GetCurrentEpoch(&bind.CallOpts{})
			balanceBig, err = vault.GetPoolSize(&bind.CallOpts{}, currEpoch)
			if err != nil {
				return err
			}
			reward, err = vault.TOTALDISTRIBUTEDAMOUNT(&bind.CallOpts{})
			numEpochsBig, err := vault.NROFEPOCHS(&bind.CallOpts{})
			numEpochs = float64(numEpochsBig.Int64())
		}
		poolBalance, _ = new(big.Float).Quo(new(big.Float).SetInt(balanceBig), new(big.Float).SetFloat64(1e18)).Float64()
		totalReward = float64(reward.Int64()) / numEpochs

		var pr models.FarmingPool
		pr.Rate = totalReward
		pr.Balance = poolBalance
		pr.ProtocolName = bd.scraper.poolName
		pr.BlockNumber = header.Number.Int64()
		pr.PoolID = pool.PoolID
		pr.TimeStamp = time.Now()
		pr.OutputAsset = []string{"BOND"}
		pr.InputAsset = pool.InTokens
		bd.scraper.chanPoolInfo <- &pr
	}
	return
}

func (bd *BARNBRIDGEScraper) getBARNBRIDGEPools() (pools []*BARNBRIDGEPoolDetail) {
	pools = append(pools, &BARNBRIDGEPoolDetail{InTokens: []string{"USDC", "DAI", "sUSD"}, VaultAddress: "0xB3F7abF8FA1Df0fF61C5AC38d35e20490419f4bb", PoolID: "STABLECOIN"})
	pools = append(pools, &BARNBRIDGEPoolDetail{InTokens: []string{"USDC_BOND_UNI_LP"}, VaultAddress: "0xC25c37c387C5C909a94055F4f16184ca325D3a76", PoolID: "USDC_BOND_UNI_LP"})
	pools = append(pools, &BARNBRIDGEPoolDetail{InTokens: []string{"BOND"}, VaultAddress: "0x3FdFb07472ea4771E1aD66FD3b87b265Cd4ec112", PoolID: "BOND"})
	return
}

type BARNBRIDGEPoolDetail struct {
	InTokens     []string
	VaultAddress string
	PoolID       string
}
