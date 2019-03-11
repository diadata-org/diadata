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
	rpcEndpoint  = flag.String("rpc", "https://mainnet.infura.io/v3/7db804679c8644fcbfbfd440ed2332eb", "geth RPC endpoint")
	dev          = flag.Bool("dev", false, "Dev mode - prints to stdout instead of sending to dia")
	symbol       = "MKR"
	address      = "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2"
	ownerAddress = "0x0000000000000000000000000000000000000000"
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
	token, err := NewMKR(common.HexToAddress(address), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	//Check for unexpected errors while loading config file or a new client
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
			//Perhaps this should not be a fatal error?
			log.Fatalf("Failed to retrieve token supply: %v", err)
		}

		ownerSupply, err := token.BalanceOf(nil, common.HexToAddress(ownerAddress))
		if err != nil {
			log.Fatalf("Failed to retrieve token ownerSupply: %v", err)
		}

		tmpSupply := supply.Sub(supply, ownerSupply)
		resultSupply := tmpSupply

		decimals, err := token.Decimals(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve token decimal: %v", err)
		}

		result := toFloat(resultSupply, decimals)

		fmt.Printf("Symbol: %s ; totalSupply: %f\n", symbol, result)

		client.SendSupply(&dia.Supply{
			Symbol: symbol,
			CirculatingSupply: result,
		})

		time.Sleep(time.Second * 15)
	}
}
