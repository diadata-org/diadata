package main

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/PaddyMc/go-tron/calculations"
	"time"
	//"fmt"
)

const (
	SERVER_HOST = "tron"
	SERVER_PORT = "50051"
	SYMBOL      = "TRX"
)

func main() {
	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldnt load config")
	}

	client := dia.NewClient(config)
	if client == nil {
		panic("Couldnt load client")
	}

	calculationsClient := calculations.NewClaculationClient(SERVER_HOST+ ":" +SERVER_PORT)
	var prevResult int64 = 0
	for{
		totalCirculatingSupply, err := calculationsClient.CalculateCirculatingSupply()
		
		if err == nil{
			if prevResult != totalCirculatingSupply {
				//fmt.Printf("Total Circulating Supply: %v\t\n", totalCirculatingSupply)
				client.SendSupply(&dia.Supply{
					Symbol:            SYMBOL,
					CirculatingSupply: float64(totalCirculatingSupply),
				})
				prevResult = totalCirculatingSupply
			}
		}
		time.Sleep(time.Second * 10)
		// time.Sleep(time.Hour * 1)
	}
}