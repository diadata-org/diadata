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

type Response struct {
	Result string `json:"result"`
	Count  int    `json:"count"`
	Marker string `json:"marker"`
	Row    []Row  `json:"rows"`
}

type Row struct {
	Date          string `json:"date"`
	Total         string `json:"total"`
	Distributed   string `json:"distributed"`
	Undistributed string `json:"undistributed"`
	Escrowed      string `json:"escrowed"`
}

const endpoint = "https://data.ripple.com/v2/network/xrp_distribution?limit=1&descending=true"
const symbol = "XRP"

func main() {
	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldnt load config")
	}
	client := dia.NewClient(config)
	if client == nil {
		panic("Couldnt load client")
	}
	for {
		response, err := http.Get(endpoint)
		if err == nil {
			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Println("Failed to retrieve xrp supply: ", err)
			}
			var responseObject Response
			uerr := json.Unmarshal(responseData, &responseObject)
			if uerr != nil {
				fmt.Println("There was an response unmarshalling error:", err)
			}
			distributed := responseObject.Row[0].Distributed
			result, cerr := strconv.ParseFloat(distributed, 64)
			if cerr != nil {
				fmt.Println("There was an integer conversion error:", err)
			}
			fmt.Printf("Symbol: %s ; circulatingSupply: %f\n", symbol, result)
			client.SendSupply(&dia.Supply{
				Symbol:            symbol,
				CirculatingSupply: result,
			})
		} else {
			log.Println("Err communicating with node:", err)
		}
		time.Sleep(time.Second * 10)
	}
}
