package aave

import (
	"context"
	"math"
	"math/big"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave/bind/aave"
	contract "github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave/bind/aave"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type reserveReader struct {
	ethC *ethclient.Client

	erc20MD ERC20MetadataProvider

	protocol dia.DefiProtocol

	aaveContractAddr common.Address
	aaveContract     *contract.ILendingPool

	log *logrus.Entry

	chMsg chan *dia.DefiProtocolState
}

func newReserveReader(messageCh chan *dia.DefiProtocolState, deps *scraperDeps) (*reserveReader, error) {
	var err error

	s := &reserveReader{
		ethC:             deps.EthClient,
		aaveContractAddr: common.HexToAddress(deps.Protocol.Address),
		log:              deps.Logger.WithField("comp", "aave-reserve-reader"),
		erc20MD:          deps.ERC20MD,
		chMsg:            messageCh,
	}

	s.aaveContract, err = contract.NewILendingPool(s.aaveContractAddr, s.ethC)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *reserveReader) read(ctx context.Context) error {
	callOpts := &bind.CallOpts{Context: ctx}

	reserveList, err := s.aaveContract.GetReservesList(callOpts)
	if err != nil {
		return errors.Wrapf(err, "unable to get the list of reserves: %s", err.Error())
	}

	ethInUSD, err := utils.GetCoinPrice("ETH")
	if err != nil {
		return errors.Wrapf(err, "unable to read ETH price: %s", err.Error())
	}

	totalReserveAmountInUSD := new(big.Float)

	for i, reserve := range reserveList {
		reserveData, err := s.aaveContract.GetReserveData(&bind.CallOpts{Context: ctx}, reserve)
		if err != nil {
			return errors.Wrapf(err, "unable to get reserve date of the address %s: %s", reserve.String(), err.Error())
		}

		erc20md, err := s.erc20MD.ERC20(ctx, reserveData.ATokenAddress)
		if err != nil {
			return errors.Wrapf(err, "unable to get ERC20 metadata of the address %s: %s", reserve.String(), err.Error())
		}

		reserveToken, err := aave.NewIERC20Caller(reserveData.ATokenAddress, s.ethC)
		if err != nil {
			return errors.Wrapf(err, "unable to create contract binding for the aToken %s: %s", erc20md.Symbol(), err.Error())
		}

		reserveTotalSupply, err := reserveToken.TotalSupply(callOpts)
		if err != nil {
			return errors.Wrapf(err, "unable to get total supply of the aToken %s: %s", erc20md.Symbol(), err.Error())
		}

		s.log.Infof("%d/%d: %s ->%s", i, len(reserveList), reserve.String(), erc20md.Symbol())
		reserveTokenPriceInUSD, err := utils.GetCoinPrice(erc20md.Symbol()[1:])
		if err != nil {
			return errors.Wrapf(err, "unable to get price in USD of the aToken %s: %s", erc20md.Symbol(), err.Error())
		}

		totalReserveAmountInUSD = new(big.Float).Add(
			totalReserveAmountInUSD,
			new(big.Float).Mul(
				new(big.Float).Quo(new(big.Float).SetInt(reserveTotalSupply), big.NewFloat(math.Pow10(erc20md.Decimals()))),
				new(big.Float).SetFloat64(reserveTokenPriceInUSD),
			),
		)
	}

	totalReserveAmountInETH := new(big.Float).Quo(
		totalReserveAmountInUSD,
		new(big.Float).SetFloat64(ethInUSD),
	)

	totalLockedInUSD, _ := totalReserveAmountInUSD.Float64()
	totalLockedInETH, _ := totalReserveAmountInETH.Float64()

	s.log.WithFields(logrus.Fields{
		"total_locked_usd": decimal.NewFromFloat(totalLockedInUSD).StringFixed(2),
		"total_locked_eth": decimal.NewFromFloat(totalLockedInETH).StringFixed(2),
	}).Info("total locked reserves are updated")

	return nil
}
