package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/diadata-org/api-golang/dia"
)

type Supply struct {
	Updatad     string `json:"updatedAt"`
	Total       string `json:"totalCoins"`
	Available   string `json:"availableCoins"`
	Distributed string `json:"distributedCoins"`
	Programs    struct {
		Direct      string `json:"directProgram"`
		Bitcoin     string `json:"bitcoinProgram"`
		Partnership string `json:"partnershipProgram"`
		Build       string `json:"buildChallenge"`
	} `json:"programs"`
}

const endpoint = "https://dashboard.stellar.org/api/lumens"
const symbol = "TRX"

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
		response, err := http.Get(endpoint)
		if err == nil {
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println("Failed to retrieve ada supply: ", err)
			}
			var supplyObject Supply
			json.Unmarshal(responseData, &supplyObject)
			result, cerr := strconv.ParseFloat(supplyObject.Distributed, 64)
			if cerr != nil {
				fmt.Println("There was an integer conversion error:", err)
			}
			fmt.Printf("Symbol: %s ; circulatingSupply: %f\n", symbol, result)
			if prevResult != result {
				client.SendSupply(&dia.Supply{
					Symbol:            symbol,
					CirculatingSupply: result,
				})
				prevResult = result
			}
		} else {
			log.Println("Err communicating with node:", err)
		}
		time.Sleep(time.Second * 10)
	}
}
