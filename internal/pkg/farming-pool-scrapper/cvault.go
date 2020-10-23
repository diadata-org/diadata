package pool

import (
	"errors"
	cvaultcontract "github.com/diadata-org/diadata/internal/pkg/farming-pool-scrapper/cvault"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	logrus "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

var log = logrus.New()

const (
	restDial      = "https://mainnet.infura.io/v3/2883d1b22e0e4d62b535592dd8075fee"
	wsDial        = "wss://mainnet.infura.io/ws/v3/2883d1b22e0e4d62b535592dd8075fee"
	cvaultAddress = "0xc5cacb708425961594b63ec171f4df27a9c0d8c9"
)

type Cvault struct {
	scrapper      *PoolScraper
	RestClient    *ethclient.Client
	WsClient      *ethclient.Client
	DepositEvent  chan *cvaultcontract.CvaultpoolcontractDeposit
	WithDrawEvent chan *cvaultcontract.CvaultpoolcontractWithdraw
}

func NewCvaultScrapper(scrapper *PoolScraper) *Cvault {

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
	cv := &Cvault{scrapper: scrapper, RestClient: restClient, WsClient: wsClient, DepositEvent: deposit, WithDrawEvent: withdrwa}
	go cv.mainLoop()

	return cv

}

// runs in a goroutine until s is closed
func (cv *Cvault) mainLoop() {

	// wait for all pairs have added into s.PairScrapers
	time.Sleep(4 * time.Second)

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

func (cv *Cvault) getPool(poolID *big.Int) {
	cvg, err := cvaultcontract.NewCvaultpoolcontractCaller(common.HexToAddress(cvaultAddress), cv.RestClient)
	if err!=nil{
		log.Errorln("Error initialising cvault caller ",err)
	}
	pi, err := cvg.PoolInfo(&bind.CallOpts{}, poolID)
	if err!=nil{
		log.Errorln("Error getting poolInfo ",err)
	}

	AccCorePerShareFloat := new(big.Float).SetInt(pi.AccCorePerShare)

	var pr models.PoolRate
	pr.TimeStamp = time.Now()
	pr.Rate = AccCorePerShareFloat.Quo(AccCorePerShareFloat, new(big.Float).SetFloat64(1e12))
	pr.ProtocolName = cv.scrapper.poolName
	pr.PoolID = poolID.String()
	pr.OutputAsset = "CORE"
	cv.scrapper.chanPoolInfo <- &pr

}
func (cv *Cvault) UpdateRate() error {
	return errors.New("")
}
