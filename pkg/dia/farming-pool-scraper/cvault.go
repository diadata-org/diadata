package pool

import (
	"context"
	"github.com/diadata-org/diadata/internal/pkg/supplyService"
	"github.com/diadata-org/diadata/pkg/dia/farming-pool-scraper/cvault"
	"github.com/diadata-org/diadata/pkg/dia/farming-pool-scraper/cvault/erc"
	"github.com/diadata-org/diadata/pkg/dia/farming-pool-scraper/cvault/lptoken"

	"math/big"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	logrus "github.com/sirupsen/logrus"
)

var log = logrus.New()

const (
	cvaultAddress  = "0xc5cacb708425961594b63ec171f4df27a9c0d8c9"
	lptokenAddress = "0x32Ce7e48debdccbFE0CD037Cc89526E4382cb81b" //nolint:gosec
)

type Cvault struct {
	scraper       *PoolScraper
	RestClient    *ethclient.Client
	WsClient      *ethclient.Client
	DepositEvent  chan *cvaultpoolcontract.cvaultpoolcontract
	WithDrawEvent chan *cvaultpoolcontract.CvaultpoolcontractWithdraw
}

func NewCvaultScraper(scraper *PoolScraper) *Cvault {

	deposit := make(chan *cvaultpoolcontract.CvaultpoolcontractDeposit)
	withdrwa := make(chan *cvaultpoolcontract.CvaultpoolcontractWithdraw)
	restClient, err := ethclient.Dial(restDial)
	if err != nil {
		log.Fatal(err)
	}
	wsClient, err := ethclient.Dial(wsDial)
	if err != nil {
		log.Fatal(err)
	}
	cv := &Cvault{scraper: scraper, RestClient: restClient, WsClient: wsClient, DepositEvent: deposit, WithDrawEvent: withdrwa}
	go cv.mainLoop()

	return cv

}

// runs in a goroutine until s is closed
func (cv *Cvault) mainLoop() {

	fr, _ := cvaultpoolcontract.NewCvaultpoolcontractFilterer(common.HexToAddress(cvaultAddress), cv.WsClient)
	_, err := fr.WatchDeposit(&bind.WatchOpts{}, cv.DepositEvent, nil, nil)
	if err != nil {
		log.Errorln("Error Subscribing deposit events", err)
	}
	log.Info("Subscribed to Deposit events")

	_, err = fr.WatchWithdraw(&bind.WatchOpts{}, cv.WithDrawEvent, nil, nil)
	if err != nil {
		log.Errorln("Error Subscribing WithDraw events", err)
	}
	log.Info("Subscribed to WithDraw events")

	go func() {
		// Pool rates change per deposit and withdraw
		for {
			select {
			case deposit := <-cv.DepositEvent:
				{
					err := cv.getPool(deposit.Pid)
					if err != nil {
						log.Error(err)
					}
				}
			case withdraw := <-cv.WithDrawEvent:
				{
					err := cv.getPool(withdraw.Pid)
					if err != nil {
						log.Error(err)
					}
				}
			}

		}
	}()

	// s.cleanup(err)
}

// https://help.cvault.finance/core-formulas/formulas
func (cv *Cvault) getPool(poolID *big.Int) (err error) {
	//TODO call all pool details at once and save in cache
	log.Infoln("Getting Pool Info")
	cvg, err := cvaultpoolcontract.NewCvaultpoolcontractCaller(common.HexToAddress(cvaultAddress), cv.RestClient)
	if err != nil {
		return
	}
	header, err := cv.RestClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return
	}

	pi, err := cvg.PoolInfo(&bind.CallOpts{}, poolID)
	if err != nil {
		return
	}

	lptokenDetails, err := lptoken.NewLptokenCaller(pi.Token, cv.RestClient)
	if err != nil {
		return
	}
	token1, err := lptokenDetails.Token1(&bind.CallOpts{})
	if err != nil {
		return
	}
	token0, err := lptokenDetails.Token0(&bind.CallOpts{})
	if err != nil {
		return
	}

	token1Details, err := erc.NewCvaultcontractCaller(token1, cv.RestClient)
	if err != nil {
		return
	}
	token0Details, err := erc.NewCvaultcontractCaller(token0, cv.RestClient)
	if err != nil {
		return
	}

	token0Symbol, err := token0Details.Symbol(&bind.CallOpts{})
	if err != nil {
		return
	}
	token1Symbol, err := token1Details.Symbol(&bind.CallOpts{})
	if err != nil {
		return
	}
	balLPToken, err := supplyservice.GetWalletBalance(cvaultAddress, lptokenAddress, cv.RestClient)

	AccCorePerShareFloat := new(big.Float).SetInt(pi.AccCorePerShare)
	var pr models.FarmingPool
	if int(poolID.Int64()) == 0 {
		rate, _ := AccCorePerShareFloat.Quo(AccCorePerShareFloat, new(big.Float).SetFloat64(1e12)).Float64()
		pr.Rate = rate
	}
	if int(poolID.Int64()) == 1 {
		// cBTC pool scales in centimillis (1e-5)
		rate, _ := AccCorePerShareFloat.Quo(AccCorePerShareFloat, new(big.Float).SetFloat64(1e17)).Float64()
		pr.Rate = rate
	}
	pr.TimeStamp = time.Now()
	pr.Balance = balLPToken
	pr.ProtocolName = cv.scraper.poolName
	pr.PoolID = poolID.String()
	pr.BlockNumber = header.Number.Int64()
	pr.InputAsset = []string{token1Symbol}
	pr.OutputAsset = []string{token0Symbol}
	cv.scraper.chanPoolInfo <- &pr
	return

}
