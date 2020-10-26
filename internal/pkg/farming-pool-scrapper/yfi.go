package pool

import (
	"context"
	"errors"
	strategy "github.com/diadata-org/diadata/internal/pkg/farming-pool-scrapper/yficontracts/strategy"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

type YFIPool struct {
	scrapper   *PoolScraper
	RestClient *ethclient.Client
	WsClient   *ethclient.Client
}

func NewYFIPool(scrapper *PoolScraper) *YFIPool {
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	cv := &YFIPool{scrapper: scrapper, RestClient: restClient, WsClient: wsClient}
	go cv.mainLoop()

	return cv

}

// runs in a goroutine until s is closed
func (cv *YFIPool) mainLoop() {

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(12 * time.Hour)

	go func() {
		// Pool rates change per deposit and withdraw
		for {
			cv.scrapPools()

		}
	}()

	// s.cleanup(err)
}

func (cv *YFIPool) scrapPools() (err error) {
	for _, poolDetail := range cv.getYFIPools() {
		strategy, err := strategy.NewStrategyCaller(common.HexToAddress(poolDetail.VaultAddress), cv.RestClient)
		if err != nil {
			continue
		}

		header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
		if err != nil {
			continue
		}


		pricePerFullShareFromContract, err := strategy.GetPricePerFullShare(&bind.CallOpts{BlockNumber: header.Number})
		if err != nil {
			continue
		}
		

		pricePerFullShare := new(big.Float).SetInt(pricePerFullShareFromContract)
		var pr models.PoolRate
		pr.TimeStamp = time.Now()
		pr.Rate = pricePerFullShare.Quo(pricePerFullShare, new(big.Float).SetFloat64(1e18))
		pr.ProtocolName = cv.scrapper.poolName
		pr.PoolID = poolDetail.PoolID
		pr.OutputAsset = poolDetail.TokenName
		pr.BlockNumber = header.Number.Int64()
		pr.InputAsset = []string{poolDetail.TokenName}
		cv.scrapper.chanPoolInfo <- &pr
	}
	return

}
func (cv *YFIPool) UpdateRate() error {
	return errors.New("")
}

func (cv *YFIPool) getYFIPools() (pools []*YFIPoolDetail) {

	pools = append(pools, &YFIPoolDetail{TokenName: "yCRV", VaultAddress: "0x5dbcF33D8c2E976c6b560249878e6F1491Bca25c", PoolID: "curve.fi/y LP"})
	pools = append(pools, &YFIPoolDetail{TokenName: "USDC", VaultAddress: "0x597aD1e0c13Bfe8025993D9e79C69E1c0233522e", PoolID: "yUSDC"})

	return

}

type YFIPoolDetail struct {
	TokenName    string
	VaultAddress string
	PoolID       string
}
