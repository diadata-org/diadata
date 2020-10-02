package tokenService

import (
	"fmt"
	"math"
	"math/big"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func GetTotalSupplyfromMainNet(tokenAddress string, datastore models.Datastore) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/your-key-here")
	if err != nil {
		log.Error(err)
	}

	// Address "0xa74476443119A942dE498590Fe1f2454d7D4aC0d"
	tokenAddressInHex := common.HexToAddress(tokenAddress)
	instance, err := NewERC20(tokenAddressInHex, client)
	if err != nil {
		log.Fatal(err)
	}

	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client3")

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	// "wei: 74605500647408739782407023"
	fmt.Printf("wei: %s\n", totalSupply) // "wei: 74605500647408739782407023"
	fmt.Println("")
	fbal := new(big.Float)

	fbal.SetString(totalSupply.String())
	valuei := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	totalSup, _ := valuei.Float64()
	fmt.Printf("balance: %f", totalSup) // "balance: 74605500.647409"
	fmt.Println("")
	timeNow := time.Now()
	// Parse last updated timestamp
	//layout := "2006-01-02T15:04:05.000Z"
	//timestamp, err := time.Parse(layout, timeNow)
	//ndecimals := decimals.(int)
	//reflect.ValueOf(ndecimals ).Int64()
	if err != nil {
		log.Fatal(err)
	}

	token := models.Token{
		Symbol:      symbol,
		Name:        name,
		TotalSupply: totalSup,
		Decimals:    int(decimals),
		Source:      "Ethereum Mainnet",
		Time:        timeNow,
	}
	fmt.Println("")
	fmt.Printf("Token %v", token)
	datastore.SaveTokenDetailInflux(token)

	time.Sleep(2 * time.Second)
	tk, err := datastore.GetTokenDetailInflux(symbol, "Ethereum Mainnet", time.Now())

	fmt.Printf("Token %v", tk)
}
