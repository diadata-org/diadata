package main

import (
	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/scrapers"
)

const (
	HOST     = "dogecoin"
	PORT     = 22555
	USER     = "dogediauser"
	PASSWD   = "dogediapassword"
	SYMBOL   = "DOGE"
	TIP_TIME = 90
)

func main() {
	blockchainscrapers.RunScraper(HOST, PORT, USER, PASSWD, SYMBOL, TIP_TIME)
}
