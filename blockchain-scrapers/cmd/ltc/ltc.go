package main

//docker run -p 9332:9332 --rm uphold/litecoin-core   -printtoconsole   -rpcallowip=::/0  -rpcuser=mysecretrpcdiauser -rpcpassword=njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E=

import (
	"github.com/diadata-org/api-golang/blockchain-scrapers/scrapers"
	"github.com/diadata-org/api-golang/dia"
)

const (
	SERVER_HOST = "litecoind"
	SERVER_PORT = 9332
	USER        = "mysecretrpcdiauser"
	PASSWD      = "njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E="
	SYMBOL      = "LTC"
)

func main() {
	config := dia.GetConfigApi()
	client := dia.NewClient(config)
	b := blockchainscrapers.NewScraper(client, SYMBOL, SERVER_HOST, SERVER_PORT, USER, PASSWD)
	b.Run()
}
