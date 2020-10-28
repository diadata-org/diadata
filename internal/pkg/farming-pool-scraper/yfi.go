package pool

import (
	"context"
	"math"
	"math/big"
	"time"

	strategy "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/yficontracts/strategy"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type YFIPool struct {
	scraper    *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

func NewYFIPool(scraper *PoolScraper) *YFIPool {
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	cv := &YFIPool{scraper: scraper, RestClient: restClient, WsClient: wsClient}
	go cv.mainLoop()

	return cv

}

// runs in a goroutine until s is closed
func (cv *YFIPool) mainLoop() {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		// Pool rates change per deposit and withdraw
		for {
			select {
			case <-ticker.C:
				err := cv.scrapePools()
				if err != nil {
					log.Errorln("Error while Scrapping", err)
				}
			}

		}
	}()

	// s.cleanup(err)
}

func (cv *YFIPool) scrapePools() (err error) {
	for _, poolDetail := range cv.getYFIPools() {
		strategy, err := strategy.NewStrategyCaller(common.HexToAddress(poolDetail.VaultAddress), cv.RestClient)
		if err != nil {
			return err
		}

		header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return err

		}

		pricePerFullShareFromContract, err := strategy.GetPricePerFullShare(&bind.CallOpts{BlockNumber: header.Number})
		if err != nil {
			return err

		}
		bal, err := strategy.Balance(&bind.CallOpts{})
		if err != nil {
			log.Error(err)
			return err

		}
		decimals, err := strategy.Decimals(&bind.CallOpts{})
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

func (cv *YFIPool) getYFIPools() (pools []*YFIPoolDetail) {

	pools = append(pools, &YFIPoolDetail{TokenName: "yCRV", VaultAddress: "0x5dbcF33D8c2E976c6b560249878e6F1491Bca25c", PoolID: "curve.fi/y LP"})
	pools = append(pools, &YFIPoolDetail{TokenName: "USDC", VaultAddress: "0x597aD1e0c13Bfe8025993D9e79C69E1c0233522e", PoolID: "yUSDC"})
	pools = append(pools, &YFIPoolDetail{TokenName: "TUSD", VaultAddress: "0x37d19d1c4E1fa9DC47bD1eA12f742a0887eDa74a", PoolID: "TUSD"})
	pools = append(pools, &YFIPoolDetail{TokenName: "DAI", VaultAddress: "0xACd43E627e64355f1861cEC6d3a6688B31a6F952", PoolID: "DAI"})
	pools = append(pools, &YFIPoolDetail{TokenName: "USDT", VaultAddress: "0x2f08119C6f07c006695E079AAFc638b8789FAf18", PoolID: "USDT"})
	pools = append(pools, &YFIPoolDetail{TokenName: "YFI", VaultAddress: "0xBA2E7Fed597fd0E3e70f5130BcDbbFE06bB94fe1", PoolID: "YFI"})
	pools = append(pools, &YFIPoolDetail{TokenName: "crvBUSD", VaultAddress: "0x2994529C0652D127b7842094103715ec5299bBed", PoolID: "crvBUSD"})
	pools = append(pools, &YFIPoolDetail{TokenName: "crvBTC", VaultAddress: "0x7Ff566E1d69DEfF32a7b244aE7276b9f90e9D0f6", PoolID: "crvBTC"})
	pools = append(pools, &YFIPoolDetail{TokenName: "WETH", VaultAddress: "0xe1237aA7f535b0CC33Fd973D66cBf830354D16c7", PoolID: "WETH"})
	pools = append(pools, &YFIPoolDetail{TokenName: "3crv", VaultAddress: "0x9cA85572E6A3EbF24dEDd195623F188735A5179f", PoolID: "3crv"})
	pools = append(pools, &YFIPoolDetail{TokenName: "GUSD", VaultAddress: "0xec0d8D3ED5477106c6D4ea27D90a60e594693C90", PoolID: "GUSD"})

	return

}

type YFIPoolDetail struct {
	TokenName    string
	VaultAddress string
	PoolID       string
}
