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
	TIP_TIME    = 60 * 60 * 2
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
	scraper := blockchainscrapers.NewScraper(client, SYMBOL, SERVER_HOST, SERVER_PORT, USER, PASSWD, TIP_TIME)
	if scraper == nil {
		panic("Couldnt load scraper")
	}
	scraper.Run()
}
