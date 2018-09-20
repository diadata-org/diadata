package main

import (
	"github.com/diadata-org/api-golang/blockchain-scrapers/scrapers"
)

const (
	HOST 		= "dogecoin"
	PORT 		= 22555
	USER        = "dogediauser"
	PASSWD      = "dogediapassword"
	SYMBOL      = "DOGE"
)

func main() {
	blockchainscrapers.RunScraper(HOST, PORT, USER, PASSWD, SYMBOL)
}
