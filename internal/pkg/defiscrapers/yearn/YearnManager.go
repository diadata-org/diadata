package yearn

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	aprOracle *Aproracle
}

func NewYearnManager(aprOracle *Aproracle) *YearnManager{
	return &YearnManager{
		aprOracle: aprOracle,
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

func (y *YearnManager) Publish(protocolName, symbol string, value *big.Int, defiRateChan chan *dia.DefiRate)  {
	rateValue, _ := new(big.Float).SetInt(value).Float64()
	rate := &dia.DefiRate{
		Timestamp:     time.Now(),
		Asset:         symbol,
		Protocol:      protocolName,
		LendingRate:   rateValue,
	}

	// Send new data through channel defiRateChan
	log.Printf("Write defiRate %#v in %v\n", rate, defiRateChan)
	defiRateChan <- rate
	log.Printf("Update of  %v completed\n", symbol)
}
