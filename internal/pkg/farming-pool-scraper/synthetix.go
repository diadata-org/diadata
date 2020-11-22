package pool

import (
	"context"
	"math/big"
	"time"

	"github.com/bep/debounce"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	proxy "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/synthetix/proxy"
	models "github.com/diadata-org/diadata/pkg/model"
	// proxy "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/synthetix/proxy"
)

var (
	synthetixTokenAddress = common.HexToAddress("0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F")
)

type SynthetixScraper struct {
	scraper *PoolScraper

	restClient *ethclient.Client
	wsClient   *ethclient.Client
}

func NewSynthetixScraper(scraper *PoolScraper) *SynthetixScraper {
	// create rest and ws eth clients
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}

	sts := &SynthetixScraper{
		restClient: restClient,
		wsClient:   wsClient,

		scraper: scraper,
	}

	go sts.mainLoop()

	return sts
}

func (sts *SynthetixScraper) scrape() {
	token, err := proxy.NewProxyERC20Caller(synthetixTokenAddress, sts.restClient)
	if err != nil {
		log.WithError(err).Error("loading SNX contract caller")
		return
	}

	header, err := sts.restClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.WithError(err).Error("getting current heading block")
		return
	}

	balance, err := token.TotalSupply(nil)
	if err != nil {
		log.WithError(err).Error("getting SNX supply")
		return
	}
	balanceFloat := new(big.Float).SetInt(balance)
	balanceFloat = balanceFloat.Quo(balanceFloat, big.NewFloat(10e17))
	balanceFloatFloat, _ := balanceFloat.Float64()

	sts.scraper.chanPoolInfo <- &models.FarmingPool{
		Balance: balanceFloatFloat,

		Rate: balanceFloatFloat,

		TimeStamp:    time.Now(),
		PoolID:       synthetixTokenAddress.Hex(),
		ProtocolName: sts.scraper.poolName,

		InputAsset:  []string{"SNX"},
		OutputAsset: []string{"sUSD"},

		BlockNumber: header.Number.Int64(),
	}
}

func (sts *SynthetixScraper) mainLoop() {
	// in mainLoop, we watch for transactions mentionning the SNX token
	// upon transaction, we scrape informations about snx
	// scraping is debounced by 30 seconds

	headersChan := make(chan *types.Header)
	sus, err := sts.wsClient.SubscribeNewHead(context.Background(), headersChan)
	if err != nil {
		log.WithError(err).Error("error subscribing to blockchain head")
		return
	}

	debouncedScrape := debounce.New(30 * time.Second)

	for {
		select {
		case err := <-sus.Err():
			log.WithError(err).Error("blockchain subscription error")
		case header := <-headersChan:
			block, err := sts.restClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.WithError(err).Error("error fetching block by hash")
				continue
			}

			for _, tx := range block.Transactions() {
				if tx == nil {
					continue
				} else if tx.To() == nil {
					continue
				}

				if tx.To().Hex() == synthetixTokenAddress.Hex() {
					log.WithFields(logrus.Fields{}).Info("new transaction involving SNX token")
					go debouncedScrape(sts.scrape)
				}
			}
		}
	}
}
