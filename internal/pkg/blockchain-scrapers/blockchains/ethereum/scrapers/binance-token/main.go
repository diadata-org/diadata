package main

import (
	"flag"
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"time"
)

//Be sure to run : abigen --abi erc20.abi --pkg main --type Token --out erc20.go
//erc20.abi taken from https://github.com/ethereum/wiki/wiki/Contract-ERC20-ABI
var (
	rpcEndpoint  = flag.String("rpc", "http://geth:8545", "geth RPC endpoint")
	dev          = flag.Bool("dev", false, "Dev mode - prints to stdout instead of sending to dia")
	symbol       = "BNB"
	address      = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	ownerAddress = "0x00c5e04176d95a286fcce0e68c683ca0bfec8454"
)

func init() {
	flag.Parse()
}

func toFloat(in *big.Int, decimals uint8) float64 {
	f := new(big.Float).SetInt(in)
	fl, _ := f.Mul(f, big.NewFloat(math.Pow(float64(10), -1*float64(decimals)))).Float64()
	return fl
}

func main() {

	conn, err := ethclient.Dial(*rpcEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract
	token, err := NewBNB(common.HexToAddress(address), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	//
	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldnt load config")
	}
	client := dia.NewClient(config)
	if client == nil {
		panic("Couldnt load client")
	}
	for {
		//Infinite loop, sends tokenSupply every 10 seconds
		supply, err := token.TotalSupply(nil)
		if err != nil {
			//Perhaps these should not be a fatal error?
			log.Fatalf("Failed to retrieve token supply: %v", err)
		}

		ownerSupply, err := token.BalanceOf(nil, common.HexToAddress(ownerAddress))
		if err != nil {
			log.Fatalf("Failed to retrieve token ownerSupply: %v", err)
		}

		frozenSupply, err := token.FreezeOf(nil, common.HexToAddress(ownerAddress))
		if err != nil {
			log.Fatalf("Failed to retrieve token frozenSupply: %v", err)
		}

		tmpSupply := supply.Sub(supply, ownerSupply)
		resultSupply := supply.Sub(tmpSupply, frozenSupply)

		decimals, err := token.Decimals(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve token decimal: %v", err)
		}

		result := toFloat(resultSupply, decimals)

		fmt.Printf("Symbol: %s ; totalSupply: %f\n", symbol, result)

		client.SendSupply(&dia.Supply{
			Symbol:            symbol,
			CirculatingSupply: result,
		})

		time.Sleep(time.Second * 10)
	}
}
