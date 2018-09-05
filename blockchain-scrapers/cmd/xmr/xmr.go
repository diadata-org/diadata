package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	//"github.com/diadata-org/api-golang/dia"
)

type HeightResponse struct {
	Value int64 `json:"height"`
}

type PostData struct {
	Id      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

type Params struct {
	Height int64 `json:"height"`
	Count  int64 `json:"count"`
}

type EmissionResult struct {
	Value *big.Int `json:"emission_amount"`
}

type Result struct {
	Id      string         `json:"id"`
	JsonRpc string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Result  EmissionResult `json:"result"`
}

// Constants
const symbol = "XMR"
const endpointGetHeight = "http://localhost:18081/getheight"
const endpointJsonRpc = "http://localhost:18081/json_rpc"
const atomicConversion = 0.000000000001

// This snapshot was taken on 2018-09-04
// These values can be updated as performance degrades:
const snapshotHeight int64 = 1654077
const snapshotEmission string = "16377801401737323881"
// Can be calculated from command line easily:
// curl -X POST http://127.0.0.1:18081/json_rpc -d '{"jsonrpc":"2.0","id":"0","method":"get_coinbase_tx_sum","params":{"height":0,"count":1654077}}' -H 'Content-Type: application/json'

func getCurrentHeight() (int64, error) {
	// retrieve json from /getheight endpoint
	response, err := http.Post(endpointGetHeight, "application/json", nil)
	if err != nil {
		return 0, err
	}

	// read body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	// unmarshal json and return height
	var latestHeight HeightResponse
	json.Unmarshal(responseData, &latestHeight)
	result := latestHeight.Value

	return result, err
}

func getEmission(currentHeight int64, previousHeight int64) (*big.Rat, error) {
	// returns the emission amount between current height and previous height

	// marshall the post body json
	data := PostData{
		Id:      "0",
		JsonRpc: "2.0",
		Method:  "get_coinbase_tx_sum",
		Params: Params{
			Height: previousHeight,
			Count:  currentHeight - previousHeight,
		},
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// post to endpoint
	result, err := http.Post(endpointJsonRpc, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	// read body
	resultData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	// extract emission value from json
	var emissionResult Result
	err = json.Unmarshal(resultData, &emissionResult)
	if err != nil {
		return nil, err
	}

	// convert json number to big.Rat
	emissionFloat := new(big.Rat)
	emissionFloat.SetInt(emissionResult.Result.Value)
	return emissionFloat, err
}

func atomicToXmr(x *big.Rat) *big.Rat {
	// convert from atomic units to XMR value

	conversion := new(big.Rat)
	atomic := new(big.Rat)
	atomic.SetFloat64(atomicConversion)
	conversion.Mul(x, atomic)

	return conversion
}

func main() {

	//config := dia.GetConfigApi()
	//if config == nil {
	//	panic("Couldn't load config")
	//}
	//client := dia.NewClient(config)
	//if client == nil {
	//	panic("Couldn't load client")
	//}

	// initialize snapshot values
	var previousHeight = snapshotHeight
	var totalEmission = new(big.Rat)
	totalEmission.SetString(snapshotEmission)

	for {
		currentHeight, err := getCurrentHeight()
		if err != nil {
			log.Println("Failed to retrieve current height from node")
			continue
		}

		if currentHeight > previousHeight {
			fmt.Printf("New height; %v\n", currentHeight)
			newEmission, err := getEmission(currentHeight, previousHeight)

			if err == nil {
				// add new emission to previous emission, set previousHeight
				totalEmission.Add(totalEmission, newEmission)
				previousHeight = currentHeight

				xmrSupply := atomicToXmr(totalEmission)

				// convert XMR from big.Rat to float64
				xmrSupplyFloat, _ := xmrSupply.Float64()
				fmt.Printf("Total circulating supply: %f\n", xmrSupplyFloat)

				//client.SendSupply(&dia.Supply{
				//	Symbol:            symbol,
				//	CirculatingSupply: xmrSupplyFloat,
				//})
			} else {
				log.Println("Error communicating with node:", err)
			}

			time.Sleep(time.Second * 10)
		}
	}
}
