package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Supply struct {
	Updatad     string  `json:"updatedAt"`
	Total       float64 `json:"totalCoins"`
	Available   float64 `json:"availableCoins"`
	Distributed float64 `json:"distributedCoins"`
	Programs    struct {
		Direct      float64 `json:"directProgram"`
		Bitcoin     float64 `json:"bitcoinProgram"`
		Partnership float64 `json:"partnershipProgram"`
		Build       int     `json:"buildChallenge"`
	} `json:"programs"`
}

const endpoint = "https://dashboard.stellar.org/api/lumens"
const symbol = "XLM"

func main() {
	// config := dia.GetConfigApi()
	// if config == nil {
	// 	panic("Couldnt load config")
	// }
	// client := dia.NewClient(config)
	// if client == nil {
	// 	panic("Couldnt load client")
	// }
	// prevResult := 0.0
	for {
		response, err := http.Get(endpoint)
		if err == nil {
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println("Failed to retrieve ada supply: ", err)
			}
			var supplyObject Supply
			json.Unmarshal(responseData, &supplyObject)
			result := supplyObject.Distributed
			fmt.Printf("Symbol: %s ; circulatingSupply: %f\n", symbol, result)
			// if prevResult != result {
			// 	client.SendSupply(&dia.Supply{
			// 		Symbol:            symbol,
			// 		CirculatingSupply: result,
			// 	})
			// 	prevResult = result
			// }
		} else {
			log.Println("Err communicating with node:", err)
		}
		time.Sleep(time.Second * 10)
	}

}
