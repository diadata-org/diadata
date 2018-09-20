package main

import "github.com/diadata-org/api-golang/blockchain-scrapers/scrapers"

const (
	SERVER_HOST = "bitcoin"
	SERVER_PORT = 8332
	USER        = "mysecretrpcdiauser"
	PASSWD      = "mysecretrpcdiapassword"
	SYMBOL      = "BTC"
)

func main() {
	blockchainscrapers.RunScraper(SERVER_HOST, SERVER_PORT, USER, PASSWD, SYMBOL)
}
