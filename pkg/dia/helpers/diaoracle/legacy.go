package diaoracle

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/oracleService"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LegacyOracle struct {
	conn           *ethclient.Client
	auth           *bind.TransactOpts
	contract       *oracleService.DiaOracle
	oracleAddress  string
	symbols        []string
	diaServer      string
	defiRate       map[string]string
	defiState      []string
	dexChartPoints map[string]string
	farmingPools   map[string]string
}

func NewLegacyContract(conn *ethclient.Client, auth *bind.TransactOpts, oracleAddress string, symbols []string, defiRate map[string]string, defiState []string, dexChartPoints, farmingPools map[string]string, diaServer string) *LegacyOracle {

	lo := &LegacyOracle{conn: conn, auth: auth, oracleAddress: oracleAddress, symbols: symbols, diaServer: diaServer, defiRate: defiRate, defiState: defiState, dexChartPoints: dexChartPoints, farmingPools: farmingPools}
	return lo
}

func (lo *LegacyOracle) Update() error {
	// Get Symbols to update
	err := lo.deployOrBindContract(lo.oracleAddress)
	if err != nil {
		return err
	}

	if lo.oracleAddress != "" {

		// Update Symbols

		for _, symbol := range lo.symbols {

			err := lo.getQuotationAndSupply(symbol)
			if err != nil {
				continue
			}
		}

		// Update DefiRate

		for protocol, symbol := range lo.defiRate {
			log.Info("Protocvo", protocol)
			log.Info("Symbol", symbol)

			rawDefiRate, err := getDefiRatesFromDia(protocol, symbol)
			if err != nil {
				log.Errorf("Failed to retrieve %s data from DIA: %v", protocol, err)
				continue
			}

			err = lo.updateDefiRate(rawDefiRate)
			if err != nil {
				log.Fatalf("Failed to update %s Oracle: %v", protocol, err)
				continue
			}
		}

		// Update DefiState

		for _, symbol := range lo.defiState {
			log.Info("Symbol", symbol)

			rawDefiState, err := getDefiStateFromDia(symbol)
			if err != nil {
				log.Errorf("Failed to retrieve %s data from DIA: %v", err)
				continue
			}

			err = lo.updateDefiState(rawDefiState)
			if err != nil {
				log.Fatalf("Failed to update %s Oracle: %v", err)
				continue
			}
		}

		//Update Dex Chart Points

		for exchange, symbol := range lo.dexChartPoints {
			log.Info("exchange", exchange)
			log.Info("Symbol", symbol)

			rawDefiRate, err := getDEXFromDia(exchange, symbol)
			if err != nil {
				log.Errorf("Failed to retrieve for exchnage %s data from DIA: %v", exchange, err)
				continue
			}

			err = lo.updateDEX(rawDefiRate)
			if err != nil {
				log.Fatalf("Failed to update %s Oracle: %v", exchange, err)
				continue
			}
		}

		// Update Farming Pools

		for poolName, symbol := range lo.farmingPools {
			log.Info("exchange", poolName)
			log.Info("Symbol", symbol)

			rawDefiRate, err := getFarmingPoolFromDia(poolName, symbol)
			if err != nil {
				log.Errorf("Failed to retrieve for poolName %s data from DIA: %v", poolName, err)
				continue
			}

			err = lo.updateFarmingPool(rawDefiRate)
			if err != nil {
				log.Fatalf("Failed to update %s Oracle: %v", poolName, err)
				continue
			}
		}
	} else {

	}
	return nil

}

func (lo *LegacyOracle) getQuotationAndSupply(symbol string) error {

	rawQuotation, err := getQuotationFromDia(symbol)
	if err != nil {
		log.Errorf("Failed to retrieve %s quotation data from DIA: %v", symbol, err)
		return err
	}

	rawSupply, err := getSupplyFromDia(symbol)
	if err != nil {
		log.Errorf("Failed to retrieve %f supply data from DIA: %v", symbol, err)
		return err
	}

	err = lo.updateQuotation(rawQuotation, rawSupply)
	if err != nil {
		log.Errorf("Failed to update  Oracle: %v", err)
		return err
	}
	return nil

}

