package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/diadata-org/api-golang/dia"
)

type Supply struct {
	Value float64 `json:"Right"`
}

const endpoint = "http://cardano:8100/api/supply/ada"
const symbol = "ADA"

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
			result := supplyObject.Value
			fmt.Printf("Symbol: %s ; totalSupply: %f\n", symbol, result)
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
