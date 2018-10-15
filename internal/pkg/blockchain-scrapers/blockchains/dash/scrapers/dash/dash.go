package main

import (
	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/scrapers"
)

const (
	HOST   = "dashd"
	PORT   = 9998
	USER   = "dashrpc"
	PASSWD = "dash_rpc_521d43b"
	SYMBOL = "DASH"
	TIP_TIME = 60 * 3
)

func main() {
	blockchainscrapers.RunScraper(HOST, PORT, USER, PASSWD, SYMBOL, TIP_TIME)
}