func (lo *LegacyOracle) updateQuotation(quotation *models.Quotation, supply *dia.Supply) error {
	name := quotation.Name
	symbol := quotation.Symbol
	price := quotation.Price
	circSupply := supply.CirculatingSupply
	err := lo.updateOracle(name, symbol, int64(price*100000), int64(circSupply))
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func (lo *LegacyOracle) deployOrBindContract(deployedContract string) error {

	var err error
	if deployedContract != "" {
		lo.contract, err = oracleService.NewDiaOracle(common.HexToAddress(deployedContract), lo.conn)
		if err != nil {
			return err
		}
	} else {
		gasPrice, err := lo.conn.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// Get 110% of the gas price
		fGas := new(big.Float).SetInt(gasPrice)
		fGas.Mul(fGas, big.NewFloat(1.1))
		gasPrice, _ = fGas.Int(nil)

		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		lo.auth.GasPrice = gasPrice
		addr, tx, lo.contract, err = oracleService.DeployDiaOracle(lo.auth, lo.conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		fmt.Println(tx.GasPrice())
		log.Printf("Contract pending deploy: 0x%x\n", addr)
		log.Printf("Transaction waiting to be mined: 0x%x\n", tx.Hash())
		log.Printf("Tx Nonce: %d\n\n", tx.Nonce())
		time.Sleep(180000 * time.Millisecond)
	}
	return nil
}

func (lo *LegacyOracle) updateDefiRate(defiRate *dia.DefiRate) error {
	symbol := strings.ToUpper(defiRate.Asset)
	name := strings.ToUpper(defiRate.Protocol)
	lendingRate := defiRate.LendingRate
	borrowingRate := defiRate.BorrowingRate
	// Get 5 digits after the comma by multiplying price with 100000
	err := lo.updateOracle(name, symbol, int64(lendingRate*100000), int64(borrowingRate*100000))
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func (lo *LegacyOracle) updateOracle(
	name string,
	symbol string,
	price int64,
	supply int64) error {

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
	tx, err := lo.contract.UpdateCoinInfo(&bind.TransactOpts{
		From:     lo.auth.From,
		Signer:   lo.auth.Signer,
		GasLimit: 800725,
		GasPrice: gasPrice,
		//	Nonce: big.NewInt(time.Now().Unix()),
	}, name, symbol, big.NewInt(price), big.NewInt(supply), big.NewInt(time.Now().Unix()))
	// prices are with 5 digits after the comma
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("Symbol: %s\n", symbol)
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	log.Printf("Tx Nonce: %d\n", tx.Nonce())
	return nil
}

func (lo *LegacyOracle) updateDefiState(defiState *dia.DefiProtocolState) error {
	symbol := ""
	name := strings.ToUpper(defiState.Protocol.Name) + "-state"
	price := defiState.TotalUSD
	// Get 5 digits after the comma by multiplying price with 100000
	err := lo.updateOracle(name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func (lo *LegacyOracle) updateDEX(dexData *models.Points) error {
	if len(dexData.DataPoints[0].Series) > 0 && len(dexData.DataPoints[0].Series[0].Values) > 0 {
		symbol := strings.ToUpper(dexData.DataPoints[0].Series[0].Values[0][3].(string))
		name := dexData.DataPoints[0].Series[0].Values[0][1].(string)
		supply := 0
		price := dexData.DataPoints[0].Series[0].Values[0][4].(float64)
		// Get 5 digits after the comma by multiplying price with 100000
		// Set supply to 0, as we don't have a supply for one exchange
		err := lo.updateOracle(name, symbol, int64(price*100000), int64(supply))
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
	} else {
		err := lo.updateOracle("", "", int64(0), int64(0))
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
	}
	return nil
}

func (lo *LegacyOracle) updateFarmingPool(poolData *models.FarmingPool) error {
	protocolName := poolData.ProtocolName
	poolID := poolData.PoolID
	rate := poolData.Rate
	balance := poolData.Balance
	err := lo.updateOracle(protocolName, poolID, int64(rate*100000), int64(balance*100000))
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}
	return nil
}
