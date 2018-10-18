package main

import (
	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
)

const (
	HOST     = "bitcoin-cash"
	PORT     = 8332
	USER     = "mysecretrpcdiauser"
	PASSWD   = "njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E="
	SYMBOL   = "BCH"
	TIP_TIME = 60 * 10
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
	scraper := blockchainscrapers.NewScraper(client, SYMBOL, HOST, PORT, USER, PASSWD, TIP_TIME)
	if scraper == nil {
		panic("Couldnt load scraper")
	}
	scraper.Run()
}