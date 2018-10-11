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
	rpcEndpoint = flag.String("rpc", "http://geth:8545", "geth RPC endpoint")
	symbol      = flag.String("symbol", "", "ERC20 token name")
	address     = flag.String("address", "", "The address at which the contract was deployed")
	dev         = flag.Bool("dev", false, "Dev mode - prints to stdout instead of sending to dia")
)

func init() {
	flag.Parse()
	if *symbol == "" {
		flag.Usage()
		log.Fatal("Symbol name is required")
	}
	if *address == "" {
		flag.Usage()
		log.Fatal("Contract address is required")
	}
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
	token, err := NewToken(common.HexToAddress(*address), conn)
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

		ownerSupply, err := token.BalanceOf(nil, common.HexToAddress("0x00c5e04176d95a286fcce0e68c683ca0bfec8454"))
		if err != nil {
			log.Fatalf("Failed to retrieve token ownerSupply: %v", err)
		}

		resultSupply := supply.Sub(supply, ownerSupply)

		decimals, err := token.Decimals(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve token decimal: %v", err)
		}

		result := toFloat(resultSupply, decimals)

		fmt.Printf("Symbol: %s ; totalSupply: %f\n", *symbol, result)

		client.SendSupply(&dia.Supply{
			Symbol:            *symbol,
			CirculatingSupply: result,
		})

		time.Sleep(time.Second * 10)
	}
}
