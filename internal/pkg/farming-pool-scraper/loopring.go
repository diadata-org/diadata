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

	lrctokencontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/loopring/token/LRCTokenContract"
)

type LRCPool struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

func NewLRCPoolScrapper(scraper *PoolScraper) *YFIPool {
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
					log.Errorln("Error while Scrapping", err)
				}
			}

		}
	}()

}

func (cv *LRCPool) scrapePools() (err error) {
	for _, poolDetail := range cv.getLRCPools() {
		lc, err := lrctokencontract.NewLRCTokenContractCaller(common.HexToAddress(poolDetail.VaultAddress), cv.RestClient)
		if err != nil {
			return err
		}

		header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err

		}

		pricePerFullShareFromContract, err := lrctokencontract.GetPricePerFullShare(&bind.CallOpts{BlockNumber: header.Number})
		if err != nil {
			return err

		}
		bal, err := lrctokencontract.Balance(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
			return err

		}
		decimals, err := lrctokencontract.Decimals(&bind.CallOpts{})
		if err != nil {
			return err

		}
		poolBalance, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(bal), new(big.Float).SetFloat64(math.Pow10(int(decimals)))).Float64()

		pricePerFullShare := new(big.Float).SetInt(pricePerFullShareFromContract)
		rate, _ := big.NewFloat(0).Sub(pricePerFullShare.Quo(pricePerFullShare, new(big.Float).SetFloat64(1e18)), big.NewFloat(1)).Float64()

		var pr models.FarmingPool
		pr.TimeStamp = time.Now()
		pr.Rate = rate
		pr.Balance = poolBalance
		pr.ProtocolName = cv.scraper.poolName
		pr.PoolID = poolDetail.PoolID
		pr.OutputAsset = []string{poolDetail.TokenName}
		pr.BlockNumber = header.Number.Int64()
		pr.InputAsset = []string{poolDetail.TokenName}
		cv.scraper.chanPoolInfo <- &pr
	}
	return

}

func (cv *LRCPool) getLRCPools() (pools []*LRCPoolDetail) {
	pools = append(pools, &LRCPoolDetail{TokenName: "LRC", VaultAddress: "0xbbbbca6a901c926f240b89eacb641d8aec7aeafd", PoolID: "LRC"})
	pools = append(pools, &LRCPoolDetail{TokenName: "fLRC", VaultAddress: "0x4b89f8996892d137c3dE1312d1dD4E4F4fFcA171", PoolID: "feeLRC"})
	pools = append(pools, &LRCPoolDetail{TokenName: "LRCdex", VaultAddress: "0x944644Ea989Ec64c2Ab9eF341D383cEf586A5777", PoolID: "LRCdex"})
	pools = append(pools, &LRCPoolDetail{TokenName: "stakingLRC", VaultAddress: "0xF4662bB1C4831fD411a95b8050B3A5998d8A4A5b", PoolID: "stakingLRC"})

	return

}

type LRCPoolDetail struct {
	TokenName    string
	VaultAddress string
	PoolID       string
}
