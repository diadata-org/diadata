package yearn

import (
	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/yearn"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

type CompoundAPRResult struct {
	CDAI  *big.Int
	CBAT  *big.Int
	CETH  *big.Int
	CREP  *big.Int
	CSAI  *big.Int
	CUSDC *big.Int
	CWBTC *big.Int
	CZRC  *big.Int
}

type YearnManager struct{
	rpcUrl string
	aprOracleAddress string
	client *ethclient.Client
	aprOracle *yearn.Aproracle
}

func NewYearnManager(rpcUrl string, aprOracleAddress string) *YearnManager{
	return &YearnManager{
		rpcUrl: rpcUrl,
		aprOracleAddress: aprOracleAddress,
	}
}

func (y *YearnManager) Init(){
	client, err := ethclient.Dial(y.rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	y.client = client
	y.aprOracle, err = yearn.NewAproracle(common.HexToAddress(y.aprOracleAddress), y.client)
	if err != nil {
		log.Fatalf("Failed to bind APR oracle contract: %v", err)
	}
}

func (y *YearnManager) GetAllCompoundAPR() (*CompoundAPRResult, error){
	aprs, err := y.aprOracle.GetAllCompoundAPR(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to GetAllCompoundAPR from APR oracle contract: %v", err)
		return nil, err
	}
	return &CompoundAPRResult{
		CDAI:  aprs.CDAI,
		CBAT:  aprs.CBAT,
		CETH:  aprs.CETH,
		CREP:  aprs.CREP,
		CSAI:  aprs.CSAI,
		CUSDC: aprs.CUSDC,
		CWBTC: aprs.CWBTC,
		CZRC:  aprs.CZRC,
	}, nil
}

func (y *YearnManager) Publish(source, symbol string, value *big.Int, interestRateChan chan *models.InterestRate)  {
	rateValue, _ := new(big.Float).SetInt(value).Float64()
	now := time.Now()
	rate := &models.InterestRate{
		Symbol:          symbol,
		Value:           rateValue,
		PublicationTime: now,
		EffectiveDate:   now,
		Source:          source,
	}

	// Send new data through channel chanInterestRate
	log.Printf("Write interestRate %#v in %v\n", rate, interestRateChan)
	interestRateChan <- rate
	log.Printf("Update of  %v completed\n", symbol)

}
