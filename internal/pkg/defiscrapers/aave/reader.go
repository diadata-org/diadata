package aave

import (
	"context"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/defiscrapers/aave/contract"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	// https://docs.aave.com/developers/deployed-contracts/deployed-contracts
	// LendingPoolAddressesProviderRegistry
	addrProviderRegistry = "0x52D306e36E3B6B02c153d0266ff0f85d18BCD413"

	// we use the main market
	marketID = "Aave genesis market"
)

func Read(ctx context.Context, ethClient *ethclient.Client, protocol dia.DefiProtocol, chRate chan *dia.DefiRate, chState chan *dia.DefiProtocolState) error {
	log := logrus.WithFields(logrus.Fields{
		"comp": "aave-scraper",
	})

	blockNum, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return errors.Wrapf(err, "unable to get current block number: %s", err.Error())
	}

	blockHdrs, err := ethClient.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return errors.Wrapf(err, "unable to get block headers for the height %d: %s", blockNum, err.Error())
	}

	date := time.Unix(int64(blockHdrs.Time), 0)

	callOpts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: new(big.Int).SetUint64(blockNum),
	}

	addrProvider, err := getAddrProvider(callOpts, ethClient)
	if err != nil {
		return errors.WithStack(err)
	}

	lendingPool, err := lendingPoolContract(callOpts, ethClient, addrProvider)
	if err != nil {
		return errors.WithStack(err)
	}

	reserveAddrs, err := lendingPool.GetReservesList(callOpts)
	if err != nil {
		return errors.Wrapf(err, "unable to get the list of reserve addresses from the lending pool contract: %s", err.Error())
	}

	ethPriceInUSD, err := utils.GetCoinPrice("ETH")
	if err != nil {
		return errors.Wrapf(err, "unable to find the price in USD for ETH: %s", err.Error())
	}

	totalLockedInUSD := new(big.Float)

	for _, reserveAddr := range reserveAddrs {
		reserveData, err := lendingPool.GetReserveData(callOpts, reserveAddr)
		if err != nil {
			return errors.Wrapf(err, "unable to get reserve data for the address: %s: %s", reserveAddr.String(), err.Error())
		}

		underlyingAsset, err := contract.NewERC20(reserveAddr, ethClient)
		if err != nil {
			return errors.Wrapf(err, "unable to initiate ERC20 contract for underlying asset at %s: %s", reserveAddr.String(), err.Error())
		}

		underlyingAssetSymbol, err := underlyingAsset.Symbol(callOpts)
		if err != nil {
			// MKR contract has a bug on the Symbol calls
			if strings.Contains(err.Error(), "cannot marshal in to go slice") && reserveAddr.String() == "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2" {
				underlyingAssetSymbol = "MKR"

			} else {
				return errors.Wrapf(err, "unable to read symbol metadata of the underlying asset at %s: %s", reserveAddr.String(), err.Error())
			}
		}

		priceInUSD, err := utils.GetCoinPrice(underlyingAssetSymbol)
		if err != nil {
			return errors.Wrapf(err, "unable to find the price in USD for underlying asset %s: %s", underlyingAssetSymbol, err.Error())
		}

		aToken, err := contract.NewERC20(reserveData.ATokenAddress, ethClient)
		if err != nil {
			return errors.Wrapf(err, "unable to initiate ERC20 contract for aToken at %s: %s", reserveData.ATokenAddress.String(), err.Error())
		}

		aTokenDecimals, err := aToken.Decimals(callOpts)
		if err != nil {
			return errors.Wrapf(err, "unable to read decimals value for aToken at %s: %s", reserveData.ATokenAddress.String(), err.Error())
		}

		totalLockedAsset, err := aToken.TotalSupply(callOpts)
		if err != nil {
			return errors.Wrapf(err, "unable to read total supply for aToken at %s: %s", reserveData.ATokenAddress.String(), err.Error())
		}

		totalLockedAssetAmount := new(big.Float).Quo(
			new(big.Float).SetInt(totalLockedAsset),
			big.NewFloat(math.Pow10(int(aTokenDecimals))),
		)

		totalLockedAssetInUSD := new(big.Float).Mul(
			big.NewFloat(priceInUSD),
			totalLockedAssetAmount,
		)

		liquidityRate := new(big.Float).Quo(
			new(big.Float).SetInt(reserveData.CurrentLiquidityRate),
			big.NewFloat(math.Pow10(25)),
		)

		stableBorrowRate := new(big.Float).Quo(
			new(big.Float).SetInt(reserveData.CurrentStableBorrowRate),
			big.NewFloat(math.Pow10(25)),
		)

		log.WithFields(logrus.Fields{
			"asset":          underlyingAssetSymbol,
			"locked":         totalLockedAssetAmount.Text('f', 2),
			"locked_in_usd":  totalLockedAssetInUSD.Text('f', 2),
			"lending_rate":   liquidityRate.Text('f', 2),
			"borrowing_rate": stableBorrowRate.Text('f', 2),
		}).Debug()

		totalLockedInUSD = new(big.Float).Add(
			totalLockedInUSD,
			totalLockedAssetInUSD,
		)

		liquidityRate64, _ := liquidityRate.Float64()
		stableBorrowRate64, _ := stableBorrowRate.Float64()

		if chRate != nil {
			select {
			case <-ctx.Done():
				return errors.WithStack(ctx.Err())

			case chRate <- &dia.DefiRate{
				Timestamp:     date,
				Protocol:      protocol.Name,
				Asset:         underlyingAssetSymbol,
				LendingRate:   liquidityRate64,
				BorrowingRate: stableBorrowRate64,
			}:
			}
		}
	}

	totalLockedInETH := new(big.Float).Quo(
		totalLockedInUSD,
		big.NewFloat(ethPriceInUSD),
	)

	log.WithFields(logrus.Fields{
		"total_locked_in_usd": totalLockedInUSD.Text('f', 2),
		"total_locked_in_eth": totalLockedInETH.Text('f', 2),
	}).Debug()

	totalLockedInUSD64, _ := totalLockedInUSD.Float64()
	totalLockedInETH64, _ := totalLockedInETH.Float64()

	if chState != nil {
		select {
		case <-ctx.Done():
			return errors.WithStack(ctx.Err())

		case chState <- &dia.DefiProtocolState{
			Timestamp: date,
			Protocol:  protocol,
			TotalUSD:  totalLockedInUSD64,
			TotalETH:  totalLockedInETH64,
		}:
		}
	}

	return nil
}

