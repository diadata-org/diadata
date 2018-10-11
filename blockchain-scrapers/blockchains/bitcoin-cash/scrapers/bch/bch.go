package main

import (
	"github.com/diadata-org/diadata/blockchain-scrapers/scrapers"
)

const (
	HOST   = "bitcoin-cash"
	PORT   = 8332
	USER   = "mysecretrpcdiauser"
	PASSWD = "njTcaNX74sSf46_TXacMVlyPMJjuv9i03bqBgj9KQ8E="
	SYMBOL = "BCH"
)

func main() {
	blockchainscrapers.RunScraper(HOST, PORT, USER, PASSWD, SYMBOL)
}
