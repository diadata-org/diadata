package pool

import (
	"context"
	"math"
	"math/big"
	"time"

	staking "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/barnbridge/staking/BONDstaking"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	daiTokenAddress = common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f")
	usdTokenAddress = common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	susdTokenAddress = common.HexToAddress("0x57ab1ec28d129707052df4df418d58a2d46d5f51")
	bondTokenAddress = common.HexToAddress("0x0391D2021f89DC339F60Fff84546EA23E337750f")
	uniLPTokenAddress = common.HexToAddress("0x6591c4bcd6d7a1eb4e537da8b78676c1576ba244")
)

type BONDPool struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

func NewBONDPool(scraper *PoolScraper) *BONDPool {
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	bd := &BONDPool{scraper: scraper, RestClient: restClient, WsClient: wsClient}
	go bd.mainLoop()

	return bd

}

// runs in a goroutine until s is closed
func (bd *BONDPool) mainLoop() {

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

func (bd *BONDPool) scrapePools() (err error) {
	for _, poolDetail := range bd.getBONDPools() {
		staking, err := staking.BONDStakingCaller(common.HexToAddress(poolDetail.VaultAddress), bd.RestClient)
		if err != nil {
			return err
		}

		header, err := bd.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err

		}

		getCurrentEpoch, err := staking.GetCurrentEpoch()
		if err != nil {
			return err
		}

		var priceStablePool
		
		if getCurrentEpoch >= 7 && getCurrentEpoch <= 25 {	
			priceStablePool, err := staking.GetEpochPoolSize(daiTokenAddress,getCurrentEpoch) + staking.GetEpochPoolSize(usdTokenAddress,getCurrentEpoch) + staking.GetEpochPoolSize(susdTokenAddress,getCurrentEpoch)
			if err != nil {
				return err
			}
		}

		if getCurrentEpoch >= 3 && getCurrentEpoch <= 12 {	
			priceStablePool, err := staking.GetEpochPoolSize(bondTokenAddress,getCurrentEpoch) 
			if err != nil {
				return err
			}
		}
		
		if getCurrentEpoch >= 6 && getCurrentEpoch <= 100 {	
			priceStablePool, err := staking.GetEpochPoolSize(uniLPTokenAddress,getCurrentEpoch) 
			if err != nil {
				return err
			}
		}
		
		
		
		bal, err := staking.Balance(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
			return err

		}
		decimals, err := staking.Decimals(&bind.CallOpts{})
		if err != nil {
			return err

		}
		poolBalance, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(bal), new(big.Float).SetFloat64(math.Pow10(int(decimals)))).Float64()

		pricePerFullShare := new(big.Float).SetInt(priceStablePool)
		rate, _ := big.NewFloat(0).Sub(pricePerFullShare.Quo(pricePerFullShare, new(big.Float).SetFloat64(1e18)), big.NewFloat(1)).Float64()

		var pr models.FarmingPool
		pr.TimeStamp = time.Now()
		pr.Rate = rate
		pr.Balance = poolBalance
		pr.ProtocolName = bd.scraper.poolName
		pr.PoolID = poolDetail.PoolID
		pr.OutputAsset = []string{poolDetail.TokenName}
		pr.BlockNumber = header.Number.Int64()
		pr.InputAsset = []string{poolDetail.TokenName}
		bd.scraper.chanPoolInfo <- &pr
	}
	return

}

func (bd *BONDPool) getBONDPools() (pools []*BONDPoolDetail) {

	pools = append(pools, &BONDPoolDetail{TokenName: "USDC/DAI/sUSD", VaultAddress: "0xb0fa2beee3cf36a7ac7e99b885b48538ab364853", PoolID: "USDC/DAI/sUSD"})
	pools = append(pools, &BONDPoolDetail{TokenName: "USDC_BOND_UNI_LP", VaultAddress: "0xb0fa2beee3cf36a7ac7e99b885b48538ab364853", PoolID: "USDC_BOND_UNI_LP"})
	pools = append(pools, &BONDPoolDetail{TokenName: "BOND", VaultAddress: "0xb0fa2beee3cf36a7ac7e99b885b48538ab364853", PoolID: "BOND"})

	return

}

type BONDPoolDetail struct {
	TokenName    string
	VaultAddress string
	PoolID       string
}
