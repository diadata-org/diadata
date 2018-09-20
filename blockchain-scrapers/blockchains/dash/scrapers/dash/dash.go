package main

import (
	"github.com/diadata-org/api-golang/blockchain-scrapers/scrapers"
)

const (
	HOST 		= "dashd"
	PORT 		= 9998
	USER        = "dashrpc"
	PASSWD 		= "dash_rpc_521d43b"
	SYMBOL      = "DASH"
)

func main() {
	blockchainscrapers.RunScraper(HOST, PORT, USER, PASSWD, SYMBOL)
}
