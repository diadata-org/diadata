package blockchainscrapers

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/go-bitcoind"
	"log"
	"time"
)

const (
	USESSL = false
)

type BlockchainScraper struct {
	bitcoind              *bitcoind.Bitcoind
	client                *dia.Client
	symbol                string
	elapsedTime           time.Duration
	lastCirculatingSupply float64
}

func numberOfCoinsFor(blockNumber float64) float64 {
	subsidy := 50.0
	totalCoins := 50.0
	var i int64 = 1
	for i < int64(blockNumber) {
		if i%210000 == 0 {
			subsidy = subsidy / 2
		}
		totalCoins += subsidy
		i++
	}
	return totalCoins
}

func NewScraper(client *dia.Client, symbol string, serverHost string, serverPort int, user string, passwd string, elapsedTime int) *BlockchainScraper {
	bc, err := bitcoind.New(serverHost, serverPort, user, passwd, USESSL)
	if err != nil {
		log.Println(err)
		return nil
	}
	t := &BlockchainScraper{
		bitcoind:    bc,
		client:      client,
		symbol:      symbol,
		elapsedTime: time.Duration(elapsedTime) * time.Second,
	}
	return t
}

func (s *BlockchainScraper) Run() {
	for {
		rinfo, err := s.bitcoind.GetBlockchainInfo()
		if err == nil {
			log.Println("GetBlockchainInfo:", rinfo)
			m := time.Unix(rinfo.Mediantime, 0)
			l := time.Now().Sub(m)
			circulatingSupply := numberOfCoinsFor(rinfo.Blocks)
			log.Println("ElapsedTime block:", l, circulatingSupply, s.lastCirculatingSupply)
			if l < s.elapsedTime && s.lastCirculatingSupply != circulatingSupply {
				err = s.client.SendSupply(&dia.Supply{
					Symbol:            s.symbol,
					CirculatingSupply: circulatingSupply,
				})
				if err != nil {
					log.Println("Err communicating with api:", err)
				} else {
					s.lastCirculatingSupply = circulatingSupply
				}
			}
		} else {
			log.Println("Err communicating with node:", err)
		}
		time.Sleep(10 * time.Second)
	}
}
