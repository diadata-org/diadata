package main

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/go-bitcoind"
	"github.com/tkanos/gonfig"
	"log"
	"time"
)

const (
	SERVER_HOST       = "bitcoind"
	SERVER_PORT       = 8332
	USER              = "mysecretrpcdiauser"
	PASSWD            = "mysecretrpcdiapassword"
	USESSL            = false
	WALLET_PASSPHRASE = "WalletPassphrase"
	SYMBOL            = "BTC"
)

func numberOfCoinsFor(blockNumber float64) float64 {
	subsidy := 50.0
	totalCoins := 50.0
	var i int64 = 1
	for i < int64(blockNumber) {
		if i%210000 == 0 {
			subsidy = subsidy / 2
		}
		totalCoins += subsidy
		i++
	}
	return totalCoins
}

var api *dia.Client

func init() {
	var c dia.ConfigApi
	configFile := "/run/secrets/api_diadata"
	err := gonfig.GetConf(configFile, &c)
	if err != nil {
		configFile = "../config/secrets/api_diadata.json"
		err = gonfig.GetConf(configFile, &c)
	}
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Loaded secret in", configFile)
	}
	api = dia.NewClient(&c)
	if api == nil {
		log.Fatalln("Couldnt log in to api")
	}
}

func main() {

	bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatalln(err)
	}
	showMessage := true
	last := 0.0
	for {
		rinfo, err := bc.GetBlockchainInfo()
		if err == nil {
			if rinfo.Initialblockdownload && showMessage {
				showMessage = false
				log.Println("Node catching up with the latest block... please wait...")
			}
			if last != rinfo.Blocks && rinfo.Initialblockdownload == false {
				last = rinfo.Blocks
				err = api.SendSupply(&dia.Supply{
					Symbol:            "BTC",
					CirculatingSupply: numberOfCoinsFor(rinfo.Blocks),
				})
				if err != nil {
					log.Println("Err communicating with api:", err)
				}
				last = rinfo.Blocks
			}
		} else {
			log.Println("Err communicating with node:", err)
		}
		time.Sleep(10 * time.Second)
	}
}
