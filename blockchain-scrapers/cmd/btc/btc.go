package main

import (
	"github.com/diadata-org/api-golang/blockchain-scrapers/scrapers"
	"github.com/diadata-org/api-golang/dia"
)

const (
	SERVER_HOST = "bitcoind"
	SERVER_PORT = 8332
	USER        = "mysecretrpcdiauser"
	PASSWD      = "mysecretrpcdiapassword"
	SYMBOL      = "BTC"
)

func main() {
	config := dia.GetConfigApi()
	client := dia.NewClient(config)
	b := blockchainscrapers.NewScraper(client, SYMBOL, SERVER_HOST, SERVER_PORT, USER, PASSWD)
	b.Run()
}
