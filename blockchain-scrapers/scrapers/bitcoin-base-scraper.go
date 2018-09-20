package blockchainscrapers

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/go-bitcoind"
	"log"
	"time"
)

const (
	SLEEP_SECONDS = 10
)

func RunScraper(host string, port int, user, password, symbol string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := dia.GetConfigApi()
	if config == nil {
		panic("Couldn't load config")
	}

	client := dia.NewClient(config)
	if client == nil {
		panic("Couldn't load client")
	}

	bitcoinLib, err := bitcoind.New(host, port, user, password, USESSL)
	if err != nil {
		panic(err)
	}

	var lastSupply float64 = 0
	for {
		txOutSetInfo, err := bitcoinLib.GetTxOutsetInfo()
		if err != nil {
			panic("Error communicating with the "+ symbol +" node: "+ err.Error())
		}

		// new supply value
		if txOutSetInfo.TotalAmount > lastSupply {
			lastSupply = txOutSetInfo.TotalAmount

			err = client.SendSupply(&dia.Supply{
				Symbol: symbol,
				CirculatingSupply: lastSupply,
			})
			if err != nil {
				log.Println("Error sending supply to API: ", err)
			} else {
				log.Println("Sent supply", lastSupply, "to API")
			}
		} else {
			log.Println("No new values. Sleeping for", SLEEP_SECONDS, "seconds")
			time.Sleep(SLEEP_SECONDS *time.Second)
		}
	}
}
