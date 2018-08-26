package main

//docker run -p 11625:11625 --rm stellar/stellar-core   -printtoconsole   -rpcallowip=::/0  -rpcuser=mysecretrpcdiauser -rpcpassword=njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E=

import (
	"github.com/diadata-org/api-golang/blockchain-scrapers/scrapers"
	"github.com/diadata-org/api-golang/dia"
)

const (
	SERVER_HOST = "stellard"
	SERVER_PORT = 11625
	USER        = "mysecretrpcdiauser"
	PASSWD      = "njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E="
	SYMBOL      = "XLM"
	TIP_TIME    = 60 * 20
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
