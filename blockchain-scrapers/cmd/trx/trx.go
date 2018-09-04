package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/diadata-org/api-golang/dia"
)

// BalanceInfo : Tron foundation api json
type BalanceInfo struct {
	Total float64 `json:"total"`
}

// TrxAccount : Tronscan api account json
type TrxAccount struct {
	Balance int `json:"balance"`
}

// TrxBlock : Tronscan api block json
type TrxBlock struct {
	Number int `json:"number"`
}

const (
	tronFoundationEndpoint = "https://tron.network/api/v2/node/balance_info"
	tronscanEndpoint       = "https://api.tronscan.org/api"
	blackholeAddress       = "TLsV52sRDL79HXGGm9yzwKibb6BeruhUzy"
	symbol                 = "TRX"
	oneTrx                 = 1000000
	genesisSupply          = 100000000000
	independenceDayBurned  = 1000000000
	startFeeBurnedNum      = -9223372036854.775808
)

// GetFoundationBalance : Get the balance of founders wallets
func GetFoundationBalance() (float64, error) {
	response, err := http.Get(tronFoundationEndpoint)
	if err == nil {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return 0.0, errors.New("Failed to retrieve foundation balance")
		}
		var balanceObject BalanceInfo
		json.Unmarshal(responseData, &balanceObject)
		return balanceObject.Total, nil
	}
	return 0.0, errors.New("Err communicating with tron network api")
}

// GetWalletTrxBalance : Get the balance of wallet
func GetWalletTrxBalance(address string) (int, error) {
	response, err := http.Get(tronscanEndpoint + "/account/" + address)
	if err == nil {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return 0, errors.New("Failed to retrieve blackhole balance")
		}
		var accountObject TrxAccount
		json.Unmarshal(responseData, &accountObject)
		return accountObject.Balance, nil
	}
	return 0, errors.New("Err communicating with tron network api")
}

// GetTrxBlockHeight : Get the block height of trx blockchain
func GetTrxBlockHeight() (int, error) {
	response, err := http.Get(tronscanEndpoint + "/block/latest")
	if err == nil {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return 0, errors.New("Failed to retrieve blackhole balance")
		}
		var blockObject TrxBlock
		json.Unmarshal(responseData, &blockObject)
		return blockObject.Number, nil
	}
	return 0, errors.New("Err communicating with tron network api")
}

func main() {
	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldnt load config")
	}
	client := dia.NewClient(config)
	if client == nil {
		panic("Couldnt load client")
	}
	prevResult := 0.0
	for {
		tronFoundationTotal, err := GetFoundationBalance()
		if err != nil {
			log.Println("Err in GetFoundationBalance:", err)
		}
		// fmt.Printf("tronFoundationTotal: %f\n", tronFoundationTotal)
		blackholeBalance, err := GetWalletTrxBalance(blackholeAddress)
		if err != nil {
			log.Println("Err in GetWalletTrxBalance:", err)
		}
		feeBurnedNum := math.Abs(startFeeBurnedNum) - math.Abs(float64(blackholeBalance/oneTrx))
		// fmt.Printf("feeBurnedNum: %f\n", feeBurnedNum)
		blockHeight, err := GetTrxBlockHeight()
		if err != nil {
			log.Println("Err in GetTrxBlockHeight:", err)
		}
		nodeRewardsNum := blockHeight * 16
		// fmt.Printf("nodeRewardsNum: %d\n", nodeRewardsNum)
		blockProduceRewardsNum := blockHeight * 32
		// fmt.Printf("blockProduceRewardsNum: %d\n", blockProduceRewardsNum)
		currentTotalSupply := genesisSupply + float64(blockProduceRewardsNum) + float64(nodeRewardsNum) - float64(independenceDayBurned) - feeBurnedNum
		// fmt.Printf("currentTotalSupply: %f\n", currentTotalSupply)
		result := currentTotalSupply - tronFoundationTotal
		fmt.Printf("Symbol: %s ; circulatingSupply: %f\n", symbol, result)
		if prevResult != result {
			client.SendSupply(&dia.Supply{
				Symbol:            symbol,
				CirculatingSupply: result,
			})
			prevResult = result
		}
		time.Sleep(time.Second * 10)
	}
}
