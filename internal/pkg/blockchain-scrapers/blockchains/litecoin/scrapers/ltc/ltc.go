package main

//docker run -p 9332:9332 --rm uphold/litecoin-core   -printtoconsole   -rpcallowip=::/0  -rpcuser=mysecretrpcdiauser -rpcpassword=njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E=

import (
	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	SERVER_HOST = "litecoin"
	SERVER_PORT = 9332
	USER        = "mysecretrpcdiauser"
	PASSWD      = "njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E="
	SYMBOL      = "LTC"
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
