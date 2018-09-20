package main

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/go-bitcoind"
	"log"
	"time"
)

const (
	SERVER_HOST = "dogecoin"
	SERVER_PORT = 22555
	USER        = "dogediauser"
	PASSWD      = "dogediapassword"
	SYMBOL      = "DOGE"
	USESSL		= false
	SLEEP_SECONDS = 10
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldn't load config")
	}

	client := dia.NewClient(config)
	if client == nil {
		panic("Couldn't load client")
	}

	bitcoinLib, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		panic(err)
	}

	var lastSupply float64 = -1
	for {
		blockchainInfo, err := bitcoinLib.GetBlockchainInfo()
		if err != nil {
			panic("Error communicating with doge node: "+ err.Error())
		}

		txOutSetInfo, err := bitcoinLib.GetTxOutsetInfo()
		if err != nil {
			panic("Error communicating with doge node: "+ err.Error())
		}

		// new supply value
		if txOutSetInfo.TotalAmount > lastSupply {
			lastSupply = txOutSetInfo.TotalAmount
			log.Println(lastSupply)

			err = client.SendSupply(&dia.Supply{
				Symbol: SYMBOL,
				CirculatingSupply: lastSupply,
			})
			if err != nil {
				log.Println("Error sending supply to API: ", err)
			}
		} else {
			medianSeconds := time.Duration(blockchainInfo.Mediantime)*time.Second

			if medianSeconds < SLEEP_SECONDS*time.Second {
				time.Sleep(medianSeconds)
			} else {
				time.Sleep(SLEEP_SECONDS *time.Second)
			}
		}
	}
}
