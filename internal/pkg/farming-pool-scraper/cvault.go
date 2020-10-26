package pool

import (
	"errors"

	cvaultcontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/cvault"
	supplyservice "github.com/diadata-org/diadata/internal/pkg/supplyService"

	erc "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/cvault/erc"
	lptoken "github.com/diadata-org/diadata/internal/pkg/farming-pool-scraper/cvault/lptoken"

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
	restDial       = "https://mainnet.infura.io/v3/2883d1b22e0e4d62b535592dd8075fee"
	wsDial         = "wss://mainnet.infura.io/ws/v3/2883d1b22e0e4d62b535592dd8075fee"
	cvaultAddress  = "0xc5cacb708425961594b63ec171f4df27a9c0d8c9"
	lpTokenAddress = "0x32Ce7e48debdccbFE0CD037Cc89526E4382cb81b"
)

type Cvault struct {
	scraper       *PoolScraper
	RestClient    *ethclient.Client
	WsClient      *ethclient.Client
	DepositEvent  chan *cvaultcontract.CvaultpoolcontractDeposit
	WithDrawEvent chan *cvaultcontract.CvaultpoolcontractWithdraw
}

func NewCvaultScraper(scraper *PoolScraper) *Cvault {

	deposit := make(chan *cvaultcontract.CvaultpoolcontractDeposit)
	withdrwa := make(chan *cvaultcontract.CvaultpoolcontractWithdraw)
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

	fr, _ := cvaultcontract.NewCvaultpoolcontractFilterer(common.HexToAddress(cvaultAddress), cv.WsClient)
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
					cv.getPool(deposit.Pid)
				}
			case withdraw := <-cv.WithDrawEvent:
				{
					cv.getPool(withdraw.Pid)
				}
			}

		}
	}()

	// s.cleanup(err)
}

func (cv *Cvault) getPool(poolID *big.Int) (err error) {
	//TODO call all pool details at once and save in cache
	log.Infoln("Getting Pool Info")
	cvg, err := cvaultcontract.NewCvaultpoolcontractCaller(common.HexToAddress(cvaultAddress), cv.RestClient)
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
	balLPToken, err := supplyservice.GetWalletBalance(cvaultAddress, lpTokenAddress, cv.RestClient)

	AccCorePerShareFloat := new(big.Float).SetInt(pi.AccCorePerShare)
	rate, _ := AccCorePerShareFloat.Quo(AccCorePerShareFloat, new(big.Float).SetFloat64(1e12)).Float64()
	var pr models.FarmingPool
	pr.TimeStamp = time.Now()
	pr.Rate = rate
	pr.Balance = balLPToken
	pr.ProtocolName = cv.scraper.poolName
	pr.PoolID = poolID.String()
	pr.InputAsset = []string{token1Symbol}
	pr.OutputAsset = token0Symbol
	cv.scraper.chanPoolInfo <- &pr
	log.Infoln(pr)
	return

}
func (cv *Cvault) UpdateRate() error {
	return errors.New("")
}
