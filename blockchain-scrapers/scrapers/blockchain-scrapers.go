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
	bitcoind *bitcoind.Bitcoind
	client   *dia.Client
	symbol   string
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

func NewScraper(client *dia.Client, symbol string, SERVER_HOST string, SERVER_PORT int, USER string, PASSWD string) *BlockchainScraper {
	bc, err := bitcoind.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Println(err)
		return nil
	}
	t := &BlockchainScraper{
		bitcoind: bc,
		client:   client,
		symbol:   symbol,
	}
	return t
}

func (s *BlockchainScraper) Run() {
	for {
		rinfo, err := s.bitcoind.GetBlockchainInfo()
		if err == nil {
			log.Println("GetBlockchainInfo", rinfo)
			m := time.Unix(rinfo.Mediantime, 0)
			l := time.Now().Sub(m)
			log.Println("lapse time", l)
			if l < 60 {
				err = s.client.SendSupply(&dia.Supply{
					Symbol:            s.symbol,
					CirculatingSupply: numberOfCoinsFor(rinfo.Blocks),
				})
				if err != nil {
					log.Println("Err communicating with api:", err)
				}
			}
		} else {
			log.Println("Err communicating with node:", err)
		}
		time.Sleep(10 * time.Second)
	}
}
