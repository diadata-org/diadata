package diaoracle

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/diaPcwsOracleService"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type UnSignedOracle struct {
	conn              *ethclient.Client
	auth              *bind.TransactOpts
	contract          *diaPcwsOracleService.DIAPcwsOracle
	deviationPermille int
	updateOnDeviation bool
	oracleAddress     string
	symbols           []string
	oldPrices         map[string]float64
	indexSymbols      []string
}

func NewUnSignedOracleContract(conn *ethclient.Client, auth *bind.TransactOpts, oracleAddress string, deviationPermille int, updateOnDeviation bool, symbols, indexSymbols []string) *UnSignedOracle {
	oldPrices := make(map[string]float64)
	lo := &UnSignedOracle{conn: conn, auth: auth, oracleAddress: oracleAddress, symbols: symbols, oldPrices: oldPrices, deviationPermille: deviationPermille, updateOnDeviation: updateOnDeviation, indexSymbols: indexSymbols}
	return lo
}

func (lo *UnSignedOracle) Update() error {
	err := lo.deployOrBindContract(lo.oracleAddress)
	if err != nil {
		return err
	}

	if lo.oracleAddress != "" {

		// Update Index Values

		for _, indexSymbol := range lo.indexSymbols {
			rawIndex, err := getIndexValueFromDia(indexSymbol)
			if err != nil {
				log.Fatalf("Failed to retrieve crypto index data from DIA: %v", err)
				return err
			}
			err = lo.updateIndexValue(rawIndex)
			if err != nil {
				log.Fatalf("Failed to update Scifi index Oracle: %v", err)
				return err
			}
		}

		// Update  Prices for Symbols
		for _, symbol := range lo.symbols {

			var (
				symbolName string
				price      float64
				timestamp  int64
			)
			if strings.Contains(symbol, "/") {
				quoteQuotation, baseQuotation, err := lo.getPairQuotation(symbol)
				if err != nil {
					continue
				}
				symbolName = quoteQuotation.Symbol + "/" + baseQuotation.Symbol
				price = quoteQuotation.Price / baseQuotation.Price
				timestamp = time.Now().Unix()

				// lo.updatePair(symbol0, symbol1)

			} else {

				// Get Quotation and update if price flucs
				quotation, err := lo.getQuotation(symbol)
				if err != nil {
					continue
				}

				symbolName = quotation.Symbol
				price = quotation.Price
				timestamp = time.Now().Unix()

			}

			newPrice := price
			oldPrice := lo.oldPrices[symbolName]

			if !lo.updateOnDeviation {

				err := lo.updateOracle(symbolName, int64(price*100000), timestamp)
				lo.oldPrices[symbol] = price
				if err != nil {
					log.Fatalf("Failed to update Oracle: %v", err)
					return err
				}

			}

			if lo.updateOnDeviation {

				if (newPrice > (oldPrice * (1 + float64(lo.deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(lo.deviationPermille)/1000))) {
					err := lo.updateOracle(symbolName, int64(price*100000), timestamp)
					lo.oldPrices[symbol] = price
					if err != nil {
						log.Fatalf("Failed to update Oracle: %v", err)
						return err
					}

				} else {
					log.Infof("last Price of symbol %s is %f and new is  %f", symbolName, lo.oldPrices[symbol], newPrice)

					log.Infoln("fluctation not yet happened")
				}
			}

		}

	}
	return nil
}

func (lo *UnSignedOracle) getQuotation(symbol string) (*models.Quotation, error) {

	rawQuotation, err := getQuotationFromDia(symbol)
	if err != nil {
		log.Errorf("Failed to retrieve %s quotation data from DIA: %v", symbol, err)
		return nil, err
	}

	return rawQuotation, nil

}

func (lo *UnSignedOracle) getPairQuotation(pair string) (*models.Quotation, *models.Quotation, error) {
	log.Println("symnol,", pair)
	symbols := strings.Split(pair, "/")
	symbol0, err := getQuotationFromDia(symbols[0])
	if err != nil {
		log.Errorf("error getting quotation for %f %v ", symbol0, err)
		return nil, nil, err

	}
	symbol1, err := getQuotationFromDia(symbols[1])
	if err != nil {
		log.Errorf("error getting quotation for %f %v ", symbol1, err)
		return nil, nil, err
	}
	return symbol0, symbol1, nil

}

func (lo *UnSignedOracle) updatePair(quoteQuotation *models.Quotation, baseQuotation *models.Quotation) error {
	symbol := quoteQuotation.Symbol + "/" + baseQuotation.Symbol
	price := quoteQuotation.Price / baseQuotation.Price
	timestamp := time.Now().Unix()
	err := lo.updateOracle(symbol, int64(price*100000000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func (lo *UnSignedOracle) deployOrBindContract(deployedContract string) error {
	var err error
	if deployedContract != "" {
		lo.contract, err = diaPcwsOracleService.NewDIAPcwsOracle(common.HexToAddress(deployedContract), lo.conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, lo.contract, err = diaPcwsOracleService.DeployDIAPcwsOracle(lo.auth, lo.conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		log.Printf("Contract pending deploy: 0x%x\n", addr)
		log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
		time.Sleep(180000 * time.Millisecond)
	}
	return nil
}

func (lo *UnSignedOracle) updateOracle(
	key string,
	value int64,
	timestamp int64) error {

	gasPrice, err := lo.conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Get 110% of the gas price
	fmt.Println(gasPrice)
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.1))
	gasPrice, _ = fGas.Int(nil)
	fmt.Println(gasPrice)
	// Write values to smart contract
	tx, err := lo.contract.SetValue(&bind.TransactOpts{
		From:     lo.auth.From,
		Signer:   lo.auth.Signer,
		GasLimit: 800725,
		GasPrice: gasPrice,
	}, key, big.NewInt(value), big.NewInt(timestamp))
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("key: %s\n", key)
	log.Println("Value: ", value)

	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func (lo *UnSignedOracle) updateIndexValue(iv *models.CryptoIndex) error {
	symbol := iv.Asset.Name
	value := iv.Value
	timestamp := iv.CalculationTime.Unix()
	err := lo.updateOracle(symbol, int64(value*10000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}