func lendingPoolContract(callOpts *bind.CallOpts, backend bind.ContractBackend, addrProvider *contract.ILendingPoolAddressesProvider) (*contract.LendingPool, error) {
	lendingPoolAddr, err := addrProvider.GetLendingPool(callOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get lending pool address: %s", err.Error())
	}

	lendingPool, err := contract.NewLendingPool(lendingPoolAddr, backend)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to initiate lending pool contract at %s: %s", lendingPoolAddr.String(), err.Error())
	}

	return lendingPool, nil
}

// getAddrProvider returns the binding of address provider contract for the market specified by marketID
func getAddrProvider(callOpts *bind.CallOpts, backend bind.ContractBackend) (*contract.ILendingPoolAddressesProvider, error) {
	addProviderReg, err := contract.NewILendingPoolAddressesProviderRegistry(
		common.HexToAddress(addrProviderRegistry),
		backend,
	)

	if err != nil {
		return nil, errors.Wrapf(err, "unable to bind the contract LendingPoolAddressesProviderRegistry at %s: %s", addrProviderRegistry, err)
	}

	addrProviders, err := addProviderReg.GetAddressesProvidersList(callOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get the list of address providers: %s", err.Error())
	}

	for _, addrProviderAddr := range addrProviders {
		addrProvider, err := contract.NewILendingPoolAddressesProvider(addrProviderAddr, backend)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to bind the contract LendingPoolAddressesProvider at %s: %s", addrProviderAddr, err)
		}

		mid, err := addrProvider.GetMarketId(callOpts)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get market ID from the LendingPoolAddressesProvider at %s: %s", addrProviderAddr, err)
		}

		if mid == marketID {
			return addrProvider, nil
		}
	}

	return nil, errors.Errorf("unable to find the market at %s: %s", addrProviderRegistry, marketID)
}
