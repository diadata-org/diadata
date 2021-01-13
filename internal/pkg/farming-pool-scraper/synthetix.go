package pool

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	proxy "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/synthetix/proxy"
	synthetixcontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/synthetix/synthetix"
	models "github.com/diadata-org/diadata/pkg/model"
)

var (
	// synthetixTokenAddress = "0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F"
	synthetixBaseAddress = "0xf87A0587Fe48Ca05dd68a514Ce387C0d4d3AE31C"
)

type SynthetixScraper struct {
	scraper    *PoolScraper
	restClient *ethclient.Client
	wsClient   *ethclient.Client
}

func NewSynthetixScraper(scraper *PoolScraper) *SynthetixScraper {

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
		scraper:    scraper,
	}

	go sts.mainLoop()
	return sts
}

// runs in a goroutine until sts is closed
func (sts *SynthetixScraper) mainLoop() {
	go func() {
		for {
			select {
			case <-sts.scraper.tickerRate.C:
				err := sts.scrapePools()
				if err != nil {
					log.Errorln("Error while Scraping", err)
				}
			}

		}
	}()
}

func (sts *SynthetixScraper) scrapePools() error {
	// Mechanism of synthetic farming in DIA Methodology and in the synthetix litepaper
	// https://www.synthetix.io/uploads/synthetix_litepaper.pdf
	// https://synthetix.community/docs/claiming-rewards
	// https://docs.synthetix.io/addresses/

	sntxBase, err := synthetixcontract.NewISynthetixCaller(common.HexToAddress(synthetixBaseAddress), sts.restClient)
	if err != nil {
		log.Info(err)
		return err
	}

	// Get total system's debt
	rem, err := sntxBase.RemainingIssuableSynths(&bind.CallOpts{}, common.Address{})
	if err != nil {
		log.Info(err)
		return err
	}
	totalSystemDebt, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(rem.TotalSystemDebt), new(big.Float).SetFloat64(math.Pow10(18))).Float64()

	// Get block number
	header, err := sts.restClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.WithError(err).Error("getting current heading block")
		return err
	}

	// Iterate over all synths
	numSynths, err := sntxBase.AvailableSynthCount(&bind.CallOpts{})
	if err != nil {
		log.Info(err)
		return err
	}
	for i := 0; i < int(numSynths.Int64()); i++ {
		synthAddress, err := sntxBase.AvailableSynths(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			log.Info(err)
			return err
		}
		token, err := proxy.NewProxyERC20Caller(synthAddress, sts.restClient)
		if err != nil {
			log.WithError(err).Error("loading SNX contract caller")
			return err
		}
		symbol, err := token.Symbol(&bind.CallOpts{})
		if err != nil {
			return err
		}
		decimals, err := token.Decimals(&bind.CallOpts{})
		// Supply values for synths from "Total Synths" on https://stats.synthetix.io/#staking
		totSupp, _ := token.TotalSupply(&bind.CallOpts{})
		totalSupply, _ := new(big.Float).Quo(big.NewFloat(0).SetInt(totSupp), new(big.Float).SetFloat64(math.Pow10(int(decimals)))).Float64()

		var pr models.FarmingPool
		pr.TimeStamp = time.Now()
		pr.Rate = totalSystemDebt
		pr.Balance = totalSupply
		pr.ProtocolName = "SYNTHETIX"
		pr.PoolID = synthAddress.Hex()
		pr.OutputAsset = []string{symbol}
		pr.BlockNumber = header.Number.Int64()
		pr.InputAsset = []string{"SNX"}
		sts.scraper.chanPoolInfo <- &pr

	}
	return nil
}
