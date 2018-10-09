package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/diadata-org/api-golang/pkg/dia"
)

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

type Result struct {
	EmissionAmount *big.Int    `json:"emission_amount"`
	BlockCount     int64       `json:"count"`
	BlockHeader    BlockHeader `json:"block_header"`
}

type Response struct {
	Id      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Result  Result `json:"result"`
}

type BlockHeader struct {
	Timestamp int64 `json:"timestamp"`
	Height    int64 `json:"height"`
}

// Constants
const (
	symbol           = "XMR"
	endpointJsonRpc  = "http://monero:18081/json_rpc"
	atomicConversion = 0.000000000001

	// This snapshot was taken on 2018-09-04
	// These values can be updated as performance degrades:
	snapshotHeight   int64  = 1654077
	snapshotEmission string = "16377801401737323881"
	// snapshot can be calculated from command line easily using curl:
	// curl -X POST http://127.0.0.1:18081/json_rpc -d \
	// '{"jsonrpc":"2.0","id":"0","method":"get_coinbase_tx_sum","params":{"height":0,"count":1654077}}' \
	// -H 'Content-Type: application/json'
)

func getJsonRpcResponse(method string, params *Params) (*Response, error) {
	// marshall the post body json
	data := PostData{
		Id:      "0",
		JsonRpc: "2.0",
		Method:  method,
		Params:  *params,
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
	var response Response
	err = json.Unmarshal(resultData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func getLastBlock() (int64, int64, error) {
	params := &Params{}
	response, err := getJsonRpcResponse("get_last_block_header", params)
	if err != nil {
		return 0, 0, err
	}

	height := response.Result.BlockHeader.Height
	timestamp := response.Result.BlockHeader.Timestamp

	return height, timestamp, err
}

func getEmission(currentHeight int64, previousHeight int64) (*big.Rat, error) {
	// returns the emission amount between current height and previous height

	params := &Params{
		Height: previousHeight,
		Count:  currentHeight - previousHeight,
	}

	response, err := getJsonRpcResponse("get_coinbase_tx_sum", params)

	if err != nil {
		return nil, err
	}

	// convert json number to big.Rat
	emissionFloat := new(big.Rat)
	emissionFloat.SetInt(response.Result.EmissionAmount)
	return emissionFloat, err
}

func atomicToXmr(x *big.Rat) float64 {
	// convert from atomic units to XMR value

	conversion := new(big.Rat)
	atomic := new(big.Rat)
	atomic.SetFloat64(atomicConversion)
	conversion.Mul(x, atomic)

	// convert XMR from big.Rat to float64
	xmrSupplyFloat, _ := conversion.Float64()

	return xmrSupplyFloat
}

func publishCirculatingSupply(client *dia.Client, height int64, supply *big.Rat) {
	xmrSupply := atomicToXmr(supply)
	log.Printf("Updating circulating supply as of block %v: %f\n", height, xmrSupply)

	client.SendSupply(&dia.Supply{
		Symbol:            symbol,
		CirculatingSupply: xmrSupply,
	})
}

func main() {

	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldn't load config")
	}
	client := dia.NewClient(config)
	if client == nil {
		panic("Couldn't load client")
	}

	// initialize snapshot values
	var previousHeight = snapshotHeight
	var totalEmission = new(big.Rat)
	totalEmission.SetString(snapshotEmission)

	for {
		currentHeight, currentBlockTimestamp, err := getLastBlock()
		if err != nil {
			log.Println("Failed to retrieve current height from node")
			continue
		}

		now := time.Now()
		blockTime := time.Unix(currentBlockTimestamp, 0)
		twoMinutes := time.Duration(60*2) * time.Second

		if now.Sub(blockTime) > twoMinutes {
			log.Println("Waiting for new block")
		} else {
			if currentHeight > previousHeight {
				log.Printf("New height: %v\n", currentHeight)
				newEmission, err := getEmission(currentHeight, previousHeight)

				if err == nil {
					// add new emission to previous emission, set previousHeight
					totalEmission.Add(totalEmission, newEmission)
					previousHeight = currentHeight

					publishCirculatingSupply(client, currentHeight, totalEmission)

				} else {
					log.Println("Error communicating with node:", err)
				}

			} else if currentHeight == previousHeight {
				// no change, publish last known height/emission
				publishCirculatingSupply(client, previousHeight, totalEmission)
			}
		}

		time.Sleep(time.Second * 10)
	}
}
