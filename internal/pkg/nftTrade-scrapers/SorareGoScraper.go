package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	jsonrpc "github.com/ybbus/jsonrpc"
)

func main() {

	fmt.Println(getRecentBlock())
	//getRecentBlock()
	//request_tokenids()
}

type NFTTrade struct {
	tokenID     string
	BlockNumber int64
	PriceUSD    float64
	Exchange    string
	txhash      string
}

func getRecentBlock() []NFTTrade {
	//Keep track of NFTTrades in this block with the below var
	var Trades_in_block []NFTTrade
	rpcClient := jsonrpc.NewClient("https://mainnet.infura.io/v3/3c49f1c98b4d4de093bdfa3a898ad849")
	//Hex value is block number, use 'latest' in production to save keeping track of blocks
	response, err := rpcClient.Call("eth_getBlockByNumber", "0xAF8D8A", true)
	if err != nil {
		fmt.Printf("Error")
	}
	block := response.Result
	block_number := (block.(map[string]interface{})["number"]).(string)
	runeblock := []rune(block_number)
	var blockhex string
	for i, r := range runeblock {
		if i >= 2 {
			blockhex = blockhex + string(r)
		}
	}
	blockint, err := strconv.ParseInt(blockhex, 16, 64)
	if err != nil {
		fmt.Print("Hex Decoding Error")
	}
	block_txs := block.(map[string]interface{})["transactions"]
	tx_list := block_txs.([]interface{})
	// The above block is mostly type conversions to move the transactions into something more manageable.
	for a := 0; a < 121; a++ {
		var current_slice = tx_list[a].(map[string]interface{})
		var input_data string = current_slice["input"].(string)
		trackbool, output := decodeAIB(input_data)
		if trackbool {
			var transaction string
			transaction = current_slice["from"].(string)
			transaction += " => "
			transaction += output[1]
			NFTtraderesponse := NFTTrade{output[2], blockint, 29, transaction, current_slice["hash"].(string)}
			Trades_in_block = append(Trades_in_block, NFTtraderesponse)
		}
	}
	fmt.Println("Successfully Scraped Block")
	return Trades_in_block
}
func decodeAIB(_data string) (bool, [3]string) {
	//So this function decodes input data to see if its calling
	//The transfer method of Sorare Contracts which has a method id of 0x223da1ba
	runedata := []rune(_data)
	var methodid string
	var arg1 string
	var arg2 string
	var arg3 string
	// Method ID of Sorare NFT transfer, search for this
	var soraremethodid string = "0x223da1ba"
	//loop through rune array to check for method id
	for i, r := range runedata {
		if i >= 10 {
			//this means we have looped through the array long enough to get all needed data
			break
		}
		methodid = methodid + string(r)
	}
	//This section is to parse args from the MethodID, but until ABI decoding is better it doesn't do much
	for i, r := range runedata {
		if i <= 10 {
			//get past the initial loop
		}
		if i >= 10 && i < 74 {
			arg1 = arg1 + string(r)
		}
		if i >= 74 && i < 138 {
			arg2 = arg2 + string(r)
		}
		if i >= 138 && i < 202 {
			arg3 = arg3 + string(r)
		}

		if i >= 2020 {
			break
		}
	}
	var return_array = [3]string{arg1, arg2, arg3}
	if methodid == soraremethodid {
		return true, return_array
	} else {
		return false, return_array
	}
}
func request_tokenids() {

	resp, err := http.Get("https://api.coinranking.com/v2/nfts?dappSlug=Sorare")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//byteValue, _ := ioutil.ReadAll(body)
	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)
	//Gets all needed data into nft format, that can the be broken down even further to easily access token ids
	fmt.Println(result["data"].(map[string]interface{})["nfts"])
}
