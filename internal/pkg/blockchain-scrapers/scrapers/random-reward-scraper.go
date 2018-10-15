package blockchainscrapers

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/go-bitcoind"
	"log"
	"time"
)

const (
	SLEEP_TIME = 10
)

func RunScraper(host string, port int, user, password, symbol string, elapsedTime int) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := dia.GetConfigApi()
	if config == nil {
		log.Fatal("Couldn't load config")
	}

	client := dia.NewClient(config)
	if client == nil {
		log.Fatal("Couldn't load client")
	}

	bitcoinLib, err := bitcoind.New(host, port, user, password, USESSL)
	if err != nil {
		log.Fatal(err)
	}

	var waitTime time.Duration = time.Duration(elapsedTime) * time.Second
	var lastSupply float64 = 0

	for {
		bestHash, err := bitcoinLib.GetBestBlockhash()
		if err != nil {
			log.Println(err)
		}

		block, err := bitcoinLib.GetBlock(bestHash)
		if err != nil {
			log.Println(err)
		}

		blockTime := time.Unix(block.Time, 0)

		if time.Now().Sub(blockTime) > waitTime {
			txOutSetInfo, err := bitcoinLib.GetTxOutsetInfo()
			if err != nil {
				log.Println(err)
			}

			if txOutSetInfo.TotalAmount > lastSupply {
				lastSupply = txOutSetInfo.TotalAmount

				err = client.SendSupply(&dia.Supply{
					Symbol:            symbol,
					CirculatingSupply: lastSupply,
				})
				if err != nil {
					log.Println("Error sending supply to API: ", err)
				} else {
					log.Println("Sent supply", lastSupply, "to API")
				}
			}
		} else {
			log.Println("Block:", block.Height, "synchronized for", time.Now().Sub(blockTime))
		}

		time.Sleep(SLEEP_TIME * time.Second)
	}
}
