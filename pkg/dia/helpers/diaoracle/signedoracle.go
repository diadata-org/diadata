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

type SignedOracle struct {
	conn              *ethclient.Client
	auth              *bind.TransactOpts
	contract          *diaPcwsOracleService.DIAPcwsOracle
	oracleAddress     string
	deviationPermille int
	updateOnDeviation bool
	symbols           []string
	oldPrices         map[string]int64
}

func NewSignedOracleContract(conn *ethclient.Client, auth *bind.TransactOpts, oracleAddress string, deviationPermille int, updateOnDeviation bool, symbols []string) *SignedOracle {
	oldPrices := make(map[string]int64)
	lo := &SignedOracle{conn: conn, auth: auth, oracleAddress: oracleAddress, symbols: symbols, oldPrices: oldPrices, deviationPermille: deviationPermille, updateOnDeviation: updateOnDeviation}
	return lo
}

func (lo *SignedOracle) Update() error {
	err := lo.deployOrBindContract(lo.oracleAddress)
	if err != nil {
		return err
	}

	if lo.oracleAddress != "" {

		for _, symbol := range lo.symbols {
			if strings.Contains(symbol, "/") {
				err := lo.getPairQuotation(symbol)
				if err != nil {
					continue
				}

			}

		}

	}
	return nil

}

func (lo *SignedOracle) getPairQuotation(pair string) error {
	log.Println("symnol,", pair)
	symbols := strings.Split(pair, "/")
	symbol0, err := getQuotationFromDia(symbols[0])
	if err != nil {
		log.Errorf("error getting quotation for %f %v ", symbol0, err)
	}
	symbol1, err := getQuotationFromDia(symbols[1])
	if err != nil {
		log.Errorf("error getting quotation for %f %v ", symbol1, err)
	}
	lo.updatePair(symbol0, symbol1)
	return nil

}

func (lo *SignedOracle) updatePair(quoteQuotation *models.Quotation, baseQuotation *models.Quotation) error {
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

func (lo *SignedOracle) deployOrBindContract(deployedContract string) error {
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

func (lo *SignedOracle) updateOracle(
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
