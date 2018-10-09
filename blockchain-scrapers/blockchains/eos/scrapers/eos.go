package main

/*
 ** Please make sure to run `go get github.com/urfave/cli`
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/diadata-org/api-golang/pkg/dia"
	cli "github.com/urfave/cli"
)

type ChainInfo struct {
	ServerVersion           string `json:"server_version_string"`
	ServerVersionHex        string `json:"server_version"`
	ChainIDHex              string `json:"chain_id"`
	HeadBlock               int    `json:"head_block_num"`
	HeadBlockID             string `json:"head_block_id"`
	LastIrreversibleBlock   int    `json:"last_irreversible_block_num"`
	LastIrreversibleBlockID string `json:"last_irreversible_block_id"`
	HeadBlockTime           string `json:"head_block_time"`
	HeadBlockProducer       string `json:"head_block_producer"`
	VirtualCPULimit         int    `json:"virtual_block_cpu_limit"`
	VirtualNetLimit         int    `json:"virtual_block_net_limit"`
	CPULimit                int    `json:"block_cpu_limit"`
	NetLimit                int    `json:"block_net_limit"`
}

type CurrencyInfo struct {
	Supply    string `json:"supply"`
	MaxSupply string `json:"max_supply"`
	Issuer    string `json:"issuer"`
}

type CurrencyResponse struct {
	Info CurrencyInfo `json:"EOS"`
}

func main() {
	app := cli.NewApp()
	app.Name = "EOS Info Tracker"
	app.Usage = "Use this to interact with the EOS RPC and get useful information"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpc",
			Value: "http://0.0.0.0:8888",
			Usage: "Which RPC Endpoint do you want to query?",
		},
		cli.StringFlag{
			Name:  "action",
			Value: "info",
			Usage: "What kind of info do you want from the RPC endpoint?",
		},
	}

	app.Action = func(c *cli.Context) error {
		config := dia.GetConfigApi()
		if config == nil {
			log.Fatal("Couldnt load config")
		}

		client := dia.NewClient(config)
		if client == nil {
			log.Fatal("Couldnt load client")
		}

		fmt.Println("")
		fmt.Println("Querying RPC Endpoint:", c.String("rpc"))
		fmt.Println("Querying with action:", c.String("action"))

		if c.String("action") == "info" {
			chainInfo(c.String("rpc"))
			CurrentSupply := currencyInfo(c.String("rpc"))

			client.SendSupply(&dia.Supply{
				Symbol:            "EOS",
				CirculatingSupply: CurrentSupply,
			})
		} else {
			log.Fatal("Invalid action specified")
		}

		fmt.Println("Successfully queried RPC Endpoint", c.String("rpc"))
		fmt.Println("")

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func chainInfo(rpc string) {
	uri := rpc + "/v1/chain/get_info"
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var payload ChainInfo
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("Chain Info for EOS")
	fmt.Println("Server Version:", payload.ServerVersion)
	fmt.Println("Latest Block:", payload.HeadBlock)
	fmt.Println("Latest Block Hex:", payload.HeadBlockID)
	fmt.Println("Latest Block Time:", payload.HeadBlockTime)
	fmt.Println("Latest Block Producer:", payload.HeadBlockProducer)
	fmt.Println("CPU Limit:", payload.CPULimit)
	fmt.Println("Net Limit:", payload.NetLimit)
	fmt.Println("")
}

func currencyInfo(rpc string) float64 {
	uri := rpc + "/v1/chain/get_currency_stats"
	jsonBody := []byte(`{"symbol":"EOS","code":"eosio.token"}`)

	res, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var payload CurrencyResponse
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Println("Currency Stats for the EOS Token")
	fmt.Println("Issuer:", payload.Info.Issuer)
	fmt.Println("Current Supply:", payload.Info.Supply)
	fmt.Println("Max Supply:", payload.Info.MaxSupply)
	fmt.Println("")

	strLen := len(payload.Info.Supply)
	formattedSupply, err := strconv.ParseFloat(payload.Info.Supply[:strLen-4], 64)
	if err != nil {
		log.Fatal(err)
	}

	return formattedSupply
}
